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

type ConfigZoneAuthorityModel struct {
	DefaultTtl      types.Int64  `tfsdk:"default_ttl"`
	Expire          types.Int64  `tfsdk:"expire"`
	Mname           types.String `tfsdk:"mname"`
	NegativeTtl     types.Int64  `tfsdk:"negative_ttl"`
	ProtocolMname   types.String `tfsdk:"protocol_mname"`
	ProtocolRname   types.String `tfsdk:"protocol_rname"`
	Refresh         types.Int64  `tfsdk:"refresh"`
	Retry           types.Int64  `tfsdk:"retry"`
	Rname           types.String `tfsdk:"rname"`
	UseDefaultMname types.Bool   `tfsdk:"use_default_mname"`
}

var ConfigZoneAuthorityAttrTypes = map[string]attr.Type{
	"default_ttl":       types.Int64Type,
	"expire":            types.Int64Type,
	"mname":             types.StringType,
	"negative_ttl":      types.Int64Type,
	"protocol_mname":    types.StringType,
	"protocol_rname":    types.StringType,
	"refresh":           types.Int64Type,
	"retry":             types.Int64Type,
	"rname":             types.StringType,
	"use_default_mname": types.BoolType,
}

var ConfigZoneAuthorityResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. ZoneAuthority default ttl for resource records in zone (value in seconds).  Defaults to 28800.`,
	},
	"expire": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. ZoneAuthority expire time in seconds.  Defaults to 2419200.`,
	},
	"mname": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Defaults to empty.`,
	},
	"negative_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. ZoneAuthority negative caching (minimum) ttl in seconds.  Defaults to 900.`,
	},
	"protocol_mname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Optional. ZoneAuthority master name server in punycode.  Defaults to empty.`,
	},
	"protocol_rname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Optional. A domain name which specifies the mailbox of the person responsible for this zone.  Defaults to empty.`,
	},
	"refresh": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. ZoneAuthority refresh.  Defaults to 10800.`,
	},
	"retry": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. ZoneAuthority retry.  Defaults to 3600.`,
	},
	"rname": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Optional. ZoneAuthority rname.  Defaults to empty.`,
	},
	"use_default_mname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Use default value for master name server.  Defaults to true.`,
	},
}

func ExpandConfigZoneAuthority(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.ZoneAuthority {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigZoneAuthorityModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigZoneAuthorityModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.ZoneAuthority {
	if m == nil {
		return nil
	}
	to := &dnsconfig.ZoneAuthority{
		DefaultTtl:      flex.ExpandInt64Pointer(m.DefaultTtl),
		Expire:          flex.ExpandInt64Pointer(m.Expire),
		Mname:           flex.ExpandStringPointer(m.Mname),
		NegativeTtl:     flex.ExpandInt64Pointer(m.NegativeTtl),
		Refresh:         flex.ExpandInt64Pointer(m.Refresh),
		Retry:           flex.ExpandInt64Pointer(m.Retry),
		Rname:           flex.ExpandStringPointer(m.Rname),
		UseDefaultMname: flex.ExpandBoolPointer(m.UseDefaultMname),
	}
	return to
}

func FlattenConfigZoneAuthority(ctx context.Context, from *dnsconfig.ZoneAuthority, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigZoneAuthorityAttrTypes)
	}
	m := ConfigZoneAuthorityModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigZoneAuthorityAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigZoneAuthorityModel) Flatten(ctx context.Context, from *dnsconfig.ZoneAuthority, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigZoneAuthorityModel{}
	}
	m.DefaultTtl = flex.FlattenInt64(int64(*from.DefaultTtl))
	m.Expire = flex.FlattenInt64(int64(*from.Expire))
	m.Mname = flex.FlattenStringPointer(from.Mname)
	m.NegativeTtl = flex.FlattenInt64(int64(*from.NegativeTtl))
	m.ProtocolMname = flex.FlattenStringPointer(from.ProtocolMname)
	m.ProtocolRname = flex.FlattenStringPointer(from.ProtocolRname)
	m.Refresh = flex.FlattenInt64(int64(*from.Refresh))
	m.Retry = flex.FlattenInt64(int64(*from.Retry))
	m.Rname = flex.FlattenStringPointer(from.Rname)
	m.UseDefaultMname = types.BoolPointerValue(from.UseDefaultMname)
}
