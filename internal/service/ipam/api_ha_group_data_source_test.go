package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHaGroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ha_groups.test"
	resourceName := "bloxone_ipam_ha_group.test"
	var v ipam.IpamsvcHAGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigFilters(acctest.RandomNameWithPrefix("test_ha")),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccHaGroupDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ha_groups.test"
	resourceName := "bloxone_ipam_ha_group.test"
	var v ipam.IpamsvcHAGroup
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHaGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHaGroupDataSourceConfigTagFilters(acctest.RandomNameWithPrefix("test_ha"), "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckHaGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckHaGroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "anycast_config_id", dataSourceName, "results.0.anycast_config_id"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "hosts", dataSourceName, "results.0.hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_space", dataSourceName, "results.0.ip_space"),
		resource.TestCheckResourceAttrPair(resourceName, "mode", dataSourceName, "results.0.mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "results.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccHaGroupDataSourceConfigFilters(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test" {
  hosts = [
	{
		host = data.bloxone_dhcp_hosts.test_01.results.0.id
		role = "active"
	},
	{
		host = data.bloxone_dhcp_hosts.test_02.results.0.id
		role = "active"
	}
  ]
  name = %q
  mode = "active-active"
}

data "bloxone_ipam_ha_groups" "test" {
  filters = {
	name = %q
  }
  depends_on = [bloxone_ipam_ha_group.test]
}`, name, name)
	return strings.Join([]string{acctest.TestAccDhcpHosts("", ""), config}, "")
}

func testAccHaGroupDataSourceConfigTagFilters(name, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ha_group" "test" {
  hosts = [
	{
		host = data.bloxone_dhcp_hosts.test_01.results.0.id
		role = "active"
	},
	{
		host = data.bloxone_dhcp_hosts.test_02.results.0.id
		role = "passive"
	}
  ]
  name = %q
  mode = "active-passive"
  tags = {
	tag1 = %q
  }
}

data "bloxone_ipam_ha_groups" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_ha_group.test.tags.tag1
  }
}`, name, tagValue)

	return strings.Join([]string{acctest.TestAccDhcpHosts("", ""), config}, "")
}
