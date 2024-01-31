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

type ConfigRootNSModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

var ConfigRootNSAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

var ConfigRootNSResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `IPv4 address.`,
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `FQDN.`,
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `FQDN in punycode.`,
	},
}

func ExpandConfigRootNS(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigRootNS {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigRootNSModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigRootNSModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigRootNS {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigRootNS{
		Address: flex.ExpandString(m.Address),
		Fqdn:    flex.ExpandString(m.Fqdn),
	}
	return to
}

func FlattenConfigRootNS(ctx context.Context, from *dns_config.ConfigRootNS, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigRootNSAttrTypes)
	}
	m := ConfigRootNSModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigRootNSAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigRootNSModel) Flatten(ctx context.Context, from *dns_config.ConfigRootNS, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigRootNSModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
