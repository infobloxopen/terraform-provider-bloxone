package ipamfederation_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNextAvailableOverlappingBlockDataSource_basic(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_overlapping_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableOverlappingBlockBasic(realmName, "10.130.0.0", 16, 24),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "overlapping_block_count", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.130\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableOverlappingBlockDataSource_count(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_overlapping_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableOverlappingBlockCount(realmName, "10.131.0.0", 16, 25, 3),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "overlapping_block_count", dataSourceName, "results.#"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.131\.`)),
					resource.TestMatchResourceAttr(dataSourceName, "results.1", regexp.MustCompile(`^10\.131\.`)),
					resource.TestMatchResourceAttr(dataSourceName, "results.2", regexp.MustCompile(`^10\.131\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableOverlappingBlockDataSource_InvalidId(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "bloxone_federation_next_available_overlapping_blocks" "test" {
    id = "federation/federated_realm/invalid"
    cidr = 24
}
`,
				ExpectError: regexp.MustCompile(`invalid resource ID specified`),
			},
		},
	})
}

func testAccNextAvailableOverlappingBlockBasic(federatedRealm string, parentAddress string, parentCidr int, cidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_overlapping_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
}
`, cidr)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccNextAvailableOverlappingBlockCount(federatedRealm string, parentAddress string, parentCidr int, cidr int, count int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_overlapping_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
    overlapping_block_count = %d
}
`, cidr, count)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}
