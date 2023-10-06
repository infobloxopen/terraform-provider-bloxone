package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigExternalPrimary ExternalPrimary
//
// External DNS primary.
//
// swagger:model configExternalPrimary
func schemaConfigExternalPrimary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Required only if _type_ is _server_. IP Address of nameserver.
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Required only if _type_ is _server_. IP Address of nameserver.",
			},

			// Optional. Required only if _type_ is _server_. FQDN of nameserver.
			"fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Required only if _type_ is _server_. FQDN of nameserver.",
			},

			// The resource identifier.
			"nsg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// FQDN of nameserver in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "FQDN of nameserver in punycode.",
			},

			// Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.
			"tsig_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.",
			},

			// Optional. TSIG key.
			//
			// Error if empty while _tsig_enabled_ is _true_.
			"tsig_key": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTSIGKey(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. TSIG key.\n\nError if empty while _tsig_enabled_ is _true_.",
			},

			// Allowed values:
			// * _nsg_,
			// * _primary_.
			// Required: true
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Allowed values:\n* _nsg_,\n* _primary_.",
			},
		},
	}
}

func flattenConfigExternalPrimary(r *models.ConfigExternalPrimary) map[string]interface{} {
	if r == nil {
		return nil
	}
	return map[string]interface{}{
		"address":       r.Address,
		"fqdn":          r.Fqdn,
		"nsg":           r.Nsg,
		"protocol_fqdn": r.ProtocolFqdn,
		"tsig_enabled":  r.TsigEnabled,
		"tsig_key":      flattenConfigTSIGKey(r.TsigKey),
		"type":          r.Type,
	}
}

func expandConfigExternalPrimary(d map[string]interface{}) *models.ConfigExternalPrimary {
	if len(d) == 0 {
		return nil
	}
	return &models.ConfigExternalPrimary{
		Address:      d["address"].(string),
		Fqdn:         d["fqdn"].(string),
		Nsg:          d["nsg"].(string),
		ProtocolFqdn: d["protocol_fqdn"].(string),
		TsigEnabled:  d["tsig_enabled"].(bool),
		TsigKey:      expandConfigTSIGKey(d["tsig_key"].([]interface{})),
		Type:         swag.String(d["type"].(string)),
	}
}
