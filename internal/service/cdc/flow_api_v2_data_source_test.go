
package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccFlowApiV2DataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_flows.test"
	resourceName := "bloxone_cdc_flow.test"
	var v cdc.FlowDataV2

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2DataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
						}, testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFlowApiV2DataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_flows.test"
	resourceName := "bloxone_cdc_flow.test"
	var v cdc.FlowDataV2
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2DataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
							testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
						}, testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc{
    return []resource.TestCheckFunc{
        resource.TestCheckResourceAttrPair(resourceName, "cdc_services", dataSourceName, "results.0.cdc_services"),
        resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
        resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
        resource.TestCheckResourceAttrPair(resourceName, "destination_ha", dataSourceName, "results.0.destination_ha"),
        resource.TestCheckResourceAttrPair(resourceName, "destinations", dataSourceName, "results.0.destinations"),
        resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
        resource.TestCheckResourceAttrPair(resourceName, "etl_filters", dataSourceName, "results.0.etl_filters"),
        resource.TestCheckResourceAttrPair(resourceName, "filter_expression", dataSourceName, "results.0.filter_expression"),
        resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
        resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
        resource.TestCheckResourceAttrPair(resourceName, "schedule_option", dataSourceName, "results.0.schedule_option"),
        resource.TestCheckResourceAttrPair(resourceName, "script_schedule", dataSourceName, "results.0.script_schedule"),
        resource.TestCheckResourceAttrPair(resourceName, "source", dataSourceName, "results.0.source"),
        resource.TestCheckResourceAttrPair(resourceName, "source_data_types", dataSourceName, "results.0.source_data_types"),
        resource.TestCheckResourceAttrPair(resourceName, "source_log_types", dataSourceName, "results.0.source_log_types"),
        resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
        resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
    }
}

func testAccFlowApiV2DataSourceConfigFilters() string {
	return `
resource "bloxone_cdc_flow" "test" {
}

data "bloxone_cdc_flows" "test" {
  filters = {
	 = bloxone_cdc_flow.test.
  }
}
`
}

func testAccFlowApiV2DataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_cdc_flows" "test" {
  tag_filters = {
	tag1 = bloxone_cdc_flow.test.tags.tag1
  }
}
`,tagValue)
}

