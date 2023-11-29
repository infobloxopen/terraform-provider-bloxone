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

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationDataSourceConfigFilters("test.com"),
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
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationDataSourceConfigTagFilters("test.com", "value1"),
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

func testAccDelegationDataSourceConfigFilters(fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test" {
  fqdn = %q
  delegation_servers = [{
	address = "12.0.0.0"
	fqdn = "test-delegation.com"
  }]
}

data "bloxone_dns_delegations" "test" {
  filters = {
	fqdn = bloxone_dns_delegation.test.fqdn
  }
}
`, fqdn)
}

func testAccDelegationDataSourceConfigTagFilters(fqdn string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone"

resource "bloxone_dns_delegation" "test" {
  fqdn = %q
  parent = 
  delegation_servers = [{
	address = "12.0.0.0"
	fqdn = "test-delegation.com"
  }]
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_delegations" "test" {
  tag_filters = {
	tag1 = bloxone_dns_delegation.test.tags.tag1
  }
}
`, fqdn, tagValue)
}
