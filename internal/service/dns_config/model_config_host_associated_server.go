package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigHostAssociatedServerModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

var ConfigHostAssociatedServerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

var ConfigHostAssociatedServerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "DNS server name.",
	},
}

func ExpandConfigHostAssociatedServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.HostAssociatedServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigHostAssociatedServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigHostAssociatedServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.HostAssociatedServer {
	if m == nil {
		return nil
	}
	to := &dnsconfig.HostAssociatedServer{}
	return to
}

func FlattenConfigHostAssociatedServer(ctx context.Context, from *dnsconfig.HostAssociatedServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigHostAssociatedServerAttrTypes)
	}
	m := ConfigHostAssociatedServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigHostAssociatedServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigHostAssociatedServerModel) Flatten(ctx context.Context, from *dnsconfig.HostAssociatedServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigHostAssociatedServerModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
}
