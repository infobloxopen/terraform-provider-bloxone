package dns_config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &HostDataSource{}

func NewHostDataSource() datasource.DataSource {
	return &HostDataSource{}
}

var HostAttrTypes = map[string]attr.Type{
	"absolute_name":          types.StringType,
	"address":                types.StringType,
	"anycast_addresses":      types.ListType{ElemType: types.StringType},
	"associated_server":      types.ObjectType{AttrTypes: ConfigHostAssociatedServerAttrTypes},
	"comment":                types.StringType,
	"current_version":        types.StringType,
	"dfp":                    types.BoolType,
	"dfp_service":            types.StringType,
	"id":                     types.StringType,
	"inheritance_sources":    types.ObjectType{AttrTypes: ConfigHostInheritanceAttrTypes},
	"kerberos_keys":          types.ListType{ElemType: types.ObjectType{AttrTypes: ConfigKerberosKeyAttrTypes}},
	"name":                   types.StringType,
	"ophid":                  types.StringType,
	"protocol_absolute_name": types.StringType,
	"provider_id":            types.StringType,
	"server":                 types.StringType,
	"site_id":                types.StringType,
	"tags":                   types.MapType{ElemType: types.StringType},
	"tags_all":               types.MapType{ElemType: types.StringType},
	"type":                   types.StringType,
}

func HostDataSourceSchemaAttributes(diag *diag.Diagnostics) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"absolute_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host FQDN.",
		},
		"address": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host's primary IP Address.",
		},
		"anycast_addresses": schema.ListAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
		},
		"associated_server": schema.SingleNestedAttribute{
			Attributes:          utils.DataSourceAttributeMap(ConfigHostAssociatedServerResourceSchemaAttributes, diag),
			Optional:            true,
			MarkdownDescription: "Host associated server configuration.",
		},
		"comment": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host description.",
		},
		"current_version": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host current version.",
		},
		"dfp": schema.BoolAttribute{
			Computed:            true,
			MarkdownDescription: "Below _dfp_ field is deprecated and not supported anymore. The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field.",
		},
		"dfp_service": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP. Allowed values:  * _unavailable_: BloxOne TD DFP application is not available,  * _enabled_: BloxOne TD DFP application is available and enabled,  * _disabled_: BloxOne TD DFP application is available but disabled.",
		},
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"inheritance_sources": schema.SingleNestedAttribute{
			Attributes:          utils.DataSourceAttributeMap(ConfigHostInheritanceResourceSchemaAttributes, diag),
			Optional:            true,
			MarkdownDescription: "Optional. Inheritance configuration.",
		},
		"kerberos_keys": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: utils.DataSourceAttributeMap(ConfigKerberosKeyResourceSchemaAttributes, diag),
			},
			Optional:            true,
			MarkdownDescription: "Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host display name.",
		},
		"ophid": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "On-Prem Host ID.",
		},
		"protocol_absolute_name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host FQDN in punycode.",
		},
		"provider_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "External provider identifier.",
		},
		"server": schema.StringAttribute{
			Optional:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"site_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Host site ID.",
		},
		"tags": schema.MapAttribute{
			ElementType:         types.StringType,
			Optional:            true,
			MarkdownDescription: "Host tagging specifics.",
		},
		"tags_all": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "Host tagging specifics includes default tags.",
		},
		"type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform.",
		},
	}
}

type HostModel struct {
	AbsoluteName         types.String `tfsdk:"absolute_name"`
	Address              types.String `tfsdk:"address"`
	AnycastAddresses     types.List   `tfsdk:"anycast_addresses"`
	AssociatedServer     types.Object `tfsdk:"associated_server"`
	Comment              types.String `tfsdk:"comment"`
	CurrentVersion       types.String `tfsdk:"current_version"`
	Dfp                  types.Bool   `tfsdk:"dfp"`
	DfpService           types.String `tfsdk:"dfp_service"`
	Id                   types.String `tfsdk:"id"`
	InheritanceSources   types.Object `tfsdk:"inheritance_sources"`
	KerberosKeys         types.List   `tfsdk:"kerberos_keys"`
	Name                 types.String `tfsdk:"name"`
	Ophid                types.String `tfsdk:"ophid"`
	ProtocolAbsoluteName types.String `tfsdk:"protocol_absolute_name"`
	ProviderId           types.String `tfsdk:"provider_id"`
	Server               types.String `tfsdk:"server"`
	SiteId               types.String `tfsdk:"site_id"`
	Tags                 types.Map    `tfsdk:"tags"`
	TagsAll              types.Map    `tfsdk:"tags_all"`
	Type                 types.String `tfsdk:"type"`
}

