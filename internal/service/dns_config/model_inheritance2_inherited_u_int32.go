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

type Inheritance2InheritedUInt32Model struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Int64  `tfsdk:"value"`
}

var Inheritance2InheritedUInt32AttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.Int64Type,
}

var Inheritance2InheritedUInt32ResourceSchemaAttributes = map[string]schema.Attribute{
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
	"value": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The inherited value.`,
	},
}

func ExpandInheritance2InheritedUInt32(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.Inheritance2InheritedUInt32 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Inheritance2InheritedUInt32Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *Inheritance2InheritedUInt32Model) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Inheritance2InheritedUInt32 {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Inheritance2InheritedUInt32{
		Action: flex.ExpandStringPointer(m.Action),
	}
	return to
}

func FlattenInheritance2InheritedUInt32(ctx context.Context, from *dnsconfig.Inheritance2InheritedUInt32, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Inheritance2InheritedUInt32AttrTypes)
	}
	m := Inheritance2InheritedUInt32Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Inheritance2InheritedUInt32AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *Inheritance2InheritedUInt32Model) Flatten(ctx context.Context, from *dnsconfig.Inheritance2InheritedUInt32, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = Inheritance2InheritedUInt32Model{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = flex.FlattenInt64(int64(*from.Value))
}
