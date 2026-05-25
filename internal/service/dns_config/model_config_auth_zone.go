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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigAuthZoneModel struct {
	Comment                   types.String      `tfsdk:"comment"`
	CompartmentId             types.String      `tfsdk:"compartment_id"`
	CreatedAt                 timetypes.RFC3339 `tfsdk:"created_at"`
	Disabled                  types.Bool        `tfsdk:"disabled"`
	DnssecKeys                types.List        `tfsdk:"dnssec_keys"`
	DnssecSigningPolicy       types.Object      `tfsdk:"dnssec_signing_policy"`
	DnssecStatus              types.String      `tfsdk:"dnssec_status"`
	ExternalPrimaries         types.List        `tfsdk:"external_primaries"`
	ExternalProviders         types.List        `tfsdk:"external_providers"`
	ExternalProvidersMetadata types.Map         `tfsdk:"external_providers_metadata"`
	ExternalSecondaries       types.List        `tfsdk:"external_secondaries"`
	Fqdn                      types.String      `tfsdk:"fqdn"`
	GridPrimaries             types.List        `tfsdk:"grid_primaries"`
	GridSecondaries           types.List        `tfsdk:"grid_secondaries"`
	GssTsigEnabled            types.Bool        `tfsdk:"gss_tsig_enabled"`
	Id                        types.String      `tfsdk:"id"`
	InheritanceAssignedHosts  types.List        `tfsdk:"inheritance_assigned_hosts"`
	InheritanceSources        types.Object      `tfsdk:"inheritance_sources"`
	InitialSoaSerial          types.Int64       `tfsdk:"initial_soa_serial"`
	InternalSecondaries       types.List        `tfsdk:"internal_secondaries"`
	MappedSubnet              types.String      `tfsdk:"mapped_subnet"`
	Mapping                   types.String      `tfsdk:"mapping"`
	MaxRecordsPerType         types.Int64       `tfsdk:"max_records_per_type"`
	MaxTypesPerName           types.Int64       `tfsdk:"max_types_per_name"`
	Nameservers               types.List        `tfsdk:"nameservers"`
	NiosGridsMetadata         types.Map         `tfsdk:"nios_grids_metadata"`
	Notify                    types.Bool        `tfsdk:"notify"`
	Nsg                       types.String      `tfsdk:"nsg"`
	Nsgs                      types.List        `tfsdk:"nsgs"`
	Parent                    types.String      `tfsdk:"parent"`
	PrimaryType               types.String      `tfsdk:"primary_type"`
	ProtocolFqdn              types.String      `tfsdk:"protocol_fqdn"`
	QueryAcl                  types.List        `tfsdk:"query_acl"`
	SecondaryZoneRecordsSync  types.Bool        `tfsdk:"secondary_zone_records_sync"`
	Tags                      types.Map         `tfsdk:"tags"`
	TagsAll                   types.Map         `tfsdk:"tags_all"`
	TransferAcl               types.List        `tfsdk:"transfer_acl"`
	UpdateAcl                 types.List        `tfsdk:"update_acl"`
	UpdatedAt                 timetypes.RFC3339 `tfsdk:"updated_at"`
	UseForwardersForSubzones  types.Bool        `tfsdk:"use_forwarders_for_subzones"`
	View                      types.String      `tfsdk:"view"`
	Warnings                  types.List        `tfsdk:"warnings"`
	ZoneAuthority             types.Object      `tfsdk:"zone_authority"`
}

