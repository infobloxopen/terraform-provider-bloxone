package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoOnpremHostModel struct {
	AnycastConfigRefs types.List        `tfsdk:"anycast_config_refs"`
	ConfigBgp         types.Object      `tfsdk:"config_bgp"`
	ConfigOspf        types.Object      `tfsdk:"config_ospf"`
	ConfigOspfv3      types.Object      `tfsdk:"config_ospfv3"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	Id                types.Int64       `tfsdk:"id"`
	IpAddress         types.String      `tfsdk:"ip_address"`
	Ipv6Address       types.String      `tfsdk:"ipv6_address"`
	Name              types.String      `tfsdk:"name"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
}

var ProtoOnpremHostAttrTypes = map[string]attr.Type{
	"anycast_config_refs": types.ListType{ElemType: types.ObjectType{AttrTypes: ProtoAnycastConfigRefAttrTypes}},
	"config_bgp":          types.ObjectType{AttrTypes: ProtoBgpConfigAttrTypes},
	"config_ospf":         types.ObjectType{AttrTypes: ProtoOspfConfigAttrTypes},
	"config_ospfv3":       types.ObjectType{AttrTypes: ProtoOspfv3ConfigAttrTypes},
	"created_at":          timetypes.RFC3339Type{},
	"id":                  types.Int64Type,
	"ip_address":          types.StringType,
	"ipv6_address":        types.StringType,
	"name":                types.StringType,
	"updated_at":          timetypes.RFC3339Type{},
}

var ProtoOnpremHostResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_refs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ProtoAnycastConfigRefResourceSchemaAttributes,
		},
		Optional: true,
	},
	"config_bgp": schema.SingleNestedAttribute{
		Attributes: ProtoBgpConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"config_ospf": schema.SingleNestedAttribute{
		Attributes: ProtoOspfConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"config_ospfv3": schema.SingleNestedAttribute{
		Attributes: ProtoOspfv3ConfigResourceSchemaAttributes,
		Optional:   true,
	},
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Optional:   true,
	},
	"id": schema.Int64Attribute{
		Computed: true,
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv4 address of the on-prem host",
	},
	"ipv6_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv6 address of the on-prem host",
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Optional:   true,
	},
}

func ExpandProtoOnpremHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtoOnpremHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoOnpremHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoOnpremHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtoOnpremHost {
	if m == nil {
		return nil
	}
	to := &anycast.ProtoOnpremHost{
		AnycastConfigRefs: flex.ExpandFrameworkListNestedBlock(ctx, m.AnycastConfigRefs, diags, ExpandProtoAnycastConfigRef),
		ConfigBgp:         ExpandProtoBgpConfig(ctx, m.ConfigBgp, diags),
		ConfigOspf:        ExpandProtoOspfConfig(ctx, m.ConfigOspf, diags),
		ConfigOspfv3:      ExpandProtoOspfv3Config(ctx, m.ConfigOspfv3, diags),
		CreatedAt:         flex.ExpandTimePointer(ctx, m.CreatedAt, diags),
		IpAddress:         flex.ExpandStringPointer(m.IpAddress),
		Ipv6Address:       flex.ExpandStringPointer(m.Ipv6Address),
		Name:              flex.ExpandStringPointer(m.Name),
		UpdatedAt:         flex.ExpandTimePointer(ctx, m.UpdatedAt, diags),
	}
	return to
}

func FlattenProtoOnpremHost(ctx context.Context, from *anycast.ProtoOnpremHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoOnpremHostAttrTypes)
	}
	m := ProtoOnpremHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoOnpremHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoOnpremHostModel) Flatten(ctx context.Context, from *anycast.ProtoOnpremHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoOnpremHostModel{}
	}
	m.AnycastConfigRefs = flex.FlattenFrameworkListNestedBlock(ctx, from.AnycastConfigRefs, ProtoAnycastConfigRefAttrTypes, diags, FlattenProtoAnycastConfigRef)
	m.ConfigBgp = FlattenProtoBgpConfig(ctx, from.ConfigBgp, diags)
	m.ConfigOspf = FlattenProtoOspfConfig(ctx, from.ConfigOspf, diags)
	m.ConfigOspfv3 = FlattenProtoOspfv3Config(ctx, from.ConfigOspfv3, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
