package types

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"golang.org/x/exp/slices"
)

var UnorderedListOfStringType = UnorderedList[basetypes.StringValue]{basetypes.ListType{ElemType: basetypes.StringType{}}}

var _ basetypes.ListValuableWithSemanticEquals = UnorderedListValue[basetypes.StringValue]{}
var _ basetypes.ListTypable = UnorderedList[basetypes.StringValue]{}

// UnorderedList is a type representing a list of values where the order of the elements is not significant.
type UnorderedList[T attr.Value] struct {
	basetypes.ListType
}

func NewUnorderedList[T attr.Value](ctx context.Context) UnorderedList[T] {
	return UnorderedList[T]{basetypes.ListType{ElemType: newAttrTypeOf[T](ctx)}}
}

func (t UnorderedList[T]) Equal(o attr.Type) bool {
	other, ok := o.(UnorderedList[T])

	if !ok {
		return false
	}

	return t.ListType.Equal(other.ListType)
}

func (UnorderedList[T]) String() string {
	return fmt.Sprintf("UnorderedList[%T]", new(T))
}

func (t UnorderedList[T]) ValueFromList(ctx context.Context, in basetypes.ListValue) (basetypes.ListValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	if in.IsNull() {
		return NewUnorderedListValueNull[T](ctx), diags
	}

	if in.IsUnknown() {
		return NewUnorderedListValueUnknown[T](ctx), diags
	}

	v, d := basetypes.NewListValue(newAttrTypeOf[T](ctx), in.Elements())
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown[T](ctx), diags
	}

	return UnorderedListValue[T]{ListValue: v}, diags
}

func (t UnorderedList[T]) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.ListType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	listValue, ok := attrValue.(basetypes.ListValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	listValuable, diags := t.ValueFromList(ctx, listValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting ListValue to ListValuable: %v", diags)
	}

	return listValuable, nil
}

func (UnorderedList[T]) ValueType(context.Context) attr.Value {
	return UnorderedListValue[T]{}
}

type UnorderedListValue[T attr.Value] struct {
	basetypes.ListValue
}

func (v UnorderedListValue[T]) ListSemanticEquals(ctx context.Context, newValuable basetypes.ListValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(UnorderedListValue[T])
	if !ok {
		return false, diags
	}

	o, d := v.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	n, d := newValue.ToListValue(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return false, diags
	}

	oldElems, newElems := o.Elements(), n.Elements()

	if len(oldElems) != len(newElems) {
		return false, diags
	}

	for _, newElem := range newElems {
		found := false
		for i, oldElem := range oldElems {
			if oldElem.Equal(newElem) {
				oldElems = slices.Delete(oldElems, i, i+1)
				found = true
				break
			}
		}
		if !found {
			return false, diags
		}
	}

	return len(oldElems) == 0, diags
}

func (v UnorderedListValue[T]) Equal(o attr.Value) bool {
	other, ok := o.(UnorderedListValue[T])

	if !ok {
		return false
	}

	return v.ListValue.Equal(other.ListValue)
}

func (UnorderedListValue[T]) Type(ctx context.Context) attr.Type {
	return NewUnorderedList[T](ctx)
}

func NewUnorderedListValueNull[T attr.Value](ctx context.Context) UnorderedListValue[T] {
	return UnorderedListValue[T]{ListValue: basetypes.NewListNull(newAttrTypeOf[T](ctx))}
}

func NewUnorderedListValueUnknown[T attr.Value](ctx context.Context) UnorderedListValue[T] {
	return UnorderedListValue[T]{ListValue: basetypes.NewListUnknown(newAttrTypeOf[T](ctx))}
}

func NewUnorderedListValue[T attr.Value](ctx context.Context, elements []attr.Value) (UnorderedListValue[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValue(newAttrTypeOf[T](ctx), elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown[T](ctx), diags
	}

	return UnorderedListValue[T]{ListValue: v}, diags
}

func NewUnorderedListValueFrom[T attr.Value](ctx context.Context, elements any) (UnorderedListValue[T], diag.Diagnostics) {
	var diags diag.Diagnostics

	v, d := basetypes.NewListValueFrom(ctx, newAttrTypeOf[T](ctx), elements)
	diags.Append(d...)
	if diags.HasError() {
		return NewUnorderedListValueUnknown[T](ctx), diags
	}

	return UnorderedListValue[T]{ListValue: v}, diags
}

func newAttrTypeOf[T attr.Value](ctx context.Context) attr.Type {
	return (*new(T)).Type(ctx)
}
