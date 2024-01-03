package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dns_config"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigACLItemModel struct {
	Access  types.String `tfsdk:"access"`
	Acl     types.String `tfsdk:"acl"`
	Address types.String `tfsdk:"address"`
	Element types.String `tfsdk:"element"`
	TsigKey types.Object `tfsdk:"tsig_key"`
}

var ConfigACLItemAttrTypes = map[string]attr.Type{
	"access":   types.StringType,
	"acl":      types.StringType,
	"address":  types.StringType,
	"element":  types.StringType,
	"tsig_key": types.ObjectType{AttrTypes: ConfigTSIGKeyAttrTypes},
}

var ConfigACLItemResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Access permission for _element_.  Allowed values:  * _allow_,  * _deny_.`,
	},
	"acl": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `The resource identifier.`,
	},
	"address": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_.`,
	},
	"element": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: `Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_.`,
	},
	"tsig_key": schema.SingleNestedAttribute{
		Attributes: ConfigTSIGKeyResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandConfigACLItem(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dns_config.ConfigACLItem {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigACLItemModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigACLItemModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dns_config.ConfigACLItem {
	if m == nil {
		return nil
	}
	to := &dns_config.ConfigACLItem{
		Access:  flex.ExpandString(m.Access),
		Acl:     flex.ExpandStringPointer(m.Acl),
		Address: flex.ExpandStringPointer(m.Address),
		Element: flex.ExpandString(m.Element),
		TsigKey: ExpandConfigTSIGKey(ctx, m.TsigKey, diags),
	}
	return to
}

func FlattenConfigACLItem(ctx context.Context, from *dns_config.ConfigACLItem, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigACLItemAttrTypes)
	}
	m := ConfigACLItemModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigACLItemAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigACLItemModel) Flatten(ctx context.Context, from *dns_config.ConfigACLItem, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigACLItemModel{}
	}
	m.Access = flex.FlattenString(from.Access)
	m.Acl = flex.FlattenStringPointer(from.Acl)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.Element = flex.FlattenString(from.Element)
	m.TsigKey = FlattenConfigTSIGKey(ctx, from.TsigKey, diags)
}
