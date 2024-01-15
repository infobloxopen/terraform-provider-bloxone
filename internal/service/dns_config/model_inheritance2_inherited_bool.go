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

type Inheritance2InheritedBoolModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Bool   `tfsdk:"value"`
}

var Inheritance2InheritedBoolAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.BoolType,
}

var Inheritance2InheritedBoolResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `The inheritance setting for a field.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The human-readable display name for the object referred to by _source_.",
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"value": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The inherited value.",
	},
}

func ExpandInheritance2InheritedBool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.Inheritance2InheritedBool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedBoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Inheritance2InheritedBoolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.Inheritance2InheritedBool {
	if m == nil {
		return nil
	}
	to := &dns_config.Inheritance2InheritedBool{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenInheritance2InheritedBool(ctx context.Context, from *dns_config.Inheritance2InheritedBool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedBoolAttrTypes)
	}
	m := Inheritance2InheritedBoolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedBoolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Inheritance2InheritedBoolModel) Flatten(ctx context.Context, from *dns_config.Inheritance2InheritedBool, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Inheritance2InheritedBoolModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = types.BoolPointerValue(from.Value)
}
