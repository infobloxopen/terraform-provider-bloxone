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

func TestAccFederatedPoolDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_federated_pools.test"
	resourceName := "bloxone_federation_federated_pool.test"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	name := acctest.RandomNameWithPrefix("federated_pool")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolDataSourceConfigFilters(realmName, name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedPoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFederatedPoolDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_federated_pools.test"
	resourceName := "bloxone_federation_federated_pool.test"
	var v ipamfederation.FederatedPool
	realmName := acctest.RandomNameWithPrefix("federated_pool")
	poolName := acctest.RandomNameWithPrefix("federated_pool")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolDataSourceConfigTagFilters(realmName, poolName, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedPoolExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedPoolResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccCheckFederatedPoolResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realm", dataSourceName, "results.0.federated_realm"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_compliant", dataSourceName, "results.0.network_compliant"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "region", dataSourceName, "results.0.region"),
		resource.TestCheckResourceAttrPair(resourceName, "state", dataSourceName, "results.0.state"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "results.0.utilization"),
	}
}

func testAccFederatedPoolDataSourceConfigFilters(realmName, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
}

data "bloxone_federation_federated_pools" "test" {
    filters = {
        name = bloxone_federation_federated_pool.test.name
    }
}
`, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolDataSourceConfigTagFilters(realmName, poolName, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = "ip4"
    region          = "us-east-1"
    name            = %q
    tags = {
        tag1 = %q
    }
}

data "bloxone_federation_federated_pools" "test" {
    tag_filters = {
        tag1 = bloxone_federation_federated_pool.test.tags.tag1
    }
}
`, poolName, tagValue)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
