package cdc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	Data        types.Map         `tfsdk:"data"`
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Tags        types.Map         `tfsdk:"tags"`
	Type        types.String      `tfsdk:"type"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
	Variables   types.Map         `tfsdk:"variables"`
}

var ApplicationDataAttrTypes = map[string]attr.Type{
	"created_at":  timetypes.RFC3339Type{},
	"data":        types.MapType{ElemType: types.StringType},
	"description": types.StringType,
	"enabled":     types.BoolType,
	"id":          types.Int64Type,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"type":        types.StringType,
	"updated_at":  timetypes.RFC3339Type{},
	"variables":   types.MapType{ElemType: types.StringType},
}

var ApplicationDataResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	// data/variables are arbitrary, per-key nested JSON objects
	// (map[string]map[string]interface{} in the SDK) - there's no existing
	// precedent anywhere in this provider for representing that shape
	// directly, so each value here is a JSON-encoded string instead (e.g.
	// data = { key1 = jsonencode({nested = "value"}) }), decoded/encoded on
	// the way in and out.
	"data": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Arbitrary application configuration data. Each value must be a JSON-encoded string, e.g. `jsonencode({...})`.",
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.UseStateForUnknown(),
		},
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"type": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"variables": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Arbitrary application variables. Each value must be a JSON-encoded string, e.g. `jsonencode({...})`.",
		PlanModifiers: []planmodifier.Map{
			mapplanmodifier.UseStateForUnknown(),
		},
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
		Data:        expandNestedJSONMap(ctx, m.Data, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Enabled:     flex.ExpandBoolPointer(m.Enabled),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Type:        flex.ExpandStringPointer(m.Type),
		Variables:   expandNestedJSONMap(ctx, m.Variables, diags),
	}
	return to
}

func FlattenApplicationData(ctx context.Context, from *cdc.ApplicationData, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ApplicationDataAttrTypes)
	}
	m := ApplicationDataModel{}
	m.Flatten(ctx, from, diags)
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
	m.Data = flattenNestedJSONMap(ctx, from.Data, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Variables = flattenNestedJSONMap(ctx, from.Variables, diags)
}

// expandNestedJSONMap converts a map of JSON-encoded strings into the
// map[string]map[string]interface{} shape the API expects for data/variables.
func expandNestedJSONMap(ctx context.Context, m types.Map, diags *diag.Diagnostics) map[string]map[string]interface{} {
	if m.IsNull() || m.IsUnknown() {
		return nil
	}
	var raw map[string]string
	diags.Append(m.ElementsAs(ctx, &raw, false)...)
	if diags.HasError() {
		return nil
	}
	result := make(map[string]map[string]interface{}, len(raw))
	for k, v := range raw {
		var inner map[string]interface{}
		if err := json.Unmarshal([]byte(v), &inner); err != nil {
			diags.AddError("Invalid JSON value", fmt.Sprintf("value for key %q is not a valid JSON object: %s", k, err))
			continue
		}
		result[k] = inner
	}
	return result
}

// flattenNestedJSONMap converts the API's map[string]map[string]interface{}
// shape for data/variables back into a map of JSON-encoded strings.
func flattenNestedJSONMap(ctx context.Context, from map[string]map[string]interface{}, diags *diag.Diagnostics) types.Map {
	if from == nil {
		return types.MapNull(types.StringType)
	}
	raw := make(map[string]string, len(from))
	for k, v := range from {
		b, err := json.Marshal(v)
		if err != nil {
			diags.AddError("Invalid JSON value", fmt.Sprintf("unable to encode value for key %q: %s", k, err))
			continue
		}
		raw[k] = string(b)
	}
	result, d := types.MapValueFrom(ctx, types.StringType, raw)
	diags.Append(d...)
	return result
}
