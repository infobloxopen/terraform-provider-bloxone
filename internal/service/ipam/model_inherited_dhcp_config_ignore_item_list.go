package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type InheritedDHCPConfigIgnoreItemListModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

var InheritedDHCPConfigIgnoreItemListAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcIgnoreItemAttrTypes}},
}

var InheritedDHCPConfigIgnoreItemListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcIgnoreItemResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

func ExpandInheritedDHCPConfigIgnoreItemList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPConfigIgnoreItemList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritedDHCPConfigIgnoreItemListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritedDHCPConfigIgnoreItemListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPConfigIgnoreItemList {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDHCPConfigIgnoreItemList{}
	return to
}

func FlattenInheritedDHCPConfigIgnoreItemList(ctx context.Context, from *ipam.InheritedDHCPConfigIgnoreItemList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPConfigIgnoreItemListAttrTypes)
	}
	m := InheritedDHCPConfigIgnoreItemListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPConfigIgnoreItemListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDHCPConfigIgnoreItemListModel) Flatten(ctx context.Context, from *ipam.InheritedDHCPConfigIgnoreItemList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDHCPConfigIgnoreItemListModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, IpamsvcIgnoreItemAttrTypes, diags, FlattenIpamsvcIgnoreItem)
}
