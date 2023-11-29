package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
)

type ConfigHostInheritanceModel struct {
	KerberosKeys types.Object `tfsdk:"kerberos_keys"`
}

var ConfigHostInheritanceAttrTypes = map[string]attr.Type{
	"kerberos_keys": types.ObjectType{AttrTypes: ConfigInheritedKerberosKeysAttrTypes},
}

var ConfigHostInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"kerberos_keys": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedKerberosKeysResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandConfigHostInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigHostInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigHostInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigHostInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigHostInheritance {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigHostInheritance{
		KerberosKeys: ExpandConfigInheritedKerberosKeys(ctx, m.KerberosKeys, diags),
	}
	return to
}

func FlattenConfigHostInheritance(ctx context.Context, from *dns_config.ConfigHostInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigHostInheritanceAttrTypes)
	}
	m := ConfigHostInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigHostInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigHostInheritanceModel) Flatten(ctx context.Context, from *dns_config.ConfigHostInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigHostInheritanceModel{}
	}
	m.KerberosKeys = FlattenConfigInheritedKerberosKeys(ctx, from.KerberosKeys, diags)
}
