package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IPAMConfigModel struct {
	DhcpServer            types.String `tfsdk:"dhcp_server"`
	DisableIpamProjection types.Bool   `tfsdk:"disable_ipam_projection"`
	IpSpace               types.String `tfsdk:"ip_space"`
}

var IPAMConfigAttrTypes = map[string]attr.Type{
	"dhcp_server":             types.StringType,
	"disable_ipam_projection": types.BoolType,
	"ip_space":                types.StringType,
}

var IPAMConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_server": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Address of the DHCP Server.",
	},
	"disable_ipam_projection": schema.BoolAttribute{
		Optional:            true,
		Computed:			 true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls the IPAM Sync/Reconciliation for the provider.",
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "IP Space.",
	},
}

func ExpandIPAMConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.IPAMConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IPAMConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IPAMConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.IPAMConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.IPAMConfig{
		DhcpServer:            flex.ExpandStringPointer(m.DhcpServer),
		DisableIpamProjection: flex.ExpandBoolPointer(m.DisableIpamProjection),
		IpSpace:               flex.ExpandStringPointer(m.IpSpace),
	}
	return to
}

func FlattenIPAMConfig(ctx context.Context, from *clouddiscovery.IPAMConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IPAMConfigAttrTypes)
	}
	m := IPAMConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IPAMConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IPAMConfigModel) Flatten(ctx context.Context, from *clouddiscovery.IPAMConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IPAMConfigModel{}
	}
	m.DhcpServer = flex.FlattenStringPointer(from.DhcpServer)
	m.DisableIpamProjection = flex.FlattenBoolPointerFalseAsNull(from.DisableIpamProjection)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
}
