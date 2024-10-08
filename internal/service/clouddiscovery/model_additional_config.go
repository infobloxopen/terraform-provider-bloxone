package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AdditionalConfigModel struct {
	ExcludedAccounts      types.List   `tfsdk:"excluded_accounts"`
	ForwardZoneEnabled    types.Bool   `tfsdk:"forward_zone_enabled"`
	InternalRangesEnabled types.Bool   `tfsdk:"internal_ranges_enabled"`
	ObjectType            types.Object `tfsdk:"object_type"`
}

var AdditionalConfigAttrTypes = map[string]attr.Type{
	"excluded_accounts":       types.ListType{ElemType: types.StringType},
	"forward_zone_enabled":    types.BoolType,
	"internal_ranges_enabled": types.BoolType,
	"object_type":             types.ObjectType{AttrTypes: ObjectTypeAttrTypes},
}

var AdditionalConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded_accounts": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "List of account IDs to exclude from discovery.",
	},
	"forward_zone_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable/Disable forward zone discovery.",
	},
	"internal_ranges_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enable/Disable internal ranges discovery.",
	},
	"object_type": schema.SingleNestedAttribute{
		Attributes:          ObjectTypeResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Object type to discover.",
	},
}

func ExpandAdditionalConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.AdditionalConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdditionalConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdditionalConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.AdditionalConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.AdditionalConfig{
		ExcludedAccounts:      flex.ExpandFrameworkListString(ctx, m.ExcludedAccounts, diags),
		ForwardZoneEnabled:    flex.ExpandBoolPointer(m.ForwardZoneEnabled),
		InternalRangesEnabled: flex.ExpandBoolPointer(m.InternalRangesEnabled),
		ObjectType:            ExpandObjectType(ctx, m.ObjectType, diags),
	}
	return to
}

func FlattenAdditionalConfig(ctx context.Context, from *clouddiscovery.AdditionalConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdditionalConfigAttrTypes)
	}
	m := AdditionalConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdditionalConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdditionalConfigModel) Flatten(ctx context.Context, from *clouddiscovery.AdditionalConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdditionalConfigModel{}
	}
	m.ExcludedAccounts = flex.FlattenFrameworkListString(ctx, from.ExcludedAccounts, diags)
	m.ForwardZoneEnabled = flex.FlattenBoolPointerFalseAsNull(from.ForwardZoneEnabled)
	m.InternalRangesEnabled = flex.FlattenBoolPointerFalseAsNull(from.InternalRangesEnabled)
	m.ObjectType = FlattenObjectType(ctx, from.ObjectType, diags)
}
