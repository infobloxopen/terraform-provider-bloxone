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

type IpamsvcUtilizationModel struct {
	AbandonUtilization types.Int64  `tfsdk:"abandon_utilization"`
	Abandoned          types.String `tfsdk:"abandoned"`
	Dynamic            types.String `tfsdk:"dynamic"`
	Free               types.String `tfsdk:"free"`
	Static             types.String `tfsdk:"static"`
	Total              types.String `tfsdk:"total"`
	Used               types.String `tfsdk:"used"`
	Utilization        types.Int64  `tfsdk:"utilization"`
}

var IpamsvcUtilizationAttrTypes = map[string]attr.Type{
	"abandon_utilization": types.Int64Type,
	"abandoned":           types.StringType,
	"dynamic":             types.StringType,
	"free":                types.StringType,
	"static":              types.StringType,
	"total":               types.StringType,
	"used":                types.StringType,
	"utilization":         types.Int64Type,
}

var IpamsvcUtilizationResourceSchema = schema.Schema{
	MarkdownDescription: `The __Utilization__ object represents IP address usage statistics for an object.`,
	Attributes:          IpamsvcUtilizationResourceSchemaAttributes,
}

var IpamsvcUtilizationResourceSchemaAttributes = map[string]schema.Attribute{
	"abandon_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object.`,
	},
	"abandoned": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).`,
	},
	"dynamic": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.`,
	},
	"free": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of IP addresses available in the scope of the object.`,
	},
	"static": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ &#x3D; _used_ - _dynamic_.`,
	},
	"total": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The total number of IP addresses available in the scope of the object.`,
	},
	"used": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The number of IP addresses used in the scope of the object.`,
	},
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The percentage of used IP addresses relative to the total IP addresses available in the scope of the object.`,
	},
}

func expandIpamsvcUtilization(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcUtilization {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcUtilizationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcUtilizationModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcUtilization {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcUtilization{}
	return to
}

func flattenIpamsvcUtilization(ctx context.Context, from *ipam.IpamsvcUtilization, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUtilizationAttrTypes)
	}
	m := IpamsvcUtilizationModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUtilizationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUtilizationModel) flatten(ctx context.Context, from *ipam.IpamsvcUtilization, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUtilizationModel{}
	}

	m.AbandonUtilization = types.Int64Value(int64(*from.AbandonUtilization))
	m.Abandoned = types.StringPointerValue(from.Abandoned)
	m.Dynamic = types.StringPointerValue(from.Dynamic)
	m.Free = types.StringPointerValue(from.Free)
	m.Static = types.StringPointerValue(from.Static)
	m.Total = types.StringPointerValue(from.Total)
	m.Used = types.StringPointerValue(from.Used)
	m.Utilization = types.Int64Value(int64(*from.Utilization))

}
