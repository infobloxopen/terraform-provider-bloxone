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

type InheritanceInheritedStringModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.String `tfsdk:"value"`
}

var InheritanceInheritedStringAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.StringType,
}

var InheritanceInheritedStringResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type _String_.`,
	Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
}

var InheritanceInheritedStringResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func expandInheritanceInheritedString(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedString {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritanceInheritedStringModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritanceInheritedStringModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedString {
	if m == nil {
		return nil
	}

	to := &ipam.InheritanceInheritedString{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
	}
	return to
}

func flattenInheritanceInheritedString(ctx context.Context, from *ipam.InheritanceInheritedString, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedStringAttrTypes)
	}
	m := InheritanceInheritedStringModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedStringAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedStringModel) flatten(ctx context.Context, from *ipam.InheritanceInheritedString, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedStringModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = types.StringPointerValue(from.Value)

}
