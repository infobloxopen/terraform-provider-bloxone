package ipamfederation

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// federatedPoolIdPrefix identifies a next_available_id that names a Federated Pool, as opposed to
// a Federated Block. Next-available allocation from a Federated Pool has no preview (GET) support
// and is routed through a different API service than allocation from a Federated Block.
const federatedPoolIdPrefix = "federation/federated_pool/"

// resolvePoolId returns the Federated Pool id that governs this block, either from federated_pool_id
// directly, or from next_available_id when it names a Federated Pool rather than a Federated Block.
func (m *FederatedBlockModel) resolvePoolId() (string, bool) {
	if !m.FederatedPoolId.IsNull() && !m.FederatedPoolId.IsUnknown() {
		return m.FederatedPoolId.ValueString(), true
	}
	if !m.NextAvailableId.IsNull() && !m.NextAvailableId.IsUnknown() && strings.HasPrefix(m.NextAvailableId.ValueString(), federatedPoolIdPrefix) {
		return m.NextAvailableId.ValueString(), true
	}
	return "", false
}

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FederatedBlockResource{}
var _ resource.ResourceWithModifyPlan = &FederatedBlockResource{}
var _ resource.ResourceWithImportState = &FederatedBlockResource{}

func NewFederatedBlockResource() resource.Resource {
	return &FederatedBlockResource{}
}

// FederatedBlockResource defines the resource implementation.
type FederatedBlockResource struct {
	client *universalddiclient.APIClient
}

func (r *FederatedBlockResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_federated_block"
}

func (r *FederatedBlockResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Federated Block.\n\nThe Federated Block object allows a uniform representation of the address space segmentation, supporting functions such as administrative grouping, routing aggregation, delegation etc.",
		Attributes:          FederatedBlockResourceSchemaAttributes,
	}
}

func (r *FederatedBlockResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FederatedBlockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FederatedBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// network_compliance is not supported for blocks in an AWS pool. A block takes its provider from its pool,
	// so look the pool up and reject the attribute before the API does.
	if !data.NetworkCompliance.IsNull() && !data.NetworkCompliance.IsUnknown() {
		if poolId, ok := data.resolvePoolId(); ok {
			poolRes, _, err := r.client.IPAMFederationAPI.
				FederatedPoolAPI.
				Read(ctx, poolId).
				Execute()
			if err == nil {
				pool := poolRes.GetResult()
				if pool.GetProvider() == ipamfederation.PROVIDERTYPE_AWS {
					resp.Diagnostics.AddAttributeError(path.Root("network_compliance"), "Unsupported attribute for provider type",
						"network_compliance is not supported for blocks in a pool with provider_type \"AWS\".")
					return
				}
			}
		}
	}

	if !data.NextAvailableId.IsUnknown() && !data.NextAvailableId.IsNull() {
		nextAvailableId := data.NextAvailableId.ValueString()

		// Lock the mutex to serialize operations with the same key.
		// This is necessary to prevent the same block being returned.
		utils.GlobalMutexStore.Lock(nextAvailableId)
		defer utils.GlobalMutexStore.Unlock(nextAvailableId)

		body := ipamfederation.NextAvailableBlockRequest{
			Cidr:  flex.ExpandInt64Pointer(data.Cidr),
			Count: utils.Ptr(int64(1)),
			Tags:  flex.ExpandFrameworkMapString(ctx, data.Tags, &resp.Diagnostics),
		}
		if !data.Name.IsNull() {
			body.Name = data.Name.ValueStringPointer()
		}
		if !data.Comment.IsNull() {
			body.Comment = data.Comment.ValueStringPointer()
		}

		if strings.HasPrefix(nextAvailableId, federatedPoolIdPrefix) {
			// Allocating directly from a Federated Pool: the pool is already the parent, so there
			// is no separate federated_pool_id to reconcile afterwards.
			apiRes, _, err := r.client.IPAMFederationAPI.
				NextAvailablePoolBlockAPI.
				CreateNextAvailableBlocks(ctx, nextAvailableId).
				Body(body).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available FederatedBlock, got error: %s", err))
				return
			}

			results := apiRes.GetResults()
			if len(results) == 0 {
				resp.Diagnostics.AddError("Client Error", "Unable to create next available FederatedBlock, no results were returned")
				return
			}

			data.Flatten(ctx, &results[0], &resp.Diagnostics)

			// Save data into Terraform state
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}

		// Allocating under a parent Federated Block. federated_pool_id does not persist on the
		// create-next-available endpoint, so it is patched in with a follow-up Update call if set.
		configuredFederatedPoolId := data.FederatedPoolId

		apiRes, _, err := r.client.IPAMFederationAPI.
			NextAvailableFederatedBlockAPI.
			CreateNextAvailableFederatedBlocks(ctx, nextAvailableId).
			Body(body).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create next available FederatedBlock, got error: %s", err))
			return
		}

		results := apiRes.GetResults()
		if len(results) == 0 {
			resp.Diagnostics.AddError("Client Error", "Unable to create next available FederatedBlock, no results were returned")
			return
		}

		data.Flatten(ctx, &results[0], &resp.Diagnostics)

		if !configuredFederatedPoolId.IsNull() {
			// Remember what the backend actually persisted from the create call, so that if the
			// follow-up Update below fails, state reflects reality rather than the desired config.
			createdFederatedPoolId := data.FederatedPoolId

			data.FederatedPoolId = configuredFederatedPoolId
			updateRes, _, err := r.client.IPAMFederationAPI.
				FederatedBlockAPI.
				Update(ctx, data.Id.ValueString()).
				Body(*data.Expand(ctx, &resp.Diagnostics, false)).
				Execute()
			if err != nil {
				// The FederatedBlock was already created successfully above. Save it to state now,
				// using the backend's actual federated_pool_id, so it isn't orphaned outside of
				// Terraform; a subsequent apply will retry setting federated_pool_id via Update.
				data.FederatedPoolId = createdFederatedPoolId
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set federated_pool_id on next available FederatedBlock, got error: %s", err))
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
			res := updateRes.GetResult()
			data.Flatten(ctx, &res, &resp.Diagnostics)
		}

		// Save data into Terraform state
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		FederatedBlockAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics, true)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create FederatedBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FederatedBlockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FederatedBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMFederationAPI.
		FederatedBlockAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read FederatedBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FederatedBlockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FederatedBlockModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		FederatedBlockAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update FederatedBlock, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FederatedBlockResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FederatedBlockModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Blocks in cloud provider backed pools (e.g. AWS) can fail to deprovision while the provider is still
	// working on the CIDR, so retry on that specific error.
	err := retry.RetryContext(ctx, 5*time.Minute, func() *retry.RetryError {
		httpRes, err := r.client.IPAMFederationAPI.
			FederatedBlockAPI.
			Delete(ctx, data.Id.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "Failed to deprovision") {
				return retry.RetryableError(err)
			}
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete FederatedBlock, got error: %s", err))
		return
	}
}

