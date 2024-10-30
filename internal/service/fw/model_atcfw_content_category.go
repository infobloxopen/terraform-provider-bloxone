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

type AtcfwContentCategoryModel struct {
	CategoryCode    types.Int32  `tfsdk:"category_code"`
	CategoryName    types.String `tfsdk:"category_name"`
	FunctionalGroup types.String `tfsdk:"functional_group"`
}

var AtcfwContentCategoryAttrTypes = map[string]attr.Type{
	"category_code":    types.Int32Type,
	"category_name":    types.StringType,
	"functional_group": types.StringType,
}

var AtcfwContentCategoryResourceSchemaAttributes = map[string]schema.Attribute{
	"category_code": schema.Int32Attribute{
		Optional:            true,
		MarkdownDescription: "The category code.",
	},
	"category_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the category.",
	},
	"functional_group": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The functional group name of the category.",
	},
}

func ExpandAtcfwContentCategory(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.ContentCategory {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwContentCategoryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwContentCategoryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.ContentCategory {
	if m == nil {
		return nil
	}
	to := &fw.ContentCategory{
		CategoryCode:    flex.ExpandInt32Pointer(m.CategoryCode),
		CategoryName:    flex.ExpandStringPointer(m.CategoryName),
		FunctionalGroup: flex.ExpandStringPointer(m.FunctionalGroup),
	}
	return to
}

func FlattenAtcfwContentCategory(ctx context.Context, from *fw.ContentCategory, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwContentCategoryAttrTypes)
	}
	m := AtcfwContentCategoryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwContentCategoryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwContentCategoryModel) Flatten(ctx context.Context, from *fw.ContentCategory, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwContentCategoryModel{}
	}
	m.CategoryCode = flex.FlattenInt32Pointer(from.CategoryCode)
	m.CategoryName = flex.FlattenStringPointer(from.CategoryName)
	m.FunctionalGroup = flex.FlattenStringPointer(from.FunctionalGroup)
}
