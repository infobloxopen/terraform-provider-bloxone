package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// UseStateForUnknownUnlessDhcpOptionsChange behaves like
// objectplanmodifier.UseStateForUnknown, except that the prior state value is
// not reused when the resource's dhcp_options change in the plan.
//
// The API recomputes inheritance_sources.dhcp_options when dhcp_options change
// (for example, an inherited option that is now set explicitly becomes an
// override), so reusing the old value would make the planned value differ from
// the one returned after apply.
func UseStateForUnknownUnlessDhcpOptionsChange() planmodifier.Object {
	return useStateForUnknownUnlessDhcpOptionsChange{}
}

type useStateForUnknownUnlessDhcpOptionsChange struct{}

func (m useStateForUnknownUnlessDhcpOptionsChange) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change unless dhcp_options change."
}

func (m useStateForUnknownUnlessDhcpOptionsChange) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m useStateForUnknownUnlessDhcpOptionsChange) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// Do nothing on create, on destroy, or if there is a known planned value.
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value, otherwise
	// interpolation gets messed up.
	if req.ConfigValue.IsUnknown() {
		return
	}

	var planOptions, stateOptions attr.Value
	if diags := req.Plan.GetAttribute(ctx, path.Root("dhcp_options"), &planOptions); diags.HasError() {
		return
	}
	if diags := req.State.GetAttribute(ctx, path.Root("dhcp_options"), &stateOptions); diags.HasError() {
		return
	}
	if !planOptions.Equal(stateOptions) {
		return
	}

	resp.PlanValue = req.StateValue
}
