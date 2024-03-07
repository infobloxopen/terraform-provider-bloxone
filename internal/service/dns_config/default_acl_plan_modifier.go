package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func UseDefaultAclForNull() planmodifier.List {
	return useDefaultAclForNull{}
}

// useDefaultAclForNull implements the plan modifier.
type useDefaultAclForNull struct{}

// Description returns a human-readable description of the plan modifier.
func (m useDefaultAclForNull) Description(_ context.Context) string {
	return "Sets the default value for the acl attribute in the plan if the value is null."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m useDefaultAclForNull) MarkdownDescription(_ context.Context) string {
	return "Sets the default value for the acl attribute in the plan if the value is null."
}

// PlanModifyList implements the plan modification logic.
func (m useDefaultAclForNull) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Do nothing if there is no state value.
	if req.StateValue.IsNull() {
		return
	}

	// If the config value is unknown or null, use an empty string.
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		obj, diags := types.ObjectValue(ConfigACLItemAttrTypes, map[string]attr.Value{
			"access":   types.StringValue("allow"),
			"acl":      types.StringNull(),
			"address":  types.StringValue(""),
			"element":  types.StringValue("any"),
			"tsig_key": types.ObjectNull(ConfigTSIGKeyAttrTypes),
		})
		if resp.Diagnostics = append(resp.Diagnostics, diags...); resp.Diagnostics.HasError() {
			return
		}

		resp.PlanValue, diags = types.ListValue(types.ObjectType{
			AttrTypes: ConfigACLItemAttrTypes,
		}, []attr.Value{
			obj,
		})
		if resp.Diagnostics = append(resp.Diagnostics, diags...); resp.Diagnostics.HasError() {
			return
		}

	}
}
