
package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationApiDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_syslogs.test"
	resourceName := "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApiDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
						}, testAccCheckDestinationApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDestinationApiDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_destination_syslogs.test"
	resourceName := "bloxone_cdc_destination_syslog.test"
	var v cdc.DestinationSyslog
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApiDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckDestinationApiExists(context.Background(), resourceName, &v),
						}, testAccCheckDestinationApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDestinationApiResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
        resource.TestCheckResourceAttrPair(resourceName, "ca_certificate", dataSourceName, "results.0.ca_certificate"),
        resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
        resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
        resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "insecure_mode", dataSourceName, "results.0.insecure_mode"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "output_data_format", dataSourceName, "results.0.output_data_format"),
        resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "results.0.port"),
        resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
        resource.TestCheckResourceAttrPair(resourceName, "transport_protocol", dataSourceName, "results.0.transport_protocol"),
        resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
    }
}

func testAccDestinationApiDataSourceConfigFilters() string {
	return `
resource "bloxone_cdc_destination_syslog" "test" {
}

data "bloxone_cdc_destination_syslogs" "test" {
  filters = {
	 = bloxone_cdc_destination_syslog.test.
  }
}
`
}

func testAccDestinationApiDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_cdc_destination_syslogs" "test" {
  tag_filters = {
	tag1 = bloxone_cdc_destination_syslog.test.tags.tag1
  }
}
`,tagValue)
}

