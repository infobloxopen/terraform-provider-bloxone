package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationHttpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_https.test"
	resourceName := "bloxone_cdc_destination_http.test"
	name := acctest.RandomNameWithPrefix("destination-http")
	var v cdc.DestinationHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationHttpDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationHttpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationHttpDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_https.test"
	resourceName := "bloxone_cdc_destination_http.test"
	name := acctest.RandomNameWithPrefix("destination-http")
	var v cdc.DestinationHttp

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationHttpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationHttpDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDestinationHttpExists(context.Background(), resourceName, &v),
					}, testAccCheckDestinationHttpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationHttpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "authentication", dataSourceName, "results.0.authentication"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "credential_id", dataSourceName, "results.0.credential_id"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "insecure_mode", dataSourceName, "results.0.insecure_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "output_data_format", dataSourceName, "results.0.output_data_format"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "transport_protocol", dataSourceName, "results.0.transport_protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccDestinationHttpDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
}

data "bloxone_cdc_destination_https" "test" {
  filters = {
    name = bloxone_cdc_destination_http.test.name
  }
}
`, name)
}

func testAccDestinationHttpDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_http" "test" {
  name          = %[1]q
  address       = "https://%[1]s.example.com"
  credential_id = "`+testAccDestinationHttpCredentialID+`"
  tags          = {
    tag1 = %[2]q
  }
}

data "bloxone_cdc_destination_https" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_destination_http.test.tags.tag1
  }
}
`, name, tagValue)
}
