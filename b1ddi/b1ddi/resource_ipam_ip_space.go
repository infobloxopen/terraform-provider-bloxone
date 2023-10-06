package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	b1models "github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcIPSpace IPSpace
//
// An __IPSpace__ object (_ipam/ip_space_) allows customers to represent their entire managed address space with no collision. A collision arises when two or more block of addresses overlap partially or fully.
//
// swagger:model ipamsvcIPSpace
func resourceIpamsvcIPSpace() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcIPSpaceCreate,
		ReadContext:   resourceIpamsvcIPSpaceRead,
		UpdateContext: resourceIpamsvcIPSpaceUpdate,
		DeleteContext: resourceIpamsvcIPSpaceDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The Automated Scope Management configuration for the IP space.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The Automated Scope Management configuration for the IP space.",
			},

			// The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.
			// Read Only: true
			"asm_scope_flag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.",
			},

			// The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// Controls who does the DDNS updates.
			//
			// Valid values are:
			// * _client_: DHCP server updates DNS if requested by client.
			// * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _ignore_: DHCP server always updates DNS, even if the client says not to.
			// * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.
			//
			// Defaults to _client_.
			"ddns_client_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Controls who does the DDNS updates.\n\nValid values are:\n* _client_: DHCP server updates DNS if requested by client.\n* _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _ignore_: DHCP server always updates DNS, even if the client says not to.\n* _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n\nDefaults to _client_.",
				ValidateFunc: validation.StringInSlice(
					[]string{"client", "server", "ignore", "over_client_update", "over_no_update"},
					false,
				),
			},

			// The domain suffix for DDNS updates. FQDN, may be empty.
			//
			// Defaults to empty.
			"ddns_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The domain suffix for DDNS updates. FQDN, may be empty.\n\nDefaults to empty.",
			},

			// Indicates if DDNS needs to generate a hostname when not supplied by the client.
			//
			// Defaults to _false_.
			"ddns_generate_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Indicates if DDNS needs to generate a hostname when not supplied by the client.\n\nDefaults to _false_.",
			},

			// The prefix used in the generation of an FQDN.
			//
			// When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
			// where address-text is simply the lease IP address converted to a hyphenated string.
			//
			// Defaults to "myhost".
			"ddns_generated_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The prefix used in the generation of an FQDN.\n\nWhen generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].\nwhere address-text is simply the lease IP address converted to a hyphenated string.\n\nDefaults to \"myhost\".",
			},

			// Determines if DDNS updates are enabled at the IP space level.
			// Defaults to _true_.
			"ddns_send_updates": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Determines if DDNS updates are enabled at the IP space level.\nDefaults to _true_.",
			},

			// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
			//
			// Defaults to _false_.
			"ddns_update_on_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.\n\nDefaults to _false_.",
			},

			// When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.
			//
			// When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.
			//
			// Defaults to _true_.
			"ddns_use_conflict_resolution": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.\n\nWhen false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.\n\nDefaults to _true_.",
			},

			// The shared DHCP configuration for the IP space that controls how leases are issued.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The shared DHCP configuration for the IP space that controls how leases are issued.",
			},

			// The list of DHCP options for the IP space. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options for the IP space. May be either a specific option or a group of options.",
			},

			// The configuration for header option filename field.
			"header_option_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option filename field.",
			},

			// The configuration for header option server address field.
			"header_option_server_address": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
				Description:  "The configuration for header option server address field.",
			},

			// The configuration for header option server name field.
			"header_option_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server name field.",
			},

			// The character to replace non-matching characters with, when hostname rewrite is enabled.
			//
			// Any single ASCII character.
			//
			// Defaults to "_".
			"hostname_rewrite_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The character to replace non-matching characters with, when hostname rewrite is enabled.\n\nAny single ASCII character.\n\nDefaults to \"_\".",
				Default:     "_",
			},

			// Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
			//
			// Defaults to _false_.
			"hostname_rewrite_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.\n\nDefaults to _false_.",
			},

			// The regex bracket expression to match valid characters.
			//
			// Must begin with "[" and end with "]" and be a compilable POSIX regex.
			//
			// Defaults to "[^a-zA-Z0-9_.]".
			"hostname_rewrite_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The regex bracket expression to match valid characters.\n\nMust begin with \"[\" and end with \"]\" and be a compilable POSIX regex.\n\nDefaults to \"[^a-zA-Z0-9_.]\".",
				Default:     "[^a-zA-Z0-9_.]",
			},

			// The inheritance configuration.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcIPSpaceInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration.",
			},

			// The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.
			// Required: true
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.",
			},

			// The tags for the IP space in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the IP space in JSON format.",
			},

			// The utilization threshold settings for the IP space.
			// Read Only: true
			"threshold": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilizationThreshold(),
				Computed:    true,
				Description: "The utilization threshold settings for the IP space.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The utilization of IP addresses in the IP space.
			// Read Only: true
			"utilization": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilization(),
				Computed:    true,
				Description: "The utilization of IP addresses in the IP space.",
			},

			// The resource identifier.
			"vendor_specific_option_option_space": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},
		},
	}
}

func resourceIpamsvcIPSpaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	dhcpOptions := make([]*b1models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	inheritanceSources, err := expandIpamsvcIPSpaceInheritance(ctx, d.Get("inheritance_sources").([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'inheritance_sources' field. The underlying expand function returned an error.")
		return diag.FromErr(err)
	}

	s := &b1models.IpamsvcIPSpace{
		AsmConfig:                       expandIpamsvcASMConfig(d.Get("asm_config").([]interface{})),
		Comment:                         d.Get("comment").(string),
		DdnsClientUpdate:                d.Get("ddns_client_update").(string),
		DdnsDomain:                      d.Get("ddns_domain").(string),
		DdnsGenerateName:                d.Get("ddns_generate_name").(bool),
		DdnsGeneratedPrefix:             d.Get("ddns_generated_prefix").(string),
		DdnsSendUpdates:                 swag.Bool(d.Get("ddns_send_updates").(bool)),
		DdnsUpdateOnRenew:               d.Get("ddns_update_on_renew").(bool),
		DdnsUseConflictResolution:       swag.Bool(d.Get("ddns_use_conflict_resolution").(bool)),
		DhcpConfig:                      expandIpamsvcDHCPConfig(d.Get("dhcp_config").([]interface{})),
		DhcpOptions:                     dhcpOptions,
		HeaderOptionFilename:            d.Get("header_option_filename").(string),
		HeaderOptionServerAddress:       d.Get("header_option_server_address").(string),
		HeaderOptionServerName:          d.Get("header_option_server_name").(string),
		HostnameRewriteChar:             d.Get("hostname_rewrite_char").(string),
		HostnameRewriteEnabled:          d.Get("hostname_rewrite_enabled").(bool),
		HostnameRewriteRegex:            d.Get("hostname_rewrite_regex").(string),
		InheritanceSources:              inheritanceSources,
		Name:                            swag.String(d.Get("name").(string)),
		Tags:                            d.Get("tags"),
		VendorSpecificOptionOptionSpace: d.Get("vendor_specific_option_option_space").(string),
	}

	resp, err := c.IPAddressManagementAPI.IPSpace.IPSpaceCreate(&ip_space.IPSpaceCreateParams{Body: s, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcIPSpaceRead(ctx, d, m)
}

func resourceIpamsvcIPSpaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	s, err := c.IPAddressManagementAPI.IPSpace.IPSpaceRead(
		&ip_space.IPSpaceReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("asm_config", flattenIpamsvcASMConfig(s.Payload.Result.AsmConfig))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("asm_scope_flag", s.Payload.Result.AsmScopeFlag)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("comment", s.Payload.Result.Comment)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("created_at", s.Payload.Result.CreatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_client_update", s.Payload.Result.DdnsClientUpdate)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_domain", s.Payload.Result.DdnsDomain)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_generate_name", s.Payload.Result.DdnsGenerateName)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_generated_prefix", s.Payload.Result.DdnsGeneratedPrefix)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_send_updates", s.Payload.Result.DdnsSendUpdates)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_update_on_renew", s.Payload.Result.DdnsUpdateOnRenew)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("ddns_use_conflict_resolution", s.Payload.Result.DdnsUseConflictResolution)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("dhcp_config", flattenIpamsvcDHCPConfig(s.Payload.Result.DhcpConfig))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	dhcpOptions := make([]map[string]interface{}, 0, len(s.Payload.Result.DhcpOptions))
	for _, dhcpOption := range s.Payload.Result.DhcpOptions {
		dhcpOptions = append(dhcpOptions, flattenIpamsvcOptionItem(dhcpOption))
	}
	err = d.Set("dhcp_options", dhcpOptions)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("header_option_filename", s.Payload.Result.HeaderOptionFilename)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("header_option_server_address", s.Payload.Result.HeaderOptionServerAddress)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("header_option_server_name", s.Payload.Result.HeaderOptionServerName)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("hostname_rewrite_char", s.Payload.Result.HostnameRewriteChar)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("hostname_rewrite_enabled", s.Payload.Result.HostnameRewriteEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("hostname_rewrite_regex", s.Payload.Result.HostnameRewriteRegex)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("inheritance_sources", flattenIpamsvcIPSpaceInheritance(s.Payload.Result.InheritanceSources))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("name", s.Payload.Result.Name)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("tags", s.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("threshold", flattenIpamsvcUtilizationThreshold(s.Payload.Result.Threshold))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("updated_at", s.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("utilization", flattenIpamsvcUtilization(s.Payload.Result.Utilization))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	err = d.Set("vendor_specific_option_option_space", s.Payload.Result.VendorSpecificOptionOptionSpace)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamsvcIPSpaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	dhcpOptions := make([]*b1models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	inheritanceSources, err := expandIpamsvcIPSpaceInheritance(ctx, d.Get("inheritance_sources").([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'inheritance_sources' field. The underlying expand function returned an error.")
		return diag.FromErr(err)
	}

	body := &b1models.IpamsvcIPSpace{
		AsmConfig:                       expandIpamsvcASMConfig(d.Get("asm_config").([]interface{})),
		Comment:                         d.Get("comment").(string),
		DdnsClientUpdate:                d.Get("ddns_client_update").(string),
		DdnsDomain:                      d.Get("ddns_domain").(string),
		DdnsGenerateName:                d.Get("ddns_generate_name").(bool),
		DdnsGeneratedPrefix:             d.Get("ddns_generated_prefix").(string),
		DdnsSendUpdates:                 swag.Bool(d.Get("ddns_send_updates").(bool)),
		DdnsUpdateOnRenew:               d.Get("ddns_update_on_renew").(bool),
		DdnsUseConflictResolution:       swag.Bool(d.Get("ddns_use_conflict_resolution").(bool)),
		DhcpConfig:                      expandIpamsvcDHCPConfig(d.Get("dhcp_config").([]interface{})),
		DhcpOptions:                     dhcpOptions,
		HeaderOptionFilename:            d.Get("header_option_filename").(string),
		HeaderOptionServerAddress:       d.Get("header_option_server_address").(string),
		HeaderOptionServerName:          d.Get("header_option_server_name").(string),
		HostnameRewriteChar:             d.Get("hostname_rewrite_char").(string),
		HostnameRewriteEnabled:          d.Get("hostname_rewrite_enabled").(bool),
		HostnameRewriteRegex:            d.Get("hostname_rewrite_regex").(string),
		InheritanceSources:              inheritanceSources,
		Name:                            swag.String(d.Get("name").(string)),
		Tags:                            d.Get("tags"),
		VendorSpecificOptionOptionSpace: d.Get("vendor_specific_option_option_space").(string),
	}

	resp, err := c.IPAddressManagementAPI.IPSpace.IPSpaceUpdate(
		&ip_space.IPSpaceUpdateParams{ID: d.Id(), Body: body, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcIPSpaceRead(ctx, d, m)
}

func resourceIpamsvcIPSpaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	ipSpaceID := d.Id()

	_, err := c.IPAddressManagementAPI.IPSpace.IPSpaceDelete(&ip_space.IPSpaceDeleteParams{ID: ipSpaceID, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