var ConfigAuthZoneAttrTypes = map[string]attr.Type{
	"comment":                     types.StringType,
	"compartment_id":              types.StringType,
	"created_at":                  timetypes.RFC3339Type{},
	"disabled":                    types.BoolType,
	"dnssec_keys":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigDNSSECKeyAttrTypes}},
	"dnssec_signing_policy":       types.ObjectType{AttrTypes: ConfigDNSSECSigningPolicyAttrTypes},
	"dnssec_status":               types.StringType,
	"external_primaries":          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalPrimaryAttrTypes}},
	"external_providers":          types.ListType{ElemType: types.ObjectType{AttrTypes: AuthZoneExternalProviderAttrTypes}},
	"external_providers_metadata": types.MapType{ElemType: types.StringType},
	"external_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigExternalSecondaryAttrTypes}},
	"fqdn":                        types.StringType,
	"grid_primaries":              types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigMemberServerAttrTypes}},
	"grid_secondaries":            types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigMemberServerAttrTypes}},
	"gss_tsig_enabled":            types.BoolType,
	"id":                          types.StringType,
	"inheritance_assigned_hosts":  types.ListType{ElemType: types.ObjectType{AttrTypes: Inheritance2AssignedHostAttrTypes}},
	"inheritance_sources":         types.ObjectType{AttrTypes: ConfigAuthZoneInheritanceAttrTypes},
	"initial_soa_serial":          types.Int64Type,
	"internal_secondaries":        types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigInternalSecondaryAttrTypes}},
	"mapped_subnet":               types.StringType,
	"mapping":                     types.StringType,
	"max_records_per_type":        types.Int64Type,
	"max_types_per_name":          types.Int64Type,
	"nameservers":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigNameserverAttrTypes}},
	"nios_grids_metadata":         types.MapType{ElemType: types.StringType},
	"notify":                      types.BoolType,
	"nsg":                         types.StringType,
	"nsgs":                        types.ListType{ElemType: types.StringType},
	"parent":                      types.StringType,
	"primary_type":                types.StringType,
	"protocol_fqdn":               types.StringType,
	"query_acl":                   types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigACLItemAttrTypes}},
	"secondary_zone_records_sync": types.BoolType,
	"tags":                        types.MapType{ElemType: types.StringType},
	"tags_all":                    types.MapType{ElemType: types.StringType},
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
		MarkdownDescription: "Optional. Comment for zone configuration.",
	},
	"compartment_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access view associated with the object. If no access view is associated with the object, the value defaults to empty.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
	},
	"dnssec_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigDNSSECKeyResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of DNSSEC keys used by the _AuthZone_ for zone signing.",
	},
	"dnssec_signing_policy": schema.SingleNestedAttribute{
		Attributes: ConfigDNSSECSigningPolicyResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
	},
	"dnssec_status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read Only.  DNSSEC status indicates the current DNSSEC signing status of the zone.  Possible values: - _UNSIGNED_: The zone is not signed with DNSSEC - _SIGNED_: The zone is fully signed with DNSSEC - _ROLLOVER_IN_PROGRESS_: DNSSEC key rollover is currently in progress - _SIGN_IN_PROGRESS_: The zone is currently being signed with DNSSEC - _UNSIGN_IN_PROGRESS_: The zone is currently being unsigned (DNSSEC removal in progress)",
	},
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalPrimaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. DNS primaries external to BloxOne DDI. Order is not significant.",
	},
	"external_providers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AuthZoneExternalProviderResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "list of external providers for the auth zone.",
	},
	"external_providers_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "External DNS providers metadata.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigExternalSecondaryResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "DNS secondaries external to BloxOne DDI. Order is not significant.",
	},
	"fqdn": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.",
	},
	"grid_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigMemberServerResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. The list of the NIOS Grid Primaries assigned to an AuthZone, only applicable for the NIOS Zones.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigMemberServerResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. The list of the NIOS Grid Secondaries assigned to an AuthZone, only applicable for the NIOS Zones.",
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
		MarkdownDescription: "The list of the inheritance assigned hosts of the object.",
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
		MarkdownDescription: "Optional. BloxOne DDI hosts acting as internal secondaries. Order is not significant.",
	},
	"mapped_subnet": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.",
	},
	"mapping": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward.",
	},
	"max_records_per_type": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(2000),
		MarkdownDescription: "The maximum number of records that can be stored in an RRset (records of same name and type), to prevent a slowdown in query processing due to an excessive number of those RRsets. The limit is enforced when serving the zone on-prem, not at the time of record creation or update. Exceeding the limit will result in the zone failing to load or to be updated. If 0, it means there is no limit. Defauts to _2000_.",
	},
	"max_types_per_name": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(100),
		MarkdownDescription: "The maximum number of record types that can be stored for an owner name, to prevent a slowdown in query processing due to an excessive number of those records. The limit is enforced when serving the zone on-prem, not at the time of record creation or update. Exceeding the limit will result in the zone failing to load or to be updated. If 0, it means there is no limit. Defauts to _100_.",
	},
	"nameservers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigNameserverResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. A list of DNS Nameservers of various roles. Cannot be configured if _nsg_ is configured.",
	},
	"nios_grids_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "NIOS Grids Metadata holds multiple NIOS grids data.",
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: `Also notify all external secondary DNS servers if enabled.  Defaults to _false_.`,
	},
	"nsg": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"nsgs": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
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
		MarkdownDescription: "Zone FQDN in punycode.",
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.",
	},
	"secondary_zone_records_sync": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Defines if secondary zone records should be synchronized.  Defaults to _false_. Only allowed to update when primary_type is \"external\".",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: `Tagging specifics.`,
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: `Tagging specifics includes default tags.`,
	},
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Clients must match this ACL to receive zone transfers.",
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigACLItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
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
		MarkdownDescription: "The list of an auth zone warnings.",
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

