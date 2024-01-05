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

func TestAccOptionSpaceDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_option_spaces.test"
	resourceName := "bloxone_dhcp_option_space.test"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionSpaceDataSourceConfigFilters("test_option_space", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionSpaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccOptionSpaceDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_option_spaces.test"
	resourceName := "bloxone_dhcp_option_space.test"
	var v ipam.IpamsvcOptionSpace
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionSpaceDataSourceConfigTagFilters("test_option_space_tag", "ip6", "space1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					}, testAccCheckOptionSpaceResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckOptionSpaceResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccOptionSpace(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test" {
  name = %q
  protocol = %q
}`, name, protocol)
}

func testAccOptionSpaceMultiple(name1, protocol1, name2, protocol2 string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test1" {
  name = %q
  protocol = %q
}
resource "bloxone_dhcp_option_space" "test2" {
  name = %q
  protocol = %q
}`, name1, protocol1, name2, protocol2)
}

func testAccOptionSpaceDataSourceConfigFilters(name, protocol string) string {
	config := `
data "bloxone_dhcp_option_spaces" "test" {
  filters = {
	name = bloxone_dhcp_option_space.test.name
  }
}
`

	return strings.Join([]string{config, testAccOptionSpace(name, protocol)}, "")
}

func testAccOptionSpaceDataSourceConfigTagFilters(name, protocol, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test" {
  name = %q
  protocol = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dhcp_option_spaces" "test" {
  tag_filters = {
	tag1 = bloxone_dhcp_option_space.test.tags.tag1
  }
}
`, name, protocol, tagValue)
}
