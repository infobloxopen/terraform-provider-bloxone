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

type IpamsvcHardwareFilterModel struct {
	Addresses                       types.List        `tfsdk:"addresses"`
	Comment                         types.String      `tfsdk:"comment"`
	CreatedAt                       timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions                     types.List        `tfsdk:"dhcp_options"`
	HeaderOptionFilename            types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String      `tfsdk:"header_option_server_name"`
	Id                              types.String      `tfsdk:"id"`
	LeaseTime                       types.Int64       `tfsdk:"lease_time"`
	Name                            types.String      `tfsdk:"name"`
	Role                            types.String      `tfsdk:"role"`
	Tags                            types.Map         `tfsdk:"tags"`
	UpdatedAt                       timetypes.RFC3339 `tfsdk:"updated_at"`
	VendorSpecificOptionOptionSpace types.String      `tfsdk:"vendor_specific_option_option_space"`
}

var IpamsvcHardwareFilterAttrTypes = map[string]attr.Type{
	"addresses":                           types.ListType{ElemType: types.StringType},
	"comment":                             types.StringType,
	"created_at":                          timetypes.RFC3339Type{},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"id":                                  types.StringType,
	"lease_time":                          types.Int64Type,
	"name":                                types.StringType,
	"role":                                types.StringType,
	"tags":                                types.MapType{},
	"updated_at":                          timetypes.RFC3339Type{},
	"vendor_specific_option_option_space": types.StringType,
}

var IpamsvcHardwareFilterResourceSchema = schema.Schema{
	MarkdownDescription: `A __HardwareFilter__ object (_dhcp/hardware_filter_) applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.`,
	Attributes:          IpamsvcHardwareFilterResourceSchemaAttributes,
}

var IpamsvcHardwareFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The list of addresses to match for the hardware filter.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the hardware filter. May contain 0 to 1024 characters. Can include UTF-8.`,
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
		MarkdownDescription: `The list of DHCP options for the hardware filter. May be either a specific option or a group of options.`,
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
		MarkdownDescription: `The name of the hardware filter. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"role": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the hardware filter in JSON format.`,
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

func expandIpamsvcHardwareFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHardwareFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHardwareFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHardwareFilterModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHardwareFilter {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHardwareFilter{
		Addresses:                       ExpandFrameworkListString(ctx, m.Addresses, diags),
		Comment:                         m.Comment.ValueStringPointer(),
		DhcpOptions:                     ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		HeaderOptionFilename:            m.HeaderOptionFilename.ValueStringPointer(),
		HeaderOptionServerAddress:       m.HeaderOptionServerAddress.ValueStringPointer(),
		HeaderOptionServerName:          m.HeaderOptionServerName.ValueStringPointer(),
		LeaseTime:                       ptr(int64(m.LeaseTime.ValueInt64())),
		Name:                            m.Name.ValueString(),
		Role:                            m.Role.ValueStringPointer(),
		Tags:                            ExpandFrameworkMapString(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: m.VendorSpecificOptionOptionSpace.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcHardwareFilter(ctx context.Context, from *ipam.IpamsvcHardwareFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHardwareFilterAttrTypes)
	}
	m := IpamsvcHardwareFilterModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHardwareFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHardwareFilterModel) flatten(ctx context.Context, from *ipam.IpamsvcHardwareFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHardwareFilterModel{}
	}

	m.Addresses = FlattenFrameworkListString(ctx, from.Addresses, diags)
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.HeaderOptionFilename = types.StringPointerValue(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = types.StringPointerValue(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = types.StringPointerValue(from.HeaderOptionServerName)
	m.Id = types.StringPointerValue(from.Id)
	m.LeaseTime = types.Int64Value(int64(*from.LeaseTime))
	m.Name = types.StringValue(from.Name)
	m.Role = types.StringPointerValue(from.Role)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
	m.VendorSpecificOptionOptionSpace = types.StringPointerValue(from.VendorSpecificOptionOptionSpace)

}
