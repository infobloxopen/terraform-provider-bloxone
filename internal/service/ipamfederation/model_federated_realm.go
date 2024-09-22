package ipamfederation

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/bloxone-go-client/ipamfederation"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type FederatedRealmModel struct {
	AllocationV4 types.Object      `tfsdk:"allocation_v4"`
	Comment      types.String      `tfsdk:"comment"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	Id           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	Tags         types.Map         `tfsdk:"tags"`
	TagsAll      types.Map         `tfsdk:"tags_all"`
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at"`
}

var FederatedRealmAttrTypes = map[string]attr.Type{
	"allocation_v4": types.ObjectType{AttrTypes: AllocationAttrTypes},
	"comment":       types.StringType,
	"created_at":    timetypes.RFC3339Type{},
	"id":            types.StringType,
	"name":          types.StringType,
	"tags":          types.MapType{ElemType: types.StringType},
	"tags_all":      types.MapType{ElemType: types.StringType},
	"updated_at":    timetypes.RFC3339Type{},
}

var FederatedRealmResourceSchemaAttributes = map[string]schema.Attribute{
	"allocation_v4": schema.SingleNestedAttribute{
		Attributes:          AllocationResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "The aggregate of all Federated Blocks within the Realm.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description of the federated realm. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the federated realm. May contain 1 to 256 characters; can include UTF-8.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: `The tags for the federation block in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags of the federation realm in JSON format including default tags.",
	},
}

func (m *FederatedRealmModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipamfederation.FederatedRealm {
	if m == nil {
		return nil
	}
	to := &ipamfederation.FederatedRealm{
		AllocationV4: ExpandAllocation(ctx, m.AllocationV4, diags),
		Comment:      flex.ExpandStringPointer(m.Comment),
		Name:         flex.ExpandString(m.Name),
		Tags:         flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenFederatedRealm(ctx context.Context, from *ipamfederation.FederatedRealm, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FederatedRealmAttrTypes)
	}
	m := FederatedRealmModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, FederatedRealmAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *FederatedRealmModel) Flatten(ctx context.Context, from *ipamfederation.FederatedRealm, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = FederatedRealmModel{}
	}
	m.AllocationV4 = FlattenAllocation(ctx, from.AllocationV4, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
