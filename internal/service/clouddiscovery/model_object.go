package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ObjectModel struct {
	Category    types.Object `tfsdk:"category"`
	ResourceSet types.List   `tfsdk:"resource_set"`
}

var ObjectAttrTypes = map[string]attr.Type{
	"category":     types.ObjectType{AttrTypes: CategoryAttrTypes},
	"resource_set": types.ListType{ElemType: types.ObjectType{AttrTypes: ResourceAttrTypes}},
}

var ObjectResourceSchemaAttributes = map[string]schema.Attribute{
	"category": schema.SingleNestedAttribute{
		Attributes:          CategoryResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Category of the object.",
	},
	"resource_set": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ResourceResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Resource set of the object .",
	},
}

func ExpandObject(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.Object {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ObjectModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ObjectModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.Object {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.Object{
		Category:    ExpandCategory(ctx, m.Category, diags),
		ResourceSet: flex.ExpandFrameworkListNestedBlock(ctx, m.ResourceSet, diags, ExpandResource),
	}
	return to
}

func FlattenObject(ctx context.Context, from *clouddiscovery.Object, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ObjectAttrTypes)
	}
	m := ObjectModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ObjectAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ObjectModel) Flatten(ctx context.Context, from *clouddiscovery.Object, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ObjectModel{}
	}
	m.Category = FlattenCategory(ctx, from.Category, diags)
	m.ResourceSet = flex.FlattenFrameworkListNestedBlock(ctx, from.ResourceSet, ResourceAttrTypes, diags, FlattenResource)
}
