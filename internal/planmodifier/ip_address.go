package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	internaltypes "github.com/infobloxopen/terraform-provider-bloxone/internal/types"
)

func RequiresReplaceIfIPAddressChanged() planmodifier.String {
	return requiresReplaceIfIPAddressChanged{}
}

// requiresReplaceIfIPAddressChanged implements the plan modifier.
type requiresReplaceIfIPAddressChanged struct{}

// Description returns a human-readable description of the plan modifier.
func (m requiresReplaceIfIPAddressChanged) Description(_ context.Context) string {
	return "Requires resource replacement only if the IP address is not semantically equal to the current IP address."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m requiresReplaceIfIPAddressChanged) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifyString implements the plan modification logic.
func (m requiresReplaceIfIPAddressChanged) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing on resource creation or destroy.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	// Do nothing if either value is null or unknown.
	if req.StateValue.IsNull() || req.StateValue.IsUnknown() || req.PlanValue.IsNull() || req.PlanValue.IsUnknown() {
		return
	}

	// Differences in IPv6 letter case or zero compression do not require replacement.
	oldValue := internaltypes.NewIPAddressValue(req.StateValue.ValueString())
	newValue := internaltypes.NewIPAddressValue(req.PlanValue.ValueString())
	equal, _ := oldValue.StringSemanticEquals(ctx, newValue)
	if !equal {
		resp.RequiresReplace = true
	}
}
