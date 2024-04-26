package infra_mgmt

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/bloxone-go-client/inframgmt"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type InfraServiceModel struct {
	Configs         types.List        `tfsdk:"configs"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Description     types.String      `tfsdk:"description"`
	DesiredState    types.String      `tfsdk:"desired_state"`
	DesiredVersion  types.String      `tfsdk:"desired_version"`
	Id              types.String      `tfsdk:"id"`
	InterfaceLabels types.List        `tfsdk:"interface_labels"`
	Name            types.String      `tfsdk:"name"`
	PoolId          types.String      `tfsdk:"pool_id"`
	ServiceType     types.String      `tfsdk:"service_type"`
	Tags            types.Map         `tfsdk:"tags"`
	TagsAll         types.Map         `tfsdk:"tags_all"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var InfraServiceAttrTypes = map[string]attr.Type{
	"configs":          types.ListType{ElemType: types.ObjectType{AttrTypes: InfraServiceHostConfigAttrTypes}},
	"created_at":       timetypes.RFC3339Type{},
	"description":      types.StringType,
	"desired_state":    types.StringType,
	"desired_version":  types.StringType,
	"id":               types.StringType,
	"interface_labels": types.ListType{ElemType: types.StringType},
	"name":             types.StringType,
	"pool_id":          types.StringType,
	"service_type":     types.StringType,
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
	"updated_at":       timetypes.RFC3339Type{},
}

func InfraServiceResourceSchemaAttributes() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"configs": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: InfraServiceHostConfigResourceSchemaAttributes,
			},
			Computed:            true,
			MarkdownDescription: "List of Host-specific configurations of this Service.",
		},
		"created_at": schema.StringAttribute{
			CustomType:          timetypes.RFC3339Type{},
			Computed:            true,
			MarkdownDescription: "Timestamp of creation of Service.",
		},
		"description": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The description of the Service (optional).",
		},
		"desired_state": schema.StringAttribute{
			Optional:            true,
			Computed:            true,
			Default:             stringdefault.StaticString("stop"),
			MarkdownDescription: "The desired state of the Service. Should either be `\"start\"` or `\"stop\"`.",
			Validators: []validator.String{
				stringvalidator.OneOf("start", "stop"),
			},
		},
		"desired_version": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The desired version of the Service.",
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"interface_labels": schema.ListAttribute{
			ElementType:         types.StringType,
			Optional:            true,
			MarkdownDescription: "List of interfaces on which this Service can operate. Note: The list can contain custom interface labels (Example: `[\"WAN\",\"LAN\",\"label1\",\"label2\"]`)",
		},
		"name": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The name of the Service (unique).",
		},
		"pool_id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"service_type": schema.StringAttribute{
			Required: true,
			MarkdownDescription: "The type of the Service deployed on the Host (`dns`, `cdc`, etc.). The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):\n\n" +
				"  | Service name          | Service type | \n" +
				"  | --------------------- | ------------ | \n" +
				"  | Access Authentication | authn        | \n" +
				"  | Anycast               | anycast      | \n" +
				"  | Data Connector        | cdc          | \n" +
				"  | DHCP                  | dhcp         | \n" +
				"  | DNS                   | dns          | \n" +
				"  | DNS Forwarding Proxy  | dfp          | \n" +
				"  | NIOS Grid Connector   | orpheus      | \n" +
				"  | MS AD Sync            | msad         | \n" +
				"  | NTP                   | ntp          | \n" +
				"  | BGP                   | bgp          | \n" +
				"  | RIP                   | rip          | \n" +
				"  | OSPF                  | ospf         | \n" +
				"  <br>",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.RequiresReplaceIfConfigured(),
			},
			Validators: []validator.String{
				stringvalidator.OneOf("authn", "anycast", "cdc", "dhcp", "dns", "dfp", "orpheus", "msad", "ntp", "bgp", "rip", "ospf"),
			},
		},
		"tags": schema.MapAttribute{
			ElementType:         types.StringType,
			Optional:            true,
			MarkdownDescription: "Tags associated with this Service.",
		},
		"tags_all": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "Tags associated with this Service including default tags.",
		},
		"updated_at": schema.StringAttribute{
			CustomType:          timetypes.RFC3339Type{},
			Computed:            true,
			MarkdownDescription: "Timestamp of the latest update on Service.",
		},
	}
}

func ExpandInfraService(ctx context.Context, o types.Object, diags *diag.Diagnostics) *inframgmt.Service {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m InfraServiceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *InfraServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *inframgmt.Service {
	if m == nil {
		return nil
	}
	to := &inframgmt.Service{
		Description:     flex.ExpandStringPointer(m.Description),
		DesiredState:    flex.ExpandStringPointer(m.DesiredState),
		DesiredVersion:  flex.ExpandStringPointer(m.DesiredVersion),
		InterfaceLabels: flex.ExpandFrameworkListString(ctx, m.InterfaceLabels, diags),
		Name:            flex.ExpandString(m.Name),
		PoolId:          flex.ExpandString(m.PoolId),
		ServiceType:     flex.ExpandString(m.ServiceType),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.TagsAll, diags),
	}
	return to
}

