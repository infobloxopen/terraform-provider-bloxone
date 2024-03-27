package fw_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAccessCodesDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_access_codes.test"
	resourceName := "bloxone_td_access_code.test"
	var v fw.AtcfwAccessCode
	name := acctest.RandomNameWithPrefix("ac")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAccessCodesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAccessCodesDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAccessCodesExists(context.Background(), resourceName, &v),
					}, testAccCheckAccessCodesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAccessCodesResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "access_key", dataSourceName, "results.0.access_key"),
		resource.TestCheckResourceAttrPair(resourceName, "activation", dataSourceName, "results.0.activation"),
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "expiration", dataSourceName, "results.0.expiration"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "rules", dataSourceName, "results.0.rules"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
		resource.TestCheckResourceAttrPair(resourceName, "policy_ids", dataSourceName, "results.0.policy_ids"),
	}
}

func testAccAccessCodesDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_access_code" "test" {
	name = %[1]q
	activation = %[2]q
	expiration = %[3]q
	rules = [
		{
			action = "" ,
			data = "antimalware",
			description = "",
			redirect_name = "",
			type = "named_feed"
		}
	]
}

data "bloxone_td_access_codes" "test" {
  filters = {
	 name = bloxone_td_access_code.test.name
  }
}
`, name, time.Now().UTC().Format(time.RFC3339), time.Now().UTC().Add(time.Hour).Format(time.RFC3339))
}
