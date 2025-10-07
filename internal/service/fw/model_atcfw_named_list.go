package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

type AtcfwNamedListModel struct {
	ConfidenceLevel types.String      `tfsdk:"confidence_level"`
	CreatedTime     timetypes.RFC3339 `tfsdk:"created_time"`
	Description     types.String      `tfsdk:"description"`
	Id              types.Int32       `tfsdk:"id"`
	ItemCount       types.Int32       `tfsdk:"item_count"`
	Items           types.List        `tfsdk:"items"`
	ItemsDescribed  types.List        `tfsdk:"items_described"`
	Name            types.String      `tfsdk:"name"`
	Policies        types.List        `tfsdk:"policies"`
	Tags            types.Map         `tfsdk:"tags"`
	TagsAll         types.Map         `tfsdk:"tags_all"`
	ThreatLevel     types.String      `tfsdk:"threat_level"`
	Type            types.String      `tfsdk:"type"`
	UpdatedTime     timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwNamedListAttrTypes = map[string]attr.Type{
	"confidence_level": types.StringType,
	"created_time":     timetypes.RFC3339Type{},
	"description":      types.StringType,
	"id":               types.Int32Type,
	"item_count":       types.Int32Type,
	"items":            types.ListType{ElemType: types.StringType},
	"items_described":  types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwItemStructsAttrTypes}},
	"name":             types.StringType,
	"policies":         types.ListType{ElemType: types.StringType},
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
	"threat_level":     types.StringType,
	"type":             types.StringType,
	"updated_time":     timetypes.RFC3339Type{},
}

var AtcfwNamedListResourceSchemaAttributes = map[string]schema.Attribute{
	"confidence_level": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The confidence level for a custom list. The possible values are [\"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Named List object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for the named list.",
	},
	"id": schema.Int32Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int32{
			int32planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Named List object identifier.",
	},
	"item_count": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The number of items in this named list.",
	},
	"items": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of the FQDN or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection.",
	},
	"items_described": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwItemStructsResourceSchemaAttributes,
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The List of ItemStructs structure which contains the item and its description",
	},
	"name": schema.StringAttribute{
		Required:            true,
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
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the named list.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the named list, including default tags.",
	},
	"threat_level": schema.StringAttribute{
		Computed:            true,
		Optional:            true,
		MarkdownDescription: "The threat level for a custom list. The possible values are [\"INFO\", \"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"type": schema.StringAttribute{
		Computed: true,
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The type of the named list, that can be \"custom_list\", \"threat_insight\", \"fast_flux\", \"dga\", \"dnsm\", \"threat_insight_nde\", \"default_allow\", \"default_block\" or \"threat_insight_nde\".",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Named List object was last updated.",
	},
}

func (m *AtcfwNamedListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.NamedList {
	if m == nil {
		return nil
	}

	to := &fw.NamedList{
		ConfidenceLevel: flex.ExpandStringPointer(m.ConfidenceLevel),
		Description:     flex.ExpandStringPointer(m.Description),
		Items:           flex.ExpandFrameworkListString(ctx, m.Items, diags),
		ItemsDescribed:  flex.ExpandFrameworkListNestedBlock(ctx, m.ItemsDescribed, diags, ExpandAtcfwItemStructs), Name: flex.ExpandStringPointer(m.Name),
		Policies:    flex.ExpandFrameworkListString(ctx, m.Policies, diags),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		ThreatLevel: flex.ExpandStringPointer(m.ThreatLevel),
		Type:        flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenAtcfwNamedList(ctx context.Context, from *fw.NamedListRead, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwNamedListAttrTypes)
	}
	m := AtcfwNamedListModel{}
	m.FlattenRead(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, AtcfwNamedListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwNamedListModel) Flatten(ctx context.Context, from *fw.NamedList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwNamedListModel{}
	}
	m.ConfidenceLevel = flex.FlattenStringPointer(from.ConfidenceLevel)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.ItemCount = flex.FlattenInt32Pointer(from.ItemCount)
	m.Items = flex.FlattenFrameworkListString(ctx, from.Items, diags)
	planList := m.ItemsDescribed
	m.ItemsDescribed = flex.FlattenFrameworkListNestedBlock(ctx, from.ItemsDescribed, AtcfwItemStructsAttrTypes, diags, FlattenAtcfwItemStructs)
	if !planList.IsUnknown() {
		reOrderedList, diags := utils.ReorderAndFilterNestedListResponse(ctx, planList, m.ItemsDescribed, "item")
		if !diags.HasError() {
			m.ItemsDescribed = reOrderedList.(basetypes.ListValue)
		}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.ThreatLevel = flex.FlattenStringPointer(from.ThreatLevel)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}

func (m *AtcfwNamedListModel) FlattenRead(ctx context.Context, from *fw.NamedListRead, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwNamedListModel{}
	}
	m.ConfidenceLevel = flex.FlattenStringPointer(from.ConfidenceLevel)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.ItemCount = flex.FlattenInt32Pointer(from.ItemCount)
	m.Items = types.ListNull(types.StringType)
	m.ItemsDescribed = types.ListNull(types.ObjectType{AttrTypes: AtcfwItemStructsAttrTypes})
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.ThreatLevel = flex.FlattenStringPointer(from.ThreatLevel)
	m.Type = flex.FlattenStringPointer(from.Type)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
