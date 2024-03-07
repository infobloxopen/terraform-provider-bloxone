package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwInternalDomainsItemsModel struct {
	DeletedItemsDescribed  types.List  `tfsdk:"deleted_items_described"`
	Id                     types.Int64 `tfsdk:"id"`
	InsertedItemsDescribed types.List  `tfsdk:"inserted_items_described"`
}

var AtcfwInternalDomainsItemsAttrTypes = map[string]attr.Type{
	"deleted_items_described":  types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwItemStructsAttrTypes}},
	"id":                       types.Int64Type,
	"inserted_items_described": types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwItemStructsAttrTypes}},
}

var AtcfwInternalDomainsItemsResourceSchemaAttributes = map[string]schema.Attribute{
	"deleted_items_described": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwItemStructsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The List of ItemStructs structure which contains the item and its description",
	},
	"id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Internal Domain List object identifier.",
	},
	"inserted_items_described": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwItemStructsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The List of ItemStructs structure which contains the item and its description",
	},
}

func ExpandAtcfwInternalDomainsItems(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwInternalDomainsItems {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwInternalDomainsItemsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwInternalDomainsItemsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwInternalDomainsItems {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwInternalDomainsItems{
		DeletedItemsDescribed:  flex.ExpandFrameworkListNestedBlock(ctx, m.DeletedItemsDescribed, diags, ExpandAtcfwItemStructs),
		InsertedItemsDescribed: flex.ExpandFrameworkListNestedBlock(ctx, m.InsertedItemsDescribed, diags, ExpandAtcfwItemStructs),
	}
	return to
}

func FlattenAtcfwInternalDomainsItems(ctx context.Context, from *fw.AtcfwInternalDomainsItems, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwInternalDomainsItemsAttrTypes)
	}
	m := AtcfwInternalDomainsItemsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwInternalDomainsItemsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwInternalDomainsItemsModel) Flatten(ctx context.Context, from *fw.AtcfwInternalDomainsItems, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwInternalDomainsItemsModel{}
	}
	m.DeletedItemsDescribed = flex.FlattenFrameworkListNestedBlock(ctx, from.DeletedItemsDescribed, AtcfwItemStructsAttrTypes, diags, FlattenAtcfwItemStructs)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.InsertedItemsDescribed = flex.FlattenFrameworkListNestedBlock(ctx, from.InsertedItemsDescribed, AtcfwItemStructsAttrTypes, diags, FlattenAtcfwItemStructs)
}
