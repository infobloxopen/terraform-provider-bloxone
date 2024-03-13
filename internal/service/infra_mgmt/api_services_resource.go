package infra_mgmt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServicesResource{}
var _ resource.ResourceWithImportState = &ServicesResource{}

func NewServicesResource() resource.Resource {
	return &ServicesResource{}
}

// ServicesResource defines the resource implementation.
type ServicesResource struct {
	client *bloxoneclient.APIClient
}

func (r *ServicesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "infra_service"
}

func (r *ServicesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Infrastructure Service (a.k.a BloxOne Applications).",
		Attributes:          InfraServiceResourceSchemaAttributesWithTimeouts(ctx),
	}
}

func (r *ServicesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ServicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InfraServiceModelWithTimeouts

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.InfraManagementAPI.
		ServicesAPI.
		ServicesCreate(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Services, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	createTimeout, diags := data.Timeouts.Create(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.WaitForState.ValueBool() {
		r.waitServiceStartStop(ctx, data.Name.ValueString(), data.DesiredState.ValueString(), createTimeout, &resp.Diagnostics)
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InfraServiceModelWithTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.InfraManagementAPI.
		ServicesAPI.
		ServicesRead(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Services, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data InfraServiceModelWithTimeouts

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.InfraManagementAPI.
		ServicesAPI.
		ServicesUpdate(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Services, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	updateTimeout, diags := data.Timeouts.Update(ctx, 20*time.Minute)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.WaitForState.ValueBool() {
		r.waitServiceStartStop(ctx, data.Name.ValueString(), data.DesiredState.ValueString(), updateTimeout, &resp.Diagnostics)
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InfraServiceModelWithTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.InfraManagementAPI.
		ServicesAPI.
		ServicesDelete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Services, got error: %s", err))
		return
	}
}

func (r *ServicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicesResource) waitServiceStartStop(ctx context.Context, name, desiredState string, timeout time.Duration, diags *diag.Diagnostics) {
	if desiredState == "start" {
		err := r.waitServiceStarted(ctx, name, timeout)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("waiting for service to be started, got error: %s", err))
			return
		}
	} else if desiredState == "stop" {
		err := r.waitServiceStopped(ctx, name, timeout)
		if err != nil {
			diags.AddError("Client Error", fmt.Sprintf("waiting for service to be stopped, got error: %s", err))
			return
		}
	}
}

func (r *ServicesResource) waitServiceStarted(ctx context.Context, name string, timeout time.Duration) error {
	scf := retry.StateChangeConf{
		Pending: []string{"starting", "error"}, // The service can be in a temporary error state, but is expected to recover.
		Target:  []string{"started"},
		Refresh: r.stateRefreshFunc(name),
		Timeout: timeout,
	}
	_, err := scf.WaitForStateContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ServicesResource) waitServiceStopped(ctx context.Context, name string, timeout time.Duration) error {
	scf := retry.StateChangeConf{
		Pending: []string{"stopping"},
		Target:  []string{"stopped"},
		Refresh: r.stateRefreshFunc(name),
		Timeout: timeout,
	}
	_, err := scf.WaitForStateContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *ServicesResource) stateRefreshFunc(name string) retry.StateRefreshFunc {
	return func() (result interface{}, state string, err error) {
		apiRes, _, err := r.client.InfraManagementAPI.
			DetailAPI.DetailServicesList(context.Background()).
			Filter(fmt.Sprintf("name=='%s'", name)).
			Execute()
		if err != nil {
			return
		}
		if len(apiRes.GetResults()) == 0 {
			return nil, "", errors.New("not found")
		}
		serviceDetail := apiRes.GetResults()[0]
		if serviceDetail.CompositeState == nil {
			return nil, "", errors.New("state not known")
		}
		return serviceDetail, *serviceDetail.CompositeState, nil
	}
}
