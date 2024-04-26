package ipam

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	Id      types.String `tfsdk:"id"`
	Cidr    types.Int64  `tfsdk:"cidr"`
	Count   types.Int64  `tfsdk:"subnet_count"`
	Results types.List   `tfsdk:"results"`
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
				Required:            true,
				MarkdownDescription: `An application specific resource identity of a resource.`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/address_block/[0-9a-f-].*$`), "invalid resource ID specified"),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: `The cidr value of subnets to be created.`,
			},
			"subnet_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: `Number of subnets to generate. Default 1 if not set.`,
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of Next available Subnet address in the specified resource",
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

func (d *NextAvailableSubnetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpamsvcNextAvailableSubnetModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := d.client.IPAddressManagementAPI.
		AddressBlockAPI.
		ListNextAvailableSubnet(ctx, data.Id.ValueString()).
		Cidr(int32(data.Cidr.ValueInt64())).
		Count(int32(data.Count.ValueInt64())).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AddressBlock Next Available Subnet API, got error: %s", err))
		return
	}

	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
