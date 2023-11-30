package dns_config

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigForwardNSGModel struct {
	Comment            types.String `tfsdk:"comment"`
	ExternalForwarders types.List   `tfsdk:"external_forwarders"`
	ForwardersOnly     types.Bool   `tfsdk:"forwarders_only"`
	Hosts              types.List   `tfsdk:"hosts"`
	Id                 types.String `tfsdk:"id"`
	InternalForwarders types.List   `tfsdk:"internal_forwarders"`
	Name               types.String `tfsdk:"name"`
	Nsgs               types.List   `tfsdk:"nsgs"`
	Tags               types.Map    `tfsdk:"tags"`
}

var ConfigForwardNSGAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"external_forwarders": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"forwarders_only":     types.BoolType,
	"hosts":               types.ListType{ElemType: types.StringType},
	"id":                  types.StringType,
	"internal_forwarders": types.ListType{ElemType: types.StringType},
	"name":                types.StringType,
	"nsgs":                types.ListType{ElemType: types.StringType},
	"tags":                types.MapType{ElemType: types.StringType},
}

var ConfigForwardNSGResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Comment for the object.",
	},
	"external_forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. External DNS servers to forward to. Order is not significant.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.",
	},
	"hosts": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"internal_forwarders": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of the object.",
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tagging specifics.",
	},
}

func ExpandConfigForwardNSG(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigForwardNSG {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigForwardNSGModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigForwardNSGModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigForwardNSG {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigForwardNSG{
		Comment:            flex.ExpandStringPointer(m.Comment),
		ExternalForwarders: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalForwarders, diags, ExpandConfigForwarder),
		ForwardersOnly:     flex.ExpandBoolPointer(m.ForwardersOnly),
		Hosts:              flex.ExpandFrameworkListString(ctx, m.Hosts, diags),
		InternalForwarders: flex.ExpandFrameworkListString(ctx, m.InternalForwarders, diags),
		Name:               flex.ExpandString(m.Name),
		Nsgs:               flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenConfigForwardNSG(ctx context.Context, from *dns_config.ConfigForwardNSG, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigForwardNSGAttrTypes)
	}
	m := ConfigForwardNSGModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigForwardNSGAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigForwardNSGModel) Flatten(ctx context.Context, from *dns_config.ConfigForwardNSG, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigForwardNSGModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExternalForwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalForwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.Hosts = flex.FlattenFrameworkListString(ctx, from.Hosts, diags)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InternalForwarders = flex.FlattenFrameworkListString(ctx, from.InternalForwarders, diags)
	m.Name = flex.FlattenString(from.Name)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
}
