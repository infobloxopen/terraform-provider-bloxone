package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/keys"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type KeysTSIGKeyModel struct {
	Algorithm    types.String      `tfsdk:"algorithm"`
	Comment      types.String      `tfsdk:"comment"`
	CreatedAt    timetypes.RFC3339 `tfsdk:"created_at"`
	Id           types.String      `tfsdk:"id"`
	Name         types.String      `tfsdk:"name"`
	ProtocolName types.String      `tfsdk:"protocol_name"`
	Secret       types.String      `tfsdk:"secret"`
	Tags         types.Map         `tfsdk:"tags"`
	UpdatedAt    timetypes.RFC3339 `tfsdk:"updated_at"`
}

var KeysTSIGKeyAttrTypes = map[string]attr.Type{
	"algorithm":     types.StringType,
	"comment":       types.StringType,
	"created_at":    timetypes.RFC3339Type{},
	"id":            types.StringType,
	"name":          types.StringType,
	"protocol_name": types.StringType,
	"secret":        types.StringType,
	"tags":          types.MapType{ElemType: types.StringType},
	"updated_at":    timetypes.RFC3339Type{},
}

var KeysTSIGKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("hmac_sha256"),
		MarkdownDescription: "TSIG key algorithm.\n\n" +
			"  Valid values are:\n" +
			"  * _hmac_sha1_\n" +
			"  * _hmac_sha224_\n" +
			"  * _hmac_sha256_\n" +
			"  * _hmac_sha384_\n" +
			"  * _hmac_sha512_\n\n" +
			"  Defaults to _hmac_sha256_.",
		Validators: []validator.String{
			stringvalidator.OneOf("hmac_sha1", "hmac_sha224", "hmac_sha256", "hmac_sha384", "hmac_sha512"),
		},
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The TSIG key name in the absolute domain name format.",
	},
	"protocol_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The TSIG key name supplied during a create/update operation that is converted to canonical form in punycode.",
	},
	"secret": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Sensitive:           true,
		MarkdownDescription: "The TSIG key secret as a Base64 encoded string.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the TSIG key in JSON format.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandKeysTSIGKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *keys.TSIGKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KeysTSIGKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KeysTSIGKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *keys.TSIGKey {
	if m == nil {
		return nil
	}
	to := &keys.TSIGKey{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Name:      flex.ExpandString(m.Name),
		Secret:    flex.ExpandString(m.Secret),
		Tags:      flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenKeysTSIGKey(ctx context.Context, from *keys.TSIGKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KeysTSIGKeyAttrTypes)
	}
	m := KeysTSIGKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KeysTSIGKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KeysTSIGKeyModel) Flatten(ctx context.Context, from *keys.TSIGKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KeysTSIGKeyModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenString(from.Name)
	m.ProtocolName = flex.FlattenStringPointer(from.ProtocolName)
	m.Secret = flex.FlattenString(from.Secret)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
