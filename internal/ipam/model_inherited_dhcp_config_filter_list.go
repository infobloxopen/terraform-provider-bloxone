package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type InheritedDHCPConfigFilterListModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.List   `tfsdk:"value"`
}

var InheritedDHCPConfigFilterListAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ListType{ElemType: types.StringType},
}

var InheritedDHCPConfigFilterListResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for a field of type list of identifier that represent list of filter attached to _DHCPConfig_.`,
	Attributes:          InheritedDHCPConfigFilterListResourceSchemaAttributes,
}

var InheritedDHCPConfigFilterListResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandInheritedDHCPConfigFilterList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDHCPConfigFilterList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m InheritedDHCPConfigFilterListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *InheritedDHCPConfigFilterListModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDHCPConfigFilterList {
	if m == nil {
		return nil
	}

	to := &ipam.InheritedDHCPConfigFilterList{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  ExpandFrameworkListString(ctx, m.Value, diags),
	}
	return to
}

func flattenInheritedDHCPConfigFilterList(ctx context.Context, from *ipam.InheritedDHCPConfigFilterList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InheritedDHCPConfigFilterListAttrTypes)
	}
	m := InheritedDHCPConfigFilterListModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InheritedDHCPConfigFilterListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InheritedDHCPConfigFilterListModel) flatten(ctx context.Context, from *ipam.InheritedDHCPConfigFilterList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InheritedDHCPConfigFilterListModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = FlattenFrameworkListString(ctx, from.Value, diags)

}
