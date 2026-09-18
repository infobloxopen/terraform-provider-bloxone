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
	internaltypes "github.com/infobloxopen/terraform-provider-bloxone/internal/types"
)

type ForwardLookingDelegationModel struct {
	Address          types.String                     `tfsdk:"address"`
	Cidr             types.Int64                      `tfsdk:"cidr"`
	Comment          types.String                     `tfsdk:"comment"`
	CreatedAt        timetypes.RFC3339                `tfsdk:"created_at"`
	FederatedPoolId  types.String                     `tfsdk:"federated_pool_id"`
	FederatedRealms  internaltypes.UnorderedListValue `tfsdk:"federated_realms"`
	Id               types.String                     `tfsdk:"id"`
	Name             types.String                     `tfsdk:"name"`
	NetworkCompliant types.Bool                       `tfsdk:"network_compliant"`
	Protocol         types.String                     `tfsdk:"protocol"`
	Tags             types.Map                        `tfsdk:"tags"`
	TagsAll          types.Map                        `tfsdk:"tags_all"`
	UpdatedAt        timetypes.RFC3339                `tfsdk:"updated_at"`
}

var ForwardLookingDelegationAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"cidr":              types.Int64Type,
	"comment":           types.StringType,
	"created_at":        timetypes.RFC3339Type{},
	"federated_pool_id": types.StringType,
	"federated_realms":  internaltypes.UnorderedListOfStringType,
	"id":                types.StringType,
	"name":              types.StringType,
	"network_compliant": types.BoolType,
	"protocol":          types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
	"updated_at":        timetypes.RFC3339Type{},
}

var ForwardLookingDelegationResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The address field in form \"a.b.c.d/n\" where the \"/n\" may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form \"a.b.c.d\".",
	},
	"cidr": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the delegation. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the delegation. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"federated_pool_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"federated_realms": schema.ListAttribute{
		ElementType:         types.StringType,
		CustomType:          internaltypes.UnorderedListOfStringType,
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the delegation. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The compliance status of the forward looking delegation, as determined by the federation service.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of delegation (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the delegation in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the forward looking delegation in JSON format including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandForwardLookingDelegation(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.ForwardLookingDelegation {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ForwardLookingDelegationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, false)
}

func (m *ForwardLookingDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipamfederation.ForwardLookingDelegation {
	if m == nil {
		return nil
	}
	to := &ipamfederation.ForwardLookingDelegation{
		Cidr:            flex.ExpandInt64Pointer(m.Cidr),
		Comment:         flex.ExpandStringPointer(m.Comment),
		FederatedPoolId: flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealms: flex.ExpandFrameworkListString(ctx, m.FederatedRealms, diags),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	if isCreate {
		to.Address = flex.ExpandStringPointer(m.Address)
	}
	return to
}

func FlattenForwardLookingDelegation(ctx context.Context, from *ipamfederation.ForwardLookingDelegation, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ForwardLookingDelegationAttrTypes)
	}
	m := ForwardLookingDelegationModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, ForwardLookingDelegationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ForwardLookingDelegationModel) Flatten(ctx context.Context, from *ipamfederation.ForwardLookingDelegation, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ForwardLookingDelegationModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealms = flex.FlattenFrameworkUnorderedList(ctx, types.StringType, from.FederatedRealms, diags)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCompliant = types.BoolPointerValue(from.NetworkCompliant)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
