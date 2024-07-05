package dns_config

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HostResource{}
var _ resource.ResourceWithImportState = &HostResource{}

type HostModelWithRetryAndTimeouts struct {
	AbsoluteName         types.String   `tfsdk:"absolute_name"`
	Address              types.String   `tfsdk:"address"`
	AnycastAddresses     types.List     `tfsdk:"anycast_addresses"`
	AssociatedServer     types.Object   `tfsdk:"associated_server"`
	Comment              types.String   `tfsdk:"comment"`
	CurrentVersion       types.String   `tfsdk:"current_version"`
	Dfp                  types.Bool     `tfsdk:"dfp"`
	DfpService           types.String   `tfsdk:"dfp_service"`
	Id                   types.String   `tfsdk:"id"`
	InheritanceSources   types.Object   `tfsdk:"inheritance_sources"`
	KerberosKeys         types.List     `tfsdk:"kerberos_keys"`
	Name                 types.String   `tfsdk:"name"`
	Ophid                types.String   `tfsdk:"ophid"`
	ProtocolAbsoluteName types.String   `tfsdk:"protocol_absolute_name"`
	ProviderId           types.String   `tfsdk:"provider_id"`
	Server               types.String   `tfsdk:"server"`
	SiteId               types.String   `tfsdk:"site_id"`
	Tags                 types.Map      `tfsdk:"tags"`
	TagsAll              types.Map      `tfsdk:"tags_all"`
	Type                 types.String   `tfsdk:"type"`
	RetryIfNotFound      types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts             timeouts.Value `tfsdk:"timeouts"`
}

var HostResourceSchemaAttributesWithRetryAndTimeouts = map[string]schema.Attribute{
	"absolute_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host FQDN.",
	},
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host's primary IP Address.",
	},
	"anycast_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
	},
	"associated_server": schema.SingleNestedAttribute{
		Attributes:          ConfigHostAssociatedServerResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "Host associated server configuration.",
	},
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host description.",
	},
	"current_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host current version.",
	},
	"dfp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Below _dfp_ field is deprecated and not supported anymore. The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.",
	},
	"dfp_service": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP. Allowed values:  * _unavailable_: BloxOne TD DFP application is not available,  * _enabled_: BloxOne TD DFP application is available and enabled,  * _disabled_: BloxOne TD DFP application is available but disabled.",
	},
	"id": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          ConfigHostInheritanceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Inheritance configuration.",
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ConfigKerberosKeyResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host display name.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "On-Prem Host ID.",
	},
	"protocol_absolute_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host FQDN in punycode.",
	},
	"provider_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "External provider identifier.",
	},
	"server": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"site_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Host site ID.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Host tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Host tagging specifics includes default tags.",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.",
	},
	"retry_if_not_found": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to `true`, the resource will retry until a matching host is found, or until the Create Timeout expires.",
	},
}

func NewHostResource() resource.Resource {
	return &HostResource{}
}

// HostResource defines the resource implementation.
type HostResource struct {
	client *bloxoneclient.APIClient
}

func (r *HostResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_host"
}

func (r *HostResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages DNS Hosts.\n\nA DNS Host object associates a DNS Config Profile with an on-prem host.\n\nNote: This resource represents an existing backend object that cannot be created or deleted through API calls. Instead, it can only be updated. When using terraform apply the resource configuration is applied to the existing object, and no new object is created. Similarly terraform destroy removes the configuration associated with the object without actually deleting it from the backend.",
		Attributes:          HostResourceSchemaAttributesWithRetryAndTimeouts,
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(ctx, timeouts.Opts{
				Create: true,
			}),
		},
	}
}

func (r *HostResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *HostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HostModelWithRetryAndTimeouts

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
		_, httpRes, err := r.client.DNSConfigurationAPI.
			HostAPI.
			Read(ctx, data.Id.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				if data.RetryIfNotFound.ValueBool() {
					return retry.RetryableError(err)
				}
				return retry.NonRetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Host, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}

	err = retry.RetryContext(ctx, readTimeout, func() *retry.RetryError {
		apiRes, httpRes, err := r.client.DNSConfigurationAPI.
			HostAPI.
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
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create Host, got error: %s", err))
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

func (r *HostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HostModelWithRetryAndTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSConfigurationAPI.
		HostAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Host, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HostModelWithRetryAndTimeouts

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSConfigurationAPI.
		HostAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update Host, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HostModelWithRetryAndTimeouts

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Server = types.StringNull()

	apiRes, _, err := r.client.DNSConfigurationAPI.
		HostAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete Host, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *HostModelWithRetryAndTimeouts) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Host {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Host{
		Server: flex.ExpandStringPointer(m.Server),
	}
	return to
}

func (m *HostModelWithRetryAndTimeouts) Flatten(ctx context.Context, from *dnsconfig.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HostModelWithRetryAndTimeouts{}
	}
	m.AbsoluteName = flex.FlattenStringPointer(from.AbsoluteName)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenConfigHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.Dfp = types.BoolPointerValue(from.Dfp)
	m.DfpService = flex.FlattenStringPointer(from.DfpService)
	m.InheritanceSources = FlattenConfigHostInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, ConfigKerberosKeyAttrTypes, diags, FlattenConfigKerberosKey)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProtocolAbsoluteName = flex.FlattenStringPointer(from.ProtocolAbsoluteName)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
