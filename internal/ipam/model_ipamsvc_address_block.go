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

type IpamsvcAddressBlockModel struct {
	Address                    types.String      `tfsdk:"address"`
	AsmConfig                  types.Object      `tfsdk:"asm_config"`
	AsmScopeFlag               types.Int64       `tfsdk:"asm_scope_flag"`
	Cidr                       types.Int64       `tfsdk:"cidr"`
	Comment                    types.String      `tfsdk:"comment"`
	CreatedAt                  timetypes.RFC3339 `tfsdk:"created_at"`
	DdnsClientUpdate           types.String      `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode types.String      `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                 types.String      `tfsdk:"ddns_domain"`
	DdnsGenerateName           types.Bool        `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix        types.String      `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates            types.Bool        `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent             types.Float64     `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew          types.Bool        `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution  types.Bool        `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                 types.Object      `tfsdk:"dhcp_config"`
	DhcpOptions                types.List        `tfsdk:"dhcp_options"`
	DhcpUtilization            types.Object      `tfsdk:"dhcp_utilization"`
	DiscoveryAttrs             types.Map         `tfsdk:"discovery_attrs"`
	DiscoveryMetadata          types.Map         `tfsdk:"discovery_metadata"`
	HeaderOptionFilename       types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress  types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName     types.String      `tfsdk:"header_option_server_name"`
	HostnameRewriteChar        types.String      `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled     types.Bool        `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex       types.String      `tfsdk:"hostname_rewrite_regex"`
	Id                         types.String      `tfsdk:"id"`
	InheritanceParent          types.String      `tfsdk:"inheritance_parent"`
	InheritanceSources         types.Object      `tfsdk:"inheritance_sources"`
	Name                       types.String      `tfsdk:"name"`
	Parent                     types.String      `tfsdk:"parent"`
	Protocol                   types.String      `tfsdk:"protocol"`
	Space                      types.String      `tfsdk:"space"`
	Tags                       types.Map         `tfsdk:"tags"`
	Threshold                  types.Object      `tfsdk:"threshold"`
	UpdatedAt                  timetypes.RFC3339 `tfsdk:"updated_at"`
	Usage                      types.List        `tfsdk:"usage"`
	Utilization                types.Object      `tfsdk:"utilization"`
	UtilizationV6              types.Object      `tfsdk:"utilization_v6"`
}

var IpamsvcAddressBlockAttrTypes = map[string]attr.Type{
	"address":                       types.StringType,
	"asm_config":                    types.ObjectType{AttrTypes: IpamsvcASMConfigAttrTypes},
	"asm_scope_flag":                types.Int64Type,
	"cidr":                          types.Int64Type,
	"comment":                       types.StringType,
	"created_at":                    timetypes.RFC3339Type{},
	"ddns_client_update":            types.StringType,
	"ddns_conflict_resolution_mode": types.StringType,
	"ddns_domain":                   types.StringType,
	"ddns_generate_name":            types.BoolType,
	"ddns_generated_prefix":         types.StringType,
	"ddns_send_updates":             types.BoolType,
	"ddns_ttl_percent":              types.Float64Type,
	"ddns_update_on_renew":          types.BoolType,
	"ddns_use_conflict_resolution":  types.BoolType,
	"dhcp_config":                   types.ObjectType{AttrTypes: IpamsvcDHCPConfigAttrTypes},
	"dhcp_options":                  types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_utilization":              types.ObjectType{AttrTypes: IpamsvcDHCPUtilizationAttrTypes},
	"discovery_attrs":               types.MapType{},
	"discovery_metadata":            types.MapType{},
	"header_option_filename":        types.StringType,
	"header_option_server_address":  types.StringType,
	"header_option_server_name":     types.StringType,
	"hostname_rewrite_char":         types.StringType,
	"hostname_rewrite_enabled":      types.BoolType,
	"hostname_rewrite_regex":        types.StringType,
	"id":                            types.StringType,
	"inheritance_parent":            types.StringType,
	"inheritance_sources":           types.ObjectType{AttrTypes: IpamsvcDHCPInheritanceAttrTypes},
	"name":                          types.StringType,
	"parent":                        types.StringType,
	"protocol":                      types.StringType,
	"space":                         types.StringType,
	"tags":                          types.MapType{},
	"threshold":                     types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                    timetypes.RFC3339Type{},
	"usage":                         types.ListType{ElemType: types.StringType},
	"utilization":                   types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":                types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
}

var IpamsvcAddressBlockResourceSchema = schema.Schema{
	MarkdownDescription: `An __AddressBlock__ object (_ipam/address_block_) is a set of contiguous IP addresses in the same IP space with no gap, expressed as a CIDR block. Address blocks are hierarchical and may be parented to other address blocks as long as the parent block fully contains the child and no sibling overlaps. Top level address blocks are parented to an IP space.`,
	Attributes:          IpamsvcAddressBlockResourceSchemaAttributes,
}

var IpamsvcAddressBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.`,
	},
	"asm_config": schema.SingleNestedAttribute{
		Attributes:          IpamsvcASMConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"asm_scope_flag": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `Incremented by 1 if the IP address usage limits for automated scope management are exceeded for any subnets in the address block.`,
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The CIDR of the address block. This is required, if _address_ does not specify it in its input.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the address block. May contain 0 to 1024 characters. Can include UTF-8.`,
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
		MarkdownDescription: `Determines if DDNS updates are enabled at the address block level. Defaults to _true_.`,
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
		MarkdownDescription: `The list of DHCP options for the address block. May be either a specific option or a group of options.`,
	},
	"dhcp_utilization": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPUtilizationResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"discovery_attrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The discovery attributes for this address block in JSON format.`,
	},
	"discovery_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The discovery metadata for this address block in JSON format.`,
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
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPInheritanceResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the address block. May contain 1 to 256 characters. Can include UTF-8.`,
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of protocol of address block (_ip4_ or _ip6_).`,
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the address block in JSON format.`,
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
	"usage": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: `The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  AddressBlock is managed in BloxOne DDI.  _DISCOVERED_           |  AddressBlock is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.`,
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
}

func expandIpamsvcAddressBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAddressBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcAddressBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcAddressBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAddressBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcAddressBlock{
		Address:                    m.Address.ValueString(),
		AsmConfig:                  expandIpamsvcASMConfig(ctx, m.AsmConfig, diags),
		Cidr:                       ptr(int64(m.Cidr.ValueInt64())),
		Comment:                    m.Comment.ValueStringPointer(),
		DdnsClientUpdate:           m.DdnsClientUpdate.ValueStringPointer(),
		DdnsConflictResolutionMode: m.DdnsConflictResolutionMode.ValueStringPointer(),
		DdnsDomain:                 m.DdnsDomain.ValueStringPointer(),
		DdnsGenerateName:           m.DdnsGenerateName.ValueBoolPointer(),
		DdnsGeneratedPrefix:        m.DdnsGeneratedPrefix.ValueStringPointer(),
		DdnsSendUpdates:            m.DdnsSendUpdates.ValueBoolPointer(),
		DdnsTtlPercent:             ptr(float32(m.DdnsTtlPercent.ValueFloat64())),
		DdnsUpdateOnRenew:          m.DdnsUpdateOnRenew.ValueBoolPointer(),
		DdnsUseConflictResolution:  m.DdnsUseConflictResolution.ValueBoolPointer(),
		DhcpConfig:                 expandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		DhcpUtilization:            expandIpamsvcDHCPUtilization(ctx, m.DhcpUtilization, diags),
		DiscoveryAttrs:             ExpandFrameworkMapString(ctx, m.DiscoveryAttrs, diags),
		DiscoveryMetadata:          ExpandFrameworkMapString(ctx, m.DiscoveryMetadata, diags),
		HeaderOptionFilename:       m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:  m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:     m.HeaderOptionServerName.ValueStringPointer(),
		HostnameRewriteChar:        m.HostnameRewriteChar.ValueStringPointer(),
		HostnameRewriteEnabled:     m.HostnameRewriteEnabled.ValueBoolPointer(),
		HostnameRewriteRegex:       m.HostnameRewriteRegex.ValueStringPointer(),
		InheritanceParent:          m.InheritanceParent.ValueStringPointer(),
		InheritanceSources:         expandIpamsvcDHCPInheritance(ctx, m.InheritanceSources, diags),
		Name:                       m.Name.ValueStringPointer(),
		Parent:                     m.Parent.ValueStringPointer(),
		Space:                      m.Space.ValueString(),
		Tags:                       ExpandFrameworkMapString(ctx, m.Tags, diags),
		Threshold:                  expandIpamsvcUtilizationThreshold(ctx, m.Threshold, diags),
		Utilization:                expandIpamsvcUtilization(ctx, m.Utilization, diags),
		UtilizationV6:              expandIpamsvcUtilizationV6(ctx, m.UtilizationV6, diags),
	}
	return to
}

func flattenIpamsvcAddressBlock(ctx context.Context, from *ipam.IpamsvcAddressBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAddressBlockAttrTypes)
	}
	m := IpamsvcAddressBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAddressBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAddressBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcAddressBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAddressBlockModel{}
	}

	m.Address = types.StringValue(from.Address)
	m.AsmConfig = flattenIpamsvcASMConfig(ctx, from.AsmConfig, diags)
	m.AsmScopeFlag = types.Int64Value(int64(*from.AsmScopeFlag))
	m.Cidr = types.Int64Value(int64(*from.Cidr))
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
	m.DhcpUtilization = flattenIpamsvcDHCPUtilization(ctx, from.DhcpUtilization, diags)
	m.DiscoveryAttrs = FlattenFrameworkMapString(ctx, from.DiscoveryAttrs, diags)
	m.DiscoveryMetadata = FlattenFrameworkMapString(ctx, from.DiscoveryMetadata, diags)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.HostnameRewriteChar = types.StringPointerValue(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = types.StringPointerValue(from.HostnameRewriteRegex)
	m.Id = types.StringPointerValue(from.Id)
	m.InheritanceParent = types.StringPointerValue(from.InheritanceParent)
	m.InheritanceSources = flattenIpamsvcDHCPInheritance(ctx, from.InheritanceSources, diags)
	m.Name = types.StringPointerValue(from.Name)
	m.Parent = types.StringPointerValue(from.Parent)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Space = types.StringValue(from.Space)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = flattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Usage = FlattenFrameworkListString(ctx, from.Usage, diags)
	m.Utilization = flattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = flattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)

}
