package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcExclusionRange ExclusionRange
//
// The __ExclusionRange__ object represents an exclusion range inside a DHCP range.
//
// swagger:model ipamsvcExclusionRange
func schemaIpamsvcExclusionRange() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// The end address of the exclusion range.
			// Required: true
			"end": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The end address of the exclusion range.",
			},

			// The start address of the exclusion range.
			// Required: true
			"start": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The start address of the exclusion range.",
			},
		},
	}
}

func flattenIpamsvcExclusionRange(r *models.IpamsvcExclusionRange) map[string]interface{} {
	if r == nil {
		return nil
	}
	return map[string]interface{}{
		"comment": r.Comment,
		"end":     r.End,
		"start":   r.Start,
	}
}

func expandIpamsvcExclusionRange(d map[string]interface{}) *models.IpamsvcExclusionRange {
	if len(d) == 0 {
		return nil
	}

	return &models.IpamsvcExclusionRange{
		Comment: d["comment"].(string),
		End:     swag.String(d["end"].(string)),
		Start:   swag.String(d["start"].(string)),
	}
}
