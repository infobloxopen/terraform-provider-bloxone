package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DataSourceModel struct {
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at"`
	Description    types.String      `tfsdk:"description"`
	Name           types.String      `tfsdk:"name"`
	SourceFields   types.List        `tfsdk:"source_fields"`
	SourceSubtypes types.List        `tfsdk:"source_subtypes"`
	SourceType     types.String      `tfsdk:"source_type"`
	Title          types.String      `tfsdk:"title"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DataSourceAttrTypes = map[string]attr.Type{
	"created_at":      timetypes.RFC3339Type{},
	"description":     types.StringType,
	"name":            types.StringType,
	"source_fields":   types.ListType{ElemType: types.ObjectType{AttrTypes: SourceFieldAttrTypes}},
	"source_subtypes": types.ListType{ElemType: types.StringType},
	"source_type":     types.StringType,
	"title":           types.StringType,
	"updated_at":      timetypes.RFC3339Type{},
}

var DataSourceResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"description": schema.StringAttribute{
		Optional: true,
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"source_fields": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SourceFieldResourceSchemaAttributes,
		},
		Optional: true,
	},
	"source_subtypes": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"source_type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		// Default: stringdefault.StaticString(SOURCETYPE_UNKNOWN), //TODO: fix this
	},
	"title": schema.StringAttribute{
		Optional: true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
}

func ExpandDataSource(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DataSource {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DataSourceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DataSourceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DataSource {
	if m == nil {
		return nil
	}
	to := &cdc.DataSource{
		Description:    flex.ExpandStringPointer(m.Description),
		Name:           flex.ExpandStringPointer(m.Name),
		SourceFields:   flex.ExpandFrameworkListNestedBlock(ctx, m.SourceFields, diags, ExpandSourceField),
		SourceSubtypes: flex.ExpandFrameworkListString(ctx, m.SourceSubtypes, diags),
		Title:          flex.ExpandStringPointer(m.Title),
		// SourceType // TODO: should have been expanded, but generator didn't know how to.
	}
	return to
}

func FlattenDataSource(ctx context.Context, from *cdc.DataSource, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DataSourceAttrTypes)
	}
	m := DataSourceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DataSourceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DataSourceModel) Flatten(ctx context.Context, from *cdc.DataSource, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DataSourceModel{}
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SourceFields = flex.FlattenFrameworkListNestedBlock(ctx, from.SourceFields, SourceFieldAttrTypes, diags, FlattenSourceField)
	m.SourceSubtypes = flex.FlattenFrameworkListString(ctx, from.SourceSubtypes, diags)
	m.Title = flex.FlattenStringPointer(from.Title)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	// SourceType // TODO: should have been flattened, but generator didn't know how to.
}
