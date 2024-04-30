package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type Inheritance2InheritedStringModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.String `tfsdk:"value"`
}

var Inheritance2InheritedStringAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.StringType,
}

var Inheritance2InheritedStringResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		MarkdownDescription: "The inheritance setting for a field. Valid values are:\n\n" +
			"  * _inherit_: Use the inherited value.\n" +
			"  * _override_: Use the value set in the object.\n\n" +
			"  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func ExpandInheritance2InheritedString(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Inheritance2InheritedString {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedStringModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Inheritance2InheritedStringModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Inheritance2InheritedString {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Inheritance2InheritedString{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenInheritance2InheritedString(ctx context.Context, from *dnsconfig.Inheritance2InheritedString, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedStringAttrTypes)
	}
	m := Inheritance2InheritedStringModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedStringAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Inheritance2InheritedStringModel) Flatten(ctx context.Context, from *dnsconfig.Inheritance2InheritedString, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Inheritance2InheritedStringModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenStringPointer(from.Value)
}
