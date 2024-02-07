package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
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
	"tags":                                types.MapType{ElemType: types.StringType},
	"threshold":                           types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                          timetypes.RFC3339Type{},
	"utilization":                         types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":                      types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcIPSpaceResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcASMConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(IpamsvcASMConfigAttrTypes, map[string]attr.Value{
			"asm_threshold":       types.Int64Value(90),
			"enable":              types.BoolValue(true),
			"enable_notification": types.BoolValue(true),
			"forecast_period":     types.Int64Value(14),
			"growth_factor":       types.Int64Value(20),
			"growth_type":         types.StringValue("percent"),
			"history":             types.Int64Value(30),
			"min_total":           types.Int64Value(10),
			"min_unused":          types.Int64Value(10),
			"reenable_date":       timetypes.NewRFC3339ValueMust("1970-01-01T00:00:00Z"),
		})),
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
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("client"),
		MarkdownDescription: "Controls who does the DDNS updates. Valid values are:\n" +
			"  * _client_: DHCP server updates DNS if requested by client.\n" +
			"  * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n" +
			"  * _ignore_: DHCP server always updates DNS, even if the client says not to.\n" +
			"  * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n" +
			"  * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n" +
			"  Defaults to _client_.\n",
		Validators: []validator.String{
			stringvalidator.OneOf("client", "server", "ignore", "over_client_update", "over_no_update"),
		},
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("check_with_dhcid"),
		MarkdownDescription: "The mode used for resolving conflicts while performing DDNS updates. Valid values are:\n\n" +
			"  * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703.\n" +
			"  * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry.\n" +
			"  * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry.\n" +
			"  * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.\n" +
			"  Defaults to _check_with_dhcid_.",
		Validators: []validator.String{
			stringvalidator.OneOf("check_with_dhcid", "no_check_with_dhcid", "check_exists_with_dhcid"),
		},
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty.`,
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_.`,
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("myhost"),
		MarkdownDescription: `The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \"myhost\".`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Determines if DDNS updates are enabled at the IP space level. Defaults to _true_.`,
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		MarkdownDescription: `DDNS TTL value - to be calculated as a simple percentage of the lease's lifetime, using the parameter's value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.`,
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
		MarkdownDescription: `When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_.`,
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(IpamsvcDHCPConfigAttrTypes, map[string]attr.Value{
			"abandoned_reclaim_time":    types.Int64Value(3600),
			"abandoned_reclaim_time_v6": types.Int64Value(3600),
			"allow_unknown":             types.BoolValue(true),
			"allow_unknown_v6":          types.BoolValue(true),
			"filters":                   types.ListNull(types.StringType),
			"filters_v6":                types.ListNull(types.StringType),
			"ignore_client_uid":         types.BoolValue(false),
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
		MarkdownDescription: `The list of IPv4 DHCP options for IP space. May be either a specific option or a group of options.`,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"dhcp_options_v6": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of IPv6 DHCP options for IP space. May be either a specific option or a group of options.`,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
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
		Attributes: IpamsvcIPSpaceInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
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
		Attributes: IpamsvcUtilizationThresholdResourceSchemaAttributes,
		Computed:   true,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"utilization": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationResourceSchemaAttributes,
		Computed:   true,
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationV6ResourceSchemaAttributes,
		Computed:   true,
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandIpamsvcIPSpace(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcIPSpace {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcIPSpaceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcIPSpaceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcIPSpace {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcIPSpace{
		AsmConfig:                       ExpandIpamsvcASMConfig(ctx, m.AsmConfig, diags),
		Comment:                         m.Comment.ValueStringPointer(),
		DdnsClientUpdate:                m.DdnsClientUpdate.ValueStringPointer(),
		DdnsConflictResolutionMode:      m.DdnsConflictResolutionMode.ValueStringPointer(),
		DdnsDomain:                      m.DdnsDomain.ValueStringPointer(),
		DdnsGenerateName:                m.DdnsGenerateName.ValueBoolPointer(),
		DdnsGeneratedPrefix:             m.DdnsGeneratedPrefix.ValueStringPointer(),
		DdnsSendUpdates:                 m.DdnsSendUpdates.ValueBoolPointer(),
		DdnsTtlPercent:                  utils.Ptr(float32(m.DdnsTtlPercent.ValueFloat64())),
		DdnsUpdateOnRenew:               m.DdnsUpdateOnRenew.ValueBoolPointer(),
		DdnsUseConflictResolution:       m.DdnsUseConflictResolution.ValueBoolPointer(),
		DhcpConfig:                      ExpandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		DhcpOptionsV6:                   flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptionsV6, diags, ExpandIpamsvcOptionItem),
		HeaderOptionFilename:            m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:       m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:          m.HeaderOptionServerName.ValueStringPointer(),
		HostnameRewriteChar:             m.HostnameRewriteChar.ValueStringPointer(),
		HostnameRewriteEnabled:          m.HostnameRewriteEnabled.ValueBoolPointer(),
		HostnameRewriteRegex:            m.HostnameRewriteRegex.ValueStringPointer(),
		InheritanceSources:              ExpandIpamsvcIPSpaceInheritance(ctx, m.InheritanceSources, diags),
		Name:                            m.Name.ValueString(),
		Tags:                            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: m.VendorSpecificOptionOptionSpace.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcIPSpace(ctx context.Context, from *ipam.IpamsvcIPSpace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIPSpaceAttrTypes)
	}
	m := IpamsvcIPSpaceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIPSpaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIPSpaceModel) Flatten(ctx context.Context, from *ipam.IpamsvcIPSpace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIPSpaceModel{}
	}
	m.AsmConfig = FlattenIpamsvcASMConfig(ctx, from.AsmConfig, diags)
	m.AsmScopeFlag = flex.FlattenInt64(int64(*from.AsmScopeFlag))
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DdnsClientUpdate = flex.FlattenStringPointer(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = flex.FlattenStringPointer(from.DdnsConflictResolutionMode)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsTtlPercent = flex.FlattenFloat64(float64(*from.DdnsTtlPercent))
	m.DdnsUpdateOnRenew = types.BoolPointerValue(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = types.BoolPointerValue(from.DdnsUseConflictResolution)
	m.DhcpConfig = FlattenIpamsvcDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.DhcpOptionsV6 = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptionsV6, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenIpamsvcIPSpaceInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenString(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = FlattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = FlattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = FlattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)
	m.VendorSpecificOptionOptionSpace = flex.FlattenStringPointer(from.VendorSpecificOptionOptionSpace)
}
