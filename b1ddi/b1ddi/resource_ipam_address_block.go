package b1ddi

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcAddressBlock AddressBlock
//
// An __AddressBlock__ object (_ipam/address_block_) is a set of contiguous IP addresses in the same IP space with no gap, expressed as a CIDR block. Address blocks are hierarchical and may be parented to other address blocks as long as the parent block fully contains the child and no sibling overlaps. Top level address blocks are parented to an IP space.
//
// swagger:model ipamsvcAddressBlock
func resourceIpamsvcAddressBlock() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcAddressBlockCreate,
		ReadContext:   resourceIpamsvcAddressBlockRead,
		UpdateContext: resourceIpamsvcAddressBlockUpdate,
		DeleteContext: resourceIpamsvcAddressBlockDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			// The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
			},

			// The Automated Scope Management configuration for the address block.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The Automated Scope Management configuration for the address block.",
			},

			// Incremented by 1 if the IP address usage limits for automated scope management are exceeded for any subnets in the address block.
			// Read Only: true
			"asm_scope_flag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Incremented by 1 if the IP address usage limits for automated scope management are exceeded for any subnets in the address block.",
			},

			// The CIDR of the address block. This is required, if _address_ does not specify it in its input.
			// Maximum: 32
			// Minimum: 1
			"cidr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The CIDR of the address block. This is required, if _address_ does not specify it in its input.",
			},

			// The description for the address block. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the address block. May contain 0 to 1024 characters. Can include UTF-8.",
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
			},

			// The domain suffix for DDNS updates. FQDN, may be empty.
			//
			// Defaults to empty.
			"ddns_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain suffix for DDNS updates. FQDN, may be empty.\n\nDefaults to empty.",
			},

			// Indicates if DDNS needs to generate a hostname when not supplied by the client.
			//
			// Defaults to _false_.
			"ddns_generate_name": {
				Type:        schema.TypeBool,
				Optional:    true,
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

			// Determines if DDNS updates are enabled at the address block level.
			// Defaults to _true_.
			"ddns_send_updates": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Determines if DDNS updates are enabled at the address block level.\nDefaults to _true_.",
			},

			// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
			//
			// Defaults to _false_.
			"ddns_update_on_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
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

			// The shared DHCP configuration that controls how leases are issued for the address block.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The shared DHCP configuration that controls how leases are issued for the address block.",
			},

			// The list of DHCP options for the address block. May be either a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The list of DHCP options for the address block. May be either a specific option or a group of options.",
			},

			// The utilization of IP addresses within the DHCP ranges of the address block.
			// Read Only: true
			"dhcp_utilization": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPUtilization(),
				Computed:    true,
				Description: "The utilization of IP addresses within the DHCP ranges of the address block.",
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

			// The character to replace non-matching characters with, when hostname rewrite is enabled.
			//
			// Any single ASCII character.
			//
			// Defaults to "_".
			"hostname_rewrite_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The character to replace non-matching characters with, when hostname rewrite is enabled.\n\nAny single ASCII character.\n\nDefaults to \"_\".",
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
				Computed:    true,
				Description: "The regex bracket expression to match valid characters.\n\nMust begin with \"[\" and end with \"]\" and be a compilable POSIX regex.\n\nDefaults to \"[^a-zA-Z0-9_.]\".",
			},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The DHCP inheritance configuration for the address block.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP inheritance configuration for the address block.",
			},

			// The name of the address block. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the address block. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The type of protocol of address block (_ipv4_ or _ipv6_).
			// Read Only: true
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of protocol of address block (_ipv4_ or _ipv6_).",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The tags for the address block in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the address block in JSON format.",
			},

			// The IP address utilization thresholds for the address block.
			"threshold": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilizationThreshold(),
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "The IP address utilization thresholds for the address block.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The IP address utilization statistics for the address block.
			// Read Only: true
			"utilization": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilization(),
				Computed:    true,
				Description: "The IP address utilization statistics for the address block.",
			},
		},
	}
}

func resourceIpamsvcAddressBlockCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	inheritanceSources, err := expandIpamsvcDHCPInheritance(ctx, d.Get("inheritance_sources").([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'inheritance_sources' field. The underlying expand function returned an error.")
		return diag.FromErr(err)
	}

	ab := &models.IpamsvcAddressBlock{
		Address:                   swag.String(d.Get("address").(string)),
		AsmConfig:                 expandIpamsvcASMConfig(d.Get("asm_config").([]interface{})),
		Cidr:                      int64(d.Get("cidr").(int)),
		Comment:                   d.Get("comment").(string),
		DdnsClientUpdate:          d.Get("ddns_client_update").(string),
		DdnsDomain:                d.Get("ddns_domain").(string),
		DdnsGenerateName:          d.Get("ddns_generate_name").(bool),
		DdnsGeneratedPrefix:       d.Get("ddns_generated_prefix").(string),
		DdnsSendUpdates:           swag.Bool(d.Get("ddns_send_updates").(bool)),
		DdnsUpdateOnRenew:         d.Get("ddns_update_on_renew").(bool),
		DdnsUseConflictResolution: swag.Bool(d.Get("ddns_use_conflict_resolution").(bool)),
		DhcpConfig:                expandIpamsvcDHCPConfig(d.Get("dhcp_config").([]interface{})),
		DhcpOptions:               dhcpOptions,
		HeaderOptionFilename:      d.Get("header_option_filename").(string),
		HeaderOptionServerAddress: d.Get("header_option_server_address").(string),
		HeaderOptionServerName:    d.Get("header_option_server_name").(string),
		HostnameRewriteChar:       d.Get("hostname_rewrite_char").(string),
		HostnameRewriteEnabled:    d.Get("hostname_rewrite_enabled").(bool),
		HostnameRewriteRegex:      d.Get("hostname_rewrite_regex").(string),
		InheritanceParent:         d.Get("inheritance_parent").(string),
		InheritanceSources:        inheritanceSources,
		Name:                      d.Get("name").(string),
		Parent:                    d.Get("parent").(string),
		Space:                     swag.String(d.Get("space").(string)),
		Tags:                      d.Get("tags"),
		Threshold:                 expandIpamsvcUtilizationThreshold(d.Get("threshold").([]interface{})),
	}

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockCreate(&address_block.AddressBlockCreateParams{Body: ab, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcAddressBlockRead(ctx, d, m)
}

func resourceIpamsvcAddressBlockRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockRead(
		&address_block.AddressBlockReadParams{
			ID: d.Id(), Context: ctx,
		},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	err = d.Set("address", resp.Payload.Result.Address)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("asm_config", flattenIpamsvcASMConfig(resp.Payload.Result.AsmConfig))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("asm_scope_flag", resp.Payload.Result.AsmScopeFlag)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("cidr", resp.Payload.Result.Cidr)
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
	err = d.Set("ddns_client_update", resp.Payload.Result.DdnsClientUpdate)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_domain", resp.Payload.Result.DdnsDomain)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_generate_name", resp.Payload.Result.DdnsGenerateName)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_generated_prefix", resp.Payload.Result.DdnsGeneratedPrefix)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_send_updates", resp.Payload.Result.DdnsSendUpdates)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_update_on_renew", resp.Payload.Result.DdnsUpdateOnRenew)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("ddns_use_conflict_resolution", resp.Payload.Result.DdnsUseConflictResolution)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("dhcp_config", flattenIpamsvcDHCPConfig(resp.Payload.Result.DhcpConfig))
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
	err = d.Set("dhcp_utilization", flattenIpamsvcDHCPUtilization(resp.Payload.Result.DhcpUtilization))
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
	err = d.Set("hostname_rewrite_char", resp.Payload.Result.HostnameRewriteChar)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hostname_rewrite_enabled", resp.Payload.Result.HostnameRewriteEnabled)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("hostname_rewrite_regex", resp.Payload.Result.HostnameRewriteRegex)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_parent", resp.Payload.Result.InheritanceParent)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("inheritance_sources", flattenIpamsvcDHCPInheritance(resp.Payload.Result.InheritanceSources))
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
	err = d.Set("protocol", resp.Payload.Result.Protocol)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("space", resp.Payload.Result.Space)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("tags", resp.Payload.Result.Tags)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("threshold", flattenIpamsvcUtilizationThreshold(resp.Payload.Result.Threshold))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("updated_at", resp.Payload.Result.UpdatedAt.String())
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}
	err = d.Set("utilization", flattenIpamsvcUtilization(resp.Payload.Result.Utilization))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	return diags
}

func resourceIpamsvcAddressBlockUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	if d.HasChange("address") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'address' field is not allowed"))
	}

	if d.HasChange("space") {
		d.Partial(true)
		return diag.FromErr(fmt.Errorf("changing the value of 'space' field is not allowed"))
	}

	dhcpOptions := make([]*models.IpamsvcOptionItem, 0)
	for _, o := range d.Get("dhcp_options").([]interface{}) {
		if o != nil {
			dhcpOptions = append(dhcpOptions, expandIpamsvcOptionItem(o.(map[string]interface{})))
		}
	}

	inheritanceSources, err := expandIpamsvcDHCPInheritance(ctx, d.Get("inheritance_sources").([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'inheritance_sources' field. The underlying expand function returned an error.")
		return diag.FromErr(err)
	}

	ab := &models.IpamsvcAddressBlock{
		AsmConfig:                 expandIpamsvcASMConfig(d.Get("asm_config").([]interface{})),
		Cidr:                      int64(d.Get("cidr").(int)),
		Comment:                   d.Get("comment").(string),
		DdnsClientUpdate:          d.Get("ddns_client_update").(string),
		DdnsDomain:                d.Get("ddns_domain").(string),
		DdnsGenerateName:          d.Get("ddns_generate_name").(bool),
		DdnsGeneratedPrefix:       d.Get("ddns_generated_prefix").(string),
		DdnsSendUpdates:           swag.Bool(d.Get("ddns_send_updates").(bool)),
		DdnsUpdateOnRenew:         d.Get("ddns_update_on_renew").(bool),
		DdnsUseConflictResolution: swag.Bool(d.Get("ddns_use_conflict_resolution").(bool)),
		DhcpConfig:                expandIpamsvcDHCPConfig(d.Get("dhcp_config").([]interface{})),
		DhcpOptions:               dhcpOptions,
		HeaderOptionFilename:      d.Get("header_option_filename").(string),
		HeaderOptionServerAddress: d.Get("header_option_server_address").(string),
		HeaderOptionServerName:    d.Get("header_option_server_name").(string),
		HostnameRewriteChar:       d.Get("hostname_rewrite_char").(string),
		HostnameRewriteEnabled:    d.Get("hostname_rewrite_enabled").(bool),
		HostnameRewriteRegex:      d.Get("hostname_rewrite_regex").(string),
		InheritanceParent:         d.Get("inheritance_parent").(string),
		InheritanceSources:        inheritanceSources,
		Name:                      d.Get("name").(string),
		Parent:                    d.Get("parent").(string),
		Tags:                      d.Get("tags"),
		Threshold:                 expandIpamsvcUtilizationThreshold(d.Get("threshold").([]interface{})),
	}

	resp, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockUpdate(
		&address_block.AddressBlockUpdateParams{ID: d.Id(), Body: ab, Context: ctx},
		nil,
	)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	return resourceIpamsvcAddressBlockRead(ctx, d, m)
}

func resourceIpamsvcAddressBlockDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*b1ddiclient.Client)

	var diags diag.Diagnostics

	_, err := c.IPAddressManagementAPI.AddressBlock.AddressBlockDelete(&address_block.AddressBlockDeleteParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		switch err.Error() {
		case errAddressBlockNotFound, errRecordNotFound, errIncorrectUtilizationUpdateRef:
			diags = append(diags, diag.Diagnostic{Severity: diag.Warning, Summary: err.Error()})
		default:
			return diag.FromErr(err)
		}
	}

	d.SetId("")

	return nil
}
