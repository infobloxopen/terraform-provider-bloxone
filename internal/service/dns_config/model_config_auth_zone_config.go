package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigAuthZoneConfigModel struct {
	ExternalPrimaries   types.List `tfsdk:"external_primaries"`
	ExternalSecondaries types.List `tfsdk:"external_secondaries"`
	InternalSecondaries types.List `tfsdk:"internal_secondaries"`
	Nsgs                types.List `tfsdk:"nsgs"`
}

var ConfigAuthZoneConfigAttrTypes = map[string]attr.Type{
	"external_primaries":   types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalPrimaryAttrTypes}},
	"external_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalSecondaryAttrTypes}},
	"internal_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigInternalSecondaryAttrTypes}},
	"nsgs":                 types.ListType{ElemType: types.StringType},
}

var ConfigAuthZoneConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. DNS primaries external to BloxOne DDI. Order is not significant.`,
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `DNS secondaries external to BloxOne DDI. Order is not significant.`,
	},
	"internal_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigInternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.`,
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandConfigAuthZoneConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.AuthZoneConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigAuthZoneConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigAuthZoneConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.AuthZoneConfig {
	if m == nil {
		return nil
	}
	to := &dnsconfig.AuthZoneConfig{
		ExternalPrimaries:   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandConfigExternalPrimary),
		ExternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandConfigExternalSecondary),
		InternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandConfigInternalSecondary),
		Nsgs:                flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
	}
	return to
}

func FlattenConfigAuthZoneConfig(ctx context.Context, from *dnsconfig.AuthZoneConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigAuthZoneConfigAttrTypes)
	}
	m := ConfigAuthZoneConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigAuthZoneConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigAuthZoneConfigModel) Flatten(ctx context.Context, from *dnsconfig.AuthZoneConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigAuthZoneConfigModel{}
	}
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ConfigExternalPrimaryAttrTypes, diags, FlattenConfigExternalPrimary)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ConfigExternalSecondaryAttrTypes, diags, FlattenConfigExternalSecondary)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, ConfigInternalSecondaryAttrTypes, diags, FlattenConfigInternalSecondary)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
}
