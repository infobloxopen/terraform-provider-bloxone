package anycast

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
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
		Attributes:          ProtoAnycastConfigResourceSchemaAttributes,
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

	err := retry.RetryContext(ctx, 20*time.Minute, func() *retry.RetryError {
		//get  id from infra
		hostRes, _, err := r.client.InfraManagementAPI.
			HostsAPI.
			HostsList(ctx).
			Filter(fmt.Sprintf("legacy_id == '%d'", data.Id.ValueInt64())).Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
			return retry.NonRetryableError(err)
		}

		//now you add the readonly data
		data.Name = types.StringValue(hostRes.GetResults()[0].DisplayName)
		data.IpAddress = types.StringPointerValue(hostRes.GetResults()[0].IpAddress)
		return nil
	})
	//now we call put call
	_, _, err = r.client.AnycastAPI.OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateOnpremHost(ctx, data.Id.ValueInt64()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).Execute()

	err = retry.RetryContext(ctx, 20*time.Minute, func() *retry.RetryError {
		_, res, err := r.client.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerGetOnpremHost(ctx, data.Id.ValueInt64()).
			Execute()
		if err != nil {
			if res.StatusCode == http.StatusNotFound {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
		return
	}
	//compare Api res with data and update if needed, do a put call if there is a change (what is deep comparision)
	//or call the put command and override the data

	/*
		res := apiRes.GetResults()
		data.Flatten(ctx, &res, &resp.Diagnostics)*/
	//r.client.AnycastAPI.OnPremAnycastManagerAPI.OnPremAnycastManagerUpdateOnpremHost(ctx, data.Id.ValueInt64()).Body(data.OnpremHost).Execute(}
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
		OnPremAnycastManagerGetOnpremHost(ctx, data.Id.ValueInt64()).
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
		OnPremAnycastManagerUpdateOnpremHost(ctx, data.Id.ValueInt64()).
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
	var data ProtoAnycastConfigModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	//OnPremAnycastManagerDeleteAnycastConfig
	_, httpRes, err := r.client.AnycastAPI.
		OnPremAnycastManagerAPI.
		OnPremAnycastManagerDeleteOnpremHost(ctx, data.Id.ValueInt64()).
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
