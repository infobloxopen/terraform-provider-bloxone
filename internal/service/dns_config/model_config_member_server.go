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

type ConfigMemberServerModel struct {
	Host types.String `tfsdk:"host"`
}

var ConfigMemberServerAttrTypes = map[string]attr.Type{
	"host": types.StringType,
}

var ConfigMemberServerResourceSchemaAttributes = map[string]schema.Attribute{
	"host": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

func ExpandConfigMemberServer(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dnsconfig.MemberServer {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigMemberServerModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigMemberServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.MemberServer {
	if m == nil {
		return nil
	}
	to := &dnsconfig.MemberServer{
		Host: flex.ExpandString(m.Host),
	}
	return to
}

func FlattenConfigMemberServer(ctx context.Context, from *dnsconfig.MemberServer, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigMemberServerAttrTypes)
	}
	m := ConfigMemberServerModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigMemberServerAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigMemberServerModel) Flatten(ctx context.Context, from *dnsconfig.MemberServer, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigMemberServerModel{}
	}
	m.Host = flex.FlattenString(from.Host)
}
