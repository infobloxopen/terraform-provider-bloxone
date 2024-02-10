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

var _ recordResourceImplementor = &recordSVCBResource{}
var _ recordDataSourceImplementor = &recordSVCBResource{}

type svcbRecordModel struct {
	TargetName types.String `tfsdk:"target_name"`
}

var svcbRecordAttrTypes = map[string]attr.Type{
	"target_name": types.StringType,
}

type recordSVCBResource struct{}

func NewRecordSVCBResource() resource.Resource {
	return newRecordResource(&recordSVCBResource{})
}

func NewRecordSVCBDataSource() datasource.DataSource {
	return newRecordDataSource(&recordSVCBResource{})
}

func (r recordSVCBResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m svcbRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"target_name": flex.ExpandString(m.TargetName),
	}
}

func (r recordSVCBResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(svcbRecordAttrTypes)
	}
	t, d := types.ObjectValue(svcbRecordAttrTypes, map[string]attr.Value{
		"target_name": flattenRDataFieldString(from["target_name"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordSVCBResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: svcbRecordAttrTypes}
	return attrTypes
}

func (r recordSVCBResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `SVCB`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"target_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The target domain name of the SVCB record.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordSVCBResource) recordType() string {
	return "SVCB"
}

func (r recordSVCBResource) resourceName() string {
	return "dns_svcb_record"
}

func (r recordSVCBResource) dataSourceName() string {
	return "dns_svcb_records"
}
