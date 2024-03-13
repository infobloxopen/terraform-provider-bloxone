package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDelegationDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_delegations.test"
	resourceName := "bloxone_dns_delegation.test"
	var v dns_config.ConfigDelegation
	viewName := acctest.RandomNameWithPrefix("view")
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	delegationFqdn := "test." + zoneFqdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationDataSourceConfigFilters(viewName, zoneFqdn, delegationFqdn),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDelegationExists(context.Background(), resourceName, &v),
					}, testAccCheckDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDelegationDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_delegations.test"
	resourceName := "bloxone_dns_delegation.test"
	var v dns_config.ConfigDelegation
	viewName := acctest.RandomNameWithPrefix("view")
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	delegationFqdn := "test." + zoneFqdn

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationDataSourceConfigTagFilters(viewName, zoneFqdn, delegationFqdn, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDelegationExists(context.Background(), resourceName, &v),
					}, testAccCheckDelegationResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDelegationResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "delegation_servers", dataSourceName, "results.0.delegation_servers"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn", dataSourceName, "results.0.fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "parent", dataSourceName, "results.0.parent"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol_fqdn", dataSourceName, "results.0.protocol_fqdn"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "results.0.view"),
	}
}

func testAccDelegationDataSourceConfigFilters(viewName, zoneFqdn, delegationFqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
    name = %q
}

resource "bloxone_dns_auth_zone" test{
  fqdn         = %q
  primary_type = "cloud"
  view = bloxone_dns_view.test.id
}

resource "bloxone_dns_delegation" "test" {
  // test.tf-acc-test.com.
  fqdn = %q
  delegation_servers = [{
	address = "12.0.0.0"
	fqdn = "ns1.com."
  }]
  view = bloxone_dns_view.test.id
  depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}

data "bloxone_dns_delegations" "test" {
  filters = {
	fqdn = bloxone_dns_delegation.test.fqdn
  }
}
`, viewName, zoneFqdn, delegationFqdn)
}

func testAccDelegationDataSourceConfigTagFilters(viewName, zoneFqdn, delegationFqdn string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
    name = %q
}

resource "bloxone_dns_auth_zone" test {
  fqdn         = %q
  primary_type = "cloud"
  view = bloxone_dns_view.test.id
  depends_on = [bloxone_dns_view.test]
}

resource "bloxone_dns_delegation" "test" {
  fqdn = %q
  delegation_servers = [{
	address = "12.0.0.0"
	fqdn = "test-delegation.com."
  }]
  tags = {
	tag1 = %q
  }
  view = bloxone_dns_view.test.id
  depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}

data "bloxone_dns_delegations" "test" {
  tag_filters = {
	tag1 = bloxone_dns_delegation.test.tags.tag1
  }
}
`, viewName, zoneFqdn, delegationFqdn, tagValue)
}
