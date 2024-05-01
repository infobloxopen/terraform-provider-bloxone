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

type ResolverModel struct {
	Address    types.String `tfsdk:"address"`
	IsFallback types.Bool   `tfsdk:"is_fallback"`
	IsLocal    types.Bool   `tfsdk:"is_local"`
	Protocols  types.List   `tfsdk:"protocols"`
}

var ResolverAttrTypes = map[string]attr.Type{
	"address":     types.StringType,
	"is_fallback": types.BoolType,
	"is_local":    types.BoolType,
	"protocols":   types.ListType{ElemType: types.StringType},
}

var ResolverResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Address that can be used as resolver",
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

func ExpandResolver(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.Resolver {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ResolverModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ResolverModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.Resolver {
	if m == nil {
		return nil
	}
	to := &dfp.Resolver{
		Address:    flex.ExpandStringPointer(m.Address),
		IsFallback: flex.ExpandBoolPointer(m.IsFallback),
		IsLocal:    flex.ExpandBoolPointer(m.IsLocal),
		Protocols:  ExpandDNSProtocol(ctx, m.Protocols, diags),
	}
	return to
}

func ExpandDNSProtocol(ctx context.Context, tfList types.List, diags *diag.Diagnostics) []dfp.DNSProtocol {
	if tfList.IsNull() || tfList.IsUnknown() {
		return nil
	}
	var data []dfp.DNSProtocol
	diags.Append(tfList.ElementsAs(ctx, &data, false)...)
	return data
}

func FlattenResolver(ctx context.Context, from *dfp.Resolver, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ResolverAttrTypes)
	}
	m := ResolverModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ResolverAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ResolverModel) Flatten(ctx context.Context, from *dfp.Resolver, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ResolverModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.IsFallback = types.BoolPointerValue(from.IsFallback)
	m.IsLocal = types.BoolPointerValue(from.IsLocal)
	m.Protocols = FlattenDNSProtocol(ctx, from.Protocols, diags)
}

func FlattenDNSProtocol(ctx context.Context, l []dfp.DNSProtocol, diags *diag.Diagnostics) types.List {
	if len(l) == 0 {
		return types.ListNull(types.StringType)
	}
	tfList, d := types.ListValueFrom(ctx, types.StringType, l)
	diags.Append(d...)
	return tfList
}
