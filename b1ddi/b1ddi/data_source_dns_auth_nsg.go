package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_nsg"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceConfigAuthNSG() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConfigAuthNSGRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceConfigAuthNSG),
				Description: "List of DNS Auth NSGs matching filters. The schema of each element is identical to the b1ddi_dns_auth_nsg resource schema.",
			},
		},
	}
}

func dataSourceConfigAuthNSGRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.DNSConfigurationAPI.AuthNsg.AuthNsgList(
		&auth_nsg.AuthNsgListParams{
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
		results = append(results, flattenConfigAuthNSG(nsg)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenConfigAuthNSG(r *models.ConfigAuthNSG) []interface{} {
	if r == nil {
		return nil
	}

	externalPrimaries := make([]map[string]interface{}, 0, len(r.ExternalPrimaries))
	for _, ep := range r.ExternalPrimaries {
		externalPrimaries = append(externalPrimaries, flattenConfigExternalPrimary(ep))
	}

	externalSecondaries := make([]map[string]interface{}, 0, len(r.ExternalSecondaries))
	for _, es := range r.ExternalSecondaries {
		externalSecondaries = append(externalSecondaries, flattenConfigExternalSecondary(es))
	}

	internalSecondaries := make([]map[string]interface{}, 0, len(r.InternalSecondaries))
	for _, is := range r.InternalSecondaries {
		internalSecondaries = append(internalSecondaries, flattenConfigInternalSecondary(is))
	}

	return []interface{}{
		map[string]interface{}{

			"comment": r.Comment,

			"external_primaries": externalPrimaries,

			"external_secondaries": externalSecondaries,

			"id": r.ID,

			"internal_secondaries": internalSecondaries,

			"name": r.Name,

			"nsgs": r.Nsgs,

			"tags": r.Tags,
		},
	}
}
