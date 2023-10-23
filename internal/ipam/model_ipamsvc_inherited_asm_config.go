package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcInheritedASMConfigModel struct {
	AsmEnableBlock types.Object `tfsdk:"asm_enable_block"`
	AsmGrowthBlock types.Object `tfsdk:"asm_growth_block"`
	AsmThreshold   types.Object `tfsdk:"asm_threshold"`
	ForecastPeriod types.Object `tfsdk:"forecast_period"`
	History        types.Object `tfsdk:"history"`
	MinTotal       types.Object `tfsdk:"min_total"`
	MinUnused      types.Object `tfsdk:"min_unused"`
}

var IpamsvcInheritedASMConfigAttrTypes = map[string]attr.Type{
	"asm_enable_block": types.ObjectType{AttrTypes: IpamsvcInheritedAsmEnableBlockAttrTypes},
	"asm_growth_block": types.ObjectType{AttrTypes: IpamsvcInheritedAsmGrowthBlockAttrTypes},
	"asm_threshold":    types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"forecast_period":  types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"history":          types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"min_total":        types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
	"min_unused":       types.ObjectType{AttrTypes: InheritanceInheritedUInt32AttrTypes},
}

var IpamsvcInheritedASMConfigResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for the __ASMConfig__ object.`,
	Attributes:          IpamsvcInheritedASMConfigResourceSchemaAttributes,
}

var IpamsvcInheritedASMConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_enable_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedAsmEnableBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"asm_growth_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedAsmGrowthBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"asm_threshold": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"forecast_period": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"history": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"min_total": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"min_unused": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedASMConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedASMConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedASMConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedASMConfigModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedASMConfig {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedASMConfig{
		AsmEnableBlock: expandIpamsvcInheritedAsmEnableBlock(ctx, m.AsmEnableBlock, diags),
		AsmGrowthBlock: expandIpamsvcInheritedAsmGrowthBlock(ctx, m.AsmGrowthBlock, diags),
		AsmThreshold:   expandInheritanceInheritedUInt32(ctx, m.AsmThreshold, diags),
		ForecastPeriod: expandInheritanceInheritedUInt32(ctx, m.ForecastPeriod, diags),
		History:        expandInheritanceInheritedUInt32(ctx, m.History, diags),
		MinTotal:       expandInheritanceInheritedUInt32(ctx, m.MinTotal, diags),
		MinUnused:      expandInheritanceInheritedUInt32(ctx, m.MinUnused, diags),
	}
	return to
}

func flattenIpamsvcInheritedASMConfig(ctx context.Context, from *ipam.IpamsvcInheritedASMConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedASMConfigAttrTypes)
	}
	m := IpamsvcInheritedASMConfigModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedASMConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedASMConfigModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedASMConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedASMConfigModel{}
	}

	m.AsmEnableBlock = flattenIpamsvcInheritedAsmEnableBlock(ctx, from.AsmEnableBlock, diags)
	m.AsmGrowthBlock = flattenIpamsvcInheritedAsmGrowthBlock(ctx, from.AsmGrowthBlock, diags)
	m.AsmThreshold = flattenInheritanceInheritedUInt32(ctx, from.AsmThreshold, diags)
	m.ForecastPeriod = flattenInheritanceInheritedUInt32(ctx, from.ForecastPeriod, diags)
	m.History = flattenInheritanceInheritedUInt32(ctx, from.History, diags)
	m.MinTotal = flattenInheritanceInheritedUInt32(ctx, from.MinTotal, diags)
	m.MinUnused = flattenInheritanceInheritedUInt32(ctx, from.MinUnused, diags)

}
