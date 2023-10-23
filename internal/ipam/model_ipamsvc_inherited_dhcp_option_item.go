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

type IpamsvcInheritedDHCPOptionItemModel struct {
	Option          types.Object `tfsdk:"option"`
	OverridingGroup types.String `tfsdk:"overriding_group"`
}

var IpamsvcInheritedDHCPOptionItemAttrTypes = map[string]attr.Type{
	"option":           types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes},
	"overriding_group": types.StringType,
}

var IpamsvcInheritedDHCPOptionItemResourceSchema = schema.Schema{
	MarkdownDescription: `A wrapper of item (_dhcp/option_item_) in a list of Inherited DHCP options. It contains extra fields not covered by OptionItem.`,
	Attributes:          IpamsvcInheritedDHCPOptionItemResourceSchemaAttributes,
}

var IpamsvcInheritedDHCPOptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"option": schema.SingleNestedAttribute{
		Attributes:          IpamsvcOptionItemResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"overriding_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcInheritedDHCPOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedDHCPOptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedDHCPOptionItemModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOptionItem {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedDHCPOptionItem{
		Option:          expandIpamsvcOptionItem(ctx, m.Option, diags),
		OverridingGroup: m.OverridingGroup.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcInheritedDHCPOptionItem(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDHCPOptionItemAttrTypes)
	}
	m := IpamsvcInheritedDHCPOptionItemModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDHCPOptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDHCPOptionItemModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOptionItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDHCPOptionItemModel{}
	}

	m.Option = flattenIpamsvcOptionItem(ctx, from.Option, diags)
	m.OverridingGroup = types.StringPointerValue(from.OverridingGroup)

}
