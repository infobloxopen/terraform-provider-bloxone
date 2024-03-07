package fw

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/fw"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcfwInternalDomainsModel struct {
	CreatedTime     timetypes.RFC3339 `tfsdk:"created_time"`
	Description     types.String      `tfsdk:"description"`
	Id              types.Int64       `tfsdk:"id"`
	InternalDomains types.List        `tfsdk:"internal_domains"`
	IsDefault       types.Bool        `tfsdk:"is_default"`
	Name            types.String      `tfsdk:"name"`
	Tags            types.Map         `tfsdk:"tags"`
	UpdatedTime     timetypes.RFC3339 `tfsdk:"updated_time"`
}

var AtcfwInternalDomainsAttrTypes = map[string]attr.Type{
	"created_time":     timetypes.RFC3339Type{},
	"description":      types.StringType,
	"id":               types.Int64Type,
	"internal_domains": types.ListType{ElemType: types.StringType},
	"is_default":       types.BoolType,
	"name":             types.StringType,
	"tags":             types.MapType{ElemType: types.StringType},
	"updated_time":     timetypes.RFC3339Type{},
}

var AtcfwInternalDomainsResourceSchemaAttributes = map[string]schema.Attribute{
	"created_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Internal Domain lists object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		Default:             stringdefault.StaticString(""),
		Computed:            true,
		MarkdownDescription: "The brief description for the internal domain lists .",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Internal Domain object identifier.",
	},
	"internal_domains": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of internal domains, should be unique to each other and has to be read-only from the API level.",
	},
	"is_default": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if name is 'Default Bypass Domains/CIDRs' otherwise false.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The name of the internal domain lists.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"updated_time": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The time when this Internal domain lists object was last updated.",
	},
}

func ExpandAtcfwInternalDomains(ctx context.Context, o types.Object, diags *diag.Diagnostics) *fw.AtcfwInternalDomains {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcfwInternalDomainsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcfwInternalDomainsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *fw.AtcfwInternalDomains {
	if m == nil {
		return nil
	}
	to := &fw.AtcfwInternalDomains{
		Description:     flex.ExpandStringPointer(m.Description),
		InternalDomains: flex.ExpandFrameworkListString(ctx, m.InternalDomains, diags),
		IsDefault:       flex.ExpandBoolPointer(m.IsDefault),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}

	return to
}

func FlattenAtcfwInternalDomains(ctx context.Context, from *fw.AtcfwInternalDomains, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcfwInternalDomainsAttrTypes)
	}
	m := AtcfwInternalDomainsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcfwInternalDomainsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcfwInternalDomainsModel) Flatten(ctx context.Context, from *fw.AtcfwInternalDomains, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcfwInternalDomainsModel{}
	}
	m.CreatedTime = timetypes.NewRFC3339TimePointerValue(from.CreatedTime)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt32Pointer(from.Id)
	m.InternalDomains = flex.FlattenFrameworkListString(ctx, from.InternalDomains, diags)
	m.IsDefault = types.BoolPointerValue(from.IsDefault)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedTime = timetypes.NewRFC3339TimePointerValue(from.UpdatedTime)
}
