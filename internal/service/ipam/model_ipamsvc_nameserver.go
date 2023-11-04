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
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type IpamsvcNameserverModel struct {
	ClientPrincipal       types.String `tfsdk:"client_principal"`
	GssTsigFallback       types.Bool   `tfsdk:"gss_tsig_fallback"`
	KerberosRekeyInterval types.Int64  `tfsdk:"kerberos_rekey_interval"`
	KerberosRetryInterval types.Int64  `tfsdk:"kerberos_retry_interval"`
	KerberosTkeyLifetime  types.Int64  `tfsdk:"kerberos_tkey_lifetime"`
	KerberosTkeyProtocol  types.String `tfsdk:"kerberos_tkey_protocol"`
	Nameserver            types.String `tfsdk:"nameserver"`
	ServerPrincipal       types.String `tfsdk:"server_principal"`
}

var IpamsvcNameserverAttrTypes = map[string]attr.Type{
	"client_principal":        types.StringType,
	"gss_tsig_fallback":       types.BoolType,
	"kerberos_rekey_interval": types.Int64Type,
	"kerberos_retry_interval": types.Int64Type,
	"kerberos_tkey_lifetime":  types.Int64Type,
	"kerberos_tkey_protocol":  types.StringType,
	"nameserver":              types.StringType,
	"server_principal":        types.StringType,
}

var IpamsvcNameserverResourceSchemaAttributes = map[string]schema.Attribute{
	"client_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name. It uses the typical Kerberos notation: <SERVICE-NAME>/<server-domain-name>@<REALM>.  Defaults to empty.`,
	},
	"gss_tsig_fallback": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.`,
	},
	"kerberos_rekey_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the _kerberos_rekey_interval_ value.  Defaults to 120 seconds.`,
	},
	"kerberos_retry_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) to retry to create a key if any error occurred previously for any configured external DNS server.  Defaults to 30 seconds.`,
	},
	"kerberos_tkey_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Lifetime (in seconds) of GSS-TSIG keys in the TKEY protocol.  Defaults to 160 seconds.`,
	},
	"kerberos_tkey_protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Determines which protocol is used to establish the security context with the external DNS servers, TCP or UDP.  Defaults to _tcp_.`,
	},
	"nameserver": schema.StringAttribute{
		Optional: true,
	},
	"server_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name of this DNS server that will receive updates.  Defaults to empty.`,
	},
}

func ExpandIpamsvcNameserver(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcNameserver {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcNameserverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcNameserverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcNameserver {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcNameserver{
		ClientPrincipal:       m.ClientPrincipal.ValueStringPointer(),
		GssTsigFallback:       m.GssTsigFallback.ValueBoolPointer(),
		KerberosRekeyInterval: utils.Ptr(int64(m.KerberosRekeyInterval.ValueInt64())),
		KerberosRetryInterval: utils.Ptr(int64(m.KerberosRetryInterval.ValueInt64())),
		KerberosTkeyLifetime:  utils.Ptr(int64(m.KerberosTkeyLifetime.ValueInt64())),
		KerberosTkeyProtocol:  m.KerberosTkeyProtocol.ValueStringPointer(),
		Nameserver:            m.Nameserver.ValueStringPointer(),
		ServerPrincipal:       m.ServerPrincipal.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcNameserver(ctx context.Context, from *ipam.IpamsvcNameserver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcNameserverAttrTypes)
	}
	m := IpamsvcNameserverModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcNameserverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcNameserverModel) Flatten(ctx context.Context, from *ipam.IpamsvcNameserver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcNameserverModel{}
	}
	m.ClientPrincipal = flex.FlattenStringPointer(from.ClientPrincipal)
	m.GssTsigFallback = types.BoolPointerValue(from.GssTsigFallback)
	m.KerberosRekeyInterval = flex.FlattenInt64(int64(*from.KerberosRekeyInterval))
	m.KerberosRetryInterval = flex.FlattenInt64(int64(*from.KerberosRetryInterval))
	m.KerberosTkeyLifetime = flex.FlattenInt64(int64(*from.KerberosTkeyLifetime))
	m.KerberosTkeyProtocol = flex.FlattenStringPointer(from.KerberosTkeyProtocol)
	m.Nameserver = flex.FlattenStringPointer(from.Nameserver)
	m.ServerPrincipal = flex.FlattenStringPointer(from.ServerPrincipal)
}
