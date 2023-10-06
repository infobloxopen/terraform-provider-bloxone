package b1ddi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcIPSpaceInheritance IPSpaceInheritance
//
// The __IPSpaceInheritance__ object specifies how and which fields _IPSpace_ object inherits from the parent.
//
// swagger:model ipamsvcIPSpaceInheritance
func schemaIpamsvcIPSpaceInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance configuration for _asm_config_ field.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _asm_config_ field.",
			},

			// The inheritance configuration for _ddns_client_update_ field from _IPSpace_ object.
			"ddns_client_update": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_client_update_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'.
			"ddns_enabled": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'.",
			},

			// The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields from _IPSpace_ object.
			"ddns_hostname_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDDNSHostnameBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields from _IPSpace_ object.
			"ddns_update_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDDNSUpdateBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_update_on_renew_ field from _IPSpace_ object.
			"ddns_update_on_renew": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_update_on_renew_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _ddns_use_conflict_resolution_ field from _IPSpace_ object.
			"ddns_use_conflict_resolution": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedBool(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _ddns_use_conflict_resolution_ field from _IPSpace_ object.",
			},

			// The inheritance configuration for _dhcp_config_ field.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _dhcp_config_ field.",
			},

			// The inheritance configuration for _dhcp_options_ field.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPOptionList(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _dhcp_options_ field.",
			},

			// The inheritance configuration for _header_option_filename_ field.
			"header_option_filename": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_filename_ field.",
			},

			// The inheritance configuration for _header_option_server_address_ field.
			"header_option_server_address": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_address_ field.",
			},

			// The inheritance configuration for _header_option_server_name_ field.
			"header_option_server_name": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_name_ field.",
			},

			// The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields from _IPSpace_ object.
			"hostname_rewrite_block": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedHostnameRewriteBlock(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields from _IPSpace_ object.",
			},

			// The inheritance configuration for _vendor_specific_option_option_space_ field.
			"vendor_specific_option_option_space": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedIdentifier(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _vendor_specific_option_option_space_ field.",
			},
		},
	}
}

func flattenIpamsvcIPSpaceInheritance(r *models.IpamsvcIPSpaceInheritance) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"asm_config":                          flattenIpamsvcInheritedASMConfig(r.AsmConfig),
			"ddns_client_update":                  flattenInheritanceInheritedString(r.DdnsClientUpdate),
			"ddns_enabled":                        flattenInheritanceInheritedBool(r.DdnsEnabled),
			"ddns_hostname_block":                 flattenIpamsvcInheritedDDNSHostnameBlock(r.DdnsHostnameBlock),
			"ddns_update_block":                   flattenIpamsvcInheritedDDNSUpdateBlock(r.DdnsUpdateBlock),
			"ddns_update_on_renew":                flattenInheritanceInheritedBool(r.DdnsUpdateOnRenew),
			"ddns_use_conflict_resolution":        flattenInheritanceInheritedBool(r.DdnsUseConflictResolution),
			"dhcp_config":                         flattenIpamsvcInheritedDHCPConfig(r.DhcpConfig),
			"dhcp_options":                        flattenIpamsvcInheritedDHCPOptionList(r.DhcpOptions),
			"header_option_filename":              flattenInheritanceInheritedString(r.HeaderOptionFilename),
			"header_option_server_address":        flattenInheritanceInheritedString(r.HeaderOptionServerAddress),
			"header_option_server_name":           flattenInheritanceInheritedString(r.HeaderOptionServerName),
			"hostname_rewrite_block":              flattenIpamsvcInheritedHostnameRewriteBlock(r.HostnameRewriteBlock),
			"vendor_specific_option_option_space": flattenInheritanceInheritedIdentifier(r.VendorSpecificOptionOptionSpace),
		},
	}
}

func expandIpamsvcIPSpaceInheritance(ctx context.Context, d []interface{}) (*models.IpamsvcIPSpaceInheritance, error) {
	if len(d) == 0 || d[0] == nil {
		return nil, nil
	}
	in := d[0].(map[string]interface{})

	asmConfig, err := expandIpamsvcInheritedASMConfig(ctx, in["asm_config"].([]interface{}))
	if err != nil {
		tflog.Error(ctx, "Failed to parse 'asm_config' field. The underlying expand function returned an error.")
		return nil, err
	}

	return &models.IpamsvcIPSpaceInheritance{
		AsmConfig:                       asmConfig,
		DdnsClientUpdate:                expandInheritanceInheritedString(in["ddns_client_update"].([]interface{})),
		DdnsEnabled:                     expandInheritanceInheritedBool(in["ddns_enabled"].([]interface{})),
		DdnsHostnameBlock:               expandIpamsvcInheritedDDNSHostnameBlock(in["ddns_hostname_block"].([]interface{})),
		DdnsUpdateBlock:                 expandIpamsvcInheritedDDNSUpdateBlock(in["ddns_update_block"].([]interface{})),
		DdnsUpdateOnRenew:               expandInheritanceInheritedBool(in["ddns_update_on_renew"].([]interface{})),
		DdnsUseConflictResolution:       expandInheritanceInheritedBool(in["ddns_use_conflict_resolution"].([]interface{})),
		DhcpConfig:                      expandIpamsvcInheritedDHCPConfig(in["dhcp_config"].([]interface{})),
		DhcpOptions:                     expandIpamsvcInheritedDHCPOptionList(in["dhcp_options"].([]interface{})),
		HeaderOptionFilename:            expandInheritanceInheritedString(in["header_option_filename"].([]interface{})),
		HeaderOptionServerAddress:       expandInheritanceInheritedString(in["header_option_server_address"].([]interface{})),
		HeaderOptionServerName:          expandInheritanceInheritedString(in["header_option_server_name"].([]interface{})),
		HostnameRewriteBlock:            expandIpamsvcInheritedHostnameRewriteBlock(in["hostname_rewrite_block"].([]interface{})),
		VendorSpecificOptionOptionSpace: expandInheritanceInheritedIdentifier(in["vendor_specific_option_option_space"].([]interface{})),
	}, nil
}
