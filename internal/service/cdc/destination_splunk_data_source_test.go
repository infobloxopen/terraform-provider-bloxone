package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationSplunkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_splunks.test"
	resourceName := "bloxone_cdc_destination_splunk.test"
	name := acctest.RandomNameWithPrefix("destination-splunk")
	var v cdc.DestinationSplunk

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationSplunkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationSplunkDataSource_TagFilters(t *testing.T) {
	// TODO: the CDC backend does not honor tag_filters on
	// /v1/destinations/splunk. Confirmed via direct curl against the API: a
	// _tfilter query for a tag value that matches nothing returns the full,
	// unfiltered list of every Splunk destination, while the same query on
	// /v1/destinations/splunkcloud and /v1/destinations/http correctly returns
	// nothing. This is a backend bug, not something fixable in the provider.
	// Needs a CDC backend fix before this can be un-skipped.
	t.Skip("CDC backend does not honor tag_filters on /v1/destinations/splunk; see TODO above")

	dataSourceName := "data.bloxone_cdc_destination_splunks.test"
	resourceName := "bloxone_cdc_destination_splunk.test"
	name := acctest.RandomNameWithPrefix("destination-splunk")
	var v cdc.DestinationSplunk

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationSplunkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationSplunkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "indexer_name", dataSourceName, "results.0.indexer_name"),
		resource.TestCheckResourceAttrPair(resourceName, "insecure_mode", dataSourceName, "results.0.insecure_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "output_data_format", dataSourceName, "results.0.output_data_format"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccDestinationSplunkDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
}

data "bloxone_cdc_destination_splunks" "test" {
  filters = {
    name = bloxone_cdc_destination_splunk.test.name
  }
}
`, name)
}

func testAccDestinationSplunkDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
  tags         = {
    tag1 = %[2]q
  }
}

data "bloxone_cdc_destination_splunks" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_destination_splunk.test.tags.tag1
  }
}
`, name, tagValue)
}
