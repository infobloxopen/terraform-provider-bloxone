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

var _ recordResourceImplementor = &recordAResource{}
var _ recordDataSourceImplementor = &recordAResource{}

type aRecordModel struct {
	Address types.String `tfsdk:"address"`
}

var aRecordDataAttrTypes = map[string]attr.Type{
	"address": types.StringType,
}

type recordAResource struct{}

func NewRecordAResource() resource.Resource {
	return newRecordResource(&recordAResource{})
}

func NewRecordADataSource() datasource.DataSource {
	return newRecordDataSource(&recordAResource{})
}

func (r recordAResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m aRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"address": flex.ExpandString(m.Address),
	}

}

func (r recordAResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(aRecordDataAttrTypes)
	}
	t, d := types.ObjectValue(aRecordDataAttrTypes, map[string]attr.Value{
		"address": flattenRDataFieldString(from["address"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordAResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: aRecordDataAttrTypes}
	return attrTypes
}

func (r recordAResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the record. This is always `A`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The IPv4 address of the host",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordAResource) recordType() string {
	return "A"
}

func (r recordAResource) resourceName() string {
	return "dns_a_record"
}

func (r recordAResource) dataSourceName() string {
	return "dns_a_records"
}
