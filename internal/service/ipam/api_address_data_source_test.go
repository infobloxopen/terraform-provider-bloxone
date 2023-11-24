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

func TestAccAddressDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_addresses.test"
	resourceName := "bloxone_ipam_address.test"
	var v ipam.IpamsvcAddress

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAddressDataSourceConfigFilters("10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAddressExists(context.Background(), resourceName, &v),
					}, testAccCheckAddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAddressDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_addresses.test"
	resourceName := "bloxone_ipam_address.test"
	var v ipam.IpamsvcAddress
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAddressDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAddressDataSourceConfigTagFilters("10.0.0.1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAddressExists(context.Background(), resourceName, &v),
					}, testAccCheckAddressResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAddressResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_info", dataSourceName, "results.0.dhcp_info"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_dhcp", dataSourceName, "results.0.disable_dhcp"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_attrs", dataSourceName, "results.0.discovery_attrs"),
		resource.TestCheckResourceAttrPair(resourceName, "discovery_metadata", dataSourceName, "results.0.discovery_metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "host", dataSourceName, "results.0.host"),
		resource.TestCheckResourceAttrPair(resourceName, "hwaddr", dataSourceName, "results.0.hwaddr"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "interface", dataSourceName, "results.0.interface"),
		resource.TestCheckResourceAttrPair(resourceName, "names", dataSourceName, "results.0.names"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "range", dataSourceName, "results.0.range"),
		resource.TestCheckResourceAttrPair(resourceName, "space", dataSourceName, "results.0.space"),
		resource.TestCheckResourceAttrPair(resourceName, "state", dataSourceName, "results.0.state"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "usage", dataSourceName, "results.0.usage"),
	}
}

func testAccAddressDataSourceConfigFilters(address string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test" {
  address = %q
  space = bloxone_ipam_ip_space.test.id
  depends_on = [bloxone_ipam_subnet.test]
}

data "bloxone_ipam_addresses" "test" {
  filters = {
	address = bloxone_ipam_address.test.address
  }
}
`, address)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccAddressDataSourceConfigTagFilters(address, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address" "test" {
  address = %q
  space = bloxone_ipam_ip_space.test.id
  depends_on = [bloxone_ipam_subnet.test]
  tags = {
	tag1 = %q
  }
}

data "bloxone_ipam_addresses" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_address.test.tags.tag1
  }
}
`, address, tagValue)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")

}
