package b1ddi

import (
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcDHCPConfig DHCPConfig
//
// A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.
//
// swagger:model ipamsvcDHCPConfig
func schemaIpamsvcDHCPConfig() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// Disable to allow leases only for known clients, those for which a fixed address is configured.
			"allow_unknown": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Disable to allow leases only for known clients, those for which a fixed address is configured.",
			},

			// The resource identifier.
			"filters": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The list of clients to ignore requests from.
			"ignore_list": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcIgnoreItem(),
				Optional:    true,
				Computed:    true,
				Description: "The list of clients to ignore requests from.",
			},

			// The lease duration in seconds.
			"lease_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The lease duration in seconds.",
				Default:     3600,
			},
		},
	}
}

func flattenIpamsvcDHCPConfig(r *models.IpamsvcDHCPConfig) []interface{} {
	if r == nil {
		return nil
	}

	ignoreList := make([]map[string]interface{}, 0, len(r.IgnoreList))
	for _, ii := range r.IgnoreList {
		ignoreList = append(ignoreList, flattenIpamsvcIgnoreItem(ii))
	}

	return []interface{}{
		map[string]interface{}{
			"allow_unknown": r.AllowUnknown,
			"filters":       r.Filters,
			"ignore_list":   ignoreList,
			"lease_time":    r.LeaseTime,
		},
	}
}

func expandIpamsvcDHCPConfig(d []interface{}) *models.IpamsvcDHCPConfig {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	ignoreList := make([]*models.IpamsvcIgnoreItem, 0)
	for _, ignoreItem := range in["ignore_list"].([]interface{}) {
		ignoreList = append(ignoreList, expandIpamsvcIgnoreItem(ignoreItem.(map[string]interface{})))
	}

	filters := make([]string, 0)
	for _, filter := range in["filters"].([]interface{}) {
		filters = append(filters, filter.(string))
	}

	return &models.IpamsvcDHCPConfig{
		AllowUnknown: swag.Bool(in["allow_unknown"].(bool)),
		Filters:      filters,
		IgnoreList:   ignoreList,
		LeaseTime:    int64(in["lease_time"].(int)),
	}
}
