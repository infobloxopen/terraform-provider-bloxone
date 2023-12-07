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

var _ recordResourceImplementor = &recordDNAMEResource{}

type dnameRecordModel struct {
	Target types.String `tfsdk:"target"`
}

var dnameRecordAttrTypes = map[string]attr.Type{
	"target": types.StringType,
}

type recordDNAMEResource struct{}

func NewRecordDNAMEResource() resource.Resource {
	return newRecordResource(&recordDNAMEResource{})
}

func NewRecordDNAMEDataSource() datasource.DataSource {
	return newRecordDataSource(&recordDNAMEResource{})
}

func (r recordDNAMEResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m dnameRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"target": flex.ExpandString(m.Target),
	}

}

func (r recordDNAMEResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(dnameRecordAttrTypes)
	}
	t, d := types.ObjectValue(dnameRecordAttrTypes, map[string]attr.Value{
		"target": flattenRDataFieldString(from["target"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordDNAMEResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: dnameRecordAttrTypes}
	return attrTypes
}

func (r recordDNAMEResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `DNAME`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"target": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The target domain name to which the zone will be mapped. Can be empty.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordDNAMEResource) recordType() string {
	return "DNAME"
}

func (r recordDNAMEResource) resourceName() string {
	return "dns_dname_record"
}

func (r recordDNAMEResource) dataSourceName() string {
	return "dns_dname_records"
}
