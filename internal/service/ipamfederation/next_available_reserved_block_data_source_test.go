package ipamfederation_test

import (
	"fmt"
	"regexp"
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
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
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
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.2"),
				),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_reserved_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagValue := acctest.RandomNameWithPrefix("next-available-reserved-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableReservedBlockTagFilters(realmName, "10.122.0.0", 16, 24, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
				),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_TagFiltersSkipsFullParent(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_reserved_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagValue := acctest.RandomNameWithPrefix("next-available-reserved-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableReservedBlockTagFiltersSkipsFullParent(realmName, "10.144.0.0", "10.144.0.128", "10.145.0.0", 24, 25, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestMatchResourceAttr(dataSourceName, "results.0", regexp.MustCompile(`^10\.145\.`)),
				),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_TagFiltersAllParentsFull(t *testing.T) {
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagValue := acctest.RandomNameWithPrefix("next-available-reserved-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccNextAvailableReservedBlockTagFiltersAllParentsFull(realmName, "10.146.0.0", "10.147.0.0", 24, 25, tagValue),
				ExpectError: regexp.MustCompile(`Insufficient Available Reserved Blocks`),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_TagFiltersCountSpansParents(t *testing.T) {
	dataSourceName := "data.bloxone_federation_next_available_reserved_blocks.test"
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	tagValue := acctest.RandomNameWithPrefix("next-available-reserved-block")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNextAvailableReservedBlockTagFiltersCountSpansParents(realmName, "10.148.0.0", "10.149.0.0", 24, 25, tagValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "2"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.1"),
				),
			},
		},
	})
}

func TestAccNextAvailableReservedBlockDataSource_NeitherIdNorTagFilters(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
data "bloxone_federation_next_available_reserved_blocks" "test" {
    cidr = 24
}
`,
				ExpectError: regexp.MustCompile(`At least one attribute out of \[id,tag_filters\] must be specified`),
			},
		},
	})
}

func testAccNextAvailableReservedBlockTagFilters(federatedRealm string, parentAddress string, parentCidr int, cidr int, tagValue string) string {
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
data "bloxone_federation_next_available_reserved_blocks" "test" {
    tag_filters = {
        env = %q
    }
    cidr = %d

    depends_on = [bloxone_federation_federated_block.test]
}
`, tagValue, cidr)

	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), base, config}, "")
}

// testAccNextAvailableReservedBlockTagFiltersSkipsFullParent creates two candidate parent Federated
// Blocks sharing the same tag: fullParent's entire address space is consumed by two pre-created
// Reserved Blocks, leaving no room for another block of the requested cidr, while roomyParent is
// left empty. The data source should skip the full parent and find the next available block under
// the roomy one, instead of failing the whole search.
func testAccNextAvailableReservedBlockTagFiltersSkipsFullParent(federatedRealm, fullParentAddress, fullParentFillHighAddress, roomyParentAddress string, parentCidr, cidr int, tagValue string) string {
	base := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "full_parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[1]q
    cidr = %[4]d
    tags = {
        env = %[6]q
    }
}

resource "bloxone_federation_federated_block" "roomy_parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[3]q
    cidr = %[4]d
    tags = {
        env = %[6]q
    }
    # Force full_parent to be created (and thus listed) first, so the search visits it before
    # roomy_parent regardless of API list ordering.
    depends_on = [bloxone_federation_federated_block.full_parent]
}

resource "bloxone_federation_reserved_block" "fill_low" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %[1]q
    cidr            = %[5]d
    depends_on      = [bloxone_federation_federated_block.full_parent]
}

resource "bloxone_federation_reserved_block" "fill_high" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %[2]q
    cidr            = %[5]d
    depends_on      = [bloxone_federation_federated_block.full_parent]
}
`, fullParentAddress, fullParentFillHighAddress, roomyParentAddress, parentCidr, cidr, tagValue)

	config := fmt.Sprintf(`
