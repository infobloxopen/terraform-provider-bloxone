package ipam_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestDataSourceNextAvailableSubnet(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_subnets.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "bloxone_ipam_next_available_subnets" "test" {
						id = "/test/address_block/123455678"
						cidr = 25
					}	
				`,
				ExpectError: regexp.MustCompile("invalid resource ID specified"),
			},
			{
				Config: testAccDataSourceNextAvailableSubnet(1, 26),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cidr", dataSourceName, "results.0.cidr"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableSubnet(3, 27),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "subnet_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2.address"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cidr", dataSourceName, "results.0.cidr"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cidr", dataSourceName, "results.1.cidr"),
					resource.TestCheckResourceAttrPair(dataSourceName, "cidr", dataSourceName, "results.2.cidr"),
				),
			},
		},
	})
}

func testAccDataSourceNextAvailableSubnetBaseConfig() string {
	return `
	resource "bloxone_ipam_ip_space" "test" {
		name = "test_ip_space"
	}
	resource "bloxone_ipam_address_block" "test" {
		name = "test_address_block"
		address = "192.168.0.0"
		cidr = "24"
		space = bloxone_ipam_ip_space.test.id
	}
`
}
func testAccDataSourceNextAvailableSubnet(count, cidr int) string {
	var config string
	if count == 1 {
		config = fmt.Sprintf(`
	data "bloxone_ipam_next_available_subnets" "test" {
		id = bloxone_ipam_address_block.test.id
		cidr = %d
	}`, cidr)
	} else {
		config = fmt.Sprintf(`
	data "bloxone_ipam_next_available_subnets" "test" {
		id = bloxone_ipam_address_block.test.id
		cidr = %d
		subnet_count = %d
	}`, cidr, count)
	}

	return strings.Join([]string{testAccDataSourceNextAvailableSubnetBaseConfig(), config}, "")
}
