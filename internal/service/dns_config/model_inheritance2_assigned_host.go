package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type Inheritance2AssignedHostModel struct {
	DisplayName types.String `tfsdk:"display_name"`
	Host        types.String `tfsdk:"host"`
	Ophid       types.String `tfsdk:"ophid"`
}

var Inheritance2AssignedHostAttrTypes = map[string]attr.Type{
	"display_name": types.StringType,
	"host":         types.StringType,
	"ophid":        types.StringType,
}

var Inheritance2AssignedHostResourceSchemaAttributes = map[string]schema.Attribute{
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

func ExpandInheritance2AssignedHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Inheritance2AssignedHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2AssignedHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Inheritance2AssignedHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Inheritance2AssignedHost {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Inheritance2AssignedHost{
		Host: flex.ExpandStringPointer(m.Host),
	}
	return to
}

func FlattenInheritance2AssignedHost(ctx context.Context, from *dnsconfig.Inheritance2AssignedHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2AssignedHostAttrTypes)
	}
	m := Inheritance2AssignedHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2AssignedHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Inheritance2AssignedHostModel) Flatten(ctx context.Context, from *dnsconfig.Inheritance2AssignedHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Inheritance2AssignedHostModel{}
	}
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
}
