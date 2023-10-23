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

type IpamsvcOptionItemModel struct {
	Group       types.String `tfsdk:"group"`
	OptionCode  types.String `tfsdk:"option_code"`
	OptionValue types.String `tfsdk:"option_value"`
	Type        types.String `tfsdk:"type"`
}

var IpamsvcOptionItemAttrTypes = map[string]attr.Type{
	"group":        types.StringType,
	"option_code":  types.StringType,
	"option_value": types.StringType,
	"type":         types.StringType,
}

var IpamsvcOptionItemResourceSchema = schema.Schema{
	MarkdownDescription: `An item (_dhcp/option_item_) in a list of DHCP options. May be either a specific option or a group of options.`,
	Attributes:          IpamsvcOptionItemResourceSchemaAttributes,
}

var IpamsvcOptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"option_code": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"option_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The option value.`,
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The type of item.  Valid values are: * _group_ * _option_`,
	},
}

func expandIpamsvcOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionItemModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionItem {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionItem{
		Group:       m.Group.ValueStringPointer(),
		OptionCode:  m.OptionCode.ValueStringPointer(),
		OptionValue: m.OptionValue.ValueStringPointer(),
		Type:        m.Type.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcOptionItem(ctx context.Context, from *ipam.IpamsvcOptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionItemAttrTypes)
	}
	m := IpamsvcOptionItemModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionItemModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionItemModel{}
	}

	m.Group = types.StringPointerValue(from.Group)
	m.OptionCode = types.StringPointerValue(from.OptionCode)
	m.OptionValue = types.StringPointerValue(from.OptionValue)
	m.Type = types.StringPointerValue(from.Type)

}
