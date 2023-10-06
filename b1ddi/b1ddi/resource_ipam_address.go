package b1ddi

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcAddress Address
//
// An __Address__ object (_ipam/address_) represents any single IP address within a given IP space.
//
// swagger:model ipamsvcAddress
func resourceIpamsvcAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcAddressCreate,
		ReadContext:   resourceIpamsvcAddressRead,
		UpdateContext: resourceIpamsvcAddressUpdate,
		DeleteContext: resourceIpamsvcAddressDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The address in form "a.b.c.d".
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The address in form \"a.b.c.d\".",
			},

			// The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The DHCP information associated with this object.
			// Read Only: true
			"dhcp_info": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPInfo(),
				Computed:    true,
				Description: "The DHCP information associated with this object.",
			},

			// The resource identifier.
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The hardware address associated with this IP address.
			"hwaddr": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The hardware address associated with this IP address.",
			},

			// The name of the network interface card (NIC) associated with the address, if any.
			"interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the network interface card (NIC) associated with the address, if any.",
			},

			// The list of all names associated with this address.
			"names": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcName(),
				Optional:    true,
				Description: "The list of all names associated with this address.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier. Can be used to allocate the next available ip for the address object.",
			},

			// The type of protocol (_ipv4_ or _ipv6_).
			// Read Only: true
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of protocol (_ipv4_ or _ipv6_).",
			},

			// The resource identifier.
			"range": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The state of the address (_used_ or _free_).
			// Read Only: true
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the address (_used_ or _free_).",
			},

			// The tags for this address in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for this address in JSON format.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:
			//  usage indicator        | description
			//  ---------------------- | --------------------------------
			//  _IPAM_                 |  Address was created by the IPAM component.
			//  _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.
			//  _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.
			//  _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.
			//  _DHCP_                 |  Address was created by the DHCP component.
			//  _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.
			//  _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.
			//  _DNS_                  |  Address is used by one or more DNS records.
			// Read Only: true
			"usage": {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Computed:    true,
				Description: "The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:\n usage indicator        | description\n ---------------------- | --------------------------------\n _IPAM_                 |  Address was created by the IPAM component.\n _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.\n _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.\n _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.\n _DHCP_                 |  Address was created by the DHCP component.\n _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.\n _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.\n _DNS_                  |  Address is used by one or more DNS records.",
			},
		},
	}
}

func resourceIpamsvcAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	names := make([]*models.IpamsvcName, 0)
	for _, n := range d.Get("names").([]interface{}) {
		if n != nil {
			names = append(names, expandIpamsvcName(n.(map[string]interface{})))
		}
	}

	if d.Get("address").(string) != "" {

		a := &models.IpamsvcAddress{
			Address:   swag.String(d.Get("address").(string)),
			Comment:   d.Get("comment").(string),
			Host:      d.Get("host").(string),
			Hwaddr:    d.Get("hwaddr").(string),
			Interface: d.Get("interface").(string),
			Names:     names,
			Parent:    d.Get("parent").(string),
			Range:     d.Get("range").(string),
			Space:     swag.String(d.Get("space").(string)),
			Tags:      d.Get("tags"),
		}

		resp, err := c.IPAddressManagementAPI.Address.AddressCreate(&address.AddressCreateParams{Body: a, Context: ctx}, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(resp.Payload.Result.ID)
	} else {
		resp, err := c.IPAddressManagementAPI.Subnet.SubnetCreateNextAvailableIP(&subnet.SubnetCreateNextAvailableIPParams{
			ID:      d.Get("parent").(string),
			Context: ctx,
		}, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		body := &models.IpamsvcAddress{
			Address:   resp.Payload.Results[0].Address,
			Comment:   d.Get("comment").(string),
			Host:      d.Get("host").(string),
			Hwaddr:    d.Get("hwaddr").(string),
			Interface: d.Get("interface").(string),
			Names:     names,
			Range:     d.Get("range").(string),
			Tags:      d.Get("tags"),
		}

		respUpd, err := c.IPAddressManagementAPI.Address.AddressUpdate(
			&address.AddressUpdateParams{ID: resp.Payload.Results[0].ID, Body: body, Context: ctx},
			nil,
		)
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(respUpd.Payload.Result.ID)
	}

	return append(diags, resourceIpamsvcAddressRead(ctx, d, m)...)
}

func resourceIpamsvcAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.Address.AddressRead(
		&address.AddressReadParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("address", resp.Payload.Result.Address)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("comment", resp.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("created_at", resp.Payload.Result.CreatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dhcp_info", flattenIpamsvcDHCPInfo(resp.Payload.Result.DhcpInfo))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("host", resp.Payload.Result.Host)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hwaddr", resp.Payload.Result.Hwaddr)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("interface", resp.Payload.Result.Interface)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	names := make([]map[string]interface{}, 0, len(resp.Payload.Result.Names))
	for _, n := range resp.Payload.Result.Names {
		names = append(names, flattenIpamsvcName(n))
	}
	err = d.Set("names", names)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("parent", resp.Payload.Result.Parent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("protocol", resp.Payload.Result.Protocol)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("range", resp.Payload.Result.Range)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("space", resp.Payload.Result.Space)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("state", resp.Payload.Result.State)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("updated_at", resp.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	usage := make([]interface{}, 0, len(resp.Payload.Result.Usage))
	for _, u := range resp.Payload.Result.Usage {
		usage = append(usage, u)
	}
	err = d.Set("usage", usage)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamsvcAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	if d.HasChange("space") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'space' field is not allowed"))
	}

	names := make([]*models.IpamsvcName, 0)
	for _, n := range d.Get("names").([]interface{}) {
		if n != nil {
			names = append(names, expandIpamsvcName(n.(map[string]interface{})))
		}
	}

	body := &models.IpamsvcAddress{
		Address:   swag.String(d.Get("address").(string)),
		Comment:   d.Get("comment").(string),
		Host:      d.Get("host").(string),
		Hwaddr:    d.Get("hwaddr").(string),
		Interface: d.Get("interface").(string),
		Names:     names,
		Range:     d.Get("range").(string),
		Tags:      d.Get("tags"),
	}

	resp, err := c.IPAddressManagementAPI.Address.AddressUpdate(
		&address.AddressUpdateParams{ID: d.Id(), Body: body, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcAddressRead(ctx, d, m)
}

func resourceIpamsvcAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	_, err := c.IPAddressManagementAPI.Address.AddressDelete(
		&address.AddressDeleteParams{ID: d.Id(), Context: ctx},
		nil,
	)
	if err != nil {
		if err.Error() == errRecordNotFound {
			diags = append(diags, diag.Diagnostic{Severity: diag.Warning, Summary: err.Error()})
		} else {
			return diag.FromErr(err)
		}
	}

	d.SetId("")

	return diags
}
