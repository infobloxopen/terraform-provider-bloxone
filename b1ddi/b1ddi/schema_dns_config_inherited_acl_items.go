package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedACLItems InheritedACLItems
//
// Inheritance configuration for a field of type list of _ACLItem_.
//
// swagger:model configInheritedACLItems
func schemaConfigInheritedACLItems() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Inheritance setting for a field.
			// Defaults to _inherit_.
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Inheritance setting for a field.\nDefaults to _inherit_.",
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
				Elem:        schemaConfigACLItem(),
				Computed:    true,
				Description: "Inherited value.",
			},
		},
	}
}

func flattenConfigInheritedACLItems(r *models.ConfigInheritedACLItems) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	values := make([]interface{}, 0, len(r.Value))
	for _, aclItem := range r.Value {
		values = append(values, flattenConfigACLItem(aclItem))
	}

	return []interface{}{
		map[string]interface{}{
			"action":       r.Action,
			"display_name": r.DisplayName,
			"source":       r.Source,
			"value":        values,
		},
	}
}

func expandConfigInheritedACLItems(d []interface{}) *models.ConfigInheritedACLItems {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	values := make([]*models.ConfigACLItem, 0)
	for _, aclItem := range in["value"].([]interface{}) {
		if aclItem != nil {
			values = append(values, expandConfigACLItem(aclItem.(map[string]interface{})))
		}
	}

	return &models.ConfigInheritedACLItems{
		Action:      in["action"].(string),
		DisplayName: in["display_name"].(string),
		Source:      in["source"].(string),
		Value:       values,
	}
}
