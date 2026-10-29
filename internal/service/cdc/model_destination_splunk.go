package cdc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
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

type DestinationSplunkModel struct {
	Address                  types.String      `tfsdk:"address"`
	CaCertificate            types.String      `tfsdk:"ca_certificate"`
	Certificate              types.String      `tfsdk:"certificate"`
	CertificateKeyPassphrase types.String      `tfsdk:"certificate_key_passphrase"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	Description              types.String      `tfsdk:"description"`
	Enabled                  types.Bool        `tfsdk:"enabled"`
	Id                       types.Int64       `tfsdk:"id"`
	IndexerName              types.String      `tfsdk:"indexer_name"`
	InsecureMode             types.Bool        `tfsdk:"insecure_mode"`
	Name                     types.String      `tfsdk:"name"`
	OutputDataFormat         types.String      `tfsdk:"output_data_format"`
	Tags                     types.Map         `tfsdk:"tags"`
	UpdatedAt                timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationSplunkAttrTypes = map[string]attr.Type{
	"address":                    types.StringType,
	"ca_certificate":             types.StringType,
	"certificate":                types.StringType,
	"certificate_key_passphrase": types.StringType,
	"created_at":                 timetypes.RFC3339Type{},
	"description":                types.StringType,
	"enabled":                    types.BoolType,
	"id":                         types.Int64Type,
	"indexer_name":               types.StringType,
	"insecure_mode":              types.BoolType,
	"name":                       types.StringType,
	"output_data_format":         types.StringType,
	"tags":                       types.MapType{ElemType: types.StringType},
	"updated_at":                 timetypes.RFC3339Type{},
}

var DestinationSplunkResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			internalvalidator.IsValidHostPortList(),
		},
		MarkdownDescription: "The FQDN or IP address and port of the Splunk indexer, e.g. `splunk.example.com:9997`. Multiple indexers are separated by commas, e.g. `10.10.10.1:9997,10.10.10.2:9997`.",
	},
	"ca_certificate": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The CA certificate used to verify the Splunk indexer certificate, in PEM format.",
	},
	"certificate": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The Splunk forwarder certificate, in PEM format.",
	},
	"certificate_key_passphrase": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The passphrase of the Splunk forwarder certificate key.",
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
		MarkdownDescription: "The description of the Splunk destination. May contain 0 to 256 characters.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the Splunk destination is enabled. Defaults to _true_.",
	},
	"id": schema.Int64Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int64{
			int64planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"indexer_name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the Splunk index the data is sent to.",
	},
	"insecure_mode": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the Splunk indexer certificate is not verified. Defaults to _true_.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the Splunk destination.",
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("CSV"),
		Validators: []validator.String{
			stringvalidator.OneOf("CSV", "CIM"),
		},
		MarkdownDescription: "The format of the data sent to Splunk. Allowed values:\n\n  * _CSV_\n  * _CIM_\n\nDefaults to _CSV_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The tags for the Splunk destination.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandDestinationSplunk(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationSplunk {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationSplunkModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationSplunkModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationSplunk {
	if m == nil {
		return nil
	}
	to := &cdc.DestinationSplunk{
		Address:                  flex.ExpandStringPointer(m.Address),
		CaCertificate:            flex.ExpandStringPointer(m.CaCertificate),
		Certificate:              flex.ExpandStringPointer(m.Certificate),
		CertificateKeyPassphrase: flex.ExpandStringPointer(m.CertificateKeyPassphrase),
		Description:              flex.ExpandStringPointer(m.Description),
		Enabled:                  flex.ExpandBoolPointer(m.Enabled),
		IndexerName:              flex.ExpandStringPointer(m.IndexerName),
		InsecureMode:             flex.ExpandBoolPointer(m.InsecureMode),
		Name:                     flex.ExpandStringPointer(m.Name),
		OutputDataFormat:         flex.ExpandStringPointer(m.OutputDataFormat),
		Tags:                     flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenDestinationSplunk(ctx context.Context, from *cdc.DestinationSplunk, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationSplunkAttrTypes)
	}
	m := DestinationSplunkModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, DestinationSplunkAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationSplunkModel) Flatten(ctx context.Context, from *cdc.DestinationSplunk, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationSplunkModel{}
	}
	// ca_certificate, certificate, certificate_key_passphrase are never read
	// back from the API: the backend masks secrets the same way it does for
	// ca_certificate on the syslog destination.
	m.Address = flex.FlattenStringPointer(from.Address)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IndexerName = flex.FlattenStringPointer(from.IndexerName)
	m.InsecureMode = types.BoolPointerValue(from.InsecureMode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutputDataFormat = flex.FlattenStringPointer(from.OutputDataFormat)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
