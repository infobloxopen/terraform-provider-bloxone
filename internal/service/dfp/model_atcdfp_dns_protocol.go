package dfp

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/bloxone-go-client/dfp"
)

type AtcdfpDNSProtocolModel struct {
}

var AtcdfpDNSProtocolAttrTypes = map[string]attr.Type{}

var AtcdfpDNSProtocolResourceSchemaAttributes = map[string]schema.Attribute{}

func ExpandAtcdfpDNSProtocol(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.AtcdfpDNSProtocol {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDNSProtocolModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDNSProtocolModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.AtcdfpDNSProtocol {
	if m == nil {
		return nil
	}
	to := &dfp.AtcdfpDNSProtocol{}
	return to
}

func FlattenAtcdfpDNSProtocol(ctx context.Context, from *dfp.AtcdfpDNSProtocol, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDNSProtocolAttrTypes)
	}
	m := AtcdfpDNSProtocolModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDNSProtocolAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDNSProtocolModel) Flatten(ctx context.Context, from *dfp.AtcdfpDNSProtocol, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDNSProtocolModel{}
	}
}
