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

// below all TestAcc functions

func testAccCheckNetworkListsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "items", dataSourceName, "results.0.items"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_id", dataSourceName, "results.0.policy_id"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
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
