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

type IpamsvcLeaseAddressModel struct {
	Address types.String `tfsdk:"address"`
	Space   types.String `tfsdk:"space"`
}

var IpamsvcLeaseAddressAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"space":   types.StringType,
}

var IpamsvcLeaseAddressResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcLeaseAddressResourceSchemaAttributes,
}

var IpamsvcLeaseAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The IP address for the DHCP lease in IPv4 or IPv6 format within the IP space specified by _space_ field.`,
	},
	"space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func expandIpamsvcLeaseAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcLeaseAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcLeaseAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcLeaseAddressModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcLeaseAddress {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcLeaseAddress{
		Address: m.Address.ValueStringPointer(),
		Space:   m.Space.ValueStringPointer(),
	}
	return to
}

func flattenIpamsvcLeaseAddress(ctx context.Context, from *ipam.IpamsvcLeaseAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcLeaseAddressAttrTypes)
	}
	m := IpamsvcLeaseAddressModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcLeaseAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcLeaseAddressModel) flatten(ctx context.Context, from *ipam.IpamsvcLeaseAddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcLeaseAddressModel{}
	}

	m.Address = types.StringPointerValue(from.Address)
	m.Space = types.StringPointerValue(from.Space)

}
