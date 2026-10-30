package cdc

import (
	"context"
	"encoding/json"
	"fmt"

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
	Data        types.String      `tfsdk:"data"`
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
	FileData    types.String      `tfsdk:"file_data"`
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
	"data":        types.StringType,
	"description": types.StringType,
	"enabled":     types.BoolType,
	"file_data":   types.StringType,
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
	"data": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The application data containing the scripts, as a JSON-encoded object, e.g. `jsonencode({ scripts = [{ filename = \"script_1.py\", content = file(\"script_1.py\") }] })`. The value is stored exactly as configured; the data returned by the backend is available in _file_data_.",
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
	"file_data": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The application data as returned by the backend, as a JSON-encoded object.",
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
		Computed:            true,
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
		Data:        expandApplicationDataJSON(m.Data, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Enabled:     flex.ExpandBoolPointer(m.Enabled),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Variables:   flex.ExpandFrameworkMapString(ctx, m.Variables, diags),
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
	// data has no fixed schema and the backend may return it in a different
	// shape than it was sent, so the configured value is kept as is and the
	// response is exposed through file_data instead. data is only populated
	// from the response when there is no configured value, e.g. on import or
	// in the data source.
	m.FileData = flattenApplicationDataJSON(from.Data, diags)
	if m.Data.IsNull() || m.Data.IsUnknown() {
		m.Data = m.FileData
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Variables = flex.FlattenFrameworkMapString(ctx, from.Variables, diags)
}

// expandApplicationDataJSON decodes the JSON-encoded data attribute into the
// object the SDK sends to the backend.
func expandApplicationDataJSON(v types.String, diags *diag.Diagnostics) map[string]interface{} {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(v.ValueString()), &result); err != nil {
		diags.AddError("Invalid data", fmt.Sprintf("data must be a JSON-encoded object, e.g. with jsonencode(): %s", err))
		return nil
	}
	return result
}

// flattenApplicationDataJSON encodes the data returned by the backend as a
// JSON string.
func flattenApplicationDataJSON(from map[string]interface{}, diags *diag.Diagnostics) types.String {
	if from == nil {
		return types.StringNull()
	}
	b, err := json.Marshal(from)
	if err != nil {
		diags.AddError("Serialization Error", fmt.Sprintf("Unable to serialize application data: %s", err))
		return types.StringNull()
	}
	return types.StringValue(string(b))
}
