package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	b1models "github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceIpamsvcIPSpace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcIPSpaceRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcIPSpace),
				Description: "List of IP Spaces matching filters. The schema of each element is identical to the b1ddi_ip_space resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcIPSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.IPSpace.IPSpaceList(&ip_space.IPSpaceListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, space := range resp.Payload.Results {
		results = append(results, flattenIpamsvcIPSpace(space)...)
	}

	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIpamsvcIPSpace(r *b1models.IpamsvcIPSpace) []interface{} {
	if r == nil {
		return nil
	}

	dhcpOptions := make([]interface{}, 0, len(r.DhcpOptions))
	for _, dhcpOption := range r.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                                  r.ID,
			"asm_config":                          flattenIpamsvcASMConfig(r.AsmConfig),
			"asm_scope_flag":                      r.AsmScopeFlag,
			"comment":                             r.Comment,
			"created_at":                          r.CreatedAt.String(),
			"ddns_client_update":                  r.DdnsClientUpdate,
			"ddns_domain":                         r.DdnsDomain,
			"ddns_generate_name":                  r.DdnsGenerateName,
			"ddns_generated_prefix":               r.DdnsGeneratedPrefix,
			"ddns_send_updates":                   r.DdnsSendUpdates,
			"ddns_update_on_renew":                r.DdnsUpdateOnRenew,
			"ddns_use_conflict_resolution":        r.DdnsUseConflictResolution,
			"dhcp_config":                         flattenIpamsvcDHCPConfig(r.DhcpConfig),
			"dhcp_options":                        dhcpOptions,
			"header_option_filename":              r.HeaderOptionFilename,
			"header_option_server_address":        r.HeaderOptionServerAddress,
			"header_option_server_name":           r.HeaderOptionServerName,
			"hostname_rewrite_char":               r.HostnameRewriteChar,
			"hostname_rewrite_enabled":            r.HostnameRewriteEnabled,
			"hostname_rewrite_regex":              r.HostnameRewriteRegex,
			"inheritance_sources":                 flattenIpamsvcIPSpaceInheritance(r.InheritanceSources),
			"name":                                r.Name,
			"tags":                                r.Tags,
			"threshold":                           flattenIpamsvcUtilizationThreshold(r.Threshold),
			"updated_at":                          r.UpdatedAt.String(),
			"utilization":                         flattenIpamsvcUtilization(r.Utilization),
			"vendor_specific_option_option_space": r.VendorSpecificOptionOptionSpace,
		},
	}
}
