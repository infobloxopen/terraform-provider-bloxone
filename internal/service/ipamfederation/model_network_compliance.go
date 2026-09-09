package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type NetworkComplianceModel struct {
	DefaultNetmaskLength types.Int64 `tfsdk:"default_netmask_length"`
	MaximumNetmaskLength types.Int64 `tfsdk:"maximum_netmask_length"`
	MinimumNetmaskLength types.Int64 `tfsdk:"minimum_netmask_length"`
}

var NetworkComplianceAttrTypes = map[string]attr.Type{
	"default_netmask_length": types.Int64Type,
	"maximum_netmask_length": types.Int64Type,
	"minimum_netmask_length": types.Int64Type,
}

var NetworkComplianceResourceSchemaAttributes = map[string]schema.Attribute{
	"default_netmask_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The default netmask length used when allocating child blocks or pools.",
	},
	"maximum_netmask_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The maximum netmask length for allocating child blocks or pools.",
	},
	"minimum_netmask_length": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minimum netmask length for allocating child blocks or pools.",
	},
}

func ExpandNetworkCompliance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.NetworkCompliance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkComplianceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkComplianceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.NetworkCompliance {
	if m == nil {
		return nil
	}
	to := &ipamfederation.NetworkCompliance{
		DefaultNetmaskLength: flex.ExpandInt64Pointer(m.DefaultNetmaskLength),
		MaximumNetmaskLength: flex.ExpandInt64Pointer(m.MaximumNetmaskLength),
		MinimumNetmaskLength: flex.ExpandInt64Pointer(m.MinimumNetmaskLength),
	}
	return to
}

func FlattenNetworkCompliance(ctx context.Context, from *ipamfederation.NetworkCompliance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkComplianceAttrTypes)
	}
	m := NetworkComplianceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkComplianceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkComplianceModel) Flatten(ctx context.Context, from *ipamfederation.NetworkCompliance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkComplianceModel{}
	}
	m.DefaultNetmaskLength = flex.FlattenInt64Pointer(from.DefaultNetmaskLength)
	m.MaximumNetmaskLength = flex.FlattenInt64Pointer(from.MaximumNetmaskLength)
	m.MinimumNetmaskLength = flex.FlattenInt64Pointer(from.MinimumNetmaskLength)
}
