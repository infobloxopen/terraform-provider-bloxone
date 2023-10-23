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

type IpamsvcHAGroupModel struct {
	AnycastConfigId types.String      `tfsdk:"anycast_config_id"`
	Comment         types.String      `tfsdk:"comment"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Hosts           types.List        `tfsdk:"hosts"`
	Id              types.String      `tfsdk:"id"`
	IpSpace         types.String      `tfsdk:"ip_space"`
	Mode            types.String      `tfsdk:"mode"`
	Name            types.String      `tfsdk:"name"`
	Status          types.String      `tfsdk:"status"`
	Tags            types.Map         `tfsdk:"tags"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var IpamsvcHAGroupAttrTypes = map[string]attr.Type{
	"anycast_config_id": types.StringType,
	"comment":           types.StringType,
	"created_at":        timetypes.RFC3339Type{},
	"hosts":             types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHAGroupHostAttrTypes}},
	"id":                types.StringType,
	"ip_space":          types.StringType,
	"mode":              types.StringType,
	"name":              types.StringType,
	"status":            types.StringType,
	"tags":              types.MapType{},
	"updated_at":        timetypes.RFC3339Type{},
}

var IpamsvcHAGroupResourceSchema = schema.Schema{
	MarkdownDescription: `An __HAGroup__ object (_dhcp/ha_group_) represents on-prem hosts that can serve the same leases for HA.`,
	Attributes:          IpamsvcHAGroupResourceSchemaAttributes,
}

var IpamsvcHAGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The description for the HA group. May contain 0 to 1024 characters. Can include UTF-8.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been created.`,
	},
	"hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcHAGroupHostResourceSchemaAttributes,
		},
		Required:            true,
		MarkdownDescription: `The list of hosts.`,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"mode": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The mode of the HA group.  Valid values are: * _active-active_: Both on-prem hosts remain active. * _active-passive_: One on-prem host remains active and one remains passive. When the active on-prem host is down, the passive on-prem host takes over. * _advanced-active-passive_: One on-prem host may be part of multiple HA groups. When the active on-prem host is down, the passive on-prem host takes over.`,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the HA group. Must contain 1 to 256 characters. Can include UTF-8.`,
	},
	"status": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Status of the HA group. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.`,
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: `The tags for the HA group.`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Time when the object has been updated. Equals to _created_at_ if not updated after creation.`,
	},
}

func expandIpamsvcHAGroup(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroup {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHAGroupModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHAGroupModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroup {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHAGroup{
		AnycastConfigId: m.AnycastConfigId.ValueStringPointer(),
		Comment:         m.Comment.ValueStringPointer(),
		Hosts:           ExpandFrameworkListNestedBlock(ctx, m.Hosts, diags, expandIpamsvcHAGroupHost),
		IpSpace:         m.IpSpace.ValueStringPointer(),
		Mode:            m.Mode.ValueStringPointer(),
		Name:            m.Name.ValueString(),
		Status:          m.Status.ValueStringPointer(),
		Tags:            ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func flattenIpamsvcHAGroup(ctx context.Context, from *ipam.IpamsvcHAGroup, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupAttrTypes)
	}
	m := IpamsvcHAGroupModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupModel) flatten(ctx context.Context, from *ipam.IpamsvcHAGroup, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupModel{}
	}

	m.AnycastConfigId = types.StringPointerValue(from.AnycastConfigId)
	m.Comment = types.StringPointerValue(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Hosts = FlattenFrameworkListNestedBlock(ctx, from.Hosts, IpamsvcHAGroupHostAttrTypes, diags, flattenIpamsvcHAGroupHost)
	m.Id = types.StringPointerValue(from.Id)
	m.IpSpace = types.StringPointerValue(from.IpSpace)
	m.Mode = types.StringPointerValue(from.Mode)
	m.Name = types.StringValue(from.Name)
	m.Status = types.StringPointerValue(from.Status)
	m.Tags = FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)

}
