package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcdfpResolverModel struct {
	Address    types.String `tfsdk:"address"`
	IsFallback types.Bool   `tfsdk:"is_fallback"`
	IsLocal    types.Bool   `tfsdk:"is_local"`
	Protocols  types.List   `tfsdk:"protocols"`
}

var AtcdfpResolverAttrTypes = map[string]attr.Type{
	"address":     types.StringType,
	"is_fallback": types.BoolType,
	"is_local":    types.BoolType,
	"protocols":   types.ListType{ElemType: types.ObjectType{AttrTypes: AtcdfpDNSProtocolAttrTypes}},
}

var AtcdfpResolverResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "address that can be used as resolver",
	},
	"is_fallback": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Mark it true to set default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable.",
	},
	"is_local": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Mark it true to set internal or local DNS servers' IPv4 or IPv6 addresses that are used as DNS resolvers",
	},
	"protocols": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of DNS resolver communication protocols.",
	},
}

func ExpandAtcdfpResolver(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpResolver {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpResolverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpResolverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpResolver {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpResolver{
		Address:    flex.ExpandStringPointer(m.Address),
		IsFallback: flex.ExpandBoolPointer(m.IsFallback),
		IsLocal:    flex.ExpandBoolPointer(m.IsLocal),
		Protocols:  flex.ExpandFrameworkListString(ctx, m.Protocols, diags),
	}
	return to
}

func FlattenAtcdfpResolver(ctx context.Context, from *dfp.AtcdfpResolver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpResolverAttrTypes)
	}
	m := AtcdfpResolverModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpResolverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpResolverModel) Flatten(ctx context.Context, from *dfp.AtcdfpResolver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpResolverModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.IsFallback = types.BoolPointerValue(from.IsFallback)
	m.IsLocal = types.BoolPointerValue(from.IsLocal)
	m.Protocols = flex.FlattenFrameworkListString(ctx, from.Protocols, diags)
}
