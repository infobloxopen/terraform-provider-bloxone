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

type IpamsvcInheritedDDNSUpdateBlockModel struct {
	Action      types.String `tfsdk:"action"`
	DisplayName types.String `tfsdk:"display_name"`
	Source      types.String `tfsdk:"source"`
	Value       types.Object `tfsdk:"value"`
}

var IpamsvcInheritedDDNSUpdateBlockAttrTypes = map[string]attr.Type{
	"action":       types.StringType,
	"display_name": types.StringType,
	"source":       types.StringType,
	"value":        types.ObjectType{AttrTypes: IpamsvcDDNSUpdateBlockAttrTypes},
}

var IpamsvcInheritedDDNSUpdateBlockResourceSchema = schema.Schema{
	MarkdownDescription: `The inheritance configuration for ddns_domain and ddns_send_updates.`,
	Attributes:          IpamsvcInheritedDDNSUpdateBlockResourceSchemaAttributes,
}

var IpamsvcInheritedDDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The inheritance setting.  Valid values are: * _inherit_: Use the inherited value. * _override_: Use the value set in the object.  Defaults to _inherit_.`,
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.SingleNestedAttribute{
		Attributes:          IpamsvcDDNSUpdateBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: ``,
	},
}

func expandIpamsvcInheritedDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcInheritedDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcInheritedDDNSUpdateBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcInheritedDDNSUpdateBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcInheritedDDNSUpdateBlock{
		Action: m.Action.ValueStringPointer(),
		Source: m.Source.ValueStringPointer(),
		Value:  expandIpamsvcDDNSUpdateBlock(ctx, m.Value, diags),
	}
	return to
}

func flattenIpamsvcInheritedDDNSUpdateBlock(ctx context.Context, from *ipam.IpamsvcInheritedDDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDDNSUpdateBlockAttrTypes)
	}
	m := IpamsvcInheritedDDNSUpdateBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDDNSUpdateBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcInheritedDDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDDNSUpdateBlockModel{}
	}

	m.Action = types.StringPointerValue(from.Action)
	m.DisplayName = types.StringPointerValue(from.DisplayName)
	m.Source = types.StringPointerValue(from.Source)
	m.Value = flattenIpamsvcDDNSUpdateBlock(ctx, from.Value, diags)

}
