package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"github.com/infobloxopen/b1ddi-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcFixedAddress FixedAddress
//
// A __FixedAddress__ object (_dhcp/fixed_address_) reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so it can match that client.
//
// swagger:model ipamsvcFixedAddress
func resourceIpamsvcFixedAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcFixedAddressCreate,
		ReadContext:   resourceIpamsvcFixedAddressRead,
		UpdateContext: resourceIpamsvcFixedAddressUpdate,
		DeleteContext: resourceIpamsvcFixedAddressDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The reserved address.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The reserved address.",
			},

			// The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// The list of DHCP options. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options. May be either a specific option or a group of options.",
			},

			// The configuration for header option filename field.
			"header_option_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option filename field.",
			},

			// The configuration for header option server address field.
			"header_option_server_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server address field.",
			},

			// The configuration for header option server name field.
			"header_option_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server name field.",
			},

			// The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			"inheritance_assigned_hosts": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceAssignedHost(),
				Computed:    true,
				Description: "The list of the inheritance assigned hosts of the object.",
			},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcFixedAddressInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration.",
			},

			// The resource identifier.
			"ip_space": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// Indicates how to match the client:
			//  * _mac_: match the client MAC address,
			//  * _client_text_ or _client_hex_: match the client identifier,
			//  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82).
			// Required: true
			"match_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Indicates how to match the client:\n * _mac_: match the client MAC address,\n * _client_text_ or _client_hex_: match the client identifier,\n * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82).",
			},

			// The value to match.
			// Required: true
			"match_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value to match.",
			},

			// The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The resource identifier.",
			},

			// The tags for the fixed address in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the fixed address in JSON format.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},
		},
	}
}

func resourceIpamsvcFixedAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	fa := &models.IpamsvcFixedAddress{
		Address:                   swag.String(d.Get("address").(string)),
		Comment:                   d.Get("comment").(string),
		DhcpOptions:               dhcpOptions,
		HeaderOptionFilename:      d.Get("header_option_filename").(string),
		HeaderOptionServerAddress: d.Get("header_option_server_address").(string),
		HeaderOptionServerName:    d.Get("header_option_server_name").(string),
		Hostname:                  d.Get("hostname").(string),
		InheritanceParent:         d.Get("inheritance_parent").(string),
		InheritanceSources:        expandIpamsvcFixedAddressInheritance(d.Get("inheritance_sources").([]interface{})),
		IPSpace:                   d.Get("ip_space").(string),
		MatchType:                 swag.String(d.Get("match_type").(string)),
		MatchValue:                swag.String(d.Get("match_value").(string)),
		Name:                      d.Get("name").(string),
		Parent:                    d.Get("parent").(string),
		Tags:                      d.Get("tags"),
	}

	resp, err := c.IPAddressManagementAPI.FixedAddress.FixedAddressCreate(&fixed_address.FixedAddressCreateParams{Body: fa, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	resourceIpamsvcFixedAddressRead(ctx, d, m)

	return diags
}

func resourceIpamsvcFixedAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.FixedAddress.FixedAddressRead(&fixed_address.FixedAddressReadParams{ID: d.Id(), Context: ctx}, nil)
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
	dhcpOptions := make([]map[string]interface{}, 0, len(resp.Payload.Result.DhcpOptions))
	for _, dhcpOption := range resp.Payload.Result.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}
	err = d.Set("dhcp_options", dhcpOptions)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("header_option_filename", resp.Payload.Result.HeaderOptionFilename)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("header_option_server_address", resp.Payload.Result.HeaderOptionServerAddress)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("header_option_server_name", resp.Payload.Result.HeaderOptionServerName)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hostname", resp.Payload.Result.Hostname)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	inheritanceAssignedHosts := make([]interface{}, 0, len(resp.Payload.Result.InheritanceAssignedHosts))
	for _, inheritanceAssignedHost := range resp.Payload.Result.InheritanceAssignedHosts {
		inheritanceAssignedHosts = append(inheritanceAssignedHosts, flattenInheritanceAssignedHost(inheritanceAssignedHost))
	}
	err = d.Set("inheritance_assigned_hosts", inheritanceAssignedHosts)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_parent", resp.Payload.Result.InheritanceParent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_sources", flattenIpamsvcFixedAddressInheritance(resp.Payload.Result.InheritanceSources))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ip_space", resp.Payload.Result.IPSpace)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("match_type", resp.Payload.Result.MatchType)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("match_value", resp.Payload.Result.MatchValue)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("name", resp.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("parent", resp.Payload.Result.Parent)
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

	return diags
}

func resourceIpamsvcFixedAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	fa := &models.IpamsvcFixedAddress{
		Address:                   swag.String(d.Get("address").(string)),
		Comment:                   d.Get("comment").(string),
		DhcpOptions:               dhcpOptions,
		HeaderOptionFilename:      d.Get("header_option_filename").(string),
		HeaderOptionServerAddress: d.Get("header_option_server_address").(string),
		HeaderOptionServerName:    d.Get("header_option_server_name").(string),
		Hostname:                  d.Get("hostname").(string),
		InheritanceSources:        expandIpamsvcFixedAddressInheritance(d.Get("inheritance_sources").([]interface{})),
		IPSpace:                   d.Get("ip_space").(string),
		MatchType:                 swag.String(d.Get("match_type").(string)),
		MatchValue:                swag.String(d.Get("match_value").(string)),
		Name:                      d.Get("name").(string),
		Tags:                      d.Get("tags"),
	}

	resp, err := c.IPAddressManagementAPI.FixedAddress.FixedAddressUpdate(
		&fixed_address.FixedAddressUpdateParams{ID: d.Id(), Body: fa, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcFixedAddressRead(ctx, d, m)
}

func resourceIpamsvcFixedAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	_, err := c.IPAddressManagementAPI.FixedAddress.FixedAddressDelete(&fixed_address.FixedAddressDeleteParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
