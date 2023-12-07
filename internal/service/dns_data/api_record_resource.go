package dns_data

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

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

const (
	// RecordOperationTimeout is the maximum amount of time to wait for eventual consistency
	RecordOperationTimeout = 2 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RecordResource{}
var _ resource.ResourceWithImportState = &RecordResource{}

// RecordResource defines the resource implementation.
type RecordResource struct {
	client *bloxoneclient.APIClient
	impl   recordResourceImplementor
}

func newRecordResource(impl recordResourceImplementor) resource.Resource {
	return &RecordResource{
		impl: impl,
	}
}

func (r *RecordResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *RecordResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + r.impl.resourceName()
}

func (r *RecordResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          r.impl.schemaAttributes(),
	}
}

func (r *RecordResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data dataRecordModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, RecordOperationTimeout, func() *retry.RetryError {
		apiRes, _, err := r.client.DNSDataAPI.
			RecordAPI.
			RecordCreate(ctx).
			Body(*data.Expand(ctx, &resp.Diagnostics, true, r.impl)).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				tflog.Debug(ctx, "Waiting for related objects to be present, will retry", map[string]interface{}{"error": err.Error()})
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Record, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		res := apiRes.GetResult()
		data.Flatten(ctx, &res, &resp.Diagnostics, r.impl)

		return nil
	})
	if err != nil {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data dataRecordModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSDataAPI.
		RecordAPI.
		RecordRead(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Record, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics, r.impl)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data dataRecordModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, RecordOperationTimeout, func() *retry.RetryError {
		apiRes, _, err := r.client.DNSDataAPI.
			RecordAPI.
			RecordUpdate(ctx, data.Id.ValueString()).
			Body(*data.Expand(ctx, &resp.Diagnostics, false, r.impl)).
			Execute()
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				tflog.Debug(ctx, "Waiting for related objects to be present, will retry", map[string]interface{}{"error": err.Error()})
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Record, got error: %s", err))
			return retry.NonRetryableError(err)
		}

		res := apiRes.GetResult()
		data.Flatten(ctx, &res, &resp.Diagnostics, r.impl)
		return nil
	})
	if err != nil {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RecordResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data dataRecordModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSDataAPI.
		RecordAPI.
		RecordDelete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Record, got error: %s", err))
		return
	}
}

func (r *RecordResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
