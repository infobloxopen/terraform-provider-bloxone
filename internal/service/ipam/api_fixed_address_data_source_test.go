package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccFixedAddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_fixed_addresses.test"
	resourceName := "bloxone_dhcp_fixed_address.test"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedAddressDataSourceConfigFilters(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedAddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFixedAddressDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_fixed_addresses.test"
	resourceName := "bloxone_dhcp_fixed_address.test"
	var v ipam.IpamsvcFixedAddress
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFixedAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFixedAddressDataSourceConfigTagFilters(spaceName, "10.0.0.10", "mac", "aa:aa:aa:aa:aa:aa", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFixedAddressExists(context.Background(), resourceName, &v),
					}, testAccCheckFixedAddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFixedAddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_dhcp", dataSourceName, "results.0.disable_dhcp"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_filename", dataSourceName, "results.0.header_option_filename"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_address", dataSourceName, "results.0.header_option_server_address"),
		resource.TestCheckResourceAttrPair(resourceName, "header_option_server_name", dataSourceName, "results.0.header_option_server_name"),
		resource.TestCheckResourceAttrPair(resourceName, "hostname", dataSourceName, "results.0.hostname"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_assigned_hosts", dataSourceName, "results.0.inheritance_assigned_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_parent", dataSourceName, "results.0.inheritance_parent"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_space", dataSourceName, "results.0.ip_space"),
		resource.TestCheckResourceAttrPair(resourceName, "match_type", dataSourceName, "results.0.match_type"),
		resource.TestCheckResourceAttrPair(resourceName, "match_value", dataSourceName, "results.0.match_value"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccFixedAddressDataSourceConfigFilters(spaceName, address, matchType, matchValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test" {
  ip_space = bloxone_ipam_ip_space.test.id
  address = %q
  match_type = %q
  match_value = %q
  depends_on = [bloxone_ipam_subnet.test]
}

data "bloxone_dhcp_fixed_addresses" "test" {
  filters = {
	address = bloxone_dhcp_fixed_address.test.address
	ip_space = bloxone_ipam_ip_space.test.id
  }
}
`, address, matchType, matchValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}

func testAccFixedAddressDataSourceConfigTagFilters(spaceName, address, matchType, matchValue, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_fixed_address" "test" {
  ip_space = bloxone_ipam_ip_space.test.id
  address = %q
  match_type = %q
  match_value = %q
  tags = {
	tag1 = %q
  }
  depends_on = [bloxone_ipam_subnet.test]
}

data "bloxone_dhcp_fixed_addresses" "test" {
  tag_filters = {
	tag1 = bloxone_dhcp_fixed_address.test.tags.tag1
  }
}
`, address, matchType, matchValue, tagValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(spaceName), config}, "")
}
