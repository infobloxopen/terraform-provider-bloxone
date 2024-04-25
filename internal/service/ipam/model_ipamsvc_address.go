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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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
	NextAvailableId   types.String      `tfsdk:"next_available_id"`
}

var IpamsvcAddressAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"comment":            types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"dhcp_info":          types.ObjectType{AttrTypes: IpamsvcDHCPInfoAttrTypes},
	"disable_dhcp":       types.BoolType,
	"discovery_attrs":    types.MapType{ElemType: types.StringType},
	"discovery_metadata": types.MapType{ElemType: types.StringType},
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
	"tags":               types.MapType{ElemType: types.StringType},
	"updated_at":         timetypes.RFC3339Type{},
	"usage":              types.ListType{ElemType: types.StringType},
	"next_available_id":  types.StringType,
}

var IpamsvcAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address in form \"a.b.c.d\".",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
		},
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"dhcp_info": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPInfoResourceSchemaAttributes,
		Computed:   true,
	},
	"disable_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Read only. Represent the value of the same field in the associated _dhcp/fixed_address_ object.",
	},
	"discovery_attrs": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The discovery attributes for this address in JSON format.",
	},
	"discovery_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The discovery metadata for this address in JSON format.",
	},
	"host": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"hwaddr": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The hardware address associated with this IP address.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the network interface card (NIC) associated with the address, if any.",
	},
	"names": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcNameResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of all names associated with this address.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol (_ip4_ or _ip6_).",
	},
	"range": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of the address (_used_ or _free_).",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for this address in JSON format.",
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
			"  | usage indicator        | description                                                                                                 |\n" +
			"  |------------------------|-------------------------------------------------------------------------------------------------------------|\n" +
			"  | _IPAM_                 |  Address was created by the IPAM component.                                                                 |\n" +
			"  | _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.                                         |\n" +
			"  | _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.   |\n" +
			"  | _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet. |\n" +
			"  | _DHCP_                 |  Address was created by the DHCP component.                                                                 |\n" +
			"  | _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.                                                  |\n" +
			"  | _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.                                              |\n" +
			"  | _DHCP_, _DISABLED_     |  Address is disabled.                                                                                       |\n" +
			"  | _DNS_                  |  Address is used by one or more DNS records.                                                                |\n" +
			"  | _DISCOVERED_           |  Address is discovered by some network discovery probe like Network Insight or NetMRI in NIOS.              |\n" +
			"  <br>",
	},
	"next_available_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier for the address block, subnet or range where the next available address should be generated",
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(path.MatchRoot("address"), path.MatchRoot("next_available_id")),
			stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(range|subnet|address_block)/[0-9a-f-].*$`), "Should be the resource identifier of an address block, range or subnet."),
		},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
	},
}

func (m *IpamsvcAddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Address {
	if m == nil {
		return nil
	}
	to := &ipam.Address{
		Address:   flex.ExpandString(m.Address),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Hwaddr:    flex.ExpandStringPointer(m.Hwaddr),
		Interface: flex.ExpandStringPointer(m.Interface),
		Names:     flex.ExpandFrameworkListNestedBlock(ctx, m.Names, diags, ExpandIpamsvcName),
		Tags:      flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() {
			naipId := flex.ExpandString(m.NextAvailableId) + "/nextavailableip"
			to.Address = naipId
		} else {
			to.Address = flex.ExpandString(m.Address)
		}

		to.Space = flex.ExpandStringPointer(m.Space)
	}
	return to
}

func FlattenIpamsvcAddress(ctx context.Context, from *ipam.Address, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAddressAttrTypes)
	}
	m := IpamsvcAddressModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAddressModel) Flatten(ctx context.Context, from *ipam.Address, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAddressModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpInfo = FlattenIpamsvcDHCPInfo(ctx, from.DhcpInfo, diags)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.DiscoveryAttrs = flex.FlattenFrameworkMapString(ctx, from.DiscoveryAttrs, diags)
	m.DiscoveryMetadata = flex.FlattenFrameworkMapString(ctx, from.DiscoveryMetadata, diags)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Hwaddr = flex.FlattenStringPointer(from.Hwaddr)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.Names = flex.FlattenFrameworkListNestedBlock(ctx, from.Names, IpamsvcNameAttrTypes, diags, FlattenIpamsvcName)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Range = flex.FlattenStringPointer(from.Range)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.State = flex.FlattenStringPointer(from.State)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Usage = flex.FlattenFrameworkListString(ctx, from.Usage, diags)
}
