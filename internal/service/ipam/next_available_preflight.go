package ipam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// apiErrorMessage returns the first message from an API error body, or the error text if there is none.
func apiErrorMessage(err error) string {
	if e, ok := err.(interface{ Body() []byte }); ok {
		var body struct {
			Error []struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(e.Body(), &body) == nil && len(body.Error) > 0 && body.Error[0].Message != "" {
			return body.Error[0].Message
		}
	}
	return err.Error()
}

// Availability statuses reported by the next-available preflight check.
const (
	availabilityAvailable   = "Available"
	availabilityPartial     = "Partially available"
	availabilityUnavailable = "Unavailable"
)

// nextAvailableMaxCount is the largest count the next-available API accepts, so at most this many candidates
// per parent and CIDR can be previewed in one plan.
const nextAvailableMaxCount = 20

// preflightPool holds the candidates fetched for one parent and CIDR during a plan. Every instance of a resource
// with count = N is planned separately; the k-th instance takes the k-th candidate. Candidates are fetched in
// batches that double in size, so N instances cost about log2(N) API calls instead of one call per instance.
// Pools are scoped to the API client, which the provider creates anew for every Terraform run, so every plan
// starts with empty pools even when the provider process outlives a single run (as in acceptance tests).
type preflightPool struct {
	sync.Mutex
	candidates []string
	taken      int32
	available  int32  // set once the API reports fewer networks than were asked for; -1 until then
	shortage   string // the API message reporting the shortage
	err        error  // a missing parent or a CIDR that does not fit applies to every instance
	httpRes    *http.Response
}

var preflightState = struct {
	sync.Mutex
	pools   map[string]*preflightPool
	parents map[string]string
}{pools: map[string]*preflightPool{}, parents: map[string]string{}}

func getPreflightPool(scope any, kind, parentId string, cidr int64) *preflightPool {
	preflightState.Lock()
	defer preflightState.Unlock()
	key := fmt.Sprintf("%p|%s|%s|%d", scope, kind, parentId, cidr)
	if preflightState.pools[key] == nil {
		preflightState.pools[key] = &preflightPool{available: -1}
	}
	return preflightState.pools[key]
}

// next reserves the next slot in the pool for this instance. It returns the slot number, the candidate for it if
// there is one, and the number of available networks and the API message if the API reported a shortage.
// The number of available networks is -1 when the slot is beyond what the API can preview.
func (p *preflightPool) next(preview func(count int32) ([]string, *http.Response, error)) (int32, string, int32, string, *http.Response, error) {
	p.Lock()
	defer p.Unlock()
	p.taken++
	k := p.taken
	for {
		if p.err != nil {
			return k, "", 0, "", p.httpRes, p.err
		}
		if k <= int32(len(p.candidates)) {
			return k, p.candidates[k-1], 0, "", nil, nil
		}
		if p.available >= 0 {
			return k, "", p.available, p.shortage, nil, nil
		}
		if len(p.candidates) >= nextAvailableMaxCount {
			return k, "", -1, "", nil, nil
		}
		count := min(max(2*int32(len(p.candidates)), k), nextAvailableMaxCount)
		candidates, httpRes, err := preview(count)
		switch {
		case err == nil && int32(len(candidates)) < count:
			// Fewer results than asked for without an error: treat what came back as all there is.
			p.candidates, p.available = candidates, int32(len(candidates))
			p.shortage = fmt.Sprintf("%d of %d requested networks were returned.", len(candidates), count)
		case err == nil:
			p.candidates = candidates
		case httpRes != nil && httpRes.StatusCode == http.StatusBadRequest && strings.Contains(apiErrorMessage(err), "available networks"):
			available := int32(0)
			if e, ok := err.(interface{ Body() []byte }); ok {
				available = utils.ExtractAvailableCountFromError(e.Body())
			}
			p.shortage = apiErrorMessage(err)
			if available <= int32(len(p.candidates)) {
				p.available = available
				p.candidates = p.candidates[:available]
				continue
			}
			// The batch was larger than the space left; fetch exactly what is available.
			candidates, httpRes, err = preview(available)
			if err != nil {
				p.err, p.httpRes = err, httpRes
				continue
			}
			p.candidates, p.available = candidates, int32(len(candidates))
		default:
			p.err, p.httpRes = err, httpRes
		}
	}
}

// parentLabel returns the parent as "10.0.0.0/16 (ipam/address_block/...)", reading it once per plan.
// If the parent cannot be read, only its resource identifier is returned.
func parentLabel(scope any, parentId string, readParent func() (string, error)) string {
	preflightState.Lock()
	defer preflightState.Unlock()
	key := fmt.Sprintf("%p|%s", scope, parentId)
	if label, ok := preflightState.parents[key]; ok {
		return label
	}
	label := parentId
	if prefix, err := readParent(); err == nil {
		label = fmt.Sprintf("%s (%s)", prefix, parentId)
	}
	preflightState.parents[key] = label
	return label
}

// checkNextAvailable runs at plan time for an object created with next_available_id. The scope is the API client of
// the current Terraform run. It reports one of three statuses:
//   - Available: a candidate exists for this object. Reported as a warning: the candidate is a preview and is not
//     reserved until apply.
//   - Partially available: the parent has some space, but less than the objects of this size planned from it.
//     Reported as a warning, because the check runs again during apply while earlier objects are being created.
//   - Unavailable: the parent is missing, the CIDR does not fit, or no space is left. Reported as a plan error.
func checkNextAvailable(scope any, kind, parentId string, cidr int64, preview func(count int32) ([]string, *http.Response, error), readParent func() (string, error), diags *diag.Diagnostics) {
	k, candidate, available, shortage, httpRes, err := getPreflightPool(scope, kind, parentId, cidr).next(preview)
	if err == nil && candidate != "" {
		diags.AddWarning("Next-available allocation preview",
			fmt.Sprintf("Status:     %s\nCandidate:  %s/%d (not reserved; the final value is assigned at apply)\nParent:     %s",
				availabilityAvailable, candidate, cidr, parentLabel(scope, parentId, readParent)))
		return
	}
	if err == nil && available < 0 {
		diags.AddWarning("Next-available preflight check skipped",
			fmt.Sprintf("Only the first %d /%d %ss planned from one parent can be previewed; the API returns at most %d. "+
				"Availability of this %s was not checked.\nParent:     %s",
				nextAvailableMaxCount, cidr, kind, nextAvailableMaxCount, kind, parentLabel(scope, parentId, readParent)))
		return
	}
	if err == nil && available > 0 {
		diags.AddWarning("Next-available allocation partially available",
			fmt.Sprintf("Status:     %s\nAvailable:  %d of %d requested /%d %ss\nParent:     %s\n\n"+
				"This plan creates more /%d %ss from this parent than it has space for. "+
				"Only %d can be allocated at apply; the others will fail.",
				availabilityPartial, available, k, cidr, kind, parentLabel(scope, parentId, readParent), cidr, kind, available))
		return
	}
	if err == nil {
		diags.AddAttributeError(path.Root("next_available_id"), "No next-available allocation found",
			fmt.Sprintf("Status: %s\nNo /%d %s is available in parent %s.\nAPI: %s\n\n"+
				"Check that:\n- the parent has enough unallocated space\n- the requested cidr is greater than the parent's prefix length\n- the correct parent was selected",
				availabilityUnavailable, cidr, kind, parentLabel(scope, parentId, readParent), shortage))
		return
	}

	msg := apiErrorMessage(err)
	status := 0
	if httpRes != nil {
		status = httpRes.StatusCode
	}
	switch {
	case status == http.StatusNotFound:
		diags.AddAttributeError(path.Root("next_available_id"), "Parent address block not found",
			fmt.Sprintf("Status: %s\nAddress block %q used in next_available_id does not exist.\nAPI: %s",
				availabilityUnavailable, parentId, msg))
	case status == http.StatusBadRequest && strings.Contains(msg, "must be greater than parent"):
		diags.AddAttributeError(path.Root("cidr"), "Requested CIDR does not fit the parent address block",
			fmt.Sprintf("Status: %s\nA /%d %s cannot be allocated from parent %s.\nAPI: %s\n\n"+
				"Check that cidr is greater than the parent's prefix length.",
				availabilityUnavailable, cidr, kind, parentLabel(scope, parentId, readParent), msg))
	default:
		diags.AddWarning("Next-available preflight check skipped",
			fmt.Sprintf("Could not check whether a /%d %s is available in parent %s: %s",
				cidr, kind, parentLabel(scope, parentId, readParent), msg))
	}
}

// nextAvailableCreateError describes a failed next-available allocation at apply time.
func nextAvailableCreateError(kind, parentId string, cidr int64, err error) string {
	return fmt.Sprintf("Unable to allocate the next available /%d %s from address block %q: %s\n\n"+
		"Nothing was created. The space may have been allocated by another process after the plan was made. "+
		"Run terraform plan again to check current availability, then apply.",
		cidr, kind, parentId, apiErrorMessage(err))
}
