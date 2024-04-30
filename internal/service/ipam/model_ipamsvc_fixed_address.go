package ipam

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcFixedAddressModel struct {
	Address                   types.String      `tfsdk:"address"`
	Comment                   types.String      `tfsdk:"comment"`
	CreatedAt                 timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions               types.List        `tfsdk:"dhcp_options"`
	DisableDhcp               types.Bool        `tfsdk:"disable_dhcp"`
	HeaderOptionFilename      types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName    types.String      `tfsdk:"header_option_server_name"`
	Hostname                  types.String      `tfsdk:"hostname"`
	Id                        types.String      `tfsdk:"id"`
	InheritanceAssignedHosts  types.List        `tfsdk:"inheritance_assigned_hosts"`
	InheritanceParent         types.String      `tfsdk:"inheritance_parent"`
	InheritanceSources        types.Object      `tfsdk:"inheritance_sources"`
	IpSpace                   types.String      `tfsdk:"ip_space"`
	MatchType                 types.String      `tfsdk:"match_type"`
	MatchValue                types.String      `tfsdk:"match_value"`
	Name                      types.String      `tfsdk:"name"`
	Parent                    types.String      `tfsdk:"parent"`
	Tags                      types.Map         `tfsdk:"tags"`
	UpdatedAt                 timetypes.RFC3339 `tfsdk:"updated_at"`
	NextAvailableId           types.String      `tfsdk:"next_available_id"`
}

var IpamsvcFixedAddressAttrTypes = map[string]attr.Type{
	"address":                      types.StringType,
	"comment":                      types.StringType,
	"created_at":                   timetypes.RFC3339Type{},
	"dhcp_options":                 types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"disable_dhcp":                 types.BoolType,
	"header_option_filename":       types.StringType,
	"header_option_server_address": types.StringType,
	"header_option_server_name":    types.StringType,
	"hostname":                     types.StringType,
	"id":                           types.StringType,
	"inheritance_assigned_hosts":   types.ListType{ElemType: types.ObjectType{AttrTypes: InheritanceAssignedHostAttrTypes}},
	"inheritance_parent":           types.StringType,
	"inheritance_sources":          types.ObjectType{AttrTypes: IpamsvcFixedAddressInheritanceAttrTypes},
	"ip_space":                     types.StringType,
	"match_type":                   types.StringType,
	"match_value":                  types.StringType,
	"name":                         types.StringType,
	"parent":                       types.StringType,
	"tags":                         types.MapType{ElemType: types.StringType},
	"updated_at":                   timetypes.RFC3339Type{},
	"next_available_id":            types.StringType,
}

var IpamsvcFixedAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The reserved address.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP options. May be either a specific option or a group of options.",
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. The fixed address is converted to an exclusion when generating configuration.  Defaults to _false_.",
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
	"hostname": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.",
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
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: IpamsvcFixedAddressInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"match_type": schema.StringAttribute{
		Required: true,
		MarkdownDescription: "Indicates how to match the client:\n" +
			"  * _mac_: match the client MAC address for both IPv4 and IPv6\n" +
			"  * _client_text_ or _client_hex_: match the client identifier for IPv4 only\n" +
			"  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82) for IPv4 only\n" +
			"  * _duid_: match the DHCP unique identifier, currently match only for IPv6 protocol",
	},
	"match_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The value to match.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the fixed address in JSON format.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"next_available_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier for the address block where the next available fixed address should be generated",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
			stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(subnet|range)/[0-9a-f-].*$`), "Should be the resource identifier of an subnet or range."),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
	},
}

func ExpandIpamsvcFixedAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.FixedAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcFixedAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcFixedAddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.FixedAddress {
	if m == nil {
		return nil
	}
	to := &ipam.FixedAddress{
		Comment:                   flex.ExpandStringPointer(m.Comment),
		DhcpOptions:               flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		DisableDhcp:               flex.ExpandBoolPointer(m.DisableDhcp),
		HeaderOptionFilename:      flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress: flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:    flex.ExpandStringPointer(m.HeaderOptionServerName),
		Hostname:                  flex.ExpandStringPointer(m.Hostname),
		InheritanceParent:         flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources:        ExpandIpamsvcFixedAddressInheritance(ctx, m.InheritanceSources, diags),
		IpSpace:                   flex.ExpandStringPointer(m.IpSpace),
		MatchType:                 flex.ExpandString(m.MatchType),
		MatchValue:                flex.ExpandString(m.MatchValue),
		Name:                      flex.ExpandStringPointer(m.Name),
		Parent:                    flex.ExpandStringPointer(m.Parent),
		Tags:                      flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	to.Address = flex.ExpandString(m.Address)
	if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() {
		to.Address = flex.ExpandString(m.NextAvailableId) + "/nextavailableip"
	}

	return to
}

func FlattenIpamsvcFixedAddress(ctx context.Context, from *ipam.FixedAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcFixedAddressAttrTypes)
	}
	m := IpamsvcFixedAddressModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcFixedAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcFixedAddressModel) Flatten(ctx context.Context, from *ipam.FixedAddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcFixedAddressModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.Hostname = flex.FlattenStringPointer(from.Hostname)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceAssignedHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, InheritanceAssignedHostAttrTypes, diags, FlattenInheritanceAssignedHost)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenIpamsvcFixedAddressInheritance(ctx, from.InheritanceSources, diags)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.MatchType = flex.FlattenString(from.MatchType)
	m.MatchValue = flex.FlattenString(from.MatchValue)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
