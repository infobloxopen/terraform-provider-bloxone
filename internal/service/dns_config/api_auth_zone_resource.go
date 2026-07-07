package dns_config

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
	// AuthZoneOperationTimeout is the maximum amount of time to wait for eventual consistency
	AuthZoneOperationTimeout = 2 * time.Minute
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthZoneResource{}
var _ resource.ResourceWithImportState = &AuthZoneResource{}
var _ resource.ResourceWithValidateConfig = &AuthZoneResource{}

var inheritanceType = "full"

func NewAuthZoneResource() resource.Resource {
	return &AuthZoneResource{}
}

// AuthZoneResource defines the resource implementation.
type AuthZoneResource struct {
	client *bloxoneclient.APIClient
}

func (r *AuthZoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_auth_zone"
}

func (r *AuthZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: `Manages an authoritative zone.`,
		Attributes:          ConfigAuthZoneResourceSchemaAttributes,
	}
}

func (r *AuthZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AuthZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ConfigAuthZoneModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSConfigurationAPI.
		AuthZoneAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics, true)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AuthZone, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ConfigAuthZoneModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSConfigurationAPI.
		AuthZoneAPI.
		Read(ctx, data.Id.ValueString()).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AuthZone, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ConfigAuthZoneModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSConfigurationAPI.
		AuthZoneAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics, false)).
		Inherit(inheritanceType).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AuthZone, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ConfigAuthZoneModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, AuthZoneOperationTimeout, func() *retry.RetryError {
		httpRes, err := r.client.DNSConfigurationAPI.
			AuthZoneAPI.
			Delete(ctx, data.Id.ValueString()).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				return nil
			}
			if strings.Contains(err.Error(), "object is referenced by a 'Zone' object") {
				return retry.RetryableError(err)
			}
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AuthZone, got error: %s", err))
			return retry.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return
	}
}

func (r *AuthZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthZoneResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ConfigAuthZoneModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nameserversConfigured := !data.Nameservers.IsNull()
	nsgConfigured := !data.Nsg.IsNull()
	primaryTypeConfigured := !data.PrimaryType.IsNull()

	if (nameserversConfigured || nsgConfigured) && primaryTypeConfigured {
		resp.Diagnostics.AddAttributeError(
			path.Root("primary_type"),
			"Invalid attribute combination",
			"Primary Type cannot be provided unified nameservers is enabled",
		)
	}

	if nameserversConfigured && nsgConfigured {
		resp.Diagnostics.AddAttributeError(
			path.Root("nameservers"),
			"Invalid attribute combination",
			"\"nameservers\" and \"nsg\" cannot both be provided. Use one or the other to configure nameservers.",
		)
	}

	if !data.ExternalSecondaries.IsNull() && !data.ExternalSecondaries.IsUnknown() {
		var externalSecondaries []ConfigExternalSecondaryModel
		resp.Diagnostics.Append(data.ExternalSecondaries.ElementsAs(ctx, &externalSecondaries, false)...)
		if !resp.Diagnostics.HasError() {
			for i, es := range externalSecondaries {
				if es.Address.IsNull() || es.Address.IsUnknown() || es.Address.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("external_secondaries").AtListIndex(i).AtName("address"),
						"Missing required attribute",
						"When external_secondaries is configured, \"address\" must be provided for each entry.",
					)
				}
				if es.Fqdn.IsNull() || es.Fqdn.IsUnknown() || es.Fqdn.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("external_secondaries").AtListIndex(i).AtName("fqdn"),
						"Missing required attribute",
						"When external_secondaries is configured, \"fqdn\" must be provided for each entry.",
					)
				}
			}
		}
	}

	if !data.GridPrimaries.IsNull() && !data.GridPrimaries.IsUnknown() {
		var gridPrimaries []ConfigMemberServerModel
		resp.Diagnostics.Append(data.GridPrimaries.ElementsAs(ctx, &gridPrimaries, false)...)
		if !resp.Diagnostics.HasError() {
			for i, gp := range gridPrimaries {
				if gp.Host.IsNull() || gp.Host.IsUnknown() || gp.Host.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("grid_primaries").AtListIndex(i).AtName("host"),
						"Missing required attribute",
						"When grid_primaries is configured, \"host\" must be provided for each entry.",
					)
				}
			}
		}
	}

	if !data.GridSecondaries.IsNull() && !data.GridSecondaries.IsUnknown() {
		var gridSecondaries []ConfigMemberServerModel
		resp.Diagnostics.Append(data.GridSecondaries.ElementsAs(ctx, &gridSecondaries, false)...)
		if !resp.Diagnostics.HasError() {
			for i, gs := range gridSecondaries {
				if gs.Host.IsNull() || gs.Host.IsUnknown() || gs.Host.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("grid_secondaries").AtListIndex(i).AtName("host"),
						"Missing required attribute",
						"When grid_secondaries is configured, \"host\" must be provided for each entry.",
					)
				}
			}
		}
	}

	if !data.InternalSecondaries.IsNull() && !data.InternalSecondaries.IsUnknown() {
		var internalSecondaries []ConfigInternalSecondaryModel
		resp.Diagnostics.Append(data.InternalSecondaries.ElementsAs(ctx, &internalSecondaries, false)...)
		if !resp.Diagnostics.HasError() {
			for i, is := range internalSecondaries {
				if is.Host.IsNull() || is.Host.IsUnknown() || is.Host.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("internal_secondaries").AtListIndex(i).AtName("host"),
						"Missing required attribute",
						"When internal_secondaries is configured, \"host\" must be provided for each entry.",
					)
				}
			}
		}
	}

	if !data.ExternalPrimaries.IsNull() && !data.ExternalPrimaries.IsUnknown() {
		var externalPrimaries []ConfigExternalPrimaryModel
		resp.Diagnostics.Append(data.ExternalPrimaries.ElementsAs(ctx, &externalPrimaries, false)...)
		if !resp.Diagnostics.HasError() {
			for i, ep := range externalPrimaries {
				if ep.Type.IsNull() || ep.Type.IsUnknown() || ep.Type.ValueString() == "" {
					resp.Diagnostics.AddAttributeError(
						path.Root("external_primaries").AtListIndex(i).AtName("type"),
						"Missing required attribute",
						"When external_primaries is configured, \"type\" must be provided for each entry.",
					)
				}
			}
		}
	}
}
