package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	internaltypes "github.com/infobloxopen/terraform-provider-bloxone/internal/types"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoOnpremHostModel struct {
	AnycastConfigRefs internaltypes.UnorderedListValue `tfsdk:"anycast_config_refs"`
	ConfigBgp         types.Object                     `tfsdk:"config_bgp"`
	ConfigOspf        types.Object                     `tfsdk:"config_ospf"`
	ConfigOspfv3      types.Object                     `tfsdk:"config_ospfv3"`
	CreatedAt         timetypes.RFC3339                `tfsdk:"created_at"`
	Id                types.Int64                      `tfsdk:"id"`
	IpAddress         types.String                     `tfsdk:"ip_address"`
	Ipv6Address       types.String                     `tfsdk:"ipv6_address"`
	Name              types.String                     `tfsdk:"name"`
	UpdatedAt         timetypes.RFC3339                `tfsdk:"updated_at"`
}

var ProtoOnpremHostResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_refs": schema.ListNestedAttribute{
		CustomType: internaltypes.UnorderedList{ListType: basetypes.ListType{ElemType: basetypes.ObjectType{AttrTypes: ProtoAnycastConfigRefAttrTypes}}},
		NestedObject: schema.NestedAttributeObject{
			Attributes: ProtoAnycastConfigRefResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: `Array of AnycastConfigRef structures, identifying the anycast configurations that this host is a member of.`,
	},
	"config_bgp": schema.SingleNestedAttribute{
		Attributes:          ProtoBgpConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: `Struct BGP configuration; defines BGP configuration for one anycast-enabled on-prem host.`,
	},
	"config_ospf": schema.SingleNestedAttribute{
		Attributes:          ProtoOspfConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: `Struct OSPF configuration; defines OSPF configuration for one anycast-enabled on-prem host.`,
	},
	"config_ospfv3": schema.SingleNestedAttribute{
		Attributes:          ProtoOspfv3ConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: `Struct OSPFv3 configuration; defines OSPFv3 configuration for one anycast-enabled on-prem host.`,
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Date/time this host was created in anycast service database.`,
	},
	"id": schema.Int64Attribute{
		Required: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: `Numeric host identifier.`,
	},
	"ip_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "IPv4 address of the on-prem host",
	},
	"ipv6_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "IPv6 address of the on-prem host",
	},
	"name": schema.StringAttribute{
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: `User-friendly name of the host @example "dns-host-1", "Central Office Server".`,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: `Date/time this host was last updated in anycast service database.`,
	},
}

func (m *ProtoOnpremHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.OnpremHost {
	if m == nil {
		return nil
	}
	to := &anycast.OnpremHost{
		AnycastConfigRefs: flex.ExpandFrameworkListNestedBlock(ctx, m.AnycastConfigRefs, diags, ExpandProtoAnycastConfigRef),
		ConfigBgp:         ExpandProtoBgpConfig(ctx, m.ConfigBgp, diags),
		ConfigOspf:        ExpandProtoOspfConfig(ctx, m.ConfigOspf, diags),
		ConfigOspfv3:      ExpandProtoOspfv3Config(ctx, m.ConfigOspfv3, diags),
		IpAddress:         flex.ExpandStringPointer(m.IpAddress),
		Ipv6Address:       flex.ExpandStringPointer(m.Ipv6Address),
		Name:              flex.ExpandStringPointer(m.Name),
		Id:                flex.ExpandInt64Pointer(m.Id),
	}
	return to
}

func (m *ProtoOnpremHostModel) Flatten(ctx context.Context, from *anycast.OnpremHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoOnpremHostModel{}
	}
	m.AnycastConfigRefs = flex.FlattenFrameworkUnorderedListNestedBlock(ctx, from.AnycastConfigRefs, ProtoAnycastConfigRefAttrTypes, diags, FlattenProtoAnycastConfigRef)
	m.ConfigOspf = FlattenProtoOspfConfig(ctx, from.ConfigOspf, diags)
	m.ConfigOspfv3 = FlattenProtoOspfv3Config(ctx, from.ConfigOspfv3, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
