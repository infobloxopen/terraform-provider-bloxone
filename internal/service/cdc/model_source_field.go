package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SourceFieldModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Required types.Bool   `tfsdk:"required"`
}

var SourceFieldAttrTypes = map[string]attr.Type{
	"id":       types.StringType,
	"name":     types.StringType,
	"required": types.BoolType,
}

var SourceFieldResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed: true,
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"required": schema.BoolAttribute{
		Optional: true,
	},
}

func ExpandSourceField(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.SourceField {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceFieldModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SourceFieldModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.SourceField {
	if m == nil {
		return nil
	}
	to := &cdc.SourceField{
		Name:     flex.ExpandStringPointer(m.Name),
		Required: flex.ExpandBoolPointer(m.Required),
	}
	return to
}

func FlattenSourceField(ctx context.Context, from *cdc.SourceField, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceFieldAttrTypes)
	}
	m := SourceFieldModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceFieldAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceFieldModel) Flatten(ctx context.Context, from *cdc.SourceField, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SourceFieldModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Required = types.BoolPointerValue(from.Required)
}
