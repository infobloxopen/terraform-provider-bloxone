package fw

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwAccessCodeModel struct {
	AccessKey   types.String      `tfsdk:"access_key"`
	Activation  timetypes.RFC3339 `tfsdk:"activation"`
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Description types.String      `tfsdk:"description"`
	Expiration  timetypes.RFC3339 `tfsdk:"expiration"`
	Name        types.String      `tfsdk:"name"`
	PolicyIds   types.List        `tfsdk:"policy_ids"`
	Rules       types.List        `tfsdk:"rules"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
	Id          types.String      `tfsdk:"id"`
}

var AtcfwAccessCodeAttrTypes = map[string]attr.Type{
	"access_key":   types.StringType,
	"activation":   timetypes.RFC3339Type{},
	"created_time": timetypes.RFC3339Type{},
	"description":  types.StringType,
	"expiration":   timetypes.RFC3339Type{},
	"name":         types.StringType,
	"policy_ids":   types.ListType{ElemType: types.Int64Type},
	"rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: AtcfwAccessCodeRuleAttrTypes}},
	"updated_time": timetypes.RFC3339Type{},
	"id":           types.StringType,
}

var AtcfwAccessCodeResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"access_key": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Auto generated unique Bypass Code value",
	},
	"activation": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Required:            true,
		MarkdownDescription: "The time when the Bypass Code object was activated.",
	},
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when the Bypass Code object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The brief description for an access code.",
	},
	"expiration": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Required:            true,
		MarkdownDescription: "The time when the Bypass Code object expires.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of Bypass Code",
	},
	"policy_ids": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The list of SecurityPolicy object identifiers.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AtcfwAccessCodeRuleResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: "The list of selected security rules",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when the Bypass Code object was last updated.",
	},
}

func ExpandAtcfwAccessCode(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwAccessCode {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwAccessCodeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwAccessCodeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwAccessCode {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwAccessCode{
		AccessKey:   flex.ExpandStringPointer(m.AccessKey),
		Activation:  flex.ExpandTimePointer(ctx, m.Activation, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Expiration:  flex.ExpandTimePointer(ctx, m.Expiration, diags),
		Name:        flex.ExpandStringPointer(m.Name),
		Rules:       flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandAtcfwAccessCodeRule),
		PolicyIds:   flex.ExpandFrameworkListInt32(ctx, m.PolicyIds, diags),
	}
	return to
}

func FlattenAtcfwAccessCode(ctx context.Context, from *fw.AtcfwAccessCode, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwAccessCodeAttrTypes)
	}
	m := AtcfwAccessCodeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwAccessCodeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwAccessCodeModel) Flatten(ctx context.Context, from *fw.AtcfwAccessCode, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwAccessCodeModel{}
	}
	m.AccessKey = flex.FlattenStringPointer(from.AccessKey)
	m.Activation = timetypes.NewRFC3339TimePointerValue(from.Activation)
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Expiration = timetypes.NewRFC3339TimePointerValue(from.Expiration)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Id = flex.FlattenStringPointer(from.AccessKey)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, AtcfwAccessCodeRuleAttrTypes, diags, FlattenAtcfwAccessCodeRule)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
	m.PolicyIds = flex.FlattenFrameworkListInt32(ctx, from.PolicyIds, diags)
}
