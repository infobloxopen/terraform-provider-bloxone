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

type ConfigForwardZoneConfigModel struct {
	ExternalForwarders types.List `tfsdk:"external_forwarders"`
	Hosts              types.List `tfsdk:"hosts"`
	InternalForwarders types.List `tfsdk:"internal_forwarders"`
	Nsgs               types.List `tfsdk:"nsgs"`
}

var ConfigForwardZoneConfigAttrTypes = map[string]attr.Type{
	"external_forwarders": types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"hosts":               types.ListType{ElemType: types.StringType},
	"internal_forwarders": types.ListType{ElemType: types.StringType},
	"nsgs":                types.ListType{ElemType: types.StringType},
}

var ConfigForwardZoneConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"external_forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. External DNS servers to forward to. Order is not significant.`,
	},
	"hosts": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"internal_forwarders": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandConfigForwardZoneConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigForwardZoneConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigForwardZoneConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigForwardZoneConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigForwardZoneConfig {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigForwardZoneConfig{
		ExternalForwarders: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalForwarders, diags, ExpandConfigForwarder),
		Hosts:              flex.ExpandFrameworkListString(ctx, m.Hosts, diags),
		InternalForwarders: flex.ExpandFrameworkListString(ctx, m.InternalForwarders, diags),
		Nsgs:               flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
	}
	return to
}

func FlattenConfigForwardZoneConfig(ctx context.Context, from *dns_config.ConfigForwardZoneConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigForwardZoneConfigAttrTypes)
	}
	m := ConfigForwardZoneConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigForwardZoneConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigForwardZoneConfigModel) Flatten(ctx context.Context, from *dns_config.ConfigForwardZoneConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigForwardZoneConfigModel{}
	}
	m.ExternalForwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalForwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.Hosts = flex.FlattenFrameworkListString(ctx, from.Hosts, diags)
	m.InternalForwarders = flex.FlattenFrameworkListString(ctx, from.InternalForwarders, diags)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
}
