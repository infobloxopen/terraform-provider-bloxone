package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DNSConfigModel struct {
	ConsolidatedZoneDataEnabled types.Bool   `tfsdk:"consolidated_zone_data_enabled"`
	SplitViewEnabled            types.Bool   `tfsdk:"split_view_enabled"`
	SyncType                    types.String `tfsdk:"sync_type"`
	ViewId                      types.String `tfsdk:"view_id"`
	ViewName                    types.String `tfsdk:"view_name"`
}

var DNSConfigAttrTypes = map[string]attr.Type{
	"consolidated_zone_data_enabled": types.BoolType,
	"split_view_enabled":             types.BoolType,
	"sync_type":                      types.StringType,
	"view_id":                        types.StringType,
	"view_name":                      types.StringType,
}

var DNSConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"consolidated_zone_data_enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
	},
	"split_view_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "split_view_enabled consolidates private zones into a single view, which is separate from the public zone view.",
	},
	"sync_type": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of sync.",
	},
	"view_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Unique identifier of the view.",
	},
	"view_name": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Name of the view.",
	},
}

func ExpandDNSConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *clouddiscovery.DNSConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DNSConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DNSConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *clouddiscovery.DNSConfig {
	if m == nil {
		return nil
	}
	to := &clouddiscovery.DNSConfig{
		ConsolidatedZoneDataEnabled: flex.ExpandBoolPointer(m.ConsolidatedZoneDataEnabled),
		SplitViewEnabled:            flex.ExpandBoolPointer(m.SplitViewEnabled),
		SyncType:                    flex.ExpandStringPointer(m.SyncType),
		ViewId:                      flex.ExpandStringPointer(m.ViewId),
		ViewName:                    flex.ExpandStringPointer(m.ViewName),
	}
	return to
}

func FlattenDNSConfig(ctx context.Context, from *clouddiscovery.DNSConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DNSConfigAttrTypes)
	}
	m := DNSConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DNSConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DNSConfigModel) Flatten(ctx context.Context, from *clouddiscovery.DNSConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DNSConfigModel{}
	}
	m.ConsolidatedZoneDataEnabled = types.BoolPointerValue(from.ConsolidatedZoneDataEnabled)
	m.SplitViewEnabled = types.BoolPointerValue(from.SplitViewEnabled)
	m.SyncType = flex.FlattenStringPointer(from.SyncType)
	m.ViewId = flex.FlattenStringPointer(from.ViewId)
	m.ViewName = flex.FlattenStringPointer(from.ViewName)
}