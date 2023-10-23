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

type InheritanceInheritedBoolModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Bool   `tfsdk:"value"`
}

var InheritanceInheritedBoolAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.BoolType,
}

var InheritanceInheritedBoolResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type _Bool_.`,
	Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
}

var InheritanceInheritedBoolResourceSchemaAttributes = map[string]schema.Attribute{
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
	"value": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func expandInheritanceInheritedBool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedBool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritanceInheritedBoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritanceInheritedBoolModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedBool {
	if m == nil {
		return nil
	}

	to := &ipam.InheritanceInheritedBool{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
	}
	return to
}

func flattenInheritanceInheritedBool(ctx context.Context, from *ipam.InheritanceInheritedBool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedBoolAttrTypes)
	}
	m := InheritanceInheritedBoolModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedBoolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedBoolModel) flatten(ctx context.Context, from *ipam.InheritanceInheritedBool, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedBoolModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = types.BoolPointerValue(from.Value)

}
