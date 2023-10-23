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

type IpamsvcInheritedAsmEnableBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedAsmEnableBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcAsmEnableBlockAttrTypes},
}

var IpamsvcInheritedAsmEnableBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance block for ASM fields: _enable_, _enable_notification_, _reenable_date_.`,
	Attributes:          IpamsvcInheritedAsmEnableBlockResourceSchemaAttributes,
}

var IpamsvcInheritedAsmEnableBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          IpamsvcAsmEnableBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedAsmEnableBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedAsmEnableBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedAsmEnableBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedAsmEnableBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedAsmEnableBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedAsmEnableBlock{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcAsmEnableBlock(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedAsmEnableBlock(ctx context.Context, from *ipam.IpamsvcInheritedAsmEnableBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedAsmEnableBlockAttrTypes)
	}
	m := IpamsvcInheritedAsmEnableBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedAsmEnableBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedAsmEnableBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedAsmEnableBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedAsmEnableBlockModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcAsmEnableBlock(ctx, from.Value, diags)

}
