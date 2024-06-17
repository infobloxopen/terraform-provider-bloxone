package ipam

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
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DhcpHostDataSource{}

var DhcpHostAttrTypes = map[string]attr.Type{
	"address":           types.StringType,
	"anycast_addresses": types.ListType{ElemType: types.StringType},
	"associated_server": types.ObjectType{AttrTypes: IpamsvcHostAssociatedServerAttrTypes},
	"comment":           types.StringType,
	"current_version":   types.StringType,
	"id":                types.StringType,
	"ip_space":          types.StringType,
	"name":              types.StringType,
	"ophid":             types.StringType,
	"provider_id":       types.StringType,
	"server":            types.StringType,
	"tags":              types.MapType{ElemType: types.StringType},
	"tags_all":          types.MapType{ElemType: types.StringType},
	"type":              types.StringType,
}

func DhcpHostDataSourceSchemaAttributes(diags *diag.Diagnostics) map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"address": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The primary IP address of the on-prem host.",
		},
		"anycast_addresses": schema.ListAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "Anycast address configured to the host. Order is not significant.",
		},
		"associated_server": schema.SingleNestedAttribute{
			Attributes:          utils.DataSourceAttributeMap(IpamsvcHostAssociatedServerResourceSchemaAttributes, diags),
			Computed:            true,
			MarkdownDescription: "The DHCP Config Profile for the on-prem host.",
		},
		"comment": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The description for the on-prem host.",
		},
		"current_version": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Current dhcp application version of the host.",
		},
		"id": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"ip_space": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"name": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The display name of the on-prem host.",
		},
		"ophid": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "The on-prem host ID.",
		},
		"provider_id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "External provider identifier.",
		},
		"server": schema.StringAttribute{
			Required:            true,
			MarkdownDescription: "The resource identifier.",
		},
		"tags": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "The tags of the on-prem host in JSON format.",
		},
		"tags_all": schema.MapAttribute{
			ElementType:         types.StringType,
			Computed:            true,
			MarkdownDescription: "The tags of the on-prem host in JSON format including default tags.",
		},
		"type": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory.",
		},
	}
}

func NewDhcpHostDataSource() datasource.DataSource {
	return &DhcpHostDataSource{}
}

// DhcpHostDataSource defines the data source implementation.
type DhcpHostDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *DhcpHostDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "dhcp_hosts"
}

type DhcpHostModel struct {
	Address          types.String `tfsdk:"address"`
	AnycastAddresses types.List   `tfsdk:"anycast_addresses"`
	AssociatedServer types.Object `tfsdk:"associated_server"`
	Comment          types.String `tfsdk:"comment"`
	CurrentVersion   types.String `tfsdk:"current_version"`
	Id               types.String `tfsdk:"id"`
	IpSpace          types.String `tfsdk:"ip_space"`
	Name             types.String `tfsdk:"name"`
	Ophid            types.String `tfsdk:"ophid"`
	ProviderId       types.String `tfsdk:"provider_id"`
	Server           types.String `tfsdk:"server"`
	Tags             types.Map    `tfsdk:"tags"`
	TagsAll          types.Map    `tfsdk:"tags_all"`
	Type             types.String `tfsdk:"type"`
}

type DhcpHostModelWithFilter struct {
	Filters         types.Map      `tfsdk:"filters"`
	TagFilters      types.Map      `tfsdk:"tag_filters"`
	Results         types.List     `tfsdk:"results"`
	RetryIfNotFound types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (m *DhcpHostModelWithFilter) FlattenResults(ctx context.Context, from []ipam.Host, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, DhcpHostAttrTypes, diags, FlattenDhcpHostDataSource)
}

func (d *DhcpHostDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing DHCP Hosts.\n\nA DHCP Host object associates a DHCP Config Profile with an on-prem host.",
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
					Attributes: DhcpHostDataSourceSchemaAttributes(&resp.Diagnostics),
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

func (d *DhcpHostDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *DhcpHostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data DhcpHostModelWithFilter

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
		allResults, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.Host, error) {
			apiRes, _, err := d.client.IPAddressManagementAPI.
				DhcpHostAPI.
				List(ctx).
				Filter(flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)).
				Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
				Offset(offset).
				Limit(limit).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read DhcpHost, got error: %s", err))
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

func (m *DhcpHostModel) Expand(ctx context.Context, diags *diag.Diagnostics) *ipam.Host {
	if m == nil {
		return nil
	}
	to := &ipam.Host{
		Server: flex.ExpandStringPointer(m.Server),
	}
	return to
}

func FlattenDhcpHostDataSource(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(DhcpHostAttrTypes)
	}
	m := DhcpHostModel{}
	m.Flatten(ctx, from, diags)
	m.Tags = m.TagsAll
	t, d := types.ObjectValueFrom(ctx, DhcpHostAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *DhcpHostModel) Flatten(ctx context.Context, from *ipam.Host, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = DhcpHostModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.AnycastAddresses = flex.FlattenFrameworkListString(ctx, from.AnycastAddresses, diags)
	m.AssociatedServer = FlattenIpamsvcHostAssociatedServer(ctx, from.AssociatedServer, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CurrentVersion = flex.FlattenStringPointer(from.CurrentVersion)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ProviderId = flex.FlattenStringPointer(from.ProviderId)
	m.Server = flex.FlattenStringPointer(from.Server)
	m.Tags = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.TagsAll = flex.FlattenFrameworkMapString(ctx, from.Tags, diags)
	m.Type = flex.FlattenStringPointer(from.Type)
}
