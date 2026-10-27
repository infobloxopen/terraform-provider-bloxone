package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccApplicationApiDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_applications.test"
	resourceName := "bloxone_cdc_application.test"
	name := acctest.RandomNameWithPrefix("application")
	var v cdc.ApplicationData

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationApiDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					}, testAccCheckApplicationApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccApplicationApiDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_applications.test"
	resourceName := "bloxone_cdc_application.test"
	name := acctest.RandomNameWithPrefix("application")
	var v cdc.ApplicationData
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationApiDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					}, testAccCheckApplicationApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckApplicationApiResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "data.scripts.0.filename", dataSourceName, "results.0.data.scripts.0.filename"),
		resource.TestCheckResourceAttrPair(resourceName, "data.scripts.0.content", dataSourceName, "results.0.data.scripts.0.content"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "results.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "variables.var1", dataSourceName, "results.0.variables.var1"),
	}
}

func testAccApplicationApiDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test" {
  name = %q
  data = {
      scripts = [{
          filename = "script_1.py"
          content  = "print(\"hello\")"
      }]
  }
  variables = {
      var1 = "value1"
  }
}

data "bloxone_cdc_applications" "test" {
  filters = {
	name = bloxone_cdc_application.test.name
  }
}
`, name)
}

func testAccApplicationApiDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test" {
  name = %q
  data = {
      scripts = [{
          filename = "script_1.py"
          content  = "print(\"hello\")"
      }]
  }
  tags = {
	tag1 = %q
  }
}

data "bloxone_cdc_applications" "test" {
  tag_filters = {
	tag1 = bloxone_cdc_application.test.tags.tag1
  }
}
`, name, tagValue)
}
