package ipamfederation

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableOverlappingBlockDataSource{}

func NewNextAvailableOverlappingBlockDataSource() datasource.DataSource {
	return &NextAvailableOverlappingBlockDataSource{}
}

// NextAvailableOverlappingBlockDataSource defines the data source implementation.
type NextAvailableOverlappingBlockDataSource struct {
	client *universalddiclient.APIClient
}

type NextAvailableOverlappingBlockModel struct {
	Id                    types.String `tfsdk:"id"`
	Cidr                  types.Int64  `tfsdk:"cidr"`
	OverlappingBlockCount types.Int64  `tfsdk:"overlapping_block_count"`
	Name                  types.String `tfsdk:"name"`
	Comment               types.String `tfsdk:"comment"`
	Results               types.List   `tfsdk:"results"`
}

func (m *NextAvailableOverlappingBlockModel) FlattenResults(ctx context.Context, from []ipamfederation.OverlappingBlock, diags *diag.Diagnostics) {
	var addresses []string
	for _, block := range from {
		if block.Address != nil {
			addresses = append(addresses, *block.Address)
		}
	}
	m.Results = flex.FlattenFrameworkListString(ctx, addresses, diags)
}

func (d *NextAvailableOverlappingBlockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_next_available_overlapping_blocks"
}

func (d *NextAvailableOverlappingBlockDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available __OverlappingBlock__ objects from the specified parent __FederatedBlock__, without allocating them.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Block.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^federation/federated_block/[0-9a-f-].*$`), "invalid resource ID specified"),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The CIDR of the overlapping blocks to be retrieved.",
			},
			"overlapping_block_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The count of overlapping blocks required. If not provided, it will default to 1.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The name to be provided.",
			},
			"comment": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description for the overlapping block. May contain 0 to 1024 characters. Can include UTF-8.",
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available overlapping block's addresses in the specified parent Federated Block.",
			},
		},
	}
}

func (d *NextAvailableOverlappingBlockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*universalddiclient.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *universalddiclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = client
}

func (d *NextAvailableOverlappingBlockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableOverlappingBlockModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.OverlappingBlockCount.IsNull() {
		data.OverlappingBlockCount = types.Int64Value(1)
	}

	// Uses the GET (preview) endpoint only, so reading this data source never allocates blocks.
	apiReq := d.client.IPAMFederationAPI.
		NextAvailableOverlappingBlockAPI.
		ListNextAvailableOverlappingBlocks(ctx, data.Id.ValueString()).
		Cidr(data.Cidr.ValueInt64()).
		Count(data.OverlappingBlockCount.ValueInt64())
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		apiReq = apiReq.Name(data.Name.ValueString())
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		apiReq = apiReq.Comment(data.Comment.ValueString())
	}

	apiRes, _, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NextAvailableOverlappingBlock, got error: %s", err))
		return
	}

	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
