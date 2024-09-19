package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcHAGroupHeartbeatsModel struct {
	Peer                  types.String `tfsdk:"peer"`
	SuccessfulHeartbeat   types.String `tfsdk:"successful_heartbeat"`
	SuccessfulHeartbeatV6 types.String `tfsdk:"successful_heartbeat_v6"`
}

var IpamsvcHAGroupHeartbeatsAttrTypes = map[string]attr.Type{
	"peer":                    types.StringType,
	"successful_heartbeat":    types.StringType,
	"successful_heartbeat_v6": types.StringType,
}

var IpamsvcHAGroupHeartbeatsResourceSchemaAttributes = map[string]schema.Attribute{
	"peer": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the peer.",
	},
	"successful_heartbeat": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The timestamp as a string of the last successful heartbeat received from the peer above.",
	},
	"successful_heartbeat_v6": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The timestamp as a string of the last successful DHCPv6 heartbeat received from the peer above.",
	},
}

func FlattenIpamsvcHAGroupHeartbeats(ctx context.Context, from *ipam.HAGroupHeartbeats, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupHeartbeatsAttrTypes)
	}
	m := IpamsvcHAGroupHeartbeatsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupHeartbeatsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupHeartbeatsModel) Flatten(ctx context.Context, from *ipam.HAGroupHeartbeats, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupHeartbeatsModel{}
	}
	m.Peer = flex.FlattenStringPointer(from.Peer)
	m.SuccessfulHeartbeat = flex.FlattenStringPointer(from.SuccessfulHeartbeat)
	m.SuccessfulHeartbeatV6 = flex.FlattenStringPointer(from.SuccessfulHeartbeatV6)
}
