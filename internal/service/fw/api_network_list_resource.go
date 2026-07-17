package fw

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

const (
	// NetworkListOperationTimeout is the maximum amount of time to wait for eventual consistency
	NetworkListOperationTimeout = 2 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NetworkListResource{}
var _ resource.ResourceWithImportState = &NetworkListResource{}

func NewNetworkListResource() resource.Resource {
	return &NetworkListResource{}
}

// NetworkListResource defines the resource implementation.
type NetworkListResource struct {
	client *universalddiclient.APIClient
}

func (r *NetworkListResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "td_network_list"
}

func (r *NetworkListResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Network List.",
		Attributes:          AtcfwNetworkListResourceSchemaAttributes,
	}
}

func (r *NetworkListResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *NetworkListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AtcfwNetworkListModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, NetworkListOperationTimeout, func() *retry.RetryError {
		apiRes, _, err := r.client.FWAPI.
			NetworkListsAPI.
			CreateNetworkList(ctx).
			Body(*data.Expand(ctx, &resp.Diagnostics)).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "Internal Server Error") {
				tflog.Debug(ctx, "Waiting for related objects to be present, will retry", map[string]interface{}{"error": err.Error()})
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create NetworkList, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		res := apiRes.GetResults()
		data.Flatten(ctx, &res, &resp.Diagnostics)

		return nil
	})
	if err != nil {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AtcfwNetworkListModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.FWAPI.
		NetworkListsAPI.
		ReadNetworkList(ctx, data.Id.ValueInt32()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NetworkLists, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AtcfwNetworkListModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.FWAPI.
		NetworkListsAPI.
		UpdateNetworkList(ctx, data.Id.ValueInt32()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update NetworkLists, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetworkListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AtcfwNetworkListModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, NetworkListOperationTimeout, func() *retry.RetryError {
		httpRes, err := r.client.FWAPI.
			NetworkListsAPI.
			DeleteSingleNetworkLists(ctx, data.Id.ValueInt32()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "Internal Server Error") || strings.Contains(err.Error(), "Operation timed out") {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete NetworkLists, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (r *NetworkListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
