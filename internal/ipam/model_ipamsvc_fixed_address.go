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
	"tags":                         types.MapType{},
	"updated_at":                   timetypes.RFC3339Type{},
}

var IpamsvcFixedAddressResourceSchema = schema.Schema{
	MarkdownDescription: `A __FixedAddress__ object (_dhcp/fixed_address_) reserves an address for a specific client. It must have a _match_type_ and a valid corresponding _match_value_ so it can match that client.`,
	Attributes:          IpamsvcFixedAddressResourceSchemaAttributes,
}

var IpamsvcFixedAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The reserved address.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DHCP options. May be either a specific option or a group of options.`,
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. _true_ to disable object. The fixed address is converted to an exclusion when generating configuration.  Defaults to _false_.`,
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
	"hostname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_assigned_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritanceAssignedHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `The list of the inheritance assigned hosts of the object.`,
	},
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          IpamsvcFixedAddressInheritanceResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"match_type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Indicates how to match the client:  * _mac_: match the client MAC address for both IPv4 and IPv6,  * _client_text_ or _client_hex_: match the client identifier for IPv4 only,  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82) for IPv4 only,  * _duid_: match the DHCP unique identifier, currently match only for IPv6 protocol.`,
	},
	"match_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The value to match.`,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.`,
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the fixed address in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func expandIpamsvcFixedAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcFixedAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcFixedAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcFixedAddressModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcFixedAddress {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcFixedAddress{
		Address:                   m.Address.ValueString(),
		Comment:                   m.Comment.ValueStringPointer(),
		DhcpOptions:               ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		DisableDhcp:               m.DisableDhcp.ValueBoolPointer(),
		HeaderOptionFilename:      m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress: m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:    m.HeaderOptionServerName.ValueStringPointer(),
		Hostname:                  m.Hostname.ValueStringPointer(),
		InheritanceParent:         m.InheritanceParent.ValueStringPointer(),
		InheritanceSources:        expandIpamsvcFixedAddressInheritance(ctx, m.InheritanceSources, diags),
		IpSpace:                   m.IpSpace.ValueStringPointer(),
		MatchType:                 m.MatchType.ValueString(),
		MatchValue:                m.MatchValue.ValueString(),
		Name:                      m.Name.ValueStringPointer(),
		Parent:                    m.Parent.ValueStringPointer(),
		Tags:                      ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcFixedAddress(ctx context.Context, from *ipam.IpamsvcFixedAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcFixedAddressAttrTypes)
	}
	m := IpamsvcFixedAddressModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcFixedAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcFixedAddressModel) flatten(ctx context.Context, from *ipam.IpamsvcFixedAddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcFixedAddressModel{}
	}

	m.Address = types.StringValue(from.Address)
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.Hostname = types.StringPointerValue(from.Hostname)
	m.Id = types.StringPointerValue(from.Id)
	m.InheritanceAssignedHosts = FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, InheritanceAssignedHostAttrTypes, diags, flattenInheritanceAssignedHost)
	m.InheritanceParent = types.StringPointerValue(from.InheritanceParent)
	m.InheritanceSources = flattenIpamsvcFixedAddressInheritance(ctx, from.InheritanceSources, diags)
	m.IpSpace = types.StringPointerValue(from.IpSpace)
	m.MatchType = types.StringValue(from.MatchType)
	m.MatchValue = types.StringValue(from.MatchValue)
	m.Name = types.StringPointerValue(from.Name)
	m.Parent = types.StringPointerValue(from.Parent)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
