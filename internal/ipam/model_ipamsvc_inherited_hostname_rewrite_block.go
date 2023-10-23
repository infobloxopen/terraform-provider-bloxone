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

type IpamsvcInheritedHostnameRewriteBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedHostnameRewriteBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcHostnameRewriteBlockAttrTypes},
}

var IpamsvcInheritedHostnameRewriteBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance block for fields: _hostname_rewrite_enabled_, _hostname_rewrite_regex_, _hostname_rewrite_char_.`,
	Attributes:          IpamsvcInheritedHostnameRewriteBlockResourceSchemaAttributes,
}

var IpamsvcInheritedHostnameRewriteBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          IpamsvcHostnameRewriteBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedHostnameRewriteBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedHostnameRewriteBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedHostnameRewriteBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedHostnameRewriteBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedHostnameRewriteBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedHostnameRewriteBlock{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcHostnameRewriteBlock(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedHostnameRewriteBlock(ctx context.Context, from *ipam.IpamsvcInheritedHostnameRewriteBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedHostnameRewriteBlockAttrTypes)
	}
	m := IpamsvcInheritedHostnameRewriteBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedHostnameRewriteBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedHostnameRewriteBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedHostnameRewriteBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedHostnameRewriteBlockModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcHostnameRewriteBlock(ctx, from.Value, diags)

}
