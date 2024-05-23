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

type AtcfwApplicationFilterModel struct {
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Criteria    types.List        `tfsdk:"criteria"`
	Description types.String      `tfsdk:"description"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Policies    types.List        `tfsdk:"policies"`
	Readonly    types.Bool        `tfsdk:"readonly"`
	Tags        types.Map         `tfsdk:"tags"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwApplicationFilterAttrTypes = map[string]attr.Type{
	"created_time": timetypes.RFC3339Type{},
	"criteria":     types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwApplicationCriterionAttrTypes}},
	"description":  types.StringType,
	"id":           types.Int64Type,
	"name":         types.StringType,
	"policies":     types.ListType{ElemType: types.StringType},
	"readonly":     types.BoolType,
	"tags":         types.MapType{ElemType: types.StringType},
	"updated_time": timetypes.RFC3339Type{},
}

var AtcfwApplicationFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Application Filter object was created.",
	},
	"criteria": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwApplicationCriterionResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "The array of key-value pairs specifying criteria for the search.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for the application filter.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Application Filter object identifier.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the application filter.",
	},
	"policies": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of security policy names with which the application filter is associated.",
	},
	"readonly": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if it is a predefined application filter",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Application Filter object was last updated.",
	},
}

func ExpandAtcfwApplicationFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.ApplicationFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwApplicationFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwApplicationFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.ApplicationFilter {
	if m == nil {
		return nil
	}
	to := &fw.ApplicationFilter{
		Criteria:    flex.ExpandFrameworkListNestedBlock(ctx, m.Criteria, diags, ExpandAtcfwApplicationCriterion),
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenAtcfwApplicationFilter(ctx context.Context, from *fw.ApplicationFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwApplicationFilterAttrTypes)
	}
	m := AtcfwApplicationFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwApplicationFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwApplicationFilterModel) Flatten(ctx context.Context, from *fw.ApplicationFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwApplicationFilterModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Criteria = flex.FlattenFrameworkListNestedBlock(ctx, from.Criteria, AtcfwApplicationCriterionAttrTypes, diags, FlattenAtcfwApplicationCriterion)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	m.Readonly = types.BoolPointerValue(from.Readonly)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
