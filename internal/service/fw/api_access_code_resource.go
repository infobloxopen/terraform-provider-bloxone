package fw

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

const (
	// InternalDomainListOperationTimeout is the maximum amount of time to wait for eventual consistency
	AccessCodeOperationTimeout = 2 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AccessCodeResource{}
var _ resource.ResourceWithImportState = &AccessCodeResource{}

func NewAccessCodeResource() resource.Resource {
	return &AccessCodeResource{}
}

// AccessCodeResource defines the resource implementation.
type AccessCodeResource struct {
	client *bloxoneclient.APIClient
}

func (r *AccessCodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "td_access_code"
}

func (r *AccessCodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an access code.",
		Attributes:          AtcfwAccessCodeResourceSchemaAttributes,
	}
}

func (r *AccessCodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AccessCodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AtcfwAccessCodeModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.FWAPI.
		AccessCodesAPI.
		CreateAccessCode(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AccessCodes, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AtcfwAccessCodeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.FWAPI.
		AccessCodesAPI.
		ReadAccessCode(ctx, data.AccessKey.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AccessCodes, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AtcfwAccessCodeModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.FWAPI.
		AccessCodesAPI.
		UpdateAccessCode(ctx, data.AccessKey.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AccessCodes, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AccessCodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AtcfwAccessCodeModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, AccessCodeOperationTimeout, func() *retry.RetryError {
		httpRes, err := r.client.FWAPI.
			AccessCodesAPI.
			DeleteSingleAccessCodes(ctx, data.AccessKey.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "Cannot delete bypass code assigned to policy") {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AccessCodes, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (r *AccessCodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
