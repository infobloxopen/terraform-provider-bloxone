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

type IpamsvcReadAddressBlockResponseModel struct {
	Result types.Object `tfsdk:"result"`
}

var IpamsvcReadAddressBlockResponseAttrTypes = map[string]attr.Type{
	"result": types.ObjectType{AttrTypes: IpamsvcAddressBlockAttrTypes},
}

var IpamsvcReadAddressBlockResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"result": schema.SingleNestedAttribute{
		Attributes: IpamsvcAddressBlockResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcReadAddressBlockResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcReadAddressBlockResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcReadAddressBlockResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcReadAddressBlockResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcReadAddressBlockResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcReadAddressBlockResponse{
		Result: ExpandIpamsvcAddressBlock(ctx, m.Result, diags),
	}
	return to
}

func FlattenIpamsvcReadAddressBlockResponse(ctx context.Context, from *ipam.IpamsvcReadAddressBlockResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcReadAddressBlockResponseAttrTypes)
	}
	m := IpamsvcReadAddressBlockResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcReadAddressBlockResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcReadAddressBlockResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcReadAddressBlockResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcReadAddressBlockResponseModel{}
	}
	m.Result = FlattenIpamsvcAddressBlock(ctx, from.Result, diags)
}
