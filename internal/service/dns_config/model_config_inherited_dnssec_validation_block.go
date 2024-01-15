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

type ConfigInheritedDNSSECValidationBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var ConfigInheritedDNSSECValidationBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: ConfigDNSSECValidationBlockAttrTypes},
}

var ConfigInheritedDNSSECValidationBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes: ConfigDNSSECValidationBlockResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandConfigInheritedDNSSECValidationBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigInheritedDNSSECValidationBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedDNSSECValidationBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedDNSSECValidationBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigInheritedDNSSECValidationBlock {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigInheritedDNSSECValidationBlock{
		Action: flex.ExpandStringPointer(m.Action),
		Value:  ExpandConfigDNSSECValidationBlock(ctx, m.Value, diags),
	}
	return to
}

func FlattenConfigInheritedDNSSECValidationBlock(ctx context.Context, from *dns_config.ConfigInheritedDNSSECValidationBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedDNSSECValidationBlockAttrTypes)
	}
	m := ConfigInheritedDNSSECValidationBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedDNSSECValidationBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedDNSSECValidationBlockModel) Flatten(ctx context.Context, from *dns_config.ConfigInheritedDNSSECValidationBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedDNSSECValidationBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenConfigDNSSECValidationBlock(ctx, from.Value, diags)
}
