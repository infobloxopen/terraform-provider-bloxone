package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type CategoryModel struct {
	Excluded types.Bool   `tfsdk:"excluded"`
	Id       types.String `tfsdk:"id"`
}

var CategoryAttrTypes = map[string]attr.Type{
	"excluded": types.BoolType,
	"id":       types.StringType,
}

var CategoryResourceSchemaAttributes = map[string]schema.Attribute{
	"excluded": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "If set true , the category is excluded from discovery.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		Optional: true,
		Validators: []validator.String{
			stringvalidator.OneOf("security", "networking-basics", "lbs", "compute", "azure-storage", "networking-advanced"),
		},
		MarkdownDescription: "Category ID.",
	},
}

func ExpandCategory(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.Category {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m CategoryModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *CategoryModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.Category {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.Category{
		Excluded: flex.ExpandBoolPointer(m.Excluded),
		Id:       flex.ExpandStringPointer(m.Id),
	}
	return to
}

func FlattenCategory(ctx context.Context, from *clouddiscovery.Category, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(CategoryAttrTypes)
	}
	m := CategoryModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, CategoryAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *CategoryModel) Flatten(ctx context.Context, from *clouddiscovery.Category, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = CategoryModel{}
	}
	m.Excluded = types.BoolPointerValue(from.Excluded)
	m.Id = flex.FlattenStringPointer(from.Id)
}
