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

type IpamsvcUpdateAddressBlockResponseModel struct {
	Result types.Object `tfsdk:"result"`
}

var IpamsvcUpdateAddressBlockResponseAttrTypes = map[string]attr.Type{
	"result": types.ObjectType{AttrTypes: IpamsvcAddressBlockAttrTypes},
}

var IpamsvcUpdateAddressBlockResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"result": schema.SingleNestedAttribute{
		Attributes: IpamsvcAddressBlockResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcUpdateAddressBlockResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcUpdateAddressBlockResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcUpdateAddressBlockResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcUpdateAddressBlockResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcUpdateAddressBlockResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcUpdateAddressBlockResponse{
		Result: ExpandIpamsvcAddressBlock(ctx, m.Result, diags),
	}
	return to
}

func FlattenIpamsvcUpdateAddressBlockResponse(ctx context.Context, from *ipam.IpamsvcUpdateAddressBlockResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcUpdateAddressBlockResponseAttrTypes)
	}
	m := IpamsvcUpdateAddressBlockResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcUpdateAddressBlockResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcUpdateAddressBlockResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcUpdateAddressBlockResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcUpdateAddressBlockResponseModel{}
	}
	m.Result = FlattenIpamsvcAddressBlock(ctx, from.Result, diags)
}
