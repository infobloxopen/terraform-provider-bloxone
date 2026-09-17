package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type OverlappingBlockModel struct {
	Address          types.String      `tfsdk:"address"`
	Cidr             types.Int64       `tfsdk:"cidr"`
	Comment          types.String      `tfsdk:"comment"`
	CreatedAt        timetypes.RFC3339 `tfsdk:"created_at"`
	FederatedPoolId  types.String      `tfsdk:"federated_pool_id"`
	FederatedRealm   types.String      `tfsdk:"federated_realm"`
	Id               types.String      `tfsdk:"id"`
	Name             types.String      `tfsdk:"name"`
	NetworkCompliant types.Bool        `tfsdk:"network_compliant"`
	Parent           types.String      `tfsdk:"parent"`
	Protocol         types.String      `tfsdk:"protocol"`
	Tags             types.Map         `tfsdk:"tags"`
	UpdatedAt        timetypes.RFC3339 `tfsdk:"updated_at"`
}

var OverlappingBlockAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"cidr":              types.Int64Type,
	"comment":           types.StringType,
	"created_at":        timetypes.RFC3339Type{},
	"federated_pool_id": types.StringType,
	"federated_realm":   types.StringType,
	"id":                types.StringType,
	"name":              types.StringType,
	"network_compliant": types.BoolType,
	"parent":            types.StringType,
	"protocol":          types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"updated_at":        timetypes.RFC3339Type{},
}

var OverlappingBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplace(),
		},
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. This field is immutable after creation.",
	},
	"cidr": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 128),
		},
		MarkdownDescription: "The CIDR of the overlapping block. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the overlapping block. May contain 0 to 1024 characters. Can include UTF-8.",
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
	"federated_realm": schema.StringAttribute{
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
		MarkdownDescription: "The name of the overlapping block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"network_compliant": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "The compliance status of the overlapping block, as determined by the federation service.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of overlapping block (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the overlapping block in JSON format.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandOverlappingBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.OverlappingBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OverlappingBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, false)
}

func (m *OverlappingBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipamfederation.OverlappingBlock {
	if m == nil {
		return nil
	}
	to := &ipamfederation.OverlappingBlock{
		Cidr:            flex.ExpandInt64Pointer(m.Cidr),
		Comment:         flex.ExpandStringPointer(m.Comment),
		FederatedPoolId: flex.ExpandStringPointer(m.FederatedPoolId),
		FederatedRealm:  flex.ExpandString(m.FederatedRealm),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	// Address is immutable after creation - the API rejects it in update requests.
	if isCreate {
		to.Address = flex.ExpandStringPointer(m.Address)
	}
	return to
}

func FlattenOverlappingBlock(ctx context.Context, from *ipamfederation.OverlappingBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OverlappingBlockAttrTypes)
	}
	m := OverlappingBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OverlappingBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *OverlappingBlockModel) Flatten(ctx context.Context, from *ipamfederation.OverlappingBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = OverlappingBlockModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.FederatedPoolId = flex.FlattenStringPointer(from.FederatedPoolId)
	m.FederatedRealm = flex.FlattenString(from.FederatedRealm)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetworkCompliant = types.BoolPointerValue(from.NetworkCompliant)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
