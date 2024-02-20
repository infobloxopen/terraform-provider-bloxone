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

type DfpListDfp500ResponseErrorModel struct {
	Code    types.String `tfsdk:"code"`
	Message types.String `tfsdk:"message"`
	Status  types.String `tfsdk:"status"`
}

var DfpListDfp500ResponseErrorAttrTypes = map[string]attr.Type{
	"code":    types.StringType,
	"message": types.StringType,
	"status":  types.StringType,
}

var DfpListDfp500ResponseErrorResourceSchemaAttributes = map[string]schema.Attribute{
	"code": schema.StringAttribute{
		Optional: true,
	},
	"message": schema.StringAttribute{
		Optional: true,
	},
	"status": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandDfpListDfp500ResponseError(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.DfpListDfp500ResponseError {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DfpListDfp500ResponseErrorModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DfpListDfp500ResponseErrorModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpListDfp500ResponseError {
	if m == nil {
		return nil
	}
	to := &dfp.DfpListDfp500ResponseError{
		Code:    flex.ExpandStringPointer(m.Code),
		Message: flex.ExpandStringPointer(m.Message),
		Status:  flex.ExpandStringPointer(m.Status),
	}
	return to
}

func FlattenDfpListDfp500ResponseError(ctx context.Context, from *dfp.DfpListDfp500ResponseError, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DfpListDfp500ResponseErrorAttrTypes)
	}
	m := DfpListDfp500ResponseErrorModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DfpListDfp500ResponseErrorAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DfpListDfp500ResponseErrorModel) Flatten(ctx context.Context, from *dfp.DfpListDfp500ResponseError, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DfpListDfp500ResponseErrorModel{}
	}
	m.Code = flex.FlattenStringPointer(from.Code)
	m.Message = flex.FlattenStringPointer(from.Message)
	m.Status = flex.FlattenStringPointer(from.Status)
}
