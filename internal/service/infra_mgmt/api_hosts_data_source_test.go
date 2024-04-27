package infra_mgmt_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/inframgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHostsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_hosts.test"
	resourceName := "bloxone_infra_host.test"
	var v inframgmt.Host
	name := acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHostsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHostsDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHostsExists(context.Background(), resourceName, &v),
					}, testAccCheckHostsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccHostsDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_hosts.test"
	resourceName := "bloxone_infra_host.test"
	var v inframgmt.Host
	name := acctest.RandomNameWithPrefix("host")
	tagValue := acctest.RandomNameWithPrefix("tag-value")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHostsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHostsDataSourceConfigTagFilters(name, tagValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHostsExists(context.Background(), resourceName, &v),
					}, testAccCheckHostsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHostsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "configs", dataSourceName, "results.0.configs"),
		resource.TestCheckResourceAttrPair(resourceName, "connectivity_monitor", dataSourceName, "results.0.connectivity_monitor"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "created_by", dataSourceName, "results.0.created_by"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "display_name", dataSourceName, "results.0.display_name"),
		resource.TestCheckResourceAttrPair(resourceName, "host_subtype", dataSourceName, "results.0.host_subtype"),
		resource.TestCheckResourceAttrPair(resourceName, "host_type", dataSourceName, "results.0.host_type"),
		resource.TestCheckResourceAttrPair(resourceName, "host_version", dataSourceName, "results.0.host_version"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_address", dataSourceName, "results.0.ip_address"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_space", dataSourceName, "results.0.ip_space"),
		resource.TestCheckResourceAttrPair(resourceName, "legacy_id", dataSourceName, "results.0.legacy_id"),
		resource.TestCheckResourceAttrPair(resourceName, "location_id", dataSourceName, "results.0.location_id"),
		resource.TestCheckResourceAttrPair(resourceName, "mac_address", dataSourceName, "results.0.mac_address"),
		resource.TestCheckResourceAttrPair(resourceName, "maintenance_mode", dataSourceName, "results.0.maintenance_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "nat_ip", dataSourceName, "results.0.nat_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "noa_cluster", dataSourceName, "results.0.noa_cluster"),
		resource.TestCheckResourceAttrPair(resourceName, "ophid", dataSourceName, "results.0.ophid"),
		resource.TestCheckResourceAttrPair(resourceName, "pool_id", dataSourceName, "results.0.pool_id"),
		resource.TestCheckResourceAttrPair(resourceName, "serial_number", dataSourceName, "results.0.serial_number"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "timezone", dataSourceName, "results.0.timezone"),
	}
}

func testAccHostsDataSourceConfigFilters(displayName string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test" {
  display_name = %q
}

data "bloxone_infra_hosts" "test" {
  filters = {
	display_name = bloxone_infra_host.test.display_name
  }
}
`, displayName)
}

func testAccHostsDataSourceConfigTagFilters(displayName string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test" {
  display_name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_infra_hosts" "test" {
  tag_filters = {
	tag1 = bloxone_infra_host.test.tags.tag1
  }
}
`, displayName, tagValue)
}
