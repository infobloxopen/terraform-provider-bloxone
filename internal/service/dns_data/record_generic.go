package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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

var _ recordResourceImplementor = &recordGenericResource{}
var _ recordDataSourceImplementor = &recordGenericResource{}

var validateGenericType = stringvalidator.NoneOfCaseInsensitive("A", "AAAA", "CAA", "CNAME", "DNAME", "MX", "NAPTR", "NS", "PTR", "SRV", "TXT", "HTTPS", "SVCB")

type genericRecordModel struct {
	SubFields types.List `tfsdk:"subfields"`
}

var genericRecordAttrTypes = map[string]attr.Type{
	"subfields": types.ListType{ElemType: types.ObjectType{AttrTypes: subFieldAttrTypes}},
}

type subFieldModel struct {
	Type       types.String `tfsdk:"type"`
	LengthKind types.String `tfsdk:"length_kind"`
	Value      types.String `tfsdk:"value"`
}

var subFieldAttrTypes = map[string]attr.Type{
	"type":        types.StringType,
	"length_kind": types.StringType,
	"value":       types.StringType,
}

type recordGenericResource struct{}

func NewRecordGenericResource() resource.Resource {
	return newRecordResource(&recordGenericResource{})
}

func NewRecordGenericDataSource() datasource.DataSource {
	return newRecordDataSource(&recordGenericResource{})

}
func (r recordGenericResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m genericRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return map[string]interface{}{
		"subfields": flex.ExpandFrameworkListNestedBlock(ctx, m.SubFields, diags, r.subFieldExpand),
	}
}

func (r recordGenericResource) subFieldExpand(ctx context.Context, o types.Object, diags *diag.Diagnostics) *map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m subFieldModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := &map[string]interface{}{
		"type":  flex.ExpandString(m.Type),
		"value": flex.ExpandString(m.Value),
	}
	if !m.LengthKind.IsNull() && !m.LengthKind.IsUnknown() {
		(*rdata)["length_kind"] = flex.ExpandString(m.LengthKind)
	}
	return rdata
}

func (r recordGenericResource) flattenRData(ctx context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(genericRecordAttrTypes)
	}
	t, d := types.ObjectValue(genericRecordAttrTypes, map[string]attr.Value{
		"subfields": flex.FlattenFrameworkListNestedBlock(ctx, from["subfields"].([]interface{}), subFieldAttrTypes, diags, r.subFieldFlatten),
	})
	diags.Append(d...)
	return t
}

func (r recordGenericResource) subFieldFlatten(_ context.Context, m *interface{}, diags *diag.Diagnostics) types.Object {
	from := (*m).(map[string]interface{})
	if from == nil {
		return types.ObjectNull(subFieldAttrTypes)
	}
	t, d := types.ObjectValue(subFieldAttrTypes, map[string]attr.Value{
		"type":        flattenRDataFieldString(from["type"], diags),
		"length_kind": flattenRDataFieldString(from["length_kind"], diags),
		"value":       flattenRDataFieldString(from["value"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordGenericResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: genericRecordAttrTypes}
	return attrTypes
}

func (r recordGenericResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The DNS resource record type specified in the textual mnemonic format or in the “TYPEnnn” format where “nnn” indicates the numeric type value.",
		Validators: []validator.String{
			validateGenericType,
		},
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"subfields": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required: true,
							MarkdownDescription: "Following types are supported:\n" +
								"  * _8BIT_: Unsigned 8-bit integer.\n" +
								"  * _16BIT_: Unsigned 16-bit integer.\n" +
								"  * _32BIT_: Unsigned 32-bit integer.\n" +
								"  * _IPV6_: IPv6 address. For example, \"abcd:123::abcd\".\n" +
								"  * _IPV4_: IPv4 address. For example, \"1.1.1.1\".\n" +
								"  * _DomainName_: Domain name (absolute or relative).\n" +
								"  * _TEXT_: ASCII text.\n" +
								"  * _BASE64_: Base64 encoded binary data.\n" +
								"  * _HEX_: Hex encoded binary data.\n" +
								"  * _PRESENTATION_: Presentation is a standard textual form of record data, as shown in a standard master zone file.\n\n" +
								"    For example, an IPSEC RDATA could be specified using the _PRESENTATION_ type field whose `value` is `10 1 2 192.0.2.38 AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ==`, instead of a sequence of the following subfields:\n\n" +
								"    ```\n" +
								"    8BIT: value=10\n" +
								"    8BIT: value=1\n" +
								"    8BIT: value=2\n" +
								"    IPV4: value=\"192.0.2.38\"\n" +
								"    BASE64 (without length_kind sub-subfield): value=\"AQNRU3mG7TVTO2BkR47usntb102uFJtugbo6BSGvgqt4AQ==\"\n" +
								"    ```\n" +
								"    If type is _PRESENTATION_, only one struct subfield can be specified.",
						},
						"length_kind": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "A string indicating the size in bits of a sub-subfield that is prepended to the value and encodes the length of the value. Valid values are:\n8: If type is ASCII or BASE64.\n16: If type is HEX.\nDefaults to none",
						},
						"value": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "A string representing the value for the sub-subfield",
						},
					},
				},
				Required: true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordGenericResource) recordType() string {
	return "Generic"
}

func (r recordGenericResource) resourceName() string {
	return "dns_record"
}

func (r recordGenericResource) dataSourceName() string {
	return "dns_records"
}

func (r recordGenericResource) description() string {
	return "Represents a DNS resource record in an authoritative zone."
}
