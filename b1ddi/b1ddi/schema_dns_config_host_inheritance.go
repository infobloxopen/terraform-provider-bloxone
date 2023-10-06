package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigHostInheritance HostInheritance
//
// Inheritance configuration specifies how and which fields _Host_ object inherits from _Global_ or _Server_ parent.
//
// swagger:model configHostInheritance
func schemaConfigHostInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// Optional. Field config for _kerberos_keys_ field from _Host_ object.
			"kerberos_keys": {
				Type:        schema.TypeList,
				Elem:        schemaConfigInheritedKerberosKeys(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. Field config for _kerberos_keys_ field from _Host_ object.",
			},
		},
	}
}

func flattenConfigHostInheritance(r *models.ConfigHostInheritance) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"kerberos_keys": flattenConfigInheritedKerberosKeys(r.KerberosKeys),
		},
	}
}
