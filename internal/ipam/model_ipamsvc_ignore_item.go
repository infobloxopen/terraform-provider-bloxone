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

type IpamsvcIgnoreItemModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

var IpamsvcIgnoreItemAttrTypes = map[string]attr.Type{
	"type":  types.StringType,
	"value": types.StringType,
}

var IpamsvcIgnoreItemResourceSchema = schema.Schema{
	MarkdownDescription: `An Ignore Item object (_dhcp/ignore_item_) represents an item in a DHCP ignore list.`,
	Attributes:          IpamsvcIgnoreItemResourceSchemaAttributes,
}

var IpamsvcIgnoreItemResourceSchemaAttributes = map[string]schema.Attribute{
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:  * _client_hex_,  * _client_text_,  * _hardware_.`,
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Value to match.`,
	},
}

func expandIpamsvcIgnoreItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcIgnoreItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcIgnoreItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcIgnoreItemModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcIgnoreItem {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcIgnoreItem{
		Type:  m.Type.ValueString(),
		Value: m.Value.ValueString(),
	}
	return to
}

func flattenIpamsvcIgnoreItem(ctx context.Context, from *ipam.IpamsvcIgnoreItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIgnoreItemAttrTypes)
	}
	m := IpamsvcIgnoreItemModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIgnoreItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIgnoreItemModel) flatten(ctx context.Context, from *ipam.IpamsvcIgnoreItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIgnoreItemModel{}
	}

	m.Type = types.StringValue(from.Type)
	m.Value = types.StringValue(from.Value)

}
