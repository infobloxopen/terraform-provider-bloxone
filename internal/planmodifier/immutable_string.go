package planmodifier

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

func ImmutableString() planmodifier.String {
	return immutableString{}
}

// immutableString implements the plan modifier.
type immutableString struct{}

// Description returns a human-readable description of the plan modifier.
func (m immutableString) Description(_ context.Context) string {
	return "Raises an error if the value is changed after the resource is created."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m immutableString) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

// PlanModifyString implements the plan modification logic.
func (m immutableString) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing on create or destroy.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}

	// Do nothing if the value is not configured, or there is no prior value.
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || req.StateValue.IsNull() {
		return
	}

	if req.ConfigValue.ValueString() != req.StateValue.ValueString() {
		resp.Diagnostics.AddAttributeError(req.Path, fmt.Sprintf("%s cannot be changed", req.Path),
			fmt.Sprintf("%s is immutable after creation (current: %q, configured: %q). Revert the configuration, or destroy and recreate the resource, to change it.",
				req.Path, req.StateValue.ValueString(), req.ConfigValue.ValueString()))
	}
}
