package ipamfederation

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

type FederatedPoolModel struct {
	Allocation        types.Object      `tfsdk:"allocation"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	Description       types.String      `tfsdk:"description"`
	FederatedRealm    types.String      `tfsdk:"federated_realm"`
	Id                types.String      `tfsdk:"id"`
	Metadata          types.Map         `tfsdk:"metadata"`
	Name              types.String      `tfsdk:"name"`
	NetworkCompliance types.Object      `tfsdk:"network_compliance"`
	NetworkCompliant  types.Bool        `tfsdk:"network_compliant"`
	Parent            types.String      `tfsdk:"parent"`
	PoolName          types.String      `tfsdk:"pool_name"`
	Protocol          types.String      `tfsdk:"protocol"`
	ProviderType      types.String      `tfsdk:"provider_type"`
	Region            types.String      `tfsdk:"region"`
	State             types.String      `tfsdk:"state"`
	Tags              types.Map         `tfsdk:"tags"`
	TagsAll           types.Map         `tfsdk:"tags_all"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
	Utilization       types.Int64       `tfsdk:"utilization"`
	UtilizationV6     types.Object      `tfsdk:"utilization_v6"`
}

var FederatedPoolAttrTypes = map[string]attr.Type{
	"allocation":         types.ObjectType{AttrTypes: AllocationAttrTypes},
	"created_at":         timetypes.RFC3339Type{},
	"description":        types.StringType,
	"federated_realm":    types.StringType,
	"id":                 types.StringType,
	"metadata":           types.MapType{ElemType: types.StringType},
	"name":               types.StringType,
	"network_compliance": types.ObjectType{AttrTypes: NetworkComplianceAttrTypes},
	"network_compliant":  types.BoolType,
	"parent":             types.StringType,
	"pool_name":          types.StringType,
	"protocol":           types.StringType,
	"provider_type":      types.StringType,
	"region":             types.StringType,
	"state":              types.StringType,
	"tags":               types.MapType{ElemType: types.StringType},
	"tags_all":           types.MapType{ElemType: types.StringType},
	"updated_at":         timetypes.RFC3339Type{},
	"utilization":        types.Int64Type,
	"utilization_v6":     types.ObjectType{AttrTypes: UtilizationV6AttrTypes},
}

var FederatedPoolResourceSchemaAttributes = map[string]schema.Attribute{
	"allocation": schema.SingleNestedAttribute{
		Attributes:          AllocationResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The allocation details for the __FederatedPool__.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the federated pool. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"federated_realm": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The ID of the federated realm the pool belongs to. A pool cannot be moved to another realm, so changing it replaces the pool.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"metadata": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The metadata for the federated pool in JSON format.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the federated pool. May contain 1 to 256 characters. Can include UTF-8. For `AWS` pools the API appends a suffix to the stored name; it is not part of this value (see `pool_name`).",
	},
	"network_compliance": schema.SingleNestedAttribute{
		Attributes: NetworkComplianceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Validators: []validator.Object{
			internalvalidator.ObjectInt64Order("minimum_netmask_length", "default_netmask_length", "maximum_netmask_length"),
		},
		MarkdownDescription: "The network compliance of the __FederatedPool__.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if this pool is compliant with its parent's network compliance policy. When false, a trouble dot should be displayed in the UI to indicate non-compliance.",
	},
	"parent": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The ID of the parent federated pool. Leave it unset for a top-level pool. A pool cannot be moved to another parent, so changing it replaces the pool.",
	},
	"pool_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the federated pool exactly as stored by the API. For `AWS` pools the API appends a suffix such as `(ipam-pool-0123456789abcdef0)`; `name` is this value without that suffix. Use this value to filter pools by name in the data source.",
	},
	"protocol": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The address family of the pool ('ip4', 'ip6', or 'ip4/ip6' for dual mode support on NIOS_X pools only).",
		Validators: []validator.String{
			stringvalidator.OneOf("ip4", "ip6", "ip4/ip6"),
		},
	},
	"provider_type": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The cloud provider type of the pool (e.g. `NIOS_X`, `AWS`, `AZURE`). It is always the provider type of the federated realm and cannot be set.",
	},
	"region": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The region/locale of the pool. It is always the region of the federated realm and cannot be set.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The current state of the federated pool (e.g., 'create-complete', 'create-in-progress', 'delete-in-progress').",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the federated pool in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the federated pool in JSON format including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"utilization": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The IPv4 utilization percentage of the __FederatedPool__.",
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes:          UtilizationV6ResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The IPv6 utilization metrics for the __FederatedPool__.",
	},
}

