package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigExternalSecondaryModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
	Stealth      types.Bool   `tfsdk:"stealth"`
	TsigEnabled  types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey      types.Object `tfsdk:"tsig_key"`
}

var ConfigExternalSecondaryAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
	"stealth":       types.BoolType,
	"tsig_enabled":  types.BoolType,
	"tsig_key":      types.ObjectType{AttrTypes: ConfigTSIGKeyAttrTypes},
}

var ConfigExternalSecondaryResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "IP Address of nameserver.",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "FQDN of nameserver.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
	"stealth": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes: ConfigTSIGKeyResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandConfigExternalSecondary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.ExternalSecondary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigExternalSecondaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigExternalSecondaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.ExternalSecondary {
	if m == nil {
		return nil
	}
	to := &dnsconfig.ExternalSecondary{
		Address:     flex.ExpandString(m.Address),
		Fqdn:        flex.ExpandString(m.Fqdn),
		Stealth:     flex.ExpandBoolPointer(m.Stealth),
		TsigEnabled: flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:     ExpandConfigTSIGKey(ctx, m.TsigKey, diags),
	}
	return to
}

func FlattenConfigExternalSecondary(ctx context.Context, from *dnsconfig.ExternalSecondary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigExternalSecondaryAttrTypes)
	}
	m := ConfigExternalSecondaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigExternalSecondaryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigExternalSecondaryModel) Flatten(ctx context.Context, from *dnsconfig.ExternalSecondary, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigExternalSecondaryModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.Stealth = types.BoolPointerValue(from.Stealth)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = FlattenConfigTSIGKey(ctx, from.TsigKey, diags)
}
