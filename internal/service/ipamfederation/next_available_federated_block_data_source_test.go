package ipamfederation_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNextAvailableFederatedBlockDataSource_basic(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_federated_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFederatedBlockBasic(realmName, "10.132.0.0", 16, 24),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "federated_block_count", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.132\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_count(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_federated_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFederatedBlockCount(realmName, "10.133.0.0", 16, 25, 3),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "federated_block_count", dataSourceName, "results.#"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.133\.`)),
					resource.TestMatchResourceAttr(dataSourceName, "results.1", regexp.MustCompile(`^10\.133\.`)),
					resource.TestMatchResourceAttr(dataSourceName, "results.2", regexp.MustCompile(`^10\.133\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_Name(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_federated_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	name := acctest.RandomNameWithPrefix("federated-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFederatedBlockName(realmName, "10.136.0.0", 16, 24, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", name),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.136\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_Comment(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_federated_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFederatedBlockComment(realmName, "10.137.0.0", 16, 24, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.137\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_federated_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagValue := acctest.RandomNameWithPrefix("next-available-federated-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableFederatedBlockTagFilters(realmName, "10.138.0.0", 16, 24, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.138\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_NeitherIdNorTagFilters(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "bloxone_federation_next_available_federated_blocks" "test" {
    cidr = 24
}
`,
				ExpectError: regexp.MustCompile(`At least one attribute out of \[id,tag_filters\] must be specified`),
			},
		},
	})
}

func TestAccNextAvailableFederatedBlockDataSource_InvalidId(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "bloxone_federation_next_available_federated_blocks" "test" {
    id = "federation/federated_realm/invalid"
    cidr = 24
}
`,
				ExpectError: regexp.MustCompile(`invalid resource ID specified`),
			},
		},
	})
}

func testAccNextAvailableFederatedBlockBasic(federatedRealm string, parentAddress string, parentCidr int, cidr int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_federated_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
}
`, cidr)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccNextAvailableFederatedBlockCount(federatedRealm string, parentAddress string, parentCidr int, cidr int, count int) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_federated_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
    federated_block_count = %d
}
`, cidr, count)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccNextAvailableFederatedBlockName(federatedRealm string, parentAddress string, parentCidr int, cidr int, name string) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_federated_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
    name = %q
}
`, cidr, name)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}

func testAccNextAvailableFederatedBlockTagFilters(federatedRealm string, parentAddress string, parentCidr int, cidr int, tagValue string) string {
	base := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
    tags = {
        env = %q
    }
}
`, parentAddress, parentCidr, tagValue)

	config := fmt.Sprintf(`
data "bloxone_federation_next_available_federated_blocks" "test" {
    tag_filters = {
        env = %q
    }
    cidr = %d

    depends_on = [bloxone_federation_federated_block.test]
}
`, tagValue, cidr)

	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), base, config}, "")
}

func testAccNextAvailableFederatedBlockComment(federatedRealm string, parentAddress string, parentCidr int, cidr int, comment string) string {
	config := fmt.Sprintf(`
data "bloxone_federation_next_available_federated_blocks" "test" {
    id = bloxone_federation_federated_block.test.id
    cidr = %d
    comment = %q
}
`, cidr, comment)
	return strings.Join([]string{testAccFederatedBlockBasicConfig(parentAddress, parentCidr, federatedRealm), config}, "")
}
