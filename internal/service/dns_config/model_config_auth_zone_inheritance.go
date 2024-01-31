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

type ConfigAuthZoneInheritanceModel struct {
	GssTsigEnabled           types.Object `tfsdk:"gss_tsig_enabled"`
	Notify                   types.Object `tfsdk:"notify"`
	QueryAcl                 types.Object `tfsdk:"query_acl"`
	TransferAcl              types.Object `tfsdk:"transfer_acl"`
	UpdateAcl                types.Object `tfsdk:"update_acl"`
	UseForwardersForSubzones types.Object `tfsdk:"use_forwarders_for_subzones"`
	ZoneAuthority            types.Object `tfsdk:"zone_authority"`
}

var ConfigAuthZoneInheritanceAttrTypes = map[string]attr.Type{
	"gss_tsig_enabled":            types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"notify":                      types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"query_acl":                   types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"transfer_acl":                types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"update_acl":                  types.ObjectType{AttrTypes: ConfigInheritedACLItemsAttrTypes},
	"use_forwarders_for_subzones": types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"zone_authority":              types.ObjectType{AttrTypes: ConfigInheritedZoneAuthorityAttrTypes},
}

var ConfigAuthZoneInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"gss_tsig_enabled": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"notify": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"query_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"transfer_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"update_acl": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedACLItemsResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"use_forwarders_for_subzones": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes: ConfigInheritedZoneAuthorityResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
}

func ExpandConfigAuthZoneInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigAuthZoneInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigAuthZoneInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigAuthZoneInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigAuthZoneInheritance {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigAuthZoneInheritance{
		GssTsigEnabled:           ExpandInheritance2InheritedBool(ctx, m.GssTsigEnabled, diags),
		Notify:                   ExpandInheritance2InheritedBool(ctx, m.Notify, diags),
		QueryAcl:                 ExpandConfigInheritedACLItems(ctx, m.QueryAcl, diags),
		TransferAcl:              ExpandConfigInheritedACLItems(ctx, m.TransferAcl, diags),
		UpdateAcl:                ExpandConfigInheritedACLItems(ctx, m.UpdateAcl, diags),
		UseForwardersForSubzones: ExpandInheritance2InheritedBool(ctx, m.UseForwardersForSubzones, diags),
		ZoneAuthority:            ExpandConfigInheritedZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	return to
}

func FlattenConfigAuthZoneInheritance(ctx context.Context, from *dns_config.ConfigAuthZoneInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigAuthZoneInheritanceAttrTypes)
	}
	m := ConfigAuthZoneInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigAuthZoneInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigAuthZoneInheritanceModel) Flatten(ctx context.Context, from *dns_config.ConfigAuthZoneInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigAuthZoneInheritanceModel{}
	}
	m.GssTsigEnabled = FlattenInheritance2InheritedBool(ctx, from.GssTsigEnabled, diags)
	m.Notify = FlattenInheritance2InheritedBool(ctx, from.Notify, diags)
	m.QueryAcl = FlattenConfigInheritedACLItems(ctx, from.QueryAcl, diags)
	m.TransferAcl = FlattenConfigInheritedACLItems(ctx, from.TransferAcl, diags)
	m.UpdateAcl = FlattenConfigInheritedACLItems(ctx, from.UpdateAcl, diags)
	m.UseForwardersForSubzones = FlattenInheritance2InheritedBool(ctx, from.UseForwardersForSubzones, diags)
	m.ZoneAuthority = FlattenConfigInheritedZoneAuthority(ctx, from.ZoneAuthority, diags)
}
