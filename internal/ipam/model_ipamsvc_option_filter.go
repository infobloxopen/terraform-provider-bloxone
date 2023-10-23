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

type IpamsvcOptionFilterModel struct {
	Comment                         types.String      `tfsdk:"comment"`
	CreatedAt                       timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions                     types.List        `tfsdk:"dhcp_options"`
	HeaderOptionFilename            types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String      `tfsdk:"header_option_server_name"`
	Id                              types.String      `tfsdk:"id"`
	LeaseTime                       types.Int64       `tfsdk:"lease_time"`
	Name                            types.String      `tfsdk:"name"`
	Protocol                        types.String      `tfsdk:"protocol"`
	Role                            types.String      `tfsdk:"role"`
	Rules                           types.Object      `tfsdk:"rules"`
	Tags                            types.Map         `tfsdk:"tags"`
	UpdatedAt                       timetypes.RFC3339 `tfsdk:"updated_at"`
	VendorSpecificOptionOptionSpace types.String      `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcOptionFilterAttrTypes = map[string]attr.Type{
	"comment":                             types.StringType,
	"created_at":                          timetypes.RFC3339Type{},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"id":                                  types.StringType,
	"lease_time":                          types.Int64Type,
	"name":                                types.StringType,
	"protocol":                            types.StringType,
	"role":                                types.StringType,
	"rules":                               types.ObjectType{AttrTypes: IpamsvcOptionFilterRuleListAttrTypes},
	"tags":                                types.MapType{},
	"updated_at":                          timetypes.RFC3339Type{},
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcOptionFilterResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionFilter__ object (_dhcp/option_filter_) applies options to DHCP clients with matching options. It must be configured in the _filters_ list for a scope to be effective.`,
	Attributes:          IpamsvcOptionFilterResourceSchemaAttributes,
}

var IpamsvcOptionFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the option filter. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcOptionItemResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `The list of DHCP options for the option filter. May be either a specific option or a group of options.`,
	},
	"header_option_filename": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option filename field.`,
	},
	"header_option_server_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option server address field.`,
	},
	"header_option_server_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The configuration for header option server name field.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: `The lease lifetime duration in seconds.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the option filter. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The type of protocol of option filter (_ip4_ or _ip6_).`,
	},
	"role": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.`,
	},
	"rules": schema.SingleNestedAttribute{
		Attributes:          IpamsvcOptionFilterRuleListResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: ``,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the option filter in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcOptionFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionFilterModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionFilter {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionFilter{
		Comment:                         m.Comment.ValueStringPointer(),
		DhcpOptions:                     ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		HeaderOptionFilename:            m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:       m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:          m.HeaderOptionServerName.ValueStringPointer(),
		LeaseTime:                       ptr(int64(m.LeaseTime.ValueInt64())),
		Name:                            m.Name.ValueString(),
		Protocol:                        m.Protocol.ValueStringPointer(),
		Role:                            m.Role.ValueStringPointer(),
		Rules:                           *expandIpamsvcOptionFilterRuleList(ctx, m.Rules, diags),
		Tags:                            ExpandFrameworkMapString(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: m.VendorSpecificOptionOptionSpace.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcOptionFilter(ctx context.Context, from *ipam.IpamsvcOptionFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionFilterAttrTypes)
	}
	m := IpamsvcOptionFilterModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionFilterModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionFilterModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.Id = types.StringPointerValue(from.Id)
	m.LeaseTime = types.Int64Value(int64(*from.LeaseTime))
	m.Name = types.StringValue(from.Name)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Role = types.StringPointerValue(from.Role)
	m.Rules = flattenIpamsvcOptionFilterRuleList(ctx, &from.Rules, diags)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.VendorSpecificOptionOptionSpace = types.StringPointerValue(from.VendorSpecificOptionOptionSpace)

}
