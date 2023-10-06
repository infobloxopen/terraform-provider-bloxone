package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigCustomRootNSBlock CustomRootNSBlock
//
// Block for fields: _custom_root_ns_enabled_, _custom_root_ns_.
//
// swagger:model configCustomRootNSBlock
func schemaConfigCustomRootNSBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Optional. Field config for _custom_root_ns_ field.
			"custom_root_ns": {
				Type:        schema.TypeList,
				Elem:        schemaConfigRootNS(),
				Optional:    true,
				Description: "Optional. Field config for _custom_root_ns_ field.",
			},

			// Optional. Field config for _custom_root_ns_enabled_ field.
			"custom_root_ns_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Optional. Field config for _custom_root_ns_enabled_ field.",
			},
		},
	}
}

func flattenConfigCustomRootNSBlock(r *models.ConfigCustomRootNSBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	customRootNs := make([]map[string]interface{}, 0, len(r.CustomRootNs))
	for _, ns := range r.CustomRootNs {
		customRootNs = append(customRootNs, flattenConfigRootNS(ns))
	}

	return []interface{}{
		map[string]interface{}{
			"custom_root_ns":         customRootNs,
			"custom_root_ns_enabled": r.CustomRootNsEnabled,
		},
	}
}
