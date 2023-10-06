package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcOptionItem OptionItem
//
// An item (_dhcp/option_item_) in a list of DHCP options. May be either a specific option or a group of options.
//
// swagger:model ipamsvcOptionItem
func schemaIpamsvcOptionItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The resource identifier.
			"group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			"option_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The option value.
			"option_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The option value.",
			},

			// The type of item.
			//
			// Valid values are:
			// * _group_
			// * _option_
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of item.\n\nValid values are:\n* _group_\n* _option_",
			},
		},
	}
}

func flattenIpamsvcOptionItem(r *models.IpamsvcOptionItem) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"group":        r.Group,
		"option_code":  r.OptionCode,
		"option_value": r.OptionValue,
		"type":         r.Type,
	}
}

func expandIpamsvcOptionItem(d map[string]interface{}) *models.IpamsvcOptionItem {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.IpamsvcOptionItem{
		Group:       d["group"].(string),
		OptionCode:  d["option_code"].(string),
		OptionValue: d["option_value"].(string),
		Type:        d["type"].(string),
	}
}
