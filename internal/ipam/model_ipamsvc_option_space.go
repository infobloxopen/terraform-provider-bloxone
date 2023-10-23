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

type IpamsvcOptionSpaceModel struct {
	Comment   types.String      `tfsdk:"comment"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at"`
	Id        types.String      `tfsdk:"id"`
	Name      types.String      `tfsdk:"name"`
	Protocol  types.String      `tfsdk:"protocol"`
	Tags      types.Map         `tfsdk:"tags"`
	UpdatedAt timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcOptionSpaceAttrTypes = map[string]attr.Type{
	"comment":    types.StringType,
	"created_at": timetypes.RFC3339Type{},
	"id":         types.StringType,
	"name":       types.StringType,
	"protocol":   types.StringType,
	"tags":       types.MapType{},
	"updated_at": timetypes.RFC3339Type{},
}

var IpamsvcOptionSpaceResourceSchema = schema.Schema{
	MarkdownDescription: `An __OptionSpace__ object (_dhcp/option_space_) represents a set of DHCP option codes.`,
	Attributes:          IpamsvcOptionSpaceResourceSchemaAttributes,
}

var IpamsvcOptionSpaceResourceSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the option space. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the option space. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The type of protocol for the option space (_ip4_ or _ip6_).`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the option space in JSON format.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func expandIpamsvcOptionSpace(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcOptionSpace {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcOptionSpaceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcOptionSpaceModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcOptionSpace {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcOptionSpace{
		Comment:  m.Comment.ValueStringPointer(),
		Name:     m.Name.ValueString(),
		Protocol: m.Protocol.ValueStringPointer(),
		Tags:     ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcOptionSpace(ctx context.Context, from *ipam.IpamsvcOptionSpace, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcOptionSpaceAttrTypes)
	}
	m := IpamsvcOptionSpaceModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcOptionSpaceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcOptionSpaceModel) flatten(ctx context.Context, from *ipam.IpamsvcOptionSpace, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcOptionSpaceModel{}
	}

	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringValue(from.Name)
	m.Protocol = types.StringPointerValue(from.Protocol)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
