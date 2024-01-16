package ipam

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NextAvailableIPDataSource{}

func NewIpamNextAvailableIPDataSource() datasource.DataSource {
	return &NextAvailableIPDataSource{}
}

// NextAvailableIPDataSource defines the data source implementation.
type NextAvailableIPDataSource struct {
	client *bloxoneclient.APIClient
}

type IpamsvcNextAvailableIPModel struct {
	Id         types.String `tfsdk:"id"`
	Contiguous types.Bool   `tfsdk:"contiguous"`
	Count      types.Int64  `tfsdk:"ip_count"`
	Results    types.List   `tfsdk:"results"`
}

func (m *IpamsvcNextAvailableIPModel) FlattenResults(ctx context.Context, from []ipam.IpamsvcAddress, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	var listOfAddress []string

	for _, address := range from {
		listOfAddress = append(listOfAddress, address.Address)
	}
	m.Results = flex.FlattenFrameworkListString(ctx, listOfAddress, diags)
}

func (d *NextAvailableIPDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NextAvailableIPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_next_available_ips"
}

func (d *NextAvailableIPDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: ``,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `An application specific resource identity of a resource`,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^ipam/(range|subnet|address_block)/[0-9a-f-].*$`), "invalid resource ID specified"),
				},
			},
			// Query parameter
			"contiguous": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: `Indicates whether the IP addresses should belong to a contiguous block. Defaults to false.`,
			},
			// Query parameter
			"ip_count": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: `The number of IP addresses requested. Defaults to 1.`,
			},
			"results": schema.ListAttribute{
				ElementType:         types.StringType,
				Computed:            true,
				MarkdownDescription: "List of Next available IP address in the specified resource",
			},
		},
	}
}

func (d *NextAvailableIPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var (
		data   IpamsvcNextAvailableIPModel
		apiRes *ipam.IpamsvcNextAvailableIPResponse
		err    error
	)

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	addressStr := data.Id.ValueString()

	switch addressStr[:strings.LastIndex(addressStr, "/")] {
	case "ipam/address_block":
		apiRes, _, err = d.client.IPAddressManagementAPI.AddressBlockAPI.
			AddressBlockListNextAvailableIP(ctx, data.Id.ValueString()).
			Count(int32(data.Count.ValueInt64())).
			Contiguous(data.Contiguous.ValueBool()).
			Execute()

	case "ipam/subnet":
		apiRes, _, err = d.client.IPAddressManagementAPI.SubnetAPI.
			SubnetListNextAvailableIP(ctx, data.Id.ValueString()).
			Count(int32(data.Count.ValueInt64())).
			Contiguous(data.Contiguous.ValueBool()).
			Execute()

	case "ipam/range":
		apiRes, _, err = d.client.IPAddressManagementAPI.RangeAPI.
			RangeListNextAvailableIP(ctx, data.Id.ValueString()).
			Count(int32(data.Count.ValueInt64())).
			Contiguous(data.Contiguous.ValueBool()).
			Execute()
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read Address, got error: %s", err))
		return
	}
	data.FlattenResults(ctx, apiRes.GetResults(), &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
