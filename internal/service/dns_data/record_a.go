package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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

var aRecordOptionsAttrTypes = map[string]attr.Type{
	"create_ptr": types.BoolType,
	"check_rmz":  types.BoolType,
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

func (r recordAResource) expandOptions(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	m := struct {
		CreatePTR types.Bool `tfsdk:"create_ptr"`
		CheckRMZ  types.Bool `tfsdk:"check_rmz"`
	}{}
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return map[string]interface{}{
		"create_ptr": flex.ExpandBool(m.CreatePTR),
		"check_rmz":  flex.ExpandBool(m.CheckRMZ),
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

func (r recordAResource) flattenOptions(ctx context.Context, m types.Object, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	// Preserve the state value. The API doesn't return options in the response.
	// If the state value is unknown, set null value.
	if !m.IsUnknown() && !m.IsNull() {
		return m
	}
	return types.ObjectNull(aRecordOptionsAttrTypes)
}

func (r recordAResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: aRecordDataAttrTypes}
	attrTypes["options"] = types.ObjectType{AttrTypes: aRecordOptionsAttrTypes}
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
	schemaAttrs["options"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"create_ptr": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				MarkdownDescription: "A boolean flag which can be set to _true_ for POST operation to automatically create the corresponding PTR record.",
			},
			"check_rmz": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "A boolean flag which can be set to _true_ for POST operation to check the existence of reverse zone for creating the corresponding PTR record. Only applicable if the _create_ptr_ option is set to _true_.",
			},
		},
		Optional:            true,
		MarkdownDescription: "The DNS resource record type-specific non-protocol options.",
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
