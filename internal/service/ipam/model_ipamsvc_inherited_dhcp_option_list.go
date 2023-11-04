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

type IpamsvcInheritedDHCPOptionListModel struct {
	Action types.String `tfsdk:"action"`
	Value  types.List   `tfsdk:"value"`
}

var IpamsvcInheritedDHCPOptionListAttrTypes = map[string]attr.Type{
	"action": types.StringType,
	"value":  types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionAttrTypes}},
}

var IpamsvcInheritedDHCPOptionListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don't use the inherited value.  Defaults to _inherit_.`,
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcInheritedDHCPOptionResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The inherited DHCP option values.`,
	},
}

func ExpandIpamsvcInheritedDHCPOptionList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOptionList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcInheritedDHCPOptionListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcInheritedDHCPOptionListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOptionList {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcInheritedDHCPOptionList{
		Action: m.Action.ValueStringPointer(),
		Value:  flex.ExpandFrameworkListNestedBlock(ctx, m.Value, diags, ExpandIpamsvcInheritedDHCPOption),
	}
	return to
}

func FlattenIpamsvcInheritedDHCPOptionList(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOptionList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDHCPOptionListAttrTypes)
	}
	m := IpamsvcInheritedDHCPOptionListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDHCPOptionListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDHCPOptionListModel) Flatten(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOptionList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDHCPOptionListModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, IpamsvcInheritedDHCPOptionAttrTypes, diags, FlattenIpamsvcInheritedDHCPOption)
}
