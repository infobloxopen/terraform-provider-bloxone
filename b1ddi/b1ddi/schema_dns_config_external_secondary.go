package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigExternalSecondary ExternalSecondary
//
// External DNS secondary.
//
// swagger:model configExternalSecondary
func schemaConfigExternalSecondary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// IP Address of nameserver.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP Address of nameserver.",
			},

			// FQDN of nameserver.
			// Required: true
			"fqdn": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "FQDN of nameserver.",
			},

			// FQDN of nameserver in punycode.
			// Read Only: true
			"protocol_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "FQDN of nameserver in punycode.",
			},

			// If enabled, the NS record and glue record will NOT be automatically generated
			// according to secondaries nameserver assignment.
			//
			// Default: _false_
			"stealth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If enabled, the NS record and glue record will NOT be automatically generated\naccording to secondaries nameserver assignment.\n\nDefault: _false_",
			},

			// If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.
			//
			// Default: _false_
			"tsig_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.\n\nDefault: _false_",
			},

			// TSIG key.
			//
			// Error if empty while _tsig_enabled_ is _true_.
			"tsig_key": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTSIGKey(),
				MaxItems:    1,
				Optional:    true,
				Description: "TSIG key.\n\nError if empty while _tsig_enabled_ is _true_.",
			},
		},
	}
}

func flattenConfigExternalSecondary(r *models.ConfigExternalSecondary) map[string]interface{} {
	if r == nil {
		return nil
	}
	return map[string]interface{}{
		"address":       r.Address,
		"fqdn":          r.Fqdn,
		"protocol_fqdn": r.ProtocolFqdn,
		"stealth":       r.Stealth,
		"tsig_enabled":  r.TsigEnabled,
		"tsig_key":      flattenConfigTSIGKey(r.TsigKey),
	}
}

func expandConfigExternalSecondary(d map[string]interface{}) *models.ConfigExternalSecondary {
	if d == nil || len(d) == 0 {
		return nil
	}
	return &models.ConfigExternalSecondary{
		Address:      swag.String(d["address"].(string)),
		Fqdn:         swag.String(d["fqdn"].(string)),
		ProtocolFqdn: d["protocol_fqdn"].(string),
		Stealth:      d["stealth"].(bool),
		TsigEnabled:  d["tsig_enabled"].(bool),
		TsigKey:      expandConfigTSIGKey(d["tsig_key"].([]interface{})),
	}
}
