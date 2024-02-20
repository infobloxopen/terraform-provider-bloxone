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

type DfpReadDfp404ResponseModel struct {
	Error types.Object `tfsdk:"error"`
}

var DfpReadDfp404ResponseAttrTypes = map[string]attr.Type{
	"error": types.ObjectType{AttrTypes: DfpReadDfp404ResponseErrorAttrTypes},
}

var DfpReadDfp404ResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"error": schema.SingleNestedAttribute{
		Attributes: DfpReadDfp404ResponseErrorResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandDfpReadDfp404Response(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.DfpReadDfp404Response {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DfpReadDfp404ResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DfpReadDfp404ResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpReadDfp404Response {
	if m == nil {
		return nil
	}
	to := &dfp.DfpReadDfp404Response{
		Error: ExpandDfpReadDfp404ResponseError(ctx, m.Error, diags),
	}
	return to
}

func FlattenDfpReadDfp404Response(ctx context.Context, from *dfp.DfpReadDfp404Response, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DfpReadDfp404ResponseAttrTypes)
	}
	m := DfpReadDfp404ResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DfpReadDfp404ResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DfpReadDfp404ResponseModel) Flatten(ctx context.Context, from *dfp.DfpReadDfp404Response, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DfpReadDfp404ResponseModel{}
	}
	m.Error = FlattenDfpReadDfp404ResponseError(ctx, from.Error, diags)
}
