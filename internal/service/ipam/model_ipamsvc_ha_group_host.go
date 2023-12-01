package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcHAGroupHostModel struct {
	Address    types.String `tfsdk:"address"`
	Heartbeats types.List   `tfsdk:"heartbeats"`
	Host       types.String `tfsdk:"host"`
	Port       types.Int64  `tfsdk:"port"`
	Role       types.String `tfsdk:"role"`
	State      types.String `tfsdk:"state"`
}

var IpamsvcHAGroupHostAttrTypes = map[string]attr.Type{
	"address":    types.StringType,
	"heartbeats": types.ListType{ElemType: types.ObjectType{AttrTypes: IpamsvcHAGroupHeartbeatsAttrTypes}},
	"host":       types.StringType,
	"port":       types.Int64Type,
	"role":       types.StringType,
	"state":      types.StringType,
}

var IpamsvcHAGroupHostResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The address on which this host listens.",
	},
	"heartbeats": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: IpamsvcHAGroupHeartbeatsResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "Last successful heartbeat received from its peer/s. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.",
	},
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"port": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The HA port.",
	},
	"role": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.OneOf("active", "passive"),
		},
		MarkdownDescription: "The role of this host in the HA relationship: _active_ or _passive_.",
	},
	"state": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The state of DHCP on the host. This field is set when the _collect_stats_ is set to _true_ in the _GET_ _/dhcp/ha_group_ request.",
	},
}

func ExpandIpamsvcHAGroupHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcHAGroupHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcHAGroupHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHAGroupHost {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcHAGroupHost{
		Address:    flex.ExpandStringPointer(m.Address),
		Heartbeats: flex.ExpandFrameworkListNestedBlock(ctx, m.Heartbeats, diags, ExpandIpamsvcHAGroupHeartbeats),
		Host:       flex.ExpandString(m.Host),
		Role:       flex.ExpandStringPointer(m.Role),
		State:      flex.ExpandStringPointer(m.State),
	}
	return to
}

func FlattenIpamsvcHAGroupHost(ctx context.Context, from *ipam.IpamsvcHAGroupHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHAGroupHostAttrTypes)
	}
	m := IpamsvcHAGroupHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHAGroupHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHAGroupHostModel) Flatten(ctx context.Context, from *ipam.IpamsvcHAGroupHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHAGroupHostModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Heartbeats = flex.FlattenFrameworkListNestedBlock(ctx, from.Heartbeats, IpamsvcHAGroupHeartbeatsAttrTypes, diags, FlattenIpamsvcHAGroupHeartbeats)
	m.Host = flex.FlattenString(from.Host)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Role = flex.FlattenStringPointer(from.Role)
	m.State = flex.FlattenStringPointer(from.State)
}
