package dns_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type ConfigZoneAuthorityMNameBlockModel struct {
	Mname           types.String `tfsdk:"mname"`
	ProtocolMname   types.String `tfsdk:"protocol_mname"`
	UseDefaultMname types.Bool   `tfsdk:"use_default_mname"`
}

var ConfigZoneAuthorityMNameBlockAttrTypes = map[string]attr.Type{
	"mname":             types.StringType,
	"protocol_mname":    types.StringType,
	"use_default_mname": types.BoolType,
}

var ConfigZoneAuthorityMNameBlockResourceSchemaAttributes = map[string]schema.Attribute{
	"mname": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: `Defaults to empty.`,
	},
	"protocol_mname": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: `Optional. Master name server in punycode.  Defaults to empty.`,
	},
	"use_default_mname": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: `Optional. Use default value for master name server.  Defaults to true.`,
	},
}

func ExpandConfigZoneAuthorityMNameBlock(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.ZoneAuthorityMNameBlock {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigZoneAuthorityMNameBlockModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigZoneAuthorityMNameBlockModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.ZoneAuthorityMNameBlock {
	if m == nil {
		return nil
	}
	to := &dnsconfig.ZoneAuthorityMNameBlock{
		Mname:           flex.ExpandStringPointer(m.Mname),
		UseDefaultMname: flex.ExpandBoolPointer(m.UseDefaultMname),
	}
	return to
}

func FlattenConfigZoneAuthorityMNameBlock(ctx context.Context, from *dnsconfig.ZoneAuthorityMNameBlock, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigZoneAuthorityMNameBlockAttrTypes)
	}
	m := ConfigZoneAuthorityMNameBlockModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigZoneAuthorityMNameBlockAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigZoneAuthorityMNameBlockModel) Flatten(ctx context.Context, from *dnsconfig.ZoneAuthorityMNameBlock, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigZoneAuthorityMNameBlockModel{}
	}
	m.Mname = flex.FlattenStringPointer(from.Mname)
	m.ProtocolMname = flex.FlattenStringPointer(from.ProtocolMname)
	m.UseDefaultMname = types.BoolPointerValue(from.UseDefaultMname)
}
