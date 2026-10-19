package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

// flowUpdatedAtPlanModifier preserves updated_at from prior state, like
// stringplanmodifier.UseStateForUnknown(), except when the user is actually
// requesting a change to another attribute on the flow. In that case the
// value is left unknown so the API's freshly returned timestamp is
// accepted, instead of being compared against the (now stale) prior value.
//
// Plain UseStateForUnknown() doesn't work for this field: a genuine update
// always changes updated_at server-side, so locking the plan to the old
// value causes "provider produced inconsistent result after apply" on
// every update. But dropping the modifier entirely isn't safe either:
// source_log_types (the only ListNestedAttribute on this resource) going
// unprotected at the same time causes Terraform to treat even a pure no-op
// refresh as a pending change.
//
// This compares against Config rather than Plan: several sibling
// attributes (etl_filters, filter_expression, ...) are themselves
// Optional+Computed with their own UseStateForUnknown, so their planned
// value can still be Unknown depending on plan-modifier execution order.
// Config values never have that ambiguity - they're either the literal
// value the user wrote, or null.
type flowUpdatedAtPlanModifier struct{}

func (m flowUpdatedAtPlanModifier) Description(_ context.Context) string {
	return "Preserves updated_at from prior state unless the user is changing another flow attribute."
}

func (m flowUpdatedAtPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m flowUpdatedAtPlanModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Resource is being created; nothing to preserve yet.
	if req.State.Raw.IsNull() {
		return
	}

	var stateData, configData FlowDataV2Model
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !flowConfigRequestsChange(stateData, configData) {
		resp.PlanValue = req.StateValue
	}
}

// flowConfigRequestsChange reports whether the user's config explicitly
// sets any flow attribute to a value different from prior state. Attributes
// the user left unconfigured (null in config) can't represent a
// user-requested change, regardless of what their planned value resolves
// to.
func flowConfigRequestsChange(state, config FlowDataV2Model) bool {
	return configDiffers(config.CdcServices, state.CdcServices) ||
		configDiffers(config.Description, state.Description) ||
		configDiffers(config.Destinations, state.Destinations) ||
		configDiffers(config.Enabled, state.Enabled) ||
		configDiffers(config.EtlFilters, state.EtlFilters) ||
		configDiffers(config.FilterExpression, state.FilterExpression) ||
		configDiffers(config.Name, state.Name) ||
		configDiffers(config.ScheduleOption, state.ScheduleOption) ||
		configDiffers(config.ScriptSchedule, state.ScriptSchedule) ||
		configDiffers(config.Source, state.Source) ||
		configDiffers(config.SourceDataTypes, state.SourceDataTypes) ||
		configDiffers(config.Tags, state.Tags)
}

func configDiffers(configVal, stateVal attr.Value) bool {
	return !configVal.IsNull() && !configVal.Equal(stateVal)
}

// flowSourceLogTypesPlanModifier preserves source_log_types from prior
// state, like listplanmodifier.UseStateForUnknown(), except when
// source_data_types is actually changing. source_log_types is derived
// server-side from source_data_types (confirmed: changing source_data_types
// changes the returned source_log_types too), so carrying forward a stale
// value while source_data_types changes causes "provider produced
// inconsistent result after apply".
type flowSourceLogTypesPlanModifier struct{}

func (m flowSourceLogTypesPlanModifier) Description(_ context.Context) string {
	return "Preserves source_log_types from prior state unless source_data_types is changing."
}

func (m flowSourceLogTypesPlanModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m flowSourceLogTypesPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	// Resource is being created; nothing to preserve yet.
	if req.State.Raw.IsNull() {
		return
	}

	var stateData, configData FlowDataV2Model
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &configData)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !configDiffers(configData.SourceDataTypes, stateData.SourceDataTypes) {
		resp.PlanValue = req.StateValue
	}
}
