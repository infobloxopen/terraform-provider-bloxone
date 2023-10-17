package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigTrustAnchor TrustAnchor
//
// DNSSEC trust anchor.
//
// swagger:model configTrustAnchor
func schemaConfigTrustAnchor() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Key algorithm.
			// Algorithm values are as per standards.
			// The mapping is as follows:
			//  * _RSAMD5_ = 1,
			//  * _DH_ = 2,
			//  * _DSA_ = 3,
			//  * _RSASHA1_ = 5,
			//  * _DSANSEC3SHA1_ = 6,
			//  * _RSASHA1NSEC3SHA1_ = 7,
			//  * _RSASHA256_ = 8,
			//  * _RSASHA512_ = 10,
			//  * _ECDSAP256SHA256_ = 13,
			//  * _ECDSAP384SHA384_ = 14.
			// Below algorithms are deprecated and not supported anymore
			//  * _RSAMD5_ = 1,
			//  * _DSA_ = 3,
			//  * _DSANSEC3SHA1_ = 6,
			// Required: true
			"algorithm": {
				Type:     schema.TypeInt,
				Required: true,
			},

			// Zone FQDN in punycode.
			// Read Only: true
			"protocol_zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone FQDN in punycode.",
			},

			// DNSSEC key data. Non-empty, valid base64 string.
			// Required: true
			"public_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DNSSEC key data. Non-empty, valid base64 string.",
			},

			// Optional. Secure Entry Point flag.
			//
			// Defaults to _true_.
			"sep": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Secure Entry Point flag.\n\nDefaults to _true_.",
			},

			// Zone FQDN.
			// Required: true
			"zone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone FQDN.",
			},
		},
	}
}

func flattenConfigTrustAnchor(r *models.ConfigTrustAnchor) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"algorithm":     r.Algorithm,
		"protocol_zone": r.ProtocolZone,
		"public_key":    r.PublicKey,
		"sep":           r.Sep,
		"zone":          r.Zone,
	}
}

func expandConfigTrustAnchor(d map[string]interface{}) *models.ConfigTrustAnchor {
	if len(d) == 0 {
		return nil
	}

	return &models.ConfigTrustAnchor{
		Algorithm: swag.Int64(int64(d["algorithm"].(int))),
		PublicKey: swag.String(d["public_key"].(string)),
		Sep:       d["sep"].(bool),
		Zone:      swag.String(d["zone"].(string)),
	}
}
