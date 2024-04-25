package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type IpamsvcDDNSUpdateBlockModel struct {
	DdnsDomain      types.String `tfsdk:"ddns_domain"`
	DdnsSendUpdates types.Bool   `tfsdk:"ddns_send_updates"`
}

var IpamsvcDDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"ddns_domain":       types.StringType,
	"ddns_send_updates": types.BoolType,
}

var IpamsvcDDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ddns_domain": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The domain name for DDNS.`,
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Determines if DDNS updates are enabled at this level.`,
	},
}

func ExpandIpamsvcDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.DDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcDDNSUpdateBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.DDNSUpdateBlock {
	if m == nil {
		return nil
	}
	to := &ipam.DDNSUpdateBlock{
		DdnsDomain:      m.DdnsDomain.ValueStringPointer(),
		DdnsSendUpdates: m.DdnsSendUpdates.ValueBoolPointer(),
	}
	return to
}

func FlattenIpamsvcDDNSUpdateBlock(ctx context.Context, from *ipam.DDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSUpdateBlockAttrTypes)
	}
	m := IpamsvcDDNSUpdateBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSUpdateBlockModel) Flatten(ctx context.Context, from *ipam.DDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSUpdateBlockModel{}
	}
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)
}
