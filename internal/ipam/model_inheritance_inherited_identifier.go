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

type InheritanceInheritedIdentifierModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.String `tfsdk:"value"`
}

var InheritanceInheritedIdentifierAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.StringType,
}

var InheritanceInheritedIdentifierResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type _Identifier_.`,
	Attributes:          InheritanceInheritedIdentifierResourceSchemaAttributes,
}

var InheritanceInheritedIdentifierResourceSchemaAttributes = map[string]schema.Attribute{
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
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandInheritanceInheritedIdentifier(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedIdentifier {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritanceInheritedIdentifierModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritanceInheritedIdentifierModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedIdentifier {
	if m == nil {
		return nil
	}

	to := &ipam.InheritanceInheritedIdentifier{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  m.Value.ValueStringPointer(),
	}
	return to
}

func flattenInheritanceInheritedIdentifier(ctx context.Context, from *ipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedIdentifierAttrTypes)
	}
	m := InheritanceInheritedIdentifierModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedIdentifierAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedIdentifierModel) flatten(ctx context.Context, from *ipam.InheritanceInheritedIdentifier, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedIdentifierModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = types.StringPointerValue(from.Value)

}
