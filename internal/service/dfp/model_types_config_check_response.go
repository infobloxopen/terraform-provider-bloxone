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

type TypesConfigCheckResponseModel struct {
	Results types.List `tfsdk:"results"`
}

var TypesConfigCheckResponseAttrTypes = map[string]attr.Type{
	"results": types.ListType{ElemType: types.ObjectType{AttrTypes: TypesConfigCheckResultAttrTypes}},
}

var TypesConfigCheckResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TypesConfigCheckResultResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of check result.",
	},
}

func ExpandTypesConfigCheckResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.TypesConfigCheckResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TypesConfigCheckResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TypesConfigCheckResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.TypesConfigCheckResponse {
	if m == nil {
		return nil
	}
	to := &dfp.TypesConfigCheckResponse{
		Results: flex.ExpandFrameworkListNestedBlock(ctx, m.Results, diags, ExpandTypesConfigCheckResult),
	}
	return to
}

func FlattenTypesConfigCheckResponse(ctx context.Context, from *dfp.TypesConfigCheckResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TypesConfigCheckResponseAttrTypes)
	}
	m := TypesConfigCheckResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TypesConfigCheckResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TypesConfigCheckResponseModel) Flatten(ctx context.Context, from *dfp.TypesConfigCheckResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TypesConfigCheckResponseModel{}
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from.Results, TypesConfigCheckResultAttrTypes, diags, FlattenTypesConfigCheckResult)
}
