package infra_mgmt

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/inframgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type InfraServiceHostConfigModel struct {
	CurrentVersion types.String      `tfsdk:"current_version"`
	ExtraData      types.String      `tfsdk:"extra_data"`
	HostId         types.String      `tfsdk:"host_id"`
	Id             types.String      `tfsdk:"id"`
	ServiceId      types.String      `tfsdk:"service_id"`
	ServiceType    types.String      `tfsdk:"service_type"`
	UpgradedAt     timetypes.RFC3339 `tfsdk:"upgraded_at"`
}

var InfraServiceHostConfigAttrTypes = map[string]attr.Type{
	"current_version": types.StringType,
	"extra_data":      types.StringType,
	"host_id":         types.StringType,
	"id":              types.StringType,
	"service_id":      types.StringType,
	"service_type":    types.StringType,
	"upgraded_at":     timetypes.RFC3339Type{},
}

var InfraServiceHostConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"current_version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The current version of the Service deployed on the Host.",
	},
	"extra_data": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The field to carry any extra data specific to this configuration.",
	},
	"host_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"service_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The type of the Service deployed on the Host (`dns`, `cdc`, etc.). The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):\n\n  | Service name | Service type | \n  | ------ | ------ | \n  | Access Authentication | authn | \n  | Anycast | anycast | \n  | Data Connector | cdc | \n  | DHCP | dhcp | \n  | DNS | dns | \n  | DNS Forwarding Proxy | dfp | \n  | NIOS Grid Connector | orpheus | \n  | MS AD Sync | msad | \n  | NTP | ntp | \n  | BGP | bgp | \n  | RIP | rip | \n  | OSPF | ospf | \n",
	},
	"upgraded_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Optional:            true,
		MarkdownDescription: "The timestamp of the latest upgrade of the Host-specific Service configuration.",
	},
}

func ExpandInfraServiceHostConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *inframgmt.ServiceHostConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InfraServiceHostConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InfraServiceHostConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *inframgmt.ServiceHostConfig {
	if m == nil {
		return nil
	}
	to := &inframgmt.ServiceHostConfig{
		CurrentVersion: flex.ExpandStringPointer(m.CurrentVersion),
		ExtraData:      flex.ExpandStringPointer(m.ExtraData),
		HostId:         flex.ExpandStringPointer(m.HostId),
		ServiceId:      flex.ExpandStringPointer(m.ServiceId),
		ServiceType:    flex.ExpandStringPointer(m.ServiceType),
		UpgradedAt:     flex.ExpandTimePointer(ctx, m.UpgradedAt, diags),
	}
	return to
}

func FlattenInfraServiceHostConfig(ctx context.Context, from *inframgmt.ServiceHostConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InfraServiceHostConfigAttrTypes)
	}
	m := InfraServiceHostConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, InfraServiceHostConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InfraServiceHostConfigModel) Flatten(ctx context.Context, from *inframgmt.ServiceHostConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InfraServiceHostConfigModel{}
	}
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.ExtraData = flex.FlattenStringPointer(from.ExtraData)
	m.HostId = flex.FlattenStringPointer(from.HostId)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceType = flex.FlattenStringPointer(from.ServiceType)
	m.UpgradedAt = timetypes.NewRFC3339TimePointerValue(from.UpgradedAt)
}
