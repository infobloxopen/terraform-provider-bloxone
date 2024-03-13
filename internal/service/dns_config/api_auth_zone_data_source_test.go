package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAuthZoneDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_auth_zones.test"
	resourceName := "bloxone_dns_auth_zone.test"
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	var v dns_config.ConfigAuthZone

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthZoneDataSourceConfigFilters(zoneFqdn, "cloud"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					}, testAccCheckAuthZoneResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAuthZoneDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_auth_zones.test"
	resourceName := "bloxone_dns_auth_zone.test"
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	var v dns_config.ConfigAuthZone
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthZoneDataSourceConfigTagFilters(zoneFqdn, "cloud", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					}, testAccCheckAuthZoneResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAuthZoneResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "external_primaries", dataSourceName, "results.0.external_primaries"),
		resource.TestCheckResourceAttrPair(resourceName, "external_providers", dataSourceName, "results.0.external_providers"),
		resource.TestCheckResourceAttrPair(resourceName, "external_secondaries", dataSourceName, "results.0.external_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "results.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "gss_tsig_enabled", dataSourceName, "results.0.gss_tsig_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_assigned_hosts", dataSourceName, "results.0.inheritance_assigned_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "initial_soa_serial", dataSourceName, "results.0.initial_soa_serial"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_secondaries", dataSourceName, "results.0.internal_secondaries"),
		resource.TestCheckResourceAttrPair(resourceName, "mapped_subnet", dataSourceName, "results.0.mapped_subnet"),
		resource.TestCheckResourceAttrPair(resourceName, "mapping", dataSourceName, "results.0.mapping"),
		resource.TestCheckResourceAttrPair(resourceName, "notify", dataSourceName, "results.0.notify"),
		resource.TestCheckResourceAttrPair(resourceName, "nsgs", dataSourceName, "results.0.nsgs"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "primary_type", dataSourceName, "results.0.primary_type"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol_fqdn", dataSourceName, "results.0.protocol_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "query_acl", dataSourceName, "results.0.query_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "transfer_acl", dataSourceName, "results.0.transfer_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "update_acl", dataSourceName, "results.0.update_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "use_forwarders_for_subzones", dataSourceName, "results.0.use_forwarders_for_subzones"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "results.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "warnings", dataSourceName, "results.0.warnings"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_authority", dataSourceName, "results.0.zone_authority"),
	}
}

func testAccAuthZoneDataSourceConfigFilters(fqdn, primaryType string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
  fqdn = %q
  primary_type = %q
}

data "bloxone_dns_auth_zones" "test" {
  filters = {
	fqdn = bloxone_dns_auth_zone.test.fqdn
  }
}
`, fqdn, primaryType)
}

func testAccAuthZoneDataSourceConfigTagFilters(fqdn, primaryType string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
  fqdn = %q
  primary_type = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_auth_zones" "test" {
  tag_filters = {
	tag1 = bloxone_dns_auth_zone.test.tags.tag1
  }
}
`, fqdn, primaryType, tagValue)
}
