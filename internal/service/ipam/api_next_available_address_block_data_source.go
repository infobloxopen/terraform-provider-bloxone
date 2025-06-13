package ipam

import (
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableAddressBlockDataSource{}

func NewNextAvailableAddressBlockDataSource() datasource.DataSource {
	return &NextAvailableAddressBlockDataSource{}
}

// NextAvailableAddressBlockDataSource defines the data source implementation.
type NextAvailableAddressBlockDataSource struct {
	client *bloxoneclient.APIClient
}

type IpamsvcNextAvailableAddressBlockModel struct {
	Id         types.String `tfsdk:"id"`
	Cidr       types.Int64  `tfsdk:"cidr"`
	Count      types.Int32  `tfsdk:"address_block_count"`
	Results    types.List   `tfsdk:"results"`
	TagFilters types.Map    `tfsdk:"tag_filters"`
}

func (m *IpamsvcNextAvailableAddressBlockModel) FlattenResults(ctx context.Context, from []ipam.AddressBlock, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	var listOfAddress []string

	for _, address := range from {
		listOfAddress = append(listOfAddress, types.StringValue(*address.Address).String())
	}
	m.Results = flex.FlattenFrameworkListString(ctx, listOfAddress, diags)
}

func (d *NextAvailableAddressBlockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_next_available_address_blocks"
}

func (d *NextAvailableAddressBlockDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available address blocks in the specified address block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: `An application specific resource identity of a resource.`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: `The cidr value of address blocks to be created.`,
			},
			"address_block_count": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: `Number of address blocks to generate. Default 1 if not set.`,
				Validators: []validator.Int32{
					int32validator.AtLeast(1),
				},
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available address block's addresses in the specified resource.",
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Key-value pairs to filter address blocks by tags.",
			},
		},
	}
}

func (d *NextAvailableAddressBlockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableAddressBlockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpamsvcNextAvailableAddressBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	if data.Count.IsNull() {
		data.Count = types.Int32Value(1)
	}
	count := data.Count.ValueInt32()
	cidrData := data.Cidr.ValueInt64()
	tagFilters := data.TagFilters // Check if tag filters are specified
	if len(tagFilters.Elements()) > 0 {
		// Find address blocks by tags
		tagFilterStr := flex.ExpandFrameworkMapFilterString(ctx, tagFilters, &resp.Diagnostics)

		var allAddressBlocks []ipam.AddressBlock

		allAddressBlocks, err := FetchAddressBlocksByTagFilter(ctx, d.client, tagFilterStr, &resp.Diagnostics)

		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Address Blocks with tags, got error: %s", err))

			return
		}

		if len(allAddressBlocks) == 0 {
			resp.Diagnostics.AddError("No Address Blocks Found", "No address blocks found with the given tags.")
			return
		}

		var findResults []ipam.AddressBlock

		for _, ab := range allAddressBlocks {
			if *ab.Cidr >= cidrData {
				continue
			}
			findResultsLen := int32(len(findResults))
			if findResultsLen >= count {
				break
			}

			remainingCount := count - findResultsLen
			findResult, findErr := d.findAddressBlock(ctx, *ab.Id, int32(cidrData), remainingCount)
			if findErr != nil {
				// Check if the error contains relevant information about available blocks
				errorBody := []byte(findErr.Error())
				availableCount := utils.ExtractAvailableCountFromError(errorBody)

				if availableCount > 0 {
					// Retry with the available count
					partialResult, retryErr := d.findAddressBlock(ctx, *ab.Id, int32(cidrData), availableCount)
					if retryErr != nil {
						resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error finding Address Block after retry: %s", retryErr))
						return
					}
					findResults = append(findResults, partialResult...)
				} else {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Address Block, got error: %s", findErr))
					return
				}
				continue
			}

			if len(findResult) > 0 {
				findResults = append(findResults, findResult...)
			}
		}
		finalResultsCount := int32(len(findResults))
		if finalResultsCount < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Address Blocks",
				fmt.Sprintf("Requested %d Address Blocks with CIDR %d, but only %d were found. Not enough Address Blocks available across all checked address blocks.", count, cidrData, finalResultsCount),
			)
			return
		}
		data.FlattenResults(ctx, findResults, &resp.Diagnostics)
	} else {
		// Use original next available address logic
		apiRes, _, err := d.client.IPAddressManagementAPI.
			AddressBlockAPI.
			ListNextAvailableAB(ctx, data.Id.ValueString()).
			Cidr(int32(cidrData)).
			Count(count).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AddressBlock Next Available Address Block API, got error: %s", err))
			return
		}

		data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Helper funtion to fecth address blocks by tag
func FetchAddressBlocksByTagFilter(ctx context.Context, client *bloxoneclient.APIClient, tagFilterStr string, diagnostics *diag.Diagnostics) ([]ipam.AddressBlock, error) {

	addressBlocks, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.AddressBlock, error) {
		apiRes, _, err := client.IPAddressManagementAPI.AddressBlockAPI.
			List(ctx).
			Tfilter(tagFilterStr).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			if diagnostics != nil {
				diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Address Blocks, got error: %s", err))
			}
			return nil, err
		}
		return apiRes.GetResults(), nil
	})

	return addressBlocks, err
}

// Helper function to find address blocks by ID and count
func (d *NextAvailableAddressBlockDataSource) findAddressBlock(ctx context.Context, id string, cidr int32, count int32) ([]ipam.AddressBlock, error) {
	apiRes, httpRes, err := d.client.IPAddressManagementAPI.AddressBlockAPI.
		ListNextAvailableAB(ctx, id).
		Cidr(cidr).
		Count(count).
		Execute()
	if err != nil {
		// Check for 400 status code without relying on specific error type
		if httpRes != nil && httpRes.StatusCode == 400 {
			// Convert response body to string if it's available from httpRes
			bodyBytes, _ := io.ReadAll(httpRes.Body)
			errMsg := httpRes.Body.Close() // Close the body after reading
			if errMsg != nil {
				return nil, errMsg
			}
			// Try to extract available count
			availableCount := utils.ExtractAvailableCountFromError(bodyBytes)
			if availableCount > 0 {
				// Retry with the available count
				retryRes, _, retryErr := d.client.IPAddressManagementAPI.AddressBlockAPI.
					ListNextAvailableAB(ctx, id).
					Cidr(cidr).
					Count(availableCount).
					Execute()
				if retryErr == nil {
					return retryRes.GetResults(), nil
				}
			}
		}
		return nil, err
	}

	return apiRes.GetResults(), nil
}
