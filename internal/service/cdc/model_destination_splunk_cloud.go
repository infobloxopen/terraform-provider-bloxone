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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	internalvalidator "github.com/infobloxopen/terraform-provider-bloxone/internal/validator"
)

type DestinationSplunkCloudModel struct {
	Address                  types.String      `tfsdk:"address"`
	CaCertificate            types.String      `tfsdk:"ca_certificate"`
	Certificate              types.String      `tfsdk:"certificate"`
	CertificateKeyPassphrase types.String      `tfsdk:"certificate_key_passphrase"`
	ConfigFilename           types.String      `tfsdk:"config_filename"`
	ConfigOption             types.String      `tfsdk:"config_option"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	Description              types.String      `tfsdk:"description"`
	Enabled                  types.Bool        `tfsdk:"enabled"`
	Id                       types.Int64       `tfsdk:"id"`
	IndexerName              types.String      `tfsdk:"indexer_name"`
	InsecureMode             types.Bool        `tfsdk:"insecure_mode"`
	Name                     types.String      `tfsdk:"name"`
	OutputDataFormat         types.String      `tfsdk:"output_data_format"`
	Tags                     types.Map         `tfsdk:"tags"`
	TagsAll                  types.Map         `tfsdk:"tags_all"`
	UpdatedAt                timetypes.RFC3339 `tfsdk:"updated_at"`
}

var DestinationSplunkCloudAttrTypes = map[string]attr.Type{
	"address":                    types.StringType,
	"ca_certificate":             types.StringType,
	"certificate":                types.StringType,
	"certificate_key_passphrase": types.StringType,
	"config_filename":            types.StringType,
	"config_option":              types.StringType,
	"created_at":                 timetypes.RFC3339Type{},
	"description":                types.StringType,
	"enabled":                    types.BoolType,
	"id":                         types.Int64Type,
	"indexer_name":               types.StringType,
	"insecure_mode":              types.BoolType,
	"name":                       types.StringType,
	"output_data_format":         types.StringType,
	"tags":                       types.MapType{ElemType: types.StringType},
	"tags_all":                   types.MapType{ElemType: types.StringType},
	"updated_at":                 timetypes.RFC3339Type{},
}

var DestinationSplunkCloudResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			internalvalidator.IsValidHostPortList(),
		},
		MarkdownDescription: "The FQDN or IP address and port of the Splunk Cloud indexer, e.g. `splunkcloud.example.com:9997`. Multiple indexers are separated by commas, e.g. `10.10.10.1:9997,10.10.10.2:9997`.",
	},
	"ca_certificate": schema.StringAttribute{
		Required:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The CA certificate used to verify the Splunk Cloud indexer certificate, in PEM format.",
	},
	"certificate": schema.StringAttribute{
		Required:  true,
		Sensitive: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The Splunk forwarder certificate, in PEM format.",
	},
	"certificate_key_passphrase": schema.StringAttribute{
		Optional:            true,
		Sensitive:           true,
		MarkdownDescription: "The passphrase of the Splunk forwarder certificate key.",
	},
	"config_filename": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		Default:             stringdefault.StaticString(""),
		MarkdownDescription: "The name of the Universal Forwarder Credentials file downloaded from Splunk Cloud. Required when _config_option_ is _IMPORT_.",
	},
	"config_option": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString(string(cdc.CONFIGOPTION_DEFAULT)),
		Validators: []validator.String{
			stringvalidator.OneOf(
				string(cdc.CONFIGOPTION_DEFAULT),
				string(cdc.CONFIGOPTION_MANUAL),
				string(cdc.CONFIGOPTION_IMPORT),
			),
		},
		MarkdownDescription: "The way the Splunk Cloud connection is configured. Allowed values:\n\n  * _DEFAULT_\n  * _MANUAL_\n  * _IMPORT_\n\nDefaults to _DEFAULT_.\n\n~> **Note:** With _IMPORT_, set the connection fields from the Universal Forwarder Credentials file downloaded from Splunk Cloud, and _config_filename_ to its name.",
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
		MarkdownDescription: "The description of the Splunk Cloud destination. May contain 0 to 256 characters.",
	},
	"enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Indicates whether the Splunk Cloud destination is enabled. Defaults to _true_.",
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
		MarkdownDescription: "The name of the Splunk Cloud index the data is sent to.",
	},
	"insecure_mode": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Indicates whether the Splunk Cloud indexer certificate is not verified. Always _false_: insecure mode is not supported for Splunk Cloud.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthAtLeast(1),
		},
		MarkdownDescription: "The name of the Splunk Cloud destination.",
	},
	"output_data_format": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Default:  stringdefault.StaticString("CSV"),
		Validators: []validator.String{
			stringvalidator.OneOf("CSV", "CIM"),
		},
		MarkdownDescription: "The format of the data sent to Splunk Cloud. Allowed values:\n\n  * _CSV_\n  * _CIM_\n\nDefaults to _CSV_.",
	},
	"tags": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		Computed:            true,
		Default:             mapdefault.StaticValue(types.MapNull(types.StringType)),
		MarkdownDescription: "The tags for the Splunk Cloud destination.",
	},
	"tags_all": schema.MapAttribute{
		ElementType:         types.StringType,
		Computed:            true,
		MarkdownDescription: "The tags for the Splunk Cloud destination, including default tags.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
}

func ExpandDestinationSplunkCloud(ctx context.Context, o types.Object, diags *diag.Diagnostics) *cdc.DestinationSplunkCloud {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m DestinationSplunkCloudModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *DestinationSplunkCloudModel) Expand(ctx context.Context, diags *diag.Diagnostics) *cdc.DestinationSplunkCloud {
	if m == nil {
		return nil
	}
	to := &cdc.DestinationSplunkCloud{
		Address:                  flex.ExpandStringPointer(m.Address),
		CaCertificate:            flex.ExpandStringPointer(m.CaCertificate),
		Certificate:              flex.ExpandStringPointer(m.Certificate),
		CertificateKeyPassphrase: flex.ExpandStringPointer(m.CertificateKeyPassphrase),
		ConfigFilename:           flex.ExpandStringPointer(m.ConfigFilename),
		ConfigOption:             expandConfigOption(m.ConfigOption),
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

func DataSourceFlattenDestinationSplunkCloud(ctx context.Context, from *cdc.DestinationSplunkCloud, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DestinationSplunkCloudAttrTypes)
	}
	m := DestinationSplunkCloudModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, DestinationSplunkCloudAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DestinationSplunkCloudModel) Flatten(ctx context.Context, from *cdc.DestinationSplunkCloud, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DestinationSplunkCloudModel{}
	}
	// ca_certificate, certificate, certificate_key_passphrase are never read
	// back from the API: the backend masks secrets the same way it does for
	// ca_certificate on the syslog destination.
	m.Address = flex.FlattenStringPointer(from.Address)
	m.ConfigFilename = flex.FlattenStringPointer(from.ConfigFilename)
	m.ConfigOption = flattenConfigOption(from.ConfigOption)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Id = flex.FlattenInt64Pointer(from.Id)
	m.IndexerName = flex.FlattenStringPointer(from.IndexerName)
	m.InsecureMode = types.BoolPointerValue(from.InsecureMode)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.OutputDataFormat = flex.FlattenStringPointer(from.OutputDataFormat)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}

func expandConfigOption(v types.String) *cdc.ConfigOption {
	if v.IsNull() || v.IsUnknown() {
		return nil
	}
	o := cdc.ConfigOption(v.ValueString())
	return &o
}

func flattenConfigOption(v *cdc.ConfigOption) types.String {
	if v == nil {
		return types.StringNull()
	}
	return types.StringValue(string(*v))
}
