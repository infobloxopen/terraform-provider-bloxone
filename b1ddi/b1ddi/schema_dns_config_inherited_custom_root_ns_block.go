package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedCustomRootNSBlock InheritedCustomRootNSBlock
//
// Inheritance block for fields: _custom_root_ns_enabled_, _custom_root_ns_.
//
// swagger:model configInheritedCustomRootNSBlock
func schemaConfigInheritedCustomRootNSBlock() *schema.Resource {
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
				Elem:        schemaConfigCustomRootNSBlock(),
				Computed:    true,
				Description: "Inherited value.",
			},
		},
	}
}

func flattenConfigInheritedCustomRootNSBlock(r *models.ConfigInheritedCustomRootNSBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"action":       r.Action,
			"display_name": r.DisplayName,
			"source":       r.Source,
			"value":        flattenConfigCustomRootNSBlock(r.Value),
		},
	}
}

func expandConfigInheritedCustomRootNSBlock(d []interface{}) *models.ConfigInheritedCustomRootNSBlock {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigInheritedCustomRootNSBlock{
		Action: in["action"].(string),
		Source: in["source"].(string),
	}
}
