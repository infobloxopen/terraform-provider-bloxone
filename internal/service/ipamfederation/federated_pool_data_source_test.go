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
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolDataSourceConfigFilters(realmName, "ip4", "NIOS_X", "us-east-1", poolName),
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
	realmName := acctest.RandomNameWithPrefix("federated-realm")
	poolName := acctest.RandomNameWithPrefix("federated-pool")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedPoolDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedPoolDataSourceConfigTagFilters(realmName, poolName, "ip4", "NIOS_X", "us-east-1", "value1"),
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
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realm", dataSourceName, "results.0.federated_realm"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_compliant", dataSourceName, "results.0.network_compliant"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "provider_type", dataSourceName, "results.0.provider_type"),
		resource.TestCheckResourceAttrPair(resourceName, "region", dataSourceName, "results.0.region"),
		resource.TestCheckResourceAttrPair(resourceName, "state", dataSourceName, "results.0.state"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "results.0.utilization"),
	}
}

func testAccFederatedPoolDataSourceConfigFilters(realmName, protocol, providerType, region, poolName string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    protocol        = %q
    provider_type   = %q
    region          = %q
    name            = %q
}

data "bloxone_federation_federated_pools" "test" {
    filters = {
        name = bloxone_federation_federated_pool.test.name
    }
}
`, protocol, providerType, region, poolName)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccFederatedPoolDataSourceConfigTagFilters(realmName, poolName, protocol, providerType, region, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_federated_pool" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    name            = %q
    protocol        = %q
    provider_type   = %q
    region          = %q
    tags = {
        tag1 = %q
    }
}

data "bloxone_federation_federated_pools" "test" {
    tag_filters = {
        tag1 = bloxone_federation_federated_pool.test.tags.tag1
    }
}
`, poolName, protocol, providerType, region, tagValue)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
