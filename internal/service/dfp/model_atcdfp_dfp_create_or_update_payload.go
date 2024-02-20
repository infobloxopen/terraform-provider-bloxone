package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcdfpDfpCreateOrUpdatePayloadModel struct {
	DefaultResolvers    types.List   `tfsdk:"default_resolvers"`
	ForwardingPolicy    types.String `tfsdk:"forwarding_policy"`
	Host                types.List   `tfsdk:"host"`
	Id                  types.Int64  `tfsdk:"id"`
	InternalDomainLists types.List   `tfsdk:"internal_domain_lists"`
	Name                types.String `tfsdk:"name"`
	PopRegionId         types.Int64  `tfsdk:"pop_region_id"`
	Resolvers           types.List   `tfsdk:"resolvers"`
	ResolversAll        types.List   `tfsdk:"resolvers_all"`
	ServiceId           types.String `tfsdk:"service_id"`
	ServiceName         types.String `tfsdk:"service_name"`
	SiteId              types.String `tfsdk:"site_id"`
}

var AtcdfpDfpCreateOrUpdatePayloadAttrTypes = map[string]attr.Type{
	"default_resolvers":     types.ListType{ElemType: types.StringType},
	"forwarding_policy":     types.StringType,
	"host":                  types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpDfpHostAttrTypes}},
	"id":                    types.Int64Type,
	"internal_domain_lists": types.ListType{ElemType: types.Int64Type},
	"name":                  types.StringType,
	"pop_region_id":         types.Int64Type,
	"resolvers":             types.ListType{ElemType: types.StringType},
	"resolvers_all":         types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpResolverAttrTypes}},
	"service_id":            types.StringType,
	"service_name":          types.StringType,
	"site_id":               types.StringType,
}

var AtcdfpDfpCreateOrUpdatePayloadResourceSchemaAttributes = map[string]schema.Attribute{
	"default_resolvers": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable. Deprecated DO NOT USE. Use resolvers_all.",
	},
	"forwarding_policy": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of DNS resolver as Forwarding Policy. It can hold values as ib_cloud_first, external_first or external_only The default value is ib_cloud_first. If empty string is sent then ib_cloud_first will be considered.",
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
		MarkdownDescription: "The list of internal domain list ids associated with this DFP (or resolvers)",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the DNS Forwarding Proxy.",
	},
	"pop_region_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Point of Presence (PoP) region",
	},
	"resolvers": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers. Deprecated DO NOT USE. Use resolvers_all.",
	},
	"resolvers_all": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcdfpResolverResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The DNS forwarding proxy additional resolvers used for fallback and local resolution. This field replaces resolvers and default_resolvers fields which are deprecated. Either deprecated fields or new field can be used, both can not be used at same time.",
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNS Forwarding Proxy Service ID object identifier.",
	},
	"service_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the DNS Forwarding Proxy Service.",
	},
	"site_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.",
	},
}

func ExpandAtcdfpDfpCreateOrUpdatePayload(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpDfpCreateOrUpdatePayload {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDfpCreateOrUpdatePayloadModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDfpCreateOrUpdatePayloadModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpDfpCreateOrUpdatePayload {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpDfpCreateOrUpdatePayload{
		DefaultResolvers:    flex.ExpandFrameworkListString(ctx, m.DefaultResolvers, diags),
		ForwardingPolicy:    flex.ExpandStringPointer(m.ForwardingPolicy),
		Host:                flex.ExpandFrameworkListNestedBlock(ctx, m.Host, diags, ExpandAtcdfpDfpHost),
		Name:                flex.ExpandStringPointer(m.Name),
		PopRegionId:         flex.ExpandInt32Pointer(m.PopRegionId),
		Resolvers:           flex.ExpandFrameworkListString(ctx, m.Resolvers, diags),
		ResolversAll:        flex.ExpandFrameworkListNestedBlock(ctx, m.ResolversAll, diags, ExpandAtcdfpResolver),
		ServiceId:           flex.ExpandStringPointer(m.ServiceId),
		ServiceName:         flex.ExpandStringPointer(m.ServiceName),
		SiteId:              flex.ExpandStringPointer(m.SiteId),
		InternalDomainLists: flex.ExpandFrameworkListInt32(ctx, m.InternalDomainLists, diags),
	}
	return to
}

func FlattenAtcdfpDfpCreateOrUpdatePayload(ctx context.Context, from *dfp.AtcdfpDfpCreateOrUpdatePayload, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDfpCreateOrUpdatePayloadAttrTypes)
	}
	m := AtcdfpDfpCreateOrUpdatePayloadModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDfpCreateOrUpdatePayloadAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDfpCreateOrUpdatePayloadModel) Flatten(ctx context.Context, from *dfp.AtcdfpDfpCreateOrUpdatePayload, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDfpCreateOrUpdatePayloadModel{}
	}
	m.DefaultResolvers = flex.FlattenFrameworkListString(ctx, from.DefaultResolvers, diags)
	m.ForwardingPolicy = flex.FlattenStringPointer(from.ForwardingPolicy)
	m.Host = flex.FlattenFrameworkListNestedBlock(ctx, from.Host, AtcdfpDfpHostAttrTypes, diags, FlattenAtcdfpDfpHost)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PopRegionId = flex.FlattenInt32Pointer(from.PopRegionId)
	m.Resolvers = flex.FlattenFrameworkListString(ctx, from.Resolvers, diags)
	m.ResolversAll = flex.FlattenFrameworkListNestedBlock(ctx, from.ResolversAll, AtcdfpResolverAttrTypes, diags, FlattenAtcdfpResolver)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceName = flex.FlattenStringPointer(from.ServiceName)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.InternalDomainLists = flex.FlattenFrameworkListInt32(ctx, from.InternalDomainLists, diags)
}
