package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AuthZoneExternalProviderModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

var AuthZoneExternalProviderAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
	"type": types.StringType,
}

var AuthZoneExternalProviderResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The identifier of the external provider.`,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the external provider.`,
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Defines the type of external provider. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.`,
	},
}

func ExpandAuthZoneExternalProvider(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.AuthZoneExternalProvider {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AuthZoneExternalProviderModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AuthZoneExternalProviderModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.AuthZoneExternalProvider {
	if m == nil {
		return nil
	}
	to := &dns_config.AuthZoneExternalProvider{
		Name: flex.ExpandStringPointer(m.Name),
		Type: flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenAuthZoneExternalProvider(ctx context.Context, from *dns_config.AuthZoneExternalProvider, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AuthZoneExternalProviderAttrTypes)
	}
	m := AuthZoneExternalProviderModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AuthZoneExternalProviderAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AuthZoneExternalProviderModel) Flatten(ctx context.Context, from *dns_config.AuthZoneExternalProvider, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AuthZoneExternalProviderModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Type = flex.FlattenStringPointer(from.Type)
}
