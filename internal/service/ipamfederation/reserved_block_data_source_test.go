package ipamfederation_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func TestAccReservedBlockDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_reserved_blocks.test"
	resourceName := "bloxone_federation_reserved_block.test"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckReservedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccReservedBlockDataSourceConfigFilters(realmName, "10.110.0.0", 24, acctest.RandomNameWithPrefix("reserved-block")),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					}, testAccCheckReservedBlockResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccReservedBlockDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_reserved_blocks.test"
	resourceName := "bloxone_federation_reserved_block.test"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckReservedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccReservedBlockDataSourceConfigTagFilters(realmName, "10.111.0.0", 24, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					}, testAccCheckReservedBlockResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckReservedBlockResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "cidr", dataSourceName, "results.0.cidr"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_pool_id", dataSourceName, "results.0.federated_pool_id"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realm", dataSourceName, "results.0.federated_realm"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "metadata", dataSourceName, "results.0.metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_compliant", dataSourceName, "results.0.network_compliant"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "region", dataSourceName, "results.0.region"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "tags_all", dataSourceName, "results.0.tags_all"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccReservedBlockDataSourceConfigFilters(federatedRealm string, address string, cidr int, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test" {
  federated_realm = bloxone_federation_federated_realm.test.id
  address = %q
  cidr = %d
  name = %q
}

data "bloxone_federation_reserved_blocks" "test" {
  filters = {
	name = bloxone_federation_reserved_block.test.name
  }
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockDataSourceConfigTagFilters(federatedRealm string, address string, cidr int, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test" {
  federated_realm = bloxone_federation_federated_realm.test.id
  address = %q
  cidr = %d
  tags = {
	tag1 = %q
  }
}

data "bloxone_federation_reserved_blocks" "test" {
  tag_filters = {
	tag1 = bloxone_federation_reserved_block.test.tags.tag1
  }
}
`, address, cidr, tagValue)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}
