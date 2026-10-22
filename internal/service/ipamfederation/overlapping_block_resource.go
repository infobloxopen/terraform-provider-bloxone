package ipamfederation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OverlappingBlockResource{}
var _ resource.ResourceWithImportState = &OverlappingBlockResource{}

func NewOverlappingBlockResource() resource.Resource {
	return &OverlappingBlockResource{}
}

// OverlappingBlockResource defines the resource implementation.
type OverlappingBlockResource struct {
	client *universalddiclient.APIClient
}

func (r *OverlappingBlockResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_overlapping_block"
}

func (r *OverlappingBlockResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Overlapping Block.\n\nThe Overlapping Block object is a set of contiguous IP addresses with no gap, expressed as a CIDR block. It is explicitly associated with a Federated Realm, and implicitly with a Federated Block Parent. An Overlapping Block in a given realm is said to be the child of the closest enclosing parent, and indicates an address range that may be managed independently by all participating IPAM services.",
		Attributes:          OverlappingBlockResourceSchemaAttributes,
	}
}

func (r *OverlappingBlockResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OverlappingBlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data OverlappingBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		OverlappingBlockAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics, true)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OverlappingBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OverlappingBlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OverlappingBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMFederationAPI.
		OverlappingBlockAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read OverlappingBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OverlappingBlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OverlappingBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		OverlappingBlockAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update OverlappingBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *OverlappingBlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OverlappingBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMFederationAPI.
		OverlappingBlockAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete OverlappingBlock, got error: %s", err))
		return
	}
}

func (r *OverlappingBlockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
