package keys

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

const (
	// TsigKeyOperationTimeout is the maximum amount of time to wait for eventual consistency
	TsigKeyOperationTimeout = 2 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TsigResource{}
var _ resource.ResourceWithImportState = &TsigResource{}

func NewTsigResource() resource.Resource {
	return &TsigResource{}
}

// TsigResource defines the resource implementation.
type TsigResource struct {
	client *bloxoneclient.APIClient
}

func (r *TsigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "keys_tsig"
}

func (r *TsigResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a TSIG Key.",
		Attributes:          KeysTSIGKeyResourceSchemaAttributes,
	}
}

func (r *TsigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *TsigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data KeysTSIGKeyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Secret.IsUnknown() {
		apiRes, _, err := r.client.KeysAPI.
			GenerateTsigAPI.
			GenerateTsigGenerateTSIG(ctx).
			Algorithm(data.Algorithm.ValueString()).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to generate new secret, got error: %s", err))
			return
		}
		res := apiRes.GetResult()
		data.Secret = flex.FlattenString(res.GetSecret())
	}

	apiRes, _, err := r.client.KeysAPI.
		TsigAPI.
		TsigCreate(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Tsig, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TsigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data KeysTSIGKeyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.KeysAPI.
		TsigAPI.
		TsigRead(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Tsig, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TsigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data KeysTSIGKeyModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.KeysAPI.
		TsigAPI.
		TsigUpdate(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Tsig, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TsigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data KeysTSIGKeyModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, TsigKeyOperationTimeout, func() *retry.RetryError {
		httpRes, err := r.client.KeysAPI.
			TsigAPI.
			TsigDelete(ctx, data.Id.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "object is referenced by a 'TSIG Key' object.") {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Tsig, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (r *TsigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
