package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcOptionCode OptionCode
//
// An __OptionCode__ (_dhcp/option_code_) defines a DHCP option code.
//
// swagger:model ipamsvcOptionCode
func resourceIpamsvcOptionCode() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Indicates whether the option value is an array of the type or not.
			"array": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether the option value is an array of the type or not.",
			},

			// The option code.
			// Required: true
			// Maximum: 254
			// Minimum: 1
			"code": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The option code.",
			},

			// The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the option code. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the option code. Must contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			// Required: true
			"option_space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The source for the option code.
			//
			// Valid values are:
			//  * _dhcp_server_
			//  * _reserved_
			//  * _blox_one_ddi_
			//  * _customer_
			//
			// Defaults to _customer_.
			// Read Only: true
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source for the option code.\n\nValid values are:\n * _dhcp_server_\n * _reserved_\n * _blox_one_ddi_\n * _customer_\n\nDefaults to _customer_.",
			},

			// The option type for the option code.
			//
			// Valid values are:
			// * _address4_
			// * _address6_
			// * _boolean_
			// * _empty_
			// * _fqdn_
			// * _int8_
			// * _int16_
			// * _int32_
			// * _text_
			// * _uint8_
			// * _uint16_
			// * _uint32_
			// Required: true
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The option type for the option code.\n\nValid values are:\n* _address4_\n* _address6_\n* _boolean_\n* _empty_\n* _fqdn_\n* _int8_\n* _int16_\n* _int32_\n* _text_\n* _uint8_\n* _uint16_\n* _uint32_",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},
		},
	}
}

func flattenIpamsvcOptionCode(r *models.IpamsvcOptionCode) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"array":        r.Array,
			"code":         r.Code,
			"comment":      r.Comment,
			"created_at":   r.CreatedAt.String(),
			"id":           r.ID,
			"name":         r.Name,
			"option_space": r.OptionSpace,
			"source":       r.Source,
			"type":         r.Type,
			"updated_at":   r.UpdatedAt.String(),
		},
	}
}
