package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcDHCPInfoModel struct {
	ClientHostname    types.String      `tfsdk:"client_hostname"`
	ClientHwaddr      types.String      `tfsdk:"client_hwaddr"`
	ClientId          types.String      `tfsdk:"client_id"`
	End               timetypes.RFC3339 `tfsdk:"end"`
	Fingerprint       types.String      `tfsdk:"fingerprint"`
	Iaid              types.Int64       `tfsdk:"iaid"`
	LeaseType         types.String      `tfsdk:"lease_type"`
	PreferredLifetime timetypes.RFC3339 `tfsdk:"preferred_lifetime"`
	Remain            types.Int64       `tfsdk:"remain"`
	Start             timetypes.RFC3339 `tfsdk:"start"`
	State             types.String      `tfsdk:"state"`
	StateTs           timetypes.RFC3339 `tfsdk:"state_ts"`
}

var IpamsvcDHCPInfoAttrTypes = map[string]attr.Type{
	"client_hostname":    types.StringType,
	"client_hwaddr":      types.StringType,
	"client_id":          types.StringType,
	"end":                timetypes.RFC3339Type{},
	"fingerprint":        types.StringType,
	"iaid":               types.Int64Type,
	"lease_type":         types.StringType,
	"preferred_lifetime": timetypes.RFC3339Type{},
	"remain":             types.Int64Type,
	"start":              timetypes.RFC3339Type{},
	"state":              types.StringType,
	"state_ts":           timetypes.RFC3339Type{},
}

var IpamsvcDHCPInfoResourceSchema = schema.Schema{
	MarkdownDescription: `The __DHCPInfo__ object represents the DHCP information associated with an address object.`,
	Attributes:          IpamsvcDHCPInfoResourceSchemaAttributes,
}

var IpamsvcDHCPInfoResourceSchemaAttributes = map[string]schema.Attribute{
	"client_hostname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The DHCP host name associated with this client.`,
	},
	"client_hwaddr": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The hardware address associated with this client.`,
	},
	"client_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The ID associated with this client.`,
	},
	"end": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The timestamp at which the _state_, when set to _leased_, will be changed to _free_.`,
	},
	"fingerprint": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The DHCP fingerprint for the associated lease.`,
	},
	"iaid": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `Identity Association Identifier (IAID) of the lease. Applicable only for DHCPv6.`,
	},
	"lease_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Lease type. Applicable only for address under DHCP control. The value can be empty for address not under DHCP control.  Valid values are: * _DHCPv6NonTemporaryAddress_: DHCPv6 non-temporary address (NA) * _DHCPv6TemporaryAddress_: DHCPv6 temporary address (TA) * _DHCPv6PrefixDelegation_: DHCPv6 prefix delegation (PD) * _DHCPv4_: DHCPv4 lease`,
	},
	"preferred_lifetime": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The length of time that a valid address is preferred (i.e., the time until deprecation). When the preferred lifetime expires, the address becomes deprecated on the client. It is still considered leased on the server. Applicable only for DHCPv6.`,
	},
	"remain": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state.`,
	},
	"start": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The timestamp at which _state_ was first set to _leased_.`,
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Indicates the status of this IP address from a DHCP protocol standpoint as:   * _none_: Address is not under DHCP control.   * _free_: Address is under DHCP control but has no lease currently assigned.   * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource.`,
	},
	"state_ts": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The timestamp at which the _state_ was last reported.`,
	},
}

func expandIpamsvcDHCPInfo(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPInfo {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPInfoModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPInfoModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPInfo {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPInfo{}
	return to
}

func flattenIpamsvcDHCPInfo(ctx context.Context, from *ipam.IpamsvcDHCPInfo, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPInfoAttrTypes)
	}
	m := IpamsvcDHCPInfoModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPInfoAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPInfoModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPInfo, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPInfoModel{}
	}

	m.ClientHostname = types.StringPointerValue(from.ClientHostname)
	m.ClientHwaddr = types.StringPointerValue(from.ClientHwaddr)
	m.ClientId = types.StringPointerValue(from.ClientId)
	m.End = timetypes.NewRFC3339TimePointerValue(from.End)
	m.Fingerprint = types.StringPointerValue(from.Fingerprint)
	m.Iaid = types.Int64Value(int64(*from.Iaid))
	m.LeaseType = types.StringPointerValue(from.LeaseType)
	m.PreferredLifetime = timetypes.NewRFC3339TimePointerValue(from.PreferredLifetime)
	m.Remain = types.Int64Value(int64(*from.Remain))
	m.Start = timetypes.NewRFC3339TimePointerValue(from.Start)
	m.State = types.StringPointerValue(from.State)
	m.StateTs = timetypes.NewRFC3339TimePointerValue(from.StateTs)

}
