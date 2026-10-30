package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// SourceLogTypeModel is read-only: source_log_types is derived by the backend
// from the flow's source_data_types.
type SourceLogTypeModel struct {
	LogFields   types.List   `tfsdk:"log_fields"`
	LogSubtypes types.List   `tfsdk:"log_subtypes"`
	LogType     types.String `tfsdk:"log_type"`
}

var SourceLogTypeAttrTypes = map[string]attr.Type{
	"log_fields":   types.ListType{ElemType: types.ObjectType{AttrTypes: SourceFieldAttrTypes}},
	"log_subtypes": types.ListType{ElemType: types.StringType},
	"log_type":     types.StringType,
}

var SourceLogTypeResourceSchemaAttributes = map[string]schema.Attribute{
	"log_fields": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"id": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "The identifier of the log field.",
				},
				"name": schema.StringAttribute{
					Computed:            true,
					MarkdownDescription: "The name of the log field.",
				},
				"required": schema.BoolAttribute{
					Computed:            true,
					MarkdownDescription: "Indicates whether the log field is always collected.",
				},
			},
		},
		Computed:            true,
		MarkdownDescription: "The log fields collected for the log type.",
	},
	"log_subtypes": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The log subtypes of the log type.",
	},
	"log_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The log type, e.g. `RPZ_LOG`.",
	},
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
	m.LogFields = flex.FlattenFrameworkListNestedBlock(ctx, from.LogFields, SourceFieldAttrTypes, diags, FlattenSourceField)
	m.LogSubtypes = flex.FlattenFrameworkListString(ctx, from.LogSubtypes, diags)
	m.LogType = flex.FlattenStringPointer(from.LogType)
}
