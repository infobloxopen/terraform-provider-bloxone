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

type UtilizationV6Model struct {
	Total types.String `tfsdk:"total"`
	Used  types.String `tfsdk:"used"`
}

var UtilizationV6AttrTypes = map[string]attr.Type{
	"total": types.StringType,
	"used":  types.StringType,
}

var UtilizationV6ResourceSchemaAttributes = map[string]schema.Attribute{
	"total": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Total IPv6 addresses.",
	},
	"used": schema.StringAttribute{
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
		Total: flex.ExpandStringPointer(m.Total),
		Used:  flex.ExpandStringPointer(m.Used),
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
	if m == nil {
		*m = UtilizationV6Model{}
	}
	m.Total = flex.FlattenStringPointer(from.Total)
	m.Used = flex.FlattenStringPointer(from.Used)
}
