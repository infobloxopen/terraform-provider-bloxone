package ipamfederation

import (
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableFederatedBlockDataSource{}

func NewNextAvailableFederatedBlockDataSource() datasource.DataSource {
	return &NextAvailableFederatedBlockDataSource{}
}

// NextAvailableFederatedBlockDataSource defines the data source implementation.
type NextAvailableFederatedBlockDataSource struct {
	client *universalddiclient.APIClient
}

type NextAvailableFederatedBlockModel struct {
	Id                  types.String `tfsdk:"id"`
	Cidr                types.Int64  `tfsdk:"cidr"`
	FederatedBlockCount types.Int64  `tfsdk:"federated_block_count"`
	Name                types.String `tfsdk:"name"`
	Comment             types.String `tfsdk:"comment"`
	Results             types.List   `tfsdk:"results"`
	TagFilters          types.Map    `tfsdk:"tag_filters"`
}

func (m *NextAvailableFederatedBlockModel) FlattenResults(ctx context.Context, from []ipamfederation.FederatedBlock, diags *diag.Diagnostics) {
	var addresses []string
	for _, block := range from {
		if block.Address != nil {
			addresses = append(addresses, *block.Address)
		}
	}
	m.Results = flex.FlattenFrameworkListString(ctx, addresses, diags)
}

func (d *NextAvailableFederatedBlockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_next_available_federated_blocks"
}

func (d *NextAvailableFederatedBlockDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available __FederatedBlock__ objects from the specified parent __FederatedBlock__, without allocating them.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Block.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^federation/federated_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
					stringvalidator.AtLeastOneOf(path.MatchRoot("id"), path.MatchRoot("tag_filters")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The CIDR of the federated blocks to be retrieved.",
			},
			"federated_block_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The count of federated blocks required. If not provided, it will default to 1.",
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
				MarkdownDescription: "The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8.",
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available federated block's addresses in the specified parent Federated Block.",
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Key-value pairs to filter the candidate parent Federated Blocks by tags.",
			},
		},
	}
}

func (d *NextAvailableFederatedBlockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableFederatedBlockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableFederatedBlockModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.FederatedBlockCount.IsNull() {
		data.FederatedBlockCount = types.Int64Value(1)
	}

	cidrData := data.Cidr.ValueInt64()
	count := data.FederatedBlockCount.ValueInt64()
	tagFilters := data.TagFilters

	if len(tagFilters.Elements()) > 0 {
		// Find candidate parent Federated Blocks by tags
		tagFilterStr := flex.ExpandFrameworkMapTagFilterString(ctx, tagFilters, &resp.Diagnostics)

		allFederatedBlocks, err := FetchFederatedBlocksByTagFilter(ctx, d.client, tagFilterStr, &resp.Diagnostics)
		if err != nil {
			// FetchFederatedBlocksByTagFilter already recorded the error in resp.Diagnostics.
			return
		}

		if len(allFederatedBlocks) == 0 {
			resp.Diagnostics.AddError("No Federated Blocks Found", "No federated blocks found with the given tags.")
			return
		}

		var findResults []ipamfederation.FederatedBlock

		for _, fb := range allFederatedBlocks {
			if fb.Cidr != nil && *fb.Cidr >= cidrData {
				continue
			}
			findResultsLen := int64(len(findResults))
			if findResultsLen >= count {
				break
			}

			remainingCount := count - findResultsLen
			findResult, findErr := d.findFederatedBlocks(ctx, *fb.Id, cidrData, remainingCount, data.Name, data.Comment)
			if findErr != nil {
				// findFederatedBlocks already retried internally with any available count the API
				// reported. This parent has no (or insufficient) space; move on to the next
				// candidate parent instead of failing the whole search.
				continue
			}

			if len(findResult) > 0 {
				findResults = append(findResults, findResult...)
			}
		}

		finalResultsCount := int64(len(findResults))
		if finalResultsCount < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Federated Blocks",
				fmt.Sprintf("Requested %d Federated Blocks with CIDR %d, but only %d were found. Not enough space available across all checked Federated Blocks.", count, cidrData, finalResultsCount),
			)
			return
		}
		data.FlattenResults(ctx, findResults, &resp.Diagnostics)
	} else {
		findResult, err := d.findFederatedBlocks(ctx, data.Id.ValueString(), cidrData, count, data.Name, data.Comment)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NextAvailableFederatedBlock, got error: %s", err))
			return
		}
		data.FlattenResults(ctx, findResult, &resp.Diagnostics)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// findFederatedBlocks is a helper function to find next available federated blocks under a given parent Federated Block.
func (d *NextAvailableFederatedBlockDataSource) findFederatedBlocks(ctx context.Context, id string, cidr int64, count int64, name, comment types.String) ([]ipamfederation.FederatedBlock, error) {
	req := func(c int64) ipamfederation.NextAvailableFederatedBlockAPIListNextAvailableFederatedBlocksRequest {
		apiReq := d.client.IPAMFederationAPI.
			NextAvailableFederatedBlockAPI.
			ListNextAvailableFederatedBlocks(ctx, id).
			Cidr(cidr).
			Count(c)

		if !name.IsNull() {
			apiReq = apiReq.Name(name.ValueString())
		}
		if !comment.IsNull() {
			apiReq = apiReq.Comment(comment.ValueString())
		}
		return apiReq
	}

	apiRes, httpRes, err := req(count).Execute()
	if err != nil {
		// Check for 400 status code without relying on specific error type
		if httpRes != nil && httpRes.StatusCode == 400 {
			bodyBytes, _ := io.ReadAll(httpRes.Body)
			closeErr := httpRes.Body.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			// Try to extract available count and retry with it
			availableCount := int64(utils.ExtractAvailableCountFromError(bodyBytes))
			if availableCount > 0 {
				retryRes, _, retryErr := req(availableCount).Execute()
				if retryErr == nil {
					return retryRes.GetResults(), nil
				}
			}
		}
		return nil, err
	}

	return apiRes.GetResults(), nil
}

// FetchFederatedBlocksByTagFilter is a helper function used to fetch federated blocks by tags.
// Also used by the ReservedBlock and OverlappingBlock next-available data sources.
func FetchFederatedBlocksByTagFilter(ctx context.Context, client *universalddiclient.APIClient, tagFilterStr string, diagnostics *diag.Diagnostics) ([]ipamfederation.FederatedBlock, error) {
	federatedBlocks, err := utils.ReadWithPages(func(offset, limit int32) ([]ipamfederation.FederatedBlock, error) {
		apiRes, _, err := client.IPAMFederationAPI.FederatedBlockAPI.
			List(ctx).
			Tfilter(tagFilterStr).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			if diagnostics != nil {
				diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Federated Blocks, got error: %s", err))
			}
			return nil, err
		}
		return apiRes.GetResults(), nil
	})

	return federatedBlocks, err
}
