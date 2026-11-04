package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccEtlFilterIpNetworkDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_ip_networks.test"
	resourceName := "bloxone_cdc_etl_filter_ip_network.test"
	name := acctest.RandomNameWithPrefix("etl-filter-ip-network")
	var v cdc.EtlFilterIpNetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterIpNetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterIpNetworkDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterIpNetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterIpNetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccEtlFilterIpNetworkDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_etl_filter_ip_networks.test"
	resourceName := "bloxone_cdc_etl_filter_ip_network.test"
	name := acctest.RandomNameWithPrefix("etl-filter-ip-network")
	var v cdc.EtlFilterIpNetwork

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterIpNetworkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterIpNetworkDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckEtlFilterIpNetworkExists(context.Background(), resourceName, &v),
					}, testAccCheckEtlFilterIpNetworkResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckEtlFilterIpNetworkResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
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

func testAccEtlFilterIpNetworkDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_ip_network" "test" {
  name = %q
  data = ["10.0.0.0/24"]
}

data "bloxone_cdc_etl_filter_ip_networks" "test" {
  filters = {
    name = bloxone_cdc_etl_filter_ip_network.test.name
  }
}
`, name)
}

func testAccEtlFilterIpNetworkDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_ip_network" "test" {
  name = %q
  data = ["10.0.0.0/24"]
  tags = {
    tag1 = %q
  }
}

data "bloxone_cdc_etl_filter_ip_networks" "test" {
  tag_filters = {
    tag1 = bloxone_cdc_etl_filter_ip_network.test.tags.tag1
  }
}
`, name, tagValue)
}
