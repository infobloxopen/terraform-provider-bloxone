package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedDNSSECValidationBlock InheritedDNSSECValidationBlock
//
// Inheritance block for fields: _dnssec_enabled_, _dnssec_enable_validation_, _dnssec_validate_expiry_, _dnssec_trust_anchors_.
//
// swagger:model configInheritedDNSSECValidationBlock
func schemaConfigInheritedDNSSECValidationBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Defaults to _inherit_.
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Defaults to _inherit_.",
			},

			// Human-readable display name for the object referred to by _source_.
			// Read Only: true
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable display name for the object referred to by _source_.",
			},

			// The resource identifier.
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Inherited value.
			// Read Only: true
			"value": {
				Type:        schema.TypeList,
				Elem:        schemaConfigDNSSECValidationBlock(),
				Computed:    true,
				Description: "Inherited value.",
			},
		},
	}
}

func flattenConfigInheritedDNSSECValidationBlock(r *models.ConfigInheritedDNSSECValidationBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"action":       r.Action,
			"display_name": r.DisplayName,
			"source":       r.Source,
			"value":        flattenConfigDNSSECValidationBlock(r.Value),
		},
	}
}

func expandConfigInheritedDNSSECValidationBlock(d []interface{}) *models.ConfigInheritedDNSSECValidationBlock {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})
	return &models.ConfigInheritedDNSSECValidationBlock{
		Action: in["action"].(string),
		Source: in["source"].(string),
	}
}
