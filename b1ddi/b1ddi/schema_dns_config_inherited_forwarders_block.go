package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedForwardersBlock InheritedForwardersBlock
//
// Inheritance block for fields: _forwarders_, _forwarders_only_.
//
// swagger:model configInheritedForwardersBlock
func schemaConfigInheritedForwardersBlock() *schema.Resource {
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
				Elem:        schemaConfigForwardersBlock(),
				Computed:    true,
				Description: "Inherited value.",
			},
		},
	}
}

func flattenConfigInheritedForwardersBlock(r *models.ConfigInheritedForwardersBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"action":       r.Action,
			"display_name": r.DisplayName,
			"source":       r.Source,
			"value":        flattenConfigForwardersBlock(r.Value),
		},
	}
}

func expandConfigInheritedForwardersBlock(d []interface{}) *models.ConfigInheritedForwardersBlock {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.ConfigInheritedForwardersBlock{
		Action: in["action"].(string),
		Source: in["source"].(string),
	}
}
