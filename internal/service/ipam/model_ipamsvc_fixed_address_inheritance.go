package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcFixedAddressInheritanceModel struct {
	DhcpOptions               types.Object `tfsdk:"dhcp_options"`
	HeaderOptionFilename      types.Object `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress types.Object `tfsdk:"header_option_server_address"`
	HeaderOptionServerName    types.Object `tfsdk:"header_option_server_name"`
}

var IpamsvcFixedAddressInheritanceAttrTypes = map[string]attr.Type{
	"dhcp_options":                 types.ObjectType{AttrTypes: IpamsvcInheritedDHCPOptionListAttrTypes},
	"header_option_filename":       types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_address": types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_name":    types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
}

var IpamsvcFixedAddressInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes: IpamsvcInheritedDHCPOptionListResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_filename": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_server_address": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
	"header_option_server_name": schema.SingleNestedAttribute{
		Attributes: InheritanceInheritedStringResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandIpamsvcFixedAddressInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcFixedAddressInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcFixedAddressInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcFixedAddressInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcFixedAddressInheritance {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcFixedAddressInheritance{
		DhcpOptions:               ExpandIpamsvcInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
		HeaderOptionFilename:      ExpandInheritanceInheritedString(ctx, m.HeaderOptionFilename, diags),
		HeaderOptionServerAddress: ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerAddress, diags),
		HeaderOptionServerName:    ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerName, diags),
	}
	return to
}

func FlattenIpamsvcFixedAddressInheritance(ctx context.Context, from *ipam.IpamsvcFixedAddressInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcFixedAddressInheritanceAttrTypes)
	}
	m := IpamsvcFixedAddressInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcFixedAddressInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcFixedAddressInheritanceModel) Flatten(ctx context.Context, from *ipam.IpamsvcFixedAddressInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcFixedAddressInheritanceModel{}
	}
	m.DhcpOptions = FlattenIpamsvcInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
	m.HeaderOptionFilename = FlattenInheritanceInheritedString(ctx, from.HeaderOptionFilename, diags)
	m.HeaderOptionServerAddress = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerAddress, diags)
	m.HeaderOptionServerName = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerName, diags)
}
