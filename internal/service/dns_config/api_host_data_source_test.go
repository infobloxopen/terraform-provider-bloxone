package dns_config_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccHostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_hosts.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "TF_TEST_HOST_01"),
				),
			},
		},
	})
}

func TestAccHostDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_hosts.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostDataSourceConfigTagFilters(),
				Check: resource.ComposeTestCheckFunc(
					// There can be more than one host with that tag, so we just check that the results is not empty
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.%"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.tags.used_for", "Terraform Provider Acceptance Tests"),
				),
			},
		},
	})
}

func testAccHostDataSourceConfigFilters() string {
	return `
data "bloxone_dns_hosts" "test" {
  filters = {
	name = "TF_TEST_HOST_01"
  }
}
`
}

func testAccHostDataSourceConfigTagFilters() string {
	return `
data "bloxone_dns_hosts" "test" {
  tag_filters = {
	used_for = "Terraform Provider Acceptance Tests"
  }
}
`
}
