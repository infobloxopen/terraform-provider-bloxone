package ipam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AddressBlockResource{}
var _ resource.ResourceWithImportState = &AddressBlockResource{}
var _ resource.ResourceWithModifyPlan = &AddressBlockResource{}

var inheritanceType = "full"

func NewAddressBlockResource() resource.Resource {
	return &AddressBlockResource{}
}

// AddressBlockResource defines the resource implementation.
type AddressBlockResource struct {
	client *universalddiclient.APIClient
}

func (r *AddressBlockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_address_block"
}

func (r *AddressBlockResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Address Block.\n\nThe Address Block object allows a uniform representation of the address space segmentation, " +
			"supporting functions such as administrative grouping, routing aggregation, delegation etc.\n\n" +
			"When `next_available_id` is set, `terraform plan` checks that an address block of the requested `cidr` is currently available in that " +
			"address block. The plan fails if the address block does not exist, the `cidr` does not fit, or no space is left. The candidate " +
			"shown in the plan is a preview only and is not reserved; the address block is allocated at apply and may differ.",
		Attributes: IpamsvcAddressBlockResourceSchemaAttributes,
	}
}

func (r *AddressBlockResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// ModifyPlan checks at plan time that a next-available allocation can be satisfied when the object is created.
func (r *AddressBlockResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if r.client == nil || !req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var data IpamsvcAddressBlockModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() ||
		data.NextAvailableId.IsNull() || data.NextAvailableId.IsUnknown() || data.Cidr.IsNull() || data.Cidr.IsUnknown() {
		return
	}
	parentId, cidr := data.NextAvailableId.ValueString(), data.Cidr.ValueInt64()
	checkNextAvailable("address block", parentId, cidr, func() (string, *http.Response, error) {
		apiRes, httpRes, err := r.client.IPAddressManagementAPI.AddressBlockAPI.
			ListNextAvailableAB(ctx, parentId).
			Cidr(int32(cidr)).
			Count(1).
			Execute()
		if err != nil || len(apiRes.GetResults()) == 0 {
			return "", httpRes, err
		}
		return apiRes.GetResults()[0].GetAddress(), httpRes, nil
	}, &resp.Diagnostics)
}

func (r *AddressBlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpamsvcAddressBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.NextAvailableId.IsUnknown() && !data.NextAvailableId.IsNull() {
		// Lock the mutex to serialize operations with the same key
		// This is necessary to prevent the same block being returned.
		utils.GlobalMutexStore.Lock(data.NextAvailableId.ValueString())
		defer utils.GlobalMutexStore.Unlock(data.NextAvailableId.ValueString())
	}

	apiRes, _, err := r.client.IPAddressManagementAPI.
		AddressBlockAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics, true)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		if !data.NextAvailableId.IsUnknown() && !data.NextAvailableId.IsNull() {
			resp.Diagnostics.AddError("Next-available allocation failed",
				nextAvailableCreateError("address block", data.NextAvailableId.ValueString(), data.Cidr.ValueInt64(), err))
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AddressBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddressBlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpamsvcAddressBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAddressManagementAPI.
		AddressBlockAPI.
		Read(ctx, data.Id.ValueString()).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AddressBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddressBlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IpamsvcAddressBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAddressManagementAPI.
		AddressBlockAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AddressBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AddressBlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpamsvcAddressBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAddressManagementAPI.
		AddressBlockAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AddressBlock, got error: %s", err))
		return
	}
}

func (r *AddressBlockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
