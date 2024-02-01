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

    resource.Test(t, resource.TestCase{
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
