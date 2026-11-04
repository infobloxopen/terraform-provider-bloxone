package cdc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EtlFilterIpNetworkResource{}
var _ resource.ResourceWithImportState = &EtlFilterIpNetworkResource{}

func NewEtlFilterIpNetworkResource() resource.Resource {
	return &EtlFilterIpNetworkResource{}
}

// EtlFilterIpNetworkResource defines the resource implementation.
type EtlFilterIpNetworkResource struct {
	client *universalddiclient.APIClient
}

func (r *EtlFilterIpNetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "cdc_etl_filter_ip_network"
}

func (r *EtlFilterIpNetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IP Network ETL Filter.\n\nThe IP Network ETL Filter object filters Cloud Data Connector (CDC) flow data by IP address or network.",
		Attributes:          EtlFilterIpNetworkResourceSchemaAttributes,
	}
}

func (r *EtlFilterIpNetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*universalddiclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *universalddiclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EtlFilterIpNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EtlFilterIpNetworkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		CreateEtlFilterIpNetwork(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create EtlFilterIpNetwork, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterIpNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EtlFilterIpNetworkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		ReadEtlFilterIpNetwork(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read EtlFilterIpNetwork, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterIpNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data EtlFilterIpNetworkModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		UpdateEtlFilterIpNetwork(ctx, data.Id.ValueInt64()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update EtlFilterIpNetwork, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterIpNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data EtlFilterIpNetworkModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		DeleteEtlFilterIpNetwork(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete EtlFilterIpNetwork, got error: %s", err))
		return
	}
}

func (r *EtlFilterIpNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	utils.ImportStatePassthroughInt64ID(ctx, req, resp)
}
