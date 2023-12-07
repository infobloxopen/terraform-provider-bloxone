package dns_data

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RecordDataSource{}

func newRecordDataSource(impl recordDataSourceImplementor) datasource.DataSource {
	return &RecordDataSource{
		impl: impl,
	}
}

// RecordDataSource defines the data source implementation.
type RecordDataSource struct {
	client *bloxoneclient.APIClient
	impl   recordDataSourceImplementor
}

func (d *RecordDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.impl.dataSourceName()
}

type DataRecordModelWithFilter struct {
	Filters    types.Map  `tfsdk:"filters"`
	TagFilters types.Map  `tfsdk:"tag_filters"`
	Results    types.List `tfsdk:"results"`
}

func (d *RecordDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"tag_filters": schema.MapAttribute{
				Description: "Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(d.impl.schemaAttributes(), &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *RecordDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*bloxoneclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *bloxoneclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *RecordDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DataRecordModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	filters := flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)
	// Add type filter by default
	if d.impl.recordType() != "Generic" {
		if len(filters) > 0 {
			filters = filters + " and "
		}
		filters = filters + "type=='" + d.impl.recordType() + "'"
	}

	apiRes, _, err := d.client.DNSDataAPI.
		RecordAPI.
		RecordList(ctx).
		Filter(filters).
		Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Record, got error: %s", err))
		return
	}

	if len(apiRes.GetResults()) == 0 {
		return
	}
	data.Results = flex.FlattenFrameworkListNestedBlock(ctx, apiRes.GetResults(), d.impl.attributeTypes(), &resp.Diagnostics, d.FlattenDataRecord)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *RecordDataSource) FlattenDataRecord(ctx context.Context, from *dns_data.DataRecord, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(d.impl.attributeTypes())
	}
	m := dataRecordModel{}
	m.Flatten(ctx, from, diags, d.impl)
	t, ds := types.ObjectValueFrom(ctx, d.impl.attributeTypes(), m)
	diags.Append(ds...)
	return t
}
