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

type InheritanceAssignedHostModel struct {
	DisplayName types.String `tfsdk:"display_name"`
	Host        types.String `tfsdk:"host"`
	Ophid       types.String `tfsdk:"ophid"`
}

var InheritanceAssignedHostAttrTypes = map[string]attr.Type{
	"display_name": types.StringType,
	"host":         types.StringType,
	"ophid":        types.StringType,
}

var InheritanceAssignedHostResourceSchema = schema.Schema{
	MarkdownDescription: `_ddi/assigned_host_ represents a BloxOne DDI host assigned to one of the following:  * Subnet (_ipam/subnet_)  * Range (_ipam/range_)  * Fixed Address (_dhcp/fixed_address_)  * Authoritative Zone (_dns/auth_zone_)`,
	Attributes:          InheritanceAssignedHostResourceSchemaAttributes,
}

var InheritanceAssignedHostResourceSchemaAttributes = map[string]schema.Attribute{
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the host referred to by _ophid_.`,
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The on-prem host ID.`,
	},
}

func expandInheritanceAssignedHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceAssignedHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritanceAssignedHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritanceAssignedHostModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceAssignedHost {
	if m == nil {
		return nil
	}

	to := &ipam.InheritanceAssignedHost{
		Host: m.Host.ValueStringPointer(),
	}
	return to
}

func flattenInheritanceAssignedHost(ctx context.Context, from *ipam.InheritanceAssignedHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceAssignedHostAttrTypes)
	}
	m := InheritanceAssignedHostModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceAssignedHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceAssignedHostModel) flatten(ctx context.Context, from *ipam.InheritanceAssignedHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceAssignedHostModel{}
	}

	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Host = types.StringPointerValue(from.Host)
	m.Ophid = types.StringPointerValue(from.Ophid)

}
