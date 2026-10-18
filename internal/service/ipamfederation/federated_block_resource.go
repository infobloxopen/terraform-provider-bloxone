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
)

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
	if !data.NetworkCompliance.IsNull() && !data.NetworkCompliance.IsUnknown() && !data.FederatedPoolId.IsNull() && !data.FederatedPoolId.IsUnknown() {
		poolRes, _, err := r.client.IPAMFederationAPI.
			FederatedPoolAPI.
			Read(ctx, data.FederatedPoolId.ValueString()).
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
	if resp.Diagnostics.HasError() || config.NetworkCompliance.IsNull() || config.NetworkCompliance.IsUnknown() ||
		config.FederatedPoolId.IsNull() || config.FederatedPoolId.IsUnknown() {
		return
	}

	poolRes, _, err := r.client.IPAMFederationAPI.
		FederatedPoolAPI.
		Read(ctx, config.FederatedPoolId.ValueString()).
		Execute()
	if err != nil {
		// Create reports a pool that cannot be read.
		return
	}
	pool := poolRes.GetResult()

	if pool.GetProvider() == ipamfederation.PROVIDERTYPE_AWS {
		resp.Diagnostics.AddAttributeError(path.Root("network_compliance"), "Unsupported attribute for provider type",
			"network_compliance is not supported for blocks in a pool with provider_type \"AWS\".")
	}
}
