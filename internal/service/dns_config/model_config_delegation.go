package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

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
	"view":               types.StringType,
}

var ConfigDelegationResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Comment for zone delegation.",
	},
	"delegation_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigDelegationServerResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "Required. DNS zone delegation servers. Order is not significant.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records.",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"parent": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Delegation FQDN in punycode.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tagging specifics.",
	},
	"view": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandConfigDelegation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigDelegation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDelegationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigDelegation {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigDelegation{
		Comment:           flex.ExpandStringPointer(m.Comment),
		DelegationServers: flex.ExpandFrameworkListNestedBlock(ctx, m.DelegationServers, diags, ExpandConfigDelegationServer),
		Disabled:          flex.ExpandBoolPointer(m.Disabled),
		Fqdn:              flex.ExpandStringPointer(m.Fqdn),
		Parent:            flex.ExpandStringPointer(m.Parent),
		Tags:              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		View:              flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenConfigDelegation(ctx context.Context, from *dns_config.ConfigDelegation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDelegationAttrTypes)
	}
	m := ConfigDelegationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDelegationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDelegationModel) Flatten(ctx context.Context, from *dns_config.ConfigDelegation, diags *diag.Diagnostics) {
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
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.View = flex.FlattenStringPointer(from.View)
}