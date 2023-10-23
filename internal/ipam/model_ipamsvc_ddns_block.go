package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcDDNSBlockModel struct {
	ClientPrincipal       types.String `tfsdk:"client_principal"`
	DdnsDomain            types.String `tfsdk:"ddns_domain"`
	DdnsEnabled           types.Bool   `tfsdk:"ddns_enabled"`
	DdnsSendUpdates       types.Bool   `tfsdk:"ddns_send_updates"`
	DdnsZones             types.List   `tfsdk:"ddns_zones"`
	GssTsigFallback       types.Bool   `tfsdk:"gss_tsig_fallback"`
	KerberosKdc           types.String `tfsdk:"kerberos_kdc"`
	KerberosKeys          types.List   `tfsdk:"kerberos_keys"`
	KerberosRekeyInterval types.Int64  `tfsdk:"kerberos_rekey_interval"`
	KerberosRetryInterval types.Int64  `tfsdk:"kerberos_retry_interval"`
	KerberosTkeyLifetime  types.Int64  `tfsdk:"kerberos_tkey_lifetime"`
	KerberosTkeyProtocol  types.String `tfsdk:"kerberos_tkey_protocol"`
	ServerPrincipal       types.String `tfsdk:"server_principal"`
}

var IpamsvcDDNSBlockAttrTypes = map[string]attr.Type{
	"client_principal":        types.StringType,
	"ddns_domain":             types.StringType,
	"ddns_enabled":            types.BoolType,
	"ddns_send_updates":       types.BoolType,
	"ddns_zones":              types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcDDNSZoneAttrTypes}},
	"gss_tsig_fallback":       types.BoolType,
	"kerberos_kdc":            types.StringType,
	"kerberos_keys":           types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcKerberosKeyAttrTypes}},
	"kerberos_rekey_interval": types.Int64Type,
	"kerberos_retry_interval": types.Int64Type,
	"kerberos_tkey_lifetime":  types.Int64Type,
	"kerberos_tkey_protocol":  types.StringType,
	"server_principal":        types.StringType,
}

var IpamsvcDDNSBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The dynamic DNS configuration.`,
	Attributes:          IpamsvcDDNSBlockResourceSchemaAttributes,
}

var IpamsvcDDNSBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"client_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name. It uses the typical Kerberos notation: &lt;SERVICE-NAME&gt;/&lt;server-domain-name&gt;@&lt;REALM&gt;.  Defaults to empty.`,
	},
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain name for DDNS.`,
	},
	"ddns_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if DDNS is enabled.`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Determines if DDNS updates are enabled at this level.`,
	},
	"ddns_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcDDNSZoneResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DDNS zones.`,
	},
	"gss_tsig_fallback": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The behavior when GSS-TSIG should be used (a matching external DNS server is configured) but no GSS-TSIG key is available. If configured to _false_ (the default) this DNS server is skipped, if configured to _true_ the DNS server is ignored and the DNS update is sent with the configured DHCP-DDNS protection e.g. TSIG key or without any protection when none was configured.  Defaults to _false_.`,
	},
	"kerberos_kdc": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Address of Kerberos Key Distribution Center.  Defaults to empty.`,
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcKerberosKeyResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.`,
	},
	"kerberos_rekey_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Time interval (in seconds) the keys for each configured external DNS server are checked for rekeying, i.e. a new key is created to replace the current usable one when its age is greater than the rekey_interval value.  Defaults to 120 seconds.`,
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
	"server_principal": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The Kerberos principal name of the external DNS server that will receive updates.  Defaults to empty.`,
	},
}

func expandIpamsvcDDNSBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDDNSBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDDNSBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDDNSBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDDNSBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDDNSBlock{
		ClientPrincipal:       m.ClientPrincipal.ValueStringPointer(),
		DdnsDomain:            m.DdnsDomain.ValueStringPointer(),
		DdnsEnabled:           m.DdnsEnabled.ValueBoolPointer(),
		DdnsSendUpdates:       m.DdnsSendUpdates.ValueBoolPointer(),
		DdnsZones:             ExpandFrameworkListNestedBlock(ctx, m.DdnsZones, diags, expandIpamsvcDDNSZone),
		GssTsigFallback:       m.GssTsigFallback.ValueBoolPointer(),
		KerberosKdc:           m.KerberosKdc.ValueStringPointer(),
		KerberosKeys:          ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, expandIpamsvcKerberosKey),
		KerberosRekeyInterval: ptr(int64(m.KerberosRekeyInterval.ValueInt64())),
		KerberosRetryInterval: ptr(int64(m.KerberosRetryInterval.ValueInt64())),
		KerberosTkeyLifetime:  ptr(int64(m.KerberosTkeyLifetime.ValueInt64())),
		KerberosTkeyProtocol:  m.KerberosTkeyProtocol.ValueStringPointer(),
		ServerPrincipal:       m.ServerPrincipal.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcDDNSBlock(ctx context.Context, from *ipam.IpamsvcDDNSBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSBlockAttrTypes)
	}
	m := IpamsvcDDNSBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcDDNSBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSBlockModel{}
	}

	m.ClientPrincipal = types.StringPointerValue(from.ClientPrincipal)
	m.DdnsDomain = types.StringPointerValue(from.DdnsDomain)
	m.DdnsEnabled = types.BoolPointerValue(from.DdnsEnabled)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
	m.DdnsZones = FlattenFrameworkListNestedBlock(ctx, from.DdnsZones, IpamsvcDDNSZoneAttrTypes, diags, flattenIpamsvcDDNSZone)
	m.GssTsigFallback = types.BoolPointerValue(from.GssTsigFallback)
	m.KerberosKdc = types.StringPointerValue(from.KerberosKdc)
	m.KerberosKeys = FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, IpamsvcKerberosKeyAttrTypes, diags, flattenIpamsvcKerberosKey)
	m.KerberosRekeyInterval = types.Int64Value(int64(*from.KerberosRekeyInterval))
	m.KerberosRetryInterval = types.Int64Value(int64(*from.KerberosRetryInterval))
	m.KerberosTkeyLifetime = types.Int64Value(int64(*from.KerberosTkeyLifetime))
	m.KerberosTkeyProtocol = types.StringPointerValue(from.KerberosTkeyProtocol)
	m.ServerPrincipal = types.StringPointerValue(from.ServerPrincipal)

}
