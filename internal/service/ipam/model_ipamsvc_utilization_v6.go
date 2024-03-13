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

type IpamsvcUtilizationV6Model struct {
	Abandoned types.String `tfsdk:"abandoned"`
	Dynamic   types.String `tfsdk:"dynamic"`
	Static    types.String `tfsdk:"static"`
	Total     types.String `tfsdk:"total"`
	Used      types.String `tfsdk:"used"`
}

var IpamsvcUtilizationV6AttrTypes = map[string]attr.Type{
	"abandoned": types.StringType,
	"dynamic":   types.StringType,
	"static":    types.StringType,
	"total":     types.StringType,
	"used":      types.StringType,
}

var IpamsvcUtilizationV6ResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned": schema.StringAttribute{
		Optional: true,
	},
	"dynamic": schema.StringAttribute{
		Optional: true,
	},
	"static": schema.StringAttribute{
		Optional: true,
	},
	"total": schema.StringAttribute{
		Optional: true,
	},
	"used": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandIpamsvcUtilizationV6(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcUtilizationV6 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcUtilizationV6Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcUtilizationV6Model) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcUtilizationV6 {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcUtilizationV6{
		Abandoned: m.Abandoned.ValueStringPointer(),
		Dynamic:   m.Dynamic.ValueStringPointer(),
		Static:    m.Static.ValueStringPointer(),
		Total:     m.Total.ValueStringPointer(),
		Used:      m.Used.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcUtilizationV6(ctx context.Context, from *ipam.IpamsvcUtilizationV6, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUtilizationV6AttrTypes)
	}
	m := IpamsvcUtilizationV6Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUtilizationV6AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUtilizationV6Model) Flatten(ctx context.Context, from *ipam.IpamsvcUtilizationV6, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUtilizationV6Model{}
	}
	m.Abandoned = flex.FlattenStringPointer(from.Abandoned)
	m.Dynamic = flex.FlattenStringPointer(from.Dynamic)
	m.Static = flex.FlattenStringPointer(from.Static)
	m.Total = flex.FlattenStringPointer(from.Total)
	m.Used = flex.FlattenStringPointer(from.Used)
}
