package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigAuthZoneInheritance config auth zone inheritance
//
// swagger:model configAuthZoneInheritance
func schemaConfigAuthZoneInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _gss_tsig_enabled_ field from _AuthZone_ object.
			"gss_tsig_enabled": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _gss_tsig_enabled_ field from _AuthZone_ object.",
			},

			// Field config for _notify_ field from _AuthZone_ object.
			"notify": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Field config for _notify_ field from _AuthZone_ object.",
			},

			// Optional. Field config for _query_acl_ field from _AuthZone_ object.
			"query_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _query_acl_ field from _AuthZone_ object.",
			},

			// Optional. Field config for _transfer_acl_ field from _AuthZone_ object.
			"transfer_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _transfer_acl_ field from _AuthZone_ object.",
			},

			// Optional. Field config for _update_acl_ field from _AuthZone_ object.
			"update_acl": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedACLItems(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _update_acl_ field from _AuthZone_ object.",
			},

			// Optional. Field config for _use_forwarders_for_subzones_ field from _AuthZone_ object.
			"use_forwarders_for_subzones": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _use_forwarders_for_subzones_ field from _AuthZone_ object.",
			},

			// Optional. Field config for _zone_authority_ field from _AuthZone_ object.
			"zone_authority": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedZoneAuthority(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _zone_authority_ field from _AuthZone_ object.",
			},
		},
	}
}

func flattenConfigAuthZoneInheritance(r *models.ConfigAuthZoneInheritance) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{}

	if r.GssTsigEnabled != nil {
		m["gss_tsig_enabled"] = flattenInheritanceInheritedBool(r.GssTsigEnabled)
	}
	if r.Notify != nil {
		m["notify"] = flattenInheritanceInheritedBool(r.Notify)
	}
	if r.QueryACL != nil {
		m["query_acl"] = flattenConfigInheritedACLItems(r.QueryACL)
	}
	if r.TransferACL != nil {
		m["transfer_acl"] = flattenConfigInheritedACLItems(r.TransferACL)
	}
	if r.UpdateACL != nil {
		m["update_acl"] = flattenConfigInheritedACLItems(r.UpdateACL)
	}
	if r.UseForwardersForSubzones != nil {
		m["use_forwarders_for_subzones"] = flattenInheritanceInheritedBool(r.UseForwardersForSubzones)
	}
	if r.ZoneAuthority != nil {
		m["zone_authority"] = flattenConfigInheritedZoneAuthority(r.ZoneAuthority)
	}
	return []interface{}{
		m,
	}
}

func expandConfigAuthZoneInheritance(d []interface{}) *models.ConfigAuthZoneInheritance {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigAuthZoneInheritance{
		GssTsigEnabled:           expandInheritance2InheritedBool(in["gss_tsig_enabled"].([]interface{})),
		Notify:                   expandInheritance2InheritedBool(in["notify"].([]interface{})),
		QueryACL:                 expandConfigInheritedACLItems(in["query_acl"].([]interface{})),
		TransferACL:              expandConfigInheritedACLItems(in["transfer_acl"].([]interface{})),
		UpdateACL:                expandConfigInheritedACLItems(in["update_acl"].([]interface{})),
		UseForwardersForSubzones: expandInheritance2InheritedBool(in["use_forwarders_for_subzones"].([]interface{})),
		ZoneAuthority:            expandConfigInheritedZoneAuthority(in["zone_authority"].([]interface{})),
	}
}
