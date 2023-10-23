package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcAsmEnableBlockModel struct {
	Enable             types.Bool        `tfsdk:"enable"`
	EnableNotification types.Bool        `tfsdk:"enable_notification"`
	ReenableDate       timetypes.RFC3339 `tfsdk:"reenable_date"`
}

var IpamsvcAsmEnableBlockAttrTypes = map[string]attr.Type{
	"enable":              types.BoolType,
	"enable_notification": types.BoolType,
	"reenable_date":       timetypes.RFC3339Type{},
}

var IpamsvcAsmEnableBlockResourceSchema = schema.Schema{
	MarkdownDescription: `ASM enable group of fields.`,
	Attributes:          IpamsvcAsmEnableBlockResourceSchemaAttributes,
}

var IpamsvcAsmEnableBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether Automated Scope Management is enabled or not.`,
	},
	"enable_notification": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Indicates whether sending notifications to the users is enabled or not.`,
	},
	"reenable_date": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Optional:            true,
		MarkdownDescription: `The date at which notifications will be re-enabled automatically.`,
	},
}

func expandIpamsvcAsmEnableBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAsmEnableBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcAsmEnableBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcAsmEnableBlockModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAsmEnableBlock {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcAsmEnableBlock{
		Enable:             m.Enable.ValueBoolPointer(),
		EnableNotification: m.EnableNotification.ValueBoolPointer(),
		ReenableDate:       ptr(ExpandTime(ctx, m.ReenableDate, diags)),
	}
	return to
}

func flattenIpamsvcAsmEnableBlock(ctx context.Context, from *ipam.IpamsvcAsmEnableBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAsmEnableBlockAttrTypes)
	}
	m := IpamsvcAsmEnableBlockModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAsmEnableBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAsmEnableBlockModel) flatten(ctx context.Context, from *ipam.IpamsvcAsmEnableBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAsmEnableBlockModel{}
	}

	m.Enable = types.BoolPointerValue(from.Enable)
	m.EnableNotification = types.BoolPointerValue(from.EnableNotification)
	m.ReenableDate = timetypes.NewRFC3339TimePointerValue(from.ReenableDate)

}
