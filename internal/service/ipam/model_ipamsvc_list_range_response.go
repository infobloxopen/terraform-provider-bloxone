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

type IpamsvcListRangeResponseModel struct {
	Results types.List `tfsdk:"results"`
}

var IpamsvcListRangeResponseAttrTypes = map[string]attr.Type{
	"results": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcRangeAttrTypes}},
}

var IpamsvcListRangeResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcRangeResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of Range objects.",
	},
}

func ExpandIpamsvcListRangeResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcListRangeResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcListRangeResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcListRangeResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcListRangeResponse {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcListRangeResponse{
		Results: flex.ExpandFrameworkListNestedBlock(ctx, m.Results, diags, ExpandIpamsvcRange),
	}
	return to
}

func FlattenIpamsvcListRangeResponse(ctx context.Context, from *ipam.IpamsvcListRangeResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcListRangeResponseAttrTypes)
	}
	m := IpamsvcListRangeResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcListRangeResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcListRangeResponseModel) Flatten(ctx context.Context, from *ipam.IpamsvcListRangeResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcListRangeResponseModel{}
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from.Results, IpamsvcRangeAttrTypes, diags, FlattenIpamsvcRange)
}
