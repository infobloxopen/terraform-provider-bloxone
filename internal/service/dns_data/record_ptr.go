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

var _ recordResourceImplementor = &recordPTRResource{}
var _ recordDataSourceImplementor = &recordPTRResource{}

type ptrRecordModel struct {
	DName types.String `tfsdk:"dname"`
}

var ptrRecordAttrTypes = map[string]attr.Type{
	"dname": types.StringType,
}

var ptrRecordOptionsAttrTypes = map[string]attr.Type{
	"address": types.StringType,
}

type recordPTRResource struct{}

func NewRecordPTRResource() resource.Resource {
	return newRecordResource(&recordPTRResource{})
}

func NewRecordPTRDataSource() datasource.DataSource {
	return newRecordDataSource(&recordPTRResource{})
}

func (r recordPTRResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m ptrRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"dname": flex.ExpandString(m.DName),
	}
}

func (r recordPTRResource) expandOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	m := struct {
		Address types.String `tfsdk:"address"`
	}{}
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return map[string]interface{}{
		"address": flex.ExpandString(m.Address),
	}
}

func (r recordPTRResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ptrRecordAttrTypes)
	}
	t, d := types.ObjectValue(ptrRecordAttrTypes, map[string]attr.Value{
		"dname": flattenRDataFieldString(from["dname"], diags),
	})
	diags.Append(d...)
	return t
}

func (r recordPTRResource) flattenOptions(ctx context.Context, m types.Object, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ptrRecordOptionsAttrTypes)
	}
	t, d := types.ObjectValue(ptrRecordOptionsAttrTypes, map[string]attr.Value{
		"address": flex.FlattenString(from["address"].(string)),
	})
	diags.Append(d...)
	return t
}

func (r recordPTRResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: ptrRecordAttrTypes}
	attrTypes["options"] = types.ObjectType{AttrTypes: ptrRecordOptionsAttrTypes}
	return attrTypes
}

func (r recordPTRResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of DNS record. This is always `PTR`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"dname": schema.StringAttribute{
				Required: true,
			},
		},
		Required: true,
	}
	schemaAttrs["options"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				Computed: true,
				Optional: true,
				MarkdownDescription: "For GET operation it contains the IPv4 or IPv6 address represented by the PTR record.\n\n" +
					"For POST and PATCH operations it can be used to create/update a PTR record based on the IP address it represents. In this case, in addition to the _address_ in the options field, need to specify the _view_ field.",
			},
		},
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The DNS resource record type-specific non-protocol options.",
	}
	return schemaAttrs
}

func (r recordPTRResource) recordType() string {
	return "PTR"
}

func (r recordPTRResource) resourceName() string {
	return "dns_ptr_record"
}

func (r recordPTRResource) dataSourceName() string {
	return "dns_ptr_records"
}
