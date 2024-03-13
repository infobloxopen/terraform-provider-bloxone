package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccIpamHostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_hosts.test"
	resourceName := "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpamHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpamHostDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					}, testAccCheckIpamHostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccIpamHostDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_hosts.test"
	resourceName := "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpamHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpamHostDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					}, testAccCheckIpamHostResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckIpamHostResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "addresses", dataSourceName, "results.0.addresses"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_generate_records", dataSourceName, "results.0.auto_generate_records"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "host_names", dataSourceName, "results.0.host_names"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccIpamHostDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test" {
  name = %q
}

data "bloxone_ipam_hosts" "test" {
  filters = {
	name = bloxone_ipam_host.test.name
  }
}
`, name)
}

func testAccIpamHostDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_ipam_hosts" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_host.test.tags.tag1
  }
}
`, name, tagValue)
}
