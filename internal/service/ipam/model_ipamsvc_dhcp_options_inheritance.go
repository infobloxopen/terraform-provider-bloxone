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

type IpamsvcDHCPOptionsInheritanceModel struct {
	DhcpOptions types.Object `tfsdk:"dhcp_options"`
}

var IpamsvcDHCPOptionsInheritanceAttrTypes = map[string]attr.Type{
	"dhcp_options": types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionListAttrTypes},
}

var IpamsvcDHCPOptionsInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandIpamsvcDHCPOptionsInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DHCPOptionsInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcDHCPOptionsInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcDHCPOptionsInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DHCPOptionsInheritance {
	if m == nil {
		return nil
	}
	to := &ipam.DHCPOptionsInheritance{
		DhcpOptions: ExpandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
	}
	return to
}

func FlattenIpamsvcDHCPOptionsInheritance(ctx context.Context, from *ipam.DHCPOptionsInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPOptionsInheritanceAttrTypes)
	}
	m := IpamsvcDHCPOptionsInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPOptionsInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPOptionsInheritanceModel) Flatten(ctx context.Context, from *ipam.DHCPOptionsInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPOptionsInheritanceModel{}
	}
	m.DhcpOptions = FlattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
}
