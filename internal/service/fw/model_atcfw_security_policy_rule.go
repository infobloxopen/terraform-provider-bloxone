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

type AtcfwSecurityPolicyRuleModel struct {
	Action       types.String `tfsdk:"action"`
	Data         types.String `tfsdk:"data"`
	ListId       types.Int64  `tfsdk:"list_id"`
	PolicyId     types.Int64  `tfsdk:"policy_id"`
	PolicyName   types.String `tfsdk:"policy_name"`
	RedirectName types.String `tfsdk:"redirect_name"`
	Type         types.String `tfsdk:"type"`
}

var AtcfwSecurityPolicyRuleAttrTypes = map[string]attr.Type{
	"action":        types.StringType,
	"data":          types.StringType,
	"list_id":       types.Int64Type,
	"policy_id":     types.Int64Type,
	"policy_name":   types.StringType,
	"redirect_name": types.StringType,
	"type":          types.StringType,
}

var AtcfwSecurityPolicyRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The action for the policy rule that can be either \"action_allow\" or \"action_log\" or \"action_redirect\" or \"action_block\" or \"action_allow_with_local_resolution\". \"action_allow_with_local_resolution\" only supported for application filter rule with enabled onprem_resolve flag on the Security policy.",
	},
	"data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The data source for the policy rule, that can be either a name of the predefined feed for \"named_feed\", custom list name for \"custom_list\" type, category filter name for \"category_filter\" type and application filter name for \"application_filter\" type.",
	},
	"list_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The Custom List object identifier with which the policy rule is associated. 0 value means no custom list is associated with this policy rule.",
	},
	"policy_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The identifier of the Security Policy object with which the policy rule is associated.",
	},
	"policy_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the security policy with which the policy rule is associated.",
	},
	"redirect_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the redirect address for redirect actions that can be either IPv4 address or a domain name.",
	},
	"type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The policy rule type that can be either \"named_feed\" or \"custom_list\" or \"category_filter\" or \"application_filter\".",
	},
}

func ExpandAtcfwSecurityPolicyRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwSecurityPolicyRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwSecurityPolicyRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwSecurityPolicyRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwSecurityPolicyRule {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwSecurityPolicyRule{
		Action:       flex.ExpandStringPointer(m.Action),
		Data:         flex.ExpandStringPointer(m.Data),
		PolicyName:   flex.ExpandStringPointer(m.PolicyName),
		RedirectName: flex.ExpandStringPointer(m.RedirectName),
		Type:         flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenAtcfwSecurityPolicyRule(ctx context.Context, from *fw.AtcfwSecurityPolicyRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwSecurityPolicyRuleAttrTypes)
	}
	m := AtcfwSecurityPolicyRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwSecurityPolicyRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwSecurityPolicyRuleModel) Flatten(ctx context.Context, from *fw.AtcfwSecurityPolicyRule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwSecurityPolicyRuleModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.Data = flex.FlattenStringPointer(from.Data)
	m.ListId = flex.FlattenInt32Pointer(from.ListId)
	m.PolicyId = flex.FlattenInt32Pointer(from.PolicyId)
	m.PolicyName = flex.FlattenStringPointer(from.PolicyName)
	m.RedirectName = flex.FlattenStringPointer(from.RedirectName)
	m.Type = flex.FlattenStringPointer(from.Type)
}
