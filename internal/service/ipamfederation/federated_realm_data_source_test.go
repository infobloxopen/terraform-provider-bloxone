package ipamfederation_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccFederatedRealmDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_federated_realms.test"
	resourceName := "bloxone_federated_realm.test"
	var v ipamfederation.FederatedRealm

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedRealmDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedRealmDataSourceConfigFilters("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedRealmResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFederatedRealmDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_federated_realms.test"
	resourceName := "bloxone_federated_realm.test"
	var v ipamfederation.FederatedRealm
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedRealmDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedRealmDataSourceConfigTagFilters("NAME_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					}, testAccCheckFederatedRealmResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFederatedRealmResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "allocation_v4", dataSourceName, "results.0.allocation_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccFederatedRealmDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test" {
  name = %q
}

data "bloxone_federated_realms" "test" {
  filters = {
	name = bloxone_federated_realm.test.name
  }
}
`, name)
}

func testAccFederatedRealmDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_federated_realms" "test" {
  tag_filters = {
	tag1 = bloxone_federated_realm.test.tags.tag1
  }
}
`, name, tagValue)
}
