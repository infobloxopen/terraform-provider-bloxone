package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

type UtilizationV6Model struct {
	Total types.Object `tfsdk:"total"`
	Used  types.Object `tfsdk:"used"`
}

var UtilizationV6AttrTypes = map[string]attr.Type{
	"total": types.ObjectType{AttrTypes: Integer128AttrTypes},
	"used":  types.ObjectType{AttrTypes: Integer128AttrTypes},
}

var UtilizationV6ResourceSchemaAttributes = map[string]schema.Attribute{
	"total": schema.SingleNestedAttribute{
		Attributes:          Integer128ResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Total IPv6 addresses.",
	},
	"used": schema.SingleNestedAttribute{
		Attributes:          Integer128ResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Used IPv6 addresses.",
	},
}

func ExpandUtilizationV6(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.UtilizationV6 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UtilizationV6Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *UtilizationV6Model) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.UtilizationV6 {
	if m == nil {
		return nil
	}
	return &ipamfederation.UtilizationV6{
		Total: ExpandInteger128(ctx, m.Total, diags),
		Used:  ExpandInteger128(ctx, m.Used, diags),
	}
}

func FlattenUtilizationV6(ctx context.Context, from *ipamfederation.UtilizationV6, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UtilizationV6AttrTypes)
	}
	m := UtilizationV6Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UtilizationV6AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *UtilizationV6Model) Flatten(ctx context.Context, from *ipamfederation.UtilizationV6, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	m.Total = FlattenInteger128(ctx, from.Total, diags)
	m.Used = FlattenInteger128(ctx, from.Used, diags)
}
