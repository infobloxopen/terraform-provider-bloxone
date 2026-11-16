package ipam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SubnetResource{}
var _ resource.ResourceWithImportState = &SubnetResource{}
var _ resource.ResourceWithModifyPlan = &SubnetResource{}

func NewSubnetResource() resource.Resource {
	return &SubnetResource{}
}

// SubnetResource defines the resource implementation.
type SubnetResource struct {
	client *universalddiclient.APIClient
}

func (r *SubnetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_subnet"
}

func (r *SubnetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Subnet.\n\nThe Subnet object represents a set of addresses from which addresses are assigned to network equipment interfaces.\n\nWhen `next_available_id` is set and `preflight_check` is `true`, `terraform plan` checks that a subnet of the requested `cidr` is currently available in that address block. The plan fails if the address block does not exist, the `cidr` does not fit, or no space is left. The candidate shown in the plan is a preview only and is not reserved; the subnet is allocated at apply and may differ.",
		Attributes:          IpamsvcSubnetResourceSchemaAttributes,
	}
}

func (r *SubnetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// ModifyPlan checks at plan time that a next-available allocation can be satisfied when the object is created
// with preflight_check enabled.
func (r *SubnetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if r.client == nil || !req.State.Raw.IsNull() || req.Plan.Raw.IsNull() {
		return
	}
	var data IpamsvcSubnetModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() || !data.PreflightCheck.ValueBool() ||
		data.NextAvailableId.IsNull() || data.NextAvailableId.IsUnknown() || data.Cidr.IsNull() || data.Cidr.IsUnknown() {
		return
	}
	parentId, cidr := data.NextAvailableId.ValueString(), data.Cidr.ValueInt64()
	checkNextAvailable(r.client, "subnet", parentId, cidr, func(count int32) ([]string, *http.Response, error) {
		apiRes, httpRes, err := r.client.IPAddressManagementAPI.AddressBlockAPI.
			ListNextAvailableSubnet(ctx, parentId).
			Cidr(int32(cidr)).
			Count(count).
			Execute()
		if err != nil {
			return nil, httpRes, err
		}
		var addresses []string
		for _, res := range apiRes.GetResults() {
			addresses = append(addresses, res.GetAddress())
		}
		return addresses, httpRes, nil
	}, func() (string, error) {
		apiRes, _, err := r.client.IPAddressManagementAPI.AddressBlockAPI.Read(ctx, parentId).Execute()
		if err != nil {
			return "", err
		}
		parent := apiRes.GetResult()
		return fmt.Sprintf("%s/%d", parent.GetAddress(), parent.GetCidr()), nil
	}, &resp.Diagnostics)
}

func (r *SubnetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpamsvcSubnetModel

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
		SubnetAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics, true)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		if !data.NextAvailableId.IsUnknown() && !data.NextAvailableId.IsNull() {
			resp.Diagnostics.AddError("Next-available allocation failed",
				nextAvailableCreateError("subnet", data.NextAvailableId.ValueString(), data.Cidr.ValueInt64(), err))
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Subnet, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubnetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpamsvcSubnetModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAddressManagementAPI.
		SubnetAPI.
		Read(ctx, data.Id.ValueString()).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Subnet, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)
	// preflight_check is not returned by the API; default it for imported or pre-upgrade state.
	if data.PreflightCheck.IsNull() {
		data.PreflightCheck = types.BoolValue(false)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubnetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IpamsvcSubnetModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAddressManagementAPI.
		SubnetAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Subnet, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubnetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpamsvcSubnetModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAddressManagementAPI.
		SubnetAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Subnet, got error: %s", err))
		return
	}
}

func (r *SubnetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
