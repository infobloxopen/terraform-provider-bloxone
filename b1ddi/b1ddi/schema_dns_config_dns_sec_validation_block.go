package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigDNSSECValidationBlock DNSSECValidationBlock
//
// Block for fields: _dnssec_enabled_, _dnssec_enable_validation_, _dnssec_validate_expiry_, _dnssec_trust_anchors_.
//
// swagger:model configDNSSECValidationBlock
func schemaConfigDNSSECValidationBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _dnssec_enable_validation_ field.
			"dnssec_enable_validation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _dnssec_enable_validation_ field.",
			},

			// Optional. Field config for _dnssec_enabled_ field.
			"dnssec_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _dnssec_enabled_ field.",
			},

			// Optional. Field config for _dnssec_trust_anchors_ field.
			"dnssec_trust_anchors": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTrustAnchor(),
				Optional:    true,
				Description: "Optional. Field config for _dnssec_trust_anchors_ field.",
			},

			// Optional. Field config for _dnssec_validate_expiry_ field.
			"dnssec_validate_expiry": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _dnssec_validate_expiry_ field.",
			},
		},
	}
}

func flattenConfigDNSSECValidationBlock(r *models.ConfigDNSSECValidationBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	dnssecTrustAnchors := make([]interface{}, 0, len(r.DnssecTrustAnchors))
	for _, ta := range r.DnssecTrustAnchors {
		dnssecTrustAnchors = append(dnssecTrustAnchors, flattenConfigTrustAnchor(ta))
	}

	return []interface{}{
		map[string]interface{}{
			"dnssec_enable_validation": r.DnssecEnableValidation,
			"dnssec_enabled":           r.DnssecEnabled,
			"dnssec_trust_anchors":     dnssecTrustAnchors,
			"dnssec_validate_expiry":   r.DnssecValidateExpiry,
		},
	}
}
