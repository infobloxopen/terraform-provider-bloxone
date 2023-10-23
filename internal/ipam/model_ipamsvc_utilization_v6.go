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

var IpamsvcUtilizationV6ResourceSchema = schema.Schema{
	MarkdownDescription: `The __UtilizationV6__ object represents IPV6 address usage statistics for an object.`,
	Attributes:          IpamsvcUtilizationV6ResourceSchemaAttributes,
}

var IpamsvcUtilizationV6ResourceSchemaAttributes = map[string]schema.Attribute{
	"abandoned": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: ``,
	},
	"dynamic": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: ``,
	},
	"static": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: ``,
	},
	"total": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: ``,
	},
	"used": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcUtilizationV6(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcUtilizationV6 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcUtilizationV6Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcUtilizationV6Model) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcUtilizationV6 {
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

func flattenIpamsvcUtilizationV6(ctx context.Context, from *ipam.IpamsvcUtilizationV6, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUtilizationV6AttrTypes)
	}
	m := IpamsvcUtilizationV6Model{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUtilizationV6AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUtilizationV6Model) flatten(ctx context.Context, from *ipam.IpamsvcUtilizationV6, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUtilizationV6Model{}
	}

	m.Abandoned = types.StringPointerValue(from.Abandoned)
	m.Dynamic = types.StringPointerValue(from.Dynamic)
	m.Static = types.StringPointerValue(from.Static)
	m.Total = types.StringPointerValue(from.Total)
	m.Used = types.StringPointerValue(from.Used)

}
