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

type ConfigDisplayViewModel struct {
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	View    types.String `tfsdk:"view"`
}

var ConfigDisplayViewAttrTypes = map[string]attr.Type{
	"comment": types.StringType,
	"name":    types.StringType,
	"view":    types.StringType,
}

var ConfigDisplayViewResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNS view description.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNS view name.",
	},
	"view": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandConfigDisplayView(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigDisplayView {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDisplayViewModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDisplayViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigDisplayView {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigDisplayView{
		View: flex.ExpandStringPointer(m.View),
	}
	return to
}

func FlattenConfigDisplayView(ctx context.Context, from *dns_config.ConfigDisplayView, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDisplayViewAttrTypes)
	}
	m := ConfigDisplayViewModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDisplayViewAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDisplayViewModel) Flatten(ctx context.Context, from *dns_config.ConfigDisplayView, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDisplayViewModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.View = flex.FlattenStringPointer(from.View)
}
