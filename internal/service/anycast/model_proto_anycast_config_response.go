package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"
)

type ProtoAnycastConfigResponseModel struct {
	Results types.Object `tfsdk:"results"`
}

var ProtoAnycastConfigResponseAttrTypes = map[string]attr.Type{
	"results": types.ObjectType{AttrTypes: ProtoAnycastConfigAttrTypes},
}

var ProtoAnycastConfigResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"results": schema.SingleNestedAttribute{
		Attributes: ProtoAnycastConfigResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandProtoAnycastConfigResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.AnycastConfigResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoAnycastConfigResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoAnycastConfigResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.AnycastConfigResponse {
	if m == nil {
		return nil
	}
	to := &anycast.AnycastConfigResponse{
		Results: ExpandProtoAnycastConfig(ctx, m.Results, diags),
	}
	return to
}

func FlattenProtoAnycastConfigResponse(ctx context.Context, from *anycast.AnycastConfigResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoAnycastConfigResponseAttrTypes)
	}
	m := ProtoAnycastConfigResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoAnycastConfigResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoAnycastConfigResponseModel) Flatten(ctx context.Context, from *anycast.AnycastConfigResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoAnycastConfigResponseModel{}
	}
	m.Results = FlattenProtoAnycastConfig(ctx, from.Results, diags)
}
