package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcHost Host
//
// A DHCP __Host__ (_dhcp/host_) object associates a DHCP Config Profile with an on-prem host.
//
// swagger:model ipamsvcHost
func schemaIpamsvcHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The resource identifier.
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The primary IP address of the on-prem host.
			// Read Only: true
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The primary IP address of the on-prem host.",
			},

			// The DHCP Config Profile for the on-prem host.
			"associated_server": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcHostAssociatedServer(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP Config Profile for the on-prem host.",
			},

			// The description for the on-prem host.
			// Read Only: true
			"comment": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description for the on-prem host.",
			},

			// Current dhcp application version of the host.
			// Read Only: true
			"current_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current dhcp application version of the host.",
			},

			// The resource identifier.
			"ip_space": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The display name of the on-prem host.
			// Read Only: true
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The display name of the on-prem host.",
			},

			// The on-prem host ID.
			// Read Only: true
			"ophid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The on-prem host ID.",
			},

			// The resource identifier.
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The tags of the on-prem host in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags of the on-prem host in JSON format.",
			},
		},
	}
}

func flattenIpamsvcHost(r *models.IpamsvcHost) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"id":                r.ID,
			"address":           r.Address,
			"associated_server": flattenIpamsvcHostAssociatedServer(r.AssociatedServer),
			"comment":           r.Comment,
			"current_version":   r.CurrentVersion,
			"ip_space":          r.IPSpace,
			"name":              r.Name,
			"ophid":             r.Ophid,
			"server":            r.Server,
			"tags":              r.Tags,
		},
	}
}
