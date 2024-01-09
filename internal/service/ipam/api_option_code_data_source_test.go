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
	var v ipam.IpamsvcOptionCode

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionCodeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionCodeDataSourceConfigFilters("234", "test_option_code", "boolean"),
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

func testAccOptionCodeDataSourceConfigFilters(code, name, type_ string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_code" "test" {
  code = %q
  name = %q
  option_space = bloxone_dhcp_option_space.test.id
  type = %q
}

data "bloxone_dhcp_option_codes" "test" {
  filters = {
    name = bloxone_dhcp_option_code.test.name
  }
}
`, code, name, type_)

	return strings.Join([]string{testAccOptionSpace("test_option_space", "ip4"), config}, "")
}
