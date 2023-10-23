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

var IpamsvcDHCPOptionsInheritanceResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration that specifies how the _dhcp_options_ field is inherited from the parent object.`,
	Attributes:          IpamsvcDHCPOptionsInheritanceResourceSchemaAttributes,
}

var IpamsvcDHCPOptionsInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcDHCPOptionsInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPOptionsInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPOptionsInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPOptionsInheritanceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPOptionsInheritance {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPOptionsInheritance{
		DhcpOptions: expandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
	}
	return to
}

func flattenIpamsvcDHCPOptionsInheritance(ctx context.Context, from *ipam.IpamsvcDHCPOptionsInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPOptionsInheritanceAttrTypes)
	}
	m := IpamsvcDHCPOptionsInheritanceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPOptionsInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPOptionsInheritanceModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPOptionsInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPOptionsInheritanceModel{}
	}

	m.DhcpOptions = flattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)

}
