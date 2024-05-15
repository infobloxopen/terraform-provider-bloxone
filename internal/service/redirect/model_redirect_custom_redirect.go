package redirect

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

	"github.com/infobloxopen/bloxone-go-client/redirect"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type RedirectCustomRedirectModel struct {
	CreatedTime timetypes.RFC3339 `tfsdk:"created_time"`
	Data        types.String      `tfsdk:"data"`
	Id          types.Int64       `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	PolicyIds   types.List        `tfsdk:"policy_ids"`
	PolicyNames types.List        `tfsdk:"policy_names"`
	UpdatedTime timetypes.RFC3339 `tfsdk:"updated_time"`
}

var RedirectCustomRedirectAttrTypes = map[string]attr.Type{
	"created_time": timetypes.RFC3339Type{},
	"data":         types.StringType,
	"id":           types.Int64Type,
	"name":         types.StringType,
	"policy_ids":   types.ListType{ElemType: types.Int64Type},
	"policy_names": types.ListType{ElemType: types.StringType},
	"updated_time": timetypes.RFC3339Type{},
}

var RedirectCustomRedirectResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Custom Redirect object was created.",
	},
	"data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The list of csv custom IPv4/IPv6 or a single domain redirect address.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Custom Redirect object identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The name of the custom redirect.",
	},
	"policy_ids": schema.ListAttribute{
		ElementType:         types.Int64Type,
		Computed:            true,
		MarkdownDescription: "The list of the security policy identifiers with which the named list is associated.",
	},
	"policy_names": schema.ListAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The list of the security policy names with which the custom redirect is associated.",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Custom Redirect object was last updated.",
	},
}

func ExpandRedirectCustomRedirect(ctx context.Context, o types.Object, diags *diag.Diagnostics) *redirect.CustomRedirect {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RedirectCustomRedirectModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *RedirectCustomRedirectModel) Expand(ctx context.Context, diags *diag.Diagnostics) *redirect.CustomRedirect {
	if m == nil {
		return nil
	}
	to := &redirect.CustomRedirect{
		Data: flex.ExpandStringPointer(m.Data),
		Name: flex.ExpandStringPointer(m.Name),
	}
	return to
}

func FlattenRedirectCustomRedirect(ctx context.Context, from *redirect.CustomRedirect, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RedirectCustomRedirectAttrTypes)
	}
	m := RedirectCustomRedirectModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RedirectCustomRedirectAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *RedirectCustomRedirectModel) Flatten(ctx context.Context, from *redirect.CustomRedirect, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = RedirectCustomRedirectModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Data = flex.FlattenStringPointer(from.Data)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.PolicyNames = flex.FlattenFrameworkListString(ctx, from.PolicyNames, diags)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
	m.PolicyIds = flex.FlattenFrameworkListInt32(ctx, from.PolicyIds, diags)
}
