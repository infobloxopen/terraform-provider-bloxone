package b1ddi

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
)

// IpamsvcNextAvailableSubnetResponse CreateNextAvailableABResponse
//
// The Next Available Subnet object create response format.
//
// swagger:model ipamsvcCreateNextAvailableSubnetResponse

func resourceIpamsvcNextAvailableSubnetResponse() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcNextAvailableSubnetResponseCreate,
		ReadContext:   resourceIpamsvcNextAvailableSubnetResponseRead,
		UpdateContext: resourceIpamsvcNextAvailableSubnetResponseUpdate,
		DeleteContext: resourceIpamsvcNextAvailableSubnetResponseDelete,
		Schema: map[string]*schema.Schema{
			"ab_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Address block ID that the subnet created has to be associated with",
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
				Elem:        resourceIpamsvcSubnet().Schema,
				Description: "List of Subnets created through ",
			},
		},
	}
}

func resourceIpamsvcNextAvailableSubnetResponseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	params := &address_block.AddressBlockCreateNextAvailableSubnetParams{
		ID:      d.Get("ab_id").(string),
		Cidr:    swag.Int32(int32(d.Get("cidr").(int))),
		Context: ctx,
	}
	name, ok := d.GetOk("name")
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
	}
	var (
		resp *address_block.AddressBlockCreateNextAvailableSubnetCreated
		err  error
	)

	count, ok := d.GetOk("subnet_count")
	if ok {
		params.Count = swag.Int32(int32(count.(int)))
	}

	resp, err = c.IPAddressManagementAPI.AddressBlock.AddressBlockCreateNextAvailableSubnet(
		params,
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}
	time.Sleep(time.Second)

	results := make([]interface{}, 0, len(resp.Payload.Results))
	for _, space := range resp.Payload.Results {
		results = append(results, flattenIpamsvcSubnet(space)...)
		fmt.Printf("Subnet ID: %s\n", space.ID)
	}
	err = d.Set("results", results)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceIpamsvcNextAvailableSubnetResponseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceIpamsvcNextAvailableSubnetResponseUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceIpamsvcNextAvailableSubnetResponseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	/*	fmt.Printf("Subnets to delete: %d\n", d.Get("subnet_count").(int))
		for i := 0; i < d.Get("subnet_count").(int); i++ {
			fmt.Printf("Deleting Subnet: %s\n", d.Id())
			_, err := c.IPAddressManagementAPI.Subnet.SubnetDelete(&subnet.SubnetDeleteParams{
				ID:      d.Id(),
				Context: ctx,
			}, nil)
			if err != nil {
				return diag.FromErr(err)
			}
		}*/
	_, err := c.IPAddressManagementAPI.Subnet.SubnetDelete(&subnet.SubnetDeleteParams{
		ID:      d.Id(),
		Context: ctx,
	}, nil)
	if err != nil {
		return nil
	}
	d.SetId("")
	return nil
}
