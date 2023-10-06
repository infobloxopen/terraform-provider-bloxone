package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceIpamsvcAddressBlock() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcAddressBlockRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcAddressBlock),
				Description: "List of Address Blocks matching filters. The schema of each element is identical to the b1ddi_address_block resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcAddressBlockRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockList(&address_block.AddressBlockListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenIpamsvcAddressBlock(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIpamsvcAddressBlock(r *models.IpamsvcAddressBlock) []interface{} {
	if r == nil {
		return nil
	}

	dhcpOptions := make([]interface{}, 0, len(r.DhcpOptions))
	for _, dhcpOption := range r.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                           r.ID,
			"address":                      r.Address,
			"asm_config":                   flattenIpamsvcASMConfig(r.AsmConfig),
			"asm_scope_flag":               r.AsmScopeFlag,
			"cidr":                         r.Cidr,
			"comment":                      r.Comment,
			"created_at":                   r.CreatedAt.String(),
			"ddns_client_update":           r.DdnsClientUpdate,
			"ddns_domain":                  r.DdnsDomain,
			"ddns_generate_name":           r.DdnsGenerateName,
			"ddns_generated_prefix":        r.DdnsGeneratedPrefix,
			"ddns_send_updates":            r.DdnsSendUpdates,
			"ddns_update_on_renew":         r.DdnsUpdateOnRenew,
			"ddns_use_conflict_resolution": r.DdnsUseConflictResolution,
			"dhcp_config":                  flattenIpamsvcDHCPConfig(r.DhcpConfig),
			"dhcp_options":                 dhcpOptions,
			"dhcp_utilization":             flattenIpamsvcDHCPUtilization(r.DhcpUtilization),
			"header_option_filename":       r.HeaderOptionFilename,
			"header_option_server_address": r.HeaderOptionServerAddress,
			"header_option_server_name":    r.HeaderOptionServerName,
			"hostname_rewrite_char":        r.HostnameRewriteChar,
			"hostname_rewrite_enabled":     r.HostnameRewriteEnabled,
			"hostname_rewrite_regex":       r.HostnameRewriteRegex,
			"inheritance_parent":           r.InheritanceParent,
			"inheritance_sources":          flattenIpamsvcDHCPInheritance(r.InheritanceSources),
			"name":                         r.Name,
			"parent":                       r.Parent,
			"protocol":                     r.Protocol,
			"space":                        r.Space,
			"tags":                         r.Tags,
			"threshold":                    flattenIpamsvcUtilizationThreshold(r.Threshold),
			"updated_at":                   r.UpdatedAt.String(),
			"utilization":                  flattenIpamsvcUtilization(r.Utilization),
		},
	}

}
