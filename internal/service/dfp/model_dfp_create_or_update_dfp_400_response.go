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

type DfpCreateOrUpdateDfp400ResponseModel struct {
	Error types.Object `tfsdk:"error"`
}

var DfpCreateOrUpdateDfp400ResponseAttrTypes = map[string]attr.Type{
	"error": types.ObjectType{AttrTypes: DfpCreateOrUpdateDfp400ResponseErrorAttrTypes},
}

var DfpCreateOrUpdateDfp400ResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"error": schema.SingleNestedAttribute{
		Attributes: DfpCreateOrUpdateDfp400ResponseErrorResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandDfpCreateOrUpdateDfp400Response(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.DfpCreateOrUpdateDfp400Response {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DfpCreateOrUpdateDfp400ResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DfpCreateOrUpdateDfp400ResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpCreateOrUpdateDfp400Response {
	if m == nil {
		return nil
	}
	to := &dfp.DfpCreateOrUpdateDfp400Response{
		Error: ExpandDfpCreateOrUpdateDfp400ResponseError(ctx, m.Error, diags),
	}
	return to
}

func FlattenDfpCreateOrUpdateDfp400Response(ctx context.Context, from *dfp.DfpCreateOrUpdateDfp400Response, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DfpCreateOrUpdateDfp400ResponseAttrTypes)
	}
	m := DfpCreateOrUpdateDfp400ResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DfpCreateOrUpdateDfp400ResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DfpCreateOrUpdateDfp400ResponseModel) Flatten(ctx context.Context, from *dfp.DfpCreateOrUpdateDfp400Response, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DfpCreateOrUpdateDfp400ResponseModel{}
	}
	m.Error = FlattenDfpCreateOrUpdateDfp400ResponseError(ctx, from.Error, diags)
}
