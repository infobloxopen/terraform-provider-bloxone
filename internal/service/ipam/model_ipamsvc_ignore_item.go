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

type IpamsvcIgnoreItemModel struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

var IpamsvcIgnoreItemAttrTypes = map[string]attr.Type{
	"type":  types.StringType,
	"value": types.StringType,
}

var IpamsvcIgnoreItemResourceSchemaAttributes = map[string]schema.Attribute{
	"type": schema.StringAttribute{
		Required: true,
		MarkdownDescription: "Type of ignore matching: client to ignore by client identifier (client hex or client text) or hardware to ignore by hardware identifier (MAC address). It can have one of the following values:\n" +
			"  * _client_hex_\n" +
			"  * _client_text_\n" +
			"  * _hardware_",
	},
	"value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Value to match.`,
	},
}

func ExpandIpamsvcIgnoreItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IgnoreItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcIgnoreItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcIgnoreItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IgnoreItem {
	if m == nil {
		return nil
	}
	to := &ipam.IgnoreItem{
		Type:  m.Type.ValueString(),
		Value: m.Value.ValueString(),
	}
	return to
}

func FlattenIpamsvcIgnoreItem(ctx context.Context, from *ipam.IgnoreItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcIgnoreItemAttrTypes)
	}
	m := IpamsvcIgnoreItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcIgnoreItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcIgnoreItemModel) Flatten(ctx context.Context, from *ipam.IgnoreItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcIgnoreItemModel{}
	}
	m.Type = flex.FlattenString(from.Type)
	m.Value = flex.FlattenString(from.Value)
}
