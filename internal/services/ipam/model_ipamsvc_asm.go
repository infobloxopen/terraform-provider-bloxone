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
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type IpamsvcASMModel struct {
	BackEnd              types.String `tfsdk:"back_end"`
	BackStart            types.String `tfsdk:"back_start"`
	BothEnd              types.String `tfsdk:"both_end"`
	BothStart            types.String `tfsdk:"both_start"`
	FrontEnd             types.String `tfsdk:"front_end"`
	FrontStart           types.String `tfsdk:"front_start"`
	Growth               types.Int64  `tfsdk:"growth"`
	Id                   types.String `tfsdk:"id"`
	Lookahead            types.Int64  `tfsdk:"lookahead"`
	RangeEnd             types.String `tfsdk:"range_end"`
	RangeId              types.String `tfsdk:"range_id"`
	RangeStart           types.String `tfsdk:"range_start"`
	SubnetAddress        types.String `tfsdk:"subnet_address"`
	SubnetCidr           types.Int64  `tfsdk:"subnet_cidr"`
	SubnetDirection      types.String `tfsdk:"subnet_direction"`
	SubnetId             types.String `tfsdk:"subnet_id"`
	SubnetRange          types.String `tfsdk:"subnet_range"`
	SubnetRangeEnd       types.String `tfsdk:"subnet_range_end"`
	SubnetRangeStart     types.String `tfsdk:"subnet_range_start"`
	Suppress             types.String `tfsdk:"suppress"`
	SuppressTime         types.Int64  `tfsdk:"suppress_time"`
	ThresholdUtilization types.Int64  `tfsdk:"threshold_utilization"`
	Update               types.String `tfsdk:"update"`
	Utilization          types.Int64  `tfsdk:"utilization"`
}

var IpamsvcASMAttrTypes = map[string]attr.Type{
	"back_end":              types.StringType,
	"back_start":            types.StringType,
	"both_end":              types.StringType,
	"both_start":            types.StringType,
	"front_end":             types.StringType,
	"front_start":           types.StringType,
	"growth":                types.Int64Type,
	"id":                    types.StringType,
	"lookahead":             types.Int64Type,
	"range_end":             types.StringType,
	"range_id":              types.StringType,
	"range_start":           types.StringType,
	"subnet_address":        types.StringType,
	"subnet_cidr":           types.Int64Type,
	"subnet_direction":      types.StringType,
	"subnet_id":             types.StringType,
	"subnet_range":          types.StringType,
	"subnet_range_end":      types.StringType,
	"subnet_range_start":    types.StringType,
	"suppress":              types.StringType,
	"suppress_time":         types.Int64Type,
	"threshold_utilization": types.Int64Type,
	"update":                types.StringType,
	"utilization":           types.Int64Type,
}

var IpamsvcASMResourceSchemaAttributes = map[string]schema.Attribute{
	"back_end": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The end IP address when adding to the back.`,
	},
	"back_start": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The start IP address when adding to the back.`,
	},
	"both_end": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The end IP address when adding to the back.`,
	},
	"both_start": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The start IP address when adding to both front and back.`,
	},
	"front_end": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The end IP address when adding to the front.`,
	},
	"front_start": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The start IP address when adding to the front.`,
	},
	"growth": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `Calculated number of addresses to grow range; its value is determined by asm_config growth factor, type, and min_unused after the expansion.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"lookahead": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `Either the value from the ASM configuration or -1 if the estimate is that utilization will not exceed the threshold.`,
	},
	"range_end": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The end IP address of the range.`,
	},
	"range_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"range_start": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The start IP address of the range.`,
	},
	"subnet_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The suggested subnet expansion.`,
	},
	"subnet_cidr": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The CIDR of the subnet.`,
	},
	"subnet_direction": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Indicates where the subnet may expand. As the subnet can only be expanded by one bit at a time, it can only grow in one of the two directions. It is set to _none_ if the subnet can't be expanded.  Valid values are: * _front_ * _back_ * _none_`,
	},
	"subnet_id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"subnet_range": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"subnet_range_end": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The suggested new range end in expanded subnet.`,
	},
	"subnet_range_start": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The suggested new range start in expanded subnet.`,
	},
	"suppress": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if future notifications for this subnet should be suppressed.  Valid values are: * _no_ * _time_ * _permanent_  If set to _permanent_ notifications are suppressed until the administrator modifies the configuration for the subnet. If set to _time_ notifications are suppressed until the specified time. Defaults to _no_.`,
	},
	"suppress_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The time duration in days to suppress the notifications for this subnet.`,
	},
	"threshold_utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The utilization threshold as per ASM configuration.`,
	},
	"update": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The object to update.  Valid values are: * _range_ * _subnet_ * _none_`,
	},
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The utilization of DHCP addresses in the subnet.`,
	},
}

func ExpandIpamsvcASM(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcASM {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcASMModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcASMModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcASM {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcASM{
		RangeEnd:     m.RangeEnd.ValueStringPointer(),
		RangeId:      m.RangeId.ValueStringPointer(),
		RangeStart:   m.RangeStart.ValueStringPointer(),
		SubnetId:     m.SubnetId.ValueString(),
		SubnetRange:  m.SubnetRange.ValueStringPointer(),
		Suppress:     m.Suppress.ValueStringPointer(),
		SuppressTime: utils.Ptr(int64(m.SuppressTime.ValueInt64())),
		Update:       m.Update.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcASM(ctx context.Context, from *ipam.IpamsvcASM, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcASMAttrTypes)
	}
	m := IpamsvcASMModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcASMAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcASMModel) Flatten(ctx context.Context, from *ipam.IpamsvcASM, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcASMModel{}
	}
	m.BackEnd = flex.FlattenStringPointer(from.BackEnd)
	m.BackStart = flex.FlattenStringPointer(from.BackStart)
	m.BothEnd = flex.FlattenStringPointer(from.BothEnd)
	m.BothStart = flex.FlattenStringPointer(from.BothStart)
	m.FrontEnd = flex.FlattenStringPointer(from.FrontEnd)
	m.FrontStart = flex.FlattenStringPointer(from.FrontStart)
	m.Growth = flex.FlattenInt64(int64(*from.Growth))
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Lookahead = flex.FlattenInt64(int64(*from.Lookahead))
	m.RangeEnd = flex.FlattenStringPointer(from.RangeEnd)
	m.RangeId = flex.FlattenStringPointer(from.RangeId)
	m.RangeStart = flex.FlattenStringPointer(from.RangeStart)
	m.SubnetAddress = flex.FlattenStringPointer(from.SubnetAddress)
	m.SubnetCidr = flex.FlattenInt64(int64(*from.SubnetCidr))
	m.SubnetDirection = flex.FlattenStringPointer(from.SubnetDirection)
	m.SubnetId = flex.FlattenString(from.SubnetId)
	m.SubnetRange = flex.FlattenStringPointer(from.SubnetRange)
	m.SubnetRangeEnd = flex.FlattenStringPointer(from.SubnetRangeEnd)
	m.SubnetRangeStart = flex.FlattenStringPointer(from.SubnetRangeStart)
	m.Suppress = flex.FlattenStringPointer(from.Suppress)
	m.SuppressTime = flex.FlattenInt64(int64(*from.SuppressTime))
	m.ThresholdUtilization = flex.FlattenInt64(int64(*from.ThresholdUtilization))
	m.Update = flex.FlattenStringPointer(from.Update)
	m.Utilization = flex.FlattenInt64(int64(*from.Utilization))

}
