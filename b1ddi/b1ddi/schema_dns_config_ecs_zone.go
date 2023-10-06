package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigECSZone ECSZone
//
// EDNS Client Subnet zone.
//
// swagger:model configECSZone
func schemaConfigECSZone() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Access control for zone.
			//
			// Allowed values:
			// * _allow_,
			// * _deny_.
			// Required: true
			"access": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access control for zone.\n\nAllowed values:\n* _allow_,\n* _deny_.",
			},

			// Zone FQDN.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone FQDN.",
			},

			// Zone FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone FQDN in punycode.",
			},
		},
	}
}

func flattenConfigECSZone(r *models.ConfigECSZone) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"access":        r.Access,
		"fqdn":          r.Fqdn,
		"protocol_fqdn": r.ProtocolFqdn,
	}
}

func expandConfigECSZone(d map[string]interface{}) *models.ConfigECSZone {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.ConfigECSZone{
		Access: swag.String(d["access"].(string)),
		Fqdn:   swag.String(d["fqdn"].(string)),
	}
}
