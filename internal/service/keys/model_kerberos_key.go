package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/keys"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type KerberosKeyModel struct {
	Algorithm  types.String `tfsdk:"algorithm"`
	Comment    types.String `tfsdk:"comment"`
	Domain     types.String `tfsdk:"domain"`
	Id         types.String `tfsdk:"id"`
	Principal  types.String `tfsdk:"principal"`
	Tags       types.Map    `tfsdk:"tags"`
	UploadedAt types.String `tfsdk:"uploaded_at"`
	Version    types.Int64  `tfsdk:"version"`
}

var KerberosKeyAttrTypes = map[string]attr.Type{
	"algorithm":   types.StringType,
	"comment":     types.StringType,
	"domain":      types.StringType,
	"id":          types.StringType,
	"principal":   types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"uploaded_at": types.StringType,
	"version":     types.Int64Type,
}

var KerberosKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Encryption algorithm of the key in accordance with RFC 3961.",
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description for Kerberos key. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"domain": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Kerberos realm of the principal.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"principal": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Kerberos principal associated with key.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the Kerberos key in JSON format.",
	},
	"uploaded_at": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Upload time for the key.",
	},
	"version": schema.Int64Attribute{
		Computed:            true,
		MarkdownDescription: "The version number (KVNO) of the key.",
	},
}

func ExpandKerberosKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *keys.KerberosKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KerberosKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KerberosKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *keys.KerberosKey {
	if m == nil {
		return nil
	}
	to := &keys.KerberosKey{
		Comment: flex.ExpandStringPointer(m.Comment),
		Tags:    flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenKerberosKey(ctx context.Context, from *keys.KerberosKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KerberosKeyAttrTypes)
	}
	m := KerberosKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KerberosKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KerberosKeyModel) Flatten(ctx context.Context, from *keys.KerberosKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KerberosKeyModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Domain = flex.FlattenStringPointer(from.Domain)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Principal = flex.FlattenStringPointer(from.Principal)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UploadedAt = flex.FlattenStringPointer(from.UploadedAt)
	m.Version = flex.FlattenInt64Pointer(from.Version)
}
