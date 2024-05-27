package redirect_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

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
					append([]resource.TestCheckFunc{
						testAccCheckCustomRedirectsExists(context.Background(), resourceName, &v),
					}, testAccCheckCustomRedirectsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckCustomRedirectsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "data", dataSourceName, "results.0.data"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_ids", dataSourceName, "results.0.policy_ids"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_names", dataSourceName, "results.0.policy_names"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
}

func testAccCustomRedirectsDataSourceConfigFilters(name, data string) string {
	return fmt.Sprintf(`
 resource "bloxone_td_custom_redirect" "test" {
  name = %q
  data = %q
}

data "bloxone_td_custom_redirects" "test" {
    depends_on = [bloxone_td_custom_redirect.test] #so that data doesn't run before resource is created
}

output "custom_redirects" {
  value = lookup(zipmap(data.bloxone_td_custom_redirects.test.results[*].name, data.bloxone_td_custom_redirects.test.results[*]), bloxone_td_custom_redirect.test.name, null)
}
 `, name, data)
}
