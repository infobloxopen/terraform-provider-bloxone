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

type IpamsvcDHCPConfigModel struct {
	AbandonedReclaimTime   types.Int64 `tfsdk:"abandoned_reclaim_time"`
	AbandonedReclaimTimeV6 types.Int64 `tfsdk:"abandoned_reclaim_time_v6"`
	AllowUnknown           types.Bool  `tfsdk:"allow_unknown"`
	AllowUnknownV6         types.Bool  `tfsdk:"allow_unknown_v6"`
	Filters                types.List  `tfsdk:"filters"`
	FiltersV6              types.List  `tfsdk:"filters_v6"`
	IgnoreClientUid        types.Bool  `tfsdk:"ignore_client_uid"`
	IgnoreList             types.List  `tfsdk:"ignore_list"`
	LeaseTime              types.Int64 `tfsdk:"lease_time"`
	LeaseTimeV6            types.Int64 `tfsdk:"lease_time_v6"`
}

var IpamsvcDHCPConfigAttrTypes = map[string]attr.Type{
	"abandoned_reclaim_time":    types.Int64Type,
	"abandoned_reclaim_time_v6": types.Int64Type,
	"allow_unknown":             types.BoolType,
	"allow_unknown_v6":          types.BoolType,
	"filters":                   types.ListType{ElemType: types.StringType},
	"filters_v6":                types.ListType{ElemType: types.StringType},
	"ignore_client_uid":         types.BoolType,
	"ignore_list":               types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcIgnoreItemAttrTypes}},
	"lease_time":                types.Int64Type,
	"lease_time_v6":             types.Int64Type,
}

var IpamsvcDHCPConfigResourceSchema = schema.Schema{
	MarkdownDescription: `A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.`,
	Attributes:          IpamsvcDHCPConfigResourceSchemaAttributes,
}

var IpamsvcDHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned_reclaim_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The abandoned reclaim time in seconds for IPV4 clients.`,
	},
	"abandoned_reclaim_time_v6": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The abandoned reclaim time in seconds for IPV6 clients.`,
	},
	"allow_unknown": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Disable to allow leases only for known IPv4 clients, those for which a fixed address is configured.`,
	},
	"allow_unknown_v6": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Disable to allow leases only for known IPV6 clients, those for which a fixed address is configured.`,
	},
	"filters": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"filters_v6": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"ignore_client_uid": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Enable to ignore the client UID when issuing a DHCP lease. Use this option to prevent assigning two IP addresses for a client which does not have a UID during one phase of PXE boot but acquires one for the other phase.`,
	},
	"ignore_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcIgnoreItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of clients to ignore requests from.`,
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The lease duration in seconds.`,
	},
	"lease_time_v6": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The lease duration in seconds for IPV6 clients.`,
	},
}

func expandIpamsvcDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPConfigModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPConfig {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPConfig{
		AbandonedReclaimTime:   ptr(int64(m.AbandonedReclaimTime.ValueInt64())),
		AbandonedReclaimTimeV6: ptr(int64(m.AbandonedReclaimTimeV6.ValueInt64())),
		AllowUnknown:           m.AllowUnknown.ValueBoolPointer(),
		AllowUnknownV6:         m.AllowUnknownV6.ValueBoolPointer(),
		Filters:                ExpandFrameworkListString(ctx, m.Filters, diags),
		FiltersV6:              ExpandFrameworkListString(ctx, m.FiltersV6, diags),
		IgnoreClientUid:        m.IgnoreClientUid.ValueBoolPointer(),
		IgnoreList:             ExpandFrameworkListNestedBlock(ctx, m.IgnoreList, diags, expandIpamsvcIgnoreItem),
		LeaseTime:              ptr(int64(m.LeaseTime.ValueInt64())),
		LeaseTimeV6:            ptr(int64(m.LeaseTimeV6.ValueInt64())),
	}
	return to
}

func flattenIpamsvcDHCPConfig(ctx context.Context, from *ipam.IpamsvcDHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPConfigAttrTypes)
	}
	m := IpamsvcDHCPConfigModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPConfigModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPConfigModel{}
	}

	m.AbandonedReclaimTime = types.Int64Value(int64(*from.AbandonedReclaimTime))
	m.AbandonedReclaimTimeV6 = types.Int64Value(int64(*from.AbandonedReclaimTimeV6))
	m.AllowUnknown = types.BoolPointerValue(from.AllowUnknown)
	m.AllowUnknownV6 = types.BoolPointerValue(from.AllowUnknownV6)
	m.Filters = FlattenFrameworkListString(ctx, from.Filters, diags)
	m.FiltersV6 = FlattenFrameworkListString(ctx, from.FiltersV6, diags)
	m.IgnoreClientUid = types.BoolPointerValue(from.IgnoreClientUid)
	m.IgnoreList = FlattenFrameworkListNestedBlock(ctx, from.IgnoreList, IpamsvcIgnoreItemAttrTypes, diags, flattenIpamsvcIgnoreItem)
	m.LeaseTime = types.Int64Value(int64(*from.LeaseTime))
	m.LeaseTimeV6 = types.Int64Value(int64(*from.LeaseTimeV6))

}
