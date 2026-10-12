package types

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable                    = IPAddressType{}
	_ basetypes.StringValuable                   = IPAddressValue{}
	_ basetypes.StringValuableWithSemanticEquals = IPAddressValue{}
)

// IPAddressType is a custom string type for a plain IPv4 or IPv6 address
// (no prefix length). Values are compared semantically, so differences in
// IPv6 letter case or zero compression do not cause drift.
type IPAddressType struct {
	basetypes.StringType
}

func (t IPAddressType) String() string {
	return "IPAddressType"
}

func (t IPAddressType) ValueType(_ context.Context) attr.Value {
	return IPAddressValue{}
}

func (t IPAddressType) Equal(o attr.Type) bool {
	other, ok := o.(IPAddressType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t IPAddressType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return IPAddressValue{StringValue: in}, nil
}

func (t IPAddressType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)
	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

// IPAddressValue is the value type for IPAddressType.
type IPAddressValue struct {
	basetypes.StringValue
}

func (v IPAddressValue) Type(_ context.Context) attr.Type {
	return IPAddressType{}
}

func (v IPAddressValue) Equal(o attr.Value) bool {
	other, ok := o.(IPAddressValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

// StringSemanticEquals reports whether both values parse to the same IP address.
func (v IPAddressValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(IPAddressValue)
	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)
		return false, diags
	}

	current, err := netip.ParseAddr(v.ValueString())
	if err != nil {
		return false, diags
	}
	updated, err := netip.ParseAddr(newValue.ValueString())
	if err != nil {
		return false, diags
	}

	return current == updated, diags
}

func NewIPAddressValue(value string) IPAddressValue {
	return IPAddressValue{StringValue: basetypes.NewStringValue(value)}
}

func NewIPAddressNull() IPAddressValue {
	return IPAddressValue{StringValue: basetypes.NewStringNull()}
}
