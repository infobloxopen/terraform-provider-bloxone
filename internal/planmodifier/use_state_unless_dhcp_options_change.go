package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// UseStateForUnknownUnlessDhcpOptionsChange behaves like
// objectplanmodifier.UseStateForUnknown, except when dhcp_options change. The
// API then turns an inherited option that the user now sets into an
// "override", so the planned inheritance_sources.dhcp_options entry for it is
// changed from "inherit" to "override". If the result can't be worked out
// (an option was removed, or a code isn't known yet) the value stays unknown.
func UseStateForUnknownUnlessDhcpOptionsChange() planmodifier.Object {
	return useStateForUnknownUnlessDhcpOptionsChange{}
}

type useStateForUnknownUnlessDhcpOptionsChange struct{}

func (useStateForUnknownUnlessDhcpOptionsChange) Description(_ context.Context) string {
	return "Uses the state value, with inherited dhcp_options that are now set explicitly planned as overrides."
}

func (m useStateForUnknownUnlessDhcpOptionsChange) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (useStateForUnknownUnlessDhcpOptionsChange) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}

	var plan, state, config attr.Value
	req.Plan.GetAttribute(ctx, path.Root("dhcp_options"), &plan)
	req.State.GetAttribute(ctx, path.Root("dhcp_options"), &state)
	req.Config.GetAttribute(ctx, path.Root("dhcp_options"), &config)
	if plan == nil || config == nil || plan.Equal(state) {
		resp.PlanValue = req.StateValue
		return
	}

	// option codes the user sets
	codes := map[string]bool{}
	cfg, _ := config.ToTerraformValue(ctx)
	var opts []tftypes.Value
	if !cfg.IsFullyKnown() || cfg.As(&opts) != nil {
		return
	}
	for _, o := range opts {
		code, _, _ := tftypes.WalkAttributePath(o, tftypes.NewAttributePath().WithAttributeName("option_code"))
		var s string
		if c, ok := code.(tftypes.Value); ok && c.As(&s) == nil {
			codes[s] = true
		}
	}

	// entries of inheritance_sources.dhcp_options.value: inherit -> override
	unknown := false
	src, _ := req.StateValue.ToTerraformValue(ctx)
	out, err := tftypes.Transform(src, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		steps := p.Steps()
		var entry map[string]tftypes.Value
		if len(steps) != 3 || steps[0] != tftypes.AttributeName("dhcp_options") || steps[1] != tftypes.AttributeName("value") || v.As(&entry) != nil {
			return v, nil
		}
		code, _, _ := tftypes.WalkAttributePath(v, tftypes.NewAttributePath().WithAttributeName("value").WithAttributeName("option").WithAttributeName("option_code"))
		var action, c string
		entry["action"].As(&action)
		if cv, ok := code.(tftypes.Value); ok {
			cv.As(&c)
		}
		switch {
		case action == "inherit" && codes[c]:
			entry["action"] = tftypes.NewValue(tftypes.String, "override")
			return tftypes.NewValue(v.Type(), entry), nil
		case action == "override" && !codes[c]: // option removed: API may switch it back
			unknown = true
		}
		return v, nil
	})
	if err != nil || unknown {
		return
	}
	if val, err := req.StateValue.Type(ctx).ValueFromTerraform(ctx, out); err == nil {
		if obj, ok := val.(basetypes.ObjectValue); ok {
			resp.PlanValue = obj
		}
	}
}
