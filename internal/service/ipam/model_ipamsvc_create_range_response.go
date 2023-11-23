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

type IpamsvcCreateRangeResponseModel struct {
	Result types.Object `tfsdk:"result"`
}

var IpamsvcCreateRangeResponseAttrTypes = map[string]attr.Type{
	"result": types.ObjectType{AttrTypes: IpamsvcRangeAttrTypes},
}

var IpamsvcCreateRangeResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"result": schema.SingleNestedAttribute{
		Attributes: IpamsvcRangeResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcCreateRangeResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcCreateRangeResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcCreateRangeResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcCreateRangeResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcCreateRangeResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcCreateRangeResponse{
		Result: ExpandIpamsvcRange(ctx, m.Result, diags),
	}
	return to
}

func FlattenIpamsvcCreateRangeResponse(ctx context.Context, from *ipam.IpamsvcCreateRangeResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcCreateRangeResponseAttrTypes)
	}
	m := IpamsvcCreateRangeResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcCreateRangeResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcCreateRangeResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcCreateRangeResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcCreateRangeResponseModel{}
	}
	m.Result = FlattenIpamsvcRange(ctx, from.Result, diags)
}
