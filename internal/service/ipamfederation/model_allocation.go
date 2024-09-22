package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AllocationModel struct {
	Allocated   types.Int64 `tfsdk:"allocated"`
	Delegated   types.Int64 `tfsdk:"delegated"`
	Overlapping types.Int64 `tfsdk:"overlapping"`
	Reserved    types.Int64 `tfsdk:"reserved"`
}

var AllocationAttrTypes = map[string]attr.Type{
	"allocated":   types.Int64Type,
	"delegated":   types.Int64Type,
	"overlapping": types.Int64Type,
	"reserved":    types.Int64Type,
}

var AllocationResourceSchemaAttributes = map[string]schema.Attribute{
	"allocated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space allocated.",
	},
	"delegated": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space delegated.",
	},
	"overlapping": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space in overlapping blocks.",
	},
	"reserved": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Percent of total space reserved.",
	},
}

func ExpandAllocation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.Allocation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AllocationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AllocationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.Allocation {
	if m == nil {
		return nil
	}
	to := &ipamfederation.Allocation{}
	return to
}

func FlattenAllocation(ctx context.Context, from *ipamfederation.Allocation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AllocationAttrTypes)
	}
	m := AllocationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AllocationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AllocationModel) Flatten(ctx context.Context, from *ipamfederation.Allocation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AllocationModel{}
	}
	m.Allocated = flex.FlattenInt64Pointer(from.Allocated)
	m.Delegated = flex.FlattenInt64Pointer(from.Delegated)
	m.Overlapping = flex.FlattenInt64Pointer(from.Overlapping)
	m.Reserved = flex.FlattenInt64Pointer(from.Reserved)
}
