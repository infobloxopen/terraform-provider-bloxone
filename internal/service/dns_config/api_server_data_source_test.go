package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccServerDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_servers.test"
	resourceName := "bloxone_dns_server.test"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-servers")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServerExists(context.Background(), resourceName, &v),
					}, testAccCheckServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccServerDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_servers.test"
	resourceName := "bloxone_dns_server.test"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-servers")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServerDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckServerExists(context.Background(), resourceName, &v),
					}, testAccCheckServerResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckServerResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "add_edns_option_in_outgoing_query", dataSourceName, "results.0.add_edns_option_in_outgoing_query"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_sort_views", dataSourceName, "results.0.auto_sort_views"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "custom_root_ns", dataSourceName, "results.0.custom_root_ns"),
		resource.TestCheckResourceAttrPair(resourceName, "custom_root_ns_enabled", dataSourceName, "results.0.custom_root_ns_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_enable_validation", dataSourceName, "results.0.dnssec_enable_validation"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_enabled", dataSourceName, "results.0.dnssec_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_root_keys", dataSourceName, "results.0.dnssec_root_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_trust_anchors", dataSourceName, "results.0.dnssec_trust_anchors"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_validate_expiry", dataSourceName, "results.0.dnssec_validate_expiry"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_enabled", dataSourceName, "results.0.ecs_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_forwarding", dataSourceName, "results.0.ecs_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_prefix_v4", dataSourceName, "results.0.ecs_prefix_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_prefix_v6", dataSourceName, "results.0.ecs_prefix_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_zones", dataSourceName, "results.0.ecs_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_acl", dataSourceName, "results.0.filter_aaaa_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_on_v4", dataSourceName, "results.0.filter_aaaa_on_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarders", dataSourceName, "results.0.forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarders_only", dataSourceName, "results.0.forwarders_only"),
		resource.TestCheckResourceAttrPair(resourceName, "gss_tsig_enabled", dataSourceName, "results.0.gss_tsig_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "kerberos_keys", dataSourceName, "results.0.kerberos_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "lame_ttl", dataSourceName, "results.0.lame_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "log_query_response", dataSourceName, "results.0.log_query_response"),
		resource.TestCheckResourceAttrPair(resourceName, "match_recursive_only", dataSourceName, "results.0.match_recursive_only"),
		resource.TestCheckResourceAttrPair(resourceName, "max_cache_ttl", dataSourceName, "results.0.max_cache_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "max_negative_ttl", dataSourceName, "results.0.max_negative_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "minimal_responses", dataSourceName, "results.0.minimal_responses"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "notify", dataSourceName, "results.0.notify"),
		resource.TestCheckResourceAttrPair(resourceName, "query_acl", dataSourceName, "results.0.query_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "query_port", dataSourceName, "results.0.query_port"),
		resource.TestCheckResourceAttrPair(resourceName, "recursion_acl", dataSourceName, "results.0.recursion_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "recursion_enabled", dataSourceName, "results.0.recursion_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "recursive_clients", dataSourceName, "results.0.recursive_clients"),
		resource.TestCheckResourceAttrPair(resourceName, "resolver_query_timeout", dataSourceName, "results.0.resolver_query_timeout"),
		resource.TestCheckResourceAttrPair(resourceName, "secondary_axfr_query_limit", dataSourceName, "results.0.secondary_axfr_query_limit"),
		resource.TestCheckResourceAttrPair(resourceName, "secondary_soa_query_limit", dataSourceName, "results.0.secondary_soa_query_limit"),
		resource.TestCheckResourceAttrPair(resourceName, "sort_list", dataSourceName, "results.0.sort_list"),
		resource.TestCheckResourceAttrPair(resourceName, "synthesize_address_records_from_https", dataSourceName, "results.0.synthesize_address_records_from_https"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "transfer_acl", dataSourceName, "results.0.transfer_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "update_acl", dataSourceName, "results.0.update_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "use_forwarders_for_subzones", dataSourceName, "results.0.use_forwarders_for_subzones"),
		resource.TestCheckResourceAttrPair(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", dataSourceName, "results.0.use_root_forwarders_for_local_resolution_with_b1td"),
		resource.TestCheckResourceAttrPair(resourceName, "views", dataSourceName, "results.0.views"),
	}
}

func testAccServerDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test" {
  name = %q
}

data "bloxone_dns_servers" "test" {
  filters = {
	name = bloxone_dns_server.test.name
  }
}
`, name)
}

func testAccServerDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_servers" "test" {
  tag_filters = {
	tag1 = bloxone_dns_server.test.tags.tag1
  }
}
`, name, tagValue)
}
