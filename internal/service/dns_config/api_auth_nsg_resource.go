package dns_config

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuthNsgResource{}
var _ resource.ResourceWithImportState = &AuthNsgResource{}
var _ resource.ResourceWithValidateConfig = &AuthNsgResource{}

func NewAuthNsgResource() resource.Resource {
	return &AuthNsgResource{}
}

// AuthNsgResource defines the resource implementation.
type AuthNsgResource struct {
	client *universalddiclient.APIClient
}

func (r *AuthNsgResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_auth_nsg"
}

func (r *AuthNsgResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Manages an Authoritative DNS Server Group for authoritative zones.",
		Attributes:          ConfigAuthNSGResourceSchemaAttributes,
	}
}

func (r *AuthNsgResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AuthNsgResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ConfigAuthNSGModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSConfigurationAPI.
		AuthNsgAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create AuthNsg, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthNsgResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ConfigAuthNSGModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DNSConfigurationAPI.
		AuthNsgAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read AuthNsg, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthNsgResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ConfigAuthNSGModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DNSConfigurationAPI.
		AuthNsgAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update AuthNsg, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthNsgResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ConfigAuthNSGModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DNSConfigurationAPI.
		AuthNsgAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete AuthNsg, got error: %s", err))
		return
	}
}

func (r *AuthNsgResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data ConfigAuthNSGModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
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
				if !gp.Host.IsUnknown() && (gp.Host.IsNull() || gp.Host.ValueString() == "") {
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
				if !gs.Host.IsUnknown() && (gs.Host.IsNull() || gs.Host.ValueString() == "") {
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
				if !is.Host.IsUnknown() && (is.Host.IsNull() || is.Host.ValueString() == "") {
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

	if !data.Nameservers.IsNull() && !data.Nameservers.IsUnknown() {
		var nameservers []ConfigNameserverModel
		resp.Diagnostics.Append(data.Nameservers.ElementsAs(ctx, &nameservers, false)...)
		if !resp.Diagnostics.HasError() {
			for i, ns := range nameservers {
				if !ns.Stealth.IsNull() && !ns.Stealth.IsUnknown() && ns.Stealth.ValueBool() &&
					!ns.Role.IsNull() && !ns.Role.IsUnknown() && ns.Role.ValueString() != "secondary" {
					resp.Diagnostics.AddAttributeError(
						path.Root("nameservers").AtListIndex(i).AtName("stealth"),
						"Invalid attribute combination",
						"\"stealth\" can only be set when \"role\" is \"secondary\".",
					)
				}
				if !ns.TsigEnabled.IsNull() && !ns.TsigEnabled.IsUnknown() && ns.TsigEnabled.ValueBool() &&
					(ns.TsigKey.IsNull() || ns.TsigKey.IsUnknown()) {
					resp.Diagnostics.AddAttributeError(
						path.Root("nameservers").AtListIndex(i).AtName("tsig_key"),
						"Missing required attribute",
						"\"tsig_key\" must be provided when \"tsig_enabled\" is true.",
					)
				}
			}
		}
	}
}

func (r *AuthNsgResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
