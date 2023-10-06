package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceIpamsvcRange() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcRangeRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcRange),
				Description: "List of IPAM Range matching filters. The schema of each element is identical to the b1ddi_range resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcRangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.RangeOperations.RangeList(&range_operations.RangeListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenIpamsvcRange(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIpamsvcRange(r *models.IpamsvcRange) []interface{} {
	if r == nil {
		return nil
	}

	dhcpOptions := make([]interface{}, 0, len(r.DhcpOptions))
	for _, dhcpOption := range r.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}

	exclusionRanges := make([]interface{}, 0, len(r.ExclusionRanges))
	for _, er := range r.ExclusionRanges {
		exclusionRanges = append(exclusionRanges, flattenIpamsvcExclusionRange(er))
	}

	inheritanceAssignedHosts := make([]interface{}, 0, len(r.InheritanceAssignedHosts))
	for _, inheritanceAssignedHost := range r.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritanceAssignedHost(inheritanceAssignedHost))
	}

	return []interface{}{
		map[string]interface{}{
			"id":                         r.ID,
			"comment":                    r.Comment,
			"created_at":                 r.CreatedAt.String(),
			"dhcp_host":                  r.DhcpHost,
			"dhcp_options":               dhcpOptions,
			"end":                        r.End,
			"exclusion_ranges":           exclusionRanges,
			"inheritance_assigned_hosts": inheritanceAssignedHosts,
			"inheritance_parent":         r.InheritanceParent,
			"inheritance_sources":        flattenIpamsvcDHCPOptionsInheritance(r.InheritanceSources),
			"name":                       r.Name,
			"parent":                     r.Parent,
			"protocol":                   r.Protocol,
			"space":                      r.Space,
			"start":                      r.Start,
			"tags":                       r.Tags,
			"threshold":                  flattenIpamsvcUtilizationThreshold(r.Threshold),
			"updated_at":                 r.UpdatedAt.String(),
			"utilization":                flattenIpamsvcUtilization(r.Utilization),
		},
	}
}
