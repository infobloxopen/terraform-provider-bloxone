package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

var aclDefaultValues = map[string]attr.Value{
	"access":   types.StringValue("allow"),
	"acl":      types.StringNull(),
	"address":  types.StringNull(),
	"element":  types.StringValue("any"),
	"tsig_key": types.ObjectNull(ConfigTSIGKeyAttrTypes),
}

type ConfigViewModel struct {
	AddEdnsOptionInOutgoingQuery                types.Bool        `tfsdk:"add_edns_option_in_outgoing_query"`
	Comment                                     types.String      `tfsdk:"comment"`
	CreatedAt                                   timetypes.RFC3339 `tfsdk:"created_at"`
	CustomRootNs                                types.List        `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled                         types.Bool        `tfsdk:"custom_root_ns_enabled"`
	Disabled                                    types.Bool        `tfsdk:"disabled"`
	DnssecEnableValidation                      types.Bool        `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled                               types.Bool        `tfsdk:"dnssec_enabled"`
	DnssecRootKeys                              types.List        `tfsdk:"dnssec_root_keys"`
	DnssecTrustAnchors                          types.List        `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry                        types.Bool        `tfsdk:"dnssec_validate_expiry"`
	DtcConfig                                   types.Object      `tfsdk:"dtc_config"`
	EcsEnabled                                  types.Bool        `tfsdk:"ecs_enabled"`
	EcsForwarding                               types.Bool        `tfsdk:"ecs_forwarding"`
	EcsPrefixV4                                 types.Int64       `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6                                 types.Int64       `tfsdk:"ecs_prefix_v6"`
	EcsZones                                    types.List        `tfsdk:"ecs_zones"`
	EdnsUdpSize                                 types.Int64       `tfsdk:"edns_udp_size"`
	FilterAaaaAcl                               types.List        `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                              types.String      `tfsdk:"filter_aaaa_on_v4"`
	Forwarders                                  types.List        `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool        `tfsdk:"forwarders_only"`
	GssTsigEnabled                              types.Bool        `tfsdk:"gss_tsig_enabled"`
	Id                                          types.String      `tfsdk:"id"`
	InheritanceSources                          types.Object      `tfsdk:"inheritance_sources"`
	IpSpaces                                    types.List        `tfsdk:"ip_spaces"`
	LameTtl                                     types.Int64       `tfsdk:"lame_ttl"`
	MatchClientsAcl                             types.List        `tfsdk:"match_clients_acl"`
	MatchDestinationsAcl                        types.List        `tfsdk:"match_destinations_acl"`
	MatchRecursiveOnly                          types.Bool        `tfsdk:"match_recursive_only"`
	MaxCacheTtl                                 types.Int64       `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                              types.Int64       `tfsdk:"max_negative_ttl"`
	MaxUdpSize                                  types.Int64       `tfsdk:"max_udp_size"`
	MinimalResponses                            types.Bool        `tfsdk:"minimal_responses"`
	Name                                        types.String      `tfsdk:"name"`
	Notify                                      types.Bool        `tfsdk:"notify"`
	QueryAcl                                    types.List        `tfsdk:"query_acl"`
	RecursionAcl                                types.List        `tfsdk:"recursion_acl"`
	RecursionEnabled                            types.Bool        `tfsdk:"recursion_enabled"`
	SortList                                    types.List        `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps           types.Bool        `tfsdk:"synthesize_address_records_from_https"`
	Tags                                        types.Map         `tfsdk:"tags"`
	TransferAcl                                 types.List        `tfsdk:"transfer_acl"`
	UpdateAcl                                   types.List        `tfsdk:"update_acl"`
	UpdatedAt                                   timetypes.RFC3339 `tfsdk:"updated_at"`
	UseForwardersForSubzones                    types.Bool        `tfsdk:"use_forwarders_for_subzones"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool        `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
	ZoneAuthority                               types.Object      `tfsdk:"zone_authority"`
}

var ConfigViewAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.BoolType,
	"comment":                               types.StringType,
	"created_at":                            timetypes.RFC3339Type{},
	"custom_root_ns":                        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigRootNSAttrTypes}},
	"custom_root_ns_enabled":                types.BoolType,
	"disabled":                              types.BoolType,
	"dnssec_enable_validation":              types.BoolType,
	"dnssec_enabled":                        types.BoolType,
	"dnssec_root_keys":                      types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigTrustAnchorAttrTypes}},
	"dnssec_trust_anchors":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigTrustAnchorAttrTypes}},
	"dnssec_validate_expiry":                types.BoolType,
	"dtc_config":                            types.ObjectType{AttrTypes: ConfigDTCConfigAttrTypes},
	"ecs_enabled":                           types.BoolType,
	"ecs_forwarding":                        types.BoolType,
	"ecs_prefix_v4":                         types.Int64Type,
	"ecs_prefix_v6":                         types.Int64Type,
	"ecs_zones":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigECSZoneAttrTypes}},
	"edns_udp_size":                         types.Int64Type,
	"filter_aaaa_acl":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"filter_aaaa_on_v4":                     types.StringType,
	"forwarders":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"forwarders_only":                       types.BoolType,
	"gss_tsig_enabled":                      types.BoolType,
	"id":                                    types.StringType,
	"inheritance_sources":                   types.ObjectType{AttrTypes: ConfigViewInheritanceAttrTypes},
	"ip_spaces":                             types.ListType{ElemType: types.StringType},
	"lame_ttl":                              types.Int64Type,
	"match_clients_acl":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"match_destinations_acl":                types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"match_recursive_only":                  types.BoolType,
	"max_cache_ttl":                         types.Int64Type,
	"max_negative_ttl":                      types.Int64Type,
	"max_udp_size":                          types.Int64Type,
	"minimal_responses":                     types.BoolType,
	"name":                                  types.StringType,
	"notify":                                types.BoolType,
	"query_acl":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"recursion_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"recursion_enabled":                     types.BoolType,
	"sort_list":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigSortListItemAttrTypes}},
	"synthesize_address_records_from_https": types.BoolType,
	"tags":                                  types.MapType{ElemType: types.StringType},
	"transfer_acl":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"update_acl":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"updated_at":                            timetypes.RFC3339Type{},
	"use_forwarders_for_subzones":           types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
	"zone_authority": types.ObjectType{AttrTypes: ConfigZoneAuthorityAttrTypes},
}

var ConfigViewResourceSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Comment for view.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The timestamp when the object has been created.`,
	},
	"custom_root_ns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigRootNSResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty.`,
	},
	"custom_root_ns_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_.`,
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.`,
	},
	"dnssec_enable_validation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_.`,
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_.`,
	},
	"dnssec_root_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigTrustAnchorResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `DNSSEC root keys. The root keys are not configurable.  A default list is provided by cloud management and included here for config generation.`,
	},
	"dnssec_trust_anchors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigTrustAnchorResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty.`,
	},
	"dnssec_validate_expiry": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_.`,
	},
	"dtc_config": schema.SingleNestedAttribute{
		Attributes: ConfigDTCConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(ConfigDTCConfigAttrTypes, map[string]attr.Value{
			"default_ttl": types.Int64Value(300),
		}),
		),
	},
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false-.`,
	},
	"ecs_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_.`,
	},
	"ecs_prefix_v4": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(24),
		MarkdownDescription: `Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24.`,
	},
	"ecs_prefix_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(56),
		MarkdownDescription: `Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56.`,
	},
	"ecs_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigECSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty.`,
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1232),
		MarkdownDescription: `Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes.`,
	},
	"filter_aaaa_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_.`,
	},
	"filter_aaaa_on_v4": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("no"),
		MarkdownDescription: "_filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients. Allowed values:\n" +
			"  * _yes_\n" +
			"  * _no_\n" +
			"  * _break_dnssec_\n\n" +
			"  Defaults to _no_",
	},
	"forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty.`,
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to only forward.  Defaults to _false_.`,
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: ConfigViewInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_spaces": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			listvalidator.SizeAtMost(1),
		},
		MarkdownDescription: `The resource identifier.`,
	},
	"lame_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: `Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600.`,
	},
	"match_clients_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}, []attr.Value{
			types.ObjectValueMust(ConfigACLItemAttrTypes, aclDefaultValues),
		})),
		MarkdownDescription: `Optional. Specifies which clients have access to the view.  Defaults to empty.`,
	},
	"match_destinations_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default: listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}, []attr.Value{
			types.ObjectValueMust(ConfigACLItemAttrTypes, aclDefaultValues),
		})),
		MarkdownDescription: `Optional. Specifies which destination addresses have access to the view.  Defaults to empty.`,
	},
	"match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_.`,
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(604800),
		MarkdownDescription: `Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d).`,
	},
	"max_negative_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10800),
		MarkdownDescription: `Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h).`,
	},
	"max_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1232),
		MarkdownDescription: `Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes.`,
	},
	"minimal_responses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_.`,
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: `Name of view.`,
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_notify_ all external secondary DNS servers.  Defaults to _false_.`,
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.`,
	},
	"recursion_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.  Defaults to empty.`,
	},
	"recursion_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. _true_ to allow recursive DNS queries.  Defaults to _true_.`,
	},
	"sort_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigSortListItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_.`,
	},
	"synthesize_address_records_from_https": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `Tagging specifics.`,
	},
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty.`,
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"use_forwarders_for_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.`,
	},
	"use_root_forwarders_for_local_resolution_with_b1td": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_.`,
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes: ConfigZoneAuthorityResourceSchemaAttributes,
		Computed:   true,
		Optional:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(ConfigZoneAuthorityAttrTypes, map[string]attr.Value{
			"default_ttl":       types.Int64Value(28800),
			"expire":            types.Int64Value(2.4192e+06),
			"mname":             types.StringValue("ns.b1ddi"),
			"negative_ttl":      types.Int64Value(900),
			"protocol_mname":    types.StringValue("ns.b1ddi"),
			"protocol_rname":    types.StringValue("hostmaster"),
			"refresh":           types.Int64Value(10800),
			"retry":             types.Int64Value(3600),
			"rname":             types.StringValue("hostmaster"),
			"use_default_mname": types.BoolValue(true),
		}),
		),
	},
}

func ExpandConfigView(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigView {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigViewModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigView {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigView{
		AddEdnsOptionInOutgoingQuery:      flex.ExpandBoolPointer(m.AddEdnsOptionInOutgoingQuery),
		Comment:                           flex.ExpandStringPointer(m.Comment),
		CustomRootNs:                      flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandConfigRootNS),
		CustomRootNsEnabled:               flex.ExpandBoolPointer(m.CustomRootNsEnabled),
		Disabled:                          flex.ExpandBoolPointer(m.Disabled),
		DnssecEnableValidation:            flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:                     flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandConfigTrustAnchor),
		DnssecValidateExpiry:              flex.ExpandBoolPointer(m.DnssecValidateExpiry),
		DtcConfig:                         ExpandConfigDTCConfig(ctx, m.DtcConfig, diags),
		EcsEnabled:                        flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding:                     flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:                       flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:                       flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:                          flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandConfigECSZone),
		EdnsUdpSize:                       flex.ExpandInt64Pointer(m.EdnsUdpSize),
		FilterAaaaAcl:                     flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaAcl, diags, ExpandConfigACLItem),
		FilterAaaaOnV4:                    flex.ExpandStringPointer(m.FilterAaaaOnV4),
		Forwarders:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandConfigForwarder),
		ForwardersOnly:                    flex.ExpandBoolPointer(m.ForwardersOnly),
		GssTsigEnabled:                    flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:                ExpandConfigViewInheritance(ctx, m.InheritanceSources, diags),
		IpSpaces:                          flex.ExpandFrameworkListString(ctx, m.IpSpaces, diags),
		LameTtl:                           flex.ExpandInt64Pointer(m.LameTtl),
		MatchClientsAcl:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MatchClientsAcl, diags, ExpandConfigACLItem),
		MatchDestinationsAcl:              flex.ExpandFrameworkListNestedBlock(ctx, m.MatchDestinationsAcl, diags, ExpandConfigACLItem),
		MatchRecursiveOnly:                flex.ExpandBoolPointer(m.MatchRecursiveOnly),
		MaxCacheTtl:                       flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNegativeTtl:                    flex.ExpandInt64Pointer(m.MaxNegativeTtl),
		MaxUdpSize:                        flex.ExpandInt64Pointer(m.MaxUdpSize),
		MinimalResponses:                  flex.ExpandBoolPointer(m.MinimalResponses),
		Name:                              flex.ExpandString(m.Name),
		Notify:                            flex.ExpandBoolPointer(m.Notify),
		QueryAcl:                          flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandConfigACLItem),
		RecursionAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.RecursionAcl, diags, ExpandConfigACLItem),
		RecursionEnabled:                  flex.ExpandBoolPointer(m.RecursionEnabled),
		SortList:                          flex.ExpandFrameworkListNestedBlock(ctx, m.SortList, diags, ExpandConfigSortListItem),
		SynthesizeAddressRecordsFromHttps: flex.ExpandBoolPointer(m.SynthesizeAddressRecordsFromHttps),
		Tags:                              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransferAcl:                       flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandConfigACLItem),
		UpdateAcl:                         flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandConfigACLItem),
		UseForwardersForSubzones:          flex.ExpandBoolPointer(m.UseForwardersForSubzones),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
		ZoneAuthority: ExpandConfigZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	return to
}

func FlattenConfigView(ctx context.Context, from *dns_config.ConfigView, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigViewAttrTypes)
	}
	m := ConfigViewModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigViewAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigViewModel) Flatten(ctx context.Context, from *dns_config.ConfigView, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigViewModel{}
	}
	m.AddEdnsOptionInOutgoingQuery = types.BoolPointerValue(from.AddEdnsOptionInOutgoingQuery)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, ConfigRootNSAttrTypes, diags, FlattenConfigRootNS)
	m.CustomRootNsEnabled = types.BoolPointerValue(from.CustomRootNsEnabled)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DnssecEnableValidation = types.BoolPointerValue(from.DnssecEnableValidation)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecRootKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecRootKeys, ConfigTrustAnchorAttrTypes, diags, FlattenConfigTrustAnchor)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, ConfigTrustAnchorAttrTypes, diags, FlattenConfigTrustAnchor)
	m.DnssecValidateExpiry = types.BoolPointerValue(from.DnssecValidateExpiry)
	m.DtcConfig = FlattenConfigDTCConfig(ctx, from.DtcConfig, diags)
	m.EcsEnabled = types.BoolPointerValue(from.EcsEnabled)
	m.EcsForwarding = types.BoolPointerValue(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64(int64(*from.EcsPrefixV4))
	m.EcsPrefixV6 = flex.FlattenInt64(int64(*from.EcsPrefixV6))
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ConfigECSZoneAttrTypes, diags, FlattenConfigECSZone)
	m.EdnsUdpSize = flex.FlattenInt64(int64(*from.EdnsUdpSize))
	m.FilterAaaaAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.FilterAaaaOnV4 = flex.FlattenStringPointer(from.FilterAaaaOnV4)
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenConfigViewInheritance(ctx, from.InheritanceSources, diags)
	m.IpSpaces = flex.FlattenFrameworkListString(ctx, from.IpSpaces, diags)
	m.LameTtl = flex.FlattenInt64(int64(*from.LameTtl))
	m.MatchClientsAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchClientsAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.MatchDestinationsAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchDestinationsAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.MatchRecursiveOnly = types.BoolPointerValue(from.MatchRecursiveOnly)
	m.MaxCacheTtl = flex.FlattenInt64(int64(*from.MaxCacheTtl))
	m.MaxNegativeTtl = flex.FlattenInt64(int64(*from.MaxNegativeTtl))
	m.MaxUdpSize = flex.FlattenInt64(int64(*from.MaxUdpSize))
	m.MinimalResponses = types.BoolPointerValue(from.MinimalResponses)
	m.Name = flex.FlattenString(from.Name)
	m.Notify = types.BoolPointerValue(from.Notify)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.RecursionAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursionAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.RecursionEnabled = types.BoolPointerValue(from.RecursionEnabled)
	m.SortList = flex.FlattenFrameworkListNestedBlock(ctx, from.SortList, ConfigSortListItemAttrTypes, diags, FlattenConfigSortListItem)
	m.SynthesizeAddressRecordsFromHttps = types.BoolPointerValue(from.SynthesizeAddressRecordsFromHttps)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.UseForwardersForSubzones = types.BoolPointerValue(from.UseForwardersForSubzones)
	m.UseRootForwardersForLocalResolutionWithB1td = types.BoolPointerValue(from.UseRootForwardersForLocalResolutionWithB1td)
	m.ZoneAuthority = FlattenConfigZoneAuthority(ctx, from.ZoneAuthority, diags)
}
