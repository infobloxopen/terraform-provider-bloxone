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

type SourceLogTypeModel struct {
	LogFields   types.List   `tfsdk:"log_fields"`
	LogSubtypes types.List   `tfsdk:"log_subtypes"`
	LogType     types.String `tfsdk:"log_type"`
}

var SourceLogTypeAttrTypes = map[string]attr.Type{
	"log_fields":   types.ListType{ElemType: types.StringType},
	"log_subtypes": types.ListType{ElemType: types.StringType},
	"log_type":     types.StringType,
}

var SourceLogTypeResourceSchemaAttributes = map[string]schema.Attribute{
	"log_fields": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"log_subtypes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"log_type": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandSourceLogType(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.SourceLogType {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceLogTypeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SourceLogTypeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.SourceLogType {
	if m == nil {
		return nil
	}
	to := &cdc.SourceLogType{
		LogFields:   flex.ExpandFrameworkListString(ctx, m.LogFields, diags),
		LogSubtypes: flex.ExpandFrameworkListString(ctx, m.LogSubtypes, diags),
		LogType:     flex.ExpandStringPointer(m.LogType),
	}
	return to
}

func FlattenSourceLogType(ctx context.Context, from *cdc.SourceLogType, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceLogTypeAttrTypes)
	}
	m := SourceLogTypeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceLogTypeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceLogTypeModel) Flatten(ctx context.Context, from *cdc.SourceLogType, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SourceLogTypeModel{}
	}
	m.LogFields = flex.FlattenFrameworkListString(ctx, from.LogFields, diags)
	m.LogSubtypes = flex.FlattenFrameworkListString(ctx, from.LogSubtypes, diags)
	m.LogType = flex.FlattenStringPointer(from.LogType)
}
