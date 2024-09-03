package ipamfederation_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccFederatedBlockDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_federated_blocks.test"
	resourceName := "bloxone_federated_block.test"
	var v ipamfederation.FederatedBlock

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedBlockDataSourceConfigFilters("FEDERATED_REALM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedBlockResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFederatedBlockDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federated_blocks.test"
	resourceName := "bloxone_federated_block.test"
	var v ipamfederation.FederatedBlock
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedBlockDataSourceConfigTagFilters("FEDERATED_REALM_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedBlockResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFederatedBlockResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "allocation_v4", dataSourceName, "results.0.allocation_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "cidr", dataSourceName, "results.0.cidr"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realm", dataSourceName, "results.0.federated_realm"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccFederatedBlockDataSourceConfigFilters(federatedRealm string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test" {
  federated_realm = %q
}

data "bloxone_federated_blocks" "test" {
  filters = {
	federated_realm = bloxone_federated_block.test.federated_realm
  }
}
`, federatedRealm)
}

func testAccFederatedBlockDataSourceConfigTagFilters(federatedRealm string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test" {
  federated_realm = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_federated_blocks" "test" {
  tag_filters = {
	tag1 = bloxone_federated_block.test.tags.tag1
  }
}
`, federatedRealm, tagValue)
}
