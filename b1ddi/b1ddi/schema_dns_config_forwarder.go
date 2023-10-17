package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigForwarder Forwarder
//
// External DNS server to forward to.
//
// swagger:model configForwarder
func schemaConfigForwarder() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Server IP address.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Server IP address.",
			},

			// Server FQDN.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Server FQDN.",
			},

			// Server FQDN in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server FQDN in punycode.",
			},
		},
	}
}

func flattenConfigForwarder(r *models.ConfigForwarder) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"address":       r.Address,
		"fqdn":          r.Fqdn,
		"protocol_fqdn": r.ProtocolFqdn,
	}
}

func expandConfigForwarder(d map[string]interface{}) *models.ConfigForwarder {
	if len(d) == 0 {
		return nil
	}

	return &models.ConfigForwarder{
		Address: swag.String(d["address"].(string)),
		Fqdn:    swag.String(d["fqdn"].(string)),
	}
}
