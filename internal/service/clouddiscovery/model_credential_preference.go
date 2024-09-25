package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type CredentialPreferenceModel struct {
	AccessIdentifierType types.String `tfsdk:"access_identifier_type"`
	CredentialType       types.String `tfsdk:"credential_type"`
}

var CredentialPreferenceAttrTypes = map[string]attr.Type{
	"access_identifier_type": types.StringType,
	"credential_type":        types.StringType,
}

var CredentialPreferenceResourceSchemaAttributes = map[string]schema.Attribute{
	"access_identifier_type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("role_arn", "tenant_id", "project_id"),
		},
		MarkdownDescription: "Access identifier type. Possible values: role_arn, tenant_id, project_id.",
	},
	"credential_type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("dynamic"),
		},
		MarkdownDescription: "Credential type. Possible values: `dynamic`. Support for Static Credentials is not present",
	},
}

func ExpandCredentialPreference(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.CredentialPreference {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CredentialPreferenceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CredentialPreferenceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.CredentialPreference {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.CredentialPreference{
		AccessIdentifierType: flex.ExpandStringPointer(m.AccessIdentifierType),
		CredentialType:       flex.ExpandStringPointer(m.CredentialType),
	}
	return to
}

func FlattenCredentialPreference(ctx context.Context, from *clouddiscovery.CredentialPreference, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CredentialPreferenceAttrTypes)
	}
	m := CredentialPreferenceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CredentialPreferenceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CredentialPreferenceModel) Flatten(ctx context.Context, from *clouddiscovery.CredentialPreference, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CredentialPreferenceModel{}
	}
	m.AccessIdentifierType = flex.FlattenStringPointer(from.AccessIdentifierType)
	m.CredentialType = flex.FlattenStringPointer(from.CredentialType)
}
