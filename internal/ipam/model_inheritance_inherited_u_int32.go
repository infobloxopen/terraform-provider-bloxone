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

type InheritanceInheritedUInt32Model struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Int64  `tfsdk:"value"`
}

var InheritanceInheritedUInt32AttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.Int64Type,
}

var InheritanceInheritedUInt32ResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type _UInt32_.`,
	Attributes:          InheritanceInheritedUInt32ResourceSchemaAttributes,
}

var InheritanceInheritedUInt32ResourceSchemaAttributes = map[string]schema.Attribute{
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
	"value": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func expandInheritanceInheritedUInt32(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedUInt32 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritanceInheritedUInt32Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritanceInheritedUInt32Model) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedUInt32 {
	if m == nil {
		return nil
	}

	to := &ipam.InheritanceInheritedUInt32{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
	}
	return to
}

func flattenInheritanceInheritedUInt32(ctx context.Context, from *ipam.InheritanceInheritedUInt32, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedUInt32AttrTypes)
	}
	m := InheritanceInheritedUInt32Model{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedUInt32AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedUInt32Model) flatten(ctx context.Context, from *ipam.InheritanceInheritedUInt32, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedUInt32Model{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = types.Int64Value(int64(*from.Value))

}
