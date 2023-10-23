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

type IpamsvcDDNSHostnameBlockModel struct {
	DdnsGenerateName    types.Bool   `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix types.String `tfsdk:"ddns_generated_prefix"`
}

var IpamsvcDDNSHostnameBlockAttrTypes = map[string]attr.Type{
	"ddns_generate_name":    types.BoolType,
	"ddns_generated_prefix": types.StringType,
}

var IpamsvcDDNSHostnameBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The dynamic DNS Hostname configuration.`,
	Attributes:          IpamsvcDDNSHostnameBlockResourceSchemaAttributes,
}

var IpamsvcDDNSHostnameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates if DDNS should generate a hostname when not supplied by the client.`,
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The prefix used in the generation of an FQDN.`,
	},
}

func expandIpamsvcDDNSHostnameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcDDNSHostnameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcDDNSHostnameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcDDNSHostnameBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcDDNSHostnameBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcDDNSHostnameBlock{
		DdnsGenerateName:    m.DdnsGenerateName.ValueBoolPointer(),
		DdnsGeneratedPrefix: m.DdnsGeneratedPrefix.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcDDNSHostnameBlock(ctx context.Context, from *ipam.IpamsvcDDNSHostnameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcDDNSHostnameBlockAttrTypes)
	}
	m := IpamsvcDDNSHostnameBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcDDNSHostnameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcDDNSHostnameBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcDDNSHostnameBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcDDNSHostnameBlockModel{}
	}

	m.DdnsGenerateName = types.BoolPointerValue(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = types.StringPointerValue(from.DdnsGeneratedPrefix)

}
