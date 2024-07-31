package infra_mgmt

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/inframgmt"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type InfraHostModel struct {
	Configs             types.List        `tfsdk:"configs"`
	ConnectivityMonitor types.Map         `tfsdk:"connectivity_monitor"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at"`
	CreatedBy           types.String      `tfsdk:"created_by"`
	Description         types.String      `tfsdk:"description"`
	DisplayName         types.String      `tfsdk:"display_name"`
	HostSubtype         types.String      `tfsdk:"host_subtype"`
	HostType            types.String      `tfsdk:"host_type"`
	HostVersion         types.String      `tfsdk:"host_version"`
	Id                  types.String      `tfsdk:"id"`
	IpAddress           types.String      `tfsdk:"ip_address"`
	IpSpace             types.String      `tfsdk:"ip_space"`
	LegacyId            types.String      `tfsdk:"legacy_id"`
	LocationId          types.String      `tfsdk:"location_id"`
	MacAddress          types.String      `tfsdk:"mac_address"`
	MaintenanceMode     types.String      `tfsdk:"maintenance_mode"`
	NatIp               types.String      `tfsdk:"nat_ip"`
	NoaCluster          types.String      `tfsdk:"noa_cluster"`
	Ophid               types.String      `tfsdk:"ophid"`
	PoolId              types.String      `tfsdk:"pool_id"`
	SerialNumber        types.String      `tfsdk:"serial_number"`
	Tags                types.Map         `tfsdk:"tags"`
	TagsAll             types.Map         `tfsdk:"tags_all"`
	Timezone            types.String      `tfsdk:"timezone"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at"`
}

var InfraHostAttrTypes = map[string]attr.Type{
	"configs":              types.ListType{ElemType: types.ObjectType{AttrTypes: InfraServiceHostConfigAttrTypes}},
	"connectivity_monitor": types.MapType{ElemType: types.StringType},
	"created_at":           timetypes.RFC3339Type{},
	"created_by":           types.StringType,
	"description":          types.StringType,
	"display_name":         types.StringType,
	"host_subtype":         types.StringType,
	"host_type":            types.StringType,
	"host_version":         types.StringType,
	"id":                   types.StringType,
	"ip_address":           types.StringType,
	"ip_space":             types.StringType,
	"legacy_id":            types.StringType,
	"location_id":          types.StringType,
	"mac_address":          types.StringType,
	"maintenance_mode":     types.StringType,
	"nat_ip":               types.StringType,
	"noa_cluster":          types.StringType,
	"ophid":                types.StringType,
	"pool_id":              types.StringType,
	"serial_number":        types.StringType,
	"tags":                 types.MapType{ElemType: types.StringType},
	"tags_all":             types.MapType{ElemType: types.StringType},
	"timezone":             types.StringType,
	"updated_at":           timetypes.RFC3339Type{},
}

