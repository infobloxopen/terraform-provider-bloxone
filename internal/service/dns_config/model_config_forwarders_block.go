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

type ConfigForwardersBlockModel struct {
	Forwarders                                  types.List `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool `tfsdk:"forwarders_only"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
}

var ConfigForwardersBlockAttrTypes = map[string]attr.Type{
	"forwarders":      types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigForwarderAttrTypes}},
	"forwarders_only": types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
}

var ConfigForwardersBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigForwarderResourceSchemaAttributes(false),
		},
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _forwarders_ field from.`,
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _forwarders_only_ field.`,
	},
	"use_root_forwarders_for_local_resolution_with_b1td": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _use_root_forwarders_for_local_resolution_with_b1td_ field.`,
	},
}

func ExpandConfigForwardersBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigForwardersBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigForwardersBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigForwardersBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigForwardersBlock {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigForwardersBlock{
		Forwarders:     flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandConfigForwarder),
		ForwardersOnly: flex.ExpandBoolPointer(m.ForwardersOnly),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
	}
	return to
}

func FlattenConfigForwardersBlock(ctx context.Context, from *dns_config.ConfigForwardersBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigForwardersBlockAttrTypes)
	}
	m := ConfigForwardersBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigForwardersBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigForwardersBlockModel) Flatten(ctx context.Context, from *dns_config.ConfigForwardersBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigForwardersBlockModel{}
	}
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ConfigForwarderAttrTypes, diags, FlattenConfigForwarder)
	m.ForwardersOnly = types.BoolPointerValue(from.ForwardersOnly)
	m.UseRootForwardersForLocalResolutionWithB1td = types.BoolPointerValue(from.UseRootForwardersForLocalResolutionWithB1td)
}
