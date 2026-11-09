package planmodifier

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// dhcpOptionAttrs are the option lists in inheritance_sources whose value is
// left unknown. A list that doesn't exist on the object (dhcp_options_v6 on a
// subnet, for example) is skipped.
var dhcpOptionAttrs = []string{"dhcp_options", "dhcp_options_v6"}

// UseStateForUnknownUnlessDhcpOptionsChange behaves like
// objectplanmodifier.UseStateForUnknown, except for the inherited option lists
// (inheritance_sources.dhcp_options.value and dhcp_options_v6.value), which are
// left unknown on update. Their entries come from the object's own options and
// from the parent, and the parent may change in the same apply, so the state
// value can't be trusted. Leaving them unknown also keeps stale "inherit"
// entries out of the update request.
func UseStateForUnknownUnlessDhcpOptionsChange() planmodifier.Object {
	return useStateForUnknownUnlessDhcpOptionsChange{}
}

type useStateForUnknownUnlessDhcpOptionsChange struct{}

func (useStateForUnknownUnlessDhcpOptionsChange) Description(_ context.Context) string {
	return "Uses the state value, with the inherited dhcp_options and dhcp_options_v6 lists left unknown on update."
}

func (m useStateForUnknownUnlessDhcpOptionsChange) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (useStateForUnknownUnlessDhcpOptionsChange) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}

	// inheritance_sources.<list>.value: unknown, the rest from state
	src, _ := req.StateValue.ToTerraformValue(ctx)
	out, err := tftypes.Transform(src, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		steps := p.Steps()
		if len(steps) != 2 || steps[1] != tftypes.AttributeName("value") {
			return v, nil
		}
		if name, _ := steps[0].(tftypes.AttributeName); slices.Contains(dhcpOptionAttrs, string(name)) {
			return tftypes.NewValue(v.Type(), tftypes.UnknownValue), nil
		}
		return v, nil
	})
	if err != nil {
		return
	}
	if val, err := req.StateValue.Type(ctx).ValueFromTerraform(ctx, out); err == nil {
		if obj, ok := val.(basetypes.ObjectValue); ok {
			resp.PlanValue = obj
		}
	}
}
