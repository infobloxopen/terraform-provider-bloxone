package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

var _ recordResourceImplementor = &recordMXResource{}
var _ recordDataSourceImplementor = &recordMXResource{}

type mxRecordModel struct {
	Exchange   types.String `tfsdk:"exchange"`
	Preference types.Int64  `tfsdk:"preference"`
}

var mxRecordAttrTypes = map[string]attr.Type{
	"exchange":   types.StringType,
	"preference": types.Int64Type,
}

type recordMXResource struct{}

func NewRecordMXResource() resource.Resource {
	return newRecordResource(&recordMXResource{})
}

func NewRecordMXDataSource() datasource.DataSource {
	return newRecordDataSource(&recordMXResource{})
}

func (r recordMXResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m mxRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"exchange":   flex.ExpandString(m.Exchange),
		"preference": flex.ExpandInt64(m.Preference),
	}

}

func (r recordMXResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(mxRecordAttrTypes)
	}
	t, d := types.ObjectValue(mxRecordAttrTypes, map[string]attr.Value{
		"exchange":   flattenRDataFieldString(from["exchange"], diags),
		"preference": flattenRDataFieldInt64(from["preference"], true, diags),
	})
	diags.Append(d...)
	return t
}

func (r recordMXResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: mxRecordAttrTypes}
	return attrTypes
}

func (r recordMXResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the record. This is always `MX`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"exchange": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "A domain name which specifies a host willing to act as a mail exchange for the owner name.",
			},
			"preference": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "An unsigned 16-bit integer which specifies the preference given to this RR among others at the same owner. Lower values are preferred. The range of the value is 0 to 65535.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordMXResource) recordType() string {
	return "MX"
}

func (r recordMXResource) resourceName() string {
	return "dns_mx_record"
}

func (r recordMXResource) dataSourceName() string {
	return "dns_mx_records"
}
