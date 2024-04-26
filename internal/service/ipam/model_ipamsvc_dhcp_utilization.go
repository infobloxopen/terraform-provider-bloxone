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

var IpamsvcDHCPUtilizationResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_free": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_.",
	},
	"dhcp_total": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The total IP addresses available in the DHCP ranges in the scope of this object.",
	},
	"dhcp_used": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The total IP addresses marked as used in the DHCP ranges in the scope of this object.",
	},
	"dhcp_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object.",
	},
}

func ExpandIpamsvcDHCPUtilization(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DHCPUtilization {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcDHCPUtilizationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcDHCPUtilizationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DHCPUtilization {
	if m == nil {
		return nil
	}
	to := &ipam.DHCPUtilization{}
	return to
}

func FlattenIpamsvcDHCPUtilization(ctx context.Context, from *ipam.DHCPUtilization, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPUtilizationAttrTypes)
	}
	m := IpamsvcDHCPUtilizationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPUtilizationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPUtilizationModel) Flatten(ctx context.Context, from *ipam.DHCPUtilization, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPUtilizationModel{}
	}
	m.DhcpFree = flex.FlattenStringPointer(from.DhcpFree)
	m.DhcpTotal = flex.FlattenStringPointer(from.DhcpTotal)
	m.DhcpUsed = flex.FlattenStringPointer(from.DhcpUsed)
	m.DhcpUtilization = flex.FlattenInt64(int64(*from.DhcpUtilization))
}
