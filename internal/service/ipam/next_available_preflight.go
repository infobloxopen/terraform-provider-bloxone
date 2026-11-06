package ipam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
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

// checkNextAvailable runs at plan time for an object created with next_available_id. It asks the API for one
// candidate and turns a missing parent, a CIDR that does not fit, or exhausted space into a plan error.
// A found candidate is reported as a warning: it is only a preview and is not reserved until apply.
func checkNextAvailable(kind, parentId string, cidr int64, preview func() (string, *http.Response, error), diags *diag.Diagnostics) {
	candidate, httpRes, err := preview()
	if err == nil {
		diags.AddWarning(
			"Next-available allocation preview",
			fmt.Sprintf("Candidate:      %s/%d (not reserved)\nRequested CIDR: /%d %s\nParent:         %s\n\n"+
				"This is a preview only. The %s is allocated at apply and may differ.",
				candidate, cidr, cidr, kind, parentId, kind),
		)
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
			fmt.Sprintf("Address block %q used in next_available_id does not exist: %s", parentId, msg))
	case status == http.StatusBadRequest && strings.Contains(msg, "must be greater than parent"):
		diags.AddAttributeError(path.Root("cidr"), "Requested CIDR does not fit the parent address block",
			fmt.Sprintf("A /%d %s cannot be allocated from address block %q: %s", cidr, kind, parentId, msg))
	case status == http.StatusBadRequest && strings.Contains(msg, "available networks"):
		diags.AddError("No next-available allocation found",
			fmt.Sprintf("No available /%d %s was found in address block %q.\n\n%s\n\n"+
				"Verify that the parent address block has enough unallocated space and that the correct parent was selected.",
				cidr, kind, parentId, msg))
	default:
		diags.AddWarning("Next-available preflight check skipped",
			fmt.Sprintf("Could not check whether a /%d %s is available in address block %q: %s", cidr, kind, parentId, msg))
	}
}

// nextAvailableCreateError describes a failed next-available allocation at apply time.
func nextAvailableCreateError(kind, parentId string, cidr int64, err error) string {
	return fmt.Sprintf("Unable to allocate the next available /%d %s from address block %q: %s\n\n"+
		"Nothing was created. The space may have been allocated by another process after the plan was made. "+
		"Run terraform plan again to check current availability, then apply.",
		cidr, kind, parentId, apiErrorMessage(err))
}
