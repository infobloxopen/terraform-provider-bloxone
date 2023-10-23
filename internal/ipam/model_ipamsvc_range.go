package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
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
	"tags":                       types.MapType{},
	"threshold":                  types.ObjectType{AttrTypes: IpamsvcUtilizationThresholdAttrTypes},
	"updated_at":                 timetypes.RFC3339Type{},
	"utilization":                types.ObjectType{AttrTypes: IpamsvcUtilizationAttrTypes},
	"utilization_v6":             types.ObjectType{AttrTypes: IpamsvcUtilizationV6AttrTypes},
}

var IpamsvcRangeResourceSchema = schema.Schema{
	MarkdownDescription: `A __Range__ object (_ipam/range_) represents a set of contiguous IP addresses in the same IP space with no gap, expressed as a (start, end) pair within a given subnet that are grouped together for administrative purpose and protocol management. The start and end values are not required to align with CIDR boundaries. `,
	Attributes:          IpamsvcRangeResourceSchemaAttributes,
}

var IpamsvcRangeResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the range. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DHCP options. May be either a specific option or a group of options.`,
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.`,
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The end IP address of the range.`,
	},
	"exclusion_ranges": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcExclusionRangeResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of all exclusion ranges in the scope of the range.`,
	},
	"filters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcAccessFilterResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of all allow/deny filters of the range.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_assigned_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InheritanceAssignedHostResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: `The list of the inheritance assigned hosts of the object.`,
	},
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDHCPOptionsInheritanceResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The name of the range. May contain 1 to 256 characters. Can include UTF-8.`,
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of protocol (_ip4_ or _ip6_).`,
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The start IP address of the range.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the range in JSON format.`,
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationThresholdResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"utilization": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
	"utilization_v6": schema.SingleNestedAttribute{
		Attributes:          IpamsvcUtilizationV6ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcRange(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcRange {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcRangeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcRangeModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcRange {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcRange{
		Comment:            m.Comment.ValueStringPointer(),
		DhcpHost:           m.DhcpHost.ValueStringPointer(),
		DhcpOptions:        ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		DisableDhcp:        m.DisableDhcp.ValueBoolPointer(),
		End:                m.End.ValueString(),
		ExclusionRanges:    ExpandFrameworkListNestedBlock(ctx, m.ExclusionRanges, diags, expandIpamsvcExclusionRange),
		Filters:            ExpandFrameworkListNestedBlock(ctx, m.Filters, diags, expandIpamsvcAccessFilter),
		InheritanceParent:  m.InheritanceParent.ValueStringPointer(),
		InheritanceSources: expandIpamsvcDHCPOptionsInheritance(ctx, m.InheritanceSources, diags),
		Name:               m.Name.ValueStringPointer(),
		Parent:             m.Parent.ValueStringPointer(),
		Space:              m.Space.ValueString(),
		Start:              m.Start.ValueString(),
		Tags:               ExpandFrameworkMapString(ctx, m.Tags, diags),
		Threshold:          expandIpamsvcUtilizationThreshold(ctx, m.Threshold, diags),
		Utilization:        expandIpamsvcUtilization(ctx, m.Utilization, diags),
		UtilizationV6:      expandIpamsvcUtilizationV6(ctx, m.UtilizationV6, diags),
	}
	return to
}

func flattenIpamsvcRange(ctx context.Context, from *ipam.IpamsvcRange, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcRangeAttrTypes)
	}
	m := IpamsvcRangeModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcRangeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcRangeModel) flatten(ctx context.Context, from *ipam.IpamsvcRange, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcRangeModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpHost = types.StringPointerValue(from.DhcpHost)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.DisableDhcp = types.BoolPointerValue(from.DisableDhcp)
	m.End = types.StringValue(from.End)
	m.ExclusionRanges = FlattenFrameworkListNestedBlock(ctx, from.ExclusionRanges, IpamsvcExclusionRangeAttrTypes, diags, flattenIpamsvcExclusionRange)
	m.Filters = FlattenFrameworkListNestedBlock(ctx, from.Filters, IpamsvcAccessFilterAttrTypes, diags, flattenIpamsvcAccessFilter)
	m.Id = types.StringPointerValue(from.Id)
	m.InheritanceAssignedHosts = FlattenFrameworkListNestedBlock(ctx, from.InheritanceAssignedHosts, InheritanceAssignedHostAttrTypes, diags, flattenInheritanceAssignedHost)
	m.InheritanceParent = types.StringPointerValue(from.InheritanceParent)
	m.InheritanceSources = flattenIpamsvcDHCPOptionsInheritance(ctx, from.InheritanceSources, diags)
	m.Name = types.StringPointerValue(from.Name)
	m.Parent = types.StringPointerValue(from.Parent)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Space = types.StringValue(from.Space)
	m.Start = types.StringValue(from.Start)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Threshold = flattenIpamsvcUtilizationThreshold(ctx, from.Threshold, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.Utilization = flattenIpamsvcUtilization(ctx, from.Utilization, diags)
	m.UtilizationV6 = flattenIpamsvcUtilizationV6(ctx, from.UtilizationV6, diags)

}
