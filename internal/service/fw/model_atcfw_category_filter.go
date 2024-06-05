package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwCategoryFilterModel struct {
	Categories  types.List        `tfsdk:"categories"`
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Description types.String      `tfsdk:"description"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Policies    types.List        `tfsdk:"policies"`
	Tags        types.Map         `tfsdk:"tags"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwCategoryFilterAttrTypes = map[string]attr.Type{
	"categories":   types.ListType{ElemType: types.StringType},
	"created_time": timetypes.RFC3339Type{},
	"description":  types.StringType,
	"id":           types.Int64Type,
	"name":         types.StringType,
	"policies":     types.ListType{ElemType: types.StringType},
	"tags":         types.MapType{ElemType: types.StringType},
	"updated_time": timetypes.RFC3339Type{},
}

var AtcfwCategoryFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"categories": schema.ListAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: "The list of content category names that falls into this category filter.",
	},
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Category Filter object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for the category filter.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Category Filter object identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the category filter.",
	},
	"policies": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of security policy names with which the category filter is associated.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Category Filter object was last updated.",
	},
}

func ExpandAtcfwCategoryFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.CategoryFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwCategoryFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwCategoryFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.CategoryFilter {
	if m == nil {
		return nil
	}
	to := &fw.CategoryFilter{
		Categories:  flex.ExpandFrameworkListString(ctx, m.Categories, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenAtcfwCategoryFilter(ctx context.Context, from *fw.CategoryFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwCategoryFilterAttrTypes)
	}
	m := AtcfwCategoryFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwCategoryFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwCategoryFilterModel) Flatten(ctx context.Context, from *fw.CategoryFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwCategoryFilterModel{}
	}
	m.Categories = flex.FlattenFrameworkListString(ctx, from.Categories, diags)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
