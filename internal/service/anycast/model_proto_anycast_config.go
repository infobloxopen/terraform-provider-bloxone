package anycast

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtoAnycastConfigModel struct {
	AccountId          types.Int64       `tfsdk:"account_id"`
	AnycastIpAddress   types.String      `tfsdk:"anycast_ip_address"`
	AnycastIpv6Address types.String      `tfsdk:"anycast_ipv6_address"`
	CreatedAt          timetypes.RFC3339 `tfsdk:"created_at"`
	Description        types.String      `tfsdk:"description"`
	Id                 types.Int64       `tfsdk:"id"`
	IsConfigured       types.Bool        `tfsdk:"is_configured"`
	Name               types.String      `tfsdk:"name"`
	OnpremHosts        types.List        `tfsdk:"onprem_hosts"`
	RuntimeStatus      types.String      `tfsdk:"runtime_status"`
	Service            types.String      `tfsdk:"service"`
	Tags               types.Map         `tfsdk:"tags"`
	UpdatedAt          timetypes.RFC3339 `tfsdk:"updated_at"`
}

var ProtoAnycastConfigAttrTypes = map[string]attr.Type{
	"account_id":           types.Int64Type,
	"anycast_ip_address":   types.StringType,
	"anycast_ipv6_address": types.StringType,
	"created_at":           timetypes.RFC3339Type{},
	"description":          types.StringType,
	"id":                   types.Int64Type,
	"is_configured":        types.BoolType,
	"name":                 types.StringType,
	"onprem_hosts":         types.ListType{ElemType: types.ObjectType{AttrTypes: ProtoOnpremHostRefAttrTypes}},
	"runtime_status":       types.StringType,
	"service":              types.StringType,
	"tags":                 types.MapType{ElemType: types.StringType},
	"updated_at":           timetypes.RFC3339Type{},
}

var ProtoAnycastConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"account_id": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The account identifier.",
	},
	"anycast_ip_address": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "IPv4 address of the host in string format.",
	},
	"anycast_ipv6_address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IPv6 address of the host in string format",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for the address object. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"is_configured": schema.BoolAttribute{
		Computed: true,
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name of the anycast configuration.`,
	},
	"onprem_hosts": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ProtoOnpremHostRefResourceSchemaAttributes,
		},
		Optional: true,
	},
	"runtime_status": schema.StringAttribute{
		Computed: true,
	},
	"service": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The type of the Service used in anycast configuration, supports (`dns`, `dhcp`, 'dfp').",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Tagging specifics.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandProtoAnycastConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtoAnycastConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtoAnycastConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtoAnycastConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtoAnycastConfig {
	if m == nil {
		return nil
	}
	to := &anycast.ProtoAnycastConfig{
		AccountId:          flex.ExpandInt64Pointer(m.AccountId),
		AnycastIpAddress:   flex.ExpandStringPointer(m.AnycastIpAddress),
		AnycastIpv6Address: flex.ExpandStringPointer(m.AnycastIpv6Address),
		CreatedAt:          flex.ExpandTimePointer(ctx, m.CreatedAt, diags),
		Description:        flex.ExpandStringPointer(m.Description),
		IsConfigured:       flex.ExpandBoolPointer(m.IsConfigured),
		Name:               flex.ExpandStringPointer(m.Name),
		OnpremHosts:        flex.ExpandFrameworkListNestedBlock(ctx, m.OnpremHosts, diags, ExpandProtoOnpremHostRef),
		RuntimeStatus:      flex.ExpandStringPointer(m.RuntimeStatus),
		Service:            flex.ExpandStringPointer(m.Service),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		UpdatedAt:          flex.ExpandTimePointer(ctx, m.UpdatedAt, diags),
	}
	return to
}

func FlattenProtoAnycastConfig(ctx context.Context, from *anycast.ProtoAnycastConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtoAnycastConfigAttrTypes)
	}
	m := ProtoAnycastConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtoAnycastConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtoAnycastConfigModel) Flatten(ctx context.Context, from *anycast.ProtoAnycastConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtoAnycastConfigModel{}
	}
	m.AccountId = flex.FlattenInt64Pointer(from.AccountId)
	m.AnycastIpAddress = flex.FlattenStringPointer(from.AnycastIpAddress)
	m.AnycastIpv6Address = flex.FlattenStringPointer(from.AnycastIpv6Address)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IsConfigured = types.BoolPointerValue(from.IsConfigured)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OnpremHosts = flex.FlattenFrameworkListNestedBlock(ctx, from.OnpremHosts, ProtoOnpremHostRefAttrTypes, diags, FlattenProtoOnpremHostRef)
	m.RuntimeStatus = flex.FlattenStringPointer(from.RuntimeStatus)
	m.Service = flex.FlattenStringPointer(from.Service)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
