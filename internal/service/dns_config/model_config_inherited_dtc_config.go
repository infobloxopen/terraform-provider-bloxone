package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
)

type ConfigInheritedDtcConfigModel struct {
	DefaultTtl types.Object `tfsdk:"default_ttl"`
}

var ConfigInheritedDtcConfigAttrTypes = map[string]attr.Type{
	"default_ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
}

var ConfigInheritedDtcConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandConfigInheritedDtcConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.InheritedDtcConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedDtcConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedDtcConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.InheritedDtcConfig {
	if m == nil {
		return nil
	}
	to := &dnsconfig.InheritedDtcConfig{
		DefaultTtl: ExpandInheritance2InheritedUInt32(ctx, m.DefaultTtl, diags),
	}
	return to
}

func FlattenConfigInheritedDtcConfig(ctx context.Context, from *dnsconfig.InheritedDtcConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedDtcConfigAttrTypes)
	}
	m := ConfigInheritedDtcConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedDtcConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedDtcConfigModel) Flatten(ctx context.Context, from *dnsconfig.InheritedDtcConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedDtcConfigModel{}
	}
	m.DefaultTtl = FlattenInheritance2InheritedUInt32(ctx, from.DefaultTtl, diags)
}
