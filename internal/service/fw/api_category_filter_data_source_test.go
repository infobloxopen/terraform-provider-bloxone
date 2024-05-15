package fw_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccCategoryFiltersDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_category_filters.test"
	resourceName := "bloxone_td_category_filter.test"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCategoryFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCategoryFiltersDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					}, testAccCheckCategoryFiltersResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccCategoryFiltersDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_category_filters.test"
	resourceName := "bloxone_td_category_filter.test"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCategoryFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCategoryFiltersDataSourceConfigTagFilters(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					}, testAccCheckCategoryFiltersResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckCategoryFiltersResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "categories", dataSourceName, "results.0.categories"),
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policies", dataSourceName, "results.0.policies"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
}

func testAccCategoryFiltersDataSourceConfigFilters(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
}

data "bloxone_td_category_filters" "test" {
  filters = {
	name = bloxone_td_category_filter.test.name
  }
}
`, name)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}

func testAccCategoryFiltersDataSourceConfigTagFilters(name, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
	tags = {
		tag1 = %q
	}
}

data "bloxone_td_category_filters" "test" {
	tag_filters = {
		tag1 = bloxone_td_category_filter.test.tags.tag1
	}
}
`, name, tagValue)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}
