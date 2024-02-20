package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcdfpNetAddrPolicyAssignmentModel struct {
	AddrNet  types.String `tfsdk:"addr_net"`
	PolicyId types.Int64  `tfsdk:"policy_id"`
}

var AtcdfpNetAddrPolicyAssignmentAttrTypes = map[string]attr.Type{
	"addr_net":  types.StringType,
	"policy_id": types.Int64Type,
}

var AtcdfpNetAddrPolicyAssignmentResourceSchemaAttributes = map[string]schema.Attribute{
	"addr_net": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "network address in IPv4 CIDR (address/bitmask length) string format",
	},
	"policy_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Identifier of the security policy associated with this address block",
	},
}

func ExpandAtcdfpNetAddrPolicyAssignment(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpNetAddrPolicyAssignment {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpNetAddrPolicyAssignmentModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpNetAddrPolicyAssignmentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpNetAddrPolicyAssignment {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpNetAddrPolicyAssignment{
		AddrNet:  flex.ExpandStringPointer(m.AddrNet),
		PolicyId: flex.ExpandInt32Pointer(m.PolicyId),
	}
	return to
}

func FlattenAtcdfpNetAddrPolicyAssignment(ctx context.Context, from *dfp.AtcdfpNetAddrPolicyAssignment, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpNetAddrPolicyAssignmentAttrTypes)
	}
	m := AtcdfpNetAddrPolicyAssignmentModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpNetAddrPolicyAssignmentAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpNetAddrPolicyAssignmentModel) Flatten(ctx context.Context, from *dfp.AtcdfpNetAddrPolicyAssignment, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpNetAddrPolicyAssignmentModel{}
	}
	m.AddrNet = flex.FlattenStringPointer(from.AddrNet)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
}
