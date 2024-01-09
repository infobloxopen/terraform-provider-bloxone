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

type ConfigInheritedSortListItemsModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

var ConfigInheritedSortListItemsAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigSortListItemAttrTypes}},
}

var ConfigInheritedSortListItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Optional. Inheritance setting for a field. Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigSortListItemResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `Inherited value.`,
	},
}

func ExpandConfigInheritedSortListItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigInheritedSortListItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedSortListItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedSortListItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigInheritedSortListItems {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigInheritedSortListItems{
		Action: flex.ExpandStringPointer(m.Action),
		Source: flex.ExpandStringPointer(m.Source),
	}
	return to
}

func FlattenConfigInheritedSortListItems(ctx context.Context, from *dns_config.ConfigInheritedSortListItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedSortListItemsAttrTypes)
	}
	m := ConfigInheritedSortListItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedSortListItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedSortListItemsModel) Flatten(ctx context.Context, from *dns_config.ConfigInheritedSortListItems, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedSortListItemsModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, ConfigSortListItemAttrTypes, diags, FlattenConfigSortListItem)
}
