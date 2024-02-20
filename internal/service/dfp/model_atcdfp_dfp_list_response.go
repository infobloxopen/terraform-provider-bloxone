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

type AtcdfpDfpListResponseModel struct {
	Results types.List `tfsdk:"results"`
}

var AtcdfpDfpListResponseAttrTypes = map[string]attr.Type{
	"results": types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpDfpAttrTypes}},
}

var AtcdfpDfpListResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcdfpDfpResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DNS Forwarding Proxy objects.",
	},
}

func ExpandAtcdfpDfpListResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpDfpListResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDfpListResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDfpListResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpDfpListResponse {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpDfpListResponse{
		Results: flex.ExpandFrameworkListNestedBlock(ctx, m.Results, diags, ExpandAtcdfpDfp),
	}
	return to
}

func FlattenAtcdfpDfpListResponse(ctx context.Context, from *dfp.AtcdfpDfpListResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDfpListResponseAttrTypes)
	}
	m := AtcdfpDfpListResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDfpListResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDfpListResponseModel) Flatten(ctx context.Context, from *dfp.AtcdfpDfpListResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDfpListResponseModel{}
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from.Results, AtcdfpDfpAttrTypes, diags, FlattenAtcdfpDfp)
}
