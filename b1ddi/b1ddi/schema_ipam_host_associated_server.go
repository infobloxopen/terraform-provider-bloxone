package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcHostAssociatedServer ipamsvc host associated server
//
// swagger:model ipamsvcHostAssociatedServer
func schemaIpamsvcHostAssociatedServer() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The DHCP Config Profile name.
			// Read Only: true
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP Config Profile name.",
			},
		},
	}
}

func flattenIpamsvcHostAssociatedServer(r *models.IpamsvcHostAssociatedServer) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"id":   r.ID,
			"name": r.Name,
		},
	}
}
