package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type CredentialConfigModel struct {
	AccessIdentifier types.String `tfsdk:"access_identifier"`
	Enclave          types.String `tfsdk:"enclave"`
	Region           types.String `tfsdk:"region"`
}

var CredentialConfigAttrTypes = map[string]attr.Type{
	"access_identifier": types.StringType,
	"enclave":           types.StringType,
	"region":            types.StringType,
}

var CredentialConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"access_identifier": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Access identifier of the account",
	},
	"enclave": schema.StringAttribute{
		Optional: true,
	},
	"region": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandCredentialConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.CredentialConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CredentialConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CredentialConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.CredentialConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.CredentialConfig{
		AccessIdentifier: flex.ExpandStringPointer(m.AccessIdentifier),
		Enclave:          flex.ExpandStringPointer(m.Enclave),
		Region:           flex.ExpandStringPointer(m.Region),
	}
	return to
}

func FlattenCredentialConfig(ctx context.Context, from *clouddiscovery.CredentialConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CredentialConfigAttrTypes)
	}
	m := CredentialConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CredentialConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CredentialConfigModel) Flatten(ctx context.Context, from *clouddiscovery.CredentialConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CredentialConfigModel{}
	}
	m.AccessIdentifier = flex.FlattenStringPointer(from.AccessIdentifier)
	m.Enclave = flex.FlattenStringPointer(from.Enclave)
	m.Region = flex.FlattenStringPointer(from.Region)
}
