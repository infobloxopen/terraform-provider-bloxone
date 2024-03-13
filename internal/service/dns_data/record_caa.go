package dns_data

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

var _ recordResourceImplementor = &recordCAAResource{}
var _ recordDataSourceImplementor = &recordCAAResource{}

type caaRecordModel struct {
	Flags types.Int64  `tfsdk:"flags"`
	Tag   types.String `tfsdk:"tag"`
	Value types.String `tfsdk:"value"`
}

var caaRecordAttrTypes = map[string]attr.Type{
	"flags": types.Int64Type,
	"tag":   types.StringType,
	"value": types.StringType,
}

type recordCAAResource struct{}

func NewRecordCAAResource() resource.Resource {
	return newRecordResource(&recordCAAResource{})
}

func NewRecordCAADataSource() datasource.DataSource {
	return newRecordDataSource(&recordCAAResource{})
}

func (r recordCAAResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m caaRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := map[string]interface{}{
		"tag":   flex.ExpandString(m.Tag),
		"value": flex.ExpandString(m.Value),
	}
	// Optional fields
	if !m.Flags.IsNull() && !m.Flags.IsUnknown() {
		rdata["flags"] = flex.ExpandInt64(m.Flags)
	}
	return rdata
}

func (r recordCAAResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(caaRecordAttrTypes)
	}

	t, d := types.ObjectValue(caaRecordAttrTypes, map[string]attr.Value{
		"flags": flattenRDataFieldInt64(from["flags"], false, diags),
		"tag":   flattenRDataFieldString(from["tag"], diags),
		"value": flattenRDataFieldString(from["value"], diags),
	})
	diags.Append(d...)
	return t
}

func flattenRDataFieldInt64(val interface{}, zeroAsNull bool, diags *diag.Diagnostics) basetypes.Int64Value {
	if val == nil {
		return types.Int64Null()
	}
	if _, ok := val.(float64); !ok {
		diags.AddError("Value conversion error", fmt.Sprintf("Failed to convert value '%v' to integer", val))
		return types.Int64Null()
	}

	if zeroAsNull {
		return flex.FlattenInt64(int64(val.(float64)))
	}
	return types.Int64Value(int64(val.(float64)))
}

func flattenRDataFieldString(val interface{}, diags *diag.Diagnostics) basetypes.StringValue {
	if val == nil {
		return types.StringNull()
	}
	if _, ok := val.(string); !ok {
		diags.AddError("Value conversion error", fmt.Sprintf("Failed to convert value '%v' to string", val))
		return types.StringNull()
	}
	if val.(string) == "" {
		// rdata is sent as a map, so we need to return a null value if the string is empty
		// if the empty string is significant, this can be wrapped around a emptyAsNull flag
		return types.StringNull()
	}
	return flex.FlattenString(val.(string))
}

func (r recordCAAResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: caaRecordAttrTypes}
	return attrTypes
}

func (r recordCAAResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the record. This is always `CAA`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"flags": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(0),
				MarkdownDescription: "An unsigned 8-bit integer which specifies the CAA record flags. RFC 6844 defines one (highest) bit in flag octet, remaining bits are deferred for future use. This bit is referenced as Critical. When the bit is set (flag value == 128), issuers must not issue certificates in case CAA records contain unknown property tags.",
			},
			"tag": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The CAA record property tag string which indicates the type of CAA record. The following property tags are defined by RFC 6844:\nissue: Used to explicitly authorize CA to issue certificates for the domain in which the property is published.\nissuewild: Used to explicitly authorize a single CA to issue wildcard certificates for the domain in which the property is published.\niodef: Used to specify an email address or URL to report invalid certificate requests or issuers’ certificate policy violations.\nNote: issuewild type takes precedence over issue.",
			},
			"value": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "A string which contains the CAA record property value.\n\nSpecifies the CA who is authorized to issue a certificate for the domain if the CAA record property tag is issue or issuewild.\n\nSpecifies the URL/email address to report CAA policy violation for the domain if the CAA record property tag is iodef.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordCAAResource) recordType() string {
	return "CAA"
}

func (r recordCAAResource) resourceName() string {
	return "dns_caa_record"
}

func (r recordCAAResource) dataSourceName() string {
	return "dns_caa_records"
}
