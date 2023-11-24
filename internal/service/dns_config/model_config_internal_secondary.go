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

type ConfigInternalSecondaryModel struct {
	Host types.String `tfsdk:"host"`
}

var ConfigInternalSecondaryAttrTypes = map[string]attr.Type{
	"host": types.StringType,
}

var ConfigInternalSecondaryResourceSchemaAttributes = map[string]schema.Attribute{
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandConfigInternalSecondary(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigInternalSecondary {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInternalSecondaryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInternalSecondaryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigInternalSecondary {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigInternalSecondary{
		Host: flex.ExpandString(m.Host),
	}
	return to
}

func FlattenConfigInternalSecondary(ctx context.Context, from *dns_config.ConfigInternalSecondary, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInternalSecondaryAttrTypes)
	}
	m := ConfigInternalSecondaryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInternalSecondaryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInternalSecondaryModel) Flatten(ctx context.Context, from *dns_config.ConfigInternalSecondary, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInternalSecondaryModel{}
	}
	m.Host = flex.FlattenString(from.Host)
}
