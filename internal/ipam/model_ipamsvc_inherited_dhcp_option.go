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

type IpamsvcInheritedDHCPOptionModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedDHCPOptionAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionItemAttrTypes},
}

var IpamsvcInheritedDHCPOptionResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type of _OptionItem_.`,
	Attributes:          IpamsvcInheritedDHCPOptionResourceSchemaAttributes,
}

var IpamsvcInheritedDHCPOptionResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _block_: Don&#39;t use the inherited value.  Defaults to _inherit_.`,
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
		Attributes:          IpamsvcInheritedDHCPOptionItemResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedDHCPOption(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOption {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedDHCPOptionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedDHCPOptionModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDHCPOption {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedDHCPOption{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcInheritedDHCPOptionItem(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedDHCPOption(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOption, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDHCPOptionAttrTypes)
	}
	m := IpamsvcInheritedDHCPOptionModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDHCPOptionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDHCPOptionModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedDHCPOption, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDHCPOptionModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcInheritedDHCPOptionItem(ctx, from.Value, diags)

}
