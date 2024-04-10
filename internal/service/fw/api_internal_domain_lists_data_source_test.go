package fw_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccInternalDomainListsDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_internal_domain_lists.test"
	resourceName := "bloxone_td_internal_domain_list.test"
	var v fw.AtcfwInternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalDomainListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccInternalDomainListsDataSourceConfigFilters(name, "example.somedomain.com"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckInternalDomainListsExists(context.Background(), resourceName, &v),
					}, testAccCheckInternalDomainListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccInternalDomainListsDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_internal_domain_lists.test"
	resourceName := "bloxone_td_internal_domain_list.test"
	var v fw.AtcfwInternalDomains
	var name = acctest.RandomNameWithPrefix("td-internal_domain_list")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckInternalDomainListsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccInternalDomainListsDataSourceConfigTagFilters(name, "example.somedomain.com", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckInternalDomainListsExists(context.Background(), resourceName, &v),
					}, testAccCheckInternalDomainListsResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckInternalDomainListsResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "created_time", dataSourceName, "results.0.created_time"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_time", dataSourceName, "results.0.updated_time"),
	}
}

func testAccInternalDomainListsDataSourceConfigFilters(name, internalDomains string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test" {
	name = %q
	internal_domains = [%q]
}

data "bloxone_td_internal_domain_lists" "test" {
  filters = {
	name = bloxone_td_internal_domain_list.test.name
  }
}
`, name, internalDomains)
}

func testAccInternalDomainListsDataSourceConfigTagFilters(name, internalDomains, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_internal_domain_list" "test" {
	name = %q
	internal_domains = [%q]
	tags = {
		tag1 = %q
	  }
}

data "bloxone_td_internal_domain_lists" "test" {
  tag_filters = {
	tag1 = bloxone_td_internal_domain_list.test.tags.tag1
  }
}
`, name, internalDomains, tagValue)
}
