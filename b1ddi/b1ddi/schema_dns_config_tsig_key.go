package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigTSIGKey TSIGKey
//
// Object representing TSIG key synced from Keys Service.
//
// swagger:model configTSIGKey
func schemaConfigTSIGKey() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// TSIG key algorithm.
			//
			// Possible values:
			//  * _hmac_sha256_,
			//  * _hmac_sha1_,
			//  * _hmac_sha224_,
			//  * _hmac_sha384_,
			//  * _hmac_sha512_.
			"algorithm": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TSIG key algorithm.\n\nPossible values:\n * _hmac_sha256_,\n * _hmac_sha1_,\n * _hmac_sha224_,\n * _hmac_sha384_,\n * _hmac_sha512_.",
			},

			// Comment for TSIG key.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment for TSIG key.",
			},

			// The resource identifier.
			// Required: true
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// TSIG key name, FQDN.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TSIG key name, FQDN.",
			},

			// TSIG key name in punycode.
			// Read Only: true
			"protocol_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TSIG key name in punycode.",
			},

			// TSIG key secret, base64 string.
			"secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TSIG key secret, base64 string.",
			},
		},
	}
}

func flattenConfigTSIGKey(r *models.ConfigTSIGKey) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"algorithm":     r.Algorithm,
			"comment":       r.Comment,
			"key":           r.Key,
			"name":          r.Name,
			"protocol_name": r.ProtocolName,
			"secret":        r.Secret,
		},
	}
}

func expandConfigTSIGKey(d []interface{}) *models.ConfigTSIGKey {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigTSIGKey{
		Algorithm: in["algorithm"].(string),
		Comment:   in["comment"].(string),
		Key:       swag.String(in["key"].(string)),
		Name:      in["name"].(string),
		Secret:    in["secret"].(string),
	}
}
