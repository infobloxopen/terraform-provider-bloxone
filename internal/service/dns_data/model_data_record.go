package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/dnsdata"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type dataRecordModel struct {
	AbsoluteNameSpec    types.String      `tfsdk:"absolute_name_spec"`
	AbsoluteZoneName    types.String      `tfsdk:"absolute_zone_name"`
	Comment             types.String      `tfsdk:"comment"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	Delegation          types.String      `tfsdk:"delegation"`
	Disabled            types.Bool        `tfsdk:"disabled"`
	DnsAbsoluteNameSpec types.String      `tfsdk:"dns_absolute_name_spec"`
	DnsAbsoluteZoneName types.String      `tfsdk:"dns_absolute_zone_name"`
	DnsNameInZone       types.String      `tfsdk:"dns_name_in_zone"`
	DnsRdata            types.String      `tfsdk:"dns_rdata"`
	Id                  types.String      `tfsdk:"id"`
	InheritanceSources  types.Object      `tfsdk:"inheritance_sources"`
	IpamHost            types.String      `tfsdk:"ipam_host"`
	NameInZone          types.String      `tfsdk:"name_in_zone"`
	Options             types.Object      `tfsdk:"options"`
	ProviderMetadata    types.Map         `tfsdk:"provider_metadata"`
	Rdata               types.Object      `tfsdk:"rdata"`
	Source              types.List        `tfsdk:"source"`
	Subtype             types.String      `tfsdk:"subtype"`
	Tags                types.Map         `tfsdk:"tags"`
	TagsAll             types.Map         `tfsdk:"tags_all"`
	Ttl                 types.Int64       `tfsdk:"ttl"`
	Type                types.String      `tfsdk:"type"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
	View                types.String      `tfsdk:"view"`
	ViewName            types.String      `tfsdk:"view_name"`
	Zone                types.String      `tfsdk:"zone"`
}

func recordCommonAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"absolute_name_spec":     types.StringType,
		"absolute_zone_name":     types.StringType,
		"comment":                types.StringType,
		"created_at":             timetypes.RFC3339Type{},
		"delegation":             types.StringType,
		"disabled":               types.BoolType,
		"dns_absolute_name_spec": types.StringType,
		"dns_absolute_zone_name": types.StringType,
		"dns_name_in_zone":       types.StringType,
		"dns_rdata":              types.StringType,
		"id":                     types.StringType,
		"inheritance_sources":    types.ObjectType{AttrTypes: DataRecordInheritanceAttrTypes},
		"ipam_host":              types.StringType,
		"name_in_zone":           types.StringType,
		"options":                types.ObjectType{},
		"provider_metadata":      types.MapType{ElemType: types.StringType},
		"source":                 types.ListType{ElemType: types.StringType},
		"subtype":                types.StringType,
		"tags":                   types.MapType{ElemType: types.StringType},
		"tags_all":               types.MapType{ElemType: types.StringType},
		"ttl":                    types.Int64Type,
		"type":                   types.StringType,
		"updated_at":             timetypes.RFC3339Type{},
		"view":                   types.StringType,
		"view_name":              types.StringType,
		"zone":                   types.StringType,
	}
}

func recordCommonSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"absolute_name_spec": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.",
			Validators: []validator.String{
				stringvalidator.AlsoRequires(path.MatchRoot("view")),
				stringvalidator.ConflictsWith(path.MatchRoot("zone"), path.MatchRoot("name_in_zone")),
			},
		},
		"absolute_zone_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The absolute domain name of the zone where this record belongs.",
		},
		"comment": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString(""),
			MarkdownDescription: "The description for the DNS resource record. May contain 0 to 1024 characters. Can include UTF-8.",
		},
		"created_at": schema.StringAttribute{
			CustomType:          timetypes.RFC3339Type{},
			Computed:            true,
			MarkdownDescription: "The timestamp when the object has been created.",
		},
		"delegation": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"disabled": schema.BoolAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "Indicates if the DNS resource record is disabled. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
			Default:             booldefault.StaticBool(false),
		},
		"dns_absolute_name_spec": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The DNS protocol textual representation of _absolute_name_spec_.",
		},
		"dns_absolute_zone_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The DNS protocol textual representation of the absolute domain name of the zone where this record belongs.",
		},
		"dns_name_in_zone": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The DNS protocol textual representation of the relative owner name for the DNS zone.",
		},
		"dns_rdata": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The DNS protocol textual representation of the DNS resource record data.",
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"inheritance_sources": schema.SingleNestedAttribute{
			Attributes: DataRecordInheritanceResourceSchemaAttributes,
			Optional:   true,
			Computed:   true,
			PlanModifiers: []planmodifier.Object{
				objectplanmodifier.UseStateForUnknown(),
			},
		},
		"ipam_host": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"name_in_zone": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The relative owner name to the zone origin. Must be specified for creating the DNS resource record and is read only for other operations.",
			Validators: []validator.String{
				stringvalidator.AlsoRequires(path.MatchRoot("zone")),
				stringvalidator.ConflictsWith(path.MatchRoot("absolute_name_spec"), path.MatchRoot("view")),
			},
		},
		"options": schema.SingleNestedAttribute{
			Attributes:          map[string]schema.Attribute{},
			Computed:            true,
			MarkdownDescription: "The DNS resource record type-specific non-protocol options.",
		},
		"provider_metadata": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "external DNS provider metadata.",
		},
		"source": schema.ListAttribute{
			ElementType: types.StringType,
			Computed:    true,
			MarkdownDescription: "Valid values are: \n\n" +
				"  | Source indicator                    | Description                                                                                                                                                     |\n" +
				"  |-------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|\n" +
				"  | _STATIC_                            | Record was created manually by API call to _dns/record_. Valid for all record types except _SOA_.                                                               |\n" +
				"  | _SYSTEM_                            | Record was created automatically based on name server assignment. Valid for _SOA_, _NS_, _A_, _AAAA_, and _PTR_ record types.                                   |\n" +
				"  | _DYNAMIC_                           | Record was created dynamically by performing dynamic update. Valid for all record types except _SOA_.                                                           |\n" +
				"  | _DELEGATED_                         | Record was created automatically based on delegation servers assignment. Always extends the _SYSTEM_ bit. Valid for _NS_, _A_, _AAAA_, and _PTR_ record types.  |\n" +
				"  | _DTC_                               | Record was created automatically based on the DTC configuration. Always extends the _SYSTEM_ bit. Valid only for _IBMETA_ record type with _LBDN_ subtype.      |\n" +
				"  | _STATIC_, _SYSTEM_                  | Record was created manually by API call but it is obfuscated by record generated based on name server assignment.                                               |\n" +
				"  | _DYNAMIC_, _SYSTEM_                 | Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment.                                                |\n" +
				"  | _DELEGATED_, _SYSTEM_               | Record was created automatically based on delegation servers assignment. _SYSTEM_ will always accompany _DELEGATED_.                                            |\n" +
				"  | _DTC_, _SYSTEM_                     | Record was created automatically based on the DTC configuration. _SYSTEM_ will always accompany _DTC_.                                                          |\n" +
				"  | _STATIC_, _SYSTEM_, _DELEGATED_     | Record was created manually by API call but it is obfuscated by record generated based on name server assignment as a result of creating a delegation.          |\n" +
				"  | _DYNAMIC_, _SYSTEM_, _DELEGATED_    | Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment as a result of creating a delegation.           |\n" +
				"  <br>",
		},
		"subtype": schema.StringAttribute{
			Computed: true,
			MarkdownDescription: "The DNS resource record subtype specified in the textual mnemonic format. Valid only in case _type_ is _IBMETA_.\n\n" +
				"  | Value | Numeric Type | Description   |\n" +
				"  |-------|--------------|---------------|\n" +
				"  |       | 0            | Default value |\n" +
				"  | LBDN  | 1            | LBDN record   |\n" +
				"  <br>",
		},
		"tags": schema.MapAttribute{
			ElementType:         types.StringType,
			Optional:            true,
			Computed:            true,
			Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
			MarkdownDescription: "The tags for the DNS resource record in JSON format.",
		},
		"tags_all": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "The tags for the DNS resource record including default tags.",
		},
		"ttl": schema.Int64Attribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The record time to live value in seconds. The range of this value is 0 to 2147483647.  Defaults to TTL value from the SOA record of the zone.",
		},
		"updated_at": schema.StringAttribute{
			CustomType:          timetypes.RFC3339Type{},
			Computed:            true,
			MarkdownDescription: "The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.",
		},
		"view": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
			},
			Validators: []validator.String{
				stringvalidator.Any(
					stringvalidator.AlsoRequires(path.MatchRoot("absolute_name_spec")),
					stringvalidator.AlsoRequires(path.MatchRoot("options").AtName("address")),
				),
				stringvalidator.ConflictsWith(path.MatchRoot("zone"), path.MatchRoot("name_in_zone")),
			},
		},
		"view_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The display name of the DNS view that contains the parent zone of the DNS resource record.",
		},
		"zone": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
			},
			Validators: []validator.String{
				stringvalidator.ConflictsWith(path.MatchRoot("absolute_name_spec"), path.MatchRoot("view")),
			},
		},
	}
}

