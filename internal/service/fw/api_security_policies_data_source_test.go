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
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPoliciesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPoliciesDataSourceConfigFilters(name),
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
	name := acctest.RandomNameWithPrefix("sec-policy")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSecurityPoliciesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityPoliciesDataSourceConfigTagFilters(name, acctest.RandomName()),
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

func testAccSecurityPoliciesDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
	name = %q
}

data "bloxone_td_security_policies" "test" {
  filters = {
	name= bloxone_td_security_policy.test.name
  }
}
`, name)
}

func testAccSecurityPoliciesDataSourceConfigTagFilters(name, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_td_security_policy" "test" {
  name = 	%q
  tags = {
	tag1 = %q
  }
}

data "bloxone_td_security_policies" "test" {
  tag_filters = {
	tag1 = bloxone_td_security_policy.test.tags.tag1
  }
}
`, name, tagValue)
}
