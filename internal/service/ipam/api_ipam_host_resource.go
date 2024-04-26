package ipam

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IpamHostResource{}
var _ resource.ResourceWithImportState = &IpamHostResource{}

func NewIpamHostResource() resource.Resource {
	return &IpamHostResource{}
}

// IpamHostResource defines the resource implementation.
type IpamHostResource struct {
	client *bloxoneclient.APIClient
}

func (r *IpamHostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_host"
}

func (r *IpamHostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an IPAM Host.\n\nThe IPAM host object represents any network connected equipment that is assigned one or more IP Addresses.",
		Attributes:          IpamsvcIpamHostResourceSchemaAttributes,
	}
}

func (r *IpamHostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IpamHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpamsvcIpamHostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if !data.Addresses.IsNull() && !data.Addresses.IsUnknown() {
		nextAvaialableIds := r.getNextAvailableIds(ctx, data)
		for _, id := range nextAvaialableIds {
			// Lock the mutex to serialize operations with the same key
			// This is necessary to prevent the same block being returned.
			utils.GlobalMutexStore.Lock(id)
			defer utils.GlobalMutexStore.Unlock(id)
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	apiRes, _, err := r.client.IPAddressManagementAPI.
		IpamHostAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create IpamHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpamHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpamsvcIpamHostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAddressManagementAPI.
		IpamHostAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IpamHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpamHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IpamsvcIpamHostModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAddressManagementAPI.
		IpamHostAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update IpamHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpamHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpamsvcIpamHostModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAddressManagementAPI.
		IpamHostAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete IpamHost, got error: %s", err))
		return
	}
}

func (r *IpamHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpamHostResource) getNextAvailableIds(ctx context.Context, data IpamsvcIpamHostModel) []string {
	// Get the list of addresses to lock the mutex
	nextAvailableIds := make(map[string]struct{})
	var hostAdresses []IpamsvcHostAddressModel
	data.Addresses.ElementsAs(ctx, &hostAdresses, false)
	for _, a := range hostAdresses {
		if !a.NextAvailableId.IsUnknown() && !a.NextAvailableId.IsNull() {
			nextAvailableIds[a.NextAvailableId.ValueString()] = struct{}{}
		}
	}

	// sort the list of ids to lock the mutex and return a list of unique ids
	// the sort is necessary to prevent deadlocks
	ids := make([]string, 0, len(nextAvailableIds))
	for id := range nextAvailableIds {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
