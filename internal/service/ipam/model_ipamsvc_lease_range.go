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

type IpamsvcLeaseRangeModel struct {
	Id types.String `tfsdk:"id"`
}

var IpamsvcLeaseRangeAttrTypes = map[string]attr.Type{
	"id": types.StringType,
}

var IpamsvcLeaseRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandIpamsvcLeaseRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcLeaseRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcLeaseRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcLeaseRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcLeaseRange {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcLeaseRange{}
	return to
}

func FlattenIpamsvcLeaseRange(ctx context.Context, from *ipam.IpamsvcLeaseRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcLeaseRangeAttrTypes)
	}
	m := IpamsvcLeaseRangeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcLeaseRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcLeaseRangeModel) Flatten(ctx context.Context, from *ipam.IpamsvcLeaseRange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcLeaseRangeModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
}
