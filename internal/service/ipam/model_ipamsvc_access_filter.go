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

var IpamsvcAccessFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The access type of DHCP filter (_allow_ or _deny_).  Defaults to _allow_.",
	},
	"hardware_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"option_filter_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandIpamsvcAccessFilter(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcAccessFilter {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcAccessFilterModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcAccessFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcAccessFilter {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcAccessFilter{
		Access:           flex.ExpandString(m.Access),
		HardwareFilterId: flex.ExpandStringPointer(m.HardwareFilterId),
		OptionFilterId:   flex.ExpandStringPointer(m.OptionFilterId),
	}
	return to
}

func FlattenIpamsvcAccessFilter(ctx context.Context, from *ipam.IpamsvcAccessFilter, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcAccessFilterAttrTypes)
	}
	m := IpamsvcAccessFilterModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcAccessFilterAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcAccessFilterModel) Flatten(ctx context.Context, from *ipam.IpamsvcAccessFilter, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcAccessFilterModel{}
	}
	m.Access = flex.FlattenString(from.Access)
	m.HardwareFilterId = flex.FlattenStringPointer(from.HardwareFilterId)
	m.OptionFilterId = flex.FlattenStringPointer(from.OptionFilterId)
}
