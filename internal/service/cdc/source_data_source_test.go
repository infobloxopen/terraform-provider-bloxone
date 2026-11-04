package cdc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

// The built-in Infoblox Cloud Source and Schedule Source exist in every
// account, so these tests need no setup.
func TestAccSourceDataSource_BuiltIn(t *testing.T) {
	for _, tc := range []struct {
		sourceType string
		name       string
	}{
		{"SOURCE_BLOXONE", "Infoblox Cloud Source"},
		{"SOURCE_SCHEDULE", "Schedule Source"},
	} {
		t.Run(tc.sourceType, func(t *testing.T) {
			dataSourceName := "data.bloxone_cdc_sources.test"

			resource.ParallelTest(t, resource.TestCase{
				PreCheck:                 func() { acctest.PreCheck(t) },
				ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: testAccSourceDataSourceConfigType(tc.sourceType),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
							resource.TestCheckResourceAttr(dataSourceName, "results.0.type", tc.sourceType),
							resource.TestCheckResourceAttr(dataSourceName, "results.0.name", tc.name),
							resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
							resource.TestCheckResourceAttrSet(dataSourceName, "results.0.created_at"),
						),
					},
				},
			})
		})
	}
}

func TestAccSourceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_sources.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSourceDataSourceConfigName("Schedule Source"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "Schedule Source"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.type", "SOURCE_SCHEDULE"),
				),
			},
		},
	})
}

func testAccSourceDataSourceConfigType(sourceType string) string {
	return fmt.Sprintf(`
data "bloxone_cdc_sources" "test" {
  filters = {
    type = %q
  }
}
`, sourceType)
}

func testAccSourceDataSourceConfigName(name string) string {
	return fmt.Sprintf(`
data "bloxone_cdc_sources" "test" {
  filters = {
    name = %q
  }
}
`, name)
}
