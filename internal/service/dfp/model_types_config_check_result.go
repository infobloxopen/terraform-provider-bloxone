package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type TypesConfigCheckResultModel struct {
	AdditionalInfo  types.String `tfsdk:"additional_info"`
	ConfigCheckType types.String `tfsdk:"config_check_type"`
	ResourceUri     types.String `tfsdk:"resource_uri"`
	ResultCode      types.String `tfsdk:"result_code"`
}

var TypesConfigCheckResultAttrTypes = map[string]attr.Type{
	"additional_info":   types.StringType,
	"config_check_type": types.StringType,
	"resource_uri":      types.StringType,
	"result_code":       types.StringType,
}

var TypesConfigCheckResultResourceSchemaAttributes = map[string]schema.Attribute{
	"additional_info": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Provides more info about the potential problem.",
	},
	"config_check_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Service check type.",
	},
	"resource_uri": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "URI of the resource that was checked.",
	},
	"result_code": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The check result.",
	},
}

func ExpandTypesConfigCheckResult(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.TypesConfigCheckResult {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TypesConfigCheckResultModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TypesConfigCheckResultModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.TypesConfigCheckResult {
	if m == nil {
		return nil
	}
	to := &dfp.TypesConfigCheckResult{
		AdditionalInfo:  flex.ExpandStringPointer(m.AdditionalInfo),
		ConfigCheckType: flex.ExpandStringPointer(m.ConfigCheckType),
		ResourceUri:     flex.ExpandStringPointer(m.ResourceUri),
		ResultCode:      flex.ExpandStringPointer(m.ResultCode),
	}
	return to
}

func FlattenTypesConfigCheckResult(ctx context.Context, from *dfp.TypesConfigCheckResult, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TypesConfigCheckResultAttrTypes)
	}
	m := TypesConfigCheckResultModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TypesConfigCheckResultAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TypesConfigCheckResultModel) Flatten(ctx context.Context, from *dfp.TypesConfigCheckResult, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TypesConfigCheckResultModel{}
	}
	m.AdditionalInfo = flex.FlattenStringPointer(from.AdditionalInfo)
	m.ConfigCheckType = flex.FlattenStringPointer(from.ConfigCheckType)
	m.ResourceUri = flex.FlattenStringPointer(from.ResourceUri)
	m.ResultCode = flex.FlattenStringPointer(from.ResultCode)
}
