package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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

var IpamsvcExclusionRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the exclusion range. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The end address of the exclusion range.",
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The start address of the exclusion range.",
	},
}

func ExpandIpamsvcExclusionRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.ExclusionRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcExclusionRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcExclusionRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.ExclusionRange {
	if m == nil {
		return nil
	}
	to := &ipam.ExclusionRange{
		Comment: flex.ExpandStringPointer(m.Comment),
		End:     flex.ExpandString(m.End),
		Start:   flex.ExpandString(m.Start),
	}
	return to
}

func FlattenIpamsvcExclusionRange(ctx context.Context, from *ipam.ExclusionRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcExclusionRangeAttrTypes)
	}
	m := IpamsvcExclusionRangeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcExclusionRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcExclusionRangeModel) Flatten(ctx context.Context, from *ipam.ExclusionRange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcExclusionRangeModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.End = flex.FlattenString(from.End)
	m.Start = flex.FlattenString(from.Start)
}
