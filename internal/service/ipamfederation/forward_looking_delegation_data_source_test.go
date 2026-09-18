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

func TestAccForwardLookingDelegationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_forward_looking_delegations.test"
	resourceName := "bloxone_federation_forward_looking_delegation.test"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	delegationName := acctest.RandomNameWithPrefix("fld")
	address := "10.90.0.0"
	cidr := 16

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardLookingDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardLookingDelegationDataSourceConfigFilters(realmName, delegationName, address, cidr),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardLookingDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccForwardLookingDelegationDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federation_forward_looking_delegations.test"
	resourceName := "bloxone_federation_forward_looking_delegation.test"
	var v ipamfederation.ForwardLookingDelegation
	realmName := acctest.RandomNameWithPrefix("fld-realm")
	address := "10.91.0.0"
	cidr := 16
	tagValue := "value1"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardLookingDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardLookingDelegationDataSourceConfigTagFilters(realmName, address, cidr, tagValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardLookingDelegationExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardLookingDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckForwardLookingDelegationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "cidr", dataSourceName, "results.0.cidr"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_pool_id", dataSourceName, "results.0.federated_pool_id"),
		resource.TestCheckResourceAttrPair(resourceName, "federated_realms.#", dataSourceName, "results.0.federated_realms.#"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_compliant", dataSourceName, "results.0.network_compliant"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccForwardLookingDelegationDataSourceConfigFilters(realmName, delegationName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test" {
  address          = %q
  cidr             = %d
  federated_realms = [bloxone_federation_federated_realm.test.id]
  name             = %q
}

data "bloxone_federation_forward_looking_delegations" "test" {
  filters = {
    name = bloxone_federation_forward_looking_delegation.test.name
  }
}
`, address, cidr, delegationName)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}

func testAccForwardLookingDelegationDataSourceConfigTagFilters(realmName, address string, cidr int, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_forward_looking_delegation" "test" {
  address          = %q
  cidr             = %d
  federated_realms = [bloxone_federation_federated_realm.test.id]
  tags = {
    tag1 = %q
  }
}

data "bloxone_federation_forward_looking_delegations" "test" {
  tag_filters = {
    tag1 = bloxone_federation_forward_looking_delegation.test.tags.tag1
  }
}
`, address, cidr, tagValue)
	return strings.Join([]string{testAccBaseWithFederatedRealm(realmName), config}, "")
}
