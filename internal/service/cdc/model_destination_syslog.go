package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
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
		Required: true,
		Validators: []validator.String{
			internalvalidator.IsValidFQDNOrIPAddress(),
		},
		MarkdownDescription: "The FQDN or IP address of the syslog server.",
	},
	"ca_certificate": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The CA certificate used to verify the syslog server certificate, in PEM format. Used for TLS, together with _transport_protocol_ _TCP_ and _insecure_mode_ _false_.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been created.",
	},
	"description": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(""),
		Validators: []validator.String{
			stringvalidator.LengthAtMost(256),
		},
		MarkdownDescription: "The description of the syslog destination. May contain 0 to 256 characters.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the syslog destination is enabled. Defaults to _true_.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"insecure_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the syslog server certificate is not verified. Defaults to _true_.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the syslog destination.",
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("CEF"),
		Validators: []validator.String{
			stringvalidator.OneOf("CEF", "LEEF"),
		},
		MarkdownDescription: "The format of the data sent to the syslog server. Allowed values:\n\n  * _CEF_\n  * _LEEF_\n\nDefaults to _CEF_.",
	},
	"port": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(514),
		Validators: []validator.Int64{
			int64validator.Between(1, 65535),
		},
		MarkdownDescription: "The port of the syslog server. Defaults to _514_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the syslog destination.",
	},
	"transport_protocol": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("TCP"),
		Validators: []validator.String{
			stringvalidator.OneOf("TCP", "UDP"),
		},
		MarkdownDescription: "The transport protocol used to connect to the syslog server. Allowed values:\n\n  * _TCP_\n  * _UDP_\n\nDefaults to _TCP_. For TLS, use _TCP_ with _insecure_mode_ set to _false_ and a _ca_certificate_.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
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
