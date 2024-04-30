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

type ConfigDelegationServerModel struct {
	Address      types.String `tfsdk:"address"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

var ConfigDelegationServerAttrTypes = map[string]attr.Type{
	"address":       types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

var ConfigDelegationServerResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. IP Address of nameserver.  Only required when fqdn of a delegation server falls under delegation fqdn",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Required. FQDN of nameserver.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "FQDN of nameserver in punycode.",
	},
}

func ExpandConfigDelegationServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.DelegationServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDelegationServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDelegationServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.DelegationServer {
	if m == nil {
		return nil
	}
	to := &dnsconfig.DelegationServer{
		Address: flex.ExpandStringPointer(m.Address),
		Fqdn:    flex.ExpandString(m.Fqdn),
	}
	return to
}

func FlattenConfigDelegationServer(ctx context.Context, from *dnsconfig.DelegationServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDelegationServerAttrTypes)
	}
	m := ConfigDelegationServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDelegationServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDelegationServerModel) Flatten(ctx context.Context, from *dnsconfig.DelegationServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDelegationServerModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
