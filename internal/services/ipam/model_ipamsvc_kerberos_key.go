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

type IpamsvcKerberosKeyModel struct {
	Algorithm  types.String `tfsdk:"algorithm"`
	Domain     types.String `tfsdk:"domain"`
	Key        types.String `tfsdk:"key"`
	Principal  types.String `tfsdk:"principal"`
	UploadedAt types.String `tfsdk:"uploaded_at"`
	Version    types.Int64  `tfsdk:"version"`
}

var IpamsvcKerberosKeyAttrTypes = map[string]attr.Type{
	"algorithm":   types.StringType,
	"domain":      types.StringType,
	"key":         types.StringType,
	"principal":   types.StringType,
	"uploaded_at": types.StringType,
	"version":     types.Int64Type,
}

var IpamsvcKerberosKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Encryption algorithm of the key in accordance with RFC 3961.`,
	},
	"domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Kerberos realm of the principal.`,
	},
	"key": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"principal": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Kerberos principal associated with key.`,
	},
	"uploaded_at": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Upload time for the key.`,
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: `The version number (KVNO) of the key.`,
	},
}

func ExpandIpamsvcKerberosKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *ipam.IpamsvcKerberosKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IpamsvcKerberosKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *IpamsvcKerberosKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.IpamsvcKerberosKey {
	if m == nil {
		return nil
	}
	to := &ipam.IpamsvcKerberosKey{
		Key: m.Key.ValueString(),
	}
	return to
}

func FlattenIpamsvcKerberosKey(ctx context.Context, from *ipam.IpamsvcKerberosKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IpamsvcKerberosKeyAttrTypes)
	}
	m := IpamsvcKerberosKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IpamsvcKerberosKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *IpamsvcKerberosKeyModel) Flatten(ctx context.Context, from *ipam.IpamsvcKerberosKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = IpamsvcKerberosKeyModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Key = flex.FlattenString(from.Key)
	m.Principal = flex.FlattenStringPointer(from.Principal)
	m.UploadedAt = flex.FlattenStringPointer(from.UploadedAt)
	m.Version = flex.FlattenInt64(int64(*from.Version))

}
