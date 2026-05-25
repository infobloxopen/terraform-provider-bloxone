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

type ConfigNameserverModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	Host         types.String `tfsdk:"host"`
	Origin       types.String `tfsdk:"origin"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
	Role         types.String `tfsdk:"role"`
	Stealth      types.Bool   `tfsdk:"stealth"`
	TsigEnabled  types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey      types.Object `tfsdk:"tsig_key"`
}

var ConfigNameserverAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"host":          types.StringType,
	"origin":        types.StringType,
	"protocol_fqdn": types.StringType,
	"role":          types.StringType,
	"stealth":       types.BoolType,
	"tsig_enabled":  types.BoolType,
	"tsig_key":      types.ObjectType{AttrTypes: ConfigTSIGKeyAttrTypes},
}

var ConfigNameserverResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Required only if _origin_ is _external_. IP Address of the nameserver.",
	},
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Required only if _origin_ is _external_. FQDN of the nameserver.",
	},
	"host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"origin": schema.StringAttribute{
		Required: true,
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of the nameserver in punycode.",
	},
	"role": schema.StringAttribute{
		Required: true,
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from a primary.",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes: ConfigTSIGKeyResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandConfigNameserver(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Nameserver {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigNameserverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigNameserverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Nameserver {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Nameserver{
		Address:     flex.ExpandStringPointer(m.Address),
		Fqdn:        flex.ExpandStringPointer(m.Fqdn),
		Host:        flex.ExpandStringPointer(m.Host),
		Origin:      flex.ExpandString(m.Origin),
		Role:        flex.ExpandString(m.Role),
		Stealth:     flex.ExpandBoolPointer(m.Stealth),
		TsigEnabled: flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:     ExpandConfigTSIGKey(ctx, m.TsigKey, diags),
	}
	return to
}

func FlattenConfigNameserver(ctx context.Context, from *dnsconfig.Nameserver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigNameserverAttrTypes)
	}
	m := ConfigNameserverModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigNameserverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigNameserverModel) Flatten(ctx context.Context, from *dnsconfig.Nameserver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigNameserverModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Host = flex.FlattenStringPointer(from.Host)
	m.Origin = flex.FlattenString(from.Origin)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.Role = flex.FlattenString(from.Role)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = FlattenConfigTSIGKey(ctx, from.TsigKey, diags)
}
