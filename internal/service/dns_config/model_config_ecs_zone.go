package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigECSZoneModel struct {
	Access       types.String `tfsdk:"access"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

var ConfigECSZoneAttrTypes = map[string]attr.Type{
	"access":        types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

var ConfigECSZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Access control for zone.  Allowed values: * _allow_, * _deny_.`,
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Zone FQDN.`,
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Zone FQDN in punycode.`,
	},
}

func ExpandConfigECSZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigECSZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigECSZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigECSZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigECSZone {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigECSZone{
		Access: flex.ExpandString(m.Access),
		Fqdn:   flex.ExpandString(m.Fqdn),
	}
	return to
}

func FlattenConfigECSZone(ctx context.Context, from *dns_config.ConfigECSZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigECSZoneAttrTypes)
	}
	m := ConfigECSZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigECSZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigECSZoneModel) Flatten(ctx context.Context, from *dns_config.ConfigECSZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigECSZoneModel{}
	}
	m.Access = flex.FlattenString(from.Access)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
