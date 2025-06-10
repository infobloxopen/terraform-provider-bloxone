package ipam_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

var randomTag = acctest.RandomNameWithPrefix("prd")

func TestDataSourceNextAvailableIP_AddressBlock(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_ips.test"
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
                    data "bloxone_ipam_next_available_ips" "test" {
                        id = "/test/address_block/123455678"
                    } 
                `,
				ExpectError: regexp.MustCompile("invalid resource ID specified"),
			},
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 1, "bloxone_ipam_address_block.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.1"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 5, "bloxone_ipam_address_block.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.1", "192.168.0.2"),
					resource.TestCheckResourceAttr(dataSourceName, "results.2", "192.168.0.3"),
					resource.TestCheckResourceAttr(dataSourceName, "results.3", "192.168.0.4"),
					resource.TestCheckResourceAttr(dataSourceName, "results.4", "192.168.0.5"),
				),
			},
		},
	})
}

func TestDataSourceNextAvailableIP_Subnet(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_ips.test"
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 1, "bloxone_ipam_subnet.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.1"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 5, "bloxone_ipam_subnet.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.1", "192.168.0.2"),
					resource.TestCheckResourceAttr(dataSourceName, "results.2", "192.168.0.3"),
					resource.TestCheckResourceAttr(dataSourceName, "results.3", "192.168.0.4"),
					resource.TestCheckResourceAttr(dataSourceName, "results.4", "192.168.0.5"),
				),
			},
		},
	})
}

func TestDataSourceNextAvailableIP_Range(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_ips.test"
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 1, "bloxone_ipam_range.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.15"),
				),
			},
			{
				Config: testAccDataSourceNextAvailableIP(spaceName, 3, "bloxone_ipam_range.test"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_count", dataSourceName, "results.#"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0", "192.168.0.15"),
					resource.TestCheckResourceAttr(dataSourceName, "results.1", "192.168.0.16"),
					resource.TestCheckResourceAttr(dataSourceName, "results.2", "192.168.0.17"),
				),
			},
		},
	})
}

// New test for tag-based IP retrieval
func TestDataSourceNextAvailableIP_ByTags(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_next_available_ips.test_tags"
	spaceName := acctest.RandomNameWithPrefix("ip-space-tags")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Test with missing resource_type
			{
				Config: `
			        data "bloxone_ipam_next_available_ips" "test_tags" {
			            tag_filters = {
			                environment = "test"
			            }
			        }
			    `,
				ExpectError: regexp.MustCompile(`Attribute "resource_type" must be specified when "tag_filters" is specified`),
			},
			//Test getting IPs from address blocks by tags
			{
				Config: testAccDataSourceNextAvailableIPByTags(spaceName, 1, "address_block"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_type", "address_block"),
				),
			},
			// Test getting IPs from subnets by tags
			{
				Config: testAccDataSourceNextAvailableIPByTags(spaceName, 2, "subnet"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_type", "subnet"),
				),
			},
			// Test getting IPs from ranges by tags
			{
				Config: testAccDataSourceNextAvailableIPByTags(spaceName, 3, "range"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "3"),
					resource.TestCheckResourceAttr(dataSourceName, "resource_type", "range"),
				),
			},
		},
	})
}

func testAccDataSourceNextAvailableIPBaseConfig(spaceName string) string {
	return fmt.Sprintf(`
    resource "bloxone_ipam_ip_space" "test" {
        name = %q
    }
    resource "bloxone_ipam_address_block" "test" {
        address = "192.168.0.0"
        cidr = "16"
        space = bloxone_ipam_ip_space.test.id
    }
    resource "bloxone_ipam_subnet" "test" {
        address = "192.168.0.0"
        cidr = "24"
        space = bloxone_ipam_ip_space.test.id
    }
    resource "bloxone_ipam_range" "test" {
        start = "192.168.0.15"
        end = "192.168.0.30"
        space = bloxone_ipam_ip_space.test.id
        depends_on = [bloxone_ipam_subnet.test]
    }
`, spaceName)
}

func testAccDataSourceNextAvailableIPBaseConfigWithTags(spaceName string) string {
	return fmt.Sprintf(`
    resource "bloxone_ipam_ip_space" "test_ipspace" {
        name = %q
    }
    resource "bloxone_ipam_address_block" "test_ab_tags" {
        address = "10.0.0.0"
        cidr = "16"
        space = bloxone_ipam_ip_space.test_ipspace.id
        tags = {
            environment = %q
            purpose = "terraform-testing"
        }
    }
    resource "bloxone_ipam_subnet" "test_subnet_tags" {
        address = "10.0.0.0"
        cidr = "24"
        space = bloxone_ipam_ip_space.test_ipspace.id
        tags = {
            environment = %q
            purpose = "terraform-testing"
        }
		depends_on = [bloxone_ipam_address_block.test_ab_tags]
    }
    resource "bloxone_ipam_range" "test_range_tags" {
        start = "10.0.0.15"
        end = "10.0.0.30"
        space = bloxone_ipam_ip_space.test_ipspace.id
        tags = {
            environment = %q
            purpose = "terraform-testing"
        }
        depends_on = [bloxone_ipam_subnet.test_subnet_tags]
    }
`, spaceName, randomTag, randomTag, randomTag)
}

func testAccDataSourceNextAvailableIP(spaceName string, count int, id string) string {
	var config string
	if count == 1 {
		config = fmt.Sprintf(`
    data "bloxone_ipam_next_available_ips" "test" {
        id = %s.id
    }`, id)
	} else {
		config = fmt.Sprintf(`
    data "bloxone_ipam_next_available_ips" "test" {
        id = %s.id
        ip_count = %d
    }`, id, count)
	}

	return strings.Join([]string{testAccDataSourceNextAvailableIPBaseConfig(spaceName), config}, "")
}

func testAccDataSourceNextAvailableIPByTags(spaceName string, count int, resourceType string) string {
	config := fmt.Sprintf(`
    data "bloxone_ipam_next_available_ips" "test_tags" {
        tag_filters = {
            environment = %q
            purpose = "terraform-testing"
        }
        resource_type = %q
        ip_count = %d
        depends_on = [
            bloxone_ipam_address_block.test_ab_tags,
            bloxone_ipam_subnet.test_subnet_tags,
            bloxone_ipam_range.test_range_tags
        ]
    }`, randomTag, resourceType, count)

	return strings.Join([]string{testAccDataSourceNextAvailableIPBaseConfigWithTags(spaceName), config}, "")
}
