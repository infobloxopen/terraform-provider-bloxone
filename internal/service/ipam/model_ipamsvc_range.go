package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcRangeModel struct {
	Comment                  types.String      `tfsdk:"comment"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpHost                 types.String      `tfsdk:"dhcp_host"`
	DhcpOptions              types.List        `tfsdk:"dhcp_options"`
	DisableDhcp              types.Bool        `tfsdk:"disable_dhcp"`
	End                      types.String      `tfsdk:"end"`
	ExclusionRanges          types.List        `tfsdk:"exclusion_ranges"`
	Filters                  types.List        `tfsdk:"filters"`
	Id                       types.String      `tfsdk:"id"`
	InheritanceAssignedHosts types.List        `tfsdk:"inheritance_assigned_hosts"`
	InheritanceParent        types.String      `tfsdk:"inheritance_parent"`
	InheritanceSources       types.Object      `tfsdk:"inheritance_sources"`
	Name                     types.String      `tfsdk:"name"`
	Parent                   types.String      `tfsdk:"parent"`
	Protocol                 types.String      `tfsdk:"protocol"`
	Space                    types.String      `tfsdk:"space"`
	Start                    types.String      `tfsdk:"start"`
	Tags                     types.Map         `tfsdk:"tags"`
	Threshold                types.Object      `tfsdk:"threshold"`
	UpdatedAt                timetypes.RFC3339 `tfsdk:"updated_at"`
	Utilization              types.Object      `tfsdk:"utilization"`
	UtilizationV6            types.Object      `tfsdk:"utilization_v6"`
}

var IpamsvcRangeAttrTypes = map[string]attr.Type{
	"comment":                    types.StringType,
	"created_at":                 timetypes.RFC3339Type{},
	"dhcp_host":                  types.StringType,
	"dhcp_options":               types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"disable_dhcp":               types.BoolType,
	"end":                        types.StringType,
	"exclusion_ranges":           types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcExclusionRangeAttrTypes}},
	"filters":                    types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcAccessFilterAttrTypes}},
	"id":                         types.StringType,
	"inheritance_assigned_hosts": types.ListType{ElemType: types.ObjectType{AttrTypes: InheritanceAssignedHostAttrTypes}},
	"inheritance_parent":         types.StringType,
	"inheritance_sources":        types.ObjectType{AttrTypes: IpamsvcDHCPOptionsInheritanceAttrTypes},
	"name":                       types.StringType,
	"parent":                     types.StringType,
	"protocol":                   types.StringType,
	"space":                      types.StringType,
	"start":                      types.StringType,
	"tags":                       types.MapType{ElemType: types.StringType},
	"threshold":                  types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                 timetypes.RFC3339Type{},
	"utilization":                types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":             types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
}

var IpamsvcRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the range. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The resource identifier.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of DHCP options. May be either a specific option or a group of options.",
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The end IP address of the range.",
	},
	"exclusion_ranges": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcExclusionRangeResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of all exclusion ranges in the scope of the range.",
	},
	"filters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcAccessFilterResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The list of all allow/deny filters of the range.",
	},
	"id": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"inheritance_assigned_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritanceAssignedHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of the inheritance assigned hosts of the object.",
	},
	"inheritance_parent": schema.StringAttribute{
		MarkdownDescription: "The resource identifier.",
		Computed:            true,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: IpamsvcDHCPOptionsInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the range. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of protocol (_ip4_ or _ip6_).",
	},
	"space": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The start IP address of the range.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the range in JSON format.",
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationThresholdResourceSchemaAttributes,
		Computed:   true,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"utilization": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationResourceSchemaAttributes,
		Computed:   true,
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes: IpamsvcUtilizationV6ResourceSchemaAttributes,
		Computed:   true,
	},
}

func ExpandIpamsvcRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.Range {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags, true)
}

func (m *IpamsvcRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *ipam.Range {
	if m == nil {
		return nil
	}
	to := &ipam.Range{
		End:                flex.ExpandString(m.End),
		Start:              flex.ExpandString(m.Start),
		Comment:            flex.ExpandStringPointer(m.Comment),
		DhcpHost:           flex.ExpandStringPointer(m.DhcpHost),
		DhcpOptions:        flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandIpamsvcOptionItem),
		DisableDhcp:        flex.ExpandBoolPointer(m.DisableDhcp),
		ExclusionRanges:    flex.ExpandFrameworkListNestedBlock(ctx, m.ExclusionRanges, diags, ExpandIpamsvcExclusionRange),
		Filters:            flex.ExpandFrameworkListNestedBlock(ctx, m.Filters, diags, ExpandIpamsvcAccessFilter),
		InheritanceParent:  flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources: ExpandIpamsvcDHCPOptionsInheritance(ctx, m.InheritanceSources, diags),
		Name:               flex.ExpandStringPointer(m.Name),
		Parent:             flex.ExpandStringPointer(m.Parent),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		Threshold:          ExpandIpamsvcUtilizationThreshold(ctx, m.Threshold, diags),
		Utilization:        ExpandIpamsvcUtilization(ctx, m.Utilization, diags),
		UtilizationV6:      ExpandIpamsvcUtilizationV6(ctx, m.UtilizationV6, diags),
	}
	if isCreate {
		to.Space = flex.ExpandStringPointer(m.Space)
	}
	return to
}

func FlattenIpamsvcRange(ctx context.Context, from *ipam.Range, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcRangeAttrTypes)
	}
	m := IpamsvcRangeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcRangeModel) Flatten(ctx context.Context, from *ipam.Range, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcRangeModel{}
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpHost = flex.FlattenStringPointerWithNilAsEmpty(from.DhcpHost)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, FlattenIpamsvcOptionItem)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.End = flex.FlattenString(from.End)
	m.ExclusionRanges = flex.FlattenFrameworkListNestedBlock(ctx, from.ExclusionRanges, IpamsvcExclusionRangeAttrTypes, diags, FlattenIpamsvcExclusionRange)
	m.Filters = flex.FlattenFrameworkListNestedBlock(ctx, from.Filters, IpamsvcAccessFilterAttrTypes, diags, FlattenIpamsvcAccessFilter)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceAssignedHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, InheritanceAssignedHostAttrTypes, diags, FlattenInheritanceAssignedHost)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenIpamsvcDHCPOptionsInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.Start = flex.FlattenString(from.Start)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = FlattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = FlattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = FlattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)
}
