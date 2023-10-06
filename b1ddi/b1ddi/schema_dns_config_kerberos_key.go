package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// ConfigKerberosKey KerberosKey
//
// A __KerberosKey__ object (_keys/kerberos_) represents a Kerberos key.
//
// swagger:model configKerberosKey
func schemaConfigKerberosKey() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Encryption algorithm of the key in accordance with RFC 3961.
			// Read Only: true
			"algorithm": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Encryption algorithm of the key in accordance with RFC 3961.",
			},

			// Kerberos realm of the principal.
			// Read Only: true
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kerberos realm of the principal.",
			},

			// The resource identifier.
			// Required: true
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// Kerberos principal associated with key.
			// Read Only: true
			"principal": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Kerberos principal associated with key.",
			},

			// Upload time for the key.
			// Read Only: true
			"uploaded_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Upload time for the key.",
			},

			// The version number (KVNO) of the key.
			// Read Only: true
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version number (KVNO) of the key.",
			},
		},
	}
}

func flattenConfigKerberosKey(r *models.ConfigKerberosKey) map[string]interface{} {
	if r == nil {
		return nil
	}

	return map[string]interface{}{
		"algorithm":   r.Algorithm,
		"domain":      r.Domain,
		"key":         r.Key,
		"principal":   r.Principal,
		"uploaded_at": r.UploadedAt,
		"version":     r.Version,
	}
}
