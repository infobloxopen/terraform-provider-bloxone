package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcOptionFilterRuleListModel struct {
	Match types.String `tfsdk:"match"`
	Rules types.List   `tfsdk:"rules"`
}

var IpamsvcOptionFilterRuleListAttrTypes = map[string]attr.Type{
	"match": types.StringType,
	"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionFilterRuleAttrTypes}},
}

var IpamsvcOptionFilterRuleListResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionFilterRuleList__ object (_dhcp/option_filter_rule_list_) represents a collection of DHCP option filter rules that supports matching all or any rules.`,
	Attributes:          IpamsvcOptionFilterRuleListResourceSchemaAttributes,
}

var IpamsvcOptionFilterRuleListResourceSchemaAttributes = map[string]schema.Attribute{
	"match": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if this list should match if any or all rules match (_any_ or _all_).`,
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionFilterRuleResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of child rules.`,
	},
}

func expandIpamsvcOptionFilterRuleList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilterRuleList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionFilterRuleListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionFilterRuleListModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilterRuleList {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionFilterRuleList{
		Match: m.Match.ValueStringPointer(),
		Rules: ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, expandIpamsvcOptionFilterRule),
	}
	return to
}

func flattenIpamsvcOptionFilterRuleList(ctx context.Context, from *ipam.IpamsvcOptionFilterRuleList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionFilterRuleListAttrTypes)
	}
	m := IpamsvcOptionFilterRuleListModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionFilterRuleListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionFilterRuleListModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionFilterRuleList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionFilterRuleListModel{}
	}

	m.Match = types.StringPointerValue(from.Match)
	m.Rules = FlattenFrameworkListNestedBlock(ctx, from.Rules, IpamsvcOptionFilterRuleAttrTypes, diags, flattenIpamsvcOptionFilterRule)

}
