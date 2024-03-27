package fw

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PoPRegionsDataSource{}

func NewPoPRegionsDataSource() datasource.DataSource {
	return &PoPRegionsDataSource{}
}

// PoPRegionsDataSource defines the data source implementation.
type PoPRegionsDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *PoPRegionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "td_pop_regions"
}

type AtcfwPoPRegionModelWithFilter struct {
	Filters types.Map  `tfsdk:"filters"`
	Results types.List `tfsdk:"results"`
}

func (m *AtcfwPoPRegionModelWithFilter) FlattenResults(ctx context.Context, from []fw.AtcfwPoPRegion, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, AtcfwPoPRegionAttrTypes, diags, FlattenAtcfwPoPRegion)
}

func (d *PoPRegionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieve information on all Point of Presence (PoP) regions available in the database.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(AtcfwPoPRegionResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *PoPRegionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *PoPRegionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AtcfwPoPRegionModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := d.client.FWAPI.
		PopRegionsAPI.
		PopRegionsListPoPRegions(ctx).
		Filter(flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Pop Regions, got error: %s", err))
		return
	}

	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
