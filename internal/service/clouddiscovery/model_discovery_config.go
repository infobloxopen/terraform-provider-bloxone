package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

type DiscoveryConfigModel struct {
	AccountPreference       types.String      `tfsdk:"account_preference"`
	AdditionalConfig        types.Object      `tfsdk:"additional_config"`
	CreatedAt               timetypes.RFC3339 `tfsdk:"created_at"`
	CredentialPreference    types.Object      `tfsdk:"credential_preference"`
	DeletedAt               timetypes.RFC3339 `tfsdk:"deleted_at"`
	Description             types.String      `tfsdk:"description"`
	DesiredState            types.String      `tfsdk:"desired_state"`
	DestinationTypesEnabled types.List        `tfsdk:"destination_types_enabled"`
	Destinations            types.List        `tfsdk:"destinations"`
	Id                      types.String      `tfsdk:"id"`
	LastSync                timetypes.RFC3339 `tfsdk:"last_sync"`
	Name                    types.String      `tfsdk:"name"`
	ProviderType            types.String      `tfsdk:"provider_type"`
	SourceConfigs           types.List        `tfsdk:"source_configs"`
	Status                  types.String      `tfsdk:"status"`
	StatusMessage           types.String      `tfsdk:"status_message"`
	SyncInterval            types.String      `tfsdk:"sync_interval"`
	Tags                    types.Map         `tfsdk:"tags"`
	TagsAll                 types.Map         `tfsdk:"tags_all"`
	UpdatedAt               timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DiscoveryConfigAttrTypes = map[string]attr.Type{
	"account_preference":        types.StringType,
	"additional_config":         types.ObjectType{AttrTypes: AdditionalConfigAttrTypes},
	"created_at":                timetypes.RFC3339Type{},
	"credential_preference":     types.ObjectType{AttrTypes: CredentialPreferenceAttrTypes},
	"deleted_at":                timetypes.RFC3339Type{},
	"description":               types.StringType,
	"desired_state":             types.StringType,
	"destination_types_enabled": types.ListType{ElemType: types.StringType},
	"destinations":              types.ListType{ElemType: types.ObjectType{AttrTypes: DestinationAttrTypes}},
	"id":                        types.StringType,
	"last_sync":                 timetypes.RFC3339Type{},
	"name":                      types.StringType,
	"provider_type":             types.StringType,
	"source_configs":            types.ListType{ElemType: types.ObjectType{AttrTypes: SourceConfigAttrTypes}},
	"status":                    types.StringType,
	"status_message":            types.StringType,
	"sync_interval":             types.StringType,
	"tags":                      types.MapType{ElemType: types.StringType},
	"tags_all":                  types.MapType{ElemType: types.StringType},
	"updated_at":                timetypes.RFC3339Type{},
}

var DiscoveryConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"account_preference": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.OneOf("single", "multiple", "auto_discover_multiple"),
		},
		MarkdownDescription: "Account preference. For ex.: single, multiple, auto_discover_multiple.",
	},
	"additional_config": schema.SingleNestedAttribute{
		Attributes:          AdditionalConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Additional configuration. Ex.: '{    \"excluded_object_types\": [],    \"exclusion_account_list\": [],    \"zone_forwarding\": \"true\" or \"false\" }'.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been created.",
	},
	"credential_preference": schema.SingleNestedAttribute{
		Attributes:          CredentialPreferenceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Credential preference. Ex.: '{    \"type\": \"dynamic\",    \"access_identifier_type\": \"role_arn\" or \"tenant_id\" or \"project_id\"  }'.",
	},
	"deleted_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Optional:            true,
		MarkdownDescription: "Timestamp when the object has been deleted.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Description of the discovery config. Optional.",
	},
	"desired_state": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("enabled"),
		MarkdownDescription: "Desired state. Default is \"enabled\".",
	},
	"destination_types_enabled": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Destinations types enabled: Ex.: DNS, IPAM and ACCOUNT.",
	},
	"destinations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DestinationResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		Validators: []validator.List{
			internalvalidator.DestinationTypeDependency(),
		},
		//Default: listdefault.StaticValue(types.ListValueMust(types.ObjectType{
		//	AttrTypes: DestinationAttrTypes,
		//}, []attr.Value{
		//	types.ObjectValueMust(DestinationAttrTypes, map[string]attr.Value{
		//		"config":           types.ObjectNull(DestinationConfigAttrTypes),
		//		"destination_type": types.StringValue("IPAM/DHCP"),
		//		"id":               types.StringNull(),
		//		"created_at":       timetypes.NewRFC3339Null(),
		//		"deleted_at":       timetypes.NewRFC3339Null(),
		//		"updated_at":       timetypes.NewRFC3339Null(),
		//	}),
		//})),
		MarkdownDescription: "Destinations For the discovery config.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto-generated unique discovery config ID. Format BloxID.",
	},
	"last_sync": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Last sync timestamp.",
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Name of the discovery config.",
	},
	"provider_type": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.OneOf("Amazon Web Services", "Google Cloud Platform", "Microsoft Azure"),
		},
		MarkdownDescription: "Provider type. Ex.: Amazon Web Services, Google Cloud Platform, Microsoft Azure.",
	},
	"source_configs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SourceConfigResourceSchemaAttributes,
		},
		Computed: true,
		Optional: true,
		PlanModifiers: []planmodifier.List{
			listplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "Source configs.",
	},
	"status": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Status of the sync operation. In single account case, Its the combined status of account & all the destinations statuses In auto discover case, Its the status of the account discovery only.",
	},
	"status_message": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Aggregate status message of the sync operation.",
	},
	"sync_interval": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("Auto"),
		MarkdownDescription: "Sync interval. Default is \"Auto\".",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Tagging specifics includes default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Timestamp when the object has been updated.",
	},
}

