package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigServerModel struct {
	AddEdnsOptionInOutgoingQuery                types.Bool        `tfsdk:"add_edns_option_in_outgoing_query"`
	AutoSortViews                               types.Bool        `tfsdk:"auto_sort_views"`
	Comment                                     types.String      `tfsdk:"comment"`
	CreatedAt                                   timetypes.RFC3339 `tfsdk:"created_at"`
	CustomRootNs                                types.List        `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled                         types.Bool        `tfsdk:"custom_root_ns_enabled"`
	DnssecEnableValidation                      types.Bool        `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled                               types.Bool        `tfsdk:"dnssec_enabled"`
	DnssecRootKeys                              types.List        `tfsdk:"dnssec_root_keys"`
	DnssecTrustAnchors                          types.List        `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry                        types.Bool        `tfsdk:"dnssec_validate_expiry"`
	EcsEnabled                                  types.Bool        `tfsdk:"ecs_enabled"`
	EcsForwarding                               types.Bool        `tfsdk:"ecs_forwarding"`
	EcsPrefixV4                                 types.Int64       `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6                                 types.Int64       `tfsdk:"ecs_prefix_v6"`
	EcsZones                                    types.List        `tfsdk:"ecs_zones"`
	FilterAaaaAcl                               types.List        `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                              types.String      `tfsdk:"filter_aaaa_on_v4"`
	Forwarders                                  types.List        `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool        `tfsdk:"forwarders_only"`
	GssTsigEnabled                              types.Bool        `tfsdk:"gss_tsig_enabled"`
	Id                                          types.String      `tfsdk:"id"`
	InheritanceSources                          types.Object      `tfsdk:"inheritance_sources"`
	KerberosKeys                                types.List        `tfsdk:"kerberos_keys"`
	LameTtl                                     types.Int64       `tfsdk:"lame_ttl"`
	LogQueryResponse                            types.Bool        `tfsdk:"log_query_response"`
	MatchRecursiveOnly                          types.Bool        `tfsdk:"match_recursive_only"`
	MaxCacheTtl                                 types.Int64       `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                              types.Int64       `tfsdk:"max_negative_ttl"`
	MinimalResponses                            types.Bool        `tfsdk:"minimal_responses"`
	Name                                        types.String      `tfsdk:"name"`
	Notify                                      types.Bool        `tfsdk:"notify"`
	QueryAcl                                    types.List        `tfsdk:"query_acl"`
	QueryPort                                   types.Int64       `tfsdk:"query_port"`
	RecursionAcl                                types.List        `tfsdk:"recursion_acl"`
	RecursionEnabled                            types.Bool        `tfsdk:"recursion_enabled"`
	RecursiveClients                            types.Int64       `tfsdk:"recursive_clients"`
	ResolverQueryTimeout                        types.Int64       `tfsdk:"resolver_query_timeout"`
	SecondaryAxfrQueryLimit                     types.Int64       `tfsdk:"secondary_axfr_query_limit"`
	SecondarySoaQueryLimit                      types.Int64       `tfsdk:"secondary_soa_query_limit"`
	SortList                                    types.List        `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps           types.Bool        `tfsdk:"synthesize_address_records_from_https"`
	Tags                                        types.Map         `tfsdk:"tags"`
	TransferAcl                                 types.List        `tfsdk:"transfer_acl"`
	UpdateAcl                                   types.List        `tfsdk:"update_acl"`
	UpdatedAt                                   timetypes.RFC3339 `tfsdk:"updated_at"`
	UseForwardersForSubzones                    types.Bool        `tfsdk:"use_forwarders_for_subzones"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool        `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
	Views                                       types.List        `tfsdk:"views"`
}

var ConfigServerAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.BoolType,
	"auto_sort_views":                       types.BoolType,
	"comment":                               types.StringType,
	"created_at":                            timetypes.RFC3339Type{},
	"custom_root_ns":                        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigRootNSAttrTypes}},
	"custom_root_ns_enabled":                types.BoolType,
	"dnssec_enable_validation":              types.BoolType,
	"dnssec_enabled":                        types.BoolType,
	"dnssec_root_keys":                      types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigTrustAnchorAttrTypes}},
	"dnssec_trust_anchors":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigTrustAnchorAttrTypes}},
	"dnssec_validate_expiry":                types.BoolType,
	"ecs_enabled":                           types.BoolType,
	"ecs_forwarding":                        types.BoolType,
	"ecs_prefix_v4":                         types.Int64Type,
	"ecs_prefix_v6":                         types.Int64Type,
	"ecs_zones":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigECSZoneAttrTypes}},
	"filter_aaaa_acl":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"filter_aaaa_on_v4":                     types.StringType,
	"forwarders":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"forwarders_only":                       types.BoolType,
	"gss_tsig_enabled":                      types.BoolType,
	"id":                                    types.StringType,
	"inheritance_sources":                   types.ObjectType{AttrTypes: ConfigServerInheritanceAttrTypes},
	"kerberos_keys":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigKerberosKeyAttrTypes}},
	"lame_ttl":                              types.Int64Type,
	"log_query_response":                    types.BoolType,
	"match_recursive_only":                  types.BoolType,
	"max_cache_ttl":                         types.Int64Type,
	"max_negative_ttl":                      types.Int64Type,
	"minimal_responses":                     types.BoolType,
	"name":                                  types.StringType,
	"notify":                                types.BoolType,
	"query_acl":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"query_port":                            types.Int64Type,
	"recursion_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"recursion_enabled":                     types.BoolType,
	"recursive_clients":                     types.Int64Type,
	"resolver_query_timeout":                types.Int64Type,
	"secondary_axfr_query_limit":            types.Int64Type,
	"secondary_soa_query_limit":             types.Int64Type,
	"sort_list":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigSortListItemAttrTypes}},
	"synthesize_address_records_from_https": types.BoolType,
	"tags":                                  types.MapType{ElemType: types.StringType},
	"transfer_acl":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"update_acl":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"updated_at":                            timetypes.RFC3339Type{},
	"use_forwarders_for_subzones":           types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
	"views": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigDisplayViewAttrTypes}},
}

var ConfigServerResourceSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_.",
	},
	"auto_sort_views": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Controls manual/automatic views ordering.  Defaults to _true_.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Comment for configuration.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"custom_root_ns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigRootNSResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty.",
	},
	"custom_root_ns_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_.",
	},
	"dnssec_enable_validation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_.",
	},
	"dnssec_root_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigTrustAnchorResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation.",
	},
	"dnssec_trust_anchors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigTrustAnchorResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty.",
	},
	"dnssec_validate_expiry": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_.",
	},
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false_.",
	},
	"ecs_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_.",
	},
	"ecs_prefix_v4": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(24),
		MarkdownDescription: "Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24.",
	},
	"ecs_prefix_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(56),
		MarkdownDescription: "Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56.",
	},
	"ecs_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigECSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty.",
	},
	"filter_aaaa_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_.",
	},
	"filter_aaaa_on_v4": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("no"),
		MarkdownDescription: "_filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_",
	},
	"forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.  Defaults to _false_.",
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: ConfigServerInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigKerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"lame_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: "Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600.",
	},
	"log_query_response": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Control DNS query/response logging functionality.  Defaults to _true_.",
	},
	"match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_.",
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(604800),
		MarkdownDescription: "Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d).",
	},
	"max_negative_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10800),
		MarkdownDescription: "Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h).",
	},
	"minimal_responses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of configuration.",
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_notify_ all external secondary DNS servers.  Defaults to _false_.",
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.",
	},
	"query_port": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Source port for outbound DNS queries. When set to 0 the port is unspecified and the implementation may randomize it using any available ports.  Defaults to 0.",
	},
	"recursion_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ field will be used instead.  Defaults to empty.",
	},
	"recursion_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to allow recursive DNS queries.  Defaults to _true_.",
	},
	"recursive_clients": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1000),
		MarkdownDescription: "Optional. Defines the number of simultaneous recursive lookups the server will perform on behalf of its clients.  Defaults to 1000.",
	},
	"resolver_query_timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "Optional. Seconds before a recursive query times out.  Unsigned integer, min 10 max 30.  Defaults to 10.",
	},
	"secondary_axfr_query_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Maximum concurrent inbound AXFRs. When set to 0 a host-dependent default will be used.  Defaults to 0.",
	},
	"secondary_soa_query_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Optional. Maximum concurrent outbound SOA queries. When set to 0 a host-dependent default will be used.  Defaults to 0.",
	},
	"sort_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigSortListItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_.",
	},
	"synthesize_address_records_from_https": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tagging specifics.",
	},
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty.",
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"use_forwarders_for_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.",
	},
	"use_root_forwarders_for_local_resolution_with_b1td": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_.",
	},
	"views": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigDisplayViewResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "Optional. Ordered list of _dns/display_view_ objects served by any of _dns/host_ assigned to a particular DNS Config Profile. Automatically determined. Allows re-ordering only.",
	},
}

func ExpandConfigServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigServer {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigServer{
		AddEdnsOptionInOutgoingQuery:      flex.ExpandBoolPointer(m.AddEdnsOptionInOutgoingQuery),
		AutoSortViews:                     flex.ExpandBoolPointer(m.AutoSortViews),
		Comment:                           flex.ExpandStringPointer(m.Comment),
		CustomRootNs:                      flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandConfigRootNS),
		CustomRootNsEnabled:               flex.ExpandBoolPointer(m.CustomRootNsEnabled),
		DnssecEnableValidation:            flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:                     flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandConfigTrustAnchor),
		DnssecValidateExpiry:              flex.ExpandBoolPointer(m.DnssecValidateExpiry),
		EcsEnabled:                        flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding:                     flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:                       flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:                       flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:                          flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandConfigECSZone),
		FilterAaaaAcl:                     flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaAcl, diags, ExpandConfigACLItem),
		FilterAaaaOnV4:                    flex.ExpandStringPointer(m.FilterAaaaOnV4),
		Forwarders:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandConfigForwarder),
		ForwardersOnly:                    flex.ExpandBoolPointer(m.ForwardersOnly),
		GssTsigEnabled:                    flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:                ExpandConfigServerInheritance(ctx, m.InheritanceSources, diags),
		KerberosKeys:                      flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandConfigKerberosKey),
		LameTtl:                           flex.ExpandInt64Pointer(m.LameTtl),
		LogQueryResponse:                  flex.ExpandBoolPointer(m.LogQueryResponse),
		MatchRecursiveOnly:                flex.ExpandBoolPointer(m.MatchRecursiveOnly),
		MaxCacheTtl:                       flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNegativeTtl:                    flex.ExpandInt64Pointer(m.MaxNegativeTtl),
		MinimalResponses:                  flex.ExpandBoolPointer(m.MinimalResponses),
		Name:                              flex.ExpandString(m.Name),
		Notify:                            flex.ExpandBoolPointer(m.Notify),
		QueryAcl:                          flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandConfigACLItem),
		QueryPort:                         flex.ExpandInt64Pointer(m.QueryPort),
		RecursionAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.RecursionAcl, diags, ExpandConfigACLItem),
		RecursionEnabled:                  flex.ExpandBoolPointer(m.RecursionEnabled),
		RecursiveClients:                  flex.ExpandInt64Pointer(m.RecursiveClients),
		ResolverQueryTimeout:              flex.ExpandInt64Pointer(m.ResolverQueryTimeout),
		SecondaryAxfrQueryLimit:           flex.ExpandInt64Pointer(m.SecondaryAxfrQueryLimit),
		SecondarySoaQueryLimit:            flex.ExpandInt64Pointer(m.SecondarySoaQueryLimit),
		SortList:                          flex.ExpandFrameworkListNestedBlock(ctx, m.SortList, diags, ExpandConfigSortListItem),
		SynthesizeAddressRecordsFromHttps: flex.ExpandBoolPointer(m.SynthesizeAddressRecordsFromHttps),
		Tags:                              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransferAcl:                       flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandConfigACLItem),
		UpdateAcl:                         flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandConfigACLItem),
		UseForwardersForSubzones:          flex.ExpandBoolPointer(m.UseForwardersForSubzones),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
	}
	return to
}

