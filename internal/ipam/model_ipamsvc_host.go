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

type IpamsvcHostModel struct {
	Address          types.String `tfsdk:"address"`
	AnycastAddresses types.List   `tfsdk:"anycast_addresses"`
	AssociatedServer types.Object `tfsdk:"associated_server"`
	Comment          types.String `tfsdk:"comment"`
	CurrentVersion   types.String `tfsdk:"current_version"`
	Id               types.String `tfsdk:"id"`
	IpSpace          types.String `tfsdk:"ip_space"`
	Name             types.String `tfsdk:"name"`
	Ophid            types.String `tfsdk:"ophid"`
	ProviderId       types.String `tfsdk:"provider_id"`
	Server           types.String `tfsdk:"server"`
	Tags             types.Map    `tfsdk:"tags"`
	Type             types.String `tfsdk:"type"`
}

var IpamsvcHostAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"anycast_addresses": types.ListType{ElemType: types.StringType},
	"associated_server": types.ObjectType{AttrTypes: IpamsvcHostAssociatedServerAttrTypes},
	"comment":           types.StringType,
	"current_version":   types.StringType,
	"id":                types.StringType,
	"ip_space":          types.StringType,
	"name":              types.StringType,
	"ophid":             types.StringType,
	"provider_id":       types.StringType,
	"server":            types.StringType,
	"tags":              types.MapType{},
	"type":              types.StringType,
}

var IpamsvcHostResourceSchema = schema.Schema{
	MarkdownDescription: `A DHCP __Host__ (_dhcp/host_) object associates a DHCP Config Profile with an on-prem host.`,
	Attributes:          IpamsvcHostResourceSchemaAttributes,
}

var IpamsvcHostResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The primary IP address of the on-prem host.`,
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: `Anycast address configured to the host. Order is not significant.`,
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          IpamsvcHostAssociatedServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The description for the on-prem host.`,
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Current dhcp application version of the host.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The display name of the on-prem host.`,
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The on-prem host ID.`,
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `External provider identifier.`,
	},
	"server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags of the on-prem host in JSON format.`,
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.`,
	},
}

func expandIpamsvcHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHostModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHost {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHost{
		AssociatedServer: expandIpamsvcHostAssociatedServer(ctx, m.AssociatedServer, diags),
		IpSpace:          m.IpSpace.ValueStringPointer(),
		Server:           m.Server.ValueStringPointer(),
		Tags:             ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcHost(ctx context.Context, from *ipam.IpamsvcHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHostAttrTypes)
	}
	m := IpamsvcHostModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHostModel) flatten(ctx context.Context, from *ipam.IpamsvcHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostModel{}
	}

	m.Address = types.StringPointerValue(from.Address)
	m.AnycastAddresses = FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = flattenIpamsvcHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = types.StringPointerValue(from.Comment)
	m.CurrentVersion = types.StringPointerValue(from.CurrentVersion)
	m.Id = types.StringPointerValue(from.Id)
	m.IpSpace = types.StringPointerValue(from.IpSpace)
	m.Name = types.StringPointerValue(from.Name)
	m.Ophid = types.StringPointerValue(from.Ophid)
	m.ProviderId = types.StringPointerValue(from.ProviderId)
	m.Server = types.StringPointerValue(from.Server)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = types.StringPointerValue(from.Type)

}
