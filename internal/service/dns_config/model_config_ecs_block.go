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

type ConfigECSBlockModel struct {
	EcsEnabled    types.Bool  `tfsdk:"ecs_enabled"`
	EcsForwarding types.Bool  `tfsdk:"ecs_forwarding"`
	EcsPrefixV4   types.Int64 `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6   types.Int64 `tfsdk:"ecs_prefix_v6"`
	EcsZones      types.List  `tfsdk:"ecs_zones"`
}

var ConfigECSBlockAttrTypes = map[string]attr.Type{
	"ecs_enabled":    types.BoolType,
	"ecs_forwarding": types.BoolType,
	"ecs_prefix_v4":  types.Int64Type,
	"ecs_prefix_v6":  types.Int64Type,
	"ecs_zones":      types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigECSZoneAttrTypes}},
}

var ConfigECSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _ecs_enabled_ field.`,
	},
	"ecs_forwarding": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _ecs_forwarding_ field.`,
	},
	"ecs_prefix_v4": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _ecs_prefix_v4_ field.`,
	},
	"ecs_prefix_v6": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _ecs_prefix_v6_ field.`,
	},
	"ecs_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigECSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _ecs_zones_ field.`,
	},
}

func ExpandConfigECSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigECSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigECSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigECSBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigECSBlock {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigECSBlock{
		EcsEnabled:    flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding: flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:   flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:   flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:      flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandConfigECSZone),
	}
	return to
}

func FlattenConfigECSBlock(ctx context.Context, from *dns_config.ConfigECSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigECSBlockAttrTypes)
	}
	m := ConfigECSBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigECSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigECSBlockModel) Flatten(ctx context.Context, from *dns_config.ConfigECSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigECSBlockModel{}
	}
	m.EcsEnabled = types.BoolPointerValue(from.EcsEnabled)
	m.EcsForwarding = types.BoolPointerValue(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64(int64(*from.EcsPrefixV4))
	m.EcsPrefixV6 = flex.FlattenInt64(int64(*from.EcsPrefixV6))
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ConfigECSZoneAttrTypes, diags, FlattenConfigECSZone)
}
