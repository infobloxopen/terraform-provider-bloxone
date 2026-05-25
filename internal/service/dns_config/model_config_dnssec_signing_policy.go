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

type ConfigDNSSECSigningPolicyModel struct {
	Keys                        types.List   `tfsdk:"keys"`
	KskAutomaticRolloverEnabled types.Bool   `tfsdk:"ksk_automatic_rollover_enabled"`
	KskNotificationEventTrigger types.String `tfsdk:"ksk_notification_event_trigger"`
	KskRolloverInterval         types.Int64  `tfsdk:"ksk_rollover_interval"`
	Nsec3Iterations             types.Int64  `tfsdk:"nsec3_iterations"`
	Nsec3SaltLength             types.Int64  `tfsdk:"nsec3_salt_length"`
	NsecType                    types.String `tfsdk:"nsec_type"`
	ZskRolloverInterval         types.Int64  `tfsdk:"zsk_rollover_interval"`
	ZskSignatureValidity        types.Int64  `tfsdk:"zsk_signature_validity"`
}

var ConfigDNSSECSigningPolicyAttrTypes = map[string]attr.Type{
	"keys":                           types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigDNSSECSigningKeyPolicyAttrTypes}},
	"ksk_automatic_rollover_enabled": types.BoolType,
	"ksk_notification_event_trigger": types.StringType,
	"ksk_rollover_interval":          types.Int64Type,
	"nsec3_iterations":               types.Int64Type,
	"nsec3_salt_length":              types.Int64Type,
	"nsec_type":                      types.StringType,
	"zsk_rollover_interval":          types.Int64Type,
	"zsk_signature_validity":         types.Int64Type,
}

var ConfigDNSSECSigningPolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigDNSSECSigningKeyPolicyResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Key settings. This defines configuration for DNSSEC keys. combination of both Key-Signing Keys (KSK) and Zone-Signing Keys (ZSK).  Defaults to empty.",
	},
	"ksk_automatic_rollover_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Flag indicating if KSK rollover should be automatic.  Defaults to _false_.",
	},
	"ksk_notification_event_trigger": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Option controls when notifications are sent for KSK rollover events.  Valid values are: * _NO_EVENTS_ - no notifications are sent * _ALL_EVENTS_ - any time KSK is rolled over, a notification is sent * _MANUAL_DS_UPDATE_EVENTS_ - a notification is sent only when DS record needs to be updated manually  Defaults to _NO_EVENTS_",
	},
	"ksk_rollover_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "KSK rollover interval in seconds.  Used to determine how often the Key-Signing Keys should be rotated. Examples: 31536000 (1 year), 7776000 (90 days), 2592000 (30 days)  Unsigned integer, min 0.  Defaults to 31536000 (1 year).",
	},
	"nsec3_iterations": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Number of additional hash iterations to perform. Increasing this value slows down both authoritative server when signing and recursive servers when verifying, but also slows down attacker's dictionary attacks.  IMPORTANT: Changing the settings for the NSEC3 number of iterations is not recommended.  Unsigned integer, min 0 max 65535.  Defaults to _0_.",
	},
	"nsec3_salt_length": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Minimum length for NSEC3 salt in octets. Used to add entropy to the hash function to defend against pre-calculated attacks.  Unsigned integer, min 0 max 65535.  Defaults to _0_.",
	},
	"nsec_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "DNSSEC resource record type for nonexistent proof. This controls which type will be used to provide proof of nonexistence.  Allowed values: * _NSEC_ * _NSEC3_  Defaults to _NSEC_",
	},
	"zsk_rollover_interval": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "ZSK rollover interval in seconds.  Used to determine how often the Zone-Signing Keys should be rotated. Examples: 2592000 (30 days), 1209600 (14 days)  Unsigned integer, min 0.  Defaults to 2592000 (30 days).",
	},
	"zsk_signature_validity": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "ZSK signature validity period in seconds.  Determines how long DNSSEC signatures remain valid. Examples: 1209600 (14 days), 604800 (7 days)  Unsigned integer, min 0.  Defaults to 1209600 (14 days).",
	},
}

func ExpandConfigDNSSECSigningPolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.DNSSECSigningPolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDNSSECSigningPolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDNSSECSigningPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.DNSSECSigningPolicy {
	if m == nil {
		return nil
	}
	to := &dnsconfig.DNSSECSigningPolicy{
		Keys:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Keys, diags, ExpandConfigDNSSECSigningKeyPolicy),
		KskAutomaticRolloverEnabled: flex.ExpandBoolPointer(m.KskAutomaticRolloverEnabled),
		KskNotificationEventTrigger: flex.ExpandStringPointer(m.KskNotificationEventTrigger),
		KskRolloverInterval:         flex.ExpandInt64Pointer(m.KskRolloverInterval),
		Nsec3Iterations:             flex.ExpandInt64Pointer(m.Nsec3Iterations),
		Nsec3SaltLength:             flex.ExpandInt64Pointer(m.Nsec3SaltLength),
		NsecType:                    flex.ExpandStringPointer(m.NsecType),
		ZskRolloverInterval:         flex.ExpandInt64Pointer(m.ZskRolloverInterval),
		ZskSignatureValidity:        flex.ExpandInt64Pointer(m.ZskSignatureValidity),
	}
	return to
}

func FlattenConfigDNSSECSigningPolicy(ctx context.Context, from *dnsconfig.DNSSECSigningPolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDNSSECSigningPolicyAttrTypes)
	}
	m := ConfigDNSSECSigningPolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDNSSECSigningPolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDNSSECSigningPolicyModel) Flatten(ctx context.Context, from *dnsconfig.DNSSECSigningPolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDNSSECSigningPolicyModel{}
	}
	m.Keys = flex.FlattenFrameworkListNestedBlock(ctx, from.Keys, ConfigDNSSECSigningKeyPolicyAttrTypes, diags, FlattenConfigDNSSECSigningKeyPolicy)
	m.KskAutomaticRolloverEnabled = types.BoolPointerValue(from.KskAutomaticRolloverEnabled)
	m.KskNotificationEventTrigger = flex.FlattenStringPointer(from.KskNotificationEventTrigger)
	m.KskRolloverInterval = flex.FlattenInt64Pointer(from.KskRolloverInterval)
	m.Nsec3Iterations = flex.FlattenInt64Pointer(from.Nsec3Iterations)
	m.Nsec3SaltLength = flex.FlattenInt64Pointer(from.Nsec3SaltLength)
	m.NsecType = flex.FlattenStringPointer(from.NsecType)
	m.ZskRolloverInterval = flex.FlattenInt64Pointer(from.ZskRolloverInterval)
	m.ZskSignatureValidity = flex.FlattenInt64Pointer(from.ZskSignatureValidity)
}
