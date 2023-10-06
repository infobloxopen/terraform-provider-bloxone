package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcInheritedDHCPOptionList InheritedDHCPOptionList
//
// The inheritance configuration for a field that contains list of _OptionItem_.
//
// swagger:model ipamsvcInheritedDHCPOptionList
func schemaIpamsvcInheritedDHCPOptionList() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance setting.
			//
			// Valid values are:
			// * _inherit_: Use the inherited value.
			// * _block_: Don't use the inherited value.
			//
			// Defaults to _inherit_.
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The inheritance setting.\n\nValid values are:\n* _inherit_: Use the inherited value.\n* _block_: Don't use the inherited value.\n\nDefaults to _inherit_.",
			},

			// The inherited DHCP option values.
			"value": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPOption(),
				Optional:    true,
				Description: "The inherited DHCP option values.",
			},
		},
	}
}

func flattenIpamsvcInheritedDHCPOptionList(r *models.IpamsvcInheritedDHCPOptionList) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	values := make([]map[string]interface{}, 0, len(r.Value))
	for _, value := range r.Value {
		values = append(values, flattenIpamsvcInheritedDHCPOption(value))
	}

	return []interface{}{
		map[string]interface{}{
			"action": r.Action,
			"value":  values,
		},
	}
}

func expandIpamsvcInheritedDHCPOptionList(d []interface{}) *models.IpamsvcInheritedDHCPOptionList {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	values := make([]*models.IpamsvcInheritedDHCPOption, 0)
	for _, value := range in["value"].([]map[string]interface{}) {
		values = append(values, expandIpamsvcInheritedDHCPOption(value))
	}

	return &models.IpamsvcInheritedDHCPOptionList{
		Action: in["action"].(string),
		Value:  values,
	}
}
