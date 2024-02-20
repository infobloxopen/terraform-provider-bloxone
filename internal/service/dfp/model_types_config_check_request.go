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

type TypesConfigCheckRequestModel struct {
	AccountId       types.Int64  `tfsdk:"account_id"`
	ConfigCheckType types.String `tfsdk:"config_check_type"`
	Notify          types.Bool   `tfsdk:"notify"`
}

var TypesConfigCheckRequestAttrTypes = map[string]attr.Type{
	"account_id":        types.Int64Type,
	"config_check_type": types.StringType,
	"notify":            types.BoolType,
}

var TypesConfigCheckRequestResourceSchemaAttributes = map[string]schema.Attribute{
	"account_id": schema.Int64Attribute{
		Optional: true,
	},
	"config_check_type": schema.StringAttribute{
		Optional: true,
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Caller sets to false if wants to suppress notification.",
	},
}

func ExpandTypesConfigCheckRequest(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.TypesConfigCheckRequest {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m TypesConfigCheckRequestModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *TypesConfigCheckRequestModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.TypesConfigCheckRequest {
	if m == nil {
		return nil
	}
	to := &dfp.TypesConfigCheckRequest{
		AccountId:       flex.ExpandInt64Pointer(m.AccountId),
		ConfigCheckType: flex.ExpandStringPointer(m.ConfigCheckType),
		Notify:          flex.ExpandBoolPointer(m.Notify),
	}
	return to
}

func FlattenTypesConfigCheckRequest(ctx context.Context, from *dfp.TypesConfigCheckRequest, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(TypesConfigCheckRequestAttrTypes)
	}
	m := TypesConfigCheckRequestModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, TypesConfigCheckRequestAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *TypesConfigCheckRequestModel) Flatten(ctx context.Context, from *dfp.TypesConfigCheckRequest, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = TypesConfigCheckRequestModel{}
	}
	m.AccountId = flex.FlattenInt64Pointer(from.AccountId)
	m.ConfigCheckType = flex.FlattenStringPointer(from.ConfigCheckType)
	m.Notify = types.BoolPointerValue(from.Notify)
}
