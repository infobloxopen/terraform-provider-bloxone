package fw_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNetworkListsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_network_lists.test"
	resourceName := "bloxone_td_network_list.test"
	var v fw.AtcfwNetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkListsDataSourceConfigFilters(name, "156.2.3.0/24"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNetworkListsDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_network_lists.test"
	resourceName := "bloxone_td_network_list.test"
	var v fw.AtcfwNetworkList
	name := acctest.RandomNameWithPrefix("nl")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckNetworkListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkListsDataSourceConfigTagFilters(name, "156.2.3.0/24", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckNetworkListsExists(context.Background(), resourceName, &v),
					}, testAccCheckNetworkListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNetworkListsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{}
}

func testAccNetworkListsDataSourceConfigFilters(name, item string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test" {
	name = %q
	items = [%q]
}

data "bloxone_td_network_lists" "test" {
	filters = {
		name = bloxone_td_network_list.test.name
	}
}
`, name, item)
}

func testAccNetworkListsDataSourceConfigTagFilters(name, item, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_network_list" "test" {
	name = %q
	items = [%q]
	tags = {
		tag1 = %q
	}
}

data "bloxone_td_network_lists" "test" {
  tag_filters = {
	tag1 = bloxone_td_network_list.test.tags.tag1
  }
}
`, name, item, tagValue)
}