// poolNameSuffix matches the "(ipam-pool-<id>)" suffix the API appends to the name of pools that are
// backed by a cloud provider, e.g. "my-pool(ipam-pool-0073c8bb836f02aec)".
var poolNameSuffix = regexp.MustCompile(`\(ipam-pool-[A-Za-z0-9]+\)$`)

func (m *FederatedPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipamfederation.FederatedPool {
	if m == nil {
		return nil
	}
	to := &ipamfederation.FederatedPool{
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if m.ProviderType.ValueString() == string(ipamfederation.PROVIDERTYPE_AWS) {
		// AWS pools do not support network_compliance (rejected instead of dropped silently, which would end in an
		// "inconsistent result after apply"), and only their name and description can be updated.
		if !isCreate {
			return &ipamfederation.FederatedPool{Name: to.Name, Description: to.Description}
		}
		if !m.NetworkCompliance.IsNull() && !m.NetworkCompliance.IsUnknown() {
			diags.AddAttributeError(path.Root("network_compliance"), "Unsupported attribute for provider type",
				"network_compliance is not supported for pools with provider_type \"AWS\".")
		}
	} else {
		to.NetworkCompliance = ExpandNetworkCompliance(ctx, m.NetworkCompliance, diags)
	}
	if isCreate {
		to.FederatedRealm = flex.ExpandStringPointer(m.FederatedRealm)
		to.Parent = flex.ExpandStringPointer(m.Parent)
		to.Protocol = flex.ExpandStringPointer(m.Protocol)
		if !m.ProviderType.IsNull() && !m.ProviderType.IsUnknown() {
			v := ipamfederation.ProviderType(m.ProviderType.ValueString())
			to.Provider = &v
		}
		to.Region = flex.ExpandStringPointer(m.Region)
	}
	return to
}

func FlattenFederatedPool(ctx context.Context, from *ipamfederation.FederatedPool, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FederatedPoolAttrTypes)
	}
	m := FederatedPoolModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, FederatedPoolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FederatedPoolModel) Flatten(ctx context.Context, from *ipamfederation.FederatedPool, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FederatedPoolModel{}
	}
	m.Allocation = FlattenAllocation(ctx, from.Allocation, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.FederatedRealm = flex.FlattenStringPointer(from.FederatedRealm)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Metadata = flex.FlattenFrameworkMapString(ctx, from.Metadata, diags)
	m.PoolName = flex.FlattenStringPointer(from.Name)
	m.Name = flattenPoolName(from.Name, from.GetProvider())
	m.NetworkCompliance = FlattenNetworkCompliance(ctx, from.NetworkCompliance, diags)
	m.NetworkCompliant = types.BoolPointerValue(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.ProviderType = types.StringValue(string(from.GetProvider()))
	m.Region = flex.FlattenStringPointer(from.Region)
	m.State = flex.FlattenStringPointer(from.State)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = flex.FlattenInt64Pointer(from.Utilization)
	m.UtilizationV6 = FlattenUtilizationV6(ctx, from.UtilizationV6, diags)
}

// stripPoolNameSuffix removes the provider-added suffix from a pool name returned by the API.
func stripPoolNameSuffix(name string) string {
	return poolNameSuffix.ReplaceAllString(name, "")
}

// flattenPoolName flattens the API pool name. The suffix is only removed for AWS pools, the only ones that get it,
// so a name chosen by the user for another pool is never altered.
func flattenPoolName(name *string, provider ipamfederation.ProviderType) types.String {
	if name == nil {
		return types.StringNull()
	}
	if provider != ipamfederation.PROVIDERTYPE_AWS {
		return types.StringValue(*name)
	}
	return flex.FlattenString(stripPoolNameSuffix(*name))
}
