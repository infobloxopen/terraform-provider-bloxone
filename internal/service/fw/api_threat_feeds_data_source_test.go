package fw_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"testing"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccThreatFeedDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_threat_feeds.test"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccThreatFeedDataSourceConfigFilters("Cryptocurrency"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "Cryptocurrency"),
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccThreatFeedDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
data "bloxone_td_threat_feeds" "test" {
  filters = {
	 name = %q
  }
}
`, name)
}
