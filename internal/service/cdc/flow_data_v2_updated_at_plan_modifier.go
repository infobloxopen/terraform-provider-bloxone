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
// sets any flow attribute to a value different from prior state.
//
// Two different comparisons are needed depending on each attribute's
// schema, because "config is null" means something different for each:
//   - Optional+Computed attributes (description, enabled, etl_filters,
//     filter_expression, schedule_option, script_schedule): a null config
//     just means "let the provider/backend decide", not "clear this value" -
//     the attribute's planned/state value can legitimately be non-null even
//     though config never set it. For these, configDiffers correctly
//     ignores a null config (it can't represent a user-requested change).
//   - Plain Optional (not Computed) attributes (cdc_services, destinations,
//     name, source, source_data_types, tags): the planned value always
//     equals the literal config value, null or not - there's no separate
//     "computed default" for Terraform to fall back on. So a transition
//     from a previously-set value to a null config (the user deliberately
//     removing/clearing the attribute) *is* a real, user-requested change,
//     and must be detected the same as any other diff. Using configDiffers
//     here would silently miss that case - the field change would still
//     force a real Update() (since it's not Computed, Terraform plans the
//     diff on its own either way), but flowConfigRequestsChange would
//     wrongly report "no change", causing updated_at to stay pinned to the
//     stale prior value while the backend genuinely bumps it - the exact
//     "provider produced inconsistent result after apply" shape this
//     modifier exists to prevent.
func flowConfigRequestsChange(state, config FlowDataV2Model) bool {
	return optionalDiffers(config.CdcServices, state.CdcServices) ||
		configDiffers(config.Description, state.Description) ||
		optionalDiffers(config.Destinations, state.Destinations) ||
		configDiffers(config.Enabled, state.Enabled) ||
		configDiffers(config.EtlFilters, state.EtlFilters) ||
		configDiffers(config.FilterExpression, state.FilterExpression) ||
		optionalDiffers(config.Name, state.Name) ||
		configDiffers(config.ScheduleOption, state.ScheduleOption) ||
		configDiffers(config.ScriptSchedule, state.ScriptSchedule) ||
		optionalDiffers(config.Source, state.Source) ||
		optionalDiffers(config.SourceDataTypes, state.SourceDataTypes) ||
		optionalDiffers(config.Tags, state.Tags)
}

// configDiffers reports a change for an Optional+Computed attribute: a null
// config can never represent a user-requested change, since the provider
// may legitimately plan a non-null value for it on its own.
func configDiffers(configVal, stateVal attr.Value) bool {
	return !configVal.IsNull() && !configVal.Equal(stateVal)
}

// optionalDiffers reports a change for a plain Optional (not Computed)
// attribute: its planned value always equals the literal config value, so
// even a transition to/from null (the user adding or clearing the
// attribute) is a real, user-requested change.
func optionalDiffers(configVal, stateVal attr.Value) bool {
	return !configVal.Equal(stateVal)
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

	if !optionalDiffers(configData.SourceDataTypes, stateData.SourceDataTypes) {
		resp.PlanValue = req.StateValue
	}
}
