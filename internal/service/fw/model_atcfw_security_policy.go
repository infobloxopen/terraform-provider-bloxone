package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwSecurityPolicyModel struct {
	AccessCodes         types.List        `tfsdk:"access_codes"`
	CreatedTime         timetypes.RFC3339 `tfsdk:"created_time"`
	DefaultAction       types.String      `tfsdk:"default_action"`
	DefaultRedirectName types.String      `tfsdk:"default_redirect_name"`
	Description         types.String      `tfsdk:"description"`
	DfpServices         types.List        `tfsdk:"dfp_services"`
	Dfps                types.List        `tfsdk:"dfps"`
	Ecs                 types.Bool        `tfsdk:"ecs"`
	Id                  types.Int64       `tfsdk:"id"`
	IsDefault           types.Bool        `tfsdk:"is_default"`
	Name                types.String      `tfsdk:"name"`
	NetAddressDfps      types.List        `tfsdk:"net_address_dfps"`
	NetworkLists        types.List        `tfsdk:"network_lists"`
	OnpremResolve       types.Bool        `tfsdk:"onprem_resolve"`
	Precedence          types.Int64       `tfsdk:"precedence"`
	RoamingDeviceGroups types.List        `tfsdk:"roaming_device_groups"`
	Rules               types.List        `tfsdk:"rules"`
	SafeSearch          types.Bool        `tfsdk:"safe_search"`
	Tags                types.Map         `tfsdk:"tags"`
	UpdatedTime         timetypes.RFC3339 `tfsdk:"updated_time"`
	UserGroups          types.List        `tfsdk:"user_groups"`
}

var AtcfwSecurityPolicyAttrTypes = map[string]attr.Type{
	"access_codes":          types.ListType{ElemType: types.StringType},
	"created_time":          timetypes.RFC3339Type{},
	"default_action":        types.StringType,
	"default_redirect_name": types.StringType,
	"description":           types.StringType,
	"dfp_services":          types.ListType{ElemType: types.StringType},
	"dfps":                  types.ListType{ElemType: types.Int64Type},
	"ecs":                   types.BoolType,
	"id":                    types.Int64Type,
	"is_default":            types.BoolType,
	"name":                  types.StringType,
	"net_address_dfps":      types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwNetAddrDfpAssignmentAttrTypes}},
	"network_lists":         types.ListType{ElemType: types.Int64Type},
	"onprem_resolve":        types.BoolType,
	"precedence":            types.Int64Type,
	"roaming_device_groups": types.ListType{ElemType: types.Int64Type},
	"rules":                 types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwSecurityPolicyRuleAttrTypes}},
	"safe_search":           types.BoolType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"updated_time":          timetypes.RFC3339Type{},
	"user_groups":           types.ListType{ElemType: types.StringType},
}

var AtcfwSecurityPolicyResourceSchemaAttributes = map[string]schema.Attribute{
	"access_codes": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Access codes assigned to Security Policy",
	},
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Security Policy object was created.",
	},
	"default_action": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString("action_allow"),
		MarkdownDescription: "The policy-level action gets applied when none of the policy rules apply/match. The default value for default_action is \"action_allow\".",
	},
	"default_redirect_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "Name of the custom redirect, if the default_action is \"action_redirect\".",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for the security policy.",
	},
	"dfp_services": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of DNS Forwarding Proxy Services object identifiers. For Internal Use only.",
	},
	"dfps": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		MarkdownDescription: "The list of DNS Forwarding Proxy object identifiers.",
	},
	"ecs": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use ECS for handling policy",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Security Policy object identifier.",
	},
	"is_default": schema.BoolAttribute{
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Flag that indicates whether this is a default security policy.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the security policy.",
	},
	"net_address_dfps": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwNetAddrDfpAssignmentResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of DFPs associated with this policy via network address (with corresponding network address)",
	},
	"network_lists": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		MarkdownDescription: "The list of Network Lists identifiers that represents networks that you want to protect from malicious attacks.",
	},
	"onprem_resolve": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use DNS resolve on onprem side",
	},
	"precedence": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Security precedence enable selection of the highest priority policy, in cases where a query matches multiple policies.",
	},
	"roaming_device_groups": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		MarkdownDescription: "The list of BloxOne Endpoint groups identifiers.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwSecurityPolicyRuleResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of Security Policy Rules objects that represent the set of rules and actions that you define to balance access and constraints so you can mitigate malicious attacks and provide security for your networks.",
	},
	"safe_search": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Apply automated rules to enforce safe search",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Security Policy object was last updated.",
	},
	"user_groups": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "List of user groups associated with this policy",
	},
}

