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

type ConfigForwarderModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

var ConfigForwarderAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

func ConfigForwarderResourceSchemaAttributes(fqdnOptional bool) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"address": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: `Server IP address.`,
		},
		"fqdn": schema.StringAttribute{
			Required:            !fqdnOptional,
			Optional:            fqdnOptional,
			MarkdownDescription: `Server FQDN.`,
		},
		"protocol_fqdn": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: `Server FQDN in punycode.`,
		},
	}
}

func ExpandConfigForwarder(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigForwarder {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigForwarderModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigForwarderModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigForwarder {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigForwarder{
		Address: flex.ExpandString(m.Address),
		Fqdn:    flex.ExpandStringPointer(m.Fqdn),
	}
	return to
}

func FlattenConfigForwarder(ctx context.Context, from *dns_config.ConfigForwarder, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigForwarderAttrTypes)
	}
	m := ConfigForwarderModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigForwarderAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigForwarderModel) Flatten(ctx context.Context, from *dns_config.ConfigForwarder, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigForwarderModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	if from.Fqdn == nil || *from.Fqdn == "" {
		m.Fqdn = types.StringNull()
	} else {
		m.Fqdn = types.StringValue(*from.Fqdn)
	}
}
