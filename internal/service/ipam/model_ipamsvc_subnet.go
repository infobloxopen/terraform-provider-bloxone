package ipam

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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

type IpamsvcSubnetModel struct {
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
	DhcpHost                   types.String      `tfsdk:"dhcp_host"`
	DhcpOptions                types.List        `tfsdk:"dhcp_options"`
	DhcpUtilization            types.Object      `tfsdk:"dhcp_utilization"`
	DisableDhcp                types.Bool        `tfsdk:"disable_dhcp"`
	DiscoveryAttrs             types.Map         `tfsdk:"discovery_attrs"`
	DiscoveryMetadata          types.Map         `tfsdk:"discovery_metadata"`
	HeaderOptionFilename       types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress  types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName     types.String      `tfsdk:"header_option_server_name"`
	HostnameRewriteChar        types.String      `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled     types.Bool        `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex       types.String      `tfsdk:"hostname_rewrite_regex"`
	Id                         types.String      `tfsdk:"id"`
	InheritanceAssignedHosts   types.List        `tfsdk:"inheritance_assigned_hosts"`
	InheritanceParent          types.String      `tfsdk:"inheritance_parent"`
	InheritanceSources         types.Object      `tfsdk:"inheritance_sources"`
	Name                       types.String      `tfsdk:"name"`
	Parent                     types.String      `tfsdk:"parent"`
	Protocol                   types.String      `tfsdk:"protocol"`
	RebindTime                 types.Int64       `tfsdk:"rebind_time"`
	RenewTime                  types.Int64       `tfsdk:"renew_time"`
	Space                      types.String      `tfsdk:"space"`
	Tags                       types.Map         `tfsdk:"tags"`
	Threshold                  types.Object      `tfsdk:"threshold"`
	UpdatedAt                  timetypes.RFC3339 `tfsdk:"updated_at"`
	Usage                      types.List        `tfsdk:"usage"`
	Utilization                types.Object      `tfsdk:"utilization"`
	UtilizationV6              types.Object      `tfsdk:"utilization_v6"`
	NextAvailableId            types.String      `tfsdk:"next_available_id"`
}

var IpamsvcSubnetAttrTypes = map[string]attr.Type{
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
	"dhcp_host":                     types.StringType,
	"dhcp_options":                  types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"dhcp_utilization":              types.ObjectType{AttrTypes: IpamsvcDHCPUtilizationAttrTypes},
	"disable_dhcp":                  types.BoolType,
	"discovery_attrs":               types.MapType{ElemType: types.StringType},
	"discovery_metadata":            types.MapType{ElemType: types.StringType},
	"header_option_filename":        types.StringType,
	"header_option_server_address":  types.StringType,
	"header_option_server_name":     types.StringType,
	"hostname_rewrite_char":         types.StringType,
	"hostname_rewrite_enabled":      types.BoolType,
	"hostname_rewrite_regex":        types.StringType,
	"id":                            types.StringType,
	"inheritance_assigned_hosts":    types.ListType{ElemType: types.ObjectType{AttrTypes: InheritanceAssignedHostAttrTypes}},
	"inheritance_parent":            types.StringType,
	"inheritance_sources":           types.ObjectType{AttrTypes: IpamsvcDHCPInheritanceAttrTypes},
	"name":                          types.StringType,
	"parent":                        types.StringType,
	"protocol":                      types.StringType,
	"rebind_time":                   types.Int64Type,
	"renew_time":                    types.Int64Type,
	"space":                         types.StringType,
	"tags":                          types.MapType{ElemType: types.StringType},
	"threshold":                     types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                    timetypes.RFC3339Type{},
	"usage":                         types.ListType{ElemType: types.StringType},
	"utilization":                   types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":                types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
	"next_available_id":             types.StringType,
}

var IpamsvcSubnetResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address of the subnet in the form “a.b.c.d”",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
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
		MarkdownDescription: "Set to 1 to indicate that the subnet may run out of addresses.",
	},
	"cidr": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the subnet. This is required if _address_ does not include CIDR.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8.",
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
			"  * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n" +
			"  Defaults to _client_.\n",
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
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty.",
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_.",
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("myhost"),
		MarkdownDescription: "The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \"myhost\".",
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if DDNS updates are enabled at the subnet level. Defaults to _true_.",
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		MarkdownDescription: "DDNS TTL value - to be calculated as a simple percentage of the lease's lifetime, using the parameter's value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.",
	},
	"ddns_update_on_renew": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_.",
	},
	"ddns_use_conflict_resolution": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_.",
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPConfigResourceSchemaAttributes(true),
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(IpamsvcDHCPConfigAttrTypes, map[string]attr.Value{
			"abandoned_reclaim_time":    types.Int64Null(), // abandoned_reclaim_time cannot be set for subnet
			"abandoned_reclaim_time_v6": types.Int64Null(), // abandoned_reclaim_time_v6 cannot be set for subnet
			"allow_unknown":             types.BoolValue(true),
			"allow_unknown_v6":          types.BoolValue(true),
			"echo_client_id":            types.BoolNull(), // echo_id cannot be set for subnet
			"filters":                   types.ListNull(types.StringType),
			"filters_v6":                types.ListNull(types.StringType),
			"ignore_client_uid":         types.BoolValue(false),
			"ignore_list":               types.ListNull(types.ObjectType{AttrTypes: IpamsvcIgnoreItemAttrTypes}),
			"lease_time":                types.Int64Value(3600),
			"lease_time_v6":             types.Int64Value(3600),
		})),
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The resource identifier.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The DHCP options of the subnet. This can either be a specific option or a group of options.",
	},
	"dhcp_utilization": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPUtilizationResourceSchemaAttributes,
		Computed:   true,
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
	},
	"discovery_attrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The discovery attributes for this subnet in JSON format.",
	},
	"discovery_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The discovery metadata for this subnet in JSON format.",
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
		MarkdownDescription: "The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \"-\".",
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.",
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("[^a-zA-Z0-9_.]"),
		MarkdownDescription: "The regex bracket expression to match valid characters.  Must begin with \"[\" and end with \"]\" and be a compilable POSIX regex.  Defaults to \"[^a-zA-Z0-9_.]\".",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_assigned_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritanceAssignedHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of the inheritance assigned hosts of the object.",
	},
	"inheritance_parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the subnet. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of the subnet (_ip4_ or _ip6_).",
	},
	"rebind_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease rebind time (T2) in seconds.",
	},
	"renew_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease renew time (T1) in seconds.",
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the subnet in JSON format.",
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationThresholdResourceSchemaAttributes,
		Computed:   true,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"usage": schema.ListAttribute{
		ElementType: types.StringType,
		Computed:    true,
		MarkdownDescription: "The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:\n\n" +
			"  | usage indicator        | description                                                                                   |\n" +
			"  |------------------------|-----------------------------------------------------------------------------------------------|\n" +
			"  | _IPAM_                 | Subnet is managed in BloxOne DDI.                                                             |\n" +
			"  | _DHCP_                 | Subnet is served by a DHCP Host.                                                              |\n" +
			"  | _DISCOVERED_           | Subnet is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.  |\n" +
			"  <br>",
	},
	"utilization": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationResourceSchemaAttributes,
		Computed:   true,
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationV6ResourceSchemaAttributes,
		Computed:   true,
	},
	"next_available_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier for the address block where the next available subnet should be generated",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
			stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`), "Should be the resource identifier of an address block."),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
	},
}

func (m *IpamsvcSubnetModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Subnet {
	if m == nil {
		return nil
	}
	to := &ipam.Subnet{
		AsmConfig:                  ExpandIpamsvcASMConfig(ctx, m.AsmConfig, diags),
		Cidr:                       flex.ExpandInt64Pointer(m.Cidr),
		Comment:                    flex.ExpandStringPointer(m.Comment),
		DdnsClientUpdate:           flex.ExpandStringPointer(m.DdnsClientUpdate),
		DdnsConflictResolutionMode: flex.ExpandStringPointer(m.DdnsConflictResolutionMode),
		DdnsDomain:                 flex.ExpandStringPointer(m.DdnsDomain),
		DdnsGenerateName:           flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix:        flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
		DdnsSendUpdates:            flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsTtlPercent:             flex.ExpandFloat32Pointer(m.DdnsTtlPercent),
		DdnsUpdateOnRenew:          flex.ExpandBoolPointer(m.DdnsUpdateOnRenew),
		DdnsUseConflictResolution:  flex.ExpandBoolPointer(m.DdnsUseConflictResolution),
		DhcpConfig:                 ExpandIpamsvcDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpHost:                   flex.ExpandStringPointer(m.DhcpHost),
		DhcpOptions:                flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		DhcpUtilization:            ExpandIpamsvcDHCPUtilization(ctx, m.DhcpUtilization, diags),
		DisableDhcp:                flex.ExpandBoolPointer(m.DisableDhcp),
		HeaderOptionFilename:       flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:  flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:     flex.ExpandStringPointer(m.HeaderOptionServerName),
		HostnameRewriteChar:        flex.ExpandStringPointer(m.HostnameRewriteChar),
		HostnameRewriteEnabled:     flex.ExpandBoolPointer(m.HostnameRewriteEnabled),
		HostnameRewriteRegex:       flex.ExpandStringPointer(m.HostnameRewriteRegex),
		InheritanceParent:          flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources:         ExpandIpamsvcDHCPInheritance(ctx, m.InheritanceSources, diags),
		Name:                       flex.ExpandStringPointer(m.Name),
		Parent:                     flex.ExpandStringPointer(m.Parent),
		RebindTime:                 flex.ExpandInt64Pointer(m.RebindTime),
		RenewTime:                  flex.ExpandInt64Pointer(m.RenewTime),
		Tags:                       flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}

	if isCreate {
		if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() {
			nasId := flex.ExpandString(m.NextAvailableId) + "/nextavailablesubnet"
			to.Address = &nasId
		} else {
			to.Address = flex.ExpandStringPointer(m.Address)
		}
		to.Space = flex.ExpandStringPointer(m.Space)
	}

	return to
}

func FlattenIpamsvcSubnet(ctx context.Context, from *ipam.Subnet, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcSubnetAttrTypes)
	}
	m := IpamsvcSubnetModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcSubnetAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcSubnetModel) Flatten(ctx context.Context, from *ipam.Subnet, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcSubnetModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AsmConfig = FlattenIpamsvcASMConfig(ctx, from.AsmConfig, diags)
	m.AsmScopeFlag = flex.FlattenInt64Pointer(from.AsmScopeFlag)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
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
	m.DhcpConfig = FlattenIpamsvcDHCPConfigForSubnetOrAddressBlock(ctx, from.DhcpConfig, diags)
	m.DhcpHost = flex.FlattenStringPointerWithNilAsEmpty(from.DhcpHost)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.DhcpUtilization = FlattenIpamsvcDHCPUtilization(ctx, from.DhcpUtilization, diags)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.DiscoveryAttrs = flex.FlattenFrameworkMapString(ctx, from.DiscoveryAttrs, diags)
	m.DiscoveryMetadata = flex.FlattenFrameworkMapString(ctx, from.DiscoveryMetadata, diags)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceAssignedHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, InheritanceAssignedHostAttrTypes, diags, FlattenInheritanceAssignedHost)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenIpamsvcDHCPInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.RebindTime = flex.FlattenInt64Pointer(from.RebindTime)
	m.RenewTime = flex.FlattenInt64Pointer(from.RenewTime)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = FlattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Usage = flex.FlattenFrameworkListString(ctx, from.Usage, diags)
	m.Utilization = FlattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = FlattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)
}
