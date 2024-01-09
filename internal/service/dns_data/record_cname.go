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

var _ recordResourceImplementor = &recordCNAMEResource{}
var _ recordDataSourceImplementor = &recordCNAMEResource{}

type cnameRecordModel struct {
	CName types.String `tfsdk:"cname"`
}

var cnameRecordAttrTypes = map[string]attr.Type{
	"cname": types.StringType,
}

type recordCNAMEResource struct{}

func NewRecordCNAMEResource() resource.Resource {
	return newRecordResource(&recordCNAMEResource{})
}

func NewRecordCNAMEDataSource() datasource.DataSource {
	return newRecordDataSource(&recordCNAMEResource{})
}

func (r recordCNAMEResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m cnameRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"cname": flex.ExpandString(m.CName),
	}

}

func (r recordCNAMEResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(cnameRecordAttrTypes)
	}
	t, d := types.ObjectValue(cnameRecordAttrTypes, map[string]attr.Value{
		"cname": flattenRDataFieldString(from["cname"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordCNAMEResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: cnameRecordAttrTypes}
	return attrTypes
}

func (r recordCNAMEResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. Always set to `CNAME`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"cname": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "A domain name which specifies the canonical or primary name for the owner. The owner name is an alias. Can be empty.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordCNAMEResource) recordType() string {
	return "CNAME"
}

func (r recordCNAMEResource) resourceName() string {
	return "dns_cname_record"
}

func (r recordCNAMEResource) dataSourceName() string {
	return "dns_cname_records"
}

func (r recordCNAMEResource) description() string {
	return "Represents a DNS CNAME resource record in an authoritative zone."
}
