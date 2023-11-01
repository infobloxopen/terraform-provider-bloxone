package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccIpSpaceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ip_spaces.test"
	resourceName := "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpSpaceDataSourceConfigFilters("ip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					}, testAccCheckIpSpaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpSpaceDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ip_spaces.test"
	resourceName := "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpSpaceDataSourceConfigTagFilters("ip_space_name", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					}, testAccCheckIpSpaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpSpaceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "asm_config", dataSourceName, "results.0.asm_config"),
		resource.TestCheckResourceAttrPair(resourceName, "asm_scope_flag", dataSourceName, "results.0.asm_scope_flag"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_client_update", dataSourceName, "results.0.ddns_client_update"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_conflict_resolution_mode", dataSourceName, "results.0.ddns_conflict_resolution_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_domain", dataSourceName, "results.0.ddns_domain"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generate_name", dataSourceName, "results.0.ddns_generate_name"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_generated_prefix", dataSourceName, "results.0.ddns_generated_prefix"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_send_updates", dataSourceName, "results.0.ddns_send_updates"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_ttl_percent", dataSourceName, "results.0.ddns_ttl_percent"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_update_on_renew", dataSourceName, "results.0.ddns_update_on_renew"),
		resource.TestCheckResourceAttrPair(resourceName, "ddns_use_conflict_resolution", dataSourceName, "results.0.ddns_use_conflict_resolution"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_config", dataSourceName, "results.0.dhcp_config"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options_v6", dataSourceName, "results.0.dhcp_options_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_filename", dataSourceName, "results.0.header_option_filename"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_address", dataSourceName, "results.0.header_option_server_address"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_name", dataSourceName, "results.0.header_option_server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_char", dataSourceName, "results.0.hostname_rewrite_char"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_enabled", dataSourceName, "results.0.hostname_rewrite_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname_rewrite_regex", dataSourceName, "results.0.hostname_rewrite_regex"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "threshold", dataSourceName, "results.0.threshold"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "results.0.utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization_v6", dataSourceName, "results.0.utilization_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "vendor_specific_option_option_space", dataSourceName, "results.0.vendor_specific_option_option_space"),
	}
}

func testAccIpSpaceDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test" {
  name = %q
}

data "bloxone_ipam_ip_spaces" "test" {
  filters = {
	name = bloxone_ipam_ip_space.test.name
  }
}
`, name)
}

func testAccIpSpaceDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_ipam_ip_spaces" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_ip_space.test.tags.tag1
  }
}
`, name, tagValue)
}
