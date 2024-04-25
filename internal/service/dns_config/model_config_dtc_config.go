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

type ConfigDTCConfigModel struct {
	DefaultTtl types.Int64 `tfsdk:"default_ttl"`
}

var ConfigDTCConfigAttrTypes = map[string]attr.Type{
	"default_ttl": types.Int64Type,
}

var ConfigDTCConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `Optional. Default TTL for synthesized DTC records (value in seconds).  Defaults to 300.`,
	},
}

func ExpandConfigDTCConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.DTCConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDTCConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDTCConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.DTCConfig {
	if m == nil {
		return nil
	}
	to := &dnsconfig.DTCConfig{
		DefaultTtl: flex.ExpandInt64Pointer(m.DefaultTtl),
	}
	return to
}

func FlattenConfigDTCConfig(ctx context.Context, from *dnsconfig.DTCConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDTCConfigAttrTypes)
	}
	m := ConfigDTCConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDTCConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDTCConfigModel) Flatten(ctx context.Context, from *dnsconfig.DTCConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDTCConfigModel{}
	}
	m.DefaultTtl = flex.FlattenInt64(int64(*from.DefaultTtl))
}
