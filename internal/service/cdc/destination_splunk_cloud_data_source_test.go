package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationSplunkCloudDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_splunk_clouds.test"
	resourceName := "bloxone_cdc_destination_splunk_cloud.test"
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")
	var v cdc.DestinationSplunkCloud

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkCloudDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkCloudDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationSplunkCloudResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationSplunkCloudDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_splunk_clouds.test"
	resourceName := "bloxone_cdc_destination_splunk_cloud.test"
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")
	var v cdc.DestinationSplunkCloud

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkCloudDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkCloudDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationSplunkCloudResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationSplunkCloudResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "config_filename", dataSourceName, "results.0.config_filename"),
		resource.TestCheckResourceAttrPair(resourceName, "config_option", dataSourceName, "results.0.config_option"),
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

func testAccDestinationSplunkCloudDataSourceConfigFilters(name string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
}

data "bloxone_cdc_destination_splunk_clouds" "test" {
  filters = {
    name = bloxone_cdc_destination_splunk_cloud.test.name
  }
}
`, name)
}

func testAccDestinationSplunkCloudDataSourceConfigTagFilters(name string, tagValue string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  tags           = {
    tag1 = %[2]q
  }
}

data "bloxone_cdc_destination_splunk_clouds" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_destination_splunk_cloud.test.tags.tag1
  }
}
`, name, tagValue)
}
