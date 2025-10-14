package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcServerModel struct {
	ClientPrincipal                 types.String      `tfsdk:"client_principal"`
	Comment                         types.String      `tfsdk:"comment"`
	CreatedAt                       timetypes.RFC3339 `tfsdk:"created_at"`
	DdnsClientUpdate                types.String      `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.String      `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                      types.String      `tfsdk:"ddns_domain"`
	DdnsEnabled                     types.Bool        `tfsdk:"ddns_enabled"`
	DdnsGenerateName                types.Bool        `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix             types.String      `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates                 types.Bool        `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent                  types.Float32     `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew               types.Bool        `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Bool        `tfsdk:"ddns_use_conflict_resolution"`
	DdnsZones                       types.List        `tfsdk:"ddns_zones"`
	DhcpConfig                      types.Object      `tfsdk:"dhcp_config"`
	DhcpOptions                     types.List        `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.List        `tfsdk:"dhcp_options_v6"`
	GssTsigFallback                 types.Bool        `tfsdk:"gss_tsig_fallback"`
	HeaderOptionFilename            types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String      `tfsdk:"header_option_server_name"`
	HostnameRewriteChar             types.String      `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled          types.Bool        `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex            types.String      `tfsdk:"hostname_rewrite_regex"`
	Id                              types.String      `tfsdk:"id"`
	InheritanceSources              types.Object      `tfsdk:"inheritance_sources"`
	KerberosKdc                     types.String      `tfsdk:"kerberos_kdc"`
	KerberosKeys                    types.List        `tfsdk:"kerberos_keys"`
	KerberosRekeyInterval           types.Int64       `tfsdk:"kerberos_rekey_interval"`
	KerberosRetryInterval           types.Int64       `tfsdk:"kerberos_retry_interval"`
	KerberosTkeyLifetime            types.Int64       `tfsdk:"kerberos_tkey_lifetime"`
	KerberosTkeyProtocol            types.String      `tfsdk:"kerberos_tkey_protocol"`
	Name                            types.String      `tfsdk:"name"`
	ProfileType                     types.String      `tfsdk:"profile_type"`
	ServerPrincipal                 types.String      `tfsdk:"server_principal"`
	Tags                            types.Map         `tfsdk:"tags"`
	TagsAll                         types.Map         `tfsdk:"tags_all"`
	UpdatedAt                       timetypes.RFC3339 `tfsdk:"updated_at"`
	VendorSpecificOptionOptionSpace types.String      `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcServerAttrTypes = map[string]attr.Type{
	"client_principal":                    types.StringType,
	"comment":                             types.StringType,
	"created_at":                          timetypes.RFC3339Type{},
	"ddns_client_update":                  types.StringType,
	"ddns_conflict_resolution_mode":       types.StringType,
	"ddns_domain":                         types.StringType,
	"ddns_enabled":                        types.BoolType,
	"ddns_generate_name":                  types.BoolType,
	"ddns_generated_prefix":               types.StringType,
	"ddns_send_updates":                   types.BoolType,
	"ddns_ttl_percent":                    types.Float32Type,
	"ddns_update_on_renew":                types.BoolType,
	"ddns_use_conflict_resolution":        types.BoolType,
	"ddns_zones":                          types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcDDNSZoneAttrTypes}},
	"dhcp_config":                         types.ObjectType{AttrTypes: IpamsvcDHCPConfigAttrTypes},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_options_v6":                     types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"gss_tsig_fallback":                   types.BoolType,
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"hostname_rewrite_char":               types.StringType,
	"hostname_rewrite_enabled":            types.BoolType,
	"hostname_rewrite_regex":              types.StringType,
	"id":                                  types.StringType,
	"inheritance_sources":                 types.ObjectType{AttrTypes: IpamsvcServerInheritanceAttrTypes},
	"kerberos_kdc":                        types.StringType,
	"kerberos_keys":                       types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcKerberosKeyAttrTypes}},
	"kerberos_rekey_interval":             types.Int64Type,
	"kerberos_retry_interval":             types.Int64Type,
	"kerberos_tkey_lifetime":              types.Int64Type,
	"kerberos_tkey_protocol":              types.StringType,
	"name":                                types.StringType,
	"profile_type":                        types.StringType,
	"server_principal":                    types.StringType,
	"tags":                                types.MapType{ElemType: types.StringType},
	"tags_all":                            types.MapType{ElemType: types.StringType},
	"updated_at":                          timetypes.RFC3339Type{},
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcServerResourceSchemaAttributes = map[string]schema.Attribute{
	"client_principal": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The Kerberos principal name. It uses the typical Kerberos notation: `<SERVICE-NAME>/<server-domain-name>@<REALM>`. Defaults to empty.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the DHCP Config Profile. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"ddns_client_update": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("client"),
		MarkdownDescription: "Controls who does the DDNS updates. Valid values are:\n" +
			"  * _client_: DHCP server updates DNS if requested by client.\n" +
			"  * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n" +
			"  * _ignore_: DHCP server always updates DNS, even if the client says not to.\n" +
			"  * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n" +
			"  * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n\n" +
			"  Defaults to _client_.",
		Validators: []validator.String{
			stringvalidator.OneOf("client", "server", "ignore", "over_client_update", "over_no_update"),
		},
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("check_with_dhcid"),
		MarkdownDescription: "The mode used for resolving conflicts while performing DDNS updates. Valid values are:\n" +
			"  * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703.\n" +
			"  * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry.\n" +
			"  * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry.\n" +
			"  * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.\n\n" +
			"  Defaults to _check_with_dhcid_.",
		Validators: []validator.String{
			stringvalidator.OneOf("check_with_dhcid", "no_check_with_dhcid", "check_exists_with_dhcid", "no_check_without_dhcid"),
		},
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The domain suffix for DDNS updates. FQDN, may be empty. Required if _ddns_enabled_ is true.  Defaults to empty.",
	},
	"ddns_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if DDNS updates should be performed for leases. All other _ddns_*_ configuration is ignored when this flag is unset. At a minimum, _ddns_domain_ and _ddns_zones_ must be configured to enable DDNS. Defaults to _false_.",
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if DDNS should generate a hostname when not supplied by the client.  Defaults to _false_.",
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("myhost"),
		MarkdownDescription: `The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string. Defaults to \"myhost\".`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Determines if DDNS updates are enabled at the IP space level. Defaults to _true_.`,
	},
	"ddns_ttl_percent": schema.Float32Attribute{
		Optional:            true,
		MarkdownDescription: "DDNS TTL value - to be calculated as a simple percentage of the lease's lifetime, using the parameter's value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.",
	},
	"ddns_update_on_renew": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_.`,
	},
	"ddns_use_conflict_resolution": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_. Can be set to true only when ddns_conflict_resolution_mode is check_with_dhcid.`,
	},
	"ddns_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcDDNSZoneResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: IpamsvcDDNSZoneAttrTypes})),
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The DNS zones that DDNS updates can be sent to. There is no resolver fallback. The target zone must be explicitly configured for the update to be performed.  Updates are sent to the closest enclosing zone.  Error if _ddns_enabled_ is _true_ and the _ddns_domain_ does not have a corresponding entry in _ddns_zones_.  Error if there are items with duplicate zone in the list.  Defaults to empty list.",
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPConfigResourceSchemaAttributes(false),
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(IpamsvcDHCPConfigAttrTypes, map[string]attr.Value{
			"abandoned_reclaim_time":    types.Int64Value(3600),
			"abandoned_reclaim_time_v6": types.Int64Value(3600),
			"allow_unknown":             types.BoolValue(true),
			"allow_unknown_v6":          types.BoolValue(true),
			"echo_client_id":            types.BoolValue(true),
			"filters":                   types.ListNull(types.StringType),
			"filters_large_selection":   types.ListNull(types.StringType),
			"filters_v6":                types.ListNull(types.StringType),
			"ignore_client_uid":         types.BoolValue(true),
			"ignore_list":               types.ListNull(types.ObjectType{AttrTypes: IpamsvcIgnoreItemAttrTypes}),
			"lease_time":                types.Int64Value(3600),
			"lease_time_v6":             types.Int64Value(3600),
		})),
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP options or group of options for IPv4. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used. Error if the graph of referenced groups contains cycles. Defaults to empty list.",
	},
	"dhcp_options_v6": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP options or group of options for IPv6. An option list is ordered and may include both option groups and specific options. Multiple occurrences of the same option or group is not an error. The last occurrence of an option in the list will be used. Error if the graph of referenced groups contains cycles. Defaults to empty list.",
	},
	"gss_tsig_fallback": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.",
	},
	"header_option_filename": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The configuration for header option filename field.",
	},
	"header_option_server_address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The configuration for header option server address field.",
	},
	"header_option_server_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The configuration for header option server name field.",
	},
	"hostname_rewrite_char": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("-"),
		MarkdownDescription: `The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \"-\".`,
		Validators: []validator.String{
			stringvalidator.LengthAtMost(1),
		},
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.`,
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("[^a-zA-Z0-9_.]"),
		MarkdownDescription: `The regex bracket expression to match valid characters.  Must begin with \"[\" and end with \"]\" and be a compilable POSIX regex.  Defaults to \"[^a-zA-Z0-9_.]\".`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: IpamsvcServerInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration.",
	},
	"kerberos_kdc": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Address of Kerberos Key Distribution Center.  Defaults to empty.",
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcKerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"kerberos_rekey_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(120),
		MarkdownDescription: "Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the _kerberos_rekey_interval_ value.  Defaults to 120 seconds.",
	},
	"kerberos_retry_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(30),
		MarkdownDescription: "Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds.",
	},
	"kerberos_tkey_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(160),
		MarkdownDescription: "Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds.",
	},
	"kerberos_tkey_protocol": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the DHCP Config Profile. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"profile_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("server"),
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The type of server object.  Defaults to _server_.  Valid values are: * _server_: The server profile type. * _subnet_: The subnet profile type.",
	},
	"server_principal": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the DHCP Config Profile in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the DHCP Config Profile in JSON format including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func (m *IpamsvcServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Server {
	if m == nil {
		return nil
	}
	to := &ipam.Server{
		ClientPrincipal:                 flex.ExpandStringPointer(m.ClientPrincipal),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		DdnsClientUpdate:                flex.ExpandStringPointer(m.DdnsClientUpdate),
		DdnsConflictResolutionMode:      flex.ExpandStringPointer(m.DdnsConflictResolutionMode),
		DdnsDomain:                      flex.ExpandStringPointer(m.DdnsDomain),
		DdnsEnabled:                     flex.ExpandBoolPointer(m.DdnsEnabled),
		DdnsGenerateName:                flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix:             flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
		DdnsSendUpdates:                 flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsTtlPercent:                  flex.ExpandFloat32Pointer(m.DdnsTtlPercent),
		DdnsUpdateOnRenew:               flex.ExpandBoolPointer(m.DdnsUpdateOnRenew),
		DdnsUseConflictResolution:       flex.ExpandBoolPointer(m.DdnsUseConflictResolution),
		DdnsZones:                       flex.ExpandFrameworkListNestedBlockNilAsEmpty(ctx, m.DdnsZones, diags, ExpandIpamsvcDDNSZone),
		DhcpConfig:                      ExpandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		DhcpOptionsV6:                   flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptionsV6, diags, ExpandIpamsvcOptionItem),
		GssTsigFallback:                 flex.ExpandBoolPointer(m.GssTsigFallback),
		HeaderOptionFilename:            flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:       flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:          flex.ExpandStringPointer(m.HeaderOptionServerName),
		HostnameRewriteChar:             flex.ExpandStringPointer(m.HostnameRewriteChar),
		HostnameRewriteEnabled:          flex.ExpandBoolPointer(m.HostnameRewriteEnabled),
		HostnameRewriteRegex:            flex.ExpandStringPointer(m.HostnameRewriteRegex),
		InheritanceSources:              ExpandIpamsvcServerInheritance(ctx, m.InheritanceSources, diags),
		KerberosKdc:                     flex.ExpandStringPointer(m.KerberosKdc),
		KerberosKeys:                    flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandIpamsvcKerberosKey),
		KerberosRekeyInterval:           flex.ExpandInt64Pointer(m.KerberosRekeyInterval),
		KerberosRetryInterval:           flex.ExpandInt64Pointer(m.KerberosRetryInterval),
		KerberosTkeyLifetime:            flex.ExpandInt64Pointer(m.KerberosTkeyLifetime),
		KerberosTkeyProtocol:            flex.ExpandStringPointer(m.KerberosTkeyProtocol),
		Name:                            flex.ExpandString(m.Name),
		ProfileType:                     flex.ExpandStringPointer(m.ProfileType),
		ServerPrincipal:                 flex.ExpandStringPointer(m.ServerPrincipal),
		Tags:                            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: flex.ExpandStringPointer(m.VendorSpecificOptionOptionSpace),
	}
	return to
}

