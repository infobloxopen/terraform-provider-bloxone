package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type EtlFilterHostnameModel struct {
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	Data                types.List        `tfsdk:"data"`
	Description         types.String      `tfsdk:"description"`
	Enabled             types.Bool        `tfsdk:"enabled"`
	FilterSource        types.String      `tfsdk:"filter_source"`
	Id                  types.Int64       `tfsdk:"id"`
	IngressPolicyAction types.String      `tfsdk:"ingress_policy_action"`
	Name                types.String      `tfsdk:"name"`
	Tags                types.Map         `tfsdk:"tags"`
	TagsAll             types.Map         `tfsdk:"tags_all"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
}

var EtlFilterHostnameAttrTypes = map[string]attr.Type{
	"created_at":            timetypes.RFC3339Type{},
	"data":                  types.ListType{ElemType: types.StringType},
	"description":           types.StringType,
	"enabled":               types.BoolType,
	"filter_source":         types.StringType,
	"id":                    types.Int64Type,
	"ingress_policy_action": types.StringType,
	"name":                  types.StringType,
	"tags":                  types.MapType{ElemType: types.StringType},
	"tags_all":              types.MapType{ElemType: types.StringType},
	"updated_at":            timetypes.RFC3339Type{},
}

var EtlFilterHostnameResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"data": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			listvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The list of hostnames to filter on.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description of the Hostname ETL filter.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates whether the Hostname ETL filter is enabled. Defaults to `false`.",
	},
	"filter_source": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(string(cdc.FILTERSOURCE_DEFAULT)),
		Validators: []validator.String{
			stringvalidator.OneOf(string(cdc.FILTERSOURCE_DEFAULT), string(cdc.FILTERSOURCE_INGRESS)),
		},
		MarkdownDescription: "The source the Hostname ETL filter applies to. Allowed values:\n" +
			"  * _SOURCE_DEFAULT_\n" +
			"  * _SOURCE_INGRESS_\n\n" +
			"Defaults to _SOURCE_DEFAULT_.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	// TODO: allowed values and server default for ingress_policy_action are not yet known.
	"ingress_policy_action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The ingress policy action of the Hostname ETL filter.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the Hostname ETL filter.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the Hostname ETL filter.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the Hostname ETL filter, including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandEtlFilterHostname(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.EtlFilterHostname {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m EtlFilterHostnameModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *EtlFilterHostnameModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.EtlFilterHostname {
	if m == nil {
		return nil
	}
	to := &cdc.EtlFilterHostname{
		Data:                flex.ExpandFrameworkListString(ctx, m.Data, diags),
		Description:         flex.ExpandStringPointer(m.Description),
		Enabled:             flex.ExpandBoolPointer(m.Enabled),
		FilterSource:        (*cdc.FilterSource)(flex.ExpandStringPointer(m.FilterSource)),
		IngressPolicyAction: flex.ExpandStringPointer(m.IngressPolicyAction),
		Name:                flex.ExpandStringPointer(m.Name),
		Tags:                flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenEtlFilterHostname(ctx context.Context, from *cdc.EtlFilterHostname, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(EtlFilterHostnameAttrTypes)
	}
	m := EtlFilterHostnameModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, EtlFilterHostnameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *EtlFilterHostnameModel) Flatten(ctx context.Context, from *cdc.EtlFilterHostname, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = EtlFilterHostnameModel{}
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Data = flex.FlattenFrameworkListString(ctx, from.Data, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.FilterSource = flex.FlattenStringPointer((*string)(from.FilterSource))
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IngressPolicyAction = flex.FlattenStringPointer(from.IngressPolicyAction)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
