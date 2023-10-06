package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigRootNS RootNS
//
// Root nameserver
//
// swagger:model configRootNS
func schemaConfigRootNS() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// IPv4 address.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IPv4 address.",
			},

			// FQDN.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN.",
			},

			// FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "FQDN in punycode.",
			},
		},
	}
}

func flattenConfigRootNS(r *models.ConfigRootNS) map[string]interface{} {
	if r == nil {
		return nil
	}
	return map[string]interface{}{
		"address":       r.Address,
		"fqdn":          r.Fqdn,
		"protocol_fqdn": r.ProtocolFqdn,
	}
}

func expandConfigRootNS(d map[string]interface{}) *models.ConfigRootNS {
	if d == nil || len(d) == 0 {
		return nil
	}
	return &models.ConfigRootNS{
		Address: swag.String(d["address"].(string)),
		Fqdn:    swag.String(d["fqdn"].(string)),
	}
}
