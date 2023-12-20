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

var _ recordResourceImplementor = &recordNSResource{}
var _ recordDataSourceImplementor = &recordNSResource{}

type nsRecordModel struct {
	DName types.String `tfsdk:"dname"`
}

var nsRecordAttrTypes = map[string]attr.Type{
	"dname": types.StringType,
}

type recordNSResource struct{}

func NewRecordNSResource() resource.Resource {
	return newRecordResource(&recordNSResource{})
}

func NewRecordNSDataSource() datasource.DataSource {
	return newRecordDataSource(&recordNSResource{})
}

func (r recordNSResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m nsRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"dname": flex.ExpandString(m.DName),
	}

}

func (r recordNSResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(nsRecordAttrTypes)
	}
	t, d := types.ObjectValue(nsRecordAttrTypes, map[string]attr.Value{
		"dname": flattenRDataFieldString(from["dname"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordNSResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: nsRecordAttrTypes}
	return attrTypes
}

func (r recordNSResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `NS`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"dname": schema.StringAttribute{
				Required: true,
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordNSResource) recordType() string {
	return "NS"
}

func (r recordNSResource) resourceName() string {
	return "dns_ns_record"
}

func (r recordNSResource) dataSourceName() string {
	return "dns_ns_records"
}

func (r recordNSResource) description() string {
	return "Represents a DNS NS resource record in an authoritative zone."
}
