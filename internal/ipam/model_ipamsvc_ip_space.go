package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcIPSpaceModel struct {
	AsmConfig                       types.Object      `tfsdk:"asm_config"`
	AsmScopeFlag                    types.Int64       `tfsdk:"asm_scope_flag"`
	Comment                         types.String      `tfsdk:"comment"`
	CreatedAt                       timetypes.RFC3339 `tfsdk:"created_at"`
	DdnsClientUpdate                types.String      `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.String      `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                      types.String      `tfsdk:"ddns_domain"`
	DdnsGenerateName                types.Bool        `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix             types.String      `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates                 types.Bool        `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent                  types.Float64     `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew               types.Bool        `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Bool        `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                      types.Object      `tfsdk:"dhcp_config"`
	DhcpOptions                     types.List        `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.List        `tfsdk:"dhcp_options_v6"`
	HeaderOptionFilename            types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String      `tfsdk:"header_option_server_name"`
	HostnameRewriteChar             types.String      `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled          types.Bool        `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex            types.String      `tfsdk:"hostname_rewrite_regex"`
	Id                              types.String      `tfsdk:"id"`
	InheritanceSources              types.Object      `tfsdk:"inheritance_sources"`
	Name                            types.String      `tfsdk:"name"`
	Tags                            types.Map         `tfsdk:"tags"`
	Threshold                       types.Object      `tfsdk:"threshold"`
	UpdatedAt                       timetypes.RFC3339 `tfsdk:"updated_at"`
	Utilization                     types.Object      `tfsdk:"utilization"`
	UtilizationV6                   types.Object      `tfsdk:"utilization_v6"`
	VendorSpecificOptionOptionSpace types.String      `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcIPSpaceAttrTypes = map[string]attr.Type{
	"asm_config":                          types.ObjectType{AttrTypes: IpamsvcASMConfigAttrTypes},
	"asm_scope_flag":                      types.Int64Type,
	"comment":                             types.StringType,
	"created_at":                          timetypes.RFC3339Type{},
	"ddns_client_update":                  types.StringType,
	"ddns_conflict_resolution_mode":       types.StringType,
	"ddns_domain":                         types.StringType,
	"ddns_generate_name":                  types.BoolType,
	"ddns_generated_prefix":               types.StringType,
	"ddns_send_updates":                   types.BoolType,
	"ddns_ttl_percent":                    types.Float64Type,
	"ddns_update_on_renew":                types.BoolType,
	"ddns_use_conflict_resolution":        types.BoolType,
	"dhcp_config":                         types.ObjectType{AttrTypes: IpamsvcDHCPConfigAttrTypes},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_options_v6":                     types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"hostname_rewrite_char":               types.StringType,
	"hostname_rewrite_enabled":            types.BoolType,
	"hostname_rewrite_regex":              types.StringType,
	"id":                                  types.StringType,
	"inheritance_sources":                 types.ObjectType{AttrTypes: IpamsvcIPSpaceInheritanceAttrTypes},
	"name":                                types.StringType,
	"tags":                                types.MapType{},
	"threshold":                           types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                          timetypes.RFC3339Type{},
	"utilization":                         types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":                      types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcIPSpaceResourceSchema = schema.Schema{
	MarkdownDescription: `An __IPSpace__ object (_ipam/ip_space_) allows customers to represent their entire managed address space with no collision. A collision arises when two or more block of addresses overlap partially or fully.`,
	Attributes:          IpamsvcIPSpaceResourceSchemaAttributes,
}

var IpamsvcIPSpaceResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcASMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"asm_scope_flag": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The number of times the automated scope management usage limits have been exceeded for any of the subnets in this IP space.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the IP space. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"ddns_client_update": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_.`,
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_.`,
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty.`,
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
		MarkdownDescription: `Determines if DDNS updates are enabled at the IP space level. Defaults to _true_.`,
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
		MarkdownDescription: `The list of IPv4 DHCP options for IP space. May be either a specific option or a group of options.`,
	},
	"dhcp_options_v6": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of IPv6 DHCP options for IP space. May be either a specific option or a group of options.`,
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
		MarkdownDescription: `The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \&quot;-\&quot;.`,
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.`,
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The regex bracket expression to match valid characters.  Must begin with \&quot;[\&quot; and end with \&quot;]\&quot; and be a compilable POSIX regex.  Defaults to \&quot;[^a-zA-Z0-9_.]\&quot;.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          IpamsvcIPSpaceInheritanceResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the IP space. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the IP space in JSON format.`,
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationThresholdResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"utilization": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationV6ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcIPSpace(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcIPSpace {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcIPSpaceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcIPSpaceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcIPSpace {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcIPSpace{
		AsmConfig:                       expandIpamsvcASMConfig(ctx, m.AsmConfig, diags),
		Comment:                         m.Comment.ValueStringPointer(),
		DdnsClientUpdate:                m.DdnsClientUpdate.ValueStringPointer(),
		DdnsConflictResolutionMode:      m.DdnsConflictResolutionMode.ValueStringPointer(),
		DdnsDomain:                      m.DdnsDomain.ValueStringPointer(),
		DdnsGenerateName:                m.DdnsGenerateName.ValueBoolPointer(),
		DdnsGeneratedPrefix:             m.DdnsGeneratedPrefix.ValueStringPointer(),
		DdnsSendUpdates:                 m.DdnsSendUpdates.ValueBoolPointer(),
		DdnsTtlPercent:                  ptr(float32(m.DdnsTtlPercent.ValueFloat64())),
		DdnsUpdateOnRenew:               m.DdnsUpdateOnRenew.ValueBoolPointer(),
		DdnsUseConflictResolution:       m.DdnsUseConflictResolution.ValueBoolPointer(),
		DhcpConfig:                      expandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		DhcpOptionsV6:                   ExpandFrameworkListNestedBlock(ctx, m.DhcpOptionsV6, diags, expandIpamsvcOptionItem),
		HeaderOptionFilename:            m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:       m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:          m.HeaderOptionServerName.ValueStringPointer(),
		HostnameRewriteChar:             m.HostnameRewriteChar.ValueStringPointer(),
		HostnameRewriteEnabled:          m.HostnameRewriteEnabled.ValueBoolPointer(),
		HostnameRewriteRegex:            m.HostnameRewriteRegex.ValueStringPointer(),
		InheritanceSources:              expandIpamsvcIPSpaceInheritance(ctx, m.InheritanceSources, diags),
		Name:                            m.Name.ValueString(),
		Tags:                            ExpandFrameworkMapString(ctx, m.Tags, diags),
		Threshold:                       expandIpamsvcUtilizationThreshold(ctx, m.Threshold, diags),
		Utilization:                     expandIpamsvcUtilization(ctx, m.Utilization, diags),
		UtilizationV6:                   expandIpamsvcUtilizationV6(ctx, m.UtilizationV6, diags),
		VendorSpecificOptionOptionSpace: m.VendorSpecificOptionOptionSpace.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcIPSpace(ctx context.Context, from *ipam.IpamsvcIPSpace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIPSpaceAttrTypes)
	}
	m := IpamsvcIPSpaceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIPSpaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIPSpaceModel) flatten(ctx context.Context, from *ipam.IpamsvcIPSpace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIPSpaceModel{}
	}

	m.AsmConfig = flattenIpamsvcASMConfig(ctx, from.AsmConfig, diags)
	m.AsmScopeFlag = types.Int64Value(int64(*from.AsmScopeFlag))
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DdnsClientUpdate = types.StringPointerValue(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = types.StringPointerValue(from.DdnsConflictResolutionMode)
	m.DdnsDomain = types.StringPointerValue(from.DdnsDomain)
	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = types.StringPointerValue(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsTtlPercent = types.Float64Value(float64(*from.DdnsTtlPercent))
	m.DdnsUpdateOnRenew = types.BoolPointerValue(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = types.BoolPointerValue(from.DdnsUseConflictResolution)
	m.DhcpConfig = flattenIpamsvcDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.DhcpOptionsV6 = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptionsV6, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.HostnameRewriteChar = types.StringPointerValue(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = types.StringPointerValue(from.HostnameRewriteRegex)
	m.Id = types.StringPointerValue(from.Id)
	m.InheritanceSources = flattenIpamsvcIPSpaceInheritance(ctx, from.InheritanceSources, diags)
	m.Name = types.StringValue(from.Name)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = flattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = flattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = flattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)
	m.VendorSpecificOptionOptionSpace = types.StringPointerValue(from.VendorSpecificOptionOptionSpace)

}
