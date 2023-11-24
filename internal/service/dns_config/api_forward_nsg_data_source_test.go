package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccForwardNsgDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_forward_nsgs.test"
	resourceName := "bloxone_dns_forward_nsg.test"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardNsgDataSourceConfigFilters("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardNsgResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccForwardNsgDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_forward_nsgs.test"
	resourceName := "bloxone_dns_forward_nsg.test"
	var v dns_config.ConfigForwardNSG
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardNsgDataSourceConfigTagFilters("NAME_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardNsgResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckForwardNsgResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "external_forwarders", dataSourceName, "results.0.external_forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarders_only", dataSourceName, "results.0.forwarders_only"),
		resource.TestCheckResourceAttrPair(resourceName, "hosts", dataSourceName, "results.0.hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_forwarders", dataSourceName, "results.0.internal_forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "nsgs", dataSourceName, "results.0.nsgs"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
	}
}

func testAccForwardNsgDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test" {
  name = %q
}

data "bloxone_dns_forward_nsgs" "test" {
  filters = {
	name = bloxone_dns_forward_nsg.test.name
  }
}
`, name)
}

func testAccForwardNsgDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_forward_nsgs" "test" {
  tag_filters = {
	tag1 = bloxone_dns_forward_nsg.test.tags.tag1
  }
}
`, name, tagValue)
}
