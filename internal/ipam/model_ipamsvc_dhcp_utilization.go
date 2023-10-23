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

type IpamsvcDHCPUtilizationModel struct {
	DhcpFree        types.String `tfsdk:"dhcp_free"`
	DhcpTotal       types.String `tfsdk:"dhcp_total"`
	DhcpUsed        types.String `tfsdk:"dhcp_used"`
	DhcpUtilization types.Int64  `tfsdk:"dhcp_utilization"`
}

var IpamsvcDHCPUtilizationAttrTypes = map[string]attr.Type{
	"dhcp_free":        types.StringType,
	"dhcp_total":       types.StringType,
	"dhcp_used":        types.StringType,
	"dhcp_utilization": types.Int64Type,
}

var IpamsvcDHCPUtilizationResourceSchema = schema.Schema{
	MarkdownDescription: `The __DHCPUtilization__ object represents DHCP utilization statistics for an object.`,
	Attributes:          IpamsvcDHCPUtilizationResourceSchemaAttributes,
}

var IpamsvcDHCPUtilizationResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_free": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_.`,
	},
	"dhcp_total": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The total IP addresses available in the DHCP ranges in the scope of this object.`,
	},
	"dhcp_used": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The total IP addresses marked as used in the DHCP ranges in the scope of this object.`,
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object.`,
	},
}

func expandIpamsvcDHCPUtilization(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPUtilization {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPUtilizationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPUtilizationModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPUtilization {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPUtilization{}
	return to
}

func flattenIpamsvcDHCPUtilization(ctx context.Context, from *ipam.IpamsvcDHCPUtilization, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPUtilizationAttrTypes)
	}
	m := IpamsvcDHCPUtilizationModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPUtilizationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPUtilizationModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPUtilization, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPUtilizationModel{}
	}

	m.DhcpFree = types.StringPointerValue(from.DhcpFree)
	m.DhcpTotal = types.StringPointerValue(from.DhcpTotal)
	m.DhcpUsed = types.StringPointerValue(from.DhcpUsed)
	m.DhcpUtilization = types.Int64Value(int64(*from.DhcpUtilization))

}
