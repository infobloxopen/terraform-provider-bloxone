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

type DfpListDfp500ResponseModel struct {
	Error types.Object `tfsdk:"error"`
}

var DfpListDfp500ResponseAttrTypes = map[string]attr.Type{
	"error": types.ObjectType{AttrTypes: DfpListDfp500ResponseErrorAttrTypes},
}

var DfpListDfp500ResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"error": schema.SingleNestedAttribute{
		Attributes: DfpListDfp500ResponseErrorResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandDfpListDfp500Response(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.DfpListDfp500Response {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DfpListDfp500ResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DfpListDfp500ResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpListDfp500Response {
	if m == nil {
		return nil
	}
	to := &dfp.DfpListDfp500Response{
		Error: ExpandDfpListDfp500ResponseError(ctx, m.Error, diags),
	}
	return to
}

func FlattenDfpListDfp500Response(ctx context.Context, from *dfp.DfpListDfp500Response, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DfpListDfp500ResponseAttrTypes)
	}
	m := DfpListDfp500ResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DfpListDfp500ResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DfpListDfp500ResponseModel) Flatten(ctx context.Context, from *dfp.DfpListDfp500Response, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DfpListDfp500ResponseModel{}
	}
	m.Error = FlattenDfpListDfp500ResponseError(ctx, from.Error, diags)
}
