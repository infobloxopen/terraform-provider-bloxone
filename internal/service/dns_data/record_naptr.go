package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

var _ recordResourceImplementor = &recordNAPTRResource{}
var _ recordDataSourceImplementor = &recordNAPTRResource{}

type naptrRecordModel struct {
	Flags       types.String `tfsdk:"flags"`
	Order       types.Int64  `tfsdk:"order"`
	Preference  types.Int64  `tfsdk:"preference"`
	Regexp      types.String `tfsdk:"regexp"`
	Replacement types.String `tfsdk:"replacement"`
	Services    types.String `tfsdk:"services"`
}

var naptrRecordAttrTypes = map[string]attr.Type{
	"flags":       types.StringType,
	"order":       types.Int64Type,
	"preference":  types.Int64Type,
	"regexp":      types.StringType,
	"replacement": types.StringType,
	"services":    types.StringType,
}

type recordNAPTRResource struct{}

func NewRecordNAPTRResource() resource.Resource {
	return newRecordResource(&recordNAPTRResource{})
}

func NewRecordNAPTRDataSource() datasource.DataSource {
	return newRecordDataSource(&recordNAPTRResource{})
}

func (r recordNAPTRResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m naptrRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := map[string]interface{}{
		"order":       flex.ExpandInt64(m.Order),
		"preference":  flex.ExpandInt64(m.Preference),
		"replacement": flex.ExpandString(m.Replacement),
		"services":    flex.ExpandString(m.Services),
	}

	// Optional fields
	if !m.Flags.IsNull() && !m.Flags.IsUnknown() {
		rdata["flags"] = flex.ExpandString(m.Flags)
	}
	if !m.Regexp.IsNull() && !m.Regexp.IsUnknown() {
		rdata["regexp"] = flex.ExpandString(m.Regexp)
	}

	return rdata
}

func (r recordNAPTRResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(naptrRecordAttrTypes)
	}
	t, d := types.ObjectValue(naptrRecordAttrTypes, map[string]attr.Value{
		"flags":       flattenRDataFieldString(from["flags"], diags),
		"order":       flattenRDataFieldInt64(from["order"], diags),
		"preference":  flattenRDataFieldInt64(from["preference"], diags),
		"regexp":      flattenRDataFieldString(from["regexp"], diags),
		"replacement": flattenRDataFieldString(from["replacement"], diags),
		"services":    flattenRDataFieldString(from["services"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordNAPTRResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: naptrRecordAttrTypes}
	return attrTypes
}

func (r recordNAPTRResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the record. This is always `NAPTR`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"flags": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A character string containing flags to control aspects of the rewriting and interpretation of the fields in the DNS resource record. The flags that are currently used are:\nU: Indicates that the output maps to a URI (Uniform Record Identifier).\nS: Indicates that the output is a domain name that has at least one SRV record. The DNS client must then send a query for the SRV record of the resulting domain name.\nA: Indicates that the output is a domain name that has at least one A or AAAA record. The DNS client must then send a query for the A or AAAA record of the resulting domain name.\nP: Indicates that the protocol specified in the services field defines the next step or phase.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1),
				},
			},
			"order": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "A 16-bit unsigned integer specifying the order in which the NAPTR records must be processed. Low numbers are processed before high numbers, and once a NAPTR is found whose rule “matches” the target, the client must not consider any NAPTRs with a higher value for order (except as noted below for the “flags” field. The range of the value is 0 to 65535.",
			},
			"preference": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "A 16-bit unsigned integer that specifies the order in which NAPTR records with equal “order” values should be processed, low numbers being processed before high numbers. This is similar to the preference field in an MX record, and is used so domain administrators can direct clients towards more capable hosts or lighter weight protocols. A client may look at records with higher preference values if it has a good reason to do so such as not understanding the preferred protocol or service. The range of the value is 0 to 65535.",
			},
			"regexp": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "A string containing a substitution expression that is applied to the original string held by the client in order to construct the next domain name to lookup.\n\nDefaults to none.",
			},
			"replacement": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The next name to query for NAPTR, SRV, or address records depending on the value of the flags field. This can be an absolute or relative domain name. Can be empty.",
			},
			"services": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Specifies the service(s) available down this rewrite path. It may also specify the particular protocol that is used to talk with a service. A protocol must be specified if the flags field states that the NAPTR is terminal. If a protocol is specified, but the flags field does not state that the NAPTR is terminal, the next lookup must be for a NAPTR. The client may choose not to perform the next lookup if the protocol is unknown, but that behavior must not be relied upon.\n\nThe service field may take any of the values below (using the Augmented BNF of RFC 2234):\n\nservice_field = [ [protocol] *(“+” rs)]\nprotocol = ALPHA * 31 ALPHANUM\nrs = ALPHA * 31 ALPHANUM\n\nThe protocol and rs fields are limited to 32 characters and must start with an alphabetic character.\n\nFor example, an optional protocol specification followed by 0 or more resolution services. Each resolution service is indicated by an initial ‘+’ character.\n\nNote that the empty string is also a valid service field. This will typically be seen at the beginning of a series of rules, when it is impossible to know what services and protocols will be offered by a particular service.\n\nThe actual format of the service request and response will be determined by the resolution protocol. Protocols need not offer all services. The labels for service requests shall be formed from the set of characters [A-Z0-9]. The case of the alphabetic characters is not significant.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordNAPTRResource) recordType() string {
	return "NAPTR"
}

func (r recordNAPTRResource) resourceName() string {
	return "dns_naptr_record"
}

func (r recordNAPTRResource) dataSourceName() string {
	return "dns_naptr_records"
}
