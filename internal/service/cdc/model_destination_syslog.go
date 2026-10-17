package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type DestinationSyslogModel struct {
	Address           types.String      `tfsdk:"address"`
	CaCertificate     types.String      `tfsdk:"ca_certificate"`
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	Description       types.String      `tfsdk:"description"`
	Enabled           types.Bool        `tfsdk:"enabled"`
	Id                types.Int64       `tfsdk:"id"`
	InsecureMode      types.Bool        `tfsdk:"insecure_mode"`
	Name              types.String      `tfsdk:"name"`
	OutputDataFormat  types.String      `tfsdk:"output_data_format"`
	Port              types.Int64       `tfsdk:"port"`
	Tags              types.Map         `tfsdk:"tags"`
	TransportProtocol types.String      `tfsdk:"transport_protocol"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationSyslogAttrTypes = map[string]attr.Type{
	"address":            types.StringType,
	"ca_certificate":     types.StringType,
	"created_at":         timetypes.RFC3339Type{},
	"description":        types.StringType,
	"enabled":            types.BoolType,
	"id":                 types.Int64Type,
	"insecure_mode":      types.BoolType,
	"name":               types.StringType,
	"output_data_format": types.StringType,
	"port":               types.Int64Type,
	"tags":               types.MapType{ElemType: types.StringType},
	"transport_protocol": types.StringType,
	"updated_at":         timetypes.RFC3339Type{},
}

var DestinationSyslogResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
	},
	"ca_certificate": schema.StringAttribute{
		Optional:  true,
		Sensitive: true,
	},
	"created_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"insecure_mode": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			boolplanmodifier.UseStateForUnknown(),
		},
	},
	"name": schema.StringAttribute{
		Optional: true,
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
	},
	"tags": schema.MapAttribute{
		ElementType: types.StringType,
		Optional:    true,
	},
	"transport_protocol": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	},
	"updated_at": schema.StringAttribute{
		CustomType: timetypes.RFC3339Type{},
		Computed:   true,
	},
}

func ExpandDestinationSyslog(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationSyslog {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationSyslogModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationSyslogModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationSyslog {
	if m == nil {
		return nil
	}
	to := &cdc.DestinationSyslog{
		Address:           flex.ExpandStringPointer(m.Address),
		CaCertificate:     flex.ExpandStringPointer(m.CaCertificate),
		Description:       flex.ExpandStringPointer(m.Description),
		Enabled:           flex.ExpandBoolPointer(m.Enabled),
		InsecureMode:      flex.ExpandBoolPointer(m.InsecureMode),
		Name:              flex.ExpandStringPointer(m.Name),
		OutputDataFormat:  flex.ExpandStringPointer(m.OutputDataFormat),
		Port:              flex.ExpandInt64Pointer(m.Port),
		Tags:              flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
		TransportProtocol: flex.ExpandStringPointer(m.TransportProtocol),
	}
	return to
}

func FlattenDestinationSyslog(ctx context.Context, from *cdc.DestinationSyslog, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationSyslogAttrTypes)
	}
	m := DestinationSyslogModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationSyslogAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationSyslogModel) Flatten(ctx context.Context, from *cdc.DestinationSyslog, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationSyslogModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	// ca_certificate is never read back from the API: the backend masks it
	// the same way it masks nios_password/scp_password on SourceNios.
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.InsecureMode = types.BoolPointerValue(from.InsecureMode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutputDataFormat = flex.FlattenStringPointer(from.OutputDataFormat)
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TransportProtocol = flex.FlattenStringPointer(from.TransportProtocol)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
