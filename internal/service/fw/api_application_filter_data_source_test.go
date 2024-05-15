package fw_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccApplicationFiltersDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_application_filters.test"
	resourceName := "bloxone_td_application_filter.test"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationFiltersDataSourceConfigFilters(name, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					}, testAccCheckApplicationFiltersResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccApplicationFiltersDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_application_filters.test"
	resourceName := "bloxone_td_application_filter.test"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationFiltersDataSourceConfigTagFilters(name, "Microsoft 365", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					}, testAccCheckApplicationFiltersResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckApplicationFiltersResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "criteria", dataSourceName, "results.0.criteria"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policies", dataSourceName, "results.0.policies"),
		resource.TestCheckResourceAttrPair(resourceName, "readonly", dataSourceName, "results.0.readonly"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
}

func testAccApplicationFiltersDataSourceConfigFilters(name, criteriaName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
}

data "bloxone_td_application_filters" "test" {
  filters = {
	name = bloxone_td_application_filter.test.name
  }
}
`, name, criteriaName)
}

func testAccApplicationFiltersDataSourceConfigTagFilters(name, criteriaName, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test" {
	name = %q
	criteria  = [
	{
		name = %q
	}
	]
	tags = {
		tag1 = %q
	}
}

data "bloxone_td_application_filters" "test" {
  tag_filters = {
	tag1 = bloxone_td_application_filter.test.tags.tag1
  }
}
`, name, criteriaName, tagValue)
}
