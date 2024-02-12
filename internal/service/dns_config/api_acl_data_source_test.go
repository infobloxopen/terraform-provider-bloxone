package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccAclDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_acls.test"
	resourceName := "bloxone_dns_acl.test"
	var v dns_config.ConfigACL

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAclDataSourceConfigFilters("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAclExists(context.Background(), resourceName, &v),
					}, testAccCheckAclResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccAclDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_acls.test"
	resourceName := "bloxone_dns_acl.test"
	var v dns_config.ConfigACL
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAclDataSourceConfigTagFilters("NAME_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckAclExists(context.Background(), resourceName, &v),
					}, testAccCheckAclResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckAclResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "list", dataSourceName, "results.0.list"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
	}
}

func testAccAclDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test" {
  name = %q
}

data "bloxone_dns_acls" "test" {
  filters = {
	name = bloxone_dns_acl.test.name
  }
}
`, name)
}

func testAccAclDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_acls" "test" {
  tag_filters = {
	tag1 = bloxone_dns_acl.test.tags.tag1
  }
}
`, name, tagValue)
}
