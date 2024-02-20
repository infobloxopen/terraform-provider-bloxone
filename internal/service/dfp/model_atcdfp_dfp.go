package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcdfpDfpModel struct {
	CreatedTime         timetypes.RFC3339 `tfsdk:"created_time"`
	DefaultResolvers    types.List        `tfsdk:"default_resolvers"`
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
	Resolvers           types.List        `tfsdk:"resolvers"`
	ResolversAll        types.List        `tfsdk:"resolvers_all"`
	ServiceId           types.String      `tfsdk:"service_id"`
	ServiceName         types.String      `tfsdk:"service_name"`
	SiteId              types.String      `tfsdk:"site_id"`
	UpdatedTime         timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcdfpDfpAttrTypes = map[string]attr.Type{
	"created_time":          timetypes.RFC3339Type{},
	"default_resolvers":     types.ListType{ElemType: types.StringType},
	"elb_ip_list":           types.ListType{ElemType: types.StringType},
	"forwarding_policy":     types.StringType,
	"host":                  types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpDfpHostAttrTypes}},
	"id":                    types.Int64Type,
	"internal_domain_lists": types.ListType{ElemType: types.Int64Type},
	"name":                  types.StringType,
	"net_addr_policy_ids":   types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpNetAddrPolicyAssignmentAttrTypes}},
	"ophid":                 types.StringType,
	"policy_id":             types.Int64Type,
	"pop_region_id":         types.Int64Type,
	"resolvers":             types.ListType{ElemType: types.StringType},
	"resolvers_all":         types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpResolverAttrTypes}},
	"service_id":            types.StringType,
	"service_name":          types.StringType,
	"site_id":               types.StringType,
	"updated_time":          timetypes.RFC3339Type{},
}

var AtcdfpDfpResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this DNS Forwarding Proxy object was created.",
	},
	"default_resolvers": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable.  Deprecated DO NOT USE. Use resolvers_all.",
	},
	"elb_ip_list": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as ELB IPs.",
	},
	"forwarding_policy": schema.StringAttribute{
		Optional: true,
	},
	"host": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcdfpDfpHostResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "host information. For internal Use only.",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
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
			Attributes: AtcdfpNetAddrPolicyAssignmentResourceSchemaAttributes,
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
	"resolvers": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers. Deprecated DO NOT USE. Use resolvers_all.",
	},
	"resolvers_all": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcdfpResolverResourceSchemaAttributes,
		},
		Optional: true,
	},
	"service_id": schema.StringAttribute{
		Computed:            true,
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

func ExpandAtcdfpDfp(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpDfp {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDfpModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDfpModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpDfp {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpDfp{
		ForwardingPolicy: flex.ExpandStringPointer(m.ForwardingPolicy),
		Host:             flex.ExpandFrameworkListNestedBlock(ctx, m.Host, diags, ExpandAtcdfpDfpHost),
		NetAddrPolicyIds: flex.ExpandFrameworkListNestedBlock(ctx, m.NetAddrPolicyIds, diags, ExpandAtcdfpNetAddrPolicyAssignment),
		ResolversAll:     flex.ExpandFrameworkListNestedBlock(ctx, m.ResolversAll, diags, ExpandAtcdfpResolver),
		// InternalDomainLists // TODO: should have been expanded, but generator didn't know how to.
	}
	return to
}

func FlattenAtcdfpDfp(ctx context.Context, from *dfp.AtcdfpDfp, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDfpAttrTypes)
	}
	m := AtcdfpDfpModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDfpAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDfpModel) Flatten(ctx context.Context, from *dfp.AtcdfpDfp, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDfpModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.DefaultResolvers = flex.FlattenFrameworkListString(ctx, from.DefaultResolvers, diags)
	m.ElbIpList = flex.FlattenFrameworkListString(ctx, from.ElbIpList, diags)
	m.ForwardingPolicy = flex.FlattenStringPointer(from.ForwardingPolicy)
	m.Host = flex.FlattenFrameworkListNestedBlock(ctx, from.Host, AtcdfpDfpHostAttrTypes, diags, FlattenAtcdfpDfpHost)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetAddrPolicyIds = flex.FlattenFrameworkListNestedBlock(ctx, from.NetAddrPolicyIds, AtcdfpNetAddrPolicyAssignmentAttrTypes, diags, FlattenAtcdfpNetAddrPolicyAssignment)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
	m.PopRegionId = flex.FlattenInt32Pointer(from.PopRegionId)
	m.Resolvers = flex.FlattenFrameworkListString(ctx, from.Resolvers, diags)
	m.ResolversAll = flex.FlattenFrameworkListNestedBlock(ctx, from.ResolversAll, AtcdfpResolverAttrTypes, diags, FlattenAtcdfpResolver)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceName = flex.FlattenStringPointer(from.ServiceName)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
	// InternalDomainLists // TODO: should have been flattened, but generator didn't know how to.
}
