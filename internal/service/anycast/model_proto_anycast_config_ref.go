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

type ProtoAnycastConfigRefModel struct {
	AnycastConfigName types.String `tfsdk:"anycast_config_name"`
	RoutingProtocols  types.List   `tfsdk:"routing_protocols"`
}

var ProtoAnycastConfigRefAttrTypes = map[string]attr.Type{
	"anycast_config_name": types.StringType,
	"routing_protocols":   types.ListType{ElemType: types.StringType},
}

var ProtoAnycastConfigRefResourceSchemaAttributes = map[string]schema.Attribute{
	"anycast_config_name": schema.StringAttribute{
		Optional: true,
	},
	"routing_protocols": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Routing protocols enabled for this anycast configuration, on a particular host. Valid protocol names are \"BGP\", \"OSPF\"/\"OSPFv2\", \"OSPFv3\".",
	},
}

func ExpandProtoAnycastConfigRef(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtoAnycastConfigRef {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoAnycastConfigRefModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoAnycastConfigRefModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtoAnycastConfigRef {
	if m == nil {
		return nil
	}
	to := &anycast.ProtoAnycastConfigRef{
		AnycastConfigName: flex.ExpandStringPointer(m.AnycastConfigName),
		RoutingProtocols:  flex.ExpandFrameworkListString(ctx, m.RoutingProtocols, diags),
	}
	return to
}

func FlattenProtoAnycastConfigRef(ctx context.Context, from *anycast.ProtoAnycastConfigRef, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoAnycastConfigRefAttrTypes)
	}
	m := ProtoAnycastConfigRefModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoAnycastConfigRefAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoAnycastConfigRefModel) Flatten(ctx context.Context, from *anycast.ProtoAnycastConfigRef, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoAnycastConfigRefModel{}
	}
	m.AnycastConfigName = flex.FlattenStringPointer(from.AnycastConfigName)
	m.RoutingProtocols = flex.FlattenFrameworkListString(ctx, from.RoutingProtocols, diags)
}