func DataSourceFlattenInfraService(ctx context.Context, from *inframgmt.Service, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(InfraServiceAttrTypes)
	}
	m := InfraServiceModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, InfraServiceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *InfraServiceModel) Flatten(ctx context.Context, from *inframgmt.Service, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InfraServiceModel{}
	}
	m.Configs = flex.FlattenFrameworkListNestedBlock(ctx, from.Configs, InfraServiceHostConfigAttrTypes, diags, FlattenInfraServiceHostConfig)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DesiredState = flex.FlattenStringPointer(from.DesiredState)
	m.DesiredVersion = flex.FlattenStringPointer(from.DesiredVersion)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InterfaceLabels = flex.FlattenFrameworkListString(ctx, from.InterfaceLabels, diags)
	m.Name = flex.FlattenString(from.Name)
	m.PoolId = flex.FlattenString(from.PoolId)
	m.ServiceType = flex.FlattenString(from.ServiceType)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}

// InfraServiceModelWithTimeouts is a helper model that also contains the timeouts for the InfraServiceModel.
// The plugin framework does not deal well with embedded structs, so have to duplicate the fields and Flatten/Expand functions here from InfraServiceModel.
// https://github.com/hashicorp/terraform-plugin-framework/issues/242
type InfraServiceModelWithTimeouts struct {
	Configs         types.List        `tfsdk:"configs"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Description     types.String      `tfsdk:"description"`
	DesiredState    types.String      `tfsdk:"desired_state"`
	DesiredVersion  types.String      `tfsdk:"desired_version"`
	Id              types.String      `tfsdk:"id"`
	InterfaceLabels types.List        `tfsdk:"interface_labels"`
	Name            types.String      `tfsdk:"name"`
	PoolId          types.String      `tfsdk:"pool_id"`
	ServiceType     types.String      `tfsdk:"service_type"`
	Tags            types.Map         `tfsdk:"tags"`
	TagsAll         types.Map         `tfsdk:"tags_all"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
	Timeouts        timeouts.Value    `tfsdk:"timeouts"`
	WaitForState    types.Bool        `tfsdk:"wait_for_state"`
}

func (m *InfraServiceModelWithTimeouts) Expand(ctx context.Context, diags *diag.Diagnostics) *inframgmt.Service {
	if m == nil {
		return nil
	}
	to := &inframgmt.Service{
		Description:     flex.ExpandStringPointer(m.Description),
		DesiredState:    flex.ExpandStringPointer(m.DesiredState),
		DesiredVersion:  flex.ExpandStringPointer(m.DesiredVersion),
		InterfaceLabels: flex.ExpandFrameworkListString(ctx, m.InterfaceLabels, diags),
		Name:            flex.ExpandString(m.Name),
		PoolId:          flex.ExpandString(m.PoolId),
		ServiceType:     flex.ExpandString(m.ServiceType),
		Tags:            flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func (m *InfraServiceModelWithTimeouts) Flatten(ctx context.Context, from *inframgmt.Service, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = InfraServiceModelWithTimeouts{}
	}
	m.Configs = flex.FlattenFrameworkListNestedBlock(ctx, from.Configs, InfraServiceHostConfigAttrTypes, diags, FlattenInfraServiceHostConfig)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.DesiredState = flex.FlattenStringPointer(from.DesiredState)
	m.DesiredVersion = flex.FlattenStringPointer(from.DesiredVersion)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InterfaceLabels = flex.FlattenFrameworkListString(ctx, from.InterfaceLabels, diags)
	m.Name = flex.FlattenString(from.Name)
	m.PoolId = flex.FlattenString(from.PoolId)
	m.ServiceType = flex.FlattenString(from.ServiceType)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}

func InfraServiceResourceSchemaAttributesWithTimeouts(ctx context.Context) map[string]schema.Attribute {
	attributes := InfraServiceResourceSchemaAttributes()
	attributes["timeouts"] = timeouts.Attributes(ctx, timeouts.Opts{
		Create:            true,
		CreateDescription: "[Duration](https://pkg.go.dev/time#ParseDuration) to wait before being considered a timeout during create operations. Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Default is 20m.",
		Update:            true,
		UpdateDescription: "[Duration](https://pkg.go.dev/time#ParseDuration) to wait before being considered a timeout during update operations. Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Default is 20m.",
	})
	attributes["wait_for_state"] = schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "If set to `true`, the resource will wait for the desired state to be reached before returning. If set to `false`, the resource will return immediately after the request is sent to the API.",
	}
	return attributes
}
