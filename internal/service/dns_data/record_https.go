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

var _ recordResourceImplementor = &recordHTTPSResource{}

type httpsRecordModel struct {
	TargetName types.String `tfsdk:"target_name"`
}

var httpsRecordAttrTypes = map[string]attr.Type{
	"target_name": types.StringType,
}

type recordHTTPSResource struct{}

func NewRecordHTTPSResource() resource.Resource {
	return newRecordResource(&recordHTTPSResource{})
}

func NewRecordHTTPSDataSource() datasource.DataSource {
	return newRecordDataSource(&recordHTTPSResource{})
}

func (r recordHTTPSResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m httpsRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"target_name": flex.ExpandString(m.TargetName),
	}

}

func (r recordHTTPSResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(httpsRecordAttrTypes)
	}
	t, d := types.ObjectValue(httpsRecordAttrTypes, map[string]attr.Value{
		"target_name": flattenRDataFieldString(from["target_name"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordHTTPSResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: httpsRecordAttrTypes}
	return attrTypes
}

func (r recordHTTPSResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `HTTPS`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"target_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The target domain name of the HTTPS record.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}
func (r recordHTTPSResource) recordType() string {
	return "HTTPS"
}

func (r recordHTTPSResource) resourceName() string {
	return "dns_https_record"
}

func (r recordHTTPSResource) dataSourceName() string {
	return "dns_https_records"
}
