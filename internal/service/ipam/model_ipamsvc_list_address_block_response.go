package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcListAddressBlockResponseModel struct {
	Results types.List `tfsdk:"results"`
}

var IpamsvcListAddressBlockResponseAttrTypes = map[string]attr.Type{
	"results": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcAddressBlockAttrTypes}},
}

var IpamsvcListAddressBlockResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcAddressBlockResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "A list of AddressBlock objects.",
	},
}

func ExpandIpamsvcListAddressBlockResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcListAddressBlockResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcListAddressBlockResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcListAddressBlockResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcListAddressBlockResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcListAddressBlockResponse{
		Results: flex.ExpandFrameworkListNestedBlock(ctx, m.Results, diags, ExpandIpamsvcAddressBlock),
	}
	return to
}

func FlattenIpamsvcListAddressBlockResponse(ctx context.Context, from *ipam.IpamsvcListAddressBlockResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcListAddressBlockResponseAttrTypes)
	}
	m := IpamsvcListAddressBlockResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcListAddressBlockResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcListAddressBlockResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcListAddressBlockResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcListAddressBlockResponseModel{}
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from.Results, IpamsvcAddressBlockAttrTypes, diags, FlattenIpamsvcAddressBlock)
}
