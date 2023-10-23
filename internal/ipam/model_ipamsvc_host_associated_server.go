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

type IpamsvcHostAssociatedServerModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

var IpamsvcHostAssociatedServerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"name": types.StringType,
}

var IpamsvcHostAssociatedServerResourceSchema = schema.Schema{
	MarkdownDescription: ``,
	Attributes:          IpamsvcHostAssociatedServerResourceSchemaAttributes,
}

var IpamsvcHostAssociatedServerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `The DHCP Config Profile name.`,
	},
}

func expandIpamsvcHostAssociatedServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcHostAssociatedServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}

	var m IpamsvcHostAssociatedServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.expand(ctx, diags)
}

func (m *IpamsvcHostAssociatedServerModel) expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcHostAssociatedServer {
	if m == nil {
		return nil
	}

	to := &ipam.IpamsvcHostAssociatedServer{}
	return to
}

func flattenIpamsvcHostAssociatedServer(ctx context.Context, from *ipam.IpamsvcHostAssociatedServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcHostAssociatedServerAttrTypes)
	}
	m := IpamsvcHostAssociatedServerModel{}
	m.flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcHostAssociatedServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcHostAssociatedServerModel) flatten(ctx context.Context, from *ipam.IpamsvcHostAssociatedServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcHostAssociatedServerModel{}
	}

	m.Id = types.StringPointerValue(from.Id)
	m.Name = types.StringPointerValue(from.Name)

}
