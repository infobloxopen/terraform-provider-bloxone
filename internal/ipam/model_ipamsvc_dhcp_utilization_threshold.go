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

type IpamsvcDHCPUtilizationThresholdModel struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	High    types.Int64 `tfsdk:"high"`
	Low     types.Int64 `tfsdk:"low"`
}

var IpamsvcDHCPUtilizationThresholdAttrTypes = map[string]attr.Type{
	"enabled": types.BoolType,
	"high":    types.Int64Type,
	"low":     types.Int64Type,
}

var IpamsvcDHCPUtilizationThresholdResourceSchema = schema.Schema{
	MarkdownDescription: `A __DHCPUtilizationThreshold__ object represents threshold settings for DHCP utilization.`,
	Attributes:          IpamsvcDHCPUtilizationThresholdResourceSchemaAttributes,
}

var IpamsvcDHCPUtilizationThresholdResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: `Indicates whether the DHCP utilization threshold is enabled or not.`,
	},
	"high": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: `The high threshold value for DHCP utilization in percentage.`,
	},
	"low": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: `The low threshold value for DHCP utilization in percentage.`,
	},
}

func expandIpamsvcDHCPUtilizationThreshold(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDHCPUtilizationThreshold {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDHCPUtilizationThresholdModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDHCPUtilizationThresholdModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDHCPUtilizationThreshold {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDHCPUtilizationThreshold{
		Enabled: m.Enabled.ValueBool(),
		High:    int64(m.High.ValueInt64()),
		Low:     int64(m.Low.ValueInt64()),
	}
	return to
}

func flattenIpamsvcDHCPUtilizationThreshold(ctx context.Context, from *ipam.IpamsvcDHCPUtilizationThreshold, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDHCPUtilizationThresholdAttrTypes)
	}
	m := IpamsvcDHCPUtilizationThresholdModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDHCPUtilizationThresholdAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDHCPUtilizationThresholdModel) flatten(ctx context.Context, from *ipam.IpamsvcDHCPUtilizationThreshold, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDHCPUtilizationThresholdModel{}
	}

	m.Enabled = types.BoolValue(from.Enabled)
	m.High = types.Int64Value(int64(from.High))
	m.Low = types.Int64Value(int64(from.Low))

}
