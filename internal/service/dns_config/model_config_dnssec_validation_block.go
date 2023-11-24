package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigDNSSECValidationBlockModel struct {
	DnssecEnableValidation types.Bool `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled          types.Bool `tfsdk:"dnssec_enabled"`
	DnssecTrustAnchors     types.List `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry   types.Bool `tfsdk:"dnssec_validate_expiry"`
}

var ConfigDNSSECValidationBlockAttrTypes = map[string]attr.Type{
	"dnssec_enable_validation": types.BoolType,
	"dnssec_enabled":           types.BoolType,
	"dnssec_trust_anchors":     types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigTrustAnchorAttrTypes}},
	"dnssec_validate_expiry":   types.BoolType,
}

var ConfigDNSSECValidationBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"dnssec_enable_validation": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _dnssec_enable_validation_ field.`,
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _dnssec_enabled_ field.`,
	},
	"dnssec_trust_anchors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigTrustAnchorResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _dnssec_trust_anchors_ field.`,
	},
	"dnssec_validate_expiry": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Field config for _dnssec_validate_expiry_ field.`,
	},
}

func ExpandConfigDNSSECValidationBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigDNSSECValidationBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigDNSSECValidationBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigDNSSECValidationBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigDNSSECValidationBlock {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigDNSSECValidationBlock{
		DnssecEnableValidation: flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:          flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:     flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandConfigTrustAnchor),
		DnssecValidateExpiry:   flex.ExpandBoolPointer(m.DnssecValidateExpiry),
	}
	return to
}

func FlattenConfigDNSSECValidationBlock(ctx context.Context, from *dns_config.ConfigDNSSECValidationBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigDNSSECValidationBlockAttrTypes)
	}
	m := ConfigDNSSECValidationBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigDNSSECValidationBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigDNSSECValidationBlockModel) Flatten(ctx context.Context, from *dns_config.ConfigDNSSECValidationBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigDNSSECValidationBlockModel{}
	}
	m.DnssecEnableValidation = types.BoolPointerValue(from.DnssecEnableValidation)
	m.DnssecEnabled = types.BoolPointerValue(from.DnssecEnabled)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, ConfigTrustAnchorAttrTypes, diags, FlattenConfigTrustAnchor)
	m.DnssecValidateExpiry = types.BoolPointerValue(from.DnssecValidateExpiry)
}