func FlattenConfigServer(ctx context.Context, from *dns_config.ConfigServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigServerAttrTypes)
	}
	m := ConfigServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigServerModel) Flatten(ctx context.Context, from *dns_config.ConfigServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigServerModel{}
	}
	m.AddEdnsOptionInOutgoingQuery = types.BoolPointerValue(from.AddEdnsOptionInOutgoingQuery)
	m.AutoSortViews = types.BoolPointerValue(from.AutoSortViews)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, ConfigRootNSAttrTypes, diags, FlattenConfigRootNS)
	m.CustomRootNsEnabled = types.BoolPointerValue(from.CustomRootNsEnabled)
	m.DnssecEnableValidation = types.BoolPointerValue(from.DnssecEnableValidation)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecRootKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecRootKeys, ConfigTrustAnchorAttrTypes, diags, FlattenConfigTrustAnchor)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, ConfigTrustAnchorAttrTypes, diags, FlattenConfigTrustAnchor)
	m.DnssecValidateExpiry = types.BoolPointerValue(from.DnssecValidateExpiry)
	m.EcsEnabled = types.BoolPointerValue(from.EcsEnabled)
	m.EcsForwarding = types.BoolPointerValue(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64Pointer(from.EcsPrefixV4)
	m.EcsPrefixV6 = flex.FlattenInt64Pointer(from.EcsPrefixV6)
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ConfigECSZoneAttrTypes, diags, FlattenConfigECSZone)
	m.FilterAaaaAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.FilterAaaaOnV4 = flex.FlattenStringPointer(from.FilterAaaaOnV4)
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenConfigServerInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, ConfigKerberosKeyAttrTypes, diags, FlattenConfigKerberosKey)
	m.LameTtl = flex.FlattenInt64Pointer(from.LameTtl)
	m.LogQueryResponse = types.BoolPointerValue(from.LogQueryResponse)
	m.MatchRecursiveOnly = types.BoolPointerValue(from.MatchRecursiveOnly)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxNegativeTtl = flex.FlattenInt64Pointer(from.MaxNegativeTtl)
	m.MinimalResponses = types.BoolPointerValue(from.MinimalResponses)
	m.Name = flex.FlattenString(from.Name)
	m.Notify = types.BoolPointerValue(from.Notify)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.QueryPort = flex.FlattenInt64Pointer(from.QueryPort)
	m.RecursionAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursionAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.RecursionEnabled = types.BoolPointerValue(from.RecursionEnabled)
	m.RecursiveClients = flex.FlattenInt64Pointer(from.RecursiveClients)
	m.ResolverQueryTimeout = flex.FlattenInt64Pointer(from.ResolverQueryTimeout)
	m.SecondaryAxfrQueryLimit = flex.FlattenInt64Pointer(from.SecondaryAxfrQueryLimit)
	m.SecondarySoaQueryLimit = flex.FlattenInt64Pointer(from.SecondarySoaQueryLimit)
	m.SortList = flex.FlattenFrameworkListNestedBlock(ctx, from.SortList, ConfigSortListItemAttrTypes, diags, FlattenConfigSortListItem)
	m.SynthesizeAddressRecordsFromHttps = types.BoolPointerValue(from.SynthesizeAddressRecordsFromHttps)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.UseForwardersForSubzones = types.BoolPointerValue(from.UseForwardersForSubzones)
	m.UseRootForwardersForLocalResolutionWithB1td = types.BoolPointerValue(from.UseRootForwardersForLocalResolutionWithB1td)
	m.Views = flex.FlattenFrameworkListNestedBlock(ctx, from.Views, ConfigDisplayViewAttrTypes, diags, FlattenConfigDisplayView)
}
