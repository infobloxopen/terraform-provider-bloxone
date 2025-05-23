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

	resource.ParallelTest(t, resource.TestCase{
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
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableSubnet(3, 27),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "subnet_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableSubnetWithSingleTagFilter(24, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableSubnetWithMultipleTagFilters(24, 3),
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

func testAccDataSourceNextAvailableSubnetBaseConfig() string {
	return fmt.Sprintf(`
	resource "bloxone_ipam_ip_space" "test" {
		name = %q
	}
	resource "bloxone_ipam_address_block" "test" {
		name = %q
		address = "192.168.0.0"
		cidr = "24"
		space = bloxone_ipam_ip_space.test.id
	}
`, acctest.RandomNameWithPrefix("nextAvailableIPSpace"), acctest.RandomNameWithPrefix("nextAvailableAB"))
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

// testAccDataSourceNextAvailableSubnetWithSingleTagFilter creates test configuration for next available subnet with a single tag filter
func testAccDataSourceNextAvailableSubnetWithSingleTagFilter(cidr, count int) string {
	config := fmt.Sprintf(`
    data "bloxone_ipam_next_available_subnets" "test" {
        cidr = %d
        subnet_count = %d
        tag_filters = {
            environment = "prd"
        }
      depends_on = [
        "bloxone_ipam_address_block.test_tagged_env_only"
    ]
    }`, cidr, count)

	return strings.Join([]string{testAccDataSourceNextAvailableSubnetWithTagsBaseConfig(), config}, "")
}

// testAccDataSourceNextAvailableSubnetWithMultipleTagFilters creates test configuration for next available subnet with multiple tag filters
func testAccDataSourceNextAvailableSubnetWithMultipleTagFilters(cidr, count int) string {
	config := fmt.Sprintf(`
    data "bloxone_ipam_next_available_subnets" "test" {
        cidr = %d
        subnet_count = %d
        tag_filters = {
            environment = "prd"
            location = "data-center-1"
        }
    depends_on = [
        "bloxone_ipam_address_block.test_tagged_env_only"
    ]
    }`, cidr, count)

	return strings.Join([]string{testAccDataSourceNextAvailableSubnetWithTagsBaseConfig(), config}, "")
}

// testAccDataSourceNextAvailableSubnetWithTagsBaseConfig creates base resources with tags for testing
func testAccDataSourceNextAvailableSubnetWithTagsBaseConfig() string {
	return `
    resource "bloxone_ipam_ip_space" "test" {
        name = "test-acc-next-available-subnets-tags"
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