func ExpandAtcfwSecurityPolicy(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwSecurityPolicy {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwSecurityPolicyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwSecurityPolicyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwSecurityPolicy {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwSecurityPolicy{
		AccessCodes:         flex.ExpandFrameworkListString(ctx, m.AccessCodes, diags),
		DefaultAction:       flex.ExpandStringPointer(m.DefaultAction),
		DefaultRedirectName: flex.ExpandStringPointer(m.DefaultRedirectName),
		Description:         flex.ExpandStringPointer(m.Description),
		DfpServices:         flex.ExpandFrameworkListString(ctx, m.DfpServices, diags),
		Ecs:                 flex.ExpandBoolPointer(m.Ecs),
		Name:                flex.ExpandStringPointer(m.Name),
		NetAddressDfps:      flex.ExpandFrameworkListNestedBlock(ctx, m.NetAddressDfps, diags, ExpandAtcfwNetAddrDfpAssignment),
		OnpremResolve:       flex.ExpandBoolPointer(m.OnpremResolve),
		Precedence:          flex.ExpandInt32Pointer(m.Precedence),
		Rules:               flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandAtcfwSecurityPolicyRule),
		SafeSearch:          flex.ExpandBoolPointer(m.SafeSearch),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		UserGroups:          flex.ExpandFrameworkListString(ctx, m.UserGroups, diags),
		Dfps:                flex.ExpandFrameworkListInt32(ctx, m.Dfps, diags),
		NetworkLists:        flex.ExpandFrameworkListInt64(ctx, m.NetworkLists, diags),
		RoamingDeviceGroups: flex.ExpandFrameworkListInt32(ctx, m.RoamingDeviceGroups, diags),
	}
	return to
}

func FlattenAtcfwSecurityPolicy(ctx context.Context, from *fw.AtcfwSecurityPolicy, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwSecurityPolicyAttrTypes)
	}
	m := AtcfwSecurityPolicyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwSecurityPolicyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwSecurityPolicyModel) Flatten(ctx context.Context, from *fw.AtcfwSecurityPolicy, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwSecurityPolicyModel{}
	}
	m.AccessCodes = flex.FlattenFrameworkListString(ctx, from.AccessCodes, diags)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.DefaultAction = flex.FlattenStringPointer(from.DefaultAction)
	m.DefaultRedirectName = flex.FlattenStringPointer(from.DefaultRedirectName)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DfpServices = flex.FlattenFrameworkListString(ctx, from.DfpServices, diags)
	m.Ecs = types.BoolPointerValue(from.Ecs)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.NetAddressDfps = flex.FlattenFrameworkListNestedBlock(ctx, from.NetAddressDfps, AtcfwNetAddrDfpAssignmentAttrTypes, diags, FlattenAtcfwNetAddrDfpAssignment)
	m.OnpremResolve = types.BoolPointerValue(from.OnpremResolve)
	m.Precedence = flex.FlattenInt32Pointer(from.Precedence)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, AtcfwSecurityPolicyRuleAttrTypes, diags, FlattenAtcfwSecurityPolicyRule)
	m.SafeSearch = types.BoolPointerValue(from.SafeSearch)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
	m.UserGroups = flex.FlattenFrameworkListString(ctx, from.UserGroups, diags)
	m.Dfps = flex.FlattenFrameworkListInt32(ctx, from.Dfps, diags)
	m.NetworkLists = flex.FlattenFrameworkListInt64(ctx, from.NetworkLists, diags)
	m.RoamingDeviceGroups = flex.FlattenFrameworkListInt32(ctx, from.RoamingDeviceGroups, diags)
}
