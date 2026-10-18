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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FederatedPoolResource{}
var _ resource.ResourceWithModifyPlan = &FederatedPoolResource{}
var _ resource.ResourceWithImportState = &FederatedPoolResource{}

func NewFederatedPoolResource() resource.Resource {
	return &FederatedPoolResource{}
}

// FederatedPoolResource defines the resource implementation.
type FederatedPoolResource struct {
	client *universalddiclient.APIClient
}

func (r *FederatedPoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "federation_federated_pool"
}

func (r *FederatedPoolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages the Federated Pool.\n\nA Federated Pool is a pool of addresses allocated within a federated realm, optionally under a parent pool.",
		Attributes:          FederatedPoolResourceSchemaAttributes,
	}
}

func (r *FederatedPoolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FederatedPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FederatedPoolModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// provider_type and region are computed: a pool always has the provider type and region of its realm.
	realmRes, _, err := r.client.IPAMFederationAPI.
		FederatedRealmAPI.
		Read(ctx, data.FederatedRealm.ValueString()).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read FederatedRealm %s for the FederatedPool, got error: %s", data.FederatedRealm.ValueString(), err))
		return
	}
	realm := realmRes.GetResult()
	providerType := string(realm.GetProvider())
	if providerType != "NIOS_X" && providerType != "AWS" {
		resp.Diagnostics.AddError("Unsupported provider type", fmt.Sprintf("The federated realm has provider_type %q. Only NIOS_X and AWS are supported.", providerType))
		return
	}
	data.ProviderType = types.StringValue(providerType)
	data.Region = types.StringValue(realm.GetRegion())

	body := data.Expand(ctx, &resp.Diagnostics, true)
	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Create(ctx).
		Body(*body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create FederatedPool, got error: %s", err))
		return
	}

	res := apiRes.GetResult()

	// Pools backed by a cloud provider are created asynchronously: poll until the pool is ready so that
	// blocks can be created in it. The pool already exists at this point, so on failure it is still saved
	// to state (tainted) instead of being left unknown to Terraform.
	waitErr := retry.RetryContext(ctx, 10*time.Minute, func() *retry.RetryError {
		readRes, _, err := r.client.IPAMFederationAPI.
			FederatedPoolAPI.
			Read(ctx, res.GetId()).
			Execute()
		if err != nil {
			return retry.NonRetryableError(err)
		}
		res = readRes.GetResult()
		state := res.GetState()
		switch {
		case state == "create-in-progress":
			return retry.RetryableError(fmt.Errorf("federated pool %s is still in state %q", res.GetId(), state))
		case strings.Contains(state, "fail") || strings.Contains(state, "error"):
			return retry.NonRetryableError(fmt.Errorf("federated pool %s ended up in state %q", res.GetId(), state))
		}
		return nil
	})

	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	if waitErr != nil {
		resp.Diagnostics.AddError("Federated Pool Not Ready", fmt.Sprintf("FederatedPool %s was created but did not become ready: %s", res.GetId(), waitErr))
	}
}

func (r *FederatedPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FederatedPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read FederatedPool, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FederatedPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FederatedPoolModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Only the name and description of an AWS pool can be updated; the request carries nothing else, so reject a
	// change of the tags here instead of letting it be dropped, which would end in an "inconsistent result after apply".
	if data.ProviderType.ValueString() == "AWS" {
		var state FederatedPoolModel
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if !data.Tags.Equal(state.Tags) {
			resp.Diagnostics.AddError("Pool attribute cannot be updated", "Only the name and description of an AWS pool can be updated. tags cannot be changed.")
			return
		}
	}

	apiRes, _, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update FederatedPool, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FederatedPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FederatedPoolModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete FederatedPool, got error: %s", err))
		return
	}
}

func (r *FederatedPoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// ModifyPlan rejects, at plan time, what the API would reject for a new pool: a realm whose provider type is not
// supported (only NIOS_X and AWS are) and network_compliance on an AWS pool. The provider type is not part of the
// configuration (a pool has the provider type of its realm), so it is taken from the realm. If the realm is not
// known yet (it is created in the same run), the check happens during apply instead, in Create and Expand.
// Existing pools are not checked here, so that a destroy is never blocked by a pending change in the configuration.
func (r *FederatedPoolResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Only new pools: not on destroy, not for pools that already exist, and not when the provider is not configured.
	if req.Plan.Raw.IsNull() || !req.State.Raw.IsNull() || r.client == nil {
		return
	}

	var config FederatedPoolModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() || config.FederatedRealm.IsNull() || config.FederatedRealm.IsUnknown() {
		return
	}

	realmRes, _, err := r.client.IPAMFederationAPI.
		FederatedRealmAPI.
		Read(ctx, config.FederatedRealm.ValueString()).
		Execute()
	if err != nil {
		// Create reports a realm that cannot be read.
		return
	}
	realm := realmRes.GetResult()

	switch providerType := string(realm.GetProvider()); providerType {
	case "NIOS_X":
	case "AWS":
		if !config.NetworkCompliance.IsNull() && !config.NetworkCompliance.IsUnknown() {
			resp.Diagnostics.AddAttributeError(path.Root("network_compliance"), "Unsupported attribute for provider type",
				"network_compliance is not supported for pools with provider_type \"AWS\".")
		}
	default:
		resp.Diagnostics.AddError("Unsupported provider type", fmt.Sprintf("The federated realm has provider_type %q. Only NIOS_X and AWS are supported.", providerType))
	}
}
