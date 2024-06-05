package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
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

var IpamsvcInheritedDDNSUpdateBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"action": schema.StringAttribute{
		Optional: true,
		Computed: true,
		MarkdownDescription: "The inheritance setting. Valid values are:\n" +
			"  * _inherit_: Use the inherited value.\n" +
			"  * _override_: Use the value set in the object.\n\n" +
			"  Defaults to _inherit_.",
	},
	"display_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The human-readable display name for the object referred to by _source_.`,
	},
	"source": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"value": schema.SingleNestedAttribute{
		Attributes: utils.ToComputedAttributeMap(IpamsvcDDNSUpdateBlockResourceSchemaAttributes),
		Computed:   true,
	},
}

func ExpandIpamsvcInheritedDDNSUpdateBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.InheritedDDNSUpdateBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcInheritedDDNSUpdateBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcInheritedDDNSUpdateBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.InheritedDDNSUpdateBlock {
	if m == nil {
		return nil
	}
	to := &ipam.InheritedDDNSUpdateBlock{
		Action: m.Action.ValueStringPointer(),
	}
	return to
}

func FlattenIpamsvcInheritedDDNSUpdateBlock(ctx context.Context, from *ipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcInheritedDDNSUpdateBlockAttrTypes)
	}
	m := IpamsvcInheritedDDNSUpdateBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcInheritedDDNSUpdateBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcInheritedDDNSUpdateBlockModel) Flatten(ctx context.Context, from *ipam.InheritedDDNSUpdateBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcInheritedDDNSUpdateBlockModel{}
	}
	m.Action = flex.FlattenStringPointer(from.Action)
	m.DisplayName = flex.FlattenStringPointer(from.DisplayName)
	m.Source = flex.FlattenStringPointer(from.Source)
	m.Value = FlattenIpamsvcDDNSUpdateBlock(ctx, from.Value, diags)
}
