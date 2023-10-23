package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcGlobalModel struct {
	AsmConfig                       types.Object  `tfsdk:"asm_config"`
	ClientPrincipal                 types.String  `tfsdk:"client_principal"`
	DdnsClientUpdate                types.String  `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.String  `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                      types.String  `tfsdk:"ddns_domain"`
	DdnsEnabled                     types.Bool    `tfsdk:"ddns_enabled"`
	DdnsGenerateName                types.Bool    `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix             types.String  `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates                 types.Bool    `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent                  types.Float64 `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew               types.Bool    `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Bool    `tfsdk:"ddns_use_conflict_resolution"`
	DdnsZones                       types.List    `tfsdk:"ddns_zones"`
	DhcpConfig                      types.Object  `tfsdk:"dhcp_config"`
	DhcpOptions                     types.List    `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.List    `tfsdk:"dhcp_options_v6"`
	DhcpThreshold                   types.Object  `tfsdk:"dhcp_threshold"`
	GssTsigFallback                 types.Bool    `tfsdk:"gss_tsig_fallback"`
	HeaderOptionFilename            types.String  `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String  `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String  `tfsdk:"header_option_server_name"`
	HostnameRewriteChar             types.String  `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled          types.Bool    `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex            types.String  `tfsdk:"hostname_rewrite_regex"`
	Id                              types.String  `tfsdk:"id"`
	KerberosKdc                     types.String  `tfsdk:"kerberos_kdc"`
	KerberosKeys                    types.List    `tfsdk:"kerberos_keys"`
	KerberosRekeyInterval           types.Int64   `tfsdk:"kerberos_rekey_interval"`
	KerberosRetryInterval           types.Int64   `tfsdk:"kerberos_retry_interval"`
	KerberosTkeyLifetime            types.Int64   `tfsdk:"kerberos_tkey_lifetime"`
	KerberosTkeyProtocol            types.String  `tfsdk:"kerberos_tkey_protocol"`
	PreferOption12                  types.Bool    `tfsdk:"prefer_option_12"`
	RemoveSuffixOption81            types.Bool    `tfsdk:"remove_suffix_option_81"`
	ServerPrincipal                 types.String  `tfsdk:"server_principal"`
	VendorSpecificOptionOptionSpace types.String  `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcGlobalAttrTypes = map[string]attr.Type{
	"asm_config":                          types.ObjectType{AttrTypes: IpamsvcASMConfigAttrTypes},
	"client_principal":                    types.StringType,
	"ddns_client_update":                  types.StringType,
	"ddns_conflict_resolution_mode":       types.StringType,
	"ddns_domain":                         types.StringType,
	"ddns_enabled":                        types.BoolType,
	"ddns_generate_name":                  types.BoolType,
	"ddns_generated_prefix":               types.StringType,
	"ddns_send_updates":                   types.BoolType,
	"ddns_ttl_percent":                    types.Float64Type,
	"ddns_update_on_renew":                types.BoolType,
	"ddns_use_conflict_resolution":        types.BoolType,
	"ddns_zones":                          types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcDDNSZoneAttrTypes}},
	"dhcp_config":                         types.ObjectType{AttrTypes: IpamsvcDHCPConfigAttrTypes},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_options_v6":                     types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_threshold":                      types.ObjectType{AttrTypes: IpamsvcDHCPUtilizationThresholdAttrTypes},
	"gss_tsig_fallback":                   types.BoolType,
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"hostname_rewrite_char":               types.StringType,
	"hostname_rewrite_enabled":            types.BoolType,
	"hostname_rewrite_regex":              types.StringType,
	"id":                                  types.StringType,
	"kerberos_kdc":                        types.StringType,
	"kerberos_keys":                       types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcKerberosKeyAttrTypes}},
	"kerberos_rekey_interval":             types.Int64Type,
	"kerberos_retry_interval":             types.Int64Type,
	"kerberos_tkey_lifetime":              types.Int64Type,
	"kerberos_tkey_protocol":              types.StringType,
	"prefer_option_12":                    types.BoolType,
	"remove_suffix_option_81":             types.BoolType,
	"server_principal":                    types.StringType,
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcGlobalResourceSchema = schema.Schema{
	MarkdownDescription: `The global DHCP configuration (_dhcp/global_). Used by default unless more specific configuration exists. There is only one instance of this object.`,
	Attributes:          IpamsvcGlobalResourceSchemaAttributes,
}

var IpamsvcGlobalResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcASMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"client_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name. It uses the typical Kerberos notation: &lt;SERVICE-NAME&gt;/&lt;server-domain-name&gt;@&lt;REALM&gt;.  Defaults to empty.`,
	},
	"ddns_client_update": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The global configuration to control who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_.`,
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_.`,
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain suffix for DDNS updates. FQDN, may be empty.  Must be specified if _ddns_enabled_ is _true_.  Defaults to empty.`,
	},
	"ddns_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if DDNS updates should be performed for leases.  All other ddns_* configuration fields are ignored when this flag is unset.  At a minimum, _ddns_domain_ and _ddns_zones_ must be configured to enable DDNS.  Defaults to _false_.`,
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_.`,
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \&quot;myhost\&quot;.`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Determines if DDNS updates are enabled at the global level. Defaults to _true_.`,
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		MarkdownDescription: `DDNS TTL value - to be calculated as a simple percentage of the lease&#39;s lifetime, using the parameter&#39;s value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.`,
	},
	"ddns_update_on_renew": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_.`,
	},
	"ddns_use_conflict_resolution": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_.`,
	},
	"ddns_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcDDNSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `DNS zones that DDNS updates can be sent to. There is no resolver fallback. The target zone must be explicitly configured for the update to be performed.  Updates are sent to the closest enclosing zone.  Error if _ddns_enabled_ is _true_ and the _ddns_domain_ does not have a corresponding entry in _ddns_zones_.  Error if there are items with duplicate zone in the list.  Defaults to empty list.`,
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DHCP options or group of options for IPv4. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used.  Error if the graph of referenced groups contains cycles.  Defaults to empty list.`,
	},
	"dhcp_options_v6": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DHCP options or group of options for IPv6. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used.  Error if the graph of referenced groups contains cycles.  Defaults to empty list.`,
	},
	"dhcp_threshold": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPUtilizationThresholdResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"gss_tsig_fallback": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.`,
	},
	"header_option_filename": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option filename field.`,
	},
	"header_option_server_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option server address field.`,
	},
	"header_option_server_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option server name field.`,
	},
	"hostname_rewrite_char": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The character to replace non-matching characters with, when hostname rewrite is enabled in global configuration.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;.`,
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The global configuration to indicate if the hostnames supplied by the client will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.`,
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The regex bracket expression to match valid characters when hostname rewrite is enabled in global configuration.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"kerberos_kdc": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Address of Kerberos Key Distribution Center.  Defaults to empty.`,
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcKerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.`,
	},
	"kerberos_rekey_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the _kerberos_rekey_interval_ value.  Defaults to 120 seconds.`,
	},
	"kerberos_retry_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds.`,
	},
	"kerberos_tkey_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds.`,
	},
	"kerberos_tkey_protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_.`,
	},
	"prefer_option_12": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `When enabled, DHCP Server will prefer option 12 over option 81 in the incoming client request.  Defaults to _false_.`,
	},
	"remove_suffix_option_81": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `When enabled, DHCP Server will remove the suffix from the option 81 in the incoming client request.  Defaults to _false_.`,
	},
	"server_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty.`,
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcGlobal(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcGlobal {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcGlobalModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcGlobalModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcGlobal {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcGlobal{
		AsmConfig:                       expandIpamsvcASMConfig(ctx, m.AsmConfig, diags),
		ClientPrincipal:                 m.ClientPrincipal.ValueStringPointer(),
		DdnsClientUpdate:                m.DdnsClientUpdate.ValueStringPointer(),
		DdnsConflictResolutionMode:      m.DdnsConflictResolutionMode.ValueStringPointer(),
		DdnsDomain:                      m.DdnsDomain.ValueStringPointer(),
		DdnsEnabled:                     m.DdnsEnabled.ValueBoolPointer(),
		DdnsGenerateName:                m.DdnsGenerateName.ValueBoolPointer(),
		DdnsGeneratedPrefix:             m.DdnsGeneratedPrefix.ValueStringPointer(),
		DdnsSendUpdates:                 m.DdnsSendUpdates.ValueBoolPointer(),
		DdnsTtlPercent:                  ptr(float32(m.DdnsTtlPercent.ValueFloat64())),
		DdnsUpdateOnRenew:               m.DdnsUpdateOnRenew.ValueBoolPointer(),
		DdnsUseConflictResolution:       m.DdnsUseConflictResolution.ValueBoolPointer(),
		DdnsZones:                       ExpandFrameworkListNestedBlock(ctx, m.DdnsZones, diags, expandIpamsvcDDNSZone),
		DhcpConfig:                      expandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		DhcpOptionsV6:                   ExpandFrameworkListNestedBlock(ctx, m.DhcpOptionsV6, diags, expandIpamsvcOptionItem),
		DhcpThreshold:                   expandIpamsvcDHCPUtilizationThreshold(ctx, m.DhcpThreshold, diags),
		GssTsigFallback:                 m.GssTsigFallback.ValueBoolPointer(),
		HeaderOptionFilename:            m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:       m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:          m.HeaderOptionServerName.ValueStringPointer(),
		HostnameRewriteChar:             m.HostnameRewriteChar.ValueStringPointer(),
		HostnameRewriteEnabled:          m.HostnameRewriteEnabled.ValueBoolPointer(),
		HostnameRewriteRegex:            m.HostnameRewriteRegex.ValueStringPointer(),
		KerberosKdc:                     m.KerberosKdc.ValueStringPointer(),
		KerberosKeys:                    ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, expandIpamsvcKerberosKey),
		KerberosRekeyInterval:           ptr(int64(m.KerberosRekeyInterval.ValueInt64())),
		KerberosRetryInterval:           ptr(int64(m.KerberosRetryInterval.ValueInt64())),
		KerberosTkeyLifetime:            ptr(int64(m.KerberosTkeyLifetime.ValueInt64())),
		KerberosTkeyProtocol:            m.KerberosTkeyProtocol.ValueStringPointer(),
		PreferOption12:                  m.PreferOption12.ValueBoolPointer(),
		RemoveSuffixOption81:            m.RemoveSuffixOption81.ValueBoolPointer(),
		ServerPrincipal:                 m.ServerPrincipal.ValueStringPointer(),
		VendorSpecificOptionOptionSpace: m.VendorSpecificOptionOptionSpace.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcGlobal(ctx context.Context, from *ipam.IpamsvcGlobal, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcGlobalAttrTypes)
	}
	m := IpamsvcGlobalModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcGlobalAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcGlobalModel) flatten(ctx context.Context, from *ipam.IpamsvcGlobal, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcGlobalModel{}
	}

	m.AsmConfig = flattenIpamsvcASMConfig(ctx, from.AsmConfig, diags)
	m.ClientPrincipal = types.StringPointerValue(from.ClientPrincipal)
	m.DdnsClientUpdate = types.StringPointerValue(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = types.StringPointerValue(from.DdnsConflictResolutionMode)
	m.DdnsDomain = types.StringPointerValue(from.DdnsDomain)
	m.DdnsEnabled = types.BoolPointerValue(from.DdnsEnabled)
	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = types.StringPointerValue(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsTtlPercent = types.Float64Value(float64(*from.DdnsTtlPercent))
	m.DdnsUpdateOnRenew = types.BoolPointerValue(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = types.BoolPointerValue(from.DdnsUseConflictResolution)
	m.DdnsZones = FlattenFrameworkListNestedBlock(ctx, from.DdnsZones, IpamsvcDDNSZoneAttrTypes, diags, flattenIpamsvcDDNSZone)
	m.DhcpConfig = flattenIpamsvcDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.DhcpOptionsV6 = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptionsV6, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.DhcpThreshold = flattenIpamsvcDHCPUtilizationThreshold(ctx, from.DhcpThreshold, diags)
	m.GssTsigFallback = types.BoolPointerValue(from.GssTsigFallback)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.HostnameRewriteChar = types.StringPointerValue(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = types.StringPointerValue(from.HostnameRewriteRegex)
	m.Id = types.StringPointerValue(from.Id)
	m.KerberosKdc = types.StringPointerValue(from.KerberosKdc)
	m.KerberosKeys = FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, IpamsvcKerberosKeyAttrTypes, diags, flattenIpamsvcKerberosKey)
	m.KerberosRekeyInterval = types.Int64Value(int64(*from.KerberosRekeyInterval))
	m.KerberosRetryInterval = types.Int64Value(int64(*from.KerberosRetryInterval))
	m.KerberosTkeyLifetime = types.Int64Value(int64(*from.KerberosTkeyLifetime))
	m.KerberosTkeyProtocol = types.StringPointerValue(from.KerberosTkeyProtocol)
	m.PreferOption12 = types.BoolPointerValue(from.PreferOption12)
	m.RemoveSuffixOption81 = types.BoolPointerValue(from.RemoveSuffixOption81)
	m.ServerPrincipal = types.StringPointerValue(from.ServerPrincipal)
	m.VendorSpecificOptionOptionSpace = types.StringPointerValue(from.VendorSpecificOptionOptionSpace)

}
