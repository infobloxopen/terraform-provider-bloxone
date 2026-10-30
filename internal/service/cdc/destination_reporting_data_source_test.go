package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationReportingDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_reportings.test"
	resourceName := "bloxone_cdc_destination_reporting.test"
	name := acctest.RandomNameWithPrefix("destination-reporting")
	var v cdc.DestinationReporting

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationReportingDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationReportingDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationReportingResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationReportingDataSource_TagFilters(t *testing.T) {
	// TODO: the CDC backend does not honor tag_filters on
	// /v1/destinations/reporting. Confirmed via direct curl against the API: a
	// _tfilter query for a tag value that matches nothing returns the full,
	// unfiltered list of every reporting destination, while the same query on
	// /v1/destinations/syslog correctly returns nothing. This is a backend bug,
	// not something fixable in the provider. Needs a CDC backend fix before
	// this can be un-skipped.
	t.Skip("CDC backend does not honor tag_filters on /v1/destinations/reporting; see TODO above")

	dataSourceName := "data.bloxone_cdc_destination_reportings.test"
	resourceName := "bloxone_cdc_destination_reporting.test"
	name := acctest.RandomNameWithPrefix("destination-reporting")
	var v cdc.DestinationReporting

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationReportingDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationReportingDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationReportingExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationReportingResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationReportingResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "reporting_appliance_address", dataSourceName, "results.0.reporting_appliance_address"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "username", dataSourceName, "results.0.username"),
	}
}

func testAccDestinationReportingDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
}

data "bloxone_cdc_destination_reportings" "test" {
  filters = {
    name = bloxone_cdc_destination_reporting.test.name
  }
}
`, name)
}

func testAccDestinationReportingDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_reporting" "test" {
  name                        = %[1]q
  address                     = "%[1]s.example.com"
  username                    = "admin"
  password                    = "example_password"
  reporting_appliance_address = "10.0.0.1"
  tags                        = {
    tag1 = %[2]q
  }
}

data "bloxone_cdc_destination_reportings" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_destination_reporting.test.tags.tag1
  }
}
`, name, tagValue)
}
