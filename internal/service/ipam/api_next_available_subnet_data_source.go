package ipam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableSubnetDataSource{}

func NewNextAvailableSubnetDataSource() datasource.DataSource {
	return &NextAvailableSubnetDataSource{}
}

// NextAvailableSubnetDataSource defines the data source implementation.
type NextAvailableSubnetDataSource struct {
	client *bloxoneclient.APIClient
}

type IpamsvcNextAvailableSubnetModel struct {
	Id         types.String            `tfsdk:"id"`
	Cidr       types.Int64             `tfsdk:"cidr"`
	Count      types.Int64             `tfsdk:"subnet_count"`
	Results    types.List              `tfsdk:"results"`
	TagFilters map[string]types.String `tfsdk:"tag_filters"`
}

func (m *IpamsvcNextAvailableSubnetModel) FlattenResults(ctx context.Context, from []ipam.Subnet, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	var listOfAddress []string

	for _, address := range from {
		listOfAddress = append(listOfAddress, types.StringValue(*address.Address).String())
	}
	m.Results = flex.FlattenFrameworkListString(ctx, listOfAddress, diags)
}

func (d *NextAvailableSubnetDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_next_available_subnets"
}

func (d *NextAvailableSubnetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available subnets in the specified address block.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: `An application specific resource identity of a resource.`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: `The cidr value of subnets to be created.`,
			},
			"subnet_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: `Number of subnets to generate. Default 1 if not set.`,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available subnet addresses in the specified resource.",
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Filter subnets by tags. Key-value pairs to match.",
			},
		},
	}
}

func (d *NextAvailableSubnetDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

// buildTagFilterString constructs a tag filter string from a map of tag key-value pairs
func (d *NextAvailableSubnetDataSource) buildTagFilterString(tagFilters map[string]types.String) string {
	if len(tagFilters) == 0 {
		return ""
	}

	filters := make([]string, 0, len(tagFilters))
	for k, v := range tagFilters {
		filters = append(filters, fmt.Sprintf("%s=='%s'", k, v.ValueString()))
	}

	return strings.Join(filters, " and ")
}

func (d *NextAvailableSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpamsvcNextAvailableSubnetModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Ensure subnet_count has a default value
	if data.Count.IsNull() {
		data.Count = types.Int64Value(1)
	}

	count := int32(data.Count.ValueInt64())

	// Check if tag filters are specified
	tagFilters := make(map[string]types.String)
	req.Config.GetAttribute(ctx, path.Root("tag_filters"), &tagFilters)

	if len(tagFilters) > 0 {
		// Find subnets by tags
		tagFilterStr := d.buildTagFilterString(tagFilters)

		var allAddressBlocks []ipam.AddressBlock
		const limit int32 = 1000
		offset := int32(0)

		// Fetch all address blocks matching the tag filters
		for {
			listAPI := d.client.IPAddressManagementAPI.AddressBlockAPI.
				List(ctx).
				Tfilter(tagFilterStr).
				Offset(offset).
				Limit(limit).
				Inherit("full")

			apiRes, _, err := listAPI.Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to list Address Blocks by tags, got error: %s", err))
				return
			}

			results := apiRes.GetResults()
			allAddressBlocks = append(allAddressBlocks, results...)

			// If we got fewer results than the limit, we've reached the end
			if int32(len(results)) < limit {
				break
			}

			offset += limit
		}

		if len(allAddressBlocks) == 0 {
			resp.Diagnostics.AddError("No Address Blocks Found", "No address blocks found with the given tags.")
			return
		}

		var findResults []ipam.Subnet
		for _, ab := range allAddressBlocks {
			if *ab.Cidr >= *flex.ExpandInt64Pointer(data.Cidr) {
				continue
			}
			if int32(len(findResults)) >= count {
				break
			}

			remainingCount := count - int32(len(findResults))
			findResult, err := d.findSubnet(ctx, *ab.Id, int32(data.Cidr.ValueInt64()), remainingCount)
			if err != nil {
				// Check if the error contains relevant information about available blocks
				if strings.Contains(err.Error(), "available networks") {
					// Extract error message body for parsing
					errorMsg := err.Error()
					// Try to extract the body from the error message
					startIdx := strings.Index(errorMsg, "{")
					if startIdx != -1 {
						errorBody := []byte(errorMsg[startIdx:])
						availableCount := d.extractAvailableCountFromError(errorBody)
						if availableCount > 0 {
							// Retry with the available count
							partialResult, retryErr := d.findSubnet(ctx, *ab.Id, int32(data.Cidr.ValueInt64()), availableCount)
							if retryErr != nil {
								resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error finding address block after retry: %s", retryErr))
								return
							}
							findResults = append(findResults, partialResult...)
						}
						continue
					}
				}
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error finding address block: %s", err))
				return
			}

			if len(findResult) > 0 {
				findResults = append(findResults, findResult...)
			}
		}
		if int32(len(findResults)) < count {
			resp.Diagnostics.AddError(
				"Insufficient Available Subnets",
				fmt.Sprintf("Requested %d subnets with CIDR %d, but only %d were found. Not enough subnets available across all checked address blocks.", count, data.Cidr.ValueInt64(), len(findResults)),
			)
			return
		}
		data.FlattenResults(ctx, findResults, &resp.Diagnostics)
	} else {
		// Use original next available address logic to find by ID
		apiRes, _, err := d.client.IPAddressManagementAPI.
			AddressBlockAPI.
			ListNextAvailableSubnet(ctx, data.Id.ValueString()).
			Cidr(int32(data.Cidr.ValueInt64())).
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

// Helper function to find subnets by ID and count
func (d *NextAvailableSubnetDataSource) findSubnet(ctx context.Context, id string, cidr int32, count int32) ([]ipam.Subnet, error) {
	apiRes, httpRes, err := d.client.IPAddressManagementAPI.AddressBlockAPI.
		ListNextAvailableSubnet(ctx, id).
		Cidr(cidr).
		Count(count).
		Execute()
	if err != nil {
		// Check for 400 status code without relying on specific error type
		if httpRes != nil && httpRes.StatusCode == 400 {
			// Convert response body to string if it's available from httpRes
			bodyBytes, _ := io.ReadAll(httpRes.Body)
			errMsg := httpRes.Body.Close()
			if errMsg != nil {
				return nil, errMsg
			} // Close the body after reading

			// Try to extract available count
			availableCount := d.extractAvailableCountFromError(bodyBytes)
			if availableCount > 0 {
				// Retry with the available count
				retryRes, _, retryErr := d.client.IPAddressManagementAPI.AddressBlockAPI.
					ListNextAvailableSubnet(ctx, id).
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

func (d *NextAvailableSubnetDataSource) extractAvailableCountFromError(body []byte) int32 {
	var errorResponse struct {
		Error []struct {
			Message string `json:"message"`
		} `json:"error"`
	}

	// Parse the JSON error body
	if err := json.Unmarshal(body, &errorResponse); err != nil {
		return 0
	}

	// Extract the available count from the error message
	for _, err := range errorResponse.Error {
		if strings.Contains(err.Message, "The available networks are:") {
			// Use regex to extract the number after "The available networks are: "
			re := regexp.MustCompile(`The available networks are: (\d+)`)
			match := re.FindStringSubmatch(err.Message)
			if len(match) > 1 {
				count, parseErr := strconv.Atoi(match[1])
				if parseErr == nil {
					return int32(count)
				}
			}
		}
	}

	return 0
}
