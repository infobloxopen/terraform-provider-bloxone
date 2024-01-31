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

type IpamsvcHAGroupHeartbeatsModel struct {
	Peer                types.String `tfsdk:"peer"`
	SuccessfulHeartbeat types.String `tfsdk:"successful_heartbeat"`
}

var IpamsvcHAGroupHeartbeatsAttrTypes = map[string]attr.Type{
	"peer":                 types.StringType,
	"successful_heartbeat": types.StringType,
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
}

func ExpandIpamsvcHAGroupHeartbeats(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHeartbeats {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcHAGroupHeartbeatsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcHAGroupHeartbeatsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHeartbeats {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcHAGroupHeartbeats{
		Peer:                flex.ExpandStringPointer(m.Peer),
		SuccessfulHeartbeat: flex.ExpandStringPointer(m.SuccessfulHeartbeat),
	}
	return to
}

func FlattenIpamsvcHAGroupHeartbeats(ctx context.Context, from *ipam.IpamsvcHAGroupHeartbeats, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupHeartbeatsAttrTypes)
	}
	m := IpamsvcHAGroupHeartbeatsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupHeartbeatsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupHeartbeatsModel) Flatten(ctx context.Context, from *ipam.IpamsvcHAGroupHeartbeats, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupHeartbeatsModel{}
	}
	m.Peer = flex.FlattenStringPointer(from.Peer)
	m.SuccessfulHeartbeat = flex.FlattenStringPointer(from.SuccessfulHeartbeat)
}
