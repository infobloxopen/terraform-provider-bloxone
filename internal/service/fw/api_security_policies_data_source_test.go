package fw_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccSecurityPoliciesDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_td_security_policies.test"
	resourceName := "bloxone_td_security_policy.test"
	var v fw.AtcfwSecurityPolicy

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPoliciesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPoliciesDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					}, testAccCheckSecurityPoliciesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSecurityPoliciesDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_td_security_policies.test"
	resourceName := "bloxone_td_security_policy.test"
	var v fw.AtcfwSecurityPolicy

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPoliciesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPoliciesDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSecurityPoliciesExists(context.Background(), resourceName, &v),
					}, testAccCheckSecurityPoliciesResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSecurityPoliciesResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{}
}

func testAccSecurityPoliciesDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
}

data "bloxone_td_security_policies" "test" {
  filters = {
	 = bloxone_td_security_policy.test.
  }
}
`)
}

func testAccSecurityPoliciesDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_td_security_policies" "test" {
  tag_filters = {
	tag1 = bloxone_td_security_policy.test.tags.tag1
  }
}
`, tagValue)
}
