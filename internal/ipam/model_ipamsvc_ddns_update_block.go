package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcDDNSUpdateBlockModel struct {
	DdnsDomain      types.String `tfsdk:"ddns_domain"`
	DdnsSendUpdates types.Bool   `tfsdk:"ddns_send_updates"`
}

var IpamsvcDDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"ddns_domain":       types.StringType,
	"ddns_send_updates": types.BoolType,
}

var IpamsvcDDNSUpdateBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The dynamic DNS configurations, ddns_domain and ddns_send_updates.`,
	Attributes:          IpamsvcDDNSUpdateBlockResourceSchemaAttributes,
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

func expandIpamsvcDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDDNSUpdateBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDDNSUpdateBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDDNSUpdateBlock{
		DdnsDomain:      m.DdnsDomain.ValueStringPointer(),
		DdnsSendUpdates: m.DdnsSendUpdates.ValueBoolPointer(),
	}
	return to
}

func flattenIpamsvcDDNSUpdateBlock(ctx context.Context, from *ipam.IpamsvcDDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSUpdateBlockAttrTypes)
	}
	m := IpamsvcDDNSUpdateBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSUpdateBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcDDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSUpdateBlockModel{}
	}

	m.DdnsDomain = types.StringPointerValue(from.DdnsDomain)
	m.DdnsSendUpdates = types.BoolPointerValue(from.DdnsSendUpdates)

}
