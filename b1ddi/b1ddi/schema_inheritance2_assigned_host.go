package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// Inheritance2AssignedHost AssignedHost
//
// _ddi/assigned_host_ represents a BloxOne DDI host assigned to an object.
//
// swagger:model inheritance2AssignedHost
func schemaInheritance2AssignedHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The human-readable display name for the host referred to by _ophid_.
			// Read Only: true
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable display name for the host referred to by _ophid_.",
			},

			// The resource identifier.
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The on-prem host ID.
			// Read Only: true
			"ophid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The on-prem host ID.",
			},
		},
	}
}

func flattenInheritance2AssignedHost(r *models.Inheritance2AssignedHost) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"display_name": r.DisplayName,
			"host":         r.Host,
			"ophid":        r.Ophid,
		},
	}
}
