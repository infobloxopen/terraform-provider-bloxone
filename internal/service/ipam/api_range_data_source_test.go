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

func TestAccRangeDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ranges.test"
	resourceName := "bloxone_ipam_range.test"
	var v ipam.IpamsvcRange

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeDataSourceConfigFilters("10.0.0.20", "10.0.0.8"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangeExists(context.Background(), resourceName, &v),
					}, testAccCheckRangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRangeDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_ipam_ranges.test"
	resourceName := "bloxone_ipam_range.test"
	var v ipam.IpamsvcRange
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRangeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRangeDataSourceConfigTagFilters("10.0.0.20", "10.0.0.8", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRangeExists(context.Background(), resourceName, &v),
					}, testAccCheckRangeResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckRangeResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_host", dataSourceName, "results.0.dhcp_host"),
		resource.TestCheckResourceAttrPair(resourceName, "dhcp_options", dataSourceName, "results.0.dhcp_options"),
		resource.TestCheckResourceAttrPair(resourceName, "disable_dhcp", dataSourceName, "results.0.disable_dhcp"),
		resource.TestCheckResourceAttrPair(resourceName, "end", dataSourceName, "results.0.end"),
		resource.TestCheckResourceAttrPair(resourceName, "exclusion_ranges", dataSourceName, "results.0.exclusion_ranges"),
		resource.TestCheckResourceAttrPair(resourceName, "filters", dataSourceName, "results.0.filters"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_assigned_hosts", dataSourceName, "results.0.inheritance_assigned_hosts"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_parent", dataSourceName, "results.0.inheritance_parent"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "results.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "space", dataSourceName, "results.0.space"),
		resource.TestCheckResourceAttrPair(resourceName, "start", dataSourceName, "results.0.start"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "threshold", dataSourceName, "results.0.threshold"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization", dataSourceName, "results.0.utilization"),
		resource.TestCheckResourceAttrPair(resourceName, "utilization_v6", dataSourceName, "results.0.utilization_v6"),
	}
}

func testAccRangeDataSourceConfigFilters(end, start string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test" {
  end = %q
  space = bloxone_ipam_ip_space.test.id
  start = %q
  depends_on = [bloxone_ipam_subnet.test]
}

data "bloxone_ipam_ranges" "test" {
  filters = {
	end = bloxone_ipam_range.test.end
  }
}`, end, start)
	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}

func testAccRangeDataSourceConfigTagFilters(end, start string, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_range" "test" {
  end = %q
  space = bloxone_ipam_ip_space.test.id
  start = %q
  tags = {
	tag1 = %q
  }
  depends_on = [bloxone_ipam_subnet.test]
}

data "bloxone_ipam_ranges" "test" {
  tag_filters = {
	tag1 = bloxone_ipam_range.test.tags.tag1
  }
}`, end, start, tagValue)

	return strings.Join([]string{testAccBaseWithIPSpaceAndSubnet(), config}, "")
}
