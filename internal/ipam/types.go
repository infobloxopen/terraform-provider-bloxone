package ipam

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func ExpandTime(_ context.Context, dt timetypes.RFC3339, diags *diag.Diagnostics) time.Time {
	t, d := dt.ValueRFC3339Time()
	diags.Append(d...)
	return t
}

func ExpandFrameworkMapString(ctx context.Context, tfMap types.Map, diags *diag.Diagnostics) map[string]interface{} {
	if tfMap.IsNull() || tfMap.IsUnknown() {
		return nil
	}
	elements := make(map[string]interface{}, len(tfMap.Elements()))
	diags.Append(tfMap.ElementsAs(ctx, &elements, false)...)
	return elements
}

func FlattenFrameworkMapString(ctx context.Context, m map[string]interface{}, diags *diag.Diagnostics) types.Map {
	if len(m) == 0 {
		return types.MapNull(types.StringType)
	}
	tfMap, d := types.MapValueFrom(ctx, types.StringType, m)
	diags.Append(d...)
	return tfMap
}

func ExpandFrameworkListString(ctx context.Context, tfList types.List, diags *diag.Diagnostics) []string {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []string
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

func FlattenFrameworkListString(ctx context.Context, l []string, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.StringType)
	}
	tfList, d := types.ListValueFrom(ctx, types.StringType, l)
	diags.Append(d...)
	return tfList
}

type FrameworkElementFlattenFunc[T any, U any] func(context.Context, T, *diag.Diagnostics) U

func FlattenFrameworkListNestedBlock[T any, U any](ctx context.Context, data []T, attrTypes map[string]attr.Type, diags *diag.Diagnostics, f FrameworkElementExpanderFunc[*T, U]) types.List {
	tfData := ApplyToAll(data, func(t T) U {
		return f(ctx, &t, diags)
	})

	tfList, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: attrTypes}, tfData)
	diags.Append(d...)
	return tfList
}

type FrameworkElementExpanderFunc[T any, U any] func(context.Context, T, *diag.Diagnostics) U

func ExpandFrameworkListNestedBlock[T any, U any](ctx context.Context, tfList types.List, diags *diag.Diagnostics, f FrameworkElementExpanderFunc[T, *U]) []U {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}

	var data []T

	diags.Append(tfList.ElementsAs(ctx, &data, false)...)

	return ApplyToAll(data, func(t T) U {
		return *f(ctx, t, diags)
	})

}

// ApplyToAll returns a new slice containing the results of applying the function `f` to each element of the original slice `s`.
func ApplyToAll[T, U any](s []T, f func(T) U) []U {
	v := make([]U, len(s))

	for i, e := range s {
		v[i] = f(e)
	}

	return v
}

func ptr[T any](t T) *T {
	return &t
}
