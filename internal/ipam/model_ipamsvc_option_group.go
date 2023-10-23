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

type IpamsvcOptionGroupModel struct {
	Comment     types.String      `tfsdk:"comment"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions types.List        `tfsdk:"dhcp_options"`
	Id          types.String      `tfsdk:"id"`
	Name        types.String      `tfsdk:"name"`
	Protocol    types.String      `tfsdk:"protocol"`
	Tags        types.Map         `tfsdk:"tags"`
	UpdatedAt   timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcOptionGroupAttrTypes = map[string]attr.Type{
	"comment":      types.StringType,
	"created_at":   timetypes.RFC3339Type{},
	"dhcp_options": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcOptionItemAttrTypes}},
	"id":           types.StringType,
	"name":         types.StringType,
	"protocol":     types.StringType,
	"tags":         types.MapType{},
	"updated_at":   timetypes.RFC3339Type{},
}

var IpamsvcOptionGroupResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionGroup__ object (_dhcp/option_group_) is a named collection of options.`,
	Attributes:          IpamsvcOptionGroupResourceSchemaAttributes,
}

var IpamsvcOptionGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.`,
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
		MarkdownDescription: `The list of DHCP options for the option group. May be either a specific option or a group of options.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The type of protocol (_ip4_ or _ip6_).`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the option group in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func expandIpamsvcOptionGroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionGroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionGroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionGroupModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionGroup {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionGroup{
		Comment:     m.Comment.ValueStringPointer(),
		DhcpOptions: ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, expandIpamsvcOptionItem),
		Name:        m.Name.ValueString(),
		Protocol:    m.Protocol.ValueStringPointer(),
		Tags:        ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcOptionGroup(ctx context.Context, from *ipam.IpamsvcOptionGroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionGroupAttrTypes)
	}
	m := IpamsvcOptionGroupModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionGroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionGroupModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionGroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionGroupModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.DhcpOptions = FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, IpamsvcOptionItemAttrTypes, diags, flattenIpamsvcOptionItem)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringValue(from.Name)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
