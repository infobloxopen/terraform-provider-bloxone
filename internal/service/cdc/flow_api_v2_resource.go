package cdc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FlowApiV2Resource{}
var _ resource.ResourceWithImportState = &FlowApiV2Resource{}

func NewFlowApiV2Resource() resource.Resource {
	return &FlowApiV2Resource{}
}

// FlowApiV2Resource defines the resource implementation.
type FlowApiV2Resource struct {
	client *universalddiclient.APIClient
}

func (r *FlowApiV2Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "cdc_flow"
}

func (r *FlowApiV2Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages a Data Flow.\n\nThe Data Flow object connects a source, one or more destinations, and optional ETL filters to define how data moves through Cloud Data Connector (CDC).",
		Attributes:          FlowDataV2ResourceSchemaAttributes,
	}
}

func (r *FlowApiV2Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FlowApiV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FlowDataV2Model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.CDCAPI.
		FlowApiV2API.
		CreateFlowDataV2(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create FlowApiV2, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	r.readBack(ctx, &res, &resp.Diagnostics)
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FlowApiV2Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FlowDataV2Model

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.CDCAPI.
		FlowApiV2API.
		ReadFlowDataV2(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read FlowApiV2, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	r.readSchedule(ctx, &res, &resp.Diagnostics)
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FlowApiV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data FlowDataV2Model

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	body := data.Expand(ctx, &resp.Diagnostics)

	apiRes, _, err := r.client.CDCAPI.
		FlowApiV2API.
		UpdateFlowDataV2(ctx, data.Id.ValueInt64()).
		Body(*body).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update FlowApiV2, got error: %s", err))
		return
	}

	res := apiRes.GetResults()
	r.readBack(ctx, &res, &resp.Diagnostics)
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FlowApiV2Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FlowDataV2Model

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.CDCAPI.
		FlowApiV2API.
		DeleteFlowDataV2(ctx, data.Id.ValueInt64()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete FlowApiV2, got error: %s", err))
		return
	}
}

func (r *FlowApiV2Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	utils.ImportStatePassthroughInt64ID(ctx, req, resp)
}

// readBack replaces res with the flow as read from the backend. The create and
// update responses do not fully populate source_log_types (log_type comes back
// empty), so the state would otherwise differ from what a later read or import
// returns.
func (r *FlowApiV2Resource) readBack(ctx context.Context, res *cdc.FlowDataV2, diags *diag.Diagnostics) {
	if res.Id == nil {
		return
	}
	apiRes, _, err := r.client.CDCAPI.
		FlowApiV2API.
		ReadFlowDataV2(ctx, *res.Id).
		Execute()
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read FlowApiV2, got error: %s", err))
		return
	}
	*res = apiRes.GetResults()
	r.readSchedule(ctx, res, diags)
}

// readSchedule fills schedule_option and script_schedule from the display
// endpoint: /v2/flows/data always returns them empty, even when they are set.
func (r *FlowApiV2Resource) readSchedule(ctx context.Context, res *cdc.FlowDataV2, diags *diag.Diagnostics) {
	if res.Id == nil {
		return
	}
	apiRes, _, err := r.client.CDCAPI.
		FlowApiV2API.
		ReadFlowV2(ctx, *res.Id).
		Execute()
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read FlowApiV2, got error: %s", err))
		return
	}
	display := apiRes.GetResults()
	res.ScheduleOption = display.ScheduleOption
	res.ScriptSchedule = display.ScriptSchedule
}
