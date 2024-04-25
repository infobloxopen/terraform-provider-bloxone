package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
)

type ConfigServerInheritanceModel struct {
	AddEdnsOptionInOutgoingQuery      types.Object `tfsdk:"add_edns_option_in_outgoing_query"`
	CustomRootNsBlock                 types.Object `tfsdk:"custom_root_ns_block"`
	DnssecValidationBlock             types.Object `tfsdk:"dnssec_validation_block"`
	EcsBlock                          types.Object `tfsdk:"ecs_block"`
	FilterAaaaAcl                     types.Object `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                    types.Object `tfsdk:"filter_aaaa_on_v4"`
	ForwardersBlock                   types.Object `tfsdk:"forwarders_block"`
	GssTsigEnabled                    types.Object `tfsdk:"gss_tsig_enabled"`
	KerberosKeys                      types.Object `tfsdk:"kerberos_keys"`
	LameTtl                           types.Object `tfsdk:"lame_ttl"`
	LogQueryResponse                  types.Object `tfsdk:"log_query_response"`
	MatchRecursiveOnly                types.Object `tfsdk:"match_recursive_only"`
	MaxCacheTtl                       types.Object `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                    types.Object `tfsdk:"max_negative_ttl"`
	MinimalResponses                  types.Object `tfsdk:"minimal_responses"`
	Notify                            types.Object `tfsdk:"notify"`
	QueryAcl                          types.Object `tfsdk:"query_acl"`
	QueryPort                         types.Object `tfsdk:"query_port"`
	RecursionAcl                      types.Object `tfsdk:"recursion_acl"`
	RecursionEnabled                  types.Object `tfsdk:"recursion_enabled"`
	RecursiveClients                  types.Object `tfsdk:"recursive_clients"`
	ResolverQueryTimeout              types.Object `tfsdk:"resolver_query_timeout"`
	SecondaryAxfrQueryLimit           types.Object `tfsdk:"secondary_axfr_query_limit"`
	SecondarySoaQueryLimit            types.Object `tfsdk:"secondary_soa_query_limit"`
	SortList                          types.Object `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps types.Object `tfsdk:"synthesize_address_records_from_https"`
	TransferAcl                       types.Object `tfsdk:"transfer_acl"`
	UpdateAcl                         types.Object `tfsdk:"update_acl"`
	UseForwardersForSubzones          types.Object `tfsdk:"use_forwarders_for_subzones"`
}

var ConfigServerInheritanceAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"custom_root_ns_block":                  types.ObjectType{AttrTypes: ConfigInheritedCustomRootNSBlockAttrTypes},
	"dnssec_validation_block":               types.ObjectType{AttrTypes: ConfigInheritedDNSSECValidationBlockAttrTypes},
	"ecs_block":                             types.ObjectType{AttrTypes: ConfigInheritedECSBlockAttrTypes},
	"filter_aaaa_acl":                       types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"filter_aaaa_on_v4":                     types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
	"forwarders_block":                      types.ObjectType{AttrTypes: ConfigInheritedForwardersBlockAttrTypes},
	"gss_tsig_enabled":                      types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"kerberos_keys":                         types.ObjectType{AttrTypes: ConfigInheritedKerberosKeysAttrTypes},
	"lame_ttl":                              types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"log_query_response":                    types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"match_recursive_only":                  types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"max_cache_ttl":                         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"max_negative_ttl":                      types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"minimal_responses":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"notify":                                types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"query_acl":                             types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"query_port":                            types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"recursion_acl":                         types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"recursion_enabled":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"recursive_clients":                     types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"resolver_query_timeout":                types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"secondary_axfr_query_limit":            types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"secondary_soa_query_limit":             types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"sort_list":                             types.ObjectType{AttrTypes: ConfigInheritedSortListItemsAttrTypes},
	"synthesize_address_records_from_https": types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"transfer_acl":                          types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"update_acl":                            types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"use_forwarders_for_subzones":           types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
}

var ConfigServerInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"custom_root_ns_block": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedCustomRootNSBlockResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"dnssec_validation_block": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedDNSSECValidationBlockResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"ecs_block": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedECSBlockResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"filter_aaaa_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"filter_aaaa_on_v4": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"forwarders_block": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedForwardersBlockResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"gss_tsig_enabled": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"kerberos_keys": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedKerberosKeysResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"lame_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"log_query_response": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"match_recursive_only": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"max_cache_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"max_negative_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"minimal_responses": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"notify": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"query_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"query_port": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"recursion_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"recursion_enabled": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"recursive_clients": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"resolver_query_timeout": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"secondary_axfr_query_limit": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"secondary_soa_query_limit": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"sort_list": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedSortListItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"synthesize_address_records_from_https": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"transfer_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"update_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"use_forwarders_for_subzones": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandConfigServerInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.ServerInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigServerInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigServerInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.ServerInheritance {
	if m == nil {
		return nil
	}
	to := &dnsconfig.ServerInheritance{
		AddEdnsOptionInOutgoingQuery:      ExpandInheritance2InheritedBool(ctx, m.AddEdnsOptionInOutgoingQuery, diags),
		CustomRootNsBlock:                 ExpandConfigInheritedCustomRootNSBlock(ctx, m.CustomRootNsBlock, diags),
		DnssecValidationBlock:             ExpandConfigInheritedDNSSECValidationBlock(ctx, m.DnssecValidationBlock, diags),
		EcsBlock:                          ExpandConfigInheritedECSBlock(ctx, m.EcsBlock, diags),
		FilterAaaaAcl:                     ExpandConfigInheritedACLItems(ctx, m.FilterAaaaAcl, diags),
		FilterAaaaOnV4:                    ExpandInheritance2InheritedString(ctx, m.FilterAaaaOnV4, diags),
		ForwardersBlock:                   ExpandConfigInheritedForwardersBlock(ctx, m.ForwardersBlock, diags),
		GssTsigEnabled:                    ExpandInheritance2InheritedBool(ctx, m.GssTsigEnabled, diags),
		KerberosKeys:                      ExpandConfigInheritedKerberosKeys(ctx, m.KerberosKeys, diags),
		LameTtl:                           ExpandInheritance2InheritedUInt32(ctx, m.LameTtl, diags),
		LogQueryResponse:                  ExpandInheritance2InheritedBool(ctx, m.LogQueryResponse, diags),
		MatchRecursiveOnly:                ExpandInheritance2InheritedBool(ctx, m.MatchRecursiveOnly, diags),
		MaxCacheTtl:                       ExpandInheritance2InheritedUInt32(ctx, m.MaxCacheTtl, diags),
		MaxNegativeTtl:                    ExpandInheritance2InheritedUInt32(ctx, m.MaxNegativeTtl, diags),
		MinimalResponses:                  ExpandInheritance2InheritedBool(ctx, m.MinimalResponses, diags),
		Notify:                            ExpandInheritance2InheritedBool(ctx, m.Notify, diags),
		QueryAcl:                          ExpandConfigInheritedACLItems(ctx, m.QueryAcl, diags),
		QueryPort:                         ExpandInheritance2InheritedUInt32(ctx, m.QueryPort, diags),
		RecursionAcl:                      ExpandConfigInheritedACLItems(ctx, m.RecursionAcl, diags),
		RecursionEnabled:                  ExpandInheritance2InheritedBool(ctx, m.RecursionEnabled, diags),
		RecursiveClients:                  ExpandInheritance2InheritedUInt32(ctx, m.RecursiveClients, diags),
		ResolverQueryTimeout:              ExpandInheritance2InheritedUInt32(ctx, m.ResolverQueryTimeout, diags),
		SecondaryAxfrQueryLimit:           ExpandInheritance2InheritedUInt32(ctx, m.SecondaryAxfrQueryLimit, diags),
		SecondarySoaQueryLimit:            ExpandInheritance2InheritedUInt32(ctx, m.SecondarySoaQueryLimit, diags),
		SortList:                          ExpandConfigInheritedSortListItems(ctx, m.SortList, diags),
		SynthesizeAddressRecordsFromHttps: ExpandInheritance2InheritedBool(ctx, m.SynthesizeAddressRecordsFromHttps, diags),
		TransferAcl:                       ExpandConfigInheritedACLItems(ctx, m.TransferAcl, diags),
		UpdateAcl:                         ExpandConfigInheritedACLItems(ctx, m.UpdateAcl, diags),
		UseForwardersForSubzones:          ExpandInheritance2InheritedBool(ctx, m.UseForwardersForSubzones, diags),
	}
	return to
}

func FlattenConfigServerInheritance(ctx context.Context, from *dnsconfig.ServerInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigServerInheritanceAttrTypes)
	}
	m := ConfigServerInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigServerInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigServerInheritanceModel) Flatten(ctx context.Context, from *dnsconfig.ServerInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigServerInheritanceModel{}
	}
	m.AddEdnsOptionInOutgoingQuery = FlattenInheritance2InheritedBool(ctx, from.AddEdnsOptionInOutgoingQuery, diags)
	m.CustomRootNsBlock = FlattenConfigInheritedCustomRootNSBlock(ctx, from.CustomRootNsBlock, diags)
	m.DnssecValidationBlock = FlattenConfigInheritedDNSSECValidationBlock(ctx, from.DnssecValidationBlock, diags)
	m.EcsBlock = FlattenConfigInheritedECSBlock(ctx, from.EcsBlock, diags)
	m.FilterAaaaAcl = FlattenConfigInheritedACLItems(ctx, from.FilterAaaaAcl, diags)
	m.FilterAaaaOnV4 = FlattenInheritance2InheritedString(ctx, from.FilterAaaaOnV4, diags)
	m.ForwardersBlock = FlattenConfigInheritedForwardersBlock(ctx, from.ForwardersBlock, diags)
	m.GssTsigEnabled = FlattenInheritance2InheritedBool(ctx, from.GssTsigEnabled, diags)
	m.KerberosKeys = FlattenConfigInheritedKerberosKeys(ctx, from.KerberosKeys, diags)
	m.LameTtl = FlattenInheritance2InheritedUInt32(ctx, from.LameTtl, diags)
	m.LogQueryResponse = FlattenInheritance2InheritedBool(ctx, from.LogQueryResponse, diags)
	m.MatchRecursiveOnly = FlattenInheritance2InheritedBool(ctx, from.MatchRecursiveOnly, diags)
	m.MaxCacheTtl = FlattenInheritance2InheritedUInt32(ctx, from.MaxCacheTtl, diags)
	m.MaxNegativeTtl = FlattenInheritance2InheritedUInt32(ctx, from.MaxNegativeTtl, diags)
	m.MinimalResponses = FlattenInheritance2InheritedBool(ctx, from.MinimalResponses, diags)
	m.Notify = FlattenInheritance2InheritedBool(ctx, from.Notify, diags)
	m.QueryAcl = FlattenConfigInheritedACLItems(ctx, from.QueryAcl, diags)
	m.QueryPort = FlattenInheritance2InheritedUInt32(ctx, from.QueryPort, diags)
	m.RecursionAcl = FlattenConfigInheritedACLItems(ctx, from.RecursionAcl, diags)
	m.RecursionEnabled = FlattenInheritance2InheritedBool(ctx, from.RecursionEnabled, diags)
	m.RecursiveClients = FlattenInheritance2InheritedUInt32(ctx, from.RecursiveClients, diags)
	m.ResolverQueryTimeout = FlattenInheritance2InheritedUInt32(ctx, from.ResolverQueryTimeout, diags)
	m.SecondaryAxfrQueryLimit = FlattenInheritance2InheritedUInt32(ctx, from.SecondaryAxfrQueryLimit, diags)
	m.SecondarySoaQueryLimit = FlattenInheritance2InheritedUInt32(ctx, from.SecondarySoaQueryLimit, diags)
	m.SortList = FlattenConfigInheritedSortListItems(ctx, from.SortList, diags)
	m.SynthesizeAddressRecordsFromHttps = FlattenInheritance2InheritedBool(ctx, from.SynthesizeAddressRecordsFromHttps, diags)
	m.TransferAcl = FlattenConfigInheritedACLItems(ctx, from.TransferAcl, diags)
	m.UpdateAcl = FlattenConfigInheritedACLItems(ctx, from.UpdateAcl, diags)
	m.UseForwardersForSubzones = FlattenInheritance2InheritedBool(ctx, from.UseForwardersForSubzones, diags)
}
