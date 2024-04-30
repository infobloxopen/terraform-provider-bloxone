package anycast

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"net/http"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OnPremAnycastHostResource{}
var _ resource.ResourceWithImportState = &OnPremAnycastHostResource{}

func NewOnPremAnycastOnpremHostResource() resource.Resource {
	return &OnPremAnycastHostResource{}
}

// OnPremAnycastManagerResource defines the resource implementation.
type OnPremAnycastHostResource struct {
	client *bloxoneclient.APIClient
}

func (r *OnPremAnycastHostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "anycast_host"
}

func (r *OnPremAnycastHostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieve an Anycast host Configurations.",
		Attributes:          ProtoOnpremHostResourceSchemaAttributes,
	}
}

func (r *OnPremAnycastHostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*bloxoneclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *bloxoneclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

//Read from ID. (ID should be mandatory)
//Retry until create_timeout, return error if not found
//If the data is different from read data from API, make an PUT call to API, to update the resource in backend.

// create is uncommented cause to create in api documentation
func (r *OnPremAnycastHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ProtoOnpremHostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// We first query the Host API to get the name and IP address of the host.
	// This is required, or any name set by the user is simply overwritten by the Host sync process.
	hostRes, _, err := r.client.InfraManagementAPI.
		HostsAPI.
		List(ctx).
		Filter(fmt.Sprintf("legacy_id == '%d'", data.Id.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
		return
	}
	if len(hostRes.GetResults()) != 1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", "Host not found"))
		return
	}
	data.Name = types.StringValue(hostRes.GetResults()[0].DisplayName)
	data.IpAddress = types.StringPointerValue(hostRes.GetResults()[0].IpAddress)

	//now we call put call
	apiRes, _, err := r.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		UpdateOnpremHost(ctx, data.Id.ValueInt64()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
		return
	}
	//compare Api res with data and update if needed, do a put call if there is a change (what is deep comparision)
	//or call the put command and override the data

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OnPremAnycastHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	//var data ProtoAnycastConfigModel
	var data ProtoOnpremHostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	//OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus
	apiRes, httpRes, err := r.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		GetOnpremHost(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read OnPremAnycastManager, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OnPremAnycastHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//var data ProtoAnycastConfigModel
	var data ProtoOnpremHostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	//OnPremAnycastManagerUpdateAnycastConfig
	apiRes, _, err := r.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		UpdateOnpremHost(ctx, data.Id.ValueInt64()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update OnPremAnycastManager, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OnPremAnycastHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProtoOnpremHostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	//OnPremAnycastManagerDeleteAnycastConfig
	_, httpRes, err := r.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		DeleteOnpremHost(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete OnPremAnycastManager, got error: %s", err))
		return
	}
}

func (r *OnPremAnycastHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
