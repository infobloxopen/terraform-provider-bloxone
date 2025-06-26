package ipam

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
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
var _ datasource.DataSource = &NextAvailableIPDataSource{}

func NewIpamNextAvailableIPDataSource() datasource.DataSource {
	return &NextAvailableIPDataSource{}
}

// NextAvailableIPDataSource defines the data source implementation.

type NextAvailableIPDataSource struct {
	client *bloxoneclient.APIClient
}

type IpamsvcNextAvailableIPModel struct {
	Id           types.String `tfsdk:"id"`
	Contiguous   types.Bool   `tfsdk:"contiguous"`
	Count        types.Int64  `tfsdk:"ip_count"`
	Results      types.List   `tfsdk:"results"`
	TagFilters   types.Map    `tfsdk:"tag_filters"`
	ResourceType types.String `tfsdk:"resource_type"`
}

func (m *IpamsvcNextAvailableIPModel) FlattenResults(ctx context.Context, from []ipam.Address, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	var listOfAddress []string

	for _, address := range from {
		listOfAddress = append(listOfAddress, address.Address)
	}
	m.Results = flex.FlattenFrameworkListString(ctx, listOfAddress, diags)
}

func (d *NextAvailableIPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableIPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_next_available_ips"
}

func (d *NextAvailableIPDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available IP addresses in the specified resource. The resource can be an address block, subnet or range.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: `An application specific resource identity of a resource`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(range|subnet|address_block)/[0-9a-f-].*$`), "invalid resource ID specified"),
					stringvalidator.ConflictsWith(path.MatchRoot("tag_filters")),
					stringvalidator.ConflictsWith(path.MatchRoot("resource_type")),
				},
			},
			// Query parameter
			"contiguous": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: `Indicates whether the IP addresses should belong to a contiguous block. Defaults to false.`,
			},
			// Query parameter
			"ip_count": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 20),
				},
				MarkdownDescription: `The number of IP addresses requested. Defaults to 1.`,
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of next available IP addresses in the specified resource",
			},
			"tag_filters": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "Map of tag key/value pairs to filter resources",
				Validators: []validator.Map{
					mapvalidator.AlsoRequires(path.MatchRoot("resource_type")),
				},
			},
			"resource_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Resource type to search when using tag filters (address_block, subnet, or range)",
				Validators: []validator.String{
					stringvalidator.OneOf("address_block", "subnet", "range"),
					stringvalidator.AlsoRequires(path.MatchRoot("tag_filters")),
				},
			},
		},
	}
}

func (d *NextAvailableIPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpamsvcNextAvailableIPModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Count.IsNull() {
		data.Count = types.Int64Value(1)
	}
	count := data.Count.ValueInt64()

	// Validate count is within allowed range (1-20)
	if count < 1 || count > 20 {
		resp.Diagnostics.AddError("Invalid count", "Count must be between 1 and 20")
		return
	}

	// Check if we're using tag_filters or direct resource ID
	if !data.TagFilters.IsNull() {
		// Using tag filters - ensure resource_type is specified
		if data.ResourceType.IsNull() {
			resp.Diagnostics.AddError("Missing resource_type", "resource_type is required when using tag_filters")
			return
		}

		// Get resources by tags and find next available IPs
		addresses, err := d.findNextAvailableIPsByTags(ctx, data)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to find resources by tags: %s", err))
			return
		}

		if len(addresses) < int(count) {
			resp.Diagnostics.AddError("Insufficient IPs", fmt.Sprintf("Not enough available IPs found in %s with the given tags", data.ResourceType.ValueString()))
			return
		}

		// Limit to requested count
		if len(addresses) > int(count) {
			addresses = addresses[:count]
		}

		data.FlattenResults(ctx, addresses, &resp.Diagnostics)
	} else if !data.Id.IsNull() {
		// Using direct resource ID
		apiRes, err := d.getNextAvailableIPsByID(ctx, data.Id.ValueString(), count, data.Contiguous.ValueBool())

		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Address, got error: %s", err))
			return
		}

		data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)
	} else {
		resp.Diagnostics.AddError("Missing Parameters", "Either id or tag_filters must be specified")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *NextAvailableIPDataSource) getNextAvailableIPsByID(ctx context.Context, id string, count int64, contiguous bool) (*ipam.NextAvailableIPResponse, error) {
	var (
		apiRes *ipam.NextAvailableIPResponse
		err    error
	)

	switch id[:strings.LastIndex(id, "/")] {
	case "ipam/address_block":
		apiRes, _, err = d.client.IPAddressManagementAPI.AddressBlockAPI.
			ListNextAvailableIP(ctx, id).
			Count(int32(count)).
			Contiguous(contiguous).
			Execute()

	case "ipam/subnet":
		apiRes, _, err = d.client.IPAddressManagementAPI.SubnetAPI.
			ListNextAvailableIP(ctx, id).
			Count(int32(count)).
			Contiguous(contiguous).
			Execute()

	case "ipam/range":
		apiRes, _, err = d.client.IPAddressManagementAPI.RangeAPI.
			ListNextAvailableIP(ctx, id).
			Count(int32(count)).
			Contiguous(contiguous).
			Execute()
	}

	return apiRes, err
}

func (d *NextAvailableIPDataSource) findNextAvailableIPsByTags(ctx context.Context, data IpamsvcNextAvailableIPModel) ([]ipam.Address, error) {
	tagFilterStr := flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, nil)

	resourceType := data.ResourceType.ValueString()
	var resources []string
	var err error

	// Get resources matching tags
	switch resourceType {
	case "address_block":
		resources, err = d.findAddressBlocksByTags(ctx, tagFilterStr)
	case "subnet":
		resources, err = d.findSubnetsByTags(ctx, tagFilterStr)
	case "range":
		resources, err = d.findRangesByTags(ctx, tagFilterStr)
	}

	if err != nil {
		return nil, err
	}

	if len(resources) == 0 {
		return nil, fmt.Errorf("no %ss found with the given tags", resourceType)
	}

	// Try to get next available IPs from each resource
	var allAddresses []ipam.Address
	count := data.Count.ValueInt64()
	contiguous := data.Contiguous.ValueBool()

	for _, resourceID := range resources {
		// First check if this resource has at least one available IP
		checkRes, checkErr := d.getNextAvailableIPsByID(ctx, resourceID, 1, contiguous)
		if checkErr != nil || len(checkRes.GetResults()) == 0 {
			continue
		}

		// Try to get as many IPs as needed from this resource
		remainingCount := count - int64(len(allAddresses))

		// Start with requesting the full remaining count
		apiRes, apiErr := d.getNextAvailableIPsByID(ctx, resourceID, remainingCount, contiguous)

		if apiErr != nil {
			// If error occurs, try with progressively smaller counts
			// This mimics the Python while loop that decrements remaining_count
			for tryCount := remainingCount - 1; tryCount > 0; tryCount-- {
				retryRes, retryErr := d.getNextAvailableIPsByID(ctx, resourceID, tryCount, contiguous)
				if retryErr == nil && len(retryRes.GetResults()) > 0 {
					// Found some IPs with the smaller count
					allAddresses = append(allAddresses, retryRes.GetResults()...)
					break
				}
			}
		} else if len(apiRes.GetResults()) > 0 {
			// Successfully got some IPs
			allAddresses = append(allAddresses, apiRes.GetResults()...)
		}

		// Stop if we have enough IPs
		if int64(len(allAddresses)) >= count {
			break
		}
	}

	// Unlike the original code, we return whatever IPs we found, even if less than requested
	return allAddresses, nil
}

func (d *NextAvailableIPDataSource) findAddressBlocksByTags(ctx context.Context, tagFilter string) ([]string, error) {

	var allResources []ipam.AddressBlock

	allResources, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.AddressBlock, error) {
		apiRes, _, err := d.client.IPAddressManagementAPI.AddressBlockAPI.
			List(ctx).
			Tfilter(tagFilter).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			return nil, err
		}

		return apiRes.GetResults(), nil
	})

	if err != nil {
		return nil, err
	}

	resourceIDs := extractResourceIDs(allResources, func(block ipam.AddressBlock) string {
		return block.GetId()
	})

	return resourceIDs, nil
}

func (d *NextAvailableIPDataSource) findSubnetsByTags(ctx context.Context, tagFilter string) ([]string, error) {

	var allResources []ipam.Subnet

	allResources, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.Subnet, error) {
		apiRes, _, err := d.client.IPAddressManagementAPI.SubnetAPI.
			List(ctx).
			Tfilter(tagFilter).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			return nil, err
		}

		return apiRes.GetResults(), nil
	})

	if err != nil {
		return nil, err
	}

	resourceIDs := extractResourceIDs(allResources, func(subnet ipam.Subnet) string {
		return subnet.GetId()
	})
	return resourceIDs, nil
}

func (d *NextAvailableIPDataSource) findRangesByTags(ctx context.Context, tagFilter string) ([]string, error) {

	var allResources []ipam.Range

	allResources, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.Range, error) {
		apiRes, _, err := d.client.IPAddressManagementAPI.RangeAPI.
			List(ctx).
			Tfilter(tagFilter).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			return nil, err
		}

		return apiRes.GetResults(), nil
	})

	if err != nil {
		return nil, err
	}

	resourceIDs := extractResourceIDs(allResources, func(rng ipam.Range) string {
		return rng.GetId()
	})

	return resourceIDs, nil
}

func extractResourceIDs[T any](resources []T, getID func(T) string) []string {
	var resourceIDs []string
	for _, resource := range resources {
		resourceIDs = append(resourceIDs, getID(resource))
	}
	return resourceIDs
}
