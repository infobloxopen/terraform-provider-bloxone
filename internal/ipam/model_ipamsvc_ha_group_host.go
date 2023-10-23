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

type IpamsvcHAGroupHostModel struct {
	Address    types.String `tfsdk:"address"`
	Heartbeats types.List   `tfsdk:"heartbeats"`
	Host       types.String `tfsdk:"host"`
	Port       types.Int64  `tfsdk:"port"`
	Role       types.String `tfsdk:"role"`
	State      types.String `tfsdk:"state"`
}

var IpamsvcHAGroupHostAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"heartbeats": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHAGroupHeartbeatsAttrTypes}},
	"host":       types.StringType,
	"port":       types.Int64Type,
	"role":       types.StringType,
	"state":      types.StringType,
}

var IpamsvcHAGroupHostResourceSchema = schema.Schema{
	MarkdownDescription: `An __HAGroupHost__ object (_dhcp/ha_group_host_) represents an on-prem host belonging to an HA Group.`,
	Attributes:          IpamsvcHAGroupHostResourceSchemaAttributes,
}

var IpamsvcHAGroupHostResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The address on which this host listens.`,
	},
	"heartbeats": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcHAGroupHeartbeatsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Last successful heartbeat received from its peer/s. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.`,
	},
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"port": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The HA port.`,
	},
	"role": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The role of this host in the HA relationship: _active_ or _passive_.`,
	},
	"state": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The state of DHCP on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.`,
	},
}

func expandIpamsvcHAGroupHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHAGroupHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHAGroupHostModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHost {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHAGroupHost{
		Address:    m.Address.ValueStringPointer(),
		Heartbeats: ExpandFrameworkListNestedBlock(ctx, m.Heartbeats, diags, expandIpamsvcHAGroupHeartbeats),
		Host:       m.Host.ValueString(),
		Role:       m.Role.ValueStringPointer(),
		State:      m.State.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcHAGroupHost(ctx context.Context, from *ipam.IpamsvcHAGroupHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupHostAttrTypes)
	}
	m := IpamsvcHAGroupHostModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupHostModel) flatten(ctx context.Context, from *ipam.IpamsvcHAGroupHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupHostModel{}
	}

	m.Address = types.StringPointerValue(from.Address)
	m.Heartbeats = FlattenFrameworkListNestedBlock(ctx, from.Heartbeats, IpamsvcHAGroupHeartbeatsAttrTypes, diags, flattenIpamsvcHAGroupHeartbeats)
	m.Host = types.StringValue(from.Host)
	m.Port = types.Int64Value(int64(*from.Port))
	m.Role = types.StringPointerValue(from.Role)
	m.State = types.StringPointerValue(from.State)

}
