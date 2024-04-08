package fw_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNamedListsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_named_lists.test"
	resourceName := "bloxone_td_named_list.test"
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListsDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{}, testAccCheckNamedListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNamedListsDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_named_lists.test"
	resourceName := "bloxone_td_named_list.test"
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListsDataSourceConfigTagFilters(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{}, testAccCheckNamedListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNamedListsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "confidence_level", dataSourceName, "results.0.confidence_level"),
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "item_count", dataSourceName, "results.0.item_count"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policies", dataSourceName, "results.0.policies"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "results.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "threat_level", dataSourceName, "results.0.threat_level"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
}

func testAccNamedListsDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test" {
	name = %q
	items_described = [
	{
		item = "tf-domain.com"
		description = "Exaample Domain"
	}
	]
	type = "custom_list"
	tags = {
		display_name = "Terraform Example Named List"
	}
}

data "bloxone_td_named_lists" "test" {
	tag_filters = {
		display_name = bloxone_td_named_list.test.tags.display_name
	}
}
`, name)
}

func testAccNamedListsDataSourceConfigTagFilters(name, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_named_list" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_td_named_lists" "test" {
  tag_filters = {
	tag1 = bloxone_td_named_list.test.tags.tag1
  }
}
`, tagValue)
}
