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

type IpamsvcLeaseSubnetModel struct {
	Id types.String `tfsdk:"id"`
}

var IpamsvcLeaseSubnetAttrTypes = map[string]attr.Type{
	"id": types.StringType,
}

var IpamsvcLeaseSubnetResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcLeaseSubnetResourceSchemaAttributes,
}

var IpamsvcLeaseSubnetResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcLeaseSubnet(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcLeaseSubnet {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcLeaseSubnetModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcLeaseSubnetModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcLeaseSubnet {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcLeaseSubnet{}
	return to
}

func flattenIpamsvcLeaseSubnet(ctx context.Context, from *ipam.IpamsvcLeaseSubnet, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcLeaseSubnetAttrTypes)
	}
	m := IpamsvcLeaseSubnetModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcLeaseSubnetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcLeaseSubnetModel) flatten(ctx context.Context, from *ipam.IpamsvcLeaseSubnet, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcLeaseSubnetModel{}
	}

	m.Id = types.StringPointerValue(from.Id)

}
