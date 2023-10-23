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

type IpamsvcAccessFilterModel struct {
	Access           types.String `tfsdk:"access"`
	HardwareFilterId types.String `tfsdk:"hardware_filter_id"`
	OptionFilterId   types.String `tfsdk:"option_filter_id"`
}

var IpamsvcAccessFilterAttrTypes = map[string]attr.Type{
	"access":             types.StringType,
	"hardware_filter_id": types.StringType,
	"option_filter_id":   types.StringType,
}

var IpamsvcAccessFilterResourceSchema = schema.Schema{
	MarkdownDescription: `The __AccessFilter__ object represents an allow/deny filter for a DHCP range.`,
	Attributes:          IpamsvcAccessFilterResourceSchemaAttributes,
}

var IpamsvcAccessFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_.`,
	},
	"hardware_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"option_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcAccessFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAccessFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcAccessFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcAccessFilterModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAccessFilter {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcAccessFilter{
		Access:           m.Access.ValueString(),
		HardwareFilterId: m.HardwareFilterId.ValueStringPointer(),
		OptionFilterId:   m.OptionFilterId.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcAccessFilter(ctx context.Context, from *ipam.IpamsvcAccessFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAccessFilterAttrTypes)
	}
	m := IpamsvcAccessFilterModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAccessFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAccessFilterModel) flatten(ctx context.Context, from *ipam.IpamsvcAccessFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAccessFilterModel{}
	}

	m.Access = types.StringValue(from.Access)
	m.HardwareFilterId = types.StringPointerValue(from.HardwareFilterId)
	m.OptionFilterId = types.StringPointerValue(from.OptionFilterId)

}
