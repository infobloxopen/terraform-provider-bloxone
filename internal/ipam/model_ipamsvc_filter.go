package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcFilterModel struct {
	Comment  types.String `tfsdk:"comment"`
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Protocol types.String `tfsdk:"protocol"`
	Role     types.String `tfsdk:"role"`
	Tags     types.Map    `tfsdk:"tags"`
	Type     types.String `tfsdk:"type"`
}

var IpamsvcFilterAttrTypes = map[string]attr.Type{
	"comment":  types.StringType,
	"id":       types.StringType,
	"name":     types.StringType,
	"protocol": types.StringType,
	"role":     types.StringType,
	"tags":     types.MapType{},
	"type":     types.StringType,
}

var IpamsvcFilterResourceSchema = schema.Schema{
	MarkdownDescription: `A DHCP Filter (_dhcp/filter_) object lists DHCP filters of all types.`,
	Attributes:          IpamsvcFilterResourceSchemaAttributes,
}

var IpamsvcFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The description for the DHCP filter. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The name of the DHCP filter.`,
	},
	"protocol": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of protocol of the filter (_ip4_ or _ip6_).`,
	},
	"role": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The role of DHCP filter (_values_ or _selection_).`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the DHCP filter in JSON format.`,
	},
	"type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The type of DHCP filter (_hardware_ or _option_).`,
	},
}

func expandIpamsvcFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcFilterModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcFilter {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcFilter{
		Role: m.Role.ValueStringPointer(),
		Tags: ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcFilter(ctx context.Context, from *ipam.IpamsvcFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcFilterAttrTypes)
	}
	m := IpamsvcFilterModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcFilterModel) flatten(ctx context.Context, from *ipam.IpamsvcFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcFilterModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringPointerValue(from.Name)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Role = types.StringPointerValue(from.Role)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = types.StringPointerValue(from.Type)

}
