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

type IpamsvcInheritedDHCPConfigModel struct {
	AbandonedReclaimTime   types.Object `tfsdk:"abandoned_reclaim_time"`
	AbandonedReclaimTimeV6 types.Object `tfsdk:"abandoned_reclaim_time_v6"`
	AllowUnknown           types.Object `tfsdk:"allow_unknown"`
	AllowUnknownV6         types.Object `tfsdk:"allow_unknown_v6"`
	Filters                types.Object `tfsdk:"filters"`
	FiltersV6              types.Object `tfsdk:"filters_v6"`
	IgnoreClientUid        types.Object `tfsdk:"ignore_client_uid"`
	IgnoreList             types.Object `tfsdk:"ignore_list"`
	LeaseTime              types.Object `tfsdk:"lease_time"`
	LeaseTimeV6            types.Object `tfsdk:"lease_time_v6"`
}

var IpamsvcInheritedDHCPConfigAttrTypes = map[string]attr.Type{
	"abandoned_reclaim_time":    types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"abandoned_reclaim_time_v6": types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"allow_unknown":             types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"allow_unknown_v6":          types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"filters":                   types.ObjectType{AttrTypes: InheritedDHCPConfigFilterListAttrTypes},
	"filters_v6":                types.ObjectType{AttrTypes: InheritedDHCPConfigFilterListAttrTypes},
	"ignore_client_uid":         types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ignore_list":               types.ObjectType{AttrTypes: InheritedDHCPConfigIgnoreItemListAttrTypes},
	"lease_time":                types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"lease_time_v6":             types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
}

var IpamsvcInheritedDHCPConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned_reclaim_time": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"abandoned_reclaim_time_v6": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"allow_unknown": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"allow_unknown_v6": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"filters": schema.SingleNestedAttribute{
		Attributes: InheritedDHCPConfigFilterListResourceSchemaAttributes,
		Optional:   true,
	},
	"filters_v6": schema.SingleNestedAttribute{
		Attributes: InheritedDHCPConfigFilterListResourceSchemaAttributes,
		Optional:   true,
	},
	"ignore_client_uid": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:   true,
	},
	"ignore_list": schema.SingleNestedAttribute{
		Attributes: InheritedDHCPConfigIgnoreItemListResourceSchemaAttributes,
		Optional:   true,
	},
	"lease_time": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"lease_time_v6": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcInheritedDHCPConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcInheritedDHCPConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcInheritedDHCPConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPConfig {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcInheritedDHCPConfig{
		AbandonedReclaimTime:   ExpandInheritanceInheritedUInt32(ctx, m.AbandonedReclaimTime, diags),
		AbandonedReclaimTimeV6: ExpandInheritanceInheritedUInt32(ctx, m.AbandonedReclaimTimeV6, diags),
		AllowUnknown:           ExpandInheritanceInheritedBool(ctx, m.AllowUnknown, diags),
		AllowUnknownV6:         ExpandInheritanceInheritedBool(ctx, m.AllowUnknownV6, diags),
		Filters:                ExpandInheritedDHCPConfigFilterList(ctx, m.Filters, diags),
		FiltersV6:              ExpandInheritedDHCPConfigFilterList(ctx, m.FiltersV6, diags),
		IgnoreClientUid:        ExpandInheritanceInheritedBool(ctx, m.IgnoreClientUid, diags),
		IgnoreList:             ExpandInheritedDHCPConfigIgnoreItemList(ctx, m.IgnoreList, diags),
		LeaseTime:              ExpandInheritanceInheritedUInt32(ctx, m.LeaseTime, diags),
		LeaseTimeV6:            ExpandInheritanceInheritedUInt32(ctx, m.LeaseTimeV6, diags),
	}
	return to
}

func FlattenIpamsvcInheritedDHCPConfig(ctx context.Context, from *ipam.IpamsvcInheritedDHCPConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDHCPConfigAttrTypes)
	}
	m := IpamsvcInheritedDHCPConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDHCPConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDHCPConfigModel) Flatten(ctx context.Context, from *ipam.IpamsvcInheritedDHCPConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDHCPConfigModel{}
	}
	m.AbandonedReclaimTime = FlattenInheritanceInheritedUInt32(ctx, from.AbandonedReclaimTime, diags)
	m.AbandonedReclaimTimeV6 = FlattenInheritanceInheritedUInt32(ctx, from.AbandonedReclaimTimeV6, diags)
	m.AllowUnknown = FlattenInheritanceInheritedBool(ctx, from.AllowUnknown, diags)
	m.AllowUnknownV6 = FlattenInheritanceInheritedBool(ctx, from.AllowUnknownV6, diags)
	m.Filters = FlattenInheritedDHCPConfigFilterList(ctx, from.Filters, diags)
	m.FiltersV6 = FlattenInheritedDHCPConfigFilterList(ctx, from.FiltersV6, diags)
	m.IgnoreClientUid = FlattenInheritanceInheritedBool(ctx, from.IgnoreClientUid, diags)
	m.IgnoreList = FlattenInheritedDHCPConfigIgnoreItemList(ctx, from.IgnoreList, diags)
	m.LeaseTime = FlattenInheritanceInheritedUInt32(ctx, from.LeaseTime, diags)
	m.LeaseTimeV6 = FlattenInheritanceInheritedUInt32(ctx, from.LeaseTimeV6, diags)
}
