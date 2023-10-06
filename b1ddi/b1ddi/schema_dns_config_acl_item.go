package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigACLItem ACLItem
//
// Element in an ACL.
//
// Error if both _acl_ and _address_ are given.
//
// swagger:model configACLItem
func schemaConfigACLItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Access permission for _element_.
			//
			// Allowed values:
			//  * _allow_,
			//  * _deny_.
			// Required: true
			"access": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access permission for _element_.\n\nAllowed values:\n * _allow_,\n * _deny_.",
			},

			// The resource identifier.
			"acl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Optional. Data for _ip_ _element_.
			//
			// Must be empty if _element_ is not _ip_.
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional. Data for _ip_ _element_.\n\nMust be empty if _element_ is not _ip_.",
			},

			// Type of element.
			//
			// Allowed values:
			//  * _any_,
			//  * _ip_,
			//  * _acl_,
			//  * _tsig_key_.
			// Required: true
			"element": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of element.\n\nAllowed values:\n * _any_,\n * _ip_,\n * _acl_,\n * _tsig_key_.",
			},

			// Optional. TSIG key.
			//
			// Must be empty if _element_ is not _tsig_key_.
			"tsig_key": {
				Type:        schema.TypeList,
				Elem:        schemaConfigTSIGKey(),
				MaxItems:    1,
				Optional:    true,
				Description: "Optional. TSIG key.\n\nMust be empty if _element_ is not _tsig_key_.",
			},
		},
	}
}

func flattenConfigACLItem(r *models.ConfigACLItem) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"access":   r.Access,
		"acl":      r.ACL,
		"address":  r.Address,
		"element":  r.Element,
		"tsig_key": flattenConfigTSIGKey(r.TsigKey),
	}
}

func expandConfigACLItem(d map[string]interface{}) *models.ConfigACLItem {
	if d == nil || len(d) == 0 {
		return nil
	}

	return &models.ConfigACLItem{
		Access:  swag.String(d["access"].(string)),
		ACL:     d["acl"].(string),
		Address: d["address"].(string),
		Element: swag.String(d["element"].(string)),
		TsigKey: expandConfigTSIGKey(d["tsig_key"].([]interface{})),
	}
}
