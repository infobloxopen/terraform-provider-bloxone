package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DfpModel struct {
	CreatedTime         timetypes.RFC3339 `tfsdk:"created_time"`
	ElbIpList           types.List        `tfsdk:"elb_ip_list"`
	ForwardingPolicy    types.String      `tfsdk:"forwarding_policy"`
	Host                types.List        `tfsdk:"host"`
	Id                  types.Int64       `tfsdk:"id"`
	InternalDomainLists types.List        `tfsdk:"internal_domain_lists"`
	Name                types.String      `tfsdk:"name"`
	NetAddrPolicyIds    types.List        `tfsdk:"net_addr_policy_ids"`
	Ophid               types.String      `tfsdk:"ophid"`
	PolicyId            types.Int64       `tfsdk:"policy_id"`
	PopRegionId         types.Int64       `tfsdk:"pop_region_id"`
	ResolversAll        types.List        `tfsdk:"resolvers_all"`
	ServiceId           types.String      `tfsdk:"service_id"`
	ServiceName         types.String      `tfsdk:"service_name"`
	SiteId              types.String      `tfsdk:"site_id"`
	UpdatedTime         timetypes.RFC3339 `tfsdk:"updated_time"`
}

var DfpAttrTypes = map[string]attr.Type{
	"created_time":          timetypes.RFC3339Type{},
	"elb_ip_list":           types.ListType{ElemType: types.StringType},
	"forwarding_policy":     types.StringType,
	"host":                  types.ListType{ElemType: types.ObjectType{AttrTypes: DfpHostAttrTypes}},
	"id":                    types.Int64Type,
	"internal_domain_lists": types.ListType{ElemType: types.Int64Type},
	"name":                  types.StringType,
	"net_addr_policy_ids":   types.ListType{ElemType: types.ObjectType{AttrTypes: NetAddrPolicyAssignmentAttrTypes}},
	"ophid":                 types.StringType,
	"policy_id":             types.Int64Type,
	"pop_region_id":         types.Int64Type,
	"resolvers_all":         types.ListType{ElemType: types.ObjectType{AttrTypes: ResolverAttrTypes}},
	"service_id":            types.StringType,
	"service_name":          types.StringType,
	"site_id":               types.StringType,
	"updated_time":          timetypes.RFC3339Type{},
}

var DfpResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this DNS Forwarding Proxy object was created.",
	},
	"elb_ip_list": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as ELB IPs.",
	},
	"forwarding_policy": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Policy Identifier for DNS Forwarding Proxy.",
	},
	"host": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DfpHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "host information. For internal Use only.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The DNS Forwarding Proxy object identifier.",
	},
	"internal_domain_lists": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		MarkdownDescription: "The list of internal domains list IDs that are associated with this DFP",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the DNS Forwarding Proxy.",
	},
	"net_addr_policy_ids": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetAddrPolicyAssignmentResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of network-address-scoped security policy assignments",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The On-Prem Host identifier.",
	},
	"policy_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The identifier of the security policy with which the DNS Forwarding Proxy is associated.",
	},
	"pop_region_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Point of Presence (PoP) region",
	},
	"resolvers_all": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ResolverResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of DNS resolvers",
	},
	"service_id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The On-Prem Application Service identifier. For internal Use only",
	},
	"service_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The On-Prem Application Service name. For internal Use only",
	},
	"site_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this DNS Forwarding Proxy object was last updated.",
	},
}

func ExpandDfp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.Dfp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DfpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DfpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.Dfp {
	if m == nil {
		return nil
	}
	to := &dfp.Dfp{
		ForwardingPolicy:    flex.ExpandStringPointer(m.ForwardingPolicy),
		Host:                flex.ExpandFrameworkListNestedBlock(ctx, m.Host, diags, ExpandDfpHost),
		NetAddrPolicyIds:    flex.ExpandFrameworkListNestedBlock(ctx, m.NetAddrPolicyIds, diags, ExpandNetAddrPolicyAssignment),
		ResolversAll:        flex.ExpandFrameworkListNestedBlock(ctx, m.ResolversAll, diags, ExpandResolver),
		InternalDomainLists: flex.ExpandFrameworkListInt32(ctx, m.InternalDomainLists, diags),
	}
	return to
}

func (m *DfpModel) ExpandCreateOrUpdatePayload(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpCreateOrUpdatePayload {
	if m == nil {
		return nil
	}
	to := &dfp.DfpCreateOrUpdatePayload{
		ForwardingPolicy:    flex.ExpandStringPointer(m.ForwardingPolicy),
		Host:                flex.ExpandFrameworkListNestedBlock(ctx, m.Host, diags, ExpandDfpHost),
		InternalDomainLists: flex.ExpandFrameworkListInt32(ctx, m.InternalDomainLists, diags),
		Name:                flex.ExpandStringPointer(m.Name),
		PopRegionId:         flex.ExpandInt32Pointer(m.PopRegionId),
		ResolversAll:        flex.ExpandFrameworkListNestedBlock(ctx, m.ResolversAll, diags, ExpandResolver),
		ServiceId:           flex.ExpandStringPointer(m.ServiceId),
		ServiceName:         flex.ExpandStringPointer(m.ServiceName),
		SiteId:              flex.ExpandStringPointer(m.SiteId),
	}
	return to
}

func FlattenDfp(ctx context.Context, from *dfp.Dfp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DfpAttrTypes)
	}
	m := DfpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DfpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DfpModel) Flatten(ctx context.Context, from *dfp.Dfp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DfpModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.ElbIpList = flex.FlattenFrameworkListString(ctx, from.ElbIpList, diags)
	m.ForwardingPolicy = flex.FlattenStringPointer(from.ForwardingPolicy)
	m.Host = flex.FlattenFrameworkListNestedBlock(ctx, from.Host, DfpHostAttrTypes, diags, FlattenDfpHost)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetAddrPolicyIds = flex.FlattenFrameworkListNestedBlock(ctx, from.NetAddrPolicyIds, NetAddrPolicyAssignmentAttrTypes, diags, FlattenNetAddrPolicyAssignment)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
	m.PopRegionId = flex.FlattenInt32Pointer(from.PopRegionId)
	m.ResolversAll = flex.FlattenFrameworkListNestedBlock(ctx, from.ResolversAll, ResolverAttrTypes, diags, FlattenResolver)
	m.ServiceName = flex.FlattenStringPointer(from.ServiceName)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
	m.InternalDomainLists = flex.FlattenFrameworkListInt32(ctx, from.InternalDomainLists, diags)
}
