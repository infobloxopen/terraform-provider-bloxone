package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// NOTE: the generator produced this file using the wrong SDK model
// (cdc.Application, from the unrelated serviceApplication definition)
// instead of cdc.ApplicationData (serviceApplicationData, what
// POST/PUT /v1/applications actually accepts/returns) - cdc.Application is
// missing data/enabled/variables entirely. This file is written by hand
// against cdc.ApplicationData instead.
type ApplicationDataModel struct {
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	Data        types.Object      `tfsdk:"data"`
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Tags        types.Map         `tfsdk:"tags"`
	TagsAll     types.Map         `tfsdk:"tags_all"`
	Type        types.String      `tfsdk:"type"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
	Variables   types.Map         `tfsdk:"variables"`
}

var ApplicationDataAttrTypes = map[string]attr.Type{
	"created_at":  timetypes.RFC3339Type{},
	"data":        types.ObjectType{AttrTypes: ApplicationDataContentAttrTypes},
	"description": types.StringType,
	"enabled":     types.BoolType,
	"id":          types.Int64Type,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"tags_all":    types.MapType{ElemType: types.StringType},
	"type":        types.StringType,
	"updated_at":  timetypes.RFC3339Type{},
	"variables":   types.MapType{ElemType: types.StringType},
}

var ApplicationDataResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"data": schema.SingleNestedAttribute{
		Attributes:          ApplicationDataContentResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "The application data containing the scripts.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description of the application.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the application is enabled.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the application.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the application.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the application, including default tags.",
	},
	"type": schema.StringAttribute{
		Computed:            true, // Backend always returns "type": "PYTHON_SCRIPT"
		MarkdownDescription: "The type of the application.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"variables": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The variables for the application.",
	},
}

func ExpandApplicationData(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.ApplicationData {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ApplicationDataModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ApplicationDataModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.ApplicationData {
	if m == nil {
		return nil
	}
	to := &cdc.ApplicationData{
		Data:        ExpandApplicationDataContent(ctx, m.Data, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Enabled:     flex.ExpandBoolPointer(m.Enabled),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Variables:   expandMapStringPointer(ctx, m.Variables, diags),
	}
	return to
}

func DataSourceFlattenApplicationData(ctx context.Context, from *cdc.ApplicationData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApplicationDataAttrTypes)
	}
	m := ApplicationDataModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, ApplicationDataAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ApplicationDataModel) Flatten(ctx context.Context, from *cdc.ApplicationData, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ApplicationDataModel{}
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Data = FlattenApplicationDataContent(ctx, from.Data, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Variables = flattenMapStringPointer(ctx, from.Variables, diags)
}

// expandMapStringPointer converts a types.Map of strings into the
// *map[string]string shape the SDK uses for variables.
func expandMapStringPointer(ctx context.Context, m types.Map, diags *diag.Diagnostics) *map[string]string {
	if m.IsNull() || m.IsUnknown() {
		return nil
	}
	var result map[string]string
	diags.Append(m.ElementsAs(ctx, &result, false)...)
	if diags.HasError() {
		return nil
	}
	return &result
}

// flattenMapStringPointer converts the SDK's *map[string]string into a
// types.Map, treating an empty map as null like flex.FlattenFrameworkMapString.
func flattenMapStringPointer(ctx context.Context, from *map[string]string, diags *diag.Diagnostics) types.Map {
	if from == nil || len(*from) == 0 {
		return types.MapNull(types.StringType)
	}
	result, d := types.MapValueFrom(ctx, types.StringType, *from)
	diags.Append(d...)
	return result
}
