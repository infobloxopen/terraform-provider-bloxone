package ipam_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestDataSourceNextAvailableAddressBlock(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_address_blocks.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "bloxone_ipam_next_available_address_blocks" "test" {
						id = "/test/address_block/123455678"
						cidr = 25
					}	
				`,
				ExpectError: regexp.MustCompile("invalid resource ID specified"),
			},
			{
				Config: testAccDataSourceNextAvailableAddressBlock(1, 26),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableAddressBlock(3, 27),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "address_block_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableAddressBlockWithSingleTagFilter(24, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableAddressBlockWithMultipleTagFilters(24, 3),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "3"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2"),
				),
			},
		},
	})
}

func testAccDataSourceNextAvailableAddressBlock(count, cidr int) string {
	var config string
	if count == 1 {
		config = fmt.Sprintf(`
	data "bloxone_ipam_next_available_address_blocks" "test" {
		id = bloxone_ipam_address_block.test.id
		cidr = %d
	}`, cidr)
	} else {
		config = fmt.Sprintf(`
	data "bloxone_ipam_next_available_address_blocks" "test" {
		id = bloxone_ipam_address_block.test.id
		cidr = %d
		address_block_count = %d
	}`, cidr, count)
	}

	return strings.Join([]string{testAccDataSourceNextAvailableSubnetBaseConfig(), config}, "")
}

// testAccDataSourceNextAvailableAddressBlockWithSingleTagFilter creates test configuration for next available address block with a single tag filter
func testAccDataSourceNextAvailableAddressBlockWithSingleTagFilter(cidr, count int) string {
	config := fmt.Sprintf(`
	data "bloxone_ipam_next_available_address_blocks" "test" {
		cidr = %d
		address_block_count = %d
		tag_filters = {
			environment = "prd"
		}
      depends_on = [
        "bloxone_ipam_address_block.test_tagged_env_only"
    ]
	}`, cidr, count)

	return strings.Join([]string{testAccDataSourceNextAvailableAddressBlockWithTagsBaseConfig(), config}, "")
}

// testAccDataSourceNextAvailableAddressBlockWithMultipleTagFilters creates test configuration for next available address block with multiple tag filters
func testAccDataSourceNextAvailableAddressBlockWithMultipleTagFilters(cidr, count int) string {
	config := fmt.Sprintf(`
	data "bloxone_ipam_next_available_address_blocks" "test" {
		cidr = %d
		address_block_count = %d
		tag_filters = {
			environment = "prd"
			location = "data-center-1"
		}
	depends_on = [
        "bloxone_ipam_address_block.test_tagged_env_only"
    ]
	}`, cidr, count)

	return strings.Join([]string{testAccDataSourceNextAvailableAddressBlockWithTagsBaseConfig(), config}, "")
}

// testAccDataSourceNextAvailableAddressBlockWithTagsBaseConfig creates base resources with tags for testing
func testAccDataSourceNextAvailableAddressBlockWithTagsBaseConfig() string {
	return `
	resource "bloxone_ipam_ip_space" "test" {
		name = "test-acc-next-available-address-blocks-tags"
	}

	resource "bloxone_ipam_address_block" "test_tagged" {
		address = "192.168.0.0"
		cidr = 16
		space = bloxone_ipam_ip_space.test.id
	}
	
	resource "bloxone_ipam_address_block" "test_tagged2" {
		address = "13.0.0.0"
		cidr = 16
		space = bloxone_ipam_ip_space.test.id
		tags = {
			environment = "prd"
			location = "data-center-1"
		}
	}

	resource "bloxone_ipam_address_block" "test_tagged_env_only" {
		address = "10.0.0.0"
		cidr = 16
		space = bloxone_ipam_ip_space.test.id
		tags = {
			environment = "prd"
		}
	}
	`
}
