package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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
	TagsAll          types.Map    `tfsdk:"tags_all"`
	Type             types.String `tfsdk:"type"`
}

type IpamsvcHostModelWithRetryAndTimeouts struct {
	Address          types.String   `tfsdk:"address"`
	AnycastAddresses types.List     `tfsdk:"anycast_addresses"`
	AssociatedServer types.Object   `tfsdk:"associated_server"`
	Comment          types.String   `tfsdk:"comment"`
	CurrentVersion   types.String   `tfsdk:"current_version"`
	Id               types.String   `tfsdk:"id"`
	IpSpace          types.String   `tfsdk:"ip_space"`
	Name             types.String   `tfsdk:"name"`
	Ophid            types.String   `tfsdk:"ophid"`
	ProviderId       types.String   `tfsdk:"provider_id"`
	Server           types.String   `tfsdk:"server"`
	Tags             types.Map      `tfsdk:"tags"`
	Type             types.String   `tfsdk:"type"`
	RetryIfNotFound  types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
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
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
	"type":              types.StringType,
}

var IpamsvcHostResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary IP address of the on-prem host.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          IpamsvcHostAssociatedServerResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The DHCP Config Profile for the on-prem host.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description for the on-prem host.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current dhcp application version of the host.",
	},
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The display name of the on-prem host.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The on-prem host ID.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the on-prem host in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the on-prem host in JSON format including default tags.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.",
	},
}

var IpamsvcHostResourceSchemaAttributesWithRetryAndTimeouts = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary IP address of the on-prem host.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          IpamsvcHostAssociatedServerResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The DHCP Config Profile for the on-prem host.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description for the on-prem host.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current dhcp application version of the host.",
	},
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The display name of the on-prem host.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The on-prem host ID.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the on-prem host in JSON format.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.",
	},
	"retry_if_not_found": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to `true`, the resource will retry until a matching host is found, or until the Create Timeout expires.",
	},
}

func ExpandIpamsvcHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Host {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Host {
	if m == nil {
		return nil
	}
	to := &ipam.Host{
		Server: flex.ExpandStringPointer(m.Server),
	}
	return to
}

func (m *IpamsvcHostModelWithRetryAndTimeouts) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Host {
	if m == nil {
		return nil
	}
	to := &ipam.Host{
		Server: flex.ExpandStringPointer(m.Server),
	}
	return to
}

func FlattenIpamsvcHostDataSource(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHostAttrTypes)
	}
	m := IpamsvcHostModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, IpamsvcHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHostModel) Flatten(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenIpamsvcHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}

func (m *IpamsvcHostModelWithRetryAndTimeouts) Flatten(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostModelWithRetryAndTimeouts{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenIpamsvcHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
