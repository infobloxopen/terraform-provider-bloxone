package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccServerDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_server.test"
	resourceName := "bloxone_dhcp_server.test"
	var v ipam.IpamsvcServer

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfigFilters("TEST-DHCP-SERVER"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServerExists(context.Background(), resourceName, &v),
					}, testAccCheckServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccServerDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_server.test"
	resourceName := "bloxone_dhcp_server.test"
	var v ipam.IpamsvcServer
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfigTagFilters("TEST-DHCP-SERVER", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServerExists(context.Background(), resourceName, &v),
					}, testAccCheckServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckServerResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "client_principal", dataSourceName, "results.0.client_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_client_update", dataSourceName, "results.0.ddns_client_update"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_conflict_resolution_mode", dataSourceName, "results.0.ddns_conflict_resolution_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domain", dataSourceName, "results.0.ddns_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_enabled", dataSourceName, "results.0.ddns_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_name", dataSourceName, "results.0.ddns_generate_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generated_prefix", dataSourceName, "results.0.ddns_generated_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_send_updates", dataSourceName, "results.0.ddns_send_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl_percent", dataSourceName, "results.0.ddns_ttl_percent"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_on_renew", dataSourceName, "results.0.ddns_update_on_renew"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_conflict_resolution", dataSourceName, "results.0.ddns_use_conflict_resolution"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_zones", dataSourceName, "results.0.ddns_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_config", dataSourceName, "results.0.dhcp_config"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options_v6", dataSourceName, "results.0.dhcp_options_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "gss_tsig_fallback", dataSourceName, "results.0.gss_tsig_fallback"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_filename", dataSourceName, "results.0.header_option_filename"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_address", dataSourceName, "results.0.header_option_server_address"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_name", dataSourceName, "results.0.header_option_server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_char", dataSourceName, "results.0.hostname_rewrite_char"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_enabled", dataSourceName, "results.0.hostname_rewrite_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_regex", dataSourceName, "results.0.hostname_rewrite_regex"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_kdc", dataSourceName, "results.0.kerberos_kdc"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_keys", dataSourceName, "results.0.kerberos_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_rekey_interval", dataSourceName, "results.0.kerberos_rekey_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_retry_interval", dataSourceName, "results.0.kerberos_retry_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_tkey_lifetime", dataSourceName, "results.0.kerberos_tkey_lifetime"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_tkey_protocol", dataSourceName, "results.0.kerberos_tkey_protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "server_principal", dataSourceName, "results.0.server_principal"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_specific_option_option_space", dataSourceName, "results.0.vendor_specific_option_option_space"),
	}
}

func testAccServerDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test" {
  name = %q
}

data "bloxone_dhcp_servers" "test" {
  filters = {
	name = bloxone_dhcp_server.test.name
  }
}
`, name)
}

func testAccServerDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dhcp_servers" "test" {
  tag_filters = {
	tag1 = bloxone_dhcp_server.test.tags.tag1
  }
}
`, name, tagValue)
}
