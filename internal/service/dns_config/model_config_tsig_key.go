package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigTSIGKeyModel struct {
	Algorithm    types.String `tfsdk:"algorithm"`
	Comment      types.String `tfsdk:"comment"`
	Key          types.String `tfsdk:"key"`
	Name         types.String `tfsdk:"name"`
	ProtocolName types.String `tfsdk:"protocol_name"`
	Secret       types.String `tfsdk:"secret"`
}

var ConfigTSIGKeyAttrTypes = map[string]attr.Type{
	"algorithm":     types.StringType,
	"comment":       types.StringType,
	"key":           types.StringType,
	"name":          types.StringType,
	"protocol_name": types.StringType,
	"secret":        types.StringType,
}

var ConfigTSIGKeyResourceSchemaAttributes = map[string]schema.Attribute{
	"algorithm": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `TSIG key algorithm.  Possible values:  * _hmac_sha256_,  * _hmac_sha1_,  * _hmac_sha224_,  * _hmac_sha384_,  * _hmac_sha512_.`,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("name"),
				path.MatchRelative().AtParent().AtName("secret")),
		},
	},
	"comment": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `Comment for TSIG key.`,
	},
	"key": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.ConflictsWith(
				path.MatchRelative().AtParent().AtName("name"),
				path.MatchRelative().AtParent().AtName("algorithm"),
				path.MatchRelative().AtParent().AtName("secret")),
		},
		MarkdownDescription: `The resource identifier.`,
	},
	"name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: `TSIG key name, FQDN.`,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("algorithm"),
				path.MatchRelative().AtParent().AtName("secret")),
		},
	},
	"protocol_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `TSIG key name in punycode.`,
	},
	"secret": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Sensitive:           true,
		MarkdownDescription: `TSIG key secret, base64 string.`,
		Validators: []validator.String{
			stringvalidator.AlsoRequires(
				path.MatchRelative().AtParent().AtName("name"),
				path.MatchRelative().AtParent().AtName("algorithm"),
			),
		},
	},
}

func ExpandConfigTSIGKey(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigTSIGKey {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigTSIGKeyModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigTSIGKeyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigTSIGKey {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigTSIGKey{
		Algorithm: flex.ExpandStringPointer(m.Algorithm),
		Comment:   flex.ExpandStringPointer(m.Comment),
		Key:       flex.ExpandStringPointer(m.Key),
		Name:      flex.ExpandStringPointer(m.Name),
		Secret:    flex.ExpandStringPointer(m.Secret),
	}
	return to
}

func FlattenConfigTSIGKey(ctx context.Context, from *dns_config.ConfigTSIGKey, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigTSIGKeyAttrTypes)
	}
	m := ConfigTSIGKeyModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigTSIGKeyAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigTSIGKeyModel) Flatten(ctx context.Context, from *dns_config.ConfigTSIGKey, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigTSIGKeyModel{}
	}
	m.Algorithm = flex.FlattenStringPointer(from.Algorithm)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Key = flex.FlattenStringPointer(from.Key)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.ProtocolName = flex.FlattenStringPointer(from.ProtocolName)
	m.Secret = flex.FlattenStringPointer(from.Secret)
}
