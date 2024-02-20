package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"
)

type AtcdfpDfpCreateOrUpdateResponseModel struct {
	Results types.Object `tfsdk:"results"`
}

var AtcdfpDfpCreateOrUpdateResponseAttrTypes = map[string]attr.Type{
	"results": types.ObjectType{AttrTypes: AtcdfpDfpAttrTypes},
}

var AtcdfpDfpCreateOrUpdateResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.SingleNestedAttribute{
		Attributes: AtcdfpDfpResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandAtcdfpDfpCreateOrUpdateResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpDfpCreateOrUpdateResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDfpCreateOrUpdateResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDfpCreateOrUpdateResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpDfpCreateOrUpdateResponse {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpDfpCreateOrUpdateResponse{
		Results: ExpandAtcdfpDfp(ctx, m.Results, diags),
	}
	return to
}

func FlattenAtcdfpDfpCreateOrUpdateResponse(ctx context.Context, from *dfp.AtcdfpDfpCreateOrUpdateResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDfpCreateOrUpdateResponseAttrTypes)
	}
	m := AtcdfpDfpCreateOrUpdateResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDfpCreateOrUpdateResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDfpCreateOrUpdateResponseModel) Flatten(ctx context.Context, from *dfp.AtcdfpDfpCreateOrUpdateResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDfpCreateOrUpdateResponseModel{}
	}
	m.Results = FlattenAtcdfpDfp(ctx, from.Results, diags)
}
