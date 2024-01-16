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
)

type InheritanceInheritedFloatModel struct {
	Action      types.String  `tfsdk:"action"`
	DisplayName types.String  `tfsdk:"display_name"`
	Source      types.String  `tfsdk:"source"`
	Value       types.Float64 `tfsdk:"value"`
}

var InheritanceInheritedFloatAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.Float64Type,
}

var InheritanceInheritedFloatResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.Float64Attribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func ExpandInheritanceInheritedFloat(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceInheritedFloat {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceInheritedFloatModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritanceInheritedFloatModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceInheritedFloat {
	if m == nil {
		return nil
	}
	to := &ipam.InheritanceInheritedFloat{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
	}
	return to
}

func FlattenInheritanceInheritedFloat(ctx context.Context, from *ipam.InheritanceInheritedFloat, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceInheritedFloatAttrTypes)
	}
	m := InheritanceInheritedFloatModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceInheritedFloatAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceInheritedFloatModel) Flatten(ctx context.Context, from *ipam.InheritanceInheritedFloat, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceInheritedFloatModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFloat64(float64(*from.Value))
}
