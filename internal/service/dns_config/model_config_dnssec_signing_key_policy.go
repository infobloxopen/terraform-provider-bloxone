package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigDNSSECSigningKeyPolicyModel struct {
	Algorithm types.Int64  `tfsdk:"algorithm"`
	Size      types.Int64  `tfsdk:"size"`
	Type      types.String `tfsdk:"type"`
}

var ConfigDNSSECSigningKeyPolicyAttrTypes = map[string]attr.Type{
	"algorithm": types.Int64Type,
	"size":      types.Int64Type,
	"type":      types.StringType,
}

var ConfigDNSSECSigningKeyPolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Algorithm used for the key.  Allowed values: * _5_ - RSASHA1 * _7_ - NSEC3RSASHA1 (RSASHA1-NSEC3-SHA1) * _8_ - RSASHA256 * _10_ - RSASHA512 * _13_ - ECDSAP256SHA256 * _14_ - ECDSAP384SHA384 * _15_ - ED25519 * _16_ - ED448  Defaults to _8_ (RSASHA256).",
	},
	"size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Key size in bits.  Value should be within allowed range for _algorithm_: * _RSASHA1_: 1024..4096 * _NSEC3RSASHA1_: 1024..4096 * _RSASHA256_: 1024..4096 * _RSASHA512_: 1024..4096 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456  Defaults are based on the _algorithm_ and _type_: For KSK:  * _RSASHA1_: 2048 * _NSEC3RSASHA1_: 2048 * _RSASHA256_: 2048 * _RSASHA512_: 2048 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456  For ZSK:  * _RSASHA1_: 1024 * _NSEC3RSASHA1_: 1024 * _RSASHA256_: 1024 * _RSASHA512_: 1024 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Key type.  Allowed values: * _KSK_: Key-Signing Key, used to sign DNSKEY records. * _ZSK_: Zone-Signing Key, used to sign all other records in the zone.",
	},
}

func ExpandConfigDNSSECSigningKeyPolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.DNSSECSigningKeyPolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDNSSECSigningKeyPolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDNSSECSigningKeyPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.DNSSECSigningKeyPolicy {
	if m == nil {
		return nil
	}
	to := &dnsconfig.DNSSECSigningKeyPolicy{
		Algorithm: flex.ExpandInt64Pointer(m.Algorithm),
		Size:      flex.ExpandInt64Pointer(m.Size),
		Type:      flex.ExpandString(m.Type),
	}
	return to
}

func FlattenConfigDNSSECSigningKeyPolicy(ctx context.Context, from *dnsconfig.DNSSECSigningKeyPolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDNSSECSigningKeyPolicyAttrTypes)
	}
	m := ConfigDNSSECSigningKeyPolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDNSSECSigningKeyPolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDNSSECSigningKeyPolicyModel) Flatten(ctx context.Context, from *dnsconfig.DNSSECSigningKeyPolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDNSSECSigningKeyPolicyModel{}
	}
	m.Algorithm = flex.FlattenInt64Pointer(from.Algorithm)
	m.Size = flex.FlattenInt64Pointer(from.Size)
	m.Type = flex.FlattenString(from.Type)
}
