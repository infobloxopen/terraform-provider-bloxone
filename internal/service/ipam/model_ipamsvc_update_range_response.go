package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcUpdateRangeResponseModel struct {
	Result types.Object `tfsdk:"result"`
}

var IpamsvcUpdateRangeResponseAttrTypes = map[string]attr.Type{
	"result": types.ObjectType{AttrTypes: IpamsvcRangeAttrTypes},
}

var IpamsvcUpdateRangeResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"result": schema.SingleNestedAttribute{
		Attributes: IpamsvcRangeResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcUpdateRangeResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcUpdateRangeResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcUpdateRangeResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcUpdateRangeResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcUpdateRangeResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcUpdateRangeResponse{
		Result: ExpandIpamsvcRange(ctx, m.Result, diags),
	}
	return to
}

func FlattenIpamsvcUpdateRangeResponse(ctx context.Context, from *ipam.IpamsvcUpdateRangeResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUpdateRangeResponseAttrTypes)
	}
	m := IpamsvcUpdateRangeResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUpdateRangeResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUpdateRangeResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcUpdateRangeResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUpdateRangeResponseModel{}
	}
	m.Result = FlattenIpamsvcRange(ctx, from.Result, diags)
}
