package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOptionGroupDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_option_groups.test"
	resourceName := "bloxone_dhcp_option_group.test"
	var v ipam.OptionGroup

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionGroupDataSourceConfigFilters(acctest.RandomNameWithPrefix("option-group"), "ip4"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccOptionGroupDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_option_groups.test"
	resourceName := "bloxone_dhcp_option_group.test"
	var v ipam.OptionGroup
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionGroupDataSourceConfigTagFilters(acctest.RandomNameWithPrefix("option-group"), "ip6", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionGroupResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckOptionGroupResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccOptionGroupDataSourceConfigFilters(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test" {
  name = %q
  protocol = %q
}

data "bloxone_dhcp_option_groups" "test" {
  filters = {
	name = bloxone_dhcp_option_group.test.name
  }
}
`, name, protocol)
}

func testAccOptionGroupDataSourceConfigTagFilters(name, protocol, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test" {
  name = %q
  protocol = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dhcp_option_groups" "test" {
  tag_filters = {
	tag1 = bloxone_dhcp_option_group.test.tags.tag1
  }
}
`, name, protocol, tagValue)
}