var InfraHostResourceSchemaAttributes = map[string]schema.Attribute{
	"configs": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: InfraServiceHostConfigResourceSchemaAttributes,
		},
		Computed:            true,
		MarkdownDescription: "The list of Host-specific configurations for each Service deployed on this Host.",
	},
	"connectivity_monitor": schema.MapAttribute{
		ElementType: types.StringType,
		Computed:    true,
		MarkdownDescription: "Represents the connectivity monitor properties of a Host, to enable/disable connectivity monitoring for redundant network interfaces.\n\n" +
			"  The \"endpoint_type\" is:\n" +
			"  - `\"csp\"` for enabling monitoring\n" +
			"  - `\"\"` for disabling monitoring (default)\n\n" +
			"  Note: Currently, all fields except \"endpoint_type\" are read-only, and will be overridden to default values in case they are edited.\n\n" +
			"  Example:\n" +
			"  ```\n{\n    \"connectivity_monitor\": {\n      \"cost\":1000000,\n      \"endpoint_type\":\"csp\",\n      \"endpoint\":\"http://csp.infoblox.com\",\n      \"interval\":15,\n      \"failure_threshold\":1,\n      \"success_threshold\":2\n    }\n  }\n```",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The timestamp of creation of Host.",
	},
	"created_by": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The creator of the Host (internal use only).",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The description of the Host (optional).",
	},
	"display_name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the Host (unique).",
	},
	"host_subtype": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The sub-type of a specific Host type.  Example: For Host type BloxOne Appliance, sub-type could be \"B105\" or \"VEP1425\"",
	},
	"host_type": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The type of Host.  Should be one of: 1. NIOS , 2. NIOS HA, 3. BloxOne VM , 4. BloxOne Appliance, 5. BloxOne Container, 6. CNIOS",
	},
	"host_version": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The version of the Host platform services.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"ip_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The IP address of the Host.",
	},
	"ip_space": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The IP Space of the Host.",
	},
	"legacy_id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The legacy Host object identifier.",
	},
	"location_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"mac_address": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The MAC address of the Host.",
	},
	"maintenance_mode": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("disabled"),
	},
	"nat_ip": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The NAT IP address of the Host.",
	},
	"noa_cluster": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The CSP cluster identifier (internal use only).",
	},
	"ophid": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The unique On-Prem Host ID generated by the On-Prem device and assigned to the Host once it is registered and logged into the Infoblox Cloud.",
	},
	"pool_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"serial_number": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The unique serial number of the Host.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "Tags associated with this Host.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "Tags associated with this Host including default tags.",
	},
	"timezone": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The timezone of the Host.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "The timestamp of the latest update on Host.",
	},
}

func ExpandInfraHost(ctx context.Context, o types.Object, diags *diag.Diagnostics) *inframgmt.Host {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InfraHostModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InfraHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *inframgmt.Host {
	if m == nil {
		return nil
	}
	to := &inframgmt.Host{
		Description:     flex.ExpandStringPointer(m.Description),
		DisplayName:     flex.ExpandString(m.DisplayName),
		IpSpace:         flex.ExpandStringPointer(m.IpSpace),
		LocationId:      flex.ExpandStringPointer(m.LocationId),
		MaintenanceMode: flex.ExpandStringPointer(m.MaintenanceMode),
		PoolId:          flex.ExpandStringPointer(m.PoolId),
		SerialNumber:    flex.ExpandStringPointer(m.SerialNumber),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func DataSourceFlattenInfraHost(ctx context.Context, from *inframgmt.Host, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InfraHostAttrTypes)
	}
	m := InfraHostModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, InfraHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InfraHostModel) Flatten(ctx context.Context, from *inframgmt.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InfraHostModel{}
	}
	m.Configs = flex.FlattenFrameworkListNestedBlock(ctx, from.Configs, InfraServiceHostConfigAttrTypes, diags, FlattenInfraServiceHostConfig)
	m.ConnectivityMonitor = flex.FlattenFrameworkMapString(ctx, from.ConnectivityMonitor, diags)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.CreatedBy = flex.FlattenStringPointer(from.CreatedBy)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DisplayName = flex.FlattenString(from.DisplayName)
	m.HostSubtype = flex.FlattenStringPointer(from.HostSubtype)
	m.HostType = flex.FlattenStringPointer(from.HostType)
	m.HostVersion = flex.FlattenStringPointer(from.HostVersion)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.IpAddress = flex.FlattenStringPointer(from.IpAddress)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.LegacyId = flex.FlattenStringPointer(from.LegacyId)
	m.LocationId = flex.FlattenStringPointer(from.LocationId)
	m.MacAddress = flex.FlattenStringPointer(from.MacAddress)
	m.MaintenanceMode = flex.FlattenStringPointer(from.MaintenanceMode)
	m.NatIp = flex.FlattenStringPointer(from.NatIp)
	m.NoaCluster = flex.FlattenStringPointer(from.NoaCluster)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.PoolId = flex.FlattenStringPointer(from.PoolId)
	m.SerialNumber = flex.FlattenStringPointer(from.SerialNumber)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Timezone = flex.FlattenStringPointer(from.Timezone)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
