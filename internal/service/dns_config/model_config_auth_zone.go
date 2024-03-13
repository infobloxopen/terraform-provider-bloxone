package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigAuthZoneModel struct {
	Comment                  types.String      `tfsdk:"comment"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	Disabled                 types.Bool        `tfsdk:"disabled"`
	ExternalPrimaries        types.List        `tfsdk:"external_primaries"`
	ExternalProviders        types.List        `tfsdk:"external_providers"`
	ExternalSecondaries      types.List        `tfsdk:"external_secondaries"`
	Fqdn                     types.String      `tfsdk:"fqdn"`
	GssTsigEnabled           types.Bool        `tfsdk:"gss_tsig_enabled"`
	Id                       types.String      `tfsdk:"id"`
	InheritanceAssignedHosts types.List        `tfsdk:"inheritance_assigned_hosts"`
	InheritanceSources       types.Object      `tfsdk:"inheritance_sources"`
	InitialSoaSerial         types.Int64       `tfsdk:"initial_soa_serial"`
	InternalSecondaries      types.List        `tfsdk:"internal_secondaries"`
	MappedSubnet             types.String      `tfsdk:"mapped_subnet"`
	Mapping                  types.String      `tfsdk:"mapping"`
	Notify                   types.Bool        `tfsdk:"notify"`
	Nsgs                     types.List        `tfsdk:"nsgs"`
	Parent                   types.String      `tfsdk:"parent"`
	PrimaryType              types.String      `tfsdk:"primary_type"`
	ProtocolFqdn             types.String      `tfsdk:"protocol_fqdn"`
	QueryAcl                 types.List        `tfsdk:"query_acl"`
	Tags                     types.Map         `tfsdk:"tags"`
	TransferAcl              types.List        `tfsdk:"transfer_acl"`
	UpdateAcl                types.List        `tfsdk:"update_acl"`
	UpdatedAt                timetypes.RFC3339 `tfsdk:"updated_at"`
	UseForwardersForSubzones types.Bool        `tfsdk:"use_forwarders_for_subzones"`
	View                     types.String      `tfsdk:"view"`
	Warnings                 types.List        `tfsdk:"warnings"`
	ZoneAuthority            types.Object      `tfsdk:"zone_authority"`
}

var ConfigAuthZoneAttrTypes = map[string]attr.Type{
	"comment":                     types.StringType,
	"created_at":                  timetypes.RFC3339Type{},
	"disabled":                    types.BoolType,
	"external_primaries":          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalPrimaryAttrTypes}},
	"external_providers":          types.ListType{ElemType: types.ObjectType{AttrTypes: AuthZoneExternalProviderAttrTypes}},
	"external_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalSecondaryAttrTypes}},
	"fqdn":                        types.StringType,
	"gss_tsig_enabled":            types.BoolType,
	"id":                          types.StringType,
	"inheritance_assigned_hosts":  types.ListType{ElemType: types.ObjectType{AttrTypes: Inheritance2AssignedHostAttrTypes}},
	"inheritance_sources":         types.ObjectType{AttrTypes: ConfigAuthZoneInheritanceAttrTypes},
	"initial_soa_serial":          types.Int64Type,
	"internal_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigInternalSecondaryAttrTypes}},
	"mapped_subnet":               types.StringType,
	"mapping":                     types.StringType,
	"notify":                      types.BoolType,
	"nsgs":                        types.ListType{ElemType: types.StringType},
	"parent":                      types.StringType,
	"primary_type":                types.StringType,
	"protocol_fqdn":               types.StringType,
	"query_acl":                   types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"tags":                        types.MapType{ElemType: types.StringType},
	"transfer_acl":                types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"update_acl":                  types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"updated_at":                  timetypes.RFC3339Type{},
	"use_forwarders_for_subzones": types.BoolType,
	"view":                        types.StringType,
	"warnings":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigWarningAttrTypes}},
	"zone_authority":              types.ObjectType{AttrTypes: ConfigZoneAuthorityAttrTypes},
}

var ConfigAuthZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: `Optional. Comment for zone configuration.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.`,
	},
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. DNS primaries external to BloxOne DDI. Order is not significant.`,
	},
	"external_providers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AuthZoneExternalProviderResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `list of external providers for the auth zone.`,
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `DNS secondaries external to BloxOne DDI. Order is not significant.`,
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: `Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.`,
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_assigned_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Inheritance2AssignedHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `The list of the inheritance assigned hosts of the object.`,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: ConfigAuthZoneInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"initial_soa_serial": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(1),
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: `On-create-only. SOA serial is allowed to be set when the authoritative zone is created.`,
	},
	"internal_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigInternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.`,
	},
	"mapped_subnet": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.`,
	},
	"mapping": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward.`,
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Also notify all external secondary DNS servers if enabled.  Defaults to _false_.`,
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"primary_type": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: `Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a BloxOne DDI host.`,
	},
	"protocol_fqdn": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Zone FQDN in punycode.`,
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `Tagging specifics.`,
	},
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Clients must match this ACL to receive zone transfers.`,
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"use_forwarders_for_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: `Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.`,
	},
	"view": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"warnings": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigWarningResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `The list of an auth zone warnings.`,
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes: ConfigZoneAuthorityResourceSchemaAttributes,
		Computed:   true,
		//Default: objectdefault.StaticValue(types.ObjectValueMust(ConfigZoneAuthorityAttrTypes, map[string]attr.Value{
		//	"default_ttl":       types.Int64Value(28800),
		//	"expire":            types.Int64Value(2.4192e+06),
		//	"mname":             types.StringNull(),
		//	"negative_ttl":      types.Int64Value(900),
		//	"protocol_mname":    types.StringValue("ns.b1ddi.tf-acc-test.com."),
		//	"protocol_rname":    types.StringValue("hostmaster.tf-acc-test.com"),
		//	"refresh":           types.Int64Value(10800),
		//	"retry":             types.Int64Value(3600),
		//	"rname":             types.StringNull(),
		//	"use_default_mname": types.BoolValue(true),
		//}),
		//),
	},
}

func (m *ConfigAuthZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns_config.ConfigAuthZone {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigAuthZone{
		Comment:                  m.Comment.ValueStringPointer(),
		Disabled:                 m.Disabled.ValueBoolPointer(),
		ExternalPrimaries:        flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandConfigExternalPrimary),
		ExternalSecondaries:      flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandConfigExternalSecondary),
		GssTsigEnabled:           m.GssTsigEnabled.ValueBoolPointer(),
		InheritanceSources:       ExpandConfigAuthZoneInheritance(ctx, m.InheritanceSources, diags),
		InternalSecondaries:      flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandConfigInternalSecondary),
		Notify:                   m.Notify.ValueBoolPointer(),
		Nsgs:                     flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		QueryAcl:                 flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandConfigACLItem),
		Tags:                     flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransferAcl:              flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandConfigACLItem),
		UpdateAcl:                flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandConfigACLItem),
		UseForwardersForSubzones: m.UseForwardersForSubzones.ValueBoolPointer(),
		ZoneAuthority:            ExpandConfigZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.PrimaryType = flex.ExpandStringPointer(m.PrimaryType)
		to.View = flex.ExpandStringPointer(m.View)
		to.InitialSoaSerial = flex.ExpandInt64Pointer(m.InitialSoaSerial)
	}

	return to
}

func FlattenConfigAuthZone(ctx context.Context, from *dns_config.ConfigAuthZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigAuthZoneAttrTypes)
	}
	m := ConfigAuthZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigAuthZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigAuthZoneModel) Flatten(ctx context.Context, from *dns_config.ConfigAuthZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigAuthZoneModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ConfigExternalPrimaryAttrTypes, diags, FlattenConfigExternalPrimary)
	m.ExternalProviders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalProviders, AuthZoneExternalProviderAttrTypes, diags, FlattenAuthZoneExternalProvider)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ConfigExternalSecondaryAttrTypes, diags, FlattenConfigExternalSecondary)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceAssignedHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, Inheritance2AssignedHostAttrTypes, diags, FlattenInheritance2AssignedHost)
	m.InheritanceSources = FlattenConfigAuthZoneInheritance(ctx, from.InheritanceSources, diags)
	m.InitialSoaSerial = flex.FlattenInt64Pointer(from.InitialSoaSerial)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, ConfigInternalSecondaryAttrTypes, diags, FlattenConfigInternalSecondary)
	m.MappedSubnet = flex.FlattenStringPointer(from.MappedSubnet)
	m.Mapping = flex.FlattenStringPointer(from.Mapping)
	m.Notify = types.BoolPointerValue(from.Notify)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.PrimaryType = flex.FlattenStringPointer(from.PrimaryType)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.UseForwardersForSubzones = types.BoolPointerValue(from.UseForwardersForSubzones)
	m.View = flex.FlattenStringPointer(from.View)
	m.Warnings = flex.FlattenFrameworkListNestedBlock(ctx, from.Warnings, ConfigWarningAttrTypes, diags, FlattenConfigWarning)
	m.ZoneAuthority = FlattenConfigZoneAuthority(ctx, from.ZoneAuthority, diags)
}
