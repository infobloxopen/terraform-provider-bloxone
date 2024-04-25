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

type ConfigInheritedACLItemsModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

var ConfigInheritedACLItemsAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
}

var ConfigInheritedACLItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Optional. Inheritance setting for a field. Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: utils.ToComputedAttributeMap(ConfigACLItemResourceSchemaAttributes),
		},
		Computed:            true,
		MarkdownDescription: "Inherited value.",
	},
}

func ExpandConfigInheritedACLItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.InheritedACLItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedACLItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedACLItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.InheritedACLItems {
	if m == nil {
		return nil
	}
	to := &dnsconfig.InheritedACLItems{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenConfigInheritedACLItems(ctx context.Context, from *dnsconfig.InheritedACLItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedACLItemsAttrTypes)
	}
	m := ConfigInheritedACLItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedACLItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedACLItemsModel) Flatten(ctx context.Context, from *dnsconfig.InheritedACLItems, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedACLItemsModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenFrameworkListNestedBlock(ctx, from.Value, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
}
