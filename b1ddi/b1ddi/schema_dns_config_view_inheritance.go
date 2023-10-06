package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigViewInheritance ViewInheritance
//
// Inheritance configuration specifies how and which fields _View_ object inherits from [ _Global_, _Server_ ] parent.
//
// swagger:model configViewInheritance
func schemaConfigViewInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _custom_root_ns_block_ field from _View_ object.
			"custom_root_ns_block": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedCustomRootNSBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _custom_root_ns_block_ field from _View_ object.",
			},

			// Optional. Field config for _dnssec_validation_block_ field from _View_ object.
			"dnssec_validation_block": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedDNSSECValidationBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _dnssec_validation_block_ field from _View_ object.",
			},

			// Optional. Field config for _ecs_block_ field from _View_ object.
			"ecs_block": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedECSBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _ecs_block_ field from _View_ object.",
			},

			// Optional. Field config for _edns_udp_size_ field from [View] object.
			"edns_udp_size": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _edns_udp_size_ field from [View] object.",
			},

			// Optional. Field config for _forwarders_block_ field from _View_ object.
			"forwarders_block": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedForwardersBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _forwarders_block_ field from _View_ object.",
			},

			// Optional. Field config for _gss_tsig_enabled_ field from _View_ object.
			"gss_tsig_enabled": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _gss_tsig_enabled_ field from _View_ object.",
			},

			// Optional. Field config for _lame_ttl_ field from _View_ object.
			"lame_ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _lame_ttl_ field from _View_ object.",
			},

			// Optional. Field config for _match_recursive_only_ field from _View_ object.
			"match_recursive_only": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _match_recursive_only_ field from _View_ object.",
			},

			// Optional. Field config for _max_cache_ttl_ field from _View_ object.
			"max_cache_ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _max_cache_ttl_ field from _View_ object.",
			},

			// Optional. Field config for _max_negative_ttl_ field from _View_ object.
			"max_negative_ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _max_negative_ttl_ field from _View_ object.",
			},

			// Optional. Field config for _max_udp_size_ field from [View] object.
			"max_udp_size": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _max_udp_size_ field from [View] object.",
			},

			// Optional. Field config for _minimal_responses_ field from _View_ object.
			"minimal_responses": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _minimal_responses_ field from _View_ object.",
			},

			// Field config for _notify_ field from _View_ object.
			"notify": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Field config for _notify_ field from _View_ object.",
			},

			// Optional. Field config for _query_acl_ field from _View_ object.
			"query_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _query_acl_ field from _View_ object.",
			},

			// Optional. Field config for _recursion_acl_ field from _View_ object.
			"recursion_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _recursion_acl_ field from _View_ object.",
			},

			// Optional. Field config for _recursion_enabled_ field from _View_ object.
			"recursion_enabled": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _recursion_enabled_ field from _View_ object.",
			},

			// Optional. Field config for _transfer_acl_ field from _View_ object.
			"transfer_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _transfer_acl_ field from _View_ object.",
			},

			// Optional. Field config for _update_acl_ field from _View_ object.
			"update_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _update_acl_ field from _View_ object.",
			},

			// Optional. Field config for _use_forwarders_for_subzones_ field from _View_ object.
			"use_forwarders_for_subzones": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _use_forwarders_for_subzones_ field from _View_ object.",
			},

			// Optional. Field config for _zone_authority_ field from _View_ object.
			"zone_authority": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedZoneAuthority(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _zone_authority_ field from _View_ object.",
			},
		},
	}
}

