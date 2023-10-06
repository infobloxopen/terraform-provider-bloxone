package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInheritedKerberosKeys InheritedKerberosKeys
//
// Inheritance configuration for a field of type list of _kerberos_key_.
//
// swagger:model configInheritedKerberosKeys
func schemaConfigInheritedKerberosKeys() *schema.Resource {
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
				Elem:        schemaConfigKerberosKey(),
				Computed:    true,
				Description: "Inherited value.",
			},
		},
	}
}

func flattenConfigInheritedKerberosKeys(r *models.ConfigInheritedKerberosKeys) []interface{} {
	if r == nil {
		return nil
	}

	value := make([]map[string]interface{}, 0, len(r.Value))
	for _, v := range r.Value {
		value = append(value, flattenConfigKerberosKey(v))
	}

	return []interface{}{
		map[string]interface{}{
			"action":       r.Action,
			"display_name": r.DisplayName,
			"source":       r.Source,
			"value":        value,
		},
	}
}
