package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigDelegationModel struct {
	Comment           types.String `tfsdk:"comment"`
	DelegationServers types.List   `tfsdk:"delegation_servers"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Fqdn              types.String `tfsdk:"fqdn"`
	Id                types.String `tfsdk:"id"`
	Parent            types.String `tfsdk:"parent"`
	ProtocolFqdn      types.String `tfsdk:"protocol_fqdn"`
	Tags              types.Map    `tfsdk:"tags"`
	TagsAll           types.Map    `tfsdk:"tags_all"`
	View              types.String `tfsdk:"view"`
}

var ConfigDelegationAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"delegation_servers": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigDelegationServerAttrTypes}},
	"disabled":           types.BoolType,
	"fqdn":               types.StringType,
	"id":                 types.StringType,
	"parent":             types.StringType,
	"protocol_fqdn":      types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"view":               types.StringType,
}

var ConfigDelegationResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Optional. Comment for zone delegation.",
	},
	"delegation_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigDelegationServerResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Required. DNS zone delegation servers. Order is not significant.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Delegation FQDN in punycode.",
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
	"view": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandConfigDelegation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Delegation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDelegationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, true)
}

func (m *ConfigDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dnsconfig.Delegation {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Delegation{
		Comment:           flex.ExpandStringPointer(m.Comment),
		DelegationServers: flex.ExpandFrameworkListNestedBlock(ctx, m.DelegationServers, diags, ExpandConfigDelegationServer),
		Disabled:          flex.ExpandBoolPointer(m.Disabled),
		Parent:            flex.ExpandStringPointer(m.Parent),
		Tags:              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func DataSourceFlattenConfigDelegation(ctx context.Context, from *dnsconfig.Delegation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDelegationAttrTypes)
	}
	m := ConfigDelegationModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, ConfigDelegationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDelegationModel) Flatten(ctx context.Context, from *dnsconfig.Delegation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDelegationModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DelegationServers = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegationServers, ConfigDelegationServerAttrTypes, diags, FlattenConfigDelegationServer)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.View = flex.FlattenStringPointer(from.View)
}
