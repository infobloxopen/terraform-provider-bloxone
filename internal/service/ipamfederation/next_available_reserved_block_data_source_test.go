package ipamfederation_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNextAvailableReservedBlockDataSource_basic(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_reserved_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableReservedBlockBasic(realmName, "10.120.0.0", 16, 24),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.cidr", "24"),
				),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_count(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_reserved_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableReservedBlockCount(realmName, "10.121.0.0", 16, 25, 3),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "3"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1.address"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2.address"),
				),
			},
		},
	})
}

func testAccNextAvailableReservedBlockBaseConfig(federatedRealm string, parentAddress string, parentCidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
}
`, parentAddress, parentCidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccNextAvailableReservedBlockBasic(federatedRealm string, parentAddress string, parentCidr int, cidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_reserved_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
}
`, cidr)
	return strings.Join([]string{testAccNextAvailableReservedBlockBaseConfig(federatedRealm, parentAddress, parentCidr), config}, "")
}

func testAccNextAvailableReservedBlockCount(federatedRealm string, parentAddress string, parentCidr int, cidr int, count int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_reserved_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
    reserved_block_count = %d
}
`, cidr, count)
	return strings.Join([]string{testAccNextAvailableReservedBlockBaseConfig(federatedRealm, parentAddress, parentCidr), config}, "")
}
