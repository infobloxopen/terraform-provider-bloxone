package anycast

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &OnPremAnycastManagerDataSource{}

func NewOnPremAnycastManagerDataSource() datasource.DataSource {
	return &OnPremAnycastManagerDataSource{}
}

// OnPremAnycastManagerDataSource defines the data source implementation.
type OnPremAnycastManagerDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *OnPremAnycastManagerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "anycast_configs"
}

type ProtoAnycastConfigModelWithFilter struct {
	Filters     types.Map    `tfsdk:"filters"` //todo : remove this
	TagFilters  types.Map    `tfsdk:"tag_filters"`
	Results     types.List   `tfsdk:"results"`
	Service     types.String `tfsdk:"service"`
	HostID      types.Int64  `tfsdk:"host_id"`
	IsConfigued types.Bool   `tfsdk:"is_configured"`
}

func (m *ProtoAnycastConfigModelWithFilter) FlattenResults(ctx context.Context, from []anycast.AnycastConfig, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, ProtoAnycastConfigAttrTypes, diags, FlattenProtoAnycastConfig)
}

func (d *OnPremAnycastManagerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieve all named anycast configurations for the account.",
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
					Attributes: utils.DataSourceAttributeMap(ProtoAnycastConfigResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
			"service": schema.StringAttribute{
				Description: "Service name.",
				Optional:    true,
			},
			"host_id": schema.Int64Attribute{
				Description: "Host ID.",
				Optional:    true,
			},
			"is_configured": schema.BoolAttribute{
				Description: "Is configured.",
				Optional:    true,
			},
		},
	}
}

func (d *OnPremAnycastManagerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *OnPremAnycastManagerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProtoAnycastConfigModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := d.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		GetAnycastConfigList(ctx).
		Service(flex.ExpandString(data.Service)).
		HostId(flex.ExpandInt64(data.HostID)).
		IsConfigured(flex.ExpandBool(data.IsConfigued)).
		Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read OnPremAnycastManager, got error: %s", err))
		return
	}

	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
