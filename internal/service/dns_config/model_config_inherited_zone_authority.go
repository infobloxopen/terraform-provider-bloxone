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

type ConfigInheritedZoneAuthorityModel struct {
	DefaultTtl    types.Object `tfsdk:"default_ttl"`
	Expire        types.Object `tfsdk:"expire"`
	MnameBlock    types.Object `tfsdk:"mname_block"`
	NegativeTtl   types.Object `tfsdk:"negative_ttl"`
	ProtocolRname types.Object `tfsdk:"protocol_rname"`
	Refresh       types.Object `tfsdk:"refresh"`
	Retry         types.Object `tfsdk:"retry"`
	Rname         types.Object `tfsdk:"rname"`
}

var ConfigInheritedZoneAuthorityAttrTypes = map[string]attr.Type{
	"default_ttl":    types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"expire":         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"mname_block":    types.ObjectType{AttrTypes: ConfigInheritedZoneAuthorityMNameBlockAttrTypes},
	"negative_ttl":   types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"protocol_rname": types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
	"refresh":        types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"retry":          types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
	"rname":          types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
}

var ConfigInheritedZoneAuthorityResourceSchemaAttributes = map[string]schema.Attribute{
	"default_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"expire": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"mname_block": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedZoneAuthorityMNameBlockResourceSchemaAttributes,
		Optional:   true,
	},
	"negative_ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"protocol_rname": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"refresh": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"retry": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
	"rname": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandConfigInheritedZoneAuthority(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigInheritedZoneAuthority {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigInheritedZoneAuthorityModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigInheritedZoneAuthorityModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigInheritedZoneAuthority {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigInheritedZoneAuthority{
		DefaultTtl:    ExpandInheritance2InheritedUInt32(ctx, m.DefaultTtl, diags),
		Expire:        ExpandInheritance2InheritedUInt32(ctx, m.Expire, diags),
		MnameBlock:    ExpandConfigInheritedZoneAuthorityMNameBlock(ctx, m.MnameBlock, diags),
		NegativeTtl:   ExpandInheritance2InheritedUInt32(ctx, m.NegativeTtl, diags),
		ProtocolRname: ExpandInheritance2InheritedString(ctx, m.ProtocolRname, diags),
		Refresh:       ExpandInheritance2InheritedUInt32(ctx, m.Refresh, diags),
		Retry:         ExpandInheritance2InheritedUInt32(ctx, m.Retry, diags),
		Rname:         ExpandInheritance2InheritedString(ctx, m.Rname, diags),
	}
	return to
}

func FlattenConfigInheritedZoneAuthority(ctx context.Context, from *dns_config.ConfigInheritedZoneAuthority, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigInheritedZoneAuthorityAttrTypes)
	}
	m := ConfigInheritedZoneAuthorityModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigInheritedZoneAuthorityAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigInheritedZoneAuthorityModel) Flatten(ctx context.Context, from *dns_config.ConfigInheritedZoneAuthority, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigInheritedZoneAuthorityModel{}
	}
	m.DefaultTtl = FlattenInheritance2InheritedUInt32(ctx, from.DefaultTtl, diags)
	m.Expire = FlattenInheritance2InheritedUInt32(ctx, from.Expire, diags)
	m.MnameBlock = FlattenConfigInheritedZoneAuthorityMNameBlock(ctx, from.MnameBlock, diags)
	m.NegativeTtl = FlattenInheritance2InheritedUInt32(ctx, from.NegativeTtl, diags)
	m.ProtocolRname = FlattenInheritance2InheritedString(ctx, from.ProtocolRname, diags)
	m.Refresh = FlattenInheritance2InheritedUInt32(ctx, from.Refresh, diags)
	m.Retry = FlattenInheritance2InheritedUInt32(ctx, from.Retry, diags)
	m.Rname = FlattenInheritance2InheritedString(ctx, from.Rname, diags)
}