// HostDataSource defines the data source implementation.
type HostDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *HostDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dns_hosts"
}

type HostModelWithFilter struct {
	Filters         types.Map      `tfsdk:"filters"`
	TagFilters      types.Map      `tfsdk:"tag_filters"`
	Results         types.List     `tfsdk:"results"`
	RetryIfNotFound types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (m *HostModelWithFilter) FlattenResults(ctx context.Context, from []dnsconfig.Host, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, HostAttrTypes, diags, FlattenConfigHost)
}

func (d *HostDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing DNS Hosts.\n\nA DNS Host object associates DNS configuration with hosts.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"tag_filters": schema.MapAttribute{
				Description: "Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: HostDataSourceSchemaAttributes(&resp.Diagnostics),
				},
				Computed: true,
			},
			"retry_if_not_found": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If set to `true`, the data source will retry until a matching host is found, or until the Read Timeout expires.",
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Read:            true,
				ReadDescription: "[Duration](https://pkg.go.dev/time#ParseDuration) to wait before being considered a timeout during read operations. Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Default is 20m.",
			}),
		},
	}
}

func (d *HostDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*bloxoneclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *bloxoneclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *HostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data HostModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	readTimeout, diags := data.Timeouts.Read(ctx, 20*time.Minute)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	err := retry.RetryContext(ctx, readTimeout, func() *retry.RetryError {
		allResults, err := utils.ReadWithPages(func(offset, limit int32) ([]dnsconfig.Host, error) {
			apiRes, _, err := d.client.DNSConfigurationAPI.
				HostAPI.
				List(ctx).
				Filter(flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)).
				Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
				Offset(offset).
				Limit(limit).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Host, got error: %s", err))
				return nil, err
			}
			return apiRes.GetResults(), nil
		})
		if err != nil {
			return retry.NonRetryableError(err)
		}

		if len(allResults) == 0 {
			if data.RetryIfNotFound.ValueBool() {
				return retry.RetryableError(errors.New("no matching hosts found; will retry"))
			}
			return nil
		}
		data.FlattenResults(ctx, allResults, &resp.Diagnostics)
		return nil
	})
	if err != nil {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (m *HostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *dnsconfig.Host {
	if m == nil {
		return nil
	}
	to := &dnsconfig.Host{
		AssociatedServer:   ExpandConfigHostAssociatedServer(ctx, m.AssociatedServer, diags),
		InheritanceSources: ExpandConfigHostInheritance(ctx, m.InheritanceSources, diags),
		KerberosKeys:       flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandConfigKerberosKey),
		Server:             flex.ExpandStringPointer(m.Server),
		Tags:               flex.ExpandFrameworkMapString(ctx, m.Tags, diags),
	}
	return to
}

func FlattenConfigHost(ctx context.Context, from *dnsconfig.Host, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(HostAttrTypes)
	}
	m := HostModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, HostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *HostModel) Flatten(ctx context.Context, from *dnsconfig.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = HostModel{}
	}
	m.AbsoluteName = flex.FlattenStringPointer(from.AbsoluteName)
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenConfigHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.Dfp = types.BoolPointerValue(from.Dfp)
	m.DfpService = flex.FlattenStringPointer(from.DfpService)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InheritanceSources = FlattenConfigHostInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, ConfigKerberosKeyAttrTypes, diags, FlattenConfigKerberosKey)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProtocolAbsoluteName = flex.FlattenStringPointer(from.ProtocolAbsoluteName)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.SiteId = flex.FlattenStringPointer(from.SiteId)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
