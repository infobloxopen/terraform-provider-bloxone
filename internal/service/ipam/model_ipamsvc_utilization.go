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
		MarkdownDescription: `The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.`,
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

func ExpandIpamsvcUtilization(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Utilization {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcUtilizationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcUtilizationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Utilization {
	if m == nil {
		return nil
	}
	to := &ipam.Utilization{}
	return to
}

func FlattenIpamsvcUtilization(ctx context.Context, from *ipam.Utilization, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUtilizationAttrTypes)
	}
	m := IpamsvcUtilizationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUtilizationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUtilizationModel) Flatten(ctx context.Context, from *ipam.Utilization, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUtilizationModel{}
	}
	m.AbandonUtilization = flex.FlattenInt64(int64(*from.AbandonUtilization))
	m.Abandoned = flex.FlattenStringPointer(from.Abandoned)
	m.Dynamic = flex.FlattenStringPointer(from.Dynamic)
	m.Free = flex.FlattenStringPointer(from.Free)
	m.Static = flex.FlattenStringPointer(from.Static)
	m.Total = flex.FlattenStringPointer(from.Total)
	m.Used = flex.FlattenStringPointer(from.Used)
	m.Utilization = flex.FlattenInt64(int64(*from.Utilization))
}
