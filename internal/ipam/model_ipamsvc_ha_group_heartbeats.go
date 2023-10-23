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

type IpamsvcHAGroupHeartbeatsModel struct {
	Peer                types.String `tfsdk:"peer"`
	SuccessfulHeartbeat types.String `tfsdk:"successful_heartbeat"`
}

var IpamsvcHAGroupHeartbeatsAttrTypes = map[string]attr.Type{
	"peer":                 types.StringType,
	"successful_heartbeat": types.StringType,
}

var IpamsvcHAGroupHeartbeatsResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcHAGroupHeartbeatsResourceSchemaAttributes,
}

var IpamsvcHAGroupHeartbeatsResourceSchemaAttributes = map[string]schema.Attribute{
	"peer": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the peer.`,
	},
	"successful_heartbeat": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The timestamp as a string of the last successful heartbeat received from the peer above.`,
	},
}

func expandIpamsvcHAGroupHeartbeats(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHeartbeats {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHAGroupHeartbeatsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHAGroupHeartbeatsModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHeartbeats {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHAGroupHeartbeats{
		Peer:                m.Peer.ValueStringPointer(),
		SuccessfulHeartbeat: m.SuccessfulHeartbeat.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcHAGroupHeartbeats(ctx context.Context, from *ipam.IpamsvcHAGroupHeartbeats, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupHeartbeatsAttrTypes)
	}
	m := IpamsvcHAGroupHeartbeatsModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupHeartbeatsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupHeartbeatsModel) flatten(ctx context.Context, from *ipam.IpamsvcHAGroupHeartbeats, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupHeartbeatsModel{}
	}

	m.Peer = types.StringPointerValue(from.Peer)
	m.SuccessfulHeartbeat = types.StringPointerValue(from.SuccessfulHeartbeat)

}
