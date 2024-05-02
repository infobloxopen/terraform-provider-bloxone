package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoAnycastVersionModel struct {
	AccountId types.Int64  `tfsdk:"account_id"`
	Version   types.String `tfsdk:"version"`
}

var ProtoAnycastVersionAttrTypes = map[string]attr.Type{
	"account_id": types.Int64Type,
	"version":    types.StringType,
}

var ProtoAnycastVersionResourceSchemaAttributes = map[string]schema.Attribute{
	"account_id": schema.Int64Attribute{
		Optional: true,
	},
	"version": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandProtoAnycastVersion(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.AnycastVersion {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoAnycastVersionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoAnycastVersionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.AnycastVersion {
	if m == nil {
		return nil
	}
	to := &anycast.AnycastVersion{
		AccountId: flex.ExpandInt64Pointer(m.AccountId),
		Version:   flex.ExpandStringPointer(m.Version),
	}
	return to
}

func FlattenProtoAnycastVersion(ctx context.Context, from *anycast.AnycastVersion, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoAnycastVersionAttrTypes)
	}
	m := ProtoAnycastVersionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoAnycastVersionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoAnycastVersionModel) Flatten(ctx context.Context, from *anycast.AnycastVersion, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoAnycastVersionModel{}
	}
	m.AccountId = flex.FlattenInt64Pointer(from.AccountId)
	m.Version = flex.FlattenStringPointer(from.Version)
}
