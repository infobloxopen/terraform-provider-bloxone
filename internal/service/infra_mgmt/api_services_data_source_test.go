package infra_mgmt_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccServicesDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_services.test"
	resourceName := "bloxone_infra_service.test"
	var v infra_mgmt.InfraService
	serviceName := acctest.RandomNameWithPrefix("service")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServicesDataSourceConfigFilters(serviceName, "dhcp"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServicesExists(context.Background(), resourceName, &v),
					}, testAccCheckServicesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccServicesDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_services.test"
	resourceName := "bloxone_infra_service.test"
	var v infra_mgmt.InfraService
	serviceName := acctest.RandomNameWithPrefix("service")
	tagValue := acctest.RandomNameWithPrefix("tag-value")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServicesDataSourceConfigTagFilters(serviceName, "dhcp", tagValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServicesExists(context.Background(), resourceName, &v),
					}, testAccCheckServicesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckServicesResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "configs", dataSourceName, "results.0.configs"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "desired_state", dataSourceName, "results.0.desired_state"),
		resource.TestCheckResourceAttrPair(resourceName, "desired_version", dataSourceName, "results.0.desired_version"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "interface_labels", dataSourceName, "results.0.interface_labels"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "pool_id", dataSourceName, "results.0.pool_id"),
		resource.TestCheckResourceAttrPair(resourceName, "service_type", dataSourceName, "results.0.service_type"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
	}
}

func testAccServicesDataSourceConfigFilters(name, serviceType string) string {
	return strings.Join([]string{
		testAccServicesBase(),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test" {
  name = %q
  pool_id = data.bloxone_infra_hosts.test.results.0.pool_id
  service_type = %q
}

data "bloxone_infra_services" "test" {
  filters = {
	name = bloxone_infra_service.test.name
  }
}
`, name, serviceType),
	}, "")
}

func testAccServicesDataSourceConfigTagFilters(name, serviceType string, tagValue string) string {
	return strings.Join([]string{
		testAccServicesBase(),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test" {
  name = %q
  pool_id = data.bloxone_infra_hosts.test.results.0.pool_id
  service_type = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_infra_services" "test" {
  tag_filters = {
	tag1 = bloxone_infra_service.test.tags.tag1
  }
}
`, name, serviceType, tagValue),
	}, "")
}
