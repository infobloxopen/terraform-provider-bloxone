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

type IpamsvcExclusionRangeModel struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

var IpamsvcExclusionRangeAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"end":     types.StringType,
	"start":   types.StringType,
}

var IpamsvcExclusionRangeResourceSchema = schema.Schema{
	MarkdownDescription: `The __ExclusionRange__ object represents an exclusion range inside a DHCP range.`,
	Attributes:          IpamsvcExclusionRangeResourceSchemaAttributes,
}

var IpamsvcExclusionRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The end address of the exclusion range.`,
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The start address of the exclusion range.`,
	},
}

func expandIpamsvcExclusionRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcExclusionRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcExclusionRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcExclusionRangeModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcExclusionRange {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcExclusionRange{
		Comment: m.Comment.ValueStringPointer(),
		End:     m.End.ValueString(),
		Start:   m.Start.ValueString(),
	}
	return to
}

func flattenIpamsvcExclusionRange(ctx context.Context, from *ipam.IpamsvcExclusionRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcExclusionRangeAttrTypes)
	}
	m := IpamsvcExclusionRangeModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcExclusionRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcExclusionRangeModel) flatten(ctx context.Context, from *ipam.IpamsvcExclusionRange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcExclusionRangeModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.End = types.StringValue(from.End)
	m.Start = types.StringValue(from.Start)

}
