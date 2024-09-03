package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
)

type ReadFederatedRealmResponseModel struct {
	Result types.Object `tfsdk:"result"`
}

var ReadFederatedRealmResponseAttrTypes = map[string]attr.Type{
	"result": types.ObjectType{AttrTypes: FederatedRealmAttrTypes},
}

var ReadFederatedRealmResponseResourceSchemaAttributes = map[string]schema.Attribute{
	"result": schema.SingleNestedAttribute{
		Attributes:          FederatedRealmResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The FederatedRealm object.",
	},
}

func ExpandReadFederatedRealmResponse(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.ReadFederatedRealmResponse {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ReadFederatedRealmResponseModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ReadFederatedRealmResponseModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.ReadFederatedRealmResponse {
	if m == nil {
		return nil
	}
	to := &ipamfederation.ReadFederatedRealmResponse{
		Result: ExpandFederatedRealm(ctx, m.Result, diags),
	}
	return to
}

func FlattenReadFederatedRealmResponse(ctx context.Context, from *ipamfederation.ReadFederatedRealmResponse, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ReadFederatedRealmResponseAttrTypes)
	}
	m := ReadFederatedRealmResponseModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ReadFederatedRealmResponseAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ReadFederatedRealmResponseModel) Flatten(ctx context.Context, from *ipamfederation.ReadFederatedRealmResponse, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ReadFederatedRealmResponseModel{}
	}
	m.Result = FlattenFederatedRealm(ctx, from.Result, diags)
}