func (m *ConfigAuthZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dnsconfig.AuthZone {
	if m == nil {
		return nil
	}
	to := &dnsconfig.AuthZone{
		Comment:                   flex.ExpandStringPointer(m.Comment),
		CompartmentId:             flex.ExpandStringPointer(m.CompartmentId),
		Disabled:                  flex.ExpandBoolPointer(m.Disabled),
		DnssecSigningPolicy:       ExpandConfigDNSSECSigningPolicy(ctx, m.DnssecSigningPolicy, diags),
		ExternalPrimaries:         flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandConfigExternalPrimary),
		ExternalProvidersMetadata: flex.ExpandFrameworkMapString(ctx, m.ExternalProvidersMetadata, diags),
		ExternalSecondaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandConfigExternalSecondary),
		GridPrimaries:             flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimaries, diags, ExpandConfigMemberServer),
		GridSecondaries:           flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandConfigMemberServer),
		GssTsigEnabled:            flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:        ExpandConfigAuthZoneInheritance(ctx, m.InheritanceSources, diags),
		InternalSecondaries:       flex.ExpandFrameworkListNestedBlock(ctx, m.InternalSecondaries, diags, ExpandConfigInternalSecondary),
		MaxRecordsPerType:         flex.ExpandInt64Pointer(m.MaxRecordsPerType),
		MaxTypesPerName:           flex.ExpandInt64Pointer(m.MaxTypesPerName),
		Nameservers:               flex.ExpandFrameworkListNestedBlock(ctx, m.Nameservers, diags, ExpandConfigNameserver),
		NiosGridsMetadata:         flex.ExpandFrameworkMapString(ctx, m.NiosGridsMetadata, diags),
		Notify:                    flex.ExpandBoolPointer(m.Notify),
		Nsg:                       flex.ExpandStringPointer(m.Nsg),
		Nsgs:                      flex.ExpandFrameworkListString(ctx, m.Nsgs, diags),
		QueryAcl:                  flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandConfigACLItem),
		SecondaryZoneRecordsSync:  flex.ExpandBoolPointer(m.SecondaryZoneRecordsSync),
		Tags:                      flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransferAcl:               flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandConfigACLItem),
		UpdateAcl:                 flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandConfigACLItem),
		UseForwardersForSubzones:  flex.ExpandBoolPointer(m.UseForwardersForSubzones),
		ZoneAuthority:             ExpandConfigZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	if isCreate {
		to.Fqdn = flex.ExpandStringPointer(m.Fqdn)
		to.PrimaryType = flex.ExpandStringPointer(m.PrimaryType)
		to.View = flex.ExpandStringPointer(m.View)
		to.InitialSoaSerial = flex.ExpandInt64Pointer(m.InitialSoaSerial)
	}

	return to
}

