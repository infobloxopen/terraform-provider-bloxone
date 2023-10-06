package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceIpamsvcAddress() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcAddressRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcAddress),
				Description: "List of Addresses matching filters. The schema of each element is identical to the b1ddi_address resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.IPAddressManagementAPI.Address.AddressList(&address.AddressListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenIpamsvcAddress(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenIpamsvcAddress(r *models.IpamsvcAddress) []interface{} {
	if r == nil {
		return nil
	}

	names := make([]interface{}, 0, len(r.Names))
	for _, n := range r.Names {
		names = append(names, flattenIpamsvcName(n))
	}

	usage := make([]interface{}, 0, len(r.Usage))
	for _, u := range r.Usage {
		usage = append(usage, u)
	}

	return []interface{}{
		map[string]interface{}{
			"id":         r.ID,
			"address":    r.Address,
			"comment":    r.Comment,
			"created_at": r.CreatedAt.String(),
			"dhcp_info":  flattenIpamsvcDHCPInfo(r.DhcpInfo),
			"host":       r.Host,
			"hwaddr":     r.Hwaddr,
			"interface":  r.Interface,
			"names":      names,
			"parent":     r.Parent,
			"protocol":   r.Protocol,
			"range":      r.Range,
			"space":      r.Space,
			"state":      r.State,
			"tags":       r.Tags,
			"updated_at": r.UpdatedAt.String(),
			"usage":      usage,
		},
	}
}
