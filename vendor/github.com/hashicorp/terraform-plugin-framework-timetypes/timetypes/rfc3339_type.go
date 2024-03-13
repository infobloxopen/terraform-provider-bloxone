// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timetypes

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	_ basetypes.StringTypable = (*RFC3339Type)(nil)
	_ xattr.TypeWithValidate  = (*RFC3339Type)(nil)
)

// RFC3339Type is an attribute type that represents a valid RFC 3339 string. Semantic equality logic is defined
// for RFC3339Type such that inconsequential differences between the `Z` suffix and a `00:00` UTC offset are ignored.
type RFC3339Type struct {
	basetypes.StringType
}

// String returns a human-readable string of the type name.
func (t RFC3339Type) String() string {
	return "timetypes.RFC3339Type"
}

// ValueType returns the Value type.
func (t RFC3339Type) ValueType(ctx context.Context) attr.Value {
	return RFC3339{}
}

// Equal returns true if the given type is equivalent.
func (t RFC3339Type) Equal(o attr.Type) bool {
	other, ok := o.(RFC3339Type)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

// Validate implements type validation. This type requires the value provided to be a String value that is
// valid RFC 3339 format. This utilizes the Go `time` library which does not strictly adhere to the RFC 3339
// standard and may allow strings that are not valid RFC 3339 strings
//
// See https://github.com/golang/go/issues/54580 for more info on the Go `time` library's RFC 3339 parsing differences.
func (t RFC3339Type) Validate(ctx context.Context, in tftypes.Value, path path.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	if in.Type() == nil {
		return diags
	}

	if !in.Type().Is(tftypes.String) {
		err := fmt.Errorf("expected String value, received %T with value: %v", in, in)
		diags.AddAttributeError(
			path,
			"RFC3339 Time Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)
		return diags
	}

	if !in.IsKnown() || in.IsNull() {
		return diags
	}

	var valueString string

	if err := in.As(&valueString); err != nil {
		diags.AddAttributeError(
			path,
			"RFC3339 Time Type Validation Error",
			"An unexpected error was encountered trying to validate an attribute value. This is always an error in the provider. "+
				"Please report the following to the provider developer:\n\n"+err.Error(),
		)

		return diags
	}

	if _, err := time.Parse(time.RFC3339, valueString); err != nil {
		diags.Append(diag.WithPath(path, rfc3339InvalidStringDiagnostic(valueString, err)))

		return diags
	}

	return diags
}

// ValueFromString returns a StringValuable type given a StringValue.
func (t RFC3339Type) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return RFC3339{
		StringValue: in,
	}, nil
}

// ValueFromTerraform returns a Value given a tftypes.Value.  This is meant to convert the tftypes.Value into a more convenient Go type
// for the provider to consume the data with.
func (t RFC3339Type) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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
