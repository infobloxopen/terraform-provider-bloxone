package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedZoneAuthority InheritedZoneAuthority
//
// Inheritance configuration for a field of type _ZoneAuthority_.
//
// swagger:model configInheritedZoneAuthority
func schemaConfigInheritedZoneAuthority() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _default_ttl_ field from _ZoneAuthority_ object.
			"default_ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _default_ttl_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _expire_ field from _ZoneAuthority_ object.
			"expire": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _expire_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _mname_ block from _ZoneAuthority_ object.
			"mname_block": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedZoneAuthorityMNameBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _mname_ block from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _negative_ttl_ field from _ZoneAuthority_ object.
			"negative_ttl": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _negative_ttl_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _protocol_rname_ field from _ZoneAuthority_ object.
			"protocol_rname": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _protocol_rname_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _refresh_ field from _ZoneAuthority_ object.
			"refresh": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _refresh_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _retry_ field from _ZoneAuthority_ object.
			"retry": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedUInt32(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _retry_ field from _ZoneAuthority_ object.",
			},

			// Optional. Field config for _rname_ field from _ZoneAuthority_ object.
			"rname": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _rname_ field from _ZoneAuthority_ object.",
			},
		},
	}
}

func flattenConfigInheritedZoneAuthority(r *models.ConfigInheritedZoneAuthority) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"default_ttl":    flattenInheritanceInheritedUInt32(r.DefaultTTL),
			"expire":         flattenInheritanceInheritedUInt32(r.Expire),
			"mname_block":    flattenConfigInheritedZoneAuthorityMNameBlock(r.MnameBlock),
			"negative_ttl":   flattenInheritanceInheritedUInt32(r.NegativeTTL),
			"protocol_rname": flattenInheritanceInheritedString(r.ProtocolRname),
			"refresh":        flattenInheritanceInheritedUInt32(r.Refresh),
			"retry":          flattenInheritanceInheritedUInt32(r.Retry),
			"rname":          flattenInheritanceInheritedString(r.Rname),
		},
	}
}

func expandConfigInheritedZoneAuthority(d []interface{}) *models.ConfigInheritedZoneAuthority {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigInheritedZoneAuthority{
		DefaultTTL:    expandInheritance2InheritedUInt32(in["default_ttl"].([]interface{})),
		Expire:        expandInheritance2InheritedUInt32(in["expire"].([]interface{})),
		MnameBlock:    expandConfigInheritedZoneAuthorityMNameBlock(in["mname_block"].([]interface{})),
		NegativeTTL:   expandInheritance2InheritedUInt32(in["negative_ttl"].([]interface{})),
		ProtocolRname: expandInheritance2InheritedString(in["protocol_rname"].([]interface{})),
		Refresh:       expandInheritance2InheritedUInt32(in["refresh"].([]interface{})),
		Retry:         expandInheritance2InheritedUInt32(in["retry"].([]interface{})),
		Rname:         expandInheritance2InheritedString(in["rname"].([]interface{})),
	}
}
