package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigExternalPrimaryModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	Nsg          types.String `tfsdk:"nsg"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
	TsigEnabled  types.Bool   `tfsdk:"tsig_enabled"`
	TsigKey      types.Object `tfsdk:"tsig_key"`
	Type         types.String `tfsdk:"type"`
}

var ConfigExternalPrimaryAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"nsg":           types.StringType,
	"protocol_fqdn": types.StringType,
	"tsig_enabled":  types.BoolType,
	"tsig_key":      types.ObjectType{AttrTypes: ConfigTSIGKeyAttrTypes},
	"type":          types.StringType,
}

var ConfigExternalPrimaryResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Required only if _type_ is _server_. IP Address of nameserver.",
	},
	"fqdn": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Required only if _type_ is _server_. FQDN of nameserver.",
	},
	"nsg": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
	"tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary.",
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes: ConfigTSIGKeyResourceSchemaAttributes,
		Optional:   true,
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Allowed values: * _nsg_, * _primary_.",
	},
}

func ExpandConfigExternalPrimary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigExternalPrimary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigExternalPrimaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigExternalPrimaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigExternalPrimary {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigExternalPrimary{
		Address:     flex.ExpandStringPointer(m.Address),
		Fqdn:        flex.ExpandStringPointer(m.Fqdn),
		Nsg:         flex.ExpandStringPointer(m.Nsg),
		TsigEnabled: flex.ExpandBoolPointer(m.TsigEnabled),
		TsigKey:     ExpandConfigTSIGKey(ctx, m.TsigKey, diags),
		Type:        flex.ExpandString(m.Type),
	}
	return to
}

func FlattenConfigExternalPrimary(ctx context.Context, from *dns_config.ConfigExternalPrimary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigExternalPrimaryAttrTypes)
	}
	m := ConfigExternalPrimaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigExternalPrimaryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigExternalPrimaryModel) Flatten(ctx context.Context, from *dns_config.ConfigExternalPrimary, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigExternalPrimaryModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.Nsg = flex.FlattenStringPointer(from.Nsg)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.TsigEnabled = types.BoolPointerValue(from.TsigEnabled)
	m.TsigKey = FlattenConfigTSIGKey(ctx, from.TsigKey, diags)
	m.Type = flex.FlattenString(from.Type)
}
