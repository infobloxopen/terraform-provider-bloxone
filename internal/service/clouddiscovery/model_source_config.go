package clouddiscovery

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SourceConfigModel struct {
	AccountScheduleId    types.String      `tfsdk:"account_schedule_id"`
	Accounts             types.List        `tfsdk:"accounts"`
	CloudCredentialId    types.String      `tfsdk:"cloud_credential_id"`
	CreatedAt            timetypes.RFC3339 `tfsdk:"created_at"`
	CredentialConfig     types.Object      `tfsdk:"credential_config"`
	DeletedAt            timetypes.RFC3339 `tfsdk:"deleted_at"`
	Id                   types.String      `tfsdk:"id"`
	RestrictedToAccounts types.List        `tfsdk:"restricted_to_accounts"`
	UpdatedAt            timetypes.RFC3339 `tfsdk:"updated_at"`
}

var SourceConfigAttrTypes = map[string]attr.Type{
	"account_schedule_id":    types.StringType,
	"accounts":               types.ListType{ElemType: types.ObjectType{AttrTypes: AccountAttrTypes}},
	"cloud_credential_id":    types.StringType,
	"created_at":             timetypes.RFC3339Type{},
	"credential_config":      types.ObjectType{AttrTypes: CredentialConfigAttrTypes},
	"deleted_at":             timetypes.RFC3339Type{},
	"id":                     types.StringType,
	"restricted_to_accounts": types.ListType{ElemType: types.StringType},
	"updated_at":             timetypes.RFC3339Type{},
}

var SourceConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"account_schedule_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Account Schedule ID.",
	},
	"accounts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AccountResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "List of accounts to be discovered.",
	},
	"cloud_credential_id": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Cloud Credential ID.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been created.",
	},
	"credential_config": schema.SingleNestedAttribute{
		Attributes:          CredentialConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Credential configuration For the discovery job.",
	},
	"deleted_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto-generated unique source config ID. Format BloxID.",
	},
	"restricted_to_accounts": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Provider account IDs such as accountID/ SubscriptionID to be restricted for a given source_config.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been updated.",
	},
}

func ExpandSourceConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.SourceConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SourceConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SourceConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.SourceConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.SourceConfig{
		CloudCredentialId:    flex.ExpandStringPointer(m.CloudCredentialId),
		CredentialConfig:     ExpandCredentialConfig(ctx, m.CredentialConfig, diags),
		RestrictedToAccounts: flex.ExpandFrameworkListString(ctx, m.RestrictedToAccounts, diags),
	}
	return to
}

func FlattenSourceConfig(ctx context.Context, from *clouddiscovery.SourceConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SourceConfigAttrTypes)
	}
	m := SourceConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SourceConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SourceConfigModel) Flatten(ctx context.Context, from *clouddiscovery.SourceConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SourceConfigModel{}
	}
	m.AccountScheduleId = flex.FlattenStringPointer(from.AccountScheduleId)
	m.Accounts = flex.FlattenFrameworkListNestedBlock(ctx, from.Accounts, AccountAttrTypes, diags, FlattenAccount)
	m.CloudCredentialId = flex.FlattenStringPointer(from.CloudCredentialId)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CredentialConfig = FlattenCredentialConfig(ctx, from.CredentialConfig, diags)
	m.DeletedAt = timetypes.NewRFC3339TimePointerValue(from.DeletedAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.RestrictedToAccounts = flex.FlattenFrameworkListString(ctx, from.RestrictedToAccounts, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
