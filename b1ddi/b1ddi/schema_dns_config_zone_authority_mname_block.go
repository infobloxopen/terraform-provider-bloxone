package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigZoneAuthorityMNameBlock ZoneAuthorityMNameBlock
//
// Block for fields: _mname_, _protocol_mname_, _use_default_mname_.
//
// swagger:model configZoneAuthorityMNameBlock
func schemaConfigZoneAuthorityMNameBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Master name server (partially qualified domain name)
			//
			// Defaults to empty.
			"mname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defaults to empty.",
			},

			// Optional. Master name server in punycode.
			//
			// Defaults to empty.
			// Read Only: true
			"protocol_mname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional. Master name server in punycode.\n\nDefaults to empty.",
			},

			// Optional. Use default value for master name server.
			//
			// Defaults to true.
			"use_default_mname": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Use default value for master name server.\n\nDefaults to true.",
			},
		},
	}
}
func flattenConfigZoneAuthorityMNameBlock(r *models.ConfigZoneAuthorityMNameBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"mname":             r.Mname,
			"protocol_mname":    r.ProtocolMname,
			"use_default_mname": r.UseDefaultMname,
		},
	}
}
