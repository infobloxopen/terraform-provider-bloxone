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

type ConfigWarningModel struct {
	Message types.String `tfsdk:"message"`
	Name    types.String `tfsdk:"name"`
}

var ConfigWarningAttrTypes = map[string]attr.Type{
	"message": types.StringType,
	"name":    types.StringType,
}

var ConfigWarningResourceSchemaAttributes = map[string]schema.Attribute{
	"message": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Warning message.`,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Name of a warning.`,
	},
}

func ExpandConfigWarning(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigWarning {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigWarningModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigWarningModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigWarning {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigWarning{
		Message: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenConfigWarning(ctx context.Context, from *dns_config.ConfigWarning, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigWarningAttrTypes)
	}
	m := ConfigWarningModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigWarningAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigWarningModel) Flatten(ctx context.Context, from *dns_config.ConfigWarning, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigWarningModel{}
	}
	m.Message = flex.FlattenStringPointer(from.Message)
	m.Name = flex.FlattenStringPointer(from.Name)
}
