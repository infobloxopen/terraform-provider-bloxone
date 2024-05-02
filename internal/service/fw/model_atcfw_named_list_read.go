package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// TODO: this model is redundant, and should be removed
type AtcfwNamedListReadModel struct {
	ConfidenceLevel types.String      `tfsdk:"confidence_level"`
	CreatedTime     timetypes.RFC3339 `tfsdk:"created_time"`
	Description     types.String      `tfsdk:"description"`
	Id              types.Int64       `tfsdk:"id"`
	ItemCount       types.Int64       `tfsdk:"item_count"`
	Name            types.String      `tfsdk:"name"`
	Policies        types.List        `tfsdk:"policies"`
	Tags            types.Map         `tfsdk:"tags"`
	TagsAll         types.Map         `tfsdk:"tags_all"`
	ThreatLevel     types.String      `tfsdk:"threat_level"`
	Type            types.String      `tfsdk:"type"`
	UpdatedTime     timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwNamedListReadAttrTypes = map[string]attr.Type{
	"confidence_level": types.StringType,
	"created_time":     timetypes.RFC3339Type{},
	"description":      types.StringType,
	"id":               types.Int64Type,
	"item_count":       types.Int64Type,
	"name":             types.StringType,
	"policies":         types.ListType{ElemType: types.StringType},
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
	"threat_level":     types.StringType,
	"type":             types.StringType,
	"updated_time":     timetypes.RFC3339Type{},
}

var AtcfwNamedListReadResourceSchemaAttributes = map[string]schema.Attribute{
	"confidence_level": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The confidence level for a custom list. The possible values are [\"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Named List object was created.",
	},
	"description": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The brief description for the named list.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Named List object identifier.",
	},
	"item_count": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The number of items in this named list.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the named list.",
	},
	"policies": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of the security policy names with which the named list is associated.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tags associated with this Named List",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Tags associated with this Named List, including default tags",
	},
	"threat_level": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The threat level for a custom list. The possible values are [\"INFO\", \"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of the named list, that can be \"custom_list\", \"threat_insight\", \"fast_flux\", \"dga\", \"dnsm\", \"threat_insight_nde\", \"default_allow\", \"default_block\".",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Named List object was last updated.",
	},
}

func ExpandAtcfwNamedListRead(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.NamedListRead {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwNamedListReadModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwNamedListReadModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.NamedListRead {
	if m == nil {
		return nil
	}
	to := &fw.NamedListRead{
		Tags: flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenAtcfwNamedListRead(ctx context.Context, from *fw.NamedListRead, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwNamedListReadAttrTypes)
	}
	m := AtcfwNamedListReadModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, AtcfwNamedListReadAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwNamedListReadModel) Flatten(ctx context.Context, from *fw.NamedListRead, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwNamedListReadModel{}
	}
	m.ConfidenceLevel = flex.FlattenStringPointer(from.ConfidenceLevel)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.ItemCount = flex.FlattenInt32Pointer(from.ItemCount)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.ThreatLevel = flex.FlattenStringPointer(from.ThreatLevel)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
