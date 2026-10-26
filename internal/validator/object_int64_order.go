package validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ validator.Object = objectInt64OrderValidator{}

// objectInt64OrderValidator validates that, for an object with int64 attributes,
// the values satisfy min <= default <= max. Null and unknown values are skipped.
type objectInt64OrderValidator struct {
	minAttr     string
	defaultAttr string
	maxAttr     string
}

func (v objectInt64OrderValidator) Description(_ context.Context) string {
	return fmt.Sprintf("%s must be <= %s, which must be <= %s", v.minAttr, v.defaultAttr, v.maxAttr)
}

func (v objectInt64OrderValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v objectInt64OrderValidator) ValidateObject(_ context.Context, req validator.ObjectRequest, resp *validator.ObjectResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	attrs := req.ConfigValue.Attributes()
	minV, minOk := knownInt64(attrs[v.minAttr])
	defV, defOk := knownInt64(attrs[v.defaultAttr])
	maxV, maxOk := knownInt64(attrs[v.maxAttr])

	check := func(lowName string, low int64, lowOk bool, highName string, high int64, highOk bool) {
		if lowOk && highOk && low > high {
			resp.Diagnostics.AddAttributeError(req.Path, "Invalid value range",
				fmt.Sprintf("%s (%d) must be less than or equal to %s (%d).", lowName, low, highName, high))
		}
	}
	check(v.minAttr, minV, minOk, v.maxAttr, maxV, maxOk)
	check(v.minAttr, minV, minOk, v.defaultAttr, defV, defOk)
	check(v.defaultAttr, defV, defOk, v.maxAttr, maxV, maxOk)
}

func knownInt64(a attr.Value) (int64, bool) {
	i, ok := a.(basetypes.Int64Value)
	if !ok || i.IsNull() || i.IsUnknown() {
		return 0, false
	}
	return i.ValueInt64(), true
}

// ObjectInt64Order returns a validator ensuring minAttr <= defaultAttr <= maxAttr
// for the int64 attributes of an object.
func ObjectInt64Order(minAttr, defaultAttr, maxAttr string) validator.Object {
	return objectInt64OrderValidator{minAttr: minAttr, defaultAttr: defaultAttr, maxAttr: maxAttr}
}
