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

type IpamsvcAddressModel struct {
	Address           types.String      `tfsdk:"address"`
	Comment           types.String      `tfsdk:"comment"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpInfo          types.Object      `tfsdk:"dhcp_info"`
	DisableDhcp       types.Bool        `tfsdk:"disable_dhcp"`
	DiscoveryAttrs    types.Map         `tfsdk:"discovery_attrs"`
	DiscoveryMetadata types.Map         `tfsdk:"discovery_metadata"`
	Host              types.String      `tfsdk:"host"`
	Hwaddr            types.String      `tfsdk:"hwaddr"`
	Id                types.String      `tfsdk:"id"`
	Interface         types.String      `tfsdk:"interface"`
	Names             types.List        `tfsdk:"names"`
	Parent            types.String      `tfsdk:"parent"`
	Protocol          types.String      `tfsdk:"protocol"`
	Range             types.String      `tfsdk:"range"`
	Space             types.String      `tfsdk:"space"`
	State             types.String      `tfsdk:"state"`
	Tags              types.Map         `tfsdk:"tags"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
	Usage             types.List        `tfsdk:"usage"`
}

var IpamsvcAddressAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"comment":            types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"dhcp_info":          types.ObjectType{AttrTypes: IpamsvcDHCPInfoAttrTypes},
	"disable_dhcp":       types.BoolType,
	"discovery_attrs":    types.MapType{},
	"discovery_metadata": types.MapType{},
	"host":               types.StringType,
	"hwaddr":             types.StringType,
	"id":                 types.StringType,
	"interface":          types.StringType,
	"names":              types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcNameAttrTypes}},
	"parent":             types.StringType,
	"protocol":           types.StringType,
	"range":              types.StringType,
	"space":              types.StringType,
	"state":              types.StringType,
	"tags":               types.MapType{},
	"updated_at":         timetypes.RFC3339Type{},
	"usage":              types.ListType{ElemType: types.StringType},
}

var IpamsvcAddressResourceSchema = schema.Schema{
	MarkdownDescription: `An __Address__ object (_ipam/address_) represents any single IP address within a given IP space.`,
	Attributes:          IpamsvcAddressResourceSchemaAttributes,
}

var IpamsvcAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The address in form \&quot;a.b.c.d\&quot;.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"dhcp_info": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPInfoResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"disable_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: `Read only. Represent the value of the same field in the associated _dhcp/fixed_address_ object.`,
	},
	"discovery_attrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The discovery attributes for this address in JSON format.`,
	},
	"discovery_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The discovery metadata for this address in JSON format.`,
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"hwaddr": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The hardware address associated with this IP address.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the network interface card (NIC) associated with the address, if any.`,
	},
	"names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcNameResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of all names associated with this address.`,
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of protocol (_ip4_ or _ip6_).`,
	},
	"range": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The state of the address (_used_ or _free_).`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for this address in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"usage": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: `The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  Address was created by the IPAM component.  _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.  _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.  _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.  _DHCP_                 |  Address was created by the DHCP component.  _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.  _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.  _DHCP_, _DISABLED_     |  Address is disabled.  _DNS_                  |  Address is used by one or more DNS records.  _DISCOVERED_           |  Address is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.`,
	},
}

func expandIpamsvcAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcAddressModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAddress {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcAddress{
		Address:           m.Address.ValueString(),
		Comment:           m.Comment.ValueStringPointer(),
		DhcpInfo:          expandIpamsvcDHCPInfo(ctx, m.DhcpInfo, diags),
		DiscoveryAttrs:    ExpandFrameworkMapString(ctx, m.DiscoveryAttrs, diags),
		DiscoveryMetadata: ExpandFrameworkMapString(ctx, m.DiscoveryMetadata, diags),
		Host:              m.Host.ValueStringPointer(),
		Hwaddr:            m.Hwaddr.ValueStringPointer(),
		Interface:         m.Interface.ValueStringPointer(),
		Names:             ExpandFrameworkListNestedBlock(ctx, m.Names, diags, expandIpamsvcName),
		Parent:            m.Parent.ValueStringPointer(),
		Range:             m.Range.ValueStringPointer(),
		Space:             m.Space.ValueString(),
		Tags:              ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcAddress(ctx context.Context, from *ipam.IpamsvcAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAddressAttrTypes)
	}
	m := IpamsvcAddressModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAddressModel) flatten(ctx context.Context, from *ipam.IpamsvcAddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAddressModel{}
	}

	m.Address = types.StringValue(from.Address)
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpInfo = flattenIpamsvcDHCPInfo(ctx, from.DhcpInfo, diags)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.DiscoveryAttrs = FlattenFrameworkMapString(ctx, from.DiscoveryAttrs, diags)
	m.DiscoveryMetadata = FlattenFrameworkMapString(ctx, from.DiscoveryMetadata, diags)
	m.Host = types.StringPointerValue(from.Host)
	m.Hwaddr = types.StringPointerValue(from.Hwaddr)
	m.Id = types.StringPointerValue(from.Id)
	m.Interface = types.StringPointerValue(from.Interface)
	m.Names = FlattenFrameworkListNestedBlock(ctx, from.Names, IpamsvcNameAttrTypes, diags, flattenIpamsvcName)
	m.Parent = types.StringPointerValue(from.Parent)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Range = types.StringPointerValue(from.Range)
	m.Space = types.StringValue(from.Space)
	m.State = types.StringPointerValue(from.State)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Usage = FlattenFrameworkListString(ctx, from.Usage, diags)

}
