package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigInternalSecondary InternalSecondary
//
// BloxOne DDI host acting as DNS secondary.
//
// swagger:model configInternalSecondary
func schemaConfigInternalSecondary() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The resource identifier.
			// Required: true
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func flattenConfigInternalSecondary(r *models.ConfigInternalSecondary) map[string]interface{} {
	if r == nil {
		return nil
	}
	return map[string]interface{}{
		"host": r.Host,
	}
}

func expandConfigInternalSecondary(d map[string]interface{}) *models.ConfigInternalSecondary {
	if len(d) == 0 {
		return nil
	}
	return &models.ConfigInternalSecondary{
		Host: swag.String(d["host"].(string)),
	}
}