func DataSourceFlattenConfigAuthZone(ctx context.Context, from *dnsconfig.AuthZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigAuthZoneAttrTypes)
	}
	m := ConfigAuthZoneModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, ConfigAuthZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigAuthZoneModel) Flatten(ctx context.Context, from *dnsconfig.AuthZone, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigAuthZoneModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DnssecKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecKeys, ConfigDNSSECKeyAttrTypes, diags, FlattenConfigDNSSECKey)
	m.DnssecSigningPolicy = FlattenConfigDNSSECSigningPolicy(ctx, from.DnssecSigningPolicy, diags)
	m.DnssecStatus = flex.FlattenStringPointer(from.DnssecStatus)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, ConfigExternalPrimaryAttrTypes, diags, FlattenConfigExternalPrimary)
	m.ExternalProviders = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalProviders, AuthZoneExternalProviderAttrTypes, diags, FlattenAuthZoneExternalProvider)
	m.ExternalProvidersMetadata = flex.FlattenFrameworkMapString(ctx, from.ExternalProvidersMetadata, diags)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, ConfigExternalSecondaryAttrTypes, diags, FlattenConfigExternalSecondary)
	m.Fqdn = flex.FlattenStringPointer(from.Fqdn)
	m.GridPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimaries, ConfigMemberServerAttrTypes, diags, FlattenConfigMemberServer)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, ConfigMemberServerAttrTypes, diags, FlattenConfigMemberServer)
	m.GssTsigEnabled = types.BoolPointerValue(from.GssTsigEnabled)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceAssignedHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, Inheritance2AssignedHostAttrTypes, diags, FlattenInheritance2AssignedHost)
	m.InheritanceSources = FlattenConfigAuthZoneInheritance(ctx, from.InheritanceSources, diags)
	m.InitialSoaSerial = flex.FlattenInt64Pointer(from.InitialSoaSerial)
	m.InternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.InternalSecondaries, ConfigInternalSecondaryAttrTypes, diags, FlattenConfigInternalSecondary)
	m.MappedSubnet = flex.FlattenStringPointer(from.MappedSubnet)
	m.Mapping = flex.FlattenStringPointer(from.Mapping)
	m.MaxRecordsPerType = flex.FlattenInt64Pointer(from.MaxRecordsPerType)
	m.MaxTypesPerName = flex.FlattenInt64Pointer(from.MaxTypesPerName)
	m.Nameservers = flex.FlattenFrameworkListNestedBlock(ctx, from.Nameservers, ConfigNameserverAttrTypes, diags, FlattenConfigNameserver)
	m.NiosGridsMetadata = flex.FlattenFrameworkMapString(ctx, from.NiosGridsMetadata, diags)
	m.Notify = types.BoolPointerValue(from.Notify)
	m.Nsg = flex.FlattenStringPointer(from.Nsg)
	m.Nsgs = flex.FlattenFrameworkListString(ctx, from.Nsgs, diags)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.PrimaryType = flex.FlattenStringPointer(from.PrimaryType)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.SecondaryZoneRecordsSync = types.BoolPointerValue(from.SecondaryZoneRecordsSync)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ConfigACLItemAttrTypes, diags, FlattenConfigACLItem)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.UseForwardersForSubzones = types.BoolPointerValue(from.UseForwardersForSubzones)
	m.View = flex.FlattenStringPointer(from.View)
	m.Warnings = flex.FlattenFrameworkListNestedBlock(ctx, from.Warnings, ConfigWarningAttrTypes, diags, FlattenConfigWarning)
	m.ZoneAuthority = FlattenConfigZoneAuthority(ctx, from.ZoneAuthority, diags)
}
