package fw_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccPopRegionDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_pop_regions.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccPopRegionDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					// check that the results is not empty
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.%"),
				),
			},
		},
	})
}

func testAccPopRegionDataSourceConfigFilters() string {
	return `
data "bloxone_td_pop_regions" "test" {
}
`
}
