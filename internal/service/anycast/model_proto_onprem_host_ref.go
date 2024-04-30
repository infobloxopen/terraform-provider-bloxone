package anycast

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoOnpremHostRefModel struct {
	Id            types.Int64  `tfsdk:"id"`
	IpAddress     types.String `tfsdk:"ip_address"`
	Ipv6Address   types.String `tfsdk:"ipv6_address"`
	Name          types.String `tfsdk:"name"`
	Ophid         types.String `tfsdk:"ophid"`
	RuntimeStatus types.String `tfsdk:"runtime_status"`
}

var ProtoOnpremHostRefAttrTypes = map[string]attr.Type{
	"id":             types.Int64Type,
	"ip_address":     types.StringType,
	"ipv6_address":   types.StringType,
	"name":           types.StringType,
	"ophid":          types.StringType,
	"runtime_status": types.StringType,
}

var ProtoOnpremHostRefResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int64Attribute{
		//Optional:            true,
		Required:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"ip_address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "IPv4 address of the host in string format",
	},
	"ipv6_address": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "IPv6 address of the host in string format",
	},
	"name": schema.StringAttribute{
		//Optional:            true,
		Computed:            true,
		MarkdownDescription: `The name of the anycast.`,
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Unique 32-character string identifier assigned to the host",
	},
	"runtime_status": schema.StringAttribute{
		Computed: true,
	},
}

func ExpandProtoOnpremHostRef(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtoOnpremHostRef {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoOnpremHostRefModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoOnpremHostRefModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtoOnpremHostRef {
	if m == nil {
		return nil
	}
	to := &anycast.ProtoOnpremHostRef{
		IpAddress:     flex.ExpandStringPointer(m.IpAddress),
		Id:            flex.ExpandInt64Pointer(m.Id),
		Ipv6Address:   flex.ExpandStringPointer(m.Ipv6Address),
		Name:          flex.ExpandStringPointer(m.Name),
		RuntimeStatus: flex.ExpandStringPointer(m.RuntimeStatus),
	}
	return to
}

func FlattenProtoOnpremHostRef(ctx context.Context, from *anycast.ProtoOnpremHostRef, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoOnpremHostRefAttrTypes)
	}
	m := ProtoOnpremHostRefModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoOnpremHostRefAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoOnpremHostRefModel) Flatten(ctx context.Context, from *anycast.ProtoOnpremHostRef, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoOnpremHostRefModel{}
	}
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.Ipv6Address = flex.FlattenStringPointer(from.Ipv6Address)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.RuntimeStatus = flex.FlattenStringPointer(from.RuntimeStatus)
}
