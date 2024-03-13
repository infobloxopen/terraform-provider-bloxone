package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	internalplanmodifier "github.com/infobloxopen/terraform-provider-bloxone/internal/planmodifier"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigSortListItemModel struct {
	Acl                 types.String `tfsdk:"acl"`
	Element             types.String `tfsdk:"element"`
	PrioritizedNetworks types.List   `tfsdk:"prioritized_networks"`
	Source              types.String `tfsdk:"source"`
}

var ConfigSortListItemAttrTypes = map[string]attr.Type{
	"acl":                  types.StringType,
	"element":              types.StringType,
	"prioritized_networks": types.ListType{ElemType: types.StringType},
	"source":               types.StringType,
}

var ConfigSortListItemResourceSchemaAttributes = map[string]schema.Attribute{
	"acl": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"element": schema.StringAttribute{
		Required: true,
		MarkdownDescription: "Type of element.\n\n" +
			"  Allowed values:\n" +
			"  * _any_\n" +
			"  * _ip_\n" +
			"  * _acl_\n",
	},
	"prioritized_networks": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `Optional. The prioritized networks. If empty, the value of _source_ or networks from _acl_ is used.`,
	},
	"source": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			internalplanmodifier.UseEmptyStringForNull(),
		},
		MarkdownDescription: `Must be empty if _element_ is not _ip_.`,
	},
}

func ExpandConfigSortListItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigSortListItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigSortListItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigSortListItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigSortListItem {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigSortListItem{
		Acl:                 flex.ExpandStringPointer(m.Acl),
		Element:             flex.ExpandString(m.Element),
		PrioritizedNetworks: flex.ExpandFrameworkListString(ctx, m.PrioritizedNetworks, diags),
		Source:              flex.ExpandStringPointer(m.Source),
	}
	return to
}

func FlattenConfigSortListItem(ctx context.Context, from *dns_config.ConfigSortListItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigSortListItemAttrTypes)
	}
	m := ConfigSortListItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigSortListItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigSortListItemModel) Flatten(ctx context.Context, from *dns_config.ConfigSortListItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigSortListItemModel{}
	}
	m.Acl = flex.FlattenStringPointer(from.Acl)
	m.Element = flex.FlattenString(from.Element)
	m.PrioritizedNetworks = flex.FlattenFrameworkListString(ctx, from.PrioritizedNetworks, diags)
	m.Source = flex.FlattenStringPointer(from.Source)
}
