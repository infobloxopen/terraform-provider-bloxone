package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigDNSSECKeyModel struct {
	Algorithm         types.Int64       `tfsdk:"algorithm"`
	KeyId             types.Int64       `tfsdk:"key_id"`
	NextRolloverEvent timetypes.RFC3339 `tfsdk:"next_rollover_event"`
	PublicKey         types.String      `tfsdk:"public_key"`
	Size              types.Int64       `tfsdk:"size"`
	Type              types.String      `tfsdk:"type"`
}

var ConfigDNSSECKeyAttrTypes = map[string]attr.Type{
	"algorithm":           types.Int64Type,
	"key_id":              types.Int64Type,
	"next_rollover_event": timetypes.RFC3339Type{},
	"public_key":          types.StringType,
	"size":                types.Int64Type,
	"type":                types.StringType,
}

var ConfigDNSSECKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Algorithm used for the key.",
	},
	"key_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Key ID (also known as Key Tag).",
	},
	"next_rollover_event": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Next Rollover Event Time.",
	},
	"public_key": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Public key in Base64 format.",
	},
	"size": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "Key size in bits.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Key type.  Allowed values: * _KSK_: Key-Signing Key, used to sign DNSKEY records. * _ZSK_: Zone-Signing Key, used to sign all other records in the zone.",
	},
}

func ExpandConfigDNSSECKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.DNSSECKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDNSSECKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDNSSECKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.DNSSECKey {
	if m == nil {
		return nil
	}
	to := &dnsconfig.DNSSECKey{}
	return to
}

func FlattenConfigDNSSECKey(ctx context.Context, from *dnsconfig.DNSSECKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDNSSECKeyAttrTypes)
	}
	m := ConfigDNSSECKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDNSSECKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDNSSECKeyModel) Flatten(ctx context.Context, from *dnsconfig.DNSSECKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDNSSECKeyModel{}
	}
	m.Algorithm = flex.FlattenInt64Pointer(from.Algorithm)
	m.KeyId = flex.FlattenInt64Pointer(from.KeyId)
	m.NextRolloverEvent = timetypes.NewRFC3339TimePointerValue(from.NextRolloverEvent)
	m.PublicKey = flex.FlattenStringPointer(from.PublicKey)
	m.Size = flex.FlattenInt64Pointer(from.Size)
	m.Type = flex.FlattenStringPointer(from.Type)
}
