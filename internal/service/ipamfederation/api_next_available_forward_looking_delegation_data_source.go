package ipamfederation

import (
	"context"
	"fmt"

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
)

var _ datasource.DataSource = &NextAvailableForwardLookingDelegationDataSource{}

func NewNextAvailableForwardLookingDelegationDataSource() datasource.DataSource {
	return &NextAvailableForwardLookingDelegationDataSource{}
}

type NextAvailableForwardLookingDelegationDataSource struct {
	client *universalddiclient.APIClient
}

type NextAvailableForwardLookingDelegationModel struct {
	FederatedBlockId types.String `tfsdk:"federated_block_id"`
	FederatedPoolId  types.String `tfsdk:"federated_pool_id"`
	Cidr             types.Int64  `tfsdk:"cidr"`
	FldCount         types.Int64  `tfsdk:"fld_count"`
	Comment          types.String `tfsdk:"comment"`
	Name             types.String `tfsdk:"name"`
	Protocol         types.String `tfsdk:"protocol"`
	Tags             types.Map    `tfsdk:"tags"`
	Results          types.List   `tfsdk:"results"`
}

func (m *NextAvailableForwardLookingDelegationModel) FlattenResults(ctx context.Context, from []ipamfederation.ForwardLookingDelegation, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	var addresses []string
	for _, fld := range from {
		if fld.Address != nil {
			addresses = append(addresses, *fld.Address)
		}
	}
	m.Results = flex.FlattenFrameworkListString(ctx, addresses, diags)
}

func (d *NextAvailableForwardLookingDelegationDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_next_available_forward_looking_delegations"
}

func (d *NextAvailableForwardLookingDelegationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Creates and retrieves the next available Forward Looking Delegation objects. Supports creation under a specific Federated Block, a specific Federated Pool, or globally across all available blocks.",
		Attributes: map[string]schema.Attribute{
			"federated_block_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Block. Mutually exclusive with `federated_pool_id`. When set, FLDs are created within this specific block.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("federated_pool_id")),
				},
			},
			"federated_pool_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The resource identifier of the parent Federated Pool. Mutually exclusive with `federated_block_id`. When set, FLDs are created within this specific pool.",
				Validators: []validator.String{
					stringvalidator.ConflictsWith(path.MatchRoot("federated_block_id")),
				},
			},
			"cidr": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The CIDR of the Forward Looking Delegations to be created.",
			},
			"fld_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Number of Forward Looking Delegation objects to generate. Default 1 if not set.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"comment": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description for the Forward Looking Delegations. May contain 0 to 1024 characters. Can include UTF-8.",
			},
			"name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The name of the Forward Looking Delegations. May contain 1 to 256 characters. Can include UTF-8.",
			},
			"protocol": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The type of protocol of delegation (_ip4_ or _ip6_). Not applicable when using `federated_block_id`.",
				Validators: []validator.String{
					stringvalidator.OneOf("ip4", "ip6"),
				},
			},
			"tags": schema.MapAttribute{
				ElementType:         types.StringType,
				Optional:            true,
				MarkdownDescription: "The tags for the Forward Looking Delegations in JSON format.",
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of addresses of the created Forward Looking Delegation objects.",
			},
		},
	}
}

func (d *NextAvailableForwardLookingDelegationDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableForwardLookingDelegationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NextAvailableForwardLookingDelegationModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var results []ipamfederation.ForwardLookingDelegation

	switch {
	case !data.FederatedBlockId.IsNull() && !data.FederatedBlockId.IsUnknown():
		results = d.createUnderBlock(ctx, data, resp)
	case !data.FederatedPoolId.IsNull() && !data.FederatedPoolId.IsUnknown():
		results = d.createUnderPool(ctx, data, resp)
	default:
		results = d.createGlobal(ctx, data, resp)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	data.FlattenResults(ctx, results, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (d *NextAvailableForwardLookingDelegationDataSource) createGlobal(ctx context.Context, data NextAvailableForwardLookingDelegationModel, resp *datasource.ReadResponse) []ipamfederation.ForwardLookingDelegation {
	body := ipamfederation.NewNextAvailableFLDRequest(data.Cidr.ValueInt64())
	if !data.FldCount.IsNull() && !data.FldCount.IsUnknown() {
		body.SetCount(data.FldCount.ValueInt64())
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		body.SetComment(data.Comment.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body.SetName(data.Name.ValueString())
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		body.SetProtocol(data.Protocol.ValueString())
	}
	if !data.Tags.IsNull() && !data.Tags.IsUnknown() {
		body.SetTags(flex.ExpandFrameworkMapString(ctx, data.Tags, &resp.Diagnostics))
		if resp.Diagnostics.HasError() {
			return nil
		}
	}

	apiRes, _, err := d.client.IPAMFederationAPI.NextAvailableFldAPI.
		CreateNextAvailableFLDBlocks(ctx).
		Body(*body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available Forward Looking Delegations, got error: %s", err))
		return nil
	}
	return apiRes.GetResults()
}

func (d *NextAvailableForwardLookingDelegationDataSource) createUnderBlock(ctx context.Context, data NextAvailableForwardLookingDelegationModel, resp *datasource.ReadResponse) []ipamfederation.ForwardLookingDelegation {
	blockId := data.FederatedBlockId.ValueString()
	body := ipamfederation.NewCreateNextAvailableFLDRequestForBlock(data.Cidr.ValueInt64(), blockId)
	if !data.FldCount.IsNull() && !data.FldCount.IsUnknown() {
		body.SetCount(data.FldCount.ValueInt64())
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		body.SetComment(data.Comment.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body.SetName(data.Name.ValueString())
	}
	if !data.Tags.IsNull() && !data.Tags.IsUnknown() {
		body.SetTags(flex.ExpandFrameworkMapString(ctx, data.Tags, &resp.Diagnostics))
		if resp.Diagnostics.HasError() {
			return nil
		}
	}

	apiRes, _, err := d.client.IPAMFederationAPI.NextAvailableFldAPI.
		CreateNextAvailableFLD(ctx, blockId).
		Body(*body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available Forward Looking Delegations under Federated Block %s, got error: %s", blockId, err))
		return nil
	}
	return apiRes.GetResults()
}

func (d *NextAvailableForwardLookingDelegationDataSource) createUnderPool(ctx context.Context, data NextAvailableForwardLookingDelegationModel, resp *datasource.ReadResponse) []ipamfederation.ForwardLookingDelegation {
	poolId := data.FederatedPoolId.ValueString()
	body := ipamfederation.NewNextAvailableFLDPoolRequest(data.Cidr.ValueInt64(), poolId)
	if !data.FldCount.IsNull() && !data.FldCount.IsUnknown() {
		body.SetCount(data.FldCount.ValueInt64())
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		body.SetComment(data.Comment.ValueString())
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		body.SetName(data.Name.ValueString())
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		body.SetProtocol(data.Protocol.ValueString())
	}
	if !data.Tags.IsNull() && !data.Tags.IsUnknown() {
		body.SetTags(flex.ExpandFrameworkMapString(ctx, data.Tags, &resp.Diagnostics))
		if resp.Diagnostics.HasError() {
			return nil
		}
	}

	apiRes, _, err := d.client.IPAMFederationAPI.NextAvailableFldAPI.
		CreateNextAvailableFLDForPool(ctx, poolId).
		Body(*body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available Forward Looking Delegations under Federated Pool %s, got error: %s", poolId, err))
		return nil
	}
	return apiRes.GetResults()
}
