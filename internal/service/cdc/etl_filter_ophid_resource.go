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
var _ resource.Resource = &EtlFilterOphidResource{}
var _ resource.ResourceWithImportState = &EtlFilterOphidResource{}

func NewEtlFilterOphidResource() resource.Resource {
	return &EtlFilterOphidResource{}
}

// EtlFilterOphidResource defines the resource implementation.
type EtlFilterOphidResource struct {
	client *universalddiclient.APIClient
}

func (r *EtlFilterOphidResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "cdc_etl_filter_ophid"
}

func (r *EtlFilterOphidResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an OPHID ETL Filter.\n\nThe OPHID ETL Filter object filters Cloud Data Connector (CDC) flow data by on-prem host ID (OPHID).",
		Attributes:          EtlFilterOphidResourceSchemaAttributes,
	}
}

func (r *EtlFilterOphidResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *EtlFilterOphidResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data EtlFilterOphidModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		CreateEtlFilterOphid(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create EtlFilterOphid, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterOphidResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data EtlFilterOphidModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		ReadEtlFilterOphid(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read EtlFilterOphid, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterOphidResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data EtlFilterOphidModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		UpdateEtlFilterOphid(ctx, data.Id.ValueInt64()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update EtlFilterOphid, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EtlFilterOphidResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data EtlFilterOphidModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.CDCAPI.
		EtlFilterApiAPI.
		DeleteEtlFilterOphid(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete EtlFilterOphid, got error: %s", err))
		return
	}
}

func (r *EtlFilterOphidResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	utils.ImportStatePassthroughInt64ID(ctx, req, resp)
}