func flattenConfigViewInheritance(r *models.ConfigViewInheritance) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"custom_root_ns_block":        flattenConfigInheritedCustomRootNSBlock(r.CustomRootNsBlock),
			"dnssec_validation_block":     flattenConfigInheritedDNSSECValidationBlock(r.DnssecValidationBlock),
			"ecs_block":                   flattenConfigInheritedECSBlock(r.EcsBlock),
			"edns_udp_size":               flattenInheritanceInheritedUInt32(r.EdnsUDPSize),
			"forwarders_block":            flattenConfigInheritedForwardersBlock(r.ForwardersBlock),
			"gss_tsig_enabled":            flattenInheritanceInheritedBool(r.GssTsigEnabled),
			"lame_ttl":                    flattenInheritanceInheritedUInt32(r.LameTTL),
			"match_recursive_only":        flattenInheritanceInheritedBool(r.MatchRecursiveOnly),
			"max_cache_ttl":               flattenInheritanceInheritedUInt32(r.MaxCacheTTL),
			"max_negative_ttl":            flattenInheritanceInheritedUInt32(r.MaxNegativeTTL),
			"max_udp_size":                flattenInheritanceInheritedUInt32(r.MaxUDPSize),
			"minimal_responses":           flattenInheritanceInheritedBool(r.MinimalResponses),
			"notify":                      flattenInheritanceInheritedBool(r.Notify),
			"query_acl":                   flattenConfigInheritedACLItems(r.QueryACL),
			"recursion_acl":               flattenConfigInheritedACLItems(r.RecursionACL),
			"recursion_enabled":           flattenInheritanceInheritedBool(r.RecursionEnabled),
			"transfer_acl":                flattenConfigInheritedACLItems(r.TransferACL),
			"update_acl":                  flattenConfigInheritedACLItems(r.UpdateACL),
			"use_forwarders_for_subzones": flattenInheritanceInheritedBool(r.UseForwardersForSubzones),
			"zone_authority":              flattenConfigInheritedZoneAuthority(r.ZoneAuthority),
		},
	}
}

func expandConfigViewInheritance(d []interface{}) *models.ConfigViewInheritance {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigViewInheritance{
		CustomRootNsBlock:        expandConfigInheritedCustomRootNSBlock(in["custom_root_ns_block"].([]interface{})),
		DnssecValidationBlock:    expandConfigInheritedDNSSECValidationBlock(in["dnssec_validation_block"].([]interface{})),
		EcsBlock:                 expandConfigInheritedECSBlock(in["ecs_block"].([]interface{})),
		EdnsUDPSize:              expandInheritance2InheritedUInt32(in["edns_udp_size"].([]interface{})),
		ForwardersBlock:          expandConfigInheritedForwardersBlock(in["forwarders_block"].([]interface{})),
		GssTsigEnabled:           expandInheritance2InheritedBool(in["gss_tsig_enabled"].([]interface{})),
		LameTTL:                  expandInheritance2InheritedUInt32(in["lame_ttl"].([]interface{})),
		MatchRecursiveOnly:       expandInheritance2InheritedBool(in["match_recursive_only"].([]interface{})),
		MaxCacheTTL:              expandInheritance2InheritedUInt32(in["max_cache_ttl"].([]interface{})),
		MaxNegativeTTL:           expandInheritance2InheritedUInt32(in["max_negative_ttl"].([]interface{})),
		MaxUDPSize:               expandInheritance2InheritedUInt32(in["max_udp_size"].([]interface{})),
		MinimalResponses:         expandInheritance2InheritedBool(in["minimal_responses"].([]interface{})),
		Notify:                   expandInheritance2InheritedBool(in["notify"].([]interface{})),
		QueryACL:                 expandConfigInheritedACLItems(in["query_acl"].([]interface{})),
		RecursionACL:             expandConfigInheritedACLItems(in["recursion_acl"].([]interface{})),
		RecursionEnabled:         expandInheritance2InheritedBool(in["recursion_enabled"].([]interface{})),
		TransferACL:              expandConfigInheritedACLItems(in["transfer_acl"].([]interface{})),
		UpdateACL:                expandConfigInheritedACLItems(in["update_acl"].([]interface{})),
		UseForwardersForSubzones: expandInheritance2InheritedBool(in["use_forwarders_for_subzones"].([]interface{})),
		ZoneAuthority:            expandConfigInheritedZoneAuthority(in["zone_authority"].([]interface{})),
	}
}
