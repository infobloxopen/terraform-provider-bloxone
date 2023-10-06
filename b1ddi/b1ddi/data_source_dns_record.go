package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_data/record"
	"github.com/infobloxopen/b1ddi-go-client/models"
	"strconv"
	"time"
)

func dataSourceDataRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDataRecordRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Configure a map of filters to be applied on the search result.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceDataRecord),
				Description: "List of DNS Records matching filters. The schema of each element is identical to the b1ddi_dns_record resource schema.",
			},
		},
	}
}

func dataSourceDataRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	filtersMap := d.Get("filters").(map[string]interface{})
	filterStr := filterFromMap(filtersMap)

	resp, err := c.DNSDataAPI.Record.RecordList(&record.RecordListParams{
		Filter:  swag.String(filterStr),
		Context: ctx,
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, ab := range resp.Payload.Results {
		results = append(results, flattenDataRecord(ab)...)
	}
	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenDataRecord(r *models.DataRecord) []interface{} {
	if r == nil {
		return nil
	}

	return []interface{}{
		map[string]interface{}{
			"id":                     r.ID,
			"absolute_name_spec":     r.AbsoluteNameSpec,
			"absolute_zone_name":     r.AbsoluteZoneName,
			"comment":                r.Comment,
			"created_at":             r.CreatedAt.String(),
			"delegation":             r.Delegation,
			"disabled":               r.Disabled,
			"dns_absolute_name_spec": r.DNSAbsoluteNameSpec,
			"dns_absolute_zone_name": r.DNSAbsoluteZoneName,
			"dns_name_in_zone":       r.DNSNameInZone,
			"dns_rdata":              r.DNSRdata,
			"inheritance_sources":    flattenDataRecordInheritance(r.InheritanceSources),
			"name_in_zone":           r.NameInZone,
			"options":                r.Options,
			"rdata":                  r.Rdata,
			"source":                 r.Source,
			"tags":                   r.Tags,
			"ttl":                    r.TTL,
			"type":                   r.Type,
			"updated_at":             r.UpdatedAt.String(),
			"view":                   r.View,
			"view_name":              r.ViewName,
			"zone":                   r.Zone,
		},
	}
}
