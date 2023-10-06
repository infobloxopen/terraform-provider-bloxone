package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_nsg"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceConfigForwardNSG() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigForwardNSGRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceConfigForwardNSG),
				Description: "List of DNS Forward NSGs matching filters. The schema of each element is identical to the b1ddi_dns_auth_nsg resource schema.",
			},
		},
	}
}

func dataSourceConfigForwardNSGRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.DNSConfigurationAPI.ForwardNsg.ForwardNsgList(
		&forward_nsg.ForwardNsgListParams{
			Filter:  swag.String(filterStr),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, nsg := range resp.Payload.Results {
		results = append(results, flattenConfigForwardNSG(nsg)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenConfigForwardNSG(r *models.ConfigForwardNSG) []interface{} {
	if r == nil {
		return nil
	}

	externalForwarders := make([]map[string]interface{}, 0, len(r.ExternalForwarders))
	for _, ef := range r.ExternalForwarders {
		externalForwarders = append(externalForwarders, flattenConfigForwarder(ef))
	}

	return []interface{}{
		map[string]interface{}{

			"comment": r.Comment,

			"external_forwarders": externalForwarders,

			"forwarders_only": r.ForwardersOnly,

			"hosts": r.Hosts,

			"id": r.ID,

			"internal_forwarders": r.InternalForwarders,

			"name": r.Name,

			"nsgs": r.Nsgs,

			"tags": r.Tags,
		},
	}
}
