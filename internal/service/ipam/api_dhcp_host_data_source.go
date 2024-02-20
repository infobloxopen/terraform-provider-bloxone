package ipam

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
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

type IpamsvcHostModelWithFilter struct {
	Filters         types.Map      `tfsdk:"filters"`
	TagFilters      types.Map      `tfsdk:"tag_filters"`
	Results         types.List     `tfsdk:"results"`
	RetryIfNotFound types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (m *IpamsvcHostModelWithFilter) FlattenResults(ctx context.Context, from []ipam.IpamsvcHost, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, IpamsvcHostAttrTypes, diags, FlattenIpamsvcHost)
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
					Attributes: utils.DataSourceAttributeMap(IpamsvcHostResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
			"retry_if_not_found": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If set to `true`, the data source will retry until a matching host is found, or until the Read Timeout expires.",
			},
			"timeouts": timeouts.Attributes(ctx, timeouts.Opts{
				Read: true,
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
	var data IpamsvcHostModelWithFilter

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
		allResults, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.IpamsvcHost, error) {
			apiRes, _, err := d.client.IPAddressManagementAPI.
				DhcpHostAPI.
				DhcpHostList(ctx).
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
