package redirect_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/redirect"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccCustomRedirectsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_custom_redirects.test"
	resourceName := "bloxone_td_custom_redirect.test"
	var v redirect.CustomRedirect
	name := acctest.RandomNameWithPrefix("custom-redirect")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCustomRedirectsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCustomRedirectsDataSourceConfigFilters(name, "156.2.3.10"),
				Check: resource.ComposeTestCheckFunc(
					[]resource.TestCheckFunc{
						testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
						testAccCheckCustomRedirectsResourceAttrPairWithIndexInOutput(resourceName, dataSourceName, "index"),
					}...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckCustomRedirectsResourceAttrPair(resourceName, dataSourceName string, index int) []resource.TestCheckFunc {
	resultKey := fmt.Sprintf("results.%d.", index)
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, resultKey+"created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "data", dataSourceName, resultKey+"data"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, resultKey+"id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, resultKey+"name"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_ids", dataSourceName, resultKey+"policy_ids"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_names", dataSourceName, resultKey+"policy_names"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, resultKey+"updated_time"),
	}
}

func testAccCheckCustomRedirectsResourceAttrPairWithIndexInOutput(resourceName, dataSourceName, outputName string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Outputs[outputName]
		if !ok {
			return fmt.Errorf("not found: %s", outputName)
		}
		if rs.Type != "string" {
			return fmt.Errorf("expected output %q to be of type string, got %q", outputName, rs.Type)
		}
		index, err := strconv.Atoi(rs.Value.(string))
		if err != nil {
			return fmt.Errorf("failed to parse output %q: %v", outputName, err)
		}
		return resource.ComposeTestCheckFunc(
			testAccCheckCustomRedirectsResourceAttrPair(resourceName, dataSourceName, index)...,
		)(state)
	}
}

func testAccCustomRedirectsDataSourceConfigFilters(name, data string) string {
	return fmt.Sprintf(`
resource "bloxone_td_custom_redirect" "test" {
	name = %q
	data = %q
}
data "bloxone_td_custom_redirects" "test" {
	depends_on = [bloxone_td_custom_redirect.test]
}
output "index" {
	value = index(data.bloxone_td_custom_redirects.test.results[*].name, bloxone_td_custom_redirect.test.name)
}
`, name, data)
}
