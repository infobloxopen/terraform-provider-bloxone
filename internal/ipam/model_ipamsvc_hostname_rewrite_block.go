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

type IpamsvcHostnameRewriteBlockModel struct {
	HostnameRewriteChar    types.String `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled types.Bool   `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex   types.String `tfsdk:"hostname_rewrite_regex"`
}

var IpamsvcHostnameRewriteBlockAttrTypes = map[string]attr.Type{
	"hostname_rewrite_char":    types.StringType,
	"hostname_rewrite_enabled": types.BoolType,
	"hostname_rewrite_regex":   types.StringType,
}

var IpamsvcHostnameRewriteBlockResourceSchema = schema.Schema{
	MarkdownDescription: `Hostname Rewrite grouping fields.`,
	Attributes:          IpamsvcHostnameRewriteBlockResourceSchemaAttributes,
}

var IpamsvcHostnameRewriteBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"hostname_rewrite_char": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance configuration for _hostname_rewrite_char_ field.`,
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance configuration for _hostname_rewrite_enabled_ field.`,
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance configuration for _hostname_rewrite_regex_ field.`,
	},
}

func expandIpamsvcHostnameRewriteBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHostnameRewriteBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHostnameRewriteBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHostnameRewriteBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHostnameRewriteBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHostnameRewriteBlock{
		HostnameRewriteChar:    m.HostnameRewriteChar.ValueStringPointer(),
		HostnameRewriteEnabled: m.HostnameRewriteEnabled.ValueBoolPointer(),
		HostnameRewriteRegex:   m.HostnameRewriteRegex.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcHostnameRewriteBlock(ctx context.Context, from *ipam.IpamsvcHostnameRewriteBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHostnameRewriteBlockAttrTypes)
	}
	m := IpamsvcHostnameRewriteBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHostnameRewriteBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHostnameRewriteBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcHostnameRewriteBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostnameRewriteBlockModel{}
	}

	m.HostnameRewriteChar = types.StringPointerValue(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = types.BoolPointerValue(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = types.StringPointerValue(from.HostnameRewriteRegex)

}