data "bloxone_federation_next_available_reserved_blocks" "test" {
    tag_filters = {
        env = %[2]q
    }
    cidr = %[1]d

    depends_on = [
        bloxone_federation_federated_block.full_parent,
        bloxone_federation_federated_block.roomy_parent,
        bloxone_federation_reserved_block.fill_low,
        bloxone_federation_reserved_block.fill_high,
    ]
}
`, cidr, tagValue)

	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), base, config}, "")
}

// testAccNextAvailableReservedBlockTagFiltersAllParentsFull creates two candidate parent Federated
// Blocks sharing the same tag, each fully consumed by one Reserved Block covering the whole /24.
// With no candidate having room for another block of the requested cidr, the search should report
// a clean "Insufficient Available Reserved Blocks" error rather than surfacing the raw per-parent
// API error (or crashing).
func testAccNextAvailableReservedBlockTagFiltersAllParentsFull(federatedRealm, parent1Address, parent2Address string, parentCidr, cidr int, tagValue string) string {
	base := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "full_parent_1" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[1]q
    cidr = %[3]d
    tags = {
        env = %[5]q
    }
}

resource "bloxone_federation_federated_block" "full_parent_2" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[2]q
    cidr = %[3]d
    tags = {
        env = %[5]q
    }
}

resource "bloxone_federation_reserved_block" "fill_parent_1" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %[1]q
    cidr            = %[3]d
    depends_on      = [bloxone_federation_federated_block.full_parent_1]
}

resource "bloxone_federation_reserved_block" "fill_parent_2" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %[2]q
    cidr            = %[3]d
    depends_on      = [bloxone_federation_federated_block.full_parent_2]
}
`, parent1Address, parent2Address, parentCidr, cidr, tagValue)

	config := fmt.Sprintf(`
data "bloxone_federation_next_available_reserved_blocks" "test" {
    tag_filters = {
        env = %[2]q
    }
    cidr = %[1]d

    depends_on = [
        bloxone_federation_federated_block.full_parent_1,
        bloxone_federation_federated_block.full_parent_2,
        bloxone_federation_reserved_block.fill_parent_1,
        bloxone_federation_reserved_block.fill_parent_2,
    ]
}
`, cidr, tagValue)

	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), base, config}, "")
}

// testAccNextAvailableReservedBlockTagFiltersCountSpansParents creates one candidate parent with
// exactly one free cidr-sized slot left (its other slot is pre-consumed) and a second, empty
// candidate parent with room to spare. Requesting reserved_block_count = 2 should combine one
// result from each parent, exercising both the per-parent partial retry and the cross-parent
// continuation together.
func testAccNextAvailableReservedBlockTagFiltersCountSpansParents(federatedRealm, partialParentAddress, roomyParentAddress string, parentCidr, cidr int, tagValue string) string {
	base := fmt.Sprintf(`
resource "bloxone_federation_federated_block" "partial_parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[1]q
    cidr = %[3]d
    tags = {
        env = %[5]q
    }
}

resource "bloxone_federation_federated_block" "roomy_parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %[2]q
    cidr = %[3]d
    tags = {
        env = %[5]q
    }
    # Force partial_parent to be created (and thus listed) first, so it is tried before
    # roomy_parent regardless of API list ordering.
    depends_on = [bloxone_federation_federated_block.partial_parent]
}

# Consumes exactly the first half of partial_parent, leaving exactly one cidr-sized slot free.
resource "bloxone_federation_reserved_block" "fill_partial_parent" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address         = %[1]q
    cidr            = %[4]d
    depends_on      = [bloxone_federation_federated_block.partial_parent]
}
`, partialParentAddress, roomyParentAddress, parentCidr, cidr, tagValue)

	config := fmt.Sprintf(`
data "bloxone_federation_next_available_reserved_blocks" "test" {
    tag_filters = {
        env = %[2]q
    }
    cidr = %[1]d
    reserved_block_count = 2

    depends_on = [
        bloxone_federation_federated_block.partial_parent,
        bloxone_federation_federated_block.roomy_parent,
        bloxone_federation_reserved_block.fill_partial_parent,
    ]
}
`, cidr, tagValue)

	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), base, config}, "")
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
