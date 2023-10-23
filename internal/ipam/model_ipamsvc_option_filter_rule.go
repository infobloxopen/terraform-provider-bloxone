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

type IpamsvcOptionFilterRuleModel struct {
	Compare         types.String `tfsdk:"compare"`
	OptionCode      types.String `tfsdk:"option_code"`
	OptionValue     types.String `tfsdk:"option_value"`
	SubstringOffset types.Int64  `tfsdk:"substring_offset"`
}

var IpamsvcOptionFilterRuleAttrTypes = map[string]attr.Type{
	"compare":          types.StringType,
	"option_code":      types.StringType,
	"option_value":     types.StringType,
	"substring_offset": types.Int64Type,
}

var IpamsvcOptionFilterRuleResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionFilterRule__ object (_dhcp/option_filter_rule_) represents a filter rule to match a DHCP client.`,
	Attributes:          IpamsvcOptionFilterRuleResourceSchemaAttributes,
}

var IpamsvcOptionFilterRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"compare": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Indicates how to compare the _option_value_ to the client option.  Success by comparison:  * _equals_: value and client option are the same,  * _not_equals_: value and client option are not the same,  * _exists_: client option exists,  * _not_exists_: client option does not exist,  * _text_substring_: value is the specified substring of the option,  * _not_text_substring_: value is not the specified substring of the option.  * _hex_substring_: value is the specified hexadecimal substring of the option,  * _not_hex_substring_: value is not the specified hexadecimal substring of the option.`,
	},
	"option_code": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"option_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The value to match against.`,
	},
	"substring_offset": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The offset where the substring match starts. This is used only if comparing the _option_value_ using any of the substring modes.`,
	},
}

func expandIpamsvcOptionFilterRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilterRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionFilterRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionFilterRuleModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilterRule {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionFilterRule{
		Compare:         m.Compare.ValueString(),
		OptionCode:      m.OptionCode.ValueString(),
		OptionValue:     m.OptionValue.ValueStringPointer(),
		SubstringOffset: ptr(int64(m.SubstringOffset.ValueInt64())),
	}
	return to
}

func flattenIpamsvcOptionFilterRule(ctx context.Context, from *ipam.IpamsvcOptionFilterRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionFilterRuleAttrTypes)
	}
	m := IpamsvcOptionFilterRuleModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionFilterRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionFilterRuleModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionFilterRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionFilterRuleModel{}
	}

	m.Compare = types.StringValue(from.Compare)
	m.OptionCode = types.StringValue(from.OptionCode)
	m.OptionValue = types.StringPointerValue(from.OptionValue)
	m.SubstringOffset = types.Int64Value(int64(*from.SubstringOffset))

}
