package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoOspfv3ConfigModel struct {
	Area               types.String `tfsdk:"area"`
	Cost               types.Int64  `tfsdk:"cost"`
	DeadInterval       types.Int64  `tfsdk:"dead_interval"`
	HelloInterval      types.Int64  `tfsdk:"hello_interval"`
	Interface          types.String `tfsdk:"interface"`
	RetransmitInterval types.Int64  `tfsdk:"retransmit_interval"`
	TransmitDelay      types.Int64  `tfsdk:"transmit_delay"`
}

var ProtoOspfv3ConfigAttrTypes = map[string]attr.Type{
	"area":                types.StringType,
	"cost":                types.Int64Type,
	"dead_interval":       types.Int64Type,
	"hello_interval":      types.Int64Type,
	"interface":           types.StringType,
	"retransmit_interval": types.Int64Type,
	"transmit_delay":      types.Int64Type,
}

var ProtoOspfv3ConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"area": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "OSPF area identifier; usually in the format of an IPv4 address (although not an address itself)",
	},
	"cost": schema.Int64Attribute{
		Optional: true,
	},
	"dead_interval": schema.Int64Attribute{
		Optional: true,
	},
	"hello_interval": schema.Int64Attribute{
		Optional: true,
	},
	"interface": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Name of the interface that is configured with external IP address of the host",
	},
	"retransmit_interval": schema.Int64Attribute{
		Optional: true,
	},
	"transmit_delay": schema.Int64Attribute{
		Optional: true,
	},
}

func ExpandProtoOspfv3Config(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.Ospfv3Config {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoOspfv3ConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoOspfv3ConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.Ospfv3Config {
	if m == nil {
		return nil
	}
	to := &anycast.Ospfv3Config{
		Area:               flex.ExpandStringPointer(m.Area),
		Cost:               flex.ExpandInt64Pointer(m.Cost),
		DeadInterval:       flex.ExpandInt64Pointer(m.DeadInterval),
		HelloInterval:      flex.ExpandInt64Pointer(m.HelloInterval),
		Interface:          flex.ExpandStringPointer(m.Interface),
		RetransmitInterval: flex.ExpandInt64Pointer(m.RetransmitInterval),
		TransmitDelay:      flex.ExpandInt64Pointer(m.TransmitDelay),
	}
	return to
}

func FlattenProtoOspfv3Config(ctx context.Context, from *anycast.Ospfv3Config, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoOspfv3ConfigAttrTypes)
	}
	m := ProtoOspfv3ConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoOspfv3ConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoOspfv3ConfigModel) Flatten(ctx context.Context, from *anycast.Ospfv3Config, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoOspfv3ConfigModel{}
	}
	m.Area = flex.FlattenStringPointer(from.Area)
	m.Cost = flex.FlattenInt64Pointer(from.Cost)
	m.DeadInterval = flex.FlattenInt64Pointer(from.DeadInterval)
	m.HelloInterval = flex.FlattenInt64Pointer(from.HelloInterval)
	m.Interface = flex.FlattenStringPointer(from.Interface)
	m.RetransmitInterval = flex.FlattenInt64Pointer(from.RetransmitInterval)
	m.TransmitDelay = flex.FlattenInt64Pointer(from.TransmitDelay)
}
