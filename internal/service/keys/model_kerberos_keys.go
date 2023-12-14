package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/keys"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type KerberosKeysModel struct {
	Items types.List `tfsdk:"items"`
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

func (m *KerberosKeysModel) Flatten(ctx context.Context, from *keys.KerberosKeys, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KerberosKeysModel{}
	}
	m.Items = flex.FlattenFrameworkListNestedBlock(ctx, from.Items, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
}
