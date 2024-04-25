package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccForwardZoneDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_forward_zones.test"
	resourceName := "bloxone_dns_forward_zone.test"
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dnsconfig.ForwardZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardZoneDataSourceConfigFilters(fqdn),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardZoneResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccForwardZoneDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_forward_zones.test"
	resourceName := "bloxone_dns_forward_zone.test"
	var v dnsconfig.ForwardZone
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardZoneDataSourceConfigTagFilters(fqdn, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					}, testAccCheckForwardZoneResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccCheckForwardZoneResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "external_forwarders", dataSourceName, "results.0.external_forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "forward_only", dataSourceName, "results.0.forward_only"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "results.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "hosts", dataSourceName, "results.0.hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_forwarders", dataSourceName, "results.0.internal_forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "mapped_subnet", dataSourceName, "results.0.mapped_subnet"),
		resource.TestCheckResourceAttrPair(resourceName, "mapping", dataSourceName, "results.0.mapping"),
		resource.TestCheckResourceAttrPair(resourceName, "nsgs", dataSourceName, "results.0.nsgs"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol_fqdn", dataSourceName, "results.0.protocol_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "results.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "warnings", dataSourceName, "results.0.warnings"),
	}
}

func testAccForwardZoneDataSourceConfigFilters(fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test" {
  fqdn = %q
}

data "bloxone_dns_forward_zones" "test" {
  filters = {
	fqdn = bloxone_dns_forward_zone.test.fqdn
  }
}
`, fqdn)
}

func testAccForwardZoneDataSourceConfigTagFilters(fqdn string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test" {
  fqdn = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_forward_zones" "test" {
  tag_filters = {
	tag1 = bloxone_dns_forward_zone.test.tags.tag1
  }
}
`, fqdn, tagValue)
}
