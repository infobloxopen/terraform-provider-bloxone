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

type IpamsvcInheritedDHCPOptionItemModel struct {
	Option          types.Object `tfsdk:"option"`
	OverridingGroup types.String `tfsdk:"overriding_group"`
}

var IpamsvcInheritedDHCPOptionItemAttrTypes = map[string]attr.Type{
	"option":           types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes},
	"overriding_group": types.StringType,
}

var IpamsvcInheritedDHCPOptionItemResourceSchemaAttributes = map[string]schema.Attribute{
	"option": schema.SingleNestedAttribute{
		Attributes:          IpamsvcOptionItemResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Option inherited from the ancestor.",
	},
	"overriding_group": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandIpamsvcInheritedDHCPOptionItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcInheritedDHCPOptionItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcInheritedDHCPOptionItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPOptionItem {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDHCPOptionItem{
		Option:          ExpandIpamsvcOptionItem(ctx, m.Option, diags),
		OverridingGroup: m.OverridingGroup.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcInheritedDHCPOptionItem(ctx context.Context, from *ipam.InheritedDHCPOptionItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDHCPOptionItemAttrTypes)
	}
	m := IpamsvcInheritedDHCPOptionItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDHCPOptionItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDHCPOptionItemModel) Flatten(ctx context.Context, from *ipam.InheritedDHCPOptionItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDHCPOptionItemModel{}
	}
	m.Option = FlattenIpamsvcOptionItem(ctx, from.Option, diags)
	m.OverridingGroup = flex.FlattenStringPointer(from.OverridingGroup)
}
