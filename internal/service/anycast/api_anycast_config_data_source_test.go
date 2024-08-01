package anycast_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAnycastConfigDataSource_Services(t *testing.T) {
	dataSourceName := "data.bloxone_anycast_configs.test"
	resourceName := "bloxone_anycast_config.test_onprem_hosts"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnycastConfigResourceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAnycastConfigDataSourceConfigService("10.1.1.2", anycastName, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAnycastConfigResourceExists(context.Background(), resourceName, &v),
					}, testAccCheckAnycastConfigResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAnycastConfigDataSource_IsConfigured(t *testing.T) {
	dataSourceName := "data.bloxone_anycast_configs.test"
	resourceName := "bloxone_anycast_config.test_onprem_hosts"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnycastConfigResourceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAnycastConfigDataSourceConfigIsConfigured("10.1.1.2", anycastName, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAnycastConfigResourceExists(context.Background(), resourceName, &v),
					}, testAccCheckAnycastConfigResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAnycastConfigDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_anycast_configs.test"
	resourceName := "bloxone_anycast_config.test"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAnycastConfigResourceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAnycastConfigDataSourceConfigTagFilters("10.1.1.2", anycastName, "DNS", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
						resource.TestCheckResourceAttr(dataSourceName, "results.0.name", anycastName),
						testAccCheckAnycastConfigResourceExists(context.Background(), resourceName, &v),
					}, testAccCheckAnycastConfigResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAnycastConfigResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "account_id", dataSourceName, "results.0.account_id"),
		resource.TestCheckResourceAttrPair(resourceName, "anycast_ip_address", dataSourceName, "results.0.anycast_ip_address"),
		resource.TestCheckResourceAttrPair(resourceName, "anycast_ipv6_address", dataSourceName, "results.0.anycast_ipv6_address"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "fields", dataSourceName, "results.0.fields"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "is_configured", dataSourceName, "results.0.is_configured"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "onprem_hosts", dataSourceName, "results.0.onprem_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "runtime_status", dataSourceName, "results.0.runtime_status"),
		resource.TestCheckResourceAttrPair(resourceName, "service", dataSourceName, "results.0.service"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccAnycastConfigDataSourceConfigService(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
data "bloxone_anycast_configs" "test" {
	service = %q
	depends_on = [bloxone_anycast_config.test]
}
`, anycastIpAddress, name, service, service)
}

func testAccAnycastConfigDataSourceConfigIsConfigured(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_onprem_hosts" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
data "bloxone_anycast_configs" "test" {
	is_configured = false
	depends_on = [bloxone_anycast_config.test_onprem_hosts]
}
`, anycastIpAddress, name, service)
}

func testAccAnycastConfigDataSourceConfigTagFilters(anycastIpAddress, name, service, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test" {
    anycast_ip_address = %q
    name = %q
    service = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_anycast_configs" "test" {
  tag_filters = {
	tag1 = bloxone_anycast_config.test.tags.tag1
  }
}
`, anycastIpAddress, name, service, tagValue)
}
