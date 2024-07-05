package ipam

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DhcpHostResource{}
var _ resource.ResourceWithImportState = &DhcpHostResource{}

type DhcpHostModelWithRetryAndTimeouts struct {
	Address          types.String   `tfsdk:"address"`
	AnycastAddresses types.List     `tfsdk:"anycast_addresses"`
	AssociatedServer types.Object   `tfsdk:"associated_server"`
	Comment          types.String   `tfsdk:"comment"`
	CurrentVersion   types.String   `tfsdk:"current_version"`
	Id               types.String   `tfsdk:"id"`
	IpSpace          types.String   `tfsdk:"ip_space"`
	Name             types.String   `tfsdk:"name"`
	Ophid            types.String   `tfsdk:"ophid"`
	ProviderId       types.String   `tfsdk:"provider_id"`
	Server           types.String   `tfsdk:"server"`
	Tags             types.Map      `tfsdk:"tags"`
	TagsAll          types.Map      `tfsdk:"tags_all"`
	Type             types.String   `tfsdk:"type"`
	RetryIfNotFound  types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts         timeouts.Value `tfsdk:"timeouts"`
}

var DhcpHostResourceSchemaAttributesWithRetry = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The primary IP address of the on-prem host.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          IpamsvcHostAssociatedServerResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The DHCP Config Profile for the on-prem host.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The description for the on-prem host.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Current dhcp application version of the host.",
	},
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_space": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The display name of the on-prem host.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The on-prem host ID.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the on-prem host in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the on-prem host in JSON format including default tags.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.",
	},
	"retry_if_not_found": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to `true`, the resource will retry until a matching host is found, or until the Create Timeout expires.",
	},
}

func NewDhcpHostResource() resource.Resource {
	return &DhcpHostResource{}
}

// DhcpHostResource defines the resource implementation.
type DhcpHostResource struct {
	client *bloxoneclient.APIClient
}

func (r *DhcpHostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_host"
}

func (r *DhcpHostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages DHCP Hosts.\n\nA DHCP Host object associates a DHCP Config Profile with an on-prem host.\n\nNote: This resource represents an existing backend object that cannot be created or deleted through API calls. Instead, it can only be updated. When using terraform apply the resource configuration is applied to the existing object, and no new object is created. Similarly terraform destroy removes the configuration associated with the object without actually deleting it from the backend.",
		Attributes:          DhcpHostResourceSchemaAttributesWithRetry,
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
			}),
		},
	}
}

func (r *DhcpHostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DhcpHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DhcpHostModelWithRetryAndTimeouts

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	readTimeout, diags := data.Timeouts.Create(ctx, 20*time.Minute)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, readTimeout, func() *retry.RetryError {
		_, httpRes, err := r.client.IPAddressManagementAPI.
			DhcpHostAPI.
			Read(ctx, data.Id.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				if data.RetryIfNotFound.ValueBool() {
					return retry.RetryableError(err)
				}
				return retry.NonRetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DhcpHost, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}

	err = retry.RetryContext(ctx, readTimeout, func() *retry.RetryError {
		apiRes, httpRes, err := r.client.IPAddressManagementAPI.
			DhcpHostAPI.
			Update(ctx, data.Id.ValueString()).
			Body(*data.Expand(ctx, &resp.Diagnostics)).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				if data.RetryIfNotFound.ValueBool() {
					return retry.RetryableError(err)
				}
				return retry.NonRetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create DhcpHost, got error: %s", err))
			return retry.NonRetryableError(err)
		}

		res := apiRes.GetResult()
		data.Flatten(ctx, &res, &resp.Diagnostics)

		return nil
	})
	if err != nil {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DhcpHostModelWithRetryAndTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.IPAddressManagementAPI.
		DhcpHostAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DhcpHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DhcpHostModelWithRetryAndTimeouts

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.IPAddressManagementAPI.
		DhcpHostAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update DhcpHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DhcpHostModelWithRetryAndTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Server = types.StringNull()

	apiRes, _, err := r.client.IPAddressManagementAPI.
		DhcpHostAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete DhcpHost, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DhcpHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *DhcpHostModelWithRetryAndTimeouts) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Host {
	if m == nil {
		return nil
	}
	to := &ipam.Host{
		Server: flex.ExpandStringPointer(m.Server),
	}
	return to
}

func (m *DhcpHostModelWithRetryAndTimeouts) Flatten(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DhcpHostModelWithRetryAndTimeouts{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenIpamsvcHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
