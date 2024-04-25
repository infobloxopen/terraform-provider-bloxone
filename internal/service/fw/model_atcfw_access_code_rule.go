package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwAccessCodeRuleModel struct {
	Action       types.String `tfsdk:"action"`
	Data         types.String `tfsdk:"data"`
	Description  types.String `tfsdk:"description"`
	RedirectName types.String `tfsdk:"redirect_name"`
	Type         types.String `tfsdk:"type"`
}

var AtcfwAccessCodeRuleAttrTypes = map[string]attr.Type{
	"action":        types.StringType,
	"data":          types.StringType,
	"description":   types.StringType,
	"redirect_name": types.StringType,
	"type":          types.StringType,
}

var AtcfwAccessCodeRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The action to be used in the rule.",
	},
	"data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The data to be used in the rule.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description of the rule.",
	},
	"redirect_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the redirect to be used in the rule.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of the rule.",
	},
}

func ExpandAtcfwAccessCodeRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AccessCodeRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwAccessCodeRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwAccessCodeRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AccessCodeRule {
	if m == nil {
		return nil
	}
	to := &fw.AccessCodeRule{
		Action:       flex.ExpandStringPointer(m.Action),
		Data:         flex.ExpandStringPointer(m.Data),
		Description:  flex.ExpandStringPointer(m.Description),
		RedirectName: flex.ExpandStringPointer(m.RedirectName),
		Type:         flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenAtcfwAccessCodeRule(ctx context.Context, from *fw.AccessCodeRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwAccessCodeRuleAttrTypes)
	}
	m := AtcfwAccessCodeRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwAccessCodeRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwAccessCodeRuleModel) Flatten(ctx context.Context, from *fw.AccessCodeRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwAccessCodeRuleModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Data = flex.FlattenStringPointer(from.Data)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.RedirectName = flex.FlattenStringPointer(from.RedirectName)
	m.Type = flex.FlattenStringPointer(from.Type)
}
