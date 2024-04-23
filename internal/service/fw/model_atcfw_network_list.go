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

type AtcfwNetworkListModel struct {
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Description types.String      `tfsdk:"description"`
	Id          types.Int64       `tfsdk:"id"`
	Items       types.List        `tfsdk:"items"`
	Name        types.String      `tfsdk:"name"`
	PolicyId    types.Int64       `tfsdk:"policy_id"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwNetworkListAttrTypes = map[string]attr.Type{
	"created_time": timetypes.RFC3339Type{},
	"description":  types.StringType,
	"id":           types.Int64Type,
	"items":        types.ListType{ElemType: types.StringType},
	"name":         types.StringType,
	"policy_id":    types.Int64Type,
	"updated_time": timetypes.RFC3339Type{},
}

var AtcfwNetworkListResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Network List object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for the network list.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Network List object identifier.",
	},
	"items": schema.ListAttribute{
		ElementType:         types.StringType,
		Required:            true,
		MarkdownDescription: "The list of networks' CIDRs that are subject for malicious attacks protection.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the network list.",
	},
	"policy_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The identifier of the security policy with which the network list is associated.",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Network List object was last updated.",
	},
}

func ExpandAtcfwNetworkList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwNetworkList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwNetworkListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwNetworkListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwNetworkList {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwNetworkList{
		Description: flex.ExpandStringPointer(m.Description),
		Items:       flex.ExpandFrameworkListString(ctx, m.Items, diags),
		Name:        flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenAtcfwNetworkList(ctx context.Context, from *fw.AtcfwNetworkList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwNetworkListAttrTypes)
	}
	m := AtcfwNetworkListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwNetworkListAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwNetworkListModel) Flatten(ctx context.Context, from *fw.AtcfwNetworkList, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwNetworkListModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Items = flex.FlattenFrameworkListString(ctx, from.Items, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
