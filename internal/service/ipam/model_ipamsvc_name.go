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

type IpamsvcNameModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

var IpamsvcNameAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"type": types.StringType,
}

var IpamsvcNameResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name expressed as a single label or FQDN.",
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The origin of the name.",
	},
}

func ExpandIpamsvcName(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcName {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcNameModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcNameModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcName {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcName{
		Name: flex.ExpandString(m.Name),
		Type: flex.ExpandString(m.Type),
	}
	return to
}

func FlattenIpamsvcName(ctx context.Context, from *ipam.IpamsvcName, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcNameAttrTypes)
	}
	m := IpamsvcNameModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcNameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcNameModel) Flatten(ctx context.Context, from *ipam.IpamsvcName, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcNameModel{}
	}
	m.Name = flex.FlattenString(from.Name)
	m.Type = flex.FlattenString(from.Type)
}
