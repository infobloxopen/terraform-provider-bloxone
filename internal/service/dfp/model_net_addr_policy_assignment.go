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

type NetAddrPolicyAssignmentModel struct {
	AddrNet  types.String `tfsdk:"addr_net"`
	PolicyId types.Int64  `tfsdk:"policy_id"`
}

var NetAddrPolicyAssignmentAttrTypes = map[string]attr.Type{
	"addr_net":  types.StringType,
	"policy_id": types.Int64Type,
}

var NetAddrPolicyAssignmentResourceSchemaAttributes = map[string]schema.Attribute{
	"addr_net": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "network address in IPv4 CIDR (address/bitmask length) string format",
	},
	"policy_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Identifier of the security policy associated with this address block",
	},
}

func ExpandNetAddrPolicyAssignment(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.NetAddrPolicyAssignment {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetAddrPolicyAssignmentModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetAddrPolicyAssignmentModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.NetAddrPolicyAssignment {
	if m == nil {
		return nil
	}
	to := &dfp.NetAddrPolicyAssignment{
		AddrNet:  flex.ExpandStringPointer(m.AddrNet),
		PolicyId: flex.ExpandInt32Pointer(m.PolicyId),
	}
	return to
}

func FlattenNetAddrPolicyAssignment(ctx context.Context, from *dfp.NetAddrPolicyAssignment, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetAddrPolicyAssignmentAttrTypes)
	}
	m := NetAddrPolicyAssignmentModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetAddrPolicyAssignmentAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetAddrPolicyAssignmentModel) Flatten(ctx context.Context, from *dfp.NetAddrPolicyAssignment, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetAddrPolicyAssignmentModel{}
	}
	m.AddrNet = flex.FlattenStringPointer(from.AddrNet)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
}
