package infra_provision

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/infra_provision"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type HostactivationJoinTokenModel struct {
	DeletedAt   timetypes.RFC3339 `tfsdk:"deleted_at"`
	Description types.String      `tfsdk:"description"`
	ExpiresAt   timetypes.RFC3339 `tfsdk:"expires_at"`
	Id          types.String      `tfsdk:"id"`
	LastUsedAt  timetypes.RFC3339 `tfsdk:"last_used_at"`
	Name        types.String      `tfsdk:"name"`
	Status      types.String      `tfsdk:"status"`
	Tags        types.Map         `tfsdk:"tags"`
	TokenId     types.String      `tfsdk:"token_id"`
	UseCounter  types.Int64       `tfsdk:"use_counter"`
}

var HostactivationJoinTokenAttrTypes = map[string]attr.Type{
	"deleted_at":   timetypes.RFC3339Type{},
	"description":  types.StringType,
	"expires_at":   timetypes.RFC3339Type{},
	"id":           types.StringType,
	"last_used_at": timetypes.RFC3339Type{},
	"name":         types.StringType,
	"status":       types.StringType,
	"tags":         types.MapType{ElemType: types.StringType},
	"token_id":     types.StringType,
	"use_counter":  types.Int64Type,
}

var HostactivationJoinTokenResourceSchemaAttributes = map[string]schema.Attribute{
	"deleted_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"description": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"expires_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Optional:   true,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"last_used_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"name": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
	},
	"status": schema.StringAttribute{
		Computed: true,
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"token_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `first half of the token.`,
	},
	"use_counter": schema.Int64Attribute{
		Computed: true,
	},
}

func ExpandHostactivationJoinToken(ctx context.Context, o types.Object, diags *diag.Diagnostics) *infra_provision.HostactivationJoinToken {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m HostactivationJoinTokenModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *HostactivationJoinTokenModel) Expand(ctx context.Context, diags *diag.Diagnostics) *infra_provision.HostactivationJoinToken {
	if m == nil {
		return nil
	}
	to := &infra_provision.HostactivationJoinToken{
		Description: m.Description.ValueStringPointer(),
		Name:        m.Name.ValueStringPointer(),
		Tags:        flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		ExpiresAt:   flex.ExpandTimePointer(ctx, m.ExpiresAt, diags),
	}
	return to
}

func FlattenHostactivationJoinToken(ctx context.Context, from *infra_provision.HostactivationJoinToken, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostactivationJoinTokenAttrTypes)
	}
	m := HostactivationJoinTokenModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, HostactivationJoinTokenAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HostactivationJoinTokenModel) Flatten(ctx context.Context, from *infra_provision.HostactivationJoinToken, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HostactivationJoinTokenModel{}
	}
	m.DeletedAt = timetypes.NewRFC3339TimePointerValue(from.DeletedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.ExpiresAt = timetypes.NewRFC3339TimePointerValue(from.ExpiresAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.LastUsedAt = timetypes.NewRFC3339TimePointerValue(from.LastUsedAt)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TokenId = flex.FlattenStringPointer(from.TokenId)
	m.UseCounter = flex.FlattenInt64(int64(*from.UseCounter))
	m.Status = flattenStatus(from.Status)
}

func flattenStatus(from *infra_provision.JoinTokenJoinTokenStatus) types.String {
	if from == nil {
		return types.StringNull()
	}
	return flex.FlattenString(string(*from))
}
