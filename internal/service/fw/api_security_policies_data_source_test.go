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
	var v fw.SecurityPolicy
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
	var v fw.SecurityPolicy
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
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.access_codes", resourceName, "access_codes"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.created_at", resourceName, "created_at"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.default_action", resourceName, "default_action"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.default_redirect_name", resourceName, "default_redirect_name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.dfp_services", resourceName, "dfp_services"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.dfps", resourceName, "dfps"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.ecs", resourceName, "ecs"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.id", resourceName, "id"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.is_default", resourceName, "is_default"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.name", resourceName, "name"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.description", resourceName, "description"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.net_address_dfps", resourceName, "net_address_dfps"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.network_lists", resourceName, "network_lists"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.onprem_resolve", resourceName, "onprem_resolve"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.precedence", resourceName, "precedence"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.roaming_device_groups", resourceName, "roaming_device_groups"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.tags", resourceName, "tags"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.rules", resourceName, "rules"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.safe_search", resourceName, "safe_search"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.updated_time", resourceName, "updated_time"),
		resource.TestCheckResourceAttrPair(dataSourceName, "results.0.user_groups", resourceName, "user_groups"),
	}
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
  name = %q
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
