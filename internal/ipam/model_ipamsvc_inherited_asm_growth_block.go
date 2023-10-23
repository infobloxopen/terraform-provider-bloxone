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

type IpamsvcInheritedAsmGrowthBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedAsmGrowthBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcAsmGrowthBlockAttrTypes},
}

var IpamsvcInheritedAsmGrowthBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance block for ASM fields: _growth_factor_ and _growth_type_.`,
	Attributes:          IpamsvcInheritedAsmGrowthBlockResourceSchemaAttributes,
}

var IpamsvcInheritedAsmGrowthBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.SingleNestedAttribute{
		Attributes:          IpamsvcAsmGrowthBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedAsmGrowthBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedAsmGrowthBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedAsmGrowthBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedAsmGrowthBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedAsmGrowthBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedAsmGrowthBlock{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcAsmGrowthBlock(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedAsmGrowthBlock(ctx context.Context, from *ipam.IpamsvcInheritedAsmGrowthBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedAsmGrowthBlockAttrTypes)
	}
	m := IpamsvcInheritedAsmGrowthBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedAsmGrowthBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedAsmGrowthBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedAsmGrowthBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedAsmGrowthBlockModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcAsmGrowthBlock(ctx, from.Value, diags)

}
