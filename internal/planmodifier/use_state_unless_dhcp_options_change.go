package planmodifier

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// dhcpOptionAttrs are the option lists whose inherited entries in
// inheritance_sources are rewritten. A list that doesn't exist on the object
// (dhcp_options_v6 on a subnet, for example) is skipped.
var dhcpOptionAttrs = []string{"dhcp_options", "dhcp_options_v6"}

// UseStateForUnknownUnlessDhcpOptionsChange behaves like
// objectplanmodifier.UseStateForUnknown, except when dhcp_options or
// dhcp_options_v6 change. The API then turns an inherited option that the user
// now sets into an "override", so the planned inheritance_sources entry for it
// (under the matching list) is changed from "inherit" to "override". If the
// result can't be worked out (an option was removed, or a code isn't known
// yet) the value stays unknown.
func UseStateForUnknownUnlessDhcpOptionsChange() planmodifier.Object {
	return useStateForUnknownUnlessDhcpOptionsChange{}
}

type useStateForUnknownUnlessDhcpOptionsChange struct{}

func (useStateForUnknownUnlessDhcpOptionsChange) Description(_ context.Context) string {
	return "Uses the state value, with inherited dhcp_options and dhcp_options_v6 that are now set explicitly planned as overrides."
}

func (m useStateForUnknownUnlessDhcpOptionsChange) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (useStateForUnknownUnlessDhcpOptionsChange) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	if req.State.Raw.IsNull() || req.Plan.Raw.IsNull() || !req.PlanValue.IsUnknown() || req.ConfigValue.IsUnknown() {
		return
	}

	// option codes the user sets, for each list that is changing
	changed := map[string]map[string]bool{}
	for _, name := range dhcpOptionAttrs {
		var plan, state, config attr.Value
		req.Plan.GetAttribute(ctx, path.Root(name), &plan)
		req.State.GetAttribute(ctx, path.Root(name), &state)
		req.Config.GetAttribute(ctx, path.Root(name), &config)
		if plan == nil || config == nil || plan.Equal(state) {
			continue
		}
		codes, ok := configuredOptionCodes(ctx, config)
		if !ok {
			return
		}
		changed[name] = codes
	}
	if len(changed) == 0 {
		resp.PlanValue = req.StateValue
		return
	}

	// entries of inheritance_sources.<list>.value: inherit -> override
	unknown := false
	src, _ := req.StateValue.ToTerraformValue(ctx)
	out, err := tftypes.Transform(src, func(p *tftypes.AttributePath, v tftypes.Value) (tftypes.Value, error) {
		steps := p.Steps()
		if len(steps) != 3 || steps[1] != tftypes.AttributeName("value") {
			return v, nil
		}
		name, _ := steps[0].(tftypes.AttributeName)
		codes, ok := changed[string(name)]
		var entry map[string]tftypes.Value
		if !ok || v.As(&entry) != nil {
			return v, nil
		}
		code, _, _ := tftypes.WalkAttributePath(v, tftypes.NewAttributePath().WithAttributeName("value").WithAttributeName("option").WithAttributeName("option_code"))
		var action, c string
		_ = entry["action"].As(&action) // left empty if unreadable, so the entry is untouched
		if cv, ok := code.(tftypes.Value); ok {
			_ = cv.As(&c)
		}
		if c == "" { // no usable option code in state, leave the entry untouched
			return v, nil
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

// configuredOptionCodes returns the option codes set in a dhcp_options style
// list. It returns false if the list isn't fully known yet.
func configuredOptionCodes(ctx context.Context, config attr.Value) (map[string]bool, bool) {
	codes := map[string]bool{}
	cfg, _ := config.ToTerraformValue(ctx)
	var opts []tftypes.Value
	if !cfg.IsFullyKnown() || cfg.As(&opts) != nil {
		return nil, false
	}
	for _, o := range opts {
		code, _, _ := tftypes.WalkAttributePath(o, tftypes.NewAttributePath().WithAttributeName("option_code"))
		var s string
		// skip null/empty option codes so they can't match anything
		if c, ok := code.(tftypes.Value); ok && c.As(&s) == nil && s != "" {
			codes[s] = true
		}
	}
	return codes, true
}
