package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/keys"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type KerberosKeysModel struct {
	Items types.List `tfsdk:"items"`
}

var KerberosKeysAttrTypes = map[string]attr.Type{
	"items": types.ListType{ElemType: types.ObjectType{AttrTypes: KerberosKeyAttrTypes}},
}

var KerberosKeysResourceSchemaAttributes = map[string]schema.Attribute{
	"items": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KerberosKeyResourceSchemaAttributes,
		},
		Optional: true,
	},
}

func ExpandKerberosKeys(ctx context.Context, o types.Object, diags *diag.Diagnostics) *keys.KerberosKeys {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KerberosKeysModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KerberosKeysModel) Expand(ctx context.Context, diags *diag.Diagnostics) *keys.KerberosKeys {
	if m == nil {
		return nil
	}
	to := &keys.KerberosKeys{
		Items: flex.ExpandFrameworkListNestedBlock(ctx, m.Items, diags, ExpandKerberosKey),
	}
	return to
}

func FlattenKerberosKeys(ctx context.Context, from *keys.KerberosKeys, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KerberosKeysAttrTypes)
	}
	m := KerberosKeysModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KerberosKeysAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KerberosKeysModel) Flatten(ctx context.Context, from *keys.KerberosKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KerberosKeysModel{}
	}
	m.Items = flex.FlattenFrameworkListNestedBlock(ctx, from.Items, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
}
