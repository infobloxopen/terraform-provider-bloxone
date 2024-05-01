package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigInheritedZoneAuthorityMNameBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var ConfigInheritedZoneAuthorityMNameBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: ConfigZoneAuthorityMNameBlockAttrTypes},
}

var ConfigInheritedZoneAuthorityMNameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.SingleNestedAttribute{
		Attributes: utils.ToComputedAttributeMap(ConfigZoneAuthorityMNameBlockResourceSchemaAttributes),
		Computed:   true,
	},
}

func ExpandConfigInheritedZoneAuthorityMNameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.InheritedZoneAuthorityMNameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedZoneAuthorityMNameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedZoneAuthorityMNameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.InheritedZoneAuthorityMNameBlock {
	if m == nil {
		return nil
	}
	to := &dnsconfig.InheritedZoneAuthorityMNameBlock{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenConfigInheritedZoneAuthorityMNameBlock(ctx context.Context, from *dnsconfig.InheritedZoneAuthorityMNameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedZoneAuthorityMNameBlockAttrTypes)
	}
	m := ConfigInheritedZoneAuthorityMNameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedZoneAuthorityMNameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedZoneAuthorityMNameBlockModel) Flatten(ctx context.Context, from *dnsconfig.InheritedZoneAuthorityMNameBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedZoneAuthorityMNameBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenConfigZoneAuthorityMNameBlock(ctx, from.Value, diags)
}
