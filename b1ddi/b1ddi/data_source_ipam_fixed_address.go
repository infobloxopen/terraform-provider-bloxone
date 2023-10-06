package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceIpamsvcFixedAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcFixedAddressRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcFixedAddress),
				Description: "List of Fixed Addresses matching filters. The schema of each element is identical to the b1ddi_fixed_address resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcFixedAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.FixedAddress.FixedAddressList(&fixed_address.FixedAddressListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenIpamsvcFixedAddress(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIpamsvcFixedAddress(r *models.IpamsvcFixedAddress) []interface{} {
	if r == nil {
		return nil
	}

	dhcpOptions := make([]interface{}, 0, len(r.DhcpOptions))
	for _, dhcpOption := range r.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}

	inheritanceAssignedHosts := make([]interface{}, 0, len(r.InheritanceAssignedHosts))
	for _, inheritanceAssignedHost := range r.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritanceAssignedHost(inheritanceAssignedHost))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                           r.ID,
			"address":                      r.Address,
			"comment":                      r.Comment,
			"created_at":                   r.CreatedAt.String(),
			"dhcp_options":                 dhcpOptions,
			"header_option_filename":       r.HeaderOptionFilename,
			"header_option_server_address": r.HeaderOptionServerAddress,
			"header_option_server_name":    r.HeaderOptionServerName,
			"hostname":                     r.Hostname,
			"inheritance_assigned_hosts":   inheritanceAssignedHosts,
			"inheritance_parent":           r.InheritanceParent,
			"inheritance_sources":          flattenIpamsvcFixedAddressInheritance(r.InheritanceSources),
			"ip_space":                     r.IPSpace,
			"match_type":                   r.MatchType,
			"match_value":                  r.MatchValue,
			"name":                         r.Name,
			"parent":                       r.Parent,
			"tags":                         r.Tags,
			"updated_at":                   r.UpdatedAt.String(),
		},
	}
}
