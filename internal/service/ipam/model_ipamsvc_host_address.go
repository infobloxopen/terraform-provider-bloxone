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

type IpamsvcHostAddressModel struct {
	Address types.String `tfsdk:"address"`
	Ref     types.String `tfsdk:"ref"`
	Space   types.String `tfsdk:"space"`
}

var IpamsvcHostAddressAttrTypes = map[string]attr.Type{
	"address": types.StringType,
	"ref":     types.StringType,
	"space":   types.StringType,
}

var IpamsvcHostAddressResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.`,
	},
	"ref": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"space": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
}

func ExpandIpamsvcHostAddress(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHostAddress {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcHostAddressModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcHostAddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHostAddress {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcHostAddress{
		Address: flex.ExpandString(m.Address),
		Ref:     flex.ExpandString(m.Ref),
		Space:   flex.ExpandString(m.Space),
	}
	return to
}

func FlattenIpamsvcHostAddress(ctx context.Context, from *ipam.IpamsvcHostAddress, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHostAddressAttrTypes)
	}
	m := IpamsvcHostAddressModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHostAddressAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHostAddressModel) Flatten(ctx context.Context, from *ipam.IpamsvcHostAddress, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostAddressModel{}
	}
	m.Address = flex.FlattenString(from.Address)
	m.Ref = flex.FlattenString(from.Ref)
	m.Space = flex.FlattenString(from.Space)

}
