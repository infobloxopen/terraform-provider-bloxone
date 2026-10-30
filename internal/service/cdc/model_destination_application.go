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

type DestinationApplicationModel struct {
	Application      types.Int64       `tfsdk:"application"`
	Context          types.Map         `tfsdk:"context"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	Description      types.String      `tfsdk:"description"`
	Enabled          types.Bool        `tfsdk:"enabled"`
	Id               types.Int64       `tfsdk:"id"`
	Name             types.String      `tfsdk:"name"`
	OutputDataFormat types.String      `tfsdk:"output_data_format"`
	Tags             types.Map         `tfsdk:"tags"`
	TagsAll          types.Map         `tfsdk:"tags_all"`
	UpdatedAt        timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationApplicationAttrTypes = map[string]attr.Type{
	"application":        types.Int64Type,
	"context":            types.MapType{ElemType: types.StringType},
	"created_at":         timetypes.RFC3339Type{},
	"description":        types.StringType,
	"enabled":            types.BoolType,
	"id":                 types.Int64Type,
	"name":               types.StringType,
	"output_data_format": types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"updated_at":         timetypes.RFC3339Type{},
}

var DestinationApplicationResourceSchemaAttributes = map[string]schema.Attribute{
	"application": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The identifier of the application (automation script) that receives the data.",
	},
	"context": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The variables passed to the application.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description of the application destination.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether the application destination is enabled. Defaults to _false_.",
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
		MarkdownDescription: "The name of the application destination.",
	},
	"output_data_format": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("STREAM"),
		MarkdownDescription: "The output data format of the application destination. Defaults to _STREAM_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the application destination.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the application destination, including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandDestinationApplication(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationApplication {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationApplicationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationApplicationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationApplication {
	if m == nil {
		return nil
	}
	// context is always sent, as {} when unset: on update the backend stores a
	// missing context as invalid JSON and fails with
	// "pq: invalid input syntax for type json". It is sent here rather than set
	// as the schema default, so the planned value stays null and matches the
	// null that an empty context flattens to.
	contextVars := flex.ExpandFrameworkMapString(ctx, m.Context, diags)
	if contextVars == nil {
		contextVars = map[string]interface{}{}
	}
	to := &cdc.DestinationApplication{
		Application:      flex.ExpandInt64Pointer(m.Application),
		Context:          contextVars,
		Description:      flex.ExpandStringPointer(m.Description),
		Enabled:          flex.ExpandBoolPointer(m.Enabled),
		Name:             flex.ExpandStringPointer(m.Name),
		OutputDataFormat: flex.ExpandStringPointer(m.OutputDataFormat),
		Tags:             flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenDestinationApplication(ctx context.Context, from *cdc.DestinationApplication, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationApplicationAttrTypes)
	}
	m := DestinationApplicationModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, DestinationApplicationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationApplicationModel) Flatten(ctx context.Context, from *cdc.DestinationApplication, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationApplicationModel{}
	}
	m.Application = flex.FlattenInt64Pointer(from.Application)
	// The UI only stores string values, but the API accepts any JSON value, so
	// non-string values are kept as their JSON representation instead of failing the read.
	m.Context = flex.FlattenFrameworkMapStringAny(ctx, from.Context, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutputDataFormat = flex.FlattenStringPointer(from.OutputDataFormat)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
