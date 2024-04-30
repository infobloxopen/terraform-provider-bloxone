package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigHostModel struct {
	AbsoluteName         types.String `tfsdk:"absolute_name"`
	Address              types.String `tfsdk:"address"`
	AnycastAddresses     types.List   `tfsdk:"anycast_addresses"`
	AssociatedServer     types.Object `tfsdk:"associated_server"`
	Comment              types.String `tfsdk:"comment"`
	CurrentVersion       types.String `tfsdk:"current_version"`
	Dfp                  types.Bool   `tfsdk:"dfp"`
	DfpService           types.String `tfsdk:"dfp_service"`
	Id                   types.String `tfsdk:"id"`
	InheritanceSources   types.Object `tfsdk:"inheritance_sources"`
	KerberosKeys         types.List   `tfsdk:"kerberos_keys"`
	Name                 types.String `tfsdk:"name"`
	Ophid                types.String `tfsdk:"ophid"`
	ProtocolAbsoluteName types.String `tfsdk:"protocol_absolute_name"`
	ProviderId           types.String `tfsdk:"provider_id"`
	Server               types.String `tfsdk:"server"`
	SiteId               types.String `tfsdk:"site_id"`
	Tags                 types.Map    `tfsdk:"tags"`
	Type                 types.String `tfsdk:"type"`
}

var ConfigHostAttrTypes = map[string]attr.Type{
	"absolute_name":          types.StringType,
	"address":                types.StringType,
	"anycast_addresses":      types.ListType{ElemType: types.StringType},
	"associated_server":      types.ObjectType{AttrTypes: ConfigHostAssociatedServerAttrTypes},
	"comment":                types.StringType,
	"current_version":        types.StringType,
	"dfp":                    types.BoolType,
	"dfp_service":            types.StringType,
	"id":                     types.StringType,
	"inheritance_sources":    types.ObjectType{AttrTypes: ConfigHostInheritanceAttrTypes},
	"kerberos_keys":          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigKerberosKeyAttrTypes}},
	"name":                   types.StringType,
	"ophid":                  types.StringType,
	"protocol_absolute_name": types.StringType,
	"provider_id":            types.StringType,
	"server":                 types.StringType,
	"site_id":                types.StringType,
	"tags":                   types.MapType{ElemType: types.StringType},
	"type":                   types.StringType,
}

var ConfigHostResourceSchemaAttributes = map[string]schema.Attribute{
	"absolute_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host FQDN.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host's primary IP Address.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          ConfigHostAssociatedServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Host associated server configuration.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host description.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host current version.",
	},
	"dfp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Below _dfp_ field is deprecated and not supported anymore. The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.",
	},
	"dfp_service": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP. Allowed values:  * _unavailable_: BloxOne TD DFP application is not available,  * _enabled_: BloxOne TD DFP application is available and enabled,  * _disabled_: BloxOne TD DFP application is available but disabled.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          ConfigHostInheritanceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Inheritance configuration.",
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigKerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host display name.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "On-Prem Host ID.",
	},
	"protocol_absolute_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host FQDN in punycode.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"site_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host site ID.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Host tagging specifics.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.",
	},
}

func ExpandConfigHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Host {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Host {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Host{
		AssociatedServer:   ExpandConfigHostAssociatedServer(ctx, m.AssociatedServer, diags),
		InheritanceSources: ExpandConfigHostInheritance(ctx, m.InheritanceSources, diags),
		KerberosKeys:       flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandConfigKerberosKey),
		Server:             flex.ExpandStringPointer(m.Server),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenConfigHost(ctx context.Context, from *dnsconfig.Host, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigHostAttrTypes)
	}
	m := ConfigHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigHostModel) Flatten(ctx context.Context, from *dnsconfig.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigHostModel{}
	}
	m.AbsoluteName = flex.FlattenStringPointer(from.AbsoluteName)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenConfigHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.Dfp = types.BoolPointerValue(from.Dfp)
	m.DfpService = flex.FlattenStringPointer(from.DfpService)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenConfigHostInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, ConfigKerberosKeyAttrTypes, diags, FlattenConfigKerberosKey)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProtocolAbsoluteName = flex.FlattenStringPointer(from.ProtocolAbsoluteName)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
