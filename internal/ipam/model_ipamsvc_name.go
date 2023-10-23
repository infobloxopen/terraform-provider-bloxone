package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/ipam"
)

type IpamsvcNameModel struct {
	Name types.String `tfsdk:"name"`
	Type types.String `tfsdk:"type"`
}

var IpamsvcNameAttrTypes = map[string]attr.Type{
	"name": types.StringType,
	"type": types.StringType,
}

var IpamsvcNameResourceSchema = schema.Schema{
	MarkdownDescription: `The __Name__ object represents a name associated with an address.`,
	Attributes:          IpamsvcNameResourceSchemaAttributes,
}

var IpamsvcNameResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The name expressed as a single label or FQDN.`,
	},
	"type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The origin of the name.`,
	},
}

func expandIpamsvcName(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcName {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcNameModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcNameModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcName {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcName{
		Name: m.Name.ValueString(),
		Type: m.Type.ValueString(),
	}
	return to
}

func flattenIpamsvcName(ctx context.Context, from *ipam.IpamsvcName, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcNameAttrTypes)
	}
	m := IpamsvcNameModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcNameAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcNameModel) flatten(ctx context.Context, from *ipam.IpamsvcName, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcNameModel{}
	}

	m.Name = types.StringValue(from.Name)
	m.Type = types.StringValue(from.Type)

}
