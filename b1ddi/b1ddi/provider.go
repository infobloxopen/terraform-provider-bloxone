package b1ddi

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
)

const (
	errAddressBlockNotFound          = "response status code indicates client error (status 400): \n{\"error\":[{\"message\":\"Failed to find address block\"}]}"
	errIncorrectUtilizationUpdateRef = "response status code indicates client error (status 400): \n{\"error\":[{\"message\":\"The 'Utilization Update' object does not refer to a valid 'IP Space' object.\"}]}"
	errRecordNotFound                = "response status code indicates client error (status 404): \n{\"error\":[{\"message\":\"record not found\"}]}"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_HOST", nil),
				Description: "BloxOne DDI host URL.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_API_KEY", nil),
				Description: "API token for authentication against the Infoblox BloxOne DDI platform.",
			},
			"base_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "api/ddi/v1",
				Description: "The base path is to indicate the API version and the product name.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"b1ddi_ip_space":        resourceIpamsvcIPSpace(),
			"b1ddi_subnet":          resourceIpamsvcSubnet(),
			"b1ddi_fixed_address":   resourceIpamsvcFixedAddress(),
			"b1ddi_address_block":   resourceIpamsvcAddressBlock(),
			"b1ddi_range":           resourceIpamsvcRange(),
			"b1ddi_address":         resourceIpamsvcAddress(),
			"b1ddi_dns_view":        resourceConfigView(),
			"b1ddi_dns_auth_zone":   resourceConfigAuthZone(),
			"b1ddi_dns_record":      resourceDataRecord(),
			"b1ddi_dns_auth_nsg":    resourceConfigAuthNSG(),
			"b1ddi_dns_forward_nsg": resourceConfigForwardNSG(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"b1ddi_ip_spaces":                   dataSourceIpamsvcIPSpace(),
			"b1ddi_subnets":                     dataSourceIpamsvcSubnet(),
			"b1ddi_fixed_addresses":             dataSourceIpamsvcFixedAddress(),
			"b1ddi_address_blocks":              dataSourceIpamsvcAddressBlock(),
			"b1ddi_ranges":                      dataSourceIpamsvcRange(),
			"b1ddi_addresses":                   dataSourceIpamsvcAddress(),
			"b1ddi_option_codes":                dataSourceIpamsvcOptionCode(),
			"b1ddi_dns_views":                   dataSourceConfigView(),
			"b1ddi_dns_auth_zones":              dataSourceConfigAuthZone(),
			"b1ddi_dns_records":                 dataSourceDataRecord(),
			"b1ddi_dhcp_hosts":                  dataSourceIpamsvcDhcpHost(),
			"b1ddi_dns_hosts":                   dataSourceDnsHost(),
			"b1ddi_dns_auth_nsgs":               dataSourceConfigAuthNSG(),
			"b1ddi_dns_forward_nsgs":            dataSourceConfigForwardNSG(),
			"b1ddi_ipam_next_available_subnets": dataSourceIpamsvcNextAvailableSubnet(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	apiKey := d.Get("api_key").(string)
	basePath := d.Get("base_path").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if host == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without the API host",
		})
		return nil, diags
	}

	if apiKey == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without the API key. Export the B1DDI_API_KEY environment variable to set the API key.",
		})
		return nil, diags
	}

	// create the transport
	transport := httptransport.New(
		host, basePath, nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1ddiclient.B1DDIAPIKey(apiKey)
	transport.DefaultAuthentication = tokenAuth

	// create the API client
	c := b1ddiclient.NewClient(transport, strfmt.Default)
	return c, diags
}

// Generates filter string for B1DDI API list request from the map
func filterFromMap(filtersMap map[string]interface{}) string {
	filters := make([]string, 0, len(filtersMap))

	for k, v := range filtersMap {
		if val, err := strconv.Atoi(v.(string)); err == nil {
			filters = append(filters, fmt.Sprintf("%s==%v", k, val))
		} else {
			filters = append(filters, fmt.Sprintf("%s=='%s'", k, v))
		}
	}

	return strings.Join(filters, " and ")
}

// dataSourceSchemaFromResource -- generates schema for results field in each data source
// This function gets the original resource, schema, and injects the ID field in it.
func dataSourceSchemaFromResource(resource func() *schema.Resource) *schema.Resource {
	// Get the resource schema
	resultSchema := resource().Schema
	// Inject id field into the resource schema
	resultSchema["id"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The resource identifier.",
	}
	return &schema.Resource{
		Schema: resultSchema,
	}
}
