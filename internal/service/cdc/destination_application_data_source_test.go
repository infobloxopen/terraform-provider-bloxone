package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationApplicationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_applications.test"
	resourceName := "bloxone_cdc_destination_application.test"
	name := acctest.RandomNameWithPrefix("destination-application")
	var v cdc.DestinationApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApplicationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApplicationDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationApplicationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationApplicationDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_applications.test"
	resourceName := "bloxone_cdc_destination_application.test"
	name := acctest.RandomNameWithPrefix("destination-application")
	var v cdc.DestinationApplication

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApplicationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApplicationDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationApplicationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationApplicationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "application", dataSourceName, "results.0.application"),
		resource.TestCheckResourceAttrPair(resourceName, "context", dataSourceName, "results.0.context"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "output_data_format", dataSourceName, "results.0.output_data_format"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccDestinationApplicationDataSourceConfigFilters(name string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test" {
  name        = %q
  application = bloxone_cdc_application.test.id
}

data "bloxone_cdc_destination_applications" "test" {
  filters = {
    name = bloxone_cdc_destination_application.test.name
  }
}
`, name)
}

func testAccDestinationApplicationDataSourceConfigTagFilters(name string, tagValue string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test" {
  name        = %q
  application = bloxone_cdc_application.test.id
  tags = {
    tag1 = %q
  }
}

data "bloxone_cdc_destination_applications" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_destination_application.test.tags.tag1
  }
}
`, name, tagValue)
}
