package dns_config

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigForwardZoneModel struct {
	Comment            types.String      `tfsdk:"comment"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at"`
	Disabled           types.Bool        `tfsdk:"disabled"`
	ExternalForwarders types.List        `tfsdk:"external_forwarders"`
	ForwardOnly        types.Bool        `tfsdk:"forward_only"`
	Fqdn               types.String      `tfsdk:"fqdn"`
	Hosts              types.List        `tfsdk:"hosts"`
	Id                 types.String      `tfsdk:"id"`
	InternalForwarders types.List        `tfsdk:"internal_forwarders"`
	MappedSubnet       types.String      `tfsdk:"mapped_subnet"`
	Mapping            types.String      `tfsdk:"mapping"`
	Nsgs               types.List        `tfsdk:"nsgs"`
	Parent             types.String      `tfsdk:"parent"`
	ProtocolFqdn       types.String      `tfsdk:"protocol_fqdn"`
	Tags               types.Map         `tfsdk:"tags"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at"`
	View               types.String      `tfsdk:"view"`
	Warnings           types.List        `tfsdk:"warnings"`
}

var ConfigForwardZoneAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"created_at":          timetypes.RFC3339Type{},
	"disabled":            types.BoolType,
	"external_forwarders": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"forward_only":        types.BoolType,
	"fqdn":                types.StringType,
	"hosts":               types.ListType{ElemType: types.StringType},
	"id":                  types.StringType,
	"internal_forwarders": types.ListType{ElemType: types.StringType},
	"mapped_subnet":       types.StringType,
	"mapping":             types.StringType,
	"nsgs":                types.ListType{ElemType: types.StringType},
	"parent":              types.StringType,
	"protocol_fqdn":       types.StringType,
	"tags":                types.MapType{ElemType: types.StringType},
	"updated_at":          timetypes.RFC3339Type{},
	"view":                types.StringType,
	"warnings":            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigWarningAttrTypes}},
}

var ConfigForwardZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Comment for zone configuration.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The timestamp when the object has been created.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
	},
	"external_forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. External DNS servers to forward to. Order is not significant.",
	},
	"forward_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.",
	},
	"fqdn": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"hosts": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"internal_forwarders": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"mapped_subnet": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.",
	},
	"mapping": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read-only. Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to _forward_.",
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Zone FQDN in punycode.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tagging specifics.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"warnings": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigWarningResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of a forward zone warnings.",
	},
}

func (m *ConfigForwardZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns_config.ConfigForwardZone {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigForwardZone{
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		ExternalForwarders: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalForwarders, diags, ExpandConfigForwarder),
		ForwardOnly:        flex.ExpandBoolPointer(m.ForwardOnly),
		Hosts:              flex.ExpandFrameworkListString(ctx, m.Hosts, diags),
		InternalForwarders: flex.ExpandFrameworkListString(ctx, m.InternalForwarders, diags),
		Nsgs:               flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Parent:             flex.ExpandStringPointer(m.Parent),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func FlattenConfigForwardZone(ctx context.Context, from *dns_config.ConfigForwardZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigForwardZoneAttrTypes)
	}
	m := ConfigForwardZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigForwardZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigForwardZoneModel) Flatten(ctx context.Context, from *dns_config.ConfigForwardZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigForwardZoneModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.ExternalForwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalForwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.ForwardOnly = types.BoolPointerValue(from.ForwardOnly)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Hosts = flex.FlattenFrameworkListString(ctx, from.Hosts, diags)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InternalForwarders = flex.FlattenFrameworkListString(ctx, from.InternalForwarders, diags)
	m.MappedSubnet = flex.FlattenStringPointer(from.MappedSubnet)
	m.Mapping = flex.FlattenStringPointer(from.Mapping)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.View = flex.FlattenStringPointer(from.View)
	m.Warnings = flex.FlattenFrameworkListNestedBlock(ctx, from.Warnings, ConfigWarningAttrTypes, diags, FlattenConfigWarning)
}
