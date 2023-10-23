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

type IpamsvcInheritedDDNSBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedDDNSBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcDDNSBlockAttrTypes},
}

var IpamsvcInheritedDDNSBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration block for dynamic DNS.`,
	Attributes:          IpamsvcInheritedDDNSBlockResourceSchemaAttributes,
}

var IpamsvcInheritedDDNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          IpamsvcDDNSBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedDDNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDDNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedDDNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedDDNSBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDDNSBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedDDNSBlock{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcDDNSBlock(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedDDNSBlock(ctx context.Context, from *ipam.IpamsvcInheritedDDNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDDNSBlockAttrTypes)
	}
	m := IpamsvcInheritedDDNSBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDDNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDDNSBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedDDNSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDDNSBlockModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcDDNSBlock(ctx, from.Value, diags)

}
