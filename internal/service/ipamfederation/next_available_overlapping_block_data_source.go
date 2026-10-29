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
	TagFilters            types.Map    `tfsdk:"tag_filters"`
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
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Block.",
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^federation/federated_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
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
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Key-value pairs to filter the candidate parent Federated Blocks by tags.",
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

	cidrData := data.Cidr.ValueInt64()
	count := data.OverlappingBlockCount.ValueInt64()
	tagFilters := data.TagFilters

	if len(tagFilters.Elements()) > 0 {
		// Find candidate parent Federated Blocks by tags
		tagFilterStr := flex.ExpandFrameworkMapTagFilterString(ctx, tagFilters, &resp.Diagnostics)

		allFederatedBlocks, err := FetchFederatedBlocksByTagFilter(ctx, d.client, tagFilterStr, &resp.Diagnostics)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Federated Blocks with tags, got error: %s", err))
			return
		}

		if len(allFederatedBlocks) == 0 {
			resp.Diagnostics.AddError("No Federated Blocks Found", "No federated blocks found with the given tags.")
			return
		}

		var findResults []ipamfederation.OverlappingBlock

		for _, fb := range allFederatedBlocks {
			if fb.Cidr != nil && *fb.Cidr >= cidrData {
				continue
			}
			findResultsLen := int64(len(findResults))
			if findResultsLen >= count {
				break
			}

			remainingCount := count - findResultsLen
			findResult, findErr := d.findOverlappingBlocks(ctx, *fb.Id, cidrData, remainingCount, data.Name, data.Comment)
			if findErr != nil {
				errorBody := []byte(findErr.Error())
				availableCount := int64(utils.ExtractAvailableCountFromError(errorBody))

				if availableCount > 0 {
					partialResult, retryErr := d.findOverlappingBlocks(ctx, *fb.Id, cidrData, availableCount, data.Name, data.Comment)
					if retryErr != nil {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error finding Overlapping Block after retry: %s", retryErr))
						return
					}
					findResults = append(findResults, partialResult...)
				} else {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NextAvailableOverlappingBlock, got error: %s", findErr))
					return
				}
				continue
			}

			if len(findResult) > 0 {
				findResults = append(findResults, findResult...)
			}
		}

		finalResultsCount := int64(len(findResults))
		if finalResultsCount < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Overlapping Blocks",
				fmt.Sprintf("Requested %d Overlapping Blocks with CIDR %d, but only %d were found. Not enough space available across all checked Federated Blocks.", count, cidrData, finalResultsCount),
			)
			return
		}
		data.FlattenResults(ctx, findResults, &resp.Diagnostics)
	} else {
		findResult, err := d.findOverlappingBlocks(ctx, data.Id.ValueString(), cidrData, count, data.Name, data.Comment)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NextAvailableOverlappingBlock, got error: %s", err))
			return
		}
		data.FlattenResults(ctx, findResult, &resp.Diagnostics)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// findOverlappingBlocks is a helper function to find next available overlapping blocks under a given parent Federated Block.
func (d *NextAvailableOverlappingBlockDataSource) findOverlappingBlocks(ctx context.Context, id string, cidr int64, count int64, name, comment types.String) ([]ipamfederation.OverlappingBlock, error) {
	req := func(c int64) ipamfederation.NextAvailableOverlappingBlockAPIListNextAvailableOverlappingBlocksRequest {
		apiReq := d.client.IPAMFederationAPI.
			NextAvailableOverlappingBlockAPI.
			ListNextAvailableOverlappingBlocks(ctx, id).
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
//
// NOTE: this is also defined by the ReservedBlock next_available_id PR (#316); once that PR
// merges, this duplicate definition should be removed in favor of the shared one.
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
