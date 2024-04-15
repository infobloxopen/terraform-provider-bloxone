package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/anycast"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ProtobufFieldMaskModel struct {
	Paths types.List `tfsdk:"paths"`
}

var ProtobufFieldMaskAttrTypes = map[string]attr.Type{
	"paths": types.ListType{ElemType: types.StringType},
}

var ProtobufFieldMaskResourceSchemaAttributes = map[string]schema.Attribute{
	"paths": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The set of field mask paths.",
	},
}

func ExpandProtobufFieldMask(ctx context.Context, o types.Object, diags *diag.Diagnostics) *anycast.ProtobufFieldMask {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ProtobufFieldMaskModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ProtobufFieldMaskModel) Expand(ctx context.Context, diags *diag.Diagnostics) *anycast.ProtobufFieldMask {
	if m == nil {
		return nil
	}
	to := &anycast.ProtobufFieldMask{
		Paths: flex.ExpandFrameworkListString(ctx, m.Paths, diags),
	}
	return to
}

func FlattenProtobufFieldMask(ctx context.Context, from *anycast.ProtobufFieldMask, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ProtobufFieldMaskAttrTypes)
	}
	m := ProtobufFieldMaskModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ProtobufFieldMaskAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ProtobufFieldMaskModel) Flatten(ctx context.Context, from *anycast.ProtobufFieldMask, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ProtobufFieldMaskModel{}
	}
	m.Paths = flex.FlattenFrameworkListString(ctx, from.Paths, diags)
}
