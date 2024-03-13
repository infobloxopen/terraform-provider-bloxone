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

var _ recordResourceImplementor = &recordAAAAResource{}
var _ recordDataSourceImplementor = &recordAAAAResource{}

type aaaaRecordModel struct {
	Address types.String `tfsdk:"address"`
}

var aaaaRecordAttrTypes = map[string]attr.Type{
	"address": types.StringType,
}

type recordAAAAResource struct{}

func NewRecordAAAAResource() resource.Resource {
	return newRecordResource(&recordAAAAResource{})
}

func NewRecordAAAADataSource() datasource.DataSource {
	return newRecordDataSource(&recordAAAAResource{})
}

func (r recordAAAAResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m aaaaRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"address": flex.ExpandString(m.Address),
	}

}

func (r recordAAAAResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(aaaaRecordAttrTypes)
	}
	t, d := types.ObjectValue(aaaaRecordAttrTypes, map[string]attr.Value{
		"address": flattenRDataFieldString(from["address"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordAAAAResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: aaaaRecordAttrTypes}
	return attrTypes
}

func (r recordAAAAResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `AAAA`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The IPv6 address of the host.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordAAAAResource) recordType() string {
	return "AAAA"
}

func (r recordAAAAResource) resourceName() string {
	return "dns_aaaa_record"
}

func (r recordAAAAResource) dataSourceName() string {
	return "dns_aaaa_records"
}