func (m *DiscoveryConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *clouddiscovery.DiscoveryConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.DiscoveryConfig{
		Name:                 flex.ExpandString(m.Name),
		CredentialPreference: ExpandCredentialPreference(ctx, m.CredentialPreference, diags),
		AccountPreference:    flex.ExpandStringPointer(m.AccountPreference),
		ProviderType:         flex.ExpandStringPointer(m.ProviderType),
		AdditionalConfig:     ExpandAdditionalConfig(ctx, m.AdditionalConfig, diags),
		Description:          flex.ExpandStringPointer(m.Description),
		DesiredState:         flex.ExpandStringPointer(m.DesiredState),
		Destinations:         flex.ExpandFrameworkListNestedBlock(ctx, m.Destinations, diags, ExpandDestination),
		SyncInterval:         flex.ExpandStringPointer(m.SyncInterval),
		Tags:                 flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.SourceConfigs = flex.ExpandFrameworkListNestedBlock(ctx, m.SourceConfigs, diags, ExpandSourceConfig)

	}
	return to
}

func FlattenDiscoveryConfig(ctx context.Context, from *clouddiscovery.DiscoveryConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DiscoveryConfigAttrTypes)
	}
	m := DiscoveryConfigModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, DiscoveryConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DiscoveryConfigModel) Flatten(ctx context.Context, from *clouddiscovery.DiscoveryConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DiscoveryConfigModel{}
	}
	m.AccountPreference = flex.FlattenStringPointer(from.AccountPreference)
	m.AdditionalConfig = FlattenAdditionalConfig(ctx, from.AdditionalConfig, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CredentialPreference = FlattenCredentialPreference(ctx, from.CredentialPreference, diags)
	m.DeletedAt = timetypes.NewRFC3339TimePointerValue(from.DeletedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DesiredState = flex.FlattenStringPointer(from.DesiredState)
	m.DestinationTypesEnabled = flex.FlattenFrameworkListString(ctx, from.DestinationTypesEnabled, diags)
	m.Destinations = flex.FlattenFrameworkListNestedBlock(ctx, from.Destinations, DestinationAttrTypes, diags, FlattenDestination)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.LastSync = timetypes.NewRFC3339TimePointerValue(from.LastSync)
	m.Name = flex.FlattenString(from.Name)
	m.ProviderType = flex.FlattenStringPointer(from.ProviderType)
	m.SourceConfigs = flex.FlattenFrameworkListNestedBlock(ctx, from.SourceConfigs, SourceConfigAttrTypes, diags, FlattenSourceConfig)
	m.Status = flex.FlattenStringPointer(from.Status)
	m.StatusMessage = flex.FlattenStringPointer(from.StatusMessage)
	m.SyncInterval = flex.FlattenStringPointer(from.SyncInterval)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
