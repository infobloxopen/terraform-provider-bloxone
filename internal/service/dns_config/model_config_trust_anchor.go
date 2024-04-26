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

type ConfigTrustAnchorModel struct {
	Algorithm    types.Int64  `tfsdk:"algorithm"`
	ProtocolZone types.String `tfsdk:"protocol_zone"`
	PublicKey    types.String `tfsdk:"public_key"`
	Sep          types.Bool   `tfsdk:"sep"`
	Zone         types.String `tfsdk:"zone"`
}

var ConfigTrustAnchorAttrTypes = map[string]attr.Type{
	"algorithm":     types.Int64Type,
	"protocol_zone": types.StringType,
	"public_key":    types.StringType,
	"sep":           types.BoolType,
	"zone":          types.StringType,
}

var ConfigTrustAnchorResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.Int64Attribute{
		Required: true,
	},
	"protocol_zone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Zone FQDN in punycode.`,
	},
	"public_key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `DNSSEC key data. Non-empty, valid base64 string.`,
	},
	"sep": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Secure Entry Point flag.  Defaults to _true_.`,
	},
	"zone": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Zone FQDN.`,
	},
}

func ExpandConfigTrustAnchor(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.TrustAnchor {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigTrustAnchorModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigTrustAnchorModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.TrustAnchor {
	if m == nil {
		return nil
	}
	to := &dnsconfig.TrustAnchor{
		Algorithm: flex.ExpandInt64(m.Algorithm),
		PublicKey: flex.ExpandString(m.PublicKey),
		Sep:       flex.ExpandBoolPointer(m.Sep),
		Zone:      flex.ExpandString(m.Zone),
	}
	return to
}

func FlattenConfigTrustAnchor(ctx context.Context, from *dnsconfig.TrustAnchor, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigTrustAnchorAttrTypes)
	}
	m := ConfigTrustAnchorModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigTrustAnchorAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigTrustAnchorModel) Flatten(ctx context.Context, from *dnsconfig.TrustAnchor, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigTrustAnchorModel{}
	}
	m.Algorithm = flex.FlattenInt64(int64(from.Algorithm))
	m.ProtocolZone = flex.FlattenStringPointer(from.ProtocolZone)
	m.PublicKey = flex.FlattenString(from.PublicKey)
	m.Sep = types.BoolPointerValue(from.Sep)
	m.Zone = flex.FlattenString(from.Zone)
}
