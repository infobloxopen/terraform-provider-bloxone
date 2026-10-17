package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type EtlFilterRecordTypeModel struct {
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	Data                types.List        `tfsdk:"data"`
	Description         types.String      `tfsdk:"description"`
	Enabled             types.Bool        `tfsdk:"enabled"`
	FilterSource        types.String      `tfsdk:"filter_source"`
	Id                  types.Int64       `tfsdk:"id"`
	IngressPolicyAction types.String      `tfsdk:"ingress_policy_action"`
	Name                types.String      `tfsdk:"name"`
	Tags                types.Map         `tfsdk:"tags"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
}

var EtlFilterRecordTypeAttrTypes = map[string]attr.Type{
	"created_at":            timetypes.RFC3339Type{},
	"data":                  types.ListType{ElemType: types.StringType},
	"description":           types.StringType,
	"enabled":               types.BoolType,
	"filter_source":         types.StringType,
	"id":                    types.Int64Type,
	"ingress_policy_action": types.StringType,
	"name":                  types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"updated_at":            timetypes.RFC3339Type{},
}

var EtlFilterRecordTypeResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"data": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"filter_source": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"ingress_policy_action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
}

func ExpandEtlFilterRecordType(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.EtlFilterRecordType {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m EtlFilterRecordTypeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *EtlFilterRecordTypeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.EtlFilterRecordType {
	if m == nil {
		return nil
	}
	to := &cdc.EtlFilterRecordType{
		Data:                flex.ExpandFrameworkListString(ctx, m.Data, diags),
		Description:         flex.ExpandStringPointer(m.Description),
		Enabled:             flex.ExpandBoolPointer(m.Enabled),
		IngressPolicyAction: flex.ExpandStringPointer(m.IngressPolicyAction),
		Name:                flex.ExpandStringPointer(m.Name),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		FilterSource:        (*cdc.FilterSource)(flex.ExpandStringPointer(m.FilterSource)),
	}
	return to
}

func FlattenEtlFilterRecordType(ctx context.Context, from *cdc.EtlFilterRecordType, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(EtlFilterRecordTypeAttrTypes)
	}
	m := EtlFilterRecordTypeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, EtlFilterRecordTypeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *EtlFilterRecordTypeModel) Flatten(ctx context.Context, from *cdc.EtlFilterRecordType, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = EtlFilterRecordTypeModel{}
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Data = flex.FlattenFrameworkListString(ctx, from.Data, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IngressPolicyAction = flex.FlattenStringPointer(from.IngressPolicyAction)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.FilterSource = flex.FlattenStringPointer((*string)(from.FilterSource))
}