func (m *dataRecordModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool, impl recordModelCommon) *dnsdata.Record {
	if m == nil {
		return nil
	}
	to := &dnsdata.Record{
		AbsoluteNameSpec:   flex.ExpandStringPointer(m.AbsoluteNameSpec),
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		InheritanceSources: ExpandDataRecordInheritance(ctx, m.InheritanceSources, diags),
		NameInZone:         flex.ExpandStringPointer(m.NameInZone),
		Rdata:              impl.expandRData(ctx, m.Rdata, diags),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
	}
	if recordWithOptionsImpl, ok := impl.(recordModelWithOptions); ok {
		to.Options = recordWithOptionsImpl.expandOptions(ctx, m.Options, diags)
	}

	if isCreate {
		rType := impl.recordType()
		to.Type = &rType
		if rType == "Generic" {
			to.Type = flex.ExpandStringPointer(m.Type)
		}
		to.Zone = flex.ExpandStringPointer(m.Zone)
		to.View = flex.ExpandStringPointer(m.View)
	}
	return to
}

func (m *dataRecordModel) Flatten(ctx context.Context, from *dnsdata.Record, diags *diag.Diagnostics, impl recordModelCommon) {
	if from == nil {
		return
	}
	if m == nil {
		*m = dataRecordModel{}
	}
	m.AbsoluteNameSpec = flex.FlattenStringPointer(from.AbsoluteNameSpec)
	m.AbsoluteZoneName = flex.FlattenStringPointer(from.AbsoluteZoneName)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Delegation = flex.FlattenStringPointer(from.Delegation)
	m.Disabled = types.BoolPointerValue(from.Disabled)
	m.DnsAbsoluteNameSpec = flex.FlattenStringPointer(from.DnsAbsoluteNameSpec)
	m.DnsAbsoluteZoneName = flex.FlattenStringPointer(from.DnsAbsoluteZoneName)
	m.DnsNameInZone = flex.FlattenStringPointer(from.DnsNameInZone)
	m.DnsRdata = flex.FlattenStringPointer(from.DnsRdata)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenDataRecordInheritance(ctx, from.InheritanceSources, diags)
	m.IpamHost = flex.FlattenStringPointer(from.IpamHost)
	m.NameInZone = flex.FlattenStringPointer(from.NameInZone)
	m.ProviderMetadata = flex.FlattenFrameworkMapString(ctx, from.ProviderMetadata, diags)
	m.Rdata = impl.flattenRData(ctx, from.Rdata, diags)
	m.Source = flex.FlattenFrameworkListString(ctx, from.Source, diags)
	m.Subtype = flex.FlattenStringPointer(from.Subtype)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.View = flex.FlattenStringPointer(from.View)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.Zone = flex.FlattenStringPointer(from.Zone)

	if recordWithOptionsImpl, ok := impl.(recordModelWithOptions); ok {
		m.Options = recordWithOptionsImpl.flattenOptions(ctx, m.Options, from.Options, diags)
	} else {
		m.Options = types.ObjectNull(nil)
	}
}
