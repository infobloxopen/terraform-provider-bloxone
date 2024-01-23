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

type ConfigCustomRootNSBlockModel struct {
	CustomRootNs        types.List `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled types.Bool `tfsdk:"custom_root_ns_enabled"`
}

var ConfigCustomRootNSBlockAttrTypes = map[string]attr.Type{
	"custom_root_ns":         types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigRootNSAttrTypes}},
	"custom_root_ns_enabled": types.BoolType,
}

var ConfigCustomRootNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"custom_root_ns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigRootNSResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _custom_root_ns_ field.`,
	},
	"custom_root_ns_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _custom_root_ns_enabled_ field.`,
	},
}

func ExpandConfigCustomRootNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigCustomRootNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigCustomRootNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigCustomRootNSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigCustomRootNSBlock {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigCustomRootNSBlock{
		CustomRootNs:        flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandConfigRootNS),
		CustomRootNsEnabled: flex.ExpandBoolPointer(m.CustomRootNsEnabled),
	}
	return to
}

func FlattenConfigCustomRootNSBlock(ctx context.Context, from *dns_config.ConfigCustomRootNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigCustomRootNSBlockAttrTypes)
	}
	m := ConfigCustomRootNSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigCustomRootNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigCustomRootNSBlockModel) Flatten(ctx context.Context, from *dns_config.ConfigCustomRootNSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigCustomRootNSBlockModel{}
	}
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, ConfigRootNSAttrTypes, diags, FlattenConfigRootNS)
	m.CustomRootNsEnabled = types.BoolPointerValue(from.CustomRootNsEnabled)
}
