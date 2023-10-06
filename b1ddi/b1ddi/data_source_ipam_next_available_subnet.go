package b1ddi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
)

func dataSourceIpamsvcNextAvailableSubnet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIpamsvcNextAvailableSubnetList,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An application specific resource identity of a resource.",
			},
			"cidr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The cidr value of subnets to be created.",
			},
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comment of next available subnets.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of next available subnets.",
			},
			"dhcp_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Reference of OnPrem Host associated with the next available subnets to be created.",
			},
			"subnet_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of subnets to generate. Default 1 if not set.",
			},
			"results": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        dataSourceSchemaFromResource(resourceIpamsvcSubnet),
				Description: "List of Subnets matching filters. The schema of each element is identical to the b1ddi_subnet resource schema.",
			},
		},
	}
}

func dataSourceIpamsvcNextAvailableSubnetList(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics
	params := &address_block.AddressBlockListNextAvailableSubnetParams{
		ID:      d.Get("id").(string),
		Cidr:    swag.Int32(int32(d.Get("cidr").(int))),
		Context: ctx,
	}
	/*name, ok := d.GetOk("name")
	if ok {
		params.Name = swag.String(name.(string))
	}
	comment, ok := d.GetOk("comment")
	if ok {
		params.Comment = swag.String(comment.(string))
	}
	dhcpHost, ok := d.GetOk("dhcp_host")
	if ok {
		params.DhcpHost = swag.String(dhcpHost.(string))
	}*/
	count, ok := d.GetOk("subnet_count")
	if ok {
		params.Count = swag.Int32(int32(count.(int)))
	}

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockListNextAvailableSubnet(params, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, space := range resp.Payload.Results {
		space.CreatedAt = &strfmt.DateTime{}
		space.UpdatedAt = &strfmt.DateTime{}
		results = append(results, flattenIpamsvcSubnet(space)...)
	}

	err = d.Set("results", results)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
