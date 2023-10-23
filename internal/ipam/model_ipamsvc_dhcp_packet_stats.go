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

type IpamsvcDHCPPacketStatsModel struct {
	DhcpPktReceived   types.String `tfsdk:"dhcp_pkt_received"`
	DhcpPktReceivedV6 types.String `tfsdk:"dhcp_pkt_received_v6"`
	DhcpPktSent       types.String `tfsdk:"dhcp_pkt_sent"`
	DhcpPktSentV6     types.String `tfsdk:"dhcp_pkt_sent_v6"`
	DhcpReqReceived   types.String `tfsdk:"dhcp_req_received"`
	DhcpReqReceivedV6 types.String `tfsdk:"dhcp_req_received_v6"`
}

var IpamsvcDHCPPacketStatsAttrTypes = map[string]attr.Type{
	"dhcp_pkt_received":    types.StringType,
	"dhcp_pkt_received_v6": types.StringType,
	"dhcp_pkt_sent":        types.StringType,
	"dhcp_pkt_sent_v6":     types.StringType,
	"dhcp_req_received":    types.StringType,
	"dhcp_req_received_v6": types.StringType,
}

var IpamsvcDHCPPacketStatsResourceSchema = schema.Schema{
	MarkdownDescription: `The DHCPPacketStats object represents DHCP packets statistics for a DHCP __Host__.`,
	Attributes:          IpamsvcDHCPPacketStatsResourceSchemaAttributes,
}

var IpamsvcDHCPPacketStatsResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_pkt_received": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP packets received.`,
	},
	"dhcp_pkt_received_v6": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP V6 packets received.`,
	},
	"dhcp_pkt_sent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP packets sent.`,
	},
	"dhcp_pkt_sent_v6": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP V6 packets sent.`,
	},
	"dhcp_req_received": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP requests received.`,
	},
	"dhcp_req_received_v6": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of DHCP V6 requests received.`,
	},
}

func expandIpamsvcDHCPPacketStats(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPPacketStats {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPPacketStatsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPPacketStatsModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPPacketStats {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPPacketStats{}
	return to
}

func flattenIpamsvcDHCPPacketStats(ctx context.Context, from *ipam.IpamsvcDHCPPacketStats, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPPacketStatsAttrTypes)
	}
	m := IpamsvcDHCPPacketStatsModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPPacketStatsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPPacketStatsModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPPacketStats, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPPacketStatsModel{}
	}

	m.DhcpPktReceived = types.StringPointerValue(from.DhcpPktReceived)
	m.DhcpPktReceivedV6 = types.StringPointerValue(from.DhcpPktReceivedV6)
	m.DhcpPktSent = types.StringPointerValue(from.DhcpPktSent)
	m.DhcpPktSentV6 = types.StringPointerValue(from.DhcpPktSentV6)
	m.DhcpReqReceived = types.StringPointerValue(from.DhcpReqReceived)
	m.DhcpReqReceivedV6 = types.StringPointerValue(from.DhcpReqReceivedV6)

}
