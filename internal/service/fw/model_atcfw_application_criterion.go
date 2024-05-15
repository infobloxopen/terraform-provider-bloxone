package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwApplicationCriterionModel struct {
	Category    types.String `tfsdk:"category"`
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Subcategory types.String `tfsdk:"subcategory"`
}

var AtcfwApplicationCriterionAttrTypes = map[string]attr.Type{
	"category":    types.StringType,
	"id":          types.StringType,
	"name":        types.StringType,
	"subcategory": types.StringType,
}

var AtcfwApplicationCriterionResourceSchemaAttributes = map[string]schema.Attribute{
	"category": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
	},
	"id": schema.StringAttribute{
		Computed: true,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name for the application. Since the name of application is unique it may be used as alternate key for the application. The 'name' is used for import-export workflow and should be resolved to the 'id' before continue processing Create/Update operations.",
	},
	"subcategory": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
	},
}

func ExpandAtcfwApplicationCriterion(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.ApplicationCriterion {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwApplicationCriterionModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwApplicationCriterionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.ApplicationCriterion {
	if m == nil {
		return nil
	}
	to := &fw.ApplicationCriterion{
		Category:    flex.ExpandStringPointer(m.Category),
		Name:        flex.ExpandStringPointer(m.Name),
		Subcategory: flex.ExpandStringPointer(m.Subcategory),
	}
	return to
}

func FlattenAtcfwApplicationCriterion(ctx context.Context, from *fw.ApplicationCriterion, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwApplicationCriterionAttrTypes)
	}
	m := AtcfwApplicationCriterionModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwApplicationCriterionAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwApplicationCriterionModel) Flatten(ctx context.Context, from *fw.ApplicationCriterion, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwApplicationCriterionModel{}
	}
	m.Category = flex.FlattenStringPointer(from.Category)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Subcategory = flex.FlattenStringPointer(from.Subcategory)
}
