package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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
	Protocol          types.String      `tfsdk:"protocol"`
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
	"protocol":           types.StringType,
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
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
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
		Optional:            true,
		MarkdownDescription: "The metadata for the federated pool in JSON format.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the federated pool. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliance": schema.SingleNestedAttribute{
		Attributes:          NetworkComplianceResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The network compliance of the __FederatedPool__.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates if this pool is compliant with its parent's network compliance policy. When false, a trouble dot should be displayed in the UI to indicate non-compliance.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The address family of the pool ('ip4', 'ip6', or 'ip4/ip6' for dual mode support on NIOS_X pools only).",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	},
	"region": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The region/locale this pool is associated with (e.g., 'us-west-1', 'eu-central-1').",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The current state of the federated pool (e.g., 'create-complete', 'create-in-progress', 'delete-in-progress').",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
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

func ExpandFederatedPool(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.FederatedPool {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FederatedPoolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, true)
}

func (m *FederatedPoolModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipamfederation.FederatedPool {
	if m == nil {
		return nil
	}
	to := &ipamfederation.FederatedPool{
		Description:       flex.ExpandStringPointer(m.Description),
		Metadata:          flex.ExpandFrameworkMapString(ctx, m.Metadata, diags),
		Name:              flex.ExpandStringPointer(m.Name),
		NetworkCompliance: ExpandNetworkCompliance(ctx, m.NetworkCompliance, diags),
		Parent:            flex.ExpandStringPointer(m.Parent),
		Tags:              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.FederatedRealm = flex.ExpandStringPointer(m.FederatedRealm)
		to.Protocol = flex.ExpandStringPointer(m.Protocol)
		providerType := ipamfederation.PROVIDERTYPE_NIOS_X
		to.Provider = &providerType
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
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCompliance = FlattenNetworkCompliance(ctx, from.NetworkCompliance, diags)
	m.NetworkCompliant = types.BoolPointerValue(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Region = flex.FlattenStringPointer(from.Region)
	m.State = flex.FlattenStringPointer(from.State)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = flex.FlattenInt64Pointer(from.Utilization)
	m.UtilizationV6 = FlattenUtilizationV6(ctx, from.UtilizationV6, diags)
}
