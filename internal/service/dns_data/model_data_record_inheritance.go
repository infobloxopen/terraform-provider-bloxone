package dns_data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
)

type DataRecordInheritanceModel struct {
	Ttl types.Object `tfsdk:"ttl"`
}

var DataRecordInheritanceAttrTypes = map[string]attr.Type{
	"ttl": types.ObjectType{AttrTypes: Inheritance2InheritedUInt32AttrTypes},
}

var DataRecordInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"ttl": schema.SingleNestedAttribute{
		Attributes: Inheritance2InheritedUInt32ResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandDataRecordInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_data.DataRecordInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DataRecordInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DataRecordInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_data.DataRecordInheritance {
	if m == nil {
		return nil
	}
	to := &dns_data.DataRecordInheritance{
		Ttl: ExpandInheritance2InheritedUInt32(ctx, m.Ttl, diags),
	}
	return to
}

func FlattenDataRecordInheritance(ctx context.Context, from *dns_data.DataRecordInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DataRecordInheritanceAttrTypes)
	}
	m := DataRecordInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DataRecordInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DataRecordInheritanceModel) Flatten(ctx context.Context, from *dns_data.DataRecordInheritance, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DataRecordInheritanceModel{}
	}
	m.Ttl = FlattenInheritance2InheritedUInt32(ctx, from.Ttl, diags)
}
