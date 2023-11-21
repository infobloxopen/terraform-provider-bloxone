package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/dns_data"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DataRecordModel struct {
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
	Options             types.Map         `tfsdk:"options"`
	ProviderMetadata    types.Map         `tfsdk:"provider_metadata"`
	Rdata               types.Map         `tfsdk:"rdata"`
	Source              types.List        `tfsdk:"source"`
	Subtype             types.String      `tfsdk:"subtype"`
	Tags                types.Map         `tfsdk:"tags"`
	Ttl                 types.Int64       `tfsdk:"ttl"`
	Type                types.String      `tfsdk:"type"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
	View                types.String      `tfsdk:"view"`
	ViewName            types.String      `tfsdk:"view_name"`
	Zone                types.String      `tfsdk:"zone"`
}

var DataRecordAttrTypes = map[string]attr.Type{
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
	"options":                types.MapType{ElemType: types.StringType},
	"provider_metadata":      types.MapType{ElemType: types.StringType},
	"rdata":                  types.ObjectType{},
	"source":                 types.ListType{ElemType: types.StringType},
	"subtype":                types.StringType,
	"tags":                   types.MapType{ElemType: types.StringType},
	"ttl":                    types.Int64Type,
	"type":                   types.StringType,
	"updated_at":             timetypes.RFC3339Type{},
	"view":                   types.StringType,
	"view_name":              types.StringType,
	"zone":                   types.StringType,
}

var DataRecordResourceSchemaAttributes = map[string]schema.Attribute{
	"absolute_name_spec": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Synthetic field, used to determine _zone_ and/or _name_in_zone_ field for records.",
	},
	"absolute_zone_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The absolute domain name of the zone where this record belongs.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
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
		Default:             booldefault.StaticBool(true),
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
	},
	"ipam_host": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name_in_zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The relative owner name to the zone origin. Must be specified for creating the DNS resource record and is read only for other operations.",
	},
	"options": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The DNS resource record type-specific non-protocol options.  Valid value for _A_ (Address) and _AAAA_ (IPv6 Address) records:  Option     | Description -----------|----------------------------------------- create_ptr | A boolean flag which can be set to _true_ for POST operation to automatically create the corresponding PTR record. check_rmz  | A boolean flag which can be set to _true_ for POST operation to check the existence of reverse zone for creating the corresponding PTR record. Only applicable if the _create_ptr_ option is set to _true_.   Valid value for _PTR_ (Pointer) records:  Option     | Description -----------|---------------------------------------- address    | For GET operation it contains the IPv4 or IPv6 address represented by the PTR record.<br><br>For POST and PATCH operations it can be used to create/update a PTR record based on the IP address it represents. In this case, in addition to the _address_ in the options field, need to specify the _view_ field. |",
	},
	"provider_metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "external DNS provider metadata.",
	},
	"rdata": schema.MapAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: "The DNS resource record data in JSON format. Certain DNS resource record-specific subfields are required for creating the DNS resource record.    Subfields for _A_ (Address) record:  Subfield | Description                           |Required ---------|---------------------------------------|-------- address  | The IPv4 address of the host.<br><br> | Yes  Subfields for _AAAA_ (IPv6 Address) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- address  | The IPv6 address of the host.<br><br> | Yes  Subfields for _CAA_ (Certification Authority Authorization) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- flags    | An unsigned 8-bit integer which specifies the CAA record flags. RFC 6844 defines one (highest) bit in flag octet, remaining bits are deferred for future use. This bit is referenced as _Critical_. When the bit is set (flag value == 128), issuers must not issue certificates in case CAA records contain unknown property tags.<br><br>Defaults to 0.<br><br> | No tag      | The CAA record property tag string which indicates the type of CAA record. The following property tags are defined by RFC 6844:<ul><li>_issue_: Used to explicitly authorize CA to issue certificates for the domain in which the property is published.</li><li>_issuewild_: Used to explicitly authorize a single CA to issue wildcard certificates for the domain in which the property is published.</li><li>_iodef_: Used to specify an email address or URL to report invalid certificate requests or issuers’ certificate policy violations.</li></ul>Note: _issuewild_ type takes precedence over _issue_.<br><br> | Yes value    | A string which contains the CAA record property value.<br><br>Specifies the CA who is authorized to issue a certificate for the domain if the CAA record property tag is _issue_ or _issuewild_.<br><br> Specifies the URL/email address to report CAA policy violation for the domain if the CAA record property tag is _iodef_.<br><br> | Yes  Subfields for _CNAME_ (Canonical Name) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- cname    | A domain name which specifies the canonical or primary name for the owner. The owner name is an alias. Can be empty.<br><br> | Yes  Subfields for _DNAME_ (Delegation Name) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- target   | The target domain name to which the zone will be mapped. Can be empty.<br><br> | Yes  Subfields for _DHCID_ (DHCP Identifier) record:  Subfield | Description                           | Required ---------|---------------------------------------|--------- dhcid    | The Base64 encoded string which contains DHCP client information.<br><br> | Yes  Subfields for _MX_ (Mail Exchanger) record:  Subfield   | Description                       | Required -----------|-----------------------------------|--------- exchange   | A domain name which specifies a host willing to act as a mail exchange for the owner name.<br><br> | Yes preference | An unsigned 16-bit integer which specifies the preference given to this RR among others at the same owner. Lower values are preferred. The range of the value is 0 to 65535. <br><br> | Yes  Subfields for _NAPTR_ (Naming Authority Pointer) record:  Subfield    | Description                         | Required ------------|-------------------------------------|--------- flags       | A character string containing flags to control aspects of the rewriting and interpretation of the fields in the DNS resource record. The flags that are currently used are: <ul><li> __U__: Indicates that the output maps to a URI (Uniform Record Identifier). </li><li> __S__: Indicates that the output is a domain name that has at least one SRV record. The DNS client must then send a query for the SRV record of the resulting domain name. </li><li> __A__: Indicates that the output is a domain name that has at least one A or AAAA record. The DNS client must then send a query for the A or AAAA record of the resulting domain name. </li><li> __P__: Indicates that the protocol specified in the _services_ field defines the next step or phase. </li></ul> | No order       | A 16-bit unsigned integer specifying the order in which the NAPTR records must be processed. Low numbers are processed before high numbers, and once a NAPTR is found whose rule \"matches\" the target, the client must not consider any NAPTRs with a higher value for order (except as noted below for the \"flags\" field. The range of the value is 0 to 65535. <br><br> | Yes preference  |A 16-bit unsigned integer that specifies the order in which NAPTR records with equal \"order\" values should be processed, low numbers being processed before high numbers. This is similar to the preference field in an MX record, and is used so domain administrators can direct clients towards more capable hosts or lighter weight protocols. A client may look at records with higher preference values if it has a good reason to do so such as not understanding the preferred protocol or service. The range of the value is 0 to 65535.<br><br> | Yes regexp      | A string containing a substitution expression that is applied to the original string held by the client in order to construct the next domain name to lookup.<br><br>Defaults to none.<br><br> | No replacement | The next name to query for NAPTR, SRV, or address records depending on the value of the _flags_ field. This can be an absolute or relative domain name. Can be empty.<br><br> | Yes services | Specifies the service(s) available down this rewrite path. It may also specify the particular protocol that is used to talk with a service. A protocol must be specified if the flags field states that the NAPTR is terminal. If a protocol is specified, but the flags field does not state that the NAPTR is terminal, the next lookup must be for a NAPTR. The client may choose not to perform the next lookup if the protocol is unknown, but that behavior must not be relied upon.<br><br>The service field may take any of the values below (using the Augmented BNF of RFC 2234):<br><br>service_field = [ [protocol] *(\"+\" rs)]<br>protocol = ALPHA * 31 ALPHANUM<br>rs = ALPHA * 31 ALPHANUM<br><br>The protocol and rs fields are limited to 32 characters and must start with an alphabetic character.<br><br> For example, an optional protocol specification followed by 0 or more resolution services. Each resolution service is indicated by an initial '+' character.<br><br> Note that the empty string is also a valid service field.  This will typically be seen at the beginning of a series of rules, when it is impossible to know what services and protocols will be offered by a particular service.<br><br> The actual format of the service request and response will be determined by the resolution protocol. Protocols need not offer all services. The labels for service requests shall be formed from the set of characters [A-Z0-9]. The case of the alphabetic characters is not significant.<br><br> | Yes  Subfields for _NS_ (Name Server) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- dname    | A domain-name which specifies a host which should be authoritative for the specified class and domain. Can be absolute or relative domain name and include UTF-8. <br><br> | Yes  Subfields for _PTR_ (Pointer) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- dname    | A domain name which points to some location in the domain name space. Can be absolute or relative domain name and include UTF-8. <br><br> | Yes  Subfields for _SOA_ (Start of Authority) record:  Subfield     | Description                         | Required ------------ |-------------------------------------|--------- expire       | The time interval in seconds after which zone data will expire and secondary server stops answering requests for the zone.<br><br> | No mname        | The domain name for the master server for the zone. Can be absolute or relative domain name.<br><br> | Yes negative_ttl | The time interval in seconds for which name servers can cache negative responses for zone. <br><br>Defaults to 900 seconds (15 minutes).<br><br> | No refresh      | The time interval in seconds that specifies how often secondary servers need to send a message to the primary server for a zone to check that their data is current, and retrieve fresh data if it is not.<br><br>Defaults to 10800 seconds (3 hours).<br><br> | No retry        | The time interval in seconds for which the secondary server will wait before attempting to recontact the primary server after a connection failure occurs.<br><br>Defaults to 3600 seconds (1 hour).<br><br> | No rname        | The domain name which specifies the mailbox of the person responsible for this zone. <br><br> | No serial       | An unsigned 32-bit integer that specifies the serial number of the zone. Used to indicate that zone data was updated, so the secondary name server can initiate zone transfer. The range of the value is 0 to 4294967295. <br><br> | No  Subfields for _SRV_ (Service) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- port     | An unsigned 16-bit integer which specifies the port on this target host of this service. The range of the value is 0 to 65535. This is often as specified in Assigned Numbers but need not be.<br><br> | Yes priority | An unsigned 16-bit integer which specifies the priority of this target host. The range of the value is 0 to 65535. A client must attempt to contact the target host with the lowest-numbered priority it can reach. Target hosts with the same priority should be tried in an order defined by the _weight_ field.<br><br>| Yes target   | The domain name of the target host. There must be one or more address records for this name, the name must not be an alias (in the sense of RFC 1034 or RFC 2181).<br><br>A target of \".\" means that the service is decidedly not available at this domain. | Yes weight   | An unsigned 16-bit integer which specifies a relative weight for entries with the same priority. The range of the value is 0 to 65535. Larger weights should be given a proportionately higher probability of being selected. Domain administrators should use weight 0 when there isn't any server selection to do, to make the RR easier to read for humans (less noisy). In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.<br><br>In the absence of a protocol whose specification calls for the use of other weighting information, a client arranges the SRV RRs of the same priority in the order in which target hosts, specified by the SRV RRs, will be contacted.<br><br>Defaults to 0.<br><br>| No  Subfields for _TXT_ (Text) record:  Subfield | Description                         | Required ---------|-------------------------------------|--------- text     | The semantics of the text depends on the domain where it is found.<br><br> | No  Generic record can be used to represent any DNS resource record not listed above.  Subfields for a generic record consist of a list of struct subfields, each having the following sub-subfields: Sub-subfield | Description                        | Required -------------|------------------------------------|--------- type         | Following types are supported:<ul><li>_8BIT_: Unsigned 8-bit integer. </li><li> _16BIT_: Unsigned 16-bit integer. </li><li> _32BIT_: Unsigned 32-bit integer. </li><li> _IPV6_: IPv6 address. For example, \"abcd:123::abcd\". </li><li> _IPV4_: IPv4 address. For example, \"1.1.1.1\". </li><li> _DomainName_: Domain name (absolute or relative). </li><li> _TEXT_: ASCII text. </li><li> _BASE64_: Base64 encoded binary data. </li><li> _HEX_: Hex encoded binary data. </li><li>_PRESENTATION_: Presentation is a standard textual form of record data, as shown in a standard master zone file. <br><br> For example, an IPSEC RDATA could be specified using the PRESENTATION type field whose value is \"10 1 2 192.0.2.38 AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ==\", instead of a sequence of the following subfields: <ul><li> 8BIT: value=10 </li><li> 8BIT: value=1 </li><li> 8BIT: value=2 </li><li> IPV4: value=\"192.0.2.38\" </li><li> BASE64 (without _length_kind_ sub-subfield): value=\"AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ==\" </li></ul></li></ul>If type is _PRESENTATION_, only one struct subfield can be specified. <br><br> | Yes length_kind  | A string indicating the size in bits of a sub-subfield that is prepended to the value and encodes the length of the value. Valid values are:<ul><li>_8_: If _type_ is _ASCII_ or _BASE64_. </li><li>_16_: If _type_ is _HEX_.</li></ul>Defaults to none. <br><br>| Only required for some types. value        | A string representing the value for the sub-subfield | Yes",
	},
	"source": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Source indicator                    | Description ------------------------------------|-------------------------------- _STATIC_                            |  Record was created manually by API call to _dns/record_. Valid for all record types except _SOA_. _SYSTEM_                            |  Record was created automatically based on name server assignment. Valid for _SOA_, _NS_, _A_, _AAAA_, and _PTR_ record types. _DYNAMIC_                           |  Record was created dynamically by performing dynamic update. Valid for all record types except _SOA_. _DELEGATED_                         |  Record was created automatically based on delegation servers assignment. Always extends the _SYSTEM_ bit. Valid for _NS_, _A_, _AAAA_, and _PTR_ record types. _DTC_                               |  Record was created automatically based on the DTC configuration. Always extends the _SYSTEM_ bit. Valid only for _IBMETA_ record type with _LBDN_ subtype. _STATIC_, _SYSTEM_                  |  Record was created manually by API call but it is obfuscated by record generated based on name server assignment. _DYNAMIC_, _SYSTEM_                 |  Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment. _DELEGATED_, _SYSTEM_               |  Record was created automatically based on delegation servers assignment. _SYSTEM_ will always accompany _DELEGATED_. _DTC_, _SYSTEM_                     |  Record was created automatically based on the DTC configuration. _SYSTEM_ will always accompany _DTC_. _STATIC_, _SYSTEM_, _DELEGATED_     |  Record was created manually by API call but it is obfuscated by record generated based on name server assignment as a result of creating a delegation. _DYNAMIC_, _SYSTEM_, _DELEGATED_    |  Record was created dynamically by DDNS but it is obfuscated by record generated based on name server assignment as a result of creating a delegation.",
	},
	"subtype": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DNS resource record subtype specified in the textual mnemonic format. Valid only in case _type_ is _IBMETA_.  Value | Numeric Type | Description ------|--------------|--------------------------------------------- | 0            | Default value LBDN  | 1            | LBDN record",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the DNS resource record in JSON format.",
	},
	"ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The record time to live value in seconds. The range of this value is 0 to 2147483647.  Defaults to TTL value from the SOA record of the zone.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The DNS resource record type specified in the textual mnemonic format or in the \"TYPEnnn\" format where \"nnn\" indicates the numeric type value.  Value  | Numeric Type | Description -------|--------------|--------------------------------------------- A      | 1            | Address record AAAA   | 28           | IPv6 Address record CAA    | 257          | Certification Authority Authorization record CNAME  | 5            | Canonical Name record DNAME  | 39           | Delegation Name record DHCID  | 49           | DHCP Identifier record MX     | 15           | Mail Exchanger record NAPTR  | 35           | Naming Authority Pointer record NS     | 2            | Name Server record PTR    | 12           | Pointer record SOA    | 6            | Start of Authority record SRV    | 33           | Service record TXT    | 16           | Text record IBMETA | 65536        | Infoblox meta records, not valid for DNS protocol (read-only)",
		Validators: []validator.String{
			stringvalidator.OneOf("A", "AAAA", "CAA", "CNAME", "DNAME", "DHCID", "MX", "NAPTR", "NS", "PTR", "SOA", "SRV", "TXT"),
		},
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
	},
	"view_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The display name of the DNS view that contains the parent zone of the DNS resource record.",
	},
	"zone": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func (m *DataRecordModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *dns_data.DataRecord {
	if m == nil {
		return nil
	}
	to := &dns_data.DataRecord{
		Comment:            flex.ExpandStringPointer(m.Comment),
		Disabled:           flex.ExpandBoolPointer(m.Disabled),
		InheritanceSources: ExpandDataRecordInheritance(ctx, m.InheritanceSources, diags),
		NameInZone:         flex.ExpandStringPointer(m.NameInZone),
		Options:            flex.ExpandFrameworkMapString(ctx, m.Options, diags),
		Rdata:              flex.ExpandFrameworkMapString(ctx, m.Rdata, diags),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Ttl:                flex.ExpandInt64Pointer(m.Ttl),
		View:               flex.ExpandStringPointer(m.View),
	}
	if isCreate {
		to.Type = flex.ExpandStringPointer(m.Type)
		to.Zone = flex.ExpandStringPointer(m.Zone)

	}
	return to
}

func FlattenDataRecord(ctx context.Context, from *dns_data.DataRecord, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DataRecordAttrTypes)
	}
	m := DataRecordModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DataRecordAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DataRecordModel) Flatten(ctx context.Context, from *dns_data.DataRecord, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DataRecordModel{}
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
	m.Options = flex.FlattenFrameworkMapString(ctx, from.Options, diags)
	m.ProviderMetadata = flex.FlattenFrameworkMapString(ctx, from.ProviderMetadata, diags)
	m.Rdata = flex.FlattenFrameworkMapString(ctx, from.Rdata, diags)
	m.Source = flex.FlattenFrameworkListString(ctx, from.Source, diags)
	m.Subtype = flex.FlattenStringPointer(from.Subtype)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Ttl = flex.FlattenInt64Pointer(from.Ttl)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.View = flex.FlattenStringPointer(from.View)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
	m.Zone = flex.FlattenStringPointer(from.Zone)
}
