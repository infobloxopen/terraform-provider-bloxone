package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccEtlFilterOphidDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_ophids.test"
	resourceName := "bloxone_cdc_etl_filter_ophid.test"
	name := acctest.RandomNameWithPrefix("etl-filter-ophid")
	var v cdc.EtlFilterOphid

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterOphidDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterOphidDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterOphidExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterOphidResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccEtlFilterOphidDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_ophids.test"
	resourceName := "bloxone_cdc_etl_filter_ophid.test"
	name := acctest.RandomNameWithPrefix("etl-filter-ophid")
	var v cdc.EtlFilterOphid

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterOphidDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterOphidDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterOphidExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterOphidResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckEtlFilterOphidResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "data", dataSourceName, "results.0.data"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_source", dataSourceName, "results.0.filter_source"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "ingress_policy_action", dataSourceName, "results.0.ingress_policy_action"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccEtlFilterOphidDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_ophid" "test" {
  name = %q
  data = ["abcd1234"]
}

data "bloxone_cdc_etl_filter_ophids" "test" {
  filters = {
    name = bloxone_cdc_etl_filter_ophid.test.name
  }
}
`, name)
}

func testAccEtlFilterOphidDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_ophid" "test" {
  name = %q
  data = ["abcd1234"]
  tags = {
    tag1 = %q
  }
}

data "bloxone_cdc_etl_filter_ophids" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_etl_filter_ophid.test.tags.tag1
  }
}
`, name, tagValue)
}
