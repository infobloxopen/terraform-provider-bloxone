package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type FlowDataV2Model struct {
	CdcServices      types.List        `tfsdk:"cdc_services"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	Description      types.String      `tfsdk:"description"`
	DestinationHa    types.Bool        `tfsdk:"destination_ha"`
	Destinations     types.List        `tfsdk:"destinations"`
	Enabled          types.Bool        `tfsdk:"enabled"`
	EtlFilters       types.List        `tfsdk:"etl_filters"`
	FilterExpression types.String      `tfsdk:"filter_expression"`
	Id               types.Int64       `tfsdk:"id"`
	Name             types.String      `tfsdk:"name"`
	ScheduleOption   types.String      `tfsdk:"schedule_option"`
	ScriptSchedule   types.String      `tfsdk:"script_schedule"`
	Source           types.Int64       `tfsdk:"source"`
	SourceDataTypes  types.List        `tfsdk:"source_data_types"`
	SourceLogTypes   types.List        `tfsdk:"source_log_types"`
	Tags             types.Map         `tfsdk:"tags"`
	UpdatedAt        timetypes.RFC3339 `tfsdk:"updated_at"`
}

var FlowDataV2AttrTypes = map[string]attr.Type{
	"cdc_services":      types.ListType{ElemType: types.StringType},
	"created_at":        timetypes.RFC3339Type{},
	"description":       types.StringType,
	"destination_ha":    types.BoolType,
	"destinations":      types.ListType{ElemType: types.Int64Type},
	"enabled":           types.BoolType,
	"etl_filters":       types.ListType{ElemType: types.Int64Type},
	"filter_expression": types.StringType,
	"id":                types.Int64Type,
	"name":              types.StringType,
	"schedule_option":   types.StringType,
	"script_schedule":   types.StringType,
	"source":            types.Int64Type,
	"source_data_types": types.ListType{ElemType: types.StringType},
	"source_log_types":  types.ListType{ElemType: types.ObjectType{AttrTypes: SourceLogTypeAttrTypes}},
	"tags":              types.MapType{ElemType: types.StringType},
	"updated_at":        timetypes.RFC3339Type{},
}

var FlowDataV2ResourceSchemaAttributes = map[string]schema.Attribute{
	"cdc_services": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"destination_ha": schema.BoolAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"destinations": schema.ListAttribute{
		ElementType: types.Int64Type,
		Optional:    true,
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"etl_filters": schema.ListAttribute{
		ElementType: types.Int64Type,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.UseStateForUnknown(),
		},
	},
	"filter_expression": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
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
	"schedule_option": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"script_schedule": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"source": schema.Int64Attribute{
		Optional: true,
	},
	"source_data_types": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"source_log_types": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SourceLogTypeResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.List{
			flowSourceLogTypesPlanModifier{},
		},
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
		PlanModifiers: []planmodifier.String{
			flowUpdatedAtPlanModifier{},
		},
	},
}

func ExpandFlowDataV2(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.FlowDataV2 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FlowDataV2Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FlowDataV2Model) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.FlowDataV2 {
	if m == nil {
		return nil
	}
	to := &cdc.FlowDataV2{
		CdcServices:      flex.ExpandFrameworkListString(ctx, m.CdcServices, diags),
		Description:      flex.ExpandStringPointer(m.Description),
		Enabled:          flex.ExpandBoolPointer(m.Enabled),
		FilterExpression: flex.ExpandStringPointer(m.FilterExpression),
		Name:             flex.ExpandStringPointer(m.Name),
		ScheduleOption:   flex.ExpandStringPointer(m.ScheduleOption),
		ScriptSchedule:   flex.ExpandStringPointer(m.ScriptSchedule),
		Source:           flex.ExpandInt64Pointer(m.Source),
		SourceDataTypes:  flex.ExpandFrameworkListString(ctx, m.SourceDataTypes, diags),
		SourceLogTypes:   flex.ExpandFrameworkListNestedBlock(ctx, m.SourceLogTypes, diags, ExpandSourceLogType),
		Tags:             flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Destinations:     flex.ExpandFrameworkListInt64(ctx, m.Destinations, diags),
		EtlFilters:       flex.ExpandFrameworkListInt64(ctx, m.EtlFilters, diags),
	}
	return to
}

func FlattenFlowDataV2(ctx context.Context, from *cdc.FlowDataV2, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FlowDataV2AttrTypes)
	}
	m := FlowDataV2Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FlowDataV2AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FlowDataV2Model) Flatten(ctx context.Context, from *cdc.FlowDataV2, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FlowDataV2Model{}
	}
	m.CdcServices = flex.FlattenFrameworkListString(ctx, from.CdcServices, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DestinationHa = types.BoolPointerValue(from.DestinationHa)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.FilterExpression = flex.FlattenStringPointer(from.FilterExpression)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ScheduleOption = flex.FlattenStringPointer(from.ScheduleOption)
	m.ScriptSchedule = flex.FlattenStringPointer(from.ScriptSchedule)
	m.Source = flex.FlattenInt64Pointer(from.Source)
	m.SourceDataTypes = flex.FlattenFrameworkListString(ctx, from.SourceDataTypes, diags)
	m.SourceLogTypes = flex.FlattenFrameworkListNestedBlock(ctx, from.SourceLogTypes, SourceLogTypeAttrTypes, diags, FlattenSourceLogType)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Destinations = flex.FlattenFrameworkListInt64(ctx, from.Destinations, diags)
	m.EtlFilters = flex.FlattenFrameworkListInt64(ctx, from.EtlFilters, diags)
}
