package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigAuthNSGModel struct {
	Comment             types.String `tfsdk:"comment"`
	ExternalPrimaries   types.List   `tfsdk:"external_primaries"`
	ExternalSecondaries types.List   `tfsdk:"external_secondaries"`
	GridPrimaries       types.List   `tfsdk:"grid_primaries"`
	GridSecondaries     types.List   `tfsdk:"grid_secondaries"`
	Id                  types.String `tfsdk:"id"`
	InternalSecondaries types.List   `tfsdk:"internal_secondaries"`
	Name                types.String `tfsdk:"name"`
	Nameservers         types.List   `tfsdk:"nameservers"`
	Nsgs                types.List   `tfsdk:"nsgs"`
	Tags                types.Map    `tfsdk:"tags"`
	TagsAll             types.Map    `tfsdk:"tags_all"`
	Version             types.String `tfsdk:"version"`
}

var ConfigAuthNSGAttrTypes = map[string]attr.Type{
	"comment":              types.StringType,
	"external_primaries":   types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalPrimaryAttrTypes}},
	"external_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalSecondaryAttrTypes}},
	"grid_primaries":       types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigMemberServerAttrTypes}},
	"grid_secondaries":     types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigMemberServerAttrTypes}},
	"id":                   types.StringType,
	"internal_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigInternalSecondaryAttrTypes}},
	"name":                 types.StringType,
	"nameservers":          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigNameserverAttrTypes}},
	"nsgs":                 types.ListType{ElemType: types.StringType},
	"tags":                 types.MapType{ElemType: types.StringType},
	"tags_all":             types.MapType{ElemType: types.StringType},
	"version":              types.StringType,
}

var ConfigAuthNSGResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Optional. Comment for the object.",
	},
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. DNS primaries external to Universal DDI. Order is not significant.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "DNS secondaries external to Universal DDI. Order is not significant.",
	},
	"grid_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigMemberServerResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. The list of the NIOS Grid Primaries assigned to an AuthNSG, only applicable for the NIOS.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigMemberServerResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. The list of the NIOS Grid Secondaries assigned to an AuthNSG, only applicable for the NIOS.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"internal_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigInternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Universal DDI hosts acting as internal secondaries. Order is not significant.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of the object.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"nameservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigNameserverResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. A list of DNS Nameservers of various roles.",
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Tagging specifics includes default tags.",
	},
	"version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read Only. Version indicates the version of the Authoritative DNS Server Group in context of DNS NSGs and nameservers that are used. Possible values:\n- _v1_: The Authoritative DNS Server Group uses original NSG model\n- _v2_: The Authoritative DNS Server Group uses new \"Unified Nameservers\" model",
	},
}

func ExpandConfigAuthNSG(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.AuthNSG {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigAuthNSGModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigAuthNSGModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.AuthNSG {
	if m == nil {
		return nil
	}
	to := &dnsconfig.AuthNSG{
		Comment:             flex.ExpandStringPointer(m.Comment),
		ExternalPrimaries:   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandConfigExternalPrimary),
		ExternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandConfigExternalSecondary),
		GridPrimaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimaries, diags, ExpandConfigMemberServer),
		GridSecondaries:     flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandConfigMemberServer),
		InternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandConfigInternalSecondary),
		Name:                flex.ExpandString(m.Name),
		Nameservers:         flex.ExpandFrameworkListNestedBlock(ctx, m.Nameservers, diags, ExpandConfigNameserver),
		Nsgs:                flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenConfigAuthNSG(ctx context.Context, from *dnsconfig.AuthNSG, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigAuthNSGAttrTypes)
	}
	m := ConfigAuthNSGModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, ConfigAuthNSGAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigAuthNSGModel) Flatten(ctx context.Context, from *dnsconfig.AuthNSG, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigAuthNSGModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ConfigExternalPrimaryAttrTypes, diags, FlattenConfigExternalPrimary)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ConfigExternalSecondaryAttrTypes, diags, FlattenConfigExternalSecondary)
	m.GridPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimaries, ConfigMemberServerAttrTypes, diags, FlattenConfigMemberServer)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ConfigMemberServerAttrTypes, diags, FlattenConfigMemberServer)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, ConfigInternalSecondaryAttrTypes, diags, FlattenConfigInternalSecondary)
	m.Name = flex.FlattenString(from.Name)
	m.Nameservers = flex.FlattenFrameworkListNestedBlock(ctx, from.Nameservers, ConfigNameserverAttrTypes, diags, FlattenConfigNameserver)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Version = flex.FlattenStringPointer(from.Version)
}
