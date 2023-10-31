package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type IpamsvcAsmGrowthBlockModel struct {
	GrowthFactor types.Int64  `tfsdk:"growth_factor"`
	GrowthType   types.String `tfsdk:"growth_type"`
}

var IpamsvcAsmGrowthBlockAttrTypes = map[string]attr.Type{
	"growth_factor": types.Int64Type,
	"growth_type":   types.StringType,
}

var IpamsvcAsmGrowthBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"growth_factor": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Either the number or percentage of addresses to grow by.`,
	},
	"growth_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The type of factor to use: _percent_ or _count_.`,
	},
}

func ExpandIpamsvcAsmGrowthBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAsmGrowthBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcAsmGrowthBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcAsmGrowthBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAsmGrowthBlock {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcAsmGrowthBlock{
		GrowthFactor: utils.Ptr(int64(m.GrowthFactor.ValueInt64())),
		GrowthType:   m.GrowthType.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcAsmGrowthBlock(ctx context.Context, from *ipam.IpamsvcAsmGrowthBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAsmGrowthBlockAttrTypes)
	}
	m := IpamsvcAsmGrowthBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAsmGrowthBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAsmGrowthBlockModel) Flatten(ctx context.Context, from *ipam.IpamsvcAsmGrowthBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAsmGrowthBlockModel{}
	}
	m.GrowthFactor = flex.FlattenInt64(int64(*from.GrowthFactor))
	m.GrowthType = flex.FlattenStringPointer(from.GrowthType)

}
