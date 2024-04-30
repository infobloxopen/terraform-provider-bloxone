package dfp_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDfpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dfp_services.test"
	resourceName := "bloxone_dfp_service.test"
	var v dfp.Dfp
	hostName := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDfpDataSourceConfigFilters(hostName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDfpExists(context.Background(), resourceName, &v),
					}, testAccCheckDfpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDfpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarding_policy", dataSourceName, "results.0.forwarding_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "host", dataSourceName, "results.0.host"),
		resource.TestCheckResourceAttrPair(resourceName, "net_addr_policy_ids", dataSourceName, "results.0.net_addr_policy_ids"),
		resource.TestCheckResourceAttrPair(resourceName, "ophid", dataSourceName, "results.0.ophid"),
		resource.TestCheckResourceAttrPair(resourceName, "internal_domain_lists", dataSourceName, "results.0.internal_domain_lists"),
		resource.TestCheckResourceAttrPair(resourceName, "pop_region_id", dataSourceName, "results.0.pop_region_id"),
		resource.TestCheckResourceAttrPair(resourceName, "resolvers_all", dataSourceName, "results.0.resolvers_all"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
		resource.TestCheckResourceAttrPair(resourceName, "elb_ip_list", dataSourceName, "results.0.elb_ip_list"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_id", dataSourceName, "results.0.policy_id"),
		resource.TestCheckResourceAttrPair(resourceName, "service_name", dataSourceName, "results.0.service_name"),
		resource.TestCheckResourceAttrPair(resourceName, "site_id", dataSourceName, "results.0.site_id"),
	}
}

func testAccDfpDataSourceConfigFilters(hostName string) string {
	config := `
resource "bloxone_dfp_service" "test" {
  service_id = bloxone_infra_service.example.id
}
data "bloxone_dfp_services" "test" {
	filters = {
		service_name = bloxone_dfp_service.test.service_name
	}
}
`
	return strings.Join([]string{testAccBaseWithInfraService(hostName), config}, "")
}
