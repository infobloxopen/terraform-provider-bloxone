package fw_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccNamedListDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_named_lists.test"
	resourceName := "bloxone_td_named_list.test"
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{}, testAccCheckNamedListResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccNamedListDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_named_lists.test"
	resourceName := "bloxone_td_named_list.test"
	name := acctest.RandomNameWithPrefix("named_list")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNamedListDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{}, testAccCheckNamedListResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckNamedListResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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

func testAccNamedListDataSourceConfigFilters(name string) string {
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

func testAccNamedListDataSourceConfigTagFilters(name, tagValue string) string {
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
`, tagValue)
}
