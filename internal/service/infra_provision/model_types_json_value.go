package infra_provision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type TypesJSONValueModel struct {
	Value types.String `tfsdk:"value"`
}

var TypesJSONValueAttrTypes = map[string]attr.Type{
	"value": types.StringType,
}

var TypesJSONValueResourceSchemaAttributes = map[string]schema.Attribute{
	"value": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandTypesJSONValue(ctx context.Context, o types.Object, diags *diag.Diagnostics) *infra_provision.TypesJSONValue {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TypesJSONValueModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TypesJSONValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *infra_provision.TypesJSONValue {
	if m == nil {
		return nil
	}
	to := &infra_provision.TypesJSONValue{
		Value: m.Value.ValueStringPointer(),
	}
	return to
}

func FlattenTypesJSONValue(ctx context.Context, from *infra_provision.TypesJSONValue, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TypesJSONValueAttrTypes)
	}
	m := TypesJSONValueModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TypesJSONValueAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TypesJSONValueModel) Flatten(ctx context.Context, from *infra_provision.TypesJSONValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TypesJSONValueModel{}
	}
	m.Value = flex.FlattenStringPointer(from.Value)
}
