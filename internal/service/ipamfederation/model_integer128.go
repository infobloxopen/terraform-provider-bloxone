package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type Integer128Model struct {
	RawValue types.String `tfsdk:"raw_value"`
}

var Integer128AttrTypes = map[string]attr.Type{
	"raw_value": types.StringType,
}

var Integer128ResourceSchemaAttributes = map[string]schema.Attribute{
	"raw_value": schema.StringAttribute{
		Computed: true,
	},
}

func ExpandInteger128(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.Integer128 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Integer128Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Integer128Model) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.Integer128 {
	if m == nil {
		return nil
	}
	to := &ipamfederation.Integer128{
		RawValue: flex.ExpandStringPointer(m.RawValue),
	}
	return to
}

func FlattenInteger128(ctx context.Context, from *ipamfederation.Integer128, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Integer128AttrTypes)
	}
	m := Integer128Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Integer128AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Integer128Model) Flatten(ctx context.Context, from *ipamfederation.Integer128, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Integer128Model{}
	}
	m.RawValue = flex.FlattenStringPointer(from.RawValue)
}
