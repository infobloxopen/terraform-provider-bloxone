package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigHostAssociatedServer config host associated server
//
// swagger:model configHostAssociatedServer
func schemaConfigHostAssociatedServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// DNS server name.
			// Read Only: true
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS server name.",
			},
		},
	}
}

func flattenConfigHostAssociatedServer(r *models.ConfigHostAssociatedServer) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"name": r.Name,
		},
	}
}