func (r *FederatedBlockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
func (r *FederatedBlockResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FederatedBlockModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.NetworkCompliance.IsNull() || data.NetworkCompliance.IsUnknown() {
		return
	}

	var nc NetworkComplianceModel
	resp.Diagnostics.Append(data.NetworkCompliance.As(ctx, &nc, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !nc.MinimumNetmaskLength.IsNull() && !nc.MinimumNetmaskLength.IsUnknown() &&
		!data.Cidr.IsNull() && !data.Cidr.IsUnknown() {
		min := nc.MinimumNetmaskLength.ValueInt64()
		cidr := data.Cidr.ValueInt64()

		if min <= cidr {
			resp.Diagnostics.AddAttributeError(
				path.Root("network_compliance").AtName("minimum_netmask_length"),
				"Invalid Network Compliance",
				fmt.Sprintf("\"minimum_netmask_length\" must be greater than the block's own \"cidr\" (/%d).", cidr),
			)
		}
	}

	if nc.MinimumNetmaskLength.IsNull() || nc.MinimumNetmaskLength.IsUnknown() ||
		nc.MaximumNetmaskLength.IsNull() || nc.MaximumNetmaskLength.IsUnknown() {
		return
	}

	min := nc.MinimumNetmaskLength.ValueInt64()
	max := nc.MaximumNetmaskLength.ValueInt64()

	if min >= max {
		resp.Diagnostics.AddAttributeError(
			path.Root("network_compliance").AtName("minimum_netmask_length"),
			"Invalid Network Compliance",
			"\"minimum_netmask_length\" must be less than \"maximum_netmask_length\".",
		)
		resp.Diagnostics.AddAttributeError(
			path.Root("network_compliance").AtName("maximum_netmask_length"),
			"Invalid Network Compliance",
			"\"maximum_netmask_length\" must be greater than \"minimum_netmask_length\".",
		)
	}

	if nc.DefaultNetmaskLength.IsNull() || nc.DefaultNetmaskLength.IsUnknown() {
		return
	}
	def := nc.DefaultNetmaskLength.ValueInt64()

	if def > max {
		resp.Diagnostics.AddAttributeError(
			path.Root("network_compliance").AtName("default_netmask_length"),
			"Invalid Network Compliance",
			"\"default_netmask_length\" must be less than or equal to \"maximum_netmask_length\".",
		)
	}
	if def < min {
		resp.Diagnostics.AddAttributeError(
			path.Root("network_compliance").AtName("default_netmask_length"),
			"Invalid Network Compliance",
			"\"default_netmask_length\" must be greater than or equal to \"minimum_netmask_length\".",
		)
	}
}

// ModifyPlan rejects network_compliance at plan time for a new block in an AWS pool. A block takes its provider type
// from its pool, so the pool is read. If the pool is not known yet (it is created in the same run), the block is
// checked during apply instead, in Create. Existing blocks are not checked here, so that a destroy is never
// blocked by a pending change in the configuration; the API rejects any update of an AWS block at apply.
func (r *FederatedBlockResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Only new blocks: not on destroy, not for blocks that already exist, and not when the provider is not configured.
	if req.Plan.Raw.IsNull() || !req.State.Raw.IsNull() || r.client == nil {
		return
	}

	var config FederatedBlockModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() || config.NetworkCompliance.IsNull() || config.NetworkCompliance.IsUnknown() {
		return
	}
	poolId, ok := config.resolvePoolId()
	if !ok {
		return
	}

	poolRes, _, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Read(ctx, poolId).
		Execute()
	if err != nil {
		// If the pool cannot be read, the create request reports the problem, so it is not checked here.
		return
	}
	pool := poolRes.GetResult()

	if pool.GetProvider() == ipamfederation.PROVIDERTYPE_AWS {
		resp.Diagnostics.AddAttributeError(path.Root("network_compliance"), "Unsupported attribute for provider type",
			"network_compliance is not supported for blocks in a pool with provider_type \"AWS\".")
	}
}
