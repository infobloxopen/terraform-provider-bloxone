package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAuthNsgDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_auth_nsgs.test"
	resourceName := "bloxone_dns_auth_nsg.test"
	var v dns_config.ConfigAuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthNsgDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					}, testAccCheckAuthNsgResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAuthNsgDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_auth_nsgs.test"
	resourceName := "bloxone_dns_auth_nsg.test"
	var v dns_config.ConfigAuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthNsgDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					}, testAccCheckAuthNsgResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAuthNsgResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "external_primaries", dataSourceName, "results.0.external_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "external_secondaries", dataSourceName, "results.0.external_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_secondaries", dataSourceName, "results.0.internal_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nsgs", dataSourceName, "results.0.nsgs"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
	}
}

func testAccAuthNsgDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test" {
  name = %q
}

data "bloxone_dns_auth_nsgs" "test" {
  filters = {
	name = bloxone_dns_auth_nsg.test.name
  }
}
`, name)
}

func testAccAuthNsgDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_auth_nsgs" "test" {
  tag_filters = {
	tag1 = bloxone_dns_auth_nsg.test.tags.tag1
  }
}
`, name, tagValue)
}
