package dfp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/dfp"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type AtcdfpDfpHostModel struct {
	LegacyHostId types.Int64  `tfsdk:"legacy_host_id"`
	Name         types.String `tfsdk:"name"`
	Ophid        types.String `tfsdk:"ophid"`
	SiteId       types.String `tfsdk:"site_id"`
}

var AtcdfpDfpHostAttrTypes = map[string]attr.Type{
	"legacy_host_id": types.Int64Type,
	"name":           types.StringType,
	"ophid":          types.StringType,
	"site_id":        types.StringType,
}

var AtcdfpDfpHostResourceSchemaAttributes = map[string]schema.Attribute{
	"legacy_host_id": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "// The DNS Forwarding Proxy legacy ID object identifier.",
	},
	"name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The name of the DNS Forwarding Proxy.",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The On-Prem Host identifier.",
	},
	"site_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.",
	},
}

func ExpandAtcdfpDfpHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *dfp.DfpHost {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AtcdfpDfpHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AtcdfpDfpHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dfp.DfpHost {
	if m == nil {
		return nil
	}
	to := &dfp.DfpHost{
		LegacyHostId: flex.ExpandInt32Pointer(m.LegacyHostId),
	}
	return to
}

func FlattenAtcdfpDfpHost(ctx context.Context, from *dfp.DfpHost, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AtcdfpDfpHostAttrTypes)
	}
	m := AtcdfpDfpHostModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AtcdfpDfpHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AtcdfpDfpHostModel) Flatten(ctx context.Context, from *dfp.DfpHost, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AtcdfpDfpHostModel{}
	}
	m.LegacyHostId = flex.FlattenInt32Pointer(from.LegacyHostId)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
}
