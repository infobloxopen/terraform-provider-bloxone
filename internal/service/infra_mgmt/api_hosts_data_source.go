package infra_mgmt

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
	"github.com/infobloxopen/bloxone-go-client/inframgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &HostsDataSource{}

func NewHostsDataSource() datasource.DataSource {
	return &HostsDataSource{}
}

// HostsDataSource defines the data source implementation.
type HostsDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *HostsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "infra_hosts"
}

type InfraHostModelWithFilter struct {
	Filters         types.Map      `tfsdk:"filters"`
	TagFilters      types.Map      `tfsdk:"tag_filters"`
	Results         types.List     `tfsdk:"results"`
	RetryIfNotFound types.Bool     `tfsdk:"retry_if_not_found"`
	Timeouts        timeouts.Value `tfsdk:"timeouts"`
}

func (m *InfraHostModelWithFilter) FlattenResults(ctx context.Context, from []inframgmt.Host, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, InfraHostAttrTypes, diags, FlattenInfraHost)
}

func (d *HostsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing Infrastructure Hosts.",
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
					Attributes: utils.DataSourceAttributeMap(InfraHostResourceSchemaAttributes, &resp.Diagnostics),
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

func (d *HostsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *HostsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data InfraHostModelWithFilter

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
		allResults, err := utils.ReadWithPages(func(offset, limit int32) ([]inframgmt.Host, error) {
			apiRes, _, err := d.client.InfraManagementAPI.
				HostsAPI.
				List(ctx).
				Filter(flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)).
				Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
				Execute()
			if err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read InfraHost, got error: %s", err))
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
