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

type IpamsvcLeaseRangeModel struct {
	Id types.String `tfsdk:"id"`
}

var IpamsvcLeaseRangeAttrTypes = map[string]attr.Type{
	"id": types.StringType,
}

var IpamsvcLeaseRangeResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcLeaseRangeResourceSchemaAttributes,
}

var IpamsvcLeaseRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcLeaseRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcLeaseRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcLeaseRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcLeaseRangeModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcLeaseRange {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcLeaseRange{}
	return to
}

func flattenIpamsvcLeaseRange(ctx context.Context, from *ipam.IpamsvcLeaseRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcLeaseRangeAttrTypes)
	}
	m := IpamsvcLeaseRangeModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcLeaseRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcLeaseRangeModel) flatten(ctx context.Context, from *ipam.IpamsvcLeaseRange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcLeaseRangeModel{}
	}

	m.Id = types.StringPointerValue(from.Id)

}
