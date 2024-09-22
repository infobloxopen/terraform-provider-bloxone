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

var IpamsvcInheritedASMConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_enable_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedAsmEnableBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The block of ASM fields: _enable_, _enable_notification_, _reenable_date_.",
	},
	"asm_growth_block": schema.SingleNestedAttribute{
		Attributes:          IpamsvcInheritedAsmGrowthBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The block of ASM fields: _growth_factor_, _growth_type_.",
	},
	"asm_threshold": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "ASM shows the number of addresses forecast to be used _forecast_period_ days in the future, if it is greater than _asm_threshold_percent_ * _dhcp_total_ (see _dhcp_utilization_) then the subnet is flagged.",
	},
	"forecast_period": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The forecast period in days.",
	},
	"history": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The minimum amount of history needed before ASM can run on this subnet.",
	},
	"min_total": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The minimum size of range needed for ASM to run on this subnet.",
	},
	"min_unused": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The minimum percentage of addresses that must be available outside of the DHCP ranges and fixed addresses when making a suggested change.",
	},
}

func ExpandIpamsvcInheritedASMConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedASMConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcInheritedASMConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcInheritedASMConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedASMConfig {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedASMConfig{
		AsmEnableBlock: ExpandIpamsvcInheritedAsmEnableBlock(ctx, m.AsmEnableBlock, diags),
		AsmGrowthBlock: ExpandIpamsvcInheritedAsmGrowthBlock(ctx, m.AsmGrowthBlock, diags),
		AsmThreshold:   ExpandInheritanceInheritedUInt32(ctx, m.AsmThreshold, diags),
		ForecastPeriod: ExpandInheritanceInheritedUInt32(ctx, m.ForecastPeriod, diags),
		History:        ExpandInheritanceInheritedUInt32(ctx, m.History, diags),
		MinTotal:       ExpandInheritanceInheritedUInt32(ctx, m.MinTotal, diags),
		MinUnused:      ExpandInheritanceInheritedUInt32(ctx, m.MinUnused, diags),
	}
	return to
}

func FlattenIpamsvcInheritedASMConfig(ctx context.Context, from *ipam.InheritedASMConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedASMConfigAttrTypes)
	}
	m := IpamsvcInheritedASMConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedASMConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedASMConfigModel) Flatten(ctx context.Context, from *ipam.InheritedASMConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedASMConfigModel{}
	}
	m.AsmEnableBlock = FlattenIpamsvcInheritedAsmEnableBlock(ctx, from.AsmEnableBlock, diags)
	m.AsmGrowthBlock = FlattenIpamsvcInheritedAsmGrowthBlock(ctx, from.AsmGrowthBlock, diags)
	m.AsmThreshold = FlattenInheritanceInheritedUInt32(ctx, from.AsmThreshold, diags)
	m.ForecastPeriod = FlattenInheritanceInheritedUInt32(ctx, from.ForecastPeriod, diags)
	m.History = FlattenInheritanceInheritedUInt32(ctx, from.History, diags)
	m.MinTotal = FlattenInheritanceInheritedUInt32(ctx, from.MinTotal, diags)
	m.MinUnused = FlattenInheritanceInheritedUInt32(ctx, from.MinUnused, diags)
}
