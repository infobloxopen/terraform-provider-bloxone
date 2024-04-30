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

type AtcfwItemStructsModel struct {
	Description types.String `tfsdk:"description"`
	Item        types.String `tfsdk:"item"`
}

var AtcfwItemStructsAttrTypes = map[string]attr.Type{
	"description": types.StringType,
	"item":        types.StringType,
}

var AtcfwItemStructsResourceSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the item",
	},
	"item": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The data of the Item",
	},
}

func ExpandAtcfwItemStructs(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.ItemStructs {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwItemStructsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwItemStructsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.ItemStructs {
	if m == nil {
		return nil
	}
	to := &fw.ItemStructs{
		Description: flex.ExpandStringPointer(m.Description),
		Item:        flex.ExpandStringPointer(m.Item),
	}
	return to
}

func FlattenAtcfwItemStructs(ctx context.Context, from *fw.ItemStructs, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwItemStructsAttrTypes)
	}
	m := AtcfwItemStructsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwItemStructsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwItemStructsModel) Flatten(ctx context.Context, from *fw.ItemStructs, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwItemStructsModel{}
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Item = flex.FlattenStringPointer(from.Item)
}
