package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type FederatedBlockModel struct {
	Address        types.String      `tfsdk:"address"`
	AllocationV4   types.Object      `tfsdk:"allocation_v4"`
	Cidr           types.Int64       `tfsdk:"cidr"`
	Comment        types.String      `tfsdk:"comment"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at"`
	FederatedRealm types.String      `tfsdk:"federated_realm"`
	Id             types.String      `tfsdk:"id"`
	Name           types.String      `tfsdk:"name"`
	Parent         types.String      `tfsdk:"parent"`
	Protocol       types.String      `tfsdk:"protocol"`
	Tags           types.Map         `tfsdk:"tags"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at"`
}

var FederatedBlockAttrTypes = map[string]attr.Type{
	"address":         types.StringType,
	"allocation_v4":   types.ObjectType{AttrTypes: AllocationAttrTypes},
	"cidr":            types.Int64Type,
	"comment":         types.StringType,
	"created_at":      timetypes.RFC3339Type{},
	"federated_realm": types.StringType,
	"id":              types.StringType,
	"name":            types.StringType,
	"parent":          types.StringType,
	"protocol":        types.StringType,
	"tags":            types.MapType{ElemType: types.StringType},
	"updated_at":      timetypes.RFC3339Type{},
}

var FederatedBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
	},
	"allocation_v4": schema.SingleNestedAttribute{
		Attributes:          AllocationResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The percentage of the Federated Block’s total address space that is consumed by Leaf Terminals.",
	},
	"cidr": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The CIDR of the federated block. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"federated_realm": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the federated block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol of federated block (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the federated block in JSON format.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandFederatedBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipamfederation.FederatedBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FederatedBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *FederatedBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.FederatedBlock {
	if m == nil {
		return nil
	}
	to := &ipamfederation.FederatedBlock{
		Address:        flex.ExpandStringPointer(m.Address),
		AllocationV4:   ExpandAllocation(ctx, m.AllocationV4, diags),
		Cidr:           flex.ExpandInt64Pointer(m.Cidr),
		Comment:        flex.ExpandStringPointer(m.Comment),
		FederatedRealm: flex.ExpandString(m.FederatedRealm),
		Name:           flex.ExpandStringPointer(m.Name),
		Parent:         flex.ExpandStringPointer(m.Parent),
		Tags:           flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenFederatedBlock(ctx context.Context, from *ipamfederation.FederatedBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FederatedBlockAttrTypes)
	}
	m := FederatedBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FederatedBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FederatedBlockModel) Flatten(ctx context.Context, from *ipamfederation.FederatedBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FederatedBlockModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AllocationV4 = FlattenAllocation(ctx, from.AllocationV4, diags)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.FederatedRealm = flex.FlattenString(from.FederatedRealm)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
