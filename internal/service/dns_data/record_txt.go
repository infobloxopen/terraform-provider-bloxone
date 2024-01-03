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

var _ recordResourceImplementor = &recordTXTResource{}
var _ recordDataSourceImplementor = &recordTXTResource{}

type txtRecordModel struct {
	Text types.String `tfsdk:"text"`
}

var txtRecordAttrTypes = map[string]attr.Type{
	"text": types.StringType,
}

type recordTXTResource struct{}

func NewRecordTXTResource() resource.Resource {
	return newRecordResource(&recordTXTResource{})
}

func NewRecordTXTDataSource() datasource.DataSource {
	return newRecordDataSource(&recordTXTResource{})
}

func (r recordTXTResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m txtRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"text": flex.ExpandString(m.Text),
	}

}

func (r recordTXTResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(txtRecordAttrTypes)
	}
	t, d := types.ObjectValue(txtRecordAttrTypes, map[string]attr.Value{
		"text": flattenRDataFieldString(from["text"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordTXTResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: txtRecordAttrTypes}
	return attrTypes
}

func (r recordTXTResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `TXT`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"text": schema.StringAttribute{
				Required: true,
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordTXTResource) recordType() string {
	return "TXT"
}

func (r recordTXTResource) resourceName() string {
	return "dns_txt_record"
}

func (r recordTXTResource) dataSourceName() string {
	return "dns_txt_records"
}

func (r recordTXTResource) description() string {
	return "Represents a DNS TXT resource record in an authoritative DNS zone."
}
