package dfp

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/infobloxopen/bloxone-go-client/dfp"
	"net/http"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
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
	resp.TypeName = req.ProviderTypeName + "_" + "td_dfp_service"
}

func (r *DfpResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a DNS Forwarding Proxy.",
		Attributes:          AtcdfpDfpResourceSchemaAttributes,
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
	var data AtcdfpDfpModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	//err := retry.RetryContext(ctx, 20*time.Minute, func() *retry.RetryError {
	//	//get  id from infra
	//	hostRes, _, err := r.client.InfraManagementAPI.
	//		HostsAPI.
	//		HostsList(ctx).
	//		Filter(fmt.Sprintf("legacy_id == '%d'", data.Id.ValueInt64())).Execute()
	//	if err != nil {
	//		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create OnPremAnycastManager, got error: %s", err))
	//		return retry.NonRetryableError(err)
	//	}
	//
	//	//now you add the readonly data
	//	data.Name = types.StringValue(hostRes.GetResults()[0].DisplayName)
	//	data.IpAddress = types.StringPointerValue(hostRes.GetResults()[0].IpAddress)
	//	return nil
	//})

	apiRes, _, err := r.client.DNSForwardingProxyAPI.
		DfpAPI.
		DfpCreateOrUpdateDfp(ctx, int32(data.Id.ValueInt64())).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
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
	var data AtcdfpDfpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSForwardingProxyAPI.
		DfpAPI.
		DfpReadDfp(ctx, int32(data.Id.ValueInt64())).
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
	var data AtcdfpDfpModel
	var data2 = dfp.NewAtcdfpDfpCreateOrUpdatePayload()

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSForwardingProxyAPI.
		DfpAPI.
		DfpCreateOrUpdateDfp(ctx, int32(data.Id.ValueInt64())).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
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
	var data AtcdfpDfpModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.State.RemoveResource(ctx)
}

func (r *DfpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
