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

var InheritanceAssignedHostResourceSchemaAttributes = map[string]schema.Attribute{
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the host referred to by _ophid_.",
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The on-prem host ID.",
	},
}

func ExpandInheritanceAssignedHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritanceAssignedHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InheritanceAssignedHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InheritanceAssignedHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritanceAssignedHost {
	if m == nil {
		return nil
	}
	to := &ipam.InheritanceAssignedHost{
		Host: flex.ExpandStringPointer(m.Host),
	}
	return to
}

func FlattenInheritanceAssignedHost(ctx context.Context, from *ipam.InheritanceAssignedHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritanceAssignedHostAttrTypes)
	}
	m := InheritanceAssignedHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritanceAssignedHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritanceAssignedHostModel) Flatten(ctx context.Context, from *ipam.InheritanceAssignedHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritanceAssignedHostModel{}
	}
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
}
