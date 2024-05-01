package dfp

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"net/http"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DfpResource{}
var _ resource.ResourceWithImportState = &DfpResource{}

func NewDfpResource() resource.Resource {
	return &DfpResource{}
}

// DfpResource defines the resource implementation.
type DfpResource struct {
	client *bloxoneclient.APIClient
}

func (r *DfpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dfp_service"
}

func (r *DfpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNS Forwarding Proxy Service.",
		Attributes:          DfpResourceSchemaAttributes,
	}
}

func (r *DfpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DfpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DfpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.DNSForwardingProxyAPI.
		InfraServicesAPI.
		CreateOrUpdateDfpService(ctx, data.ServiceId.ValueString()).
		Body(*data.ExpandCreateOrUpdatePayload(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dfp, got error: %s", err))
		return
	}

	// We call Read again, so that the fields not part of the CreateOrUpdatePayload are also populated
	apiRes, _, err := r.client.DNSForwardingProxyAPI.
		InfraServicesAPI.
		ReadDfpService(ctx, data.ServiceId.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dfp, got error: %s", err))
		return
	}
	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DfpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DfpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSForwardingProxyAPI.
		InfraServicesAPI.
		ReadDfpService(ctx, data.ServiceId.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Dfp, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DfpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DfpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, _, err := r.client.DNSForwardingProxyAPI.
		InfraServicesAPI.
		CreateOrUpdateDfpService(ctx, data.ServiceId.ValueString()).
		Body(*data.ExpandCreateOrUpdatePayload(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Dfp, got error: %s", err))
		return
	}

	// We call Read again, so that the fields not part of the CreateOrUpdatePayload are also populated
	apiRes, _, err := r.client.DNSForwardingProxyAPI.
		InfraServicesAPI.
		ReadDfpService(ctx, data.ServiceId.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Dfp, got error: %s", err))
		return
	}
	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DfpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DfpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: check the UI calls, or confirm if DELETE is called explicitly for the resource in the UI.

	resp.State.RemoveResource(ctx)
}

func (r *DfpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