func FlattenIpamsvcServerDataSource(ctx context.Context, from *ipam.Server, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcServerAttrTypes)
	}
	m := IpamsvcServerModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, IpamsvcServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcServerModel) Flatten(ctx context.Context, from *ipam.Server, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcServerModel{}
	}
	m.ClientPrincipal = flex.FlattenStringPointer(from.ClientPrincipal)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DdnsClientUpdate = flex.FlattenStringPointer(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = flex.FlattenStringPointer(from.DdnsConflictResolutionMode)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsEnabled = types.BoolPointerValue(from.DdnsEnabled)
	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsTtlPercent = flex.FlattenFloat32(*from.DdnsTtlPercent)
	m.DdnsUpdateOnRenew = types.BoolPointerValue(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = types.BoolPointerValue(from.DdnsUseConflictResolution)
	m.DdnsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.DdnsZones, IpamsvcDDNSZoneAttrTypes, diags, FlattenIpamsvcDDNSZone)
	m.DhcpConfig = FlattenIpamsvcDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.DhcpOptionsV6 = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptionsV6, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.GssTsigFallback = types.BoolPointerValue(from.GssTsigFallback)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenIpamsvcServerInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKdc = flex.FlattenStringPointer(from.KerberosKdc)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, IpamsvcKerberosKeyAttrTypes, diags, FlattenIpamsvcKerberosKey)
	m.KerberosRekeyInterval = flex.FlattenInt64Pointer(from.KerberosRekeyInterval)
	m.KerberosRetryInterval = flex.FlattenInt64Pointer(from.KerberosRetryInterval)
	m.KerberosTkeyLifetime = flex.FlattenInt64Pointer(from.KerberosTkeyLifetime)
	m.KerberosTkeyProtocol = flex.FlattenStringPointer(from.KerberosTkeyProtocol)
	m.Name = flex.FlattenString(from.Name)
	m.ProfileType = flex.FlattenStringPointer(from.ProfileType)
	m.ServerPrincipal = flex.FlattenStringPointer(from.ServerPrincipal)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.VendorSpecificOptionOptionSpace = flex.FlattenStringPointer(from.VendorSpecificOptionOptionSpace)
}
