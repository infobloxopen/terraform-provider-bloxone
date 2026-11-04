package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccEtlFilterFqdnDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_fqdns.test"
	resourceName := "bloxone_cdc_etl_filter_fqdn.test"
	name := acctest.RandomNameWithPrefix("etl-filter-fqdn")
	var v cdc.EtlFilterFqdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterFqdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterFqdnDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterFqdnExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterFqdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccEtlFilterFqdnDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_fqdns.test"
	resourceName := "bloxone_cdc_etl_filter_fqdn.test"
	name := acctest.RandomNameWithPrefix("etl-filter-fqdn")
	var v cdc.EtlFilterFqdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterFqdnDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterFqdnDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterFqdnExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterFqdnResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckEtlFilterFqdnResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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

func testAccEtlFilterFqdnDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_fqdn" "test" {
  name = %q
  data = ["example.com"]
}

data "bloxone_cdc_etl_filter_fqdns" "test" {
  filters = {
    name = bloxone_cdc_etl_filter_fqdn.test.name
  }
}
`, name)
}

func testAccEtlFilterFqdnDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_fqdn" "test" {
  name = %q
  data = ["example.com"]
  tags = {
    tag1 = %q
  }
}

data "bloxone_cdc_etl_filter_fqdns" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_etl_filter_fqdn.test.tags.tag1
  }
}
`, name, tagValue)
}
