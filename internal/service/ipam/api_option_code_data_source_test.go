package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOptionCodeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_option_codes.test"
	resourceName := "bloxone_dhcp_option_code.test"
	var v ipam.OptionCode
	optionSpaceName := acctest.RandomNameWithPrefix("option-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionCodeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionCodeDataSourceConfigFilters(optionSpaceName, "234", "test_option_code", "boolean", "name"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionCodeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
			{
				Config: testAccOptionCodeDataSourceConfigFilters(optionSpaceName, "234", "test_option_code", "boolean", "code"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionCodeExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionCodeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccCheckOptionCodeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "array", dataSourceName, "results.0.array"),
		resource.TestCheckResourceAttrPair(resourceName, "code", dataSourceName, "results.0.code"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "option_space", dataSourceName, "results.0.option_space"),
		resource.TestCheckResourceAttrPair(resourceName, "source", dataSourceName, "results.0.source"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "results.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccOptionCodeDataSourceConfigFilters(optionSpaceName, code, name, optionItemType string, filterBy string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test" {
  code = %[1]q
  name = %[2]q
  option_space = bloxone_dhcp_option_space.test.id
  type = %[3]q
}

data "bloxone_dhcp_option_codes" "test" {
  filters = {
    %[4]s = bloxone_dhcp_option_code.test.%[4]s
  }
}
`, code, name, optionItemType, filterBy)

	return strings.Join([]string{testAccBaseWithOptionSpace(optionSpaceName, "ip4"), config}, "")
}
