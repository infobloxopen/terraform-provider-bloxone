package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"golang.org/x/exp/slices"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHaGroupDataSource_Filters(t *testing.T) {
	t.Skip("Skipping this test as there is not enough DHCP hosts to run this test")

	dataSourceName := "data.bloxone_dhcp_ha_groups.test"
	resourceName := "bloxone_dhcp_ha_group.test"
	var v ipam.IpamsvcHAGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigFilters(acctest.RandomNameWithPrefix("test_ha")),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName, false)...)...,
				),
			},
		},
	})
}

func TestAccHaGroupDataSource_TagFilters(t *testing.T) {
	t.Skip("Skipping this test as there is not enough DHCP hosts to run this test")

	dataSourceName := "data.bloxone_dhcp_ha_groups.test"
	resourceName := "bloxone_dhcp_ha_group.test"
	var v ipam.IpamsvcHAGroup
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigTagFilters(acctest.RandomNameWithPrefix("test_ha"), acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName, false)...)...,
				),
			},
		},
	})
}

func TestAccHaGroupDataSource_CollectStats(t *testing.T) {
	t.Skip("Skipping this test as there is not enough DHCP hosts to run this test")

	dataSourceName := "data.bloxone_dhcp_ha_groups.test"
	resourceName := "bloxone_dhcp_ha_group.test"
	var v ipam.IpamsvcHAGroup
	name := acctest.RandomNameWithPrefix("test_ha_stats")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigCollectStats(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName, true)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName string, collectStats bool) []resource.TestCheckFunc {
	var testFunc []resource.TestCheckFunc
	if collectStats {
		testFunc = []resource.TestCheckFunc{
			resource.TestCheckResourceAttrWith(dataSourceName, "results.0.status", func(value string) error {
				status := []string{"ok", "failure", "degraded", "intermediate", "unreachable", "unknown"}
				if !slices.Contains(status, value) {
					return fmt.Errorf("status not valid")
				}
				return nil
			}),
			resource.TestCheckResourceAttrWith(dataSourceName, "results.0.hosts.0.state", testAccCheckHAState),
			resource.TestCheckResourceAttr(dataSourceName, "results.0.hosts.0.heartbeats.#", "1"),
			resource.TestCheckResourceAttrSet(dataSourceName, "results.0.hosts.0.heartbeats.0.peer"),
			resource.TestCheckResourceAttrSet(dataSourceName, "results.0.hosts.0.heartbeats.0.successful_heartbeat"),

			resource.TestCheckResourceAttrWith(dataSourceName, "results.0.hosts.1.state", testAccCheckHAState),
			resource.TestCheckResourceAttr(dataSourceName, "results.0.hosts.1.heartbeats.#", "1"),
			resource.TestCheckResourceAttrSet(dataSourceName, "results.0.hosts.1.heartbeats.0.peer"),
			resource.TestCheckResourceAttrSet(dataSourceName, "results.0.hosts.1.heartbeats.0.successful_heartbeat"),
		}

	}
	testFunc = append(testFunc, []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "anycast_config_id", dataSourceName, "results.0.anycast_config_id"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "hosts", dataSourceName, "results.0.hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_space", dataSourceName, "results.0.ip_space"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "results.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}...)

	return testFunc
}

func testAccHaGroupDataSourceConfigFilters(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_ha_group" "test" {
  hosts = [
	{
		host = data.bloxone_dhcp_hosts.test.results.0.id
		role = "active"
	},
	{
		host = data.bloxone_dhcp_hosts.test.results.1.id
		role = "active"
	}
  ]
  name = %q
  mode = "active-active"
}

data "bloxone_dhcp_ha_groups" "test" {
  filters = {
	name = %q
  }
  depends_on = [bloxone_dhcp_ha_group.test]
}`, name, name)
	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupDataSourceConfigTagFilters(name, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_ha_group" "test" {
  hosts = [
	{
		host = data.bloxone_dhcp_hosts.test.results.0.id
		role = "active"
	},
	{
		host = data.bloxone_dhcp_hosts.test.results.1.id
		role = "passive"
	}
  ]
  name = %q
  mode = "active-passive"
  tags = {
	tag1 = %q
  }
}

data "bloxone_dhcp_ha_groups" "test" {
  tag_filters = {
	tag1 = bloxone_dhcp_ha_group.test.tags.tag1
  }
}`, name, tagValue)

	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccHaGroupDataSourceConfigCollectStats(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_ha_group" "test" {
  hosts = [
	{
		host = data.bloxone_dhcp_hosts.test.results.0.id
		role = "active"
	},
	{
		host = data.bloxone_dhcp_hosts.test.results.1.id
		role = "passive"
	}
  ]
  name = %q
  mode = "active-passive"
}

data "bloxone_dhcp_ha_groups" "test" {
  filters = {
	name = %q
  }
  collect_stats = true
  depends_on = [bloxone_dhcp_ha_group.test]
}`, name, name)

	return strings.Join([]string{acctest.TestAccBaseConfig_DhcpHosts(), config}, "")
}

func testAccCheckHAState(state string) error {
	states := []string{"ready", "waiting", "terminated", "syncing", "passive-backup", "partner-down",
		"load-balancing", "backup", "hot-standby", "down", "unreachable",
		"updates-interrupted", "unknown"}

	if !slices.Contains(states, state) {
		return fmt.Errorf("state not valid, value is %s", state)
	}
	return nil
}
