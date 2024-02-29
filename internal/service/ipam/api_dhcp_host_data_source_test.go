package ipam_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDhcpHostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_hosts.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpHostDataSourceConfigFilters("TF_TEST_HOST_01"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "TF_TEST_HOST_01"),
				),
			},
		},
	})
}

func TestAccDhcpHostDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_hosts.test_tag"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpHostDataSourceConfigTagFilters("Terraform Provider Acceptance Tests"),
				Check: resource.ComposeTestCheckFunc(
					// There can be more than one host with that tag, so we just check that the results is not empty
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.%"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.tags.used_for", "Terraform Provider Acceptance Tests"),
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccDhcpHostDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
data "bloxone_dhcp_hosts" "test" {
  filters = {
	 name = %q
  }
}
`, name)
}

func testAccDhcpHostDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
data "bloxone_dhcp_hosts" "test_tag" {
  tag_filters = {
	used_for = %q
  }
}
`, tagValue)
}
