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

var _ recordResourceImplementor = &recordSRVResource{}
var _ recordDataSourceImplementor = &recordSRVResource{}

type srvRecordModel struct {
	Port     types.Int64  `tfsdk:"port"`
	Priority types.Int64  `tfsdk:"priority"`
	Target   types.String `tfsdk:"target"`
	Weight   types.Int64  `tfsdk:"weight"`
}

var srvRecordAttrTypes = map[string]attr.Type{
	"port":     types.Int64Type,
	"priority": types.Int64Type,
	"target":   types.StringType,
	"weight":   types.Int64Type,
}

type recordSRVResource struct{}

func NewRecordSRVResource() resource.Resource {
	return newRecordResource(&recordSRVResource{})
}

func NewRecordSRVDataSource() datasource.DataSource {
	return newRecordDataSource(&recordSRVResource{})
}

func (r recordSRVResource) expandRData(ctx context.Context, o types.Object, diags *diag.Diagnostics) map[string]interface{} {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m srvRecordModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	rdata := map[string]interface{}{
		"port":     flex.ExpandInt64(m.Port),
		"priority": flex.ExpandInt64(m.Priority),
		"target":   flex.ExpandString(m.Target),
	}

	// Optional fields
	if !m.Weight.IsNull() && !m.Weight.IsUnknown() {
		rdata["weight"] = flex.ExpandInt64(m.Weight)
	}
	return rdata
}

func (r recordSRVResource) flattenRData(_ context.Context, from map[string]interface{}, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(srvRecordAttrTypes)
	}
	t, d := types.ObjectValue(srvRecordAttrTypes, map[string]attr.Value{
		"port":     flattenRDataFieldInt64(from["port"], true, diags),
		"priority": flattenRDataFieldInt64(from["priority"], true, diags),
		"target":   flattenRDataFieldString(from["target"], diags),
		"weight":   flattenRDataFieldInt64(from["weight"], true, diags),
	})
	diags.Append(d...)
	return t
}

func (r recordSRVResource) attributeTypes() map[string]attr.Type {
	attrTypes := recordCommonAttrTypes()
	attrTypes["rdata"] = types.ObjectType{AttrTypes: srvRecordAttrTypes}
	return attrTypes
}

func (r recordSRVResource) schemaAttributes() map[string]schema.Attribute {
	schemaAttrs := recordCommonSchema()
	schemaAttrs["type"] = schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the record. This is always `SRV`.",
	}
	schemaAttrs["rdata"] = schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"port": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "An unsigned 16-bit integer which specifies the port on this target host of this service. The range of the value is 0 to 65535. This is often as specified in Assigned Numbers but need not be.",
			},
			"priority": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "An unsigned 16-bit integer which specifies the priority of this target host. The range of the value is 0 to 65535. A client must attempt to contact the target host with the lowest-numbered priority it can reach. Target hosts with the same priority should be tried in an order defined by the weight field.",
			},
			"target": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The domain name of the target host. There must be one or more address records for this name, the name must not be an alias (in the sense of RFC 1034 or RFC 2181).\n\nA target of “.” means that the service is decidedly not available at this domain.",
			},
			"weight": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "An unsigned 16-bit integer which specifies a relative weight for entries with the same priority. The range of the value is 0 to 65535. Larger weights should be given a proportionately higher probability of being selected. Domain administrators should use weight 0 when there isn’t any server selection to do, to make the RR easier to read for humans (less noisy). In the presence of records containing weights greater than 0, records with weight 0 should have a very small chance of being selected.\n\nIn the absence of a protocol whose specification calls for the use of other weighting information, a client arranges the SRV RRs of the same priority in the order in which target hosts, specified by the SRV RRs, will be contacted.\n\nDefaults to 0.",
			},
		},
		Required: true,
	}
	return schemaAttrs
}

func (r recordSRVResource) recordType() string {
	return "SRV"
}

func (r recordSRVResource) resourceName() string {
	return "dns_srv_record"
}

func (r recordSRVResource) dataSourceName() string {
	return "dns_srv_records"
}

func (r recordSRVResource) description() string {
	return "Represents a DNS SRV resource record in an authoritative zone."
}
