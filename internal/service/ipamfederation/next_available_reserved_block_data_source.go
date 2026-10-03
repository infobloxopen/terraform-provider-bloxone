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
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableReservedBlockDataSource{}

func NewNextAvailableReservedBlockDataSource() datasource.DataSource {
	return &NextAvailableReservedBlockDataSource{}
}

// NextAvailableReservedBlockDataSource defines the data source implementation.
type NextAvailableReservedBlockDataSource struct {
	client *universalddiclient.APIClient
}

type NextAvailableReservedBlockModel struct {
	Id                 types.String `tfsdk:"id"`
	Cidr               types.Int64  `tfsdk:"cidr"`
	ReservedBlockCount types.Int64  `tfsdk:"reserved_block_count"`
	Name               types.String `tfsdk:"name"`
	Comment            types.String `tfsdk:"comment"`
	Results            types.List   `tfsdk:"results"`
}

func (m *NextAvailableReservedBlockModel) FlattenResults(ctx context.Context, from []ipamfederation.ReservedBlock, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, ReservedBlockAttrTypes, diags, FlattenReservedBlock)
}

func (d *NextAvailableReservedBlockDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "next_available_reserved_blocks"
}

func (d *NextAvailableReservedBlockDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves the next available __ReservedBlock__ objects from the specified parent __FederatedBlock__, without allocating them.",
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
				MarkdownDescription: "The CIDR of the reserved blocks to be created.",
			},
			"reserved_block_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The count of reserved blocks required. If not provided, it will default to 1.",
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
				MarkdownDescription: "The description for the reserved block. May contain 0 to 1024 characters. Can include UTF-8.",
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(ReservedBlockResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed:            true,
				MarkdownDescription: "List of next available reserved blocks in the specified parent Federated Block.",
			},
		},
	}
}

func (d *NextAvailableReservedBlockDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableReservedBlockDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableReservedBlockModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ReservedBlockCount.IsNull() {
		data.ReservedBlockCount = types.Int64Value(1)
	}

	apiReq := d.client.IPAMFederationAPI.
		NextAvailableReservedBlockAPI.
		ListNextAvailableReservedBlocks(ctx, data.Id.ValueString()).
		Cidr(data.Cidr.ValueInt64()).
		Count(data.ReservedBlockCount.ValueInt64())

	if !data.Name.IsNull() {
		apiReq = apiReq.Name(data.Name.ValueString())
	}
	if !data.Comment.IsNull() {
		apiReq = apiReq.Comment(data.Comment.ValueString())
	}

	apiRes, _, err := apiReq.Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NextAvailableReservedBlock, got error: %s", err))
		return
	}

	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
