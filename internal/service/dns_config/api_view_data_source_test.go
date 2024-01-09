package dns_config_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccViewDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_views.test"
	resourceName := "bloxone_dns_view.test"
	var v dns_config.ConfigView
	var name = acctest.RandomNameWithPrefix("view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckViewExists(context.Background(), resourceName, &v),
					}, testAccCheckViewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccViewDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_views.test"
	resourceName := "bloxone_dns_view.test"
	var v dns_config.ConfigView
	var name = acctest.RandomNameWithPrefix("view")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewDataSourceConfigTagFilters(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckViewExists(context.Background(), resourceName, &v),
					}, testAccCheckViewResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckViewResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "add_edns_option_in_outgoing_query", dataSourceName, "results.0.add_edns_option_in_outgoing_query"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "custom_root_ns", dataSourceName, "results.0.custom_root_ns"),
		resource.TestCheckResourceAttrPair(resourceName, "custom_root_ns_enabled", dataSourceName, "results.0.custom_root_ns_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_enable_validation", dataSourceName, "results.0.dnssec_enable_validation"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_enabled", dataSourceName, "results.0.dnssec_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_root_keys", dataSourceName, "results.0.dnssec_root_keys"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_trust_anchors", dataSourceName, "results.0.dnssec_trust_anchors"),
		resource.TestCheckResourceAttrPair(resourceName, "dnssec_validate_expiry", dataSourceName, "results.0.dnssec_validate_expiry"),
		resource.TestCheckResourceAttrPair(resourceName, "dtc_config", dataSourceName, "results.0.dtc_config"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_enabled", dataSourceName, "results.0.ecs_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_forwarding", dataSourceName, "results.0.ecs_forwarding"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_prefix_v4", dataSourceName, "results.0.ecs_prefix_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_prefix_v6", dataSourceName, "results.0.ecs_prefix_v6"),
		resource.TestCheckResourceAttrPair(resourceName, "ecs_zones", dataSourceName, "results.0.ecs_zones"),
		resource.TestCheckResourceAttrPair(resourceName, "edns_udp_size", dataSourceName, "results.0.edns_udp_size"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_acl", dataSourceName, "results.0.filter_aaaa_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_on_v4", dataSourceName, "results.0.filter_aaaa_on_v4"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarders", dataSourceName, "results.0.forwarders"),
		resource.TestCheckResourceAttrPair(resourceName, "forwarders_only", dataSourceName, "results.0.forwarders_only"),
		resource.TestCheckResourceAttrPair(resourceName, "gss_tsig_enabled", dataSourceName, "results.0.gss_tsig_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "ip_spaces", dataSourceName, "results.0.ip_spaces"),
		resource.TestCheckResourceAttrPair(resourceName, "lame_ttl", dataSourceName, "results.0.lame_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "match_clients_acl", dataSourceName, "results.0.match_clients_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "match_destinations_acl", dataSourceName, "results.0.match_destinations_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "match_recursive_only", dataSourceName, "results.0.match_recursive_only"),
		resource.TestCheckResourceAttrPair(resourceName, "max_cache_ttl", dataSourceName, "results.0.max_cache_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "max_negative_ttl", dataSourceName, "results.0.max_negative_ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "max_udp_size", dataSourceName, "results.0.max_udp_size"),
		resource.TestCheckResourceAttrPair(resourceName, "minimal_responses", dataSourceName, "results.0.minimal_responses"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "notify", dataSourceName, "results.0.notify"),
		resource.TestCheckResourceAttrPair(resourceName, "query_acl", dataSourceName, "results.0.query_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "recursion_acl", dataSourceName, "results.0.recursion_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "recursion_enabled", dataSourceName, "results.0.recursion_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "sort_list", dataSourceName, "results.0.sort_list"),
		resource.TestCheckResourceAttrPair(resourceName, "synthesize_address_records_from_https", dataSourceName, "results.0.synthesize_address_records_from_https"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "transfer_acl", dataSourceName, "results.0.transfer_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "update_acl", dataSourceName, "results.0.update_acl"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "use_forwarders_for_subzones", dataSourceName, "results.0.use_forwarders_for_subzones"),
		resource.TestCheckResourceAttrPair(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", dataSourceName, "results.0.use_root_forwarders_for_local_resolution_with_b1td"),
		resource.TestCheckResourceAttrPair(resourceName, "zone_authority", dataSourceName, "results.0.zone_authority"),
	}
}

func testAccViewDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
  name = %q
}

data "bloxone_dns_views" "test" {
  filters = {
	name = bloxone_dns_view.test.name
  }
}
`, name)
}

func testAccViewDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_views" "test" {
  tag_filters = {
	tag1 = bloxone_dns_view.test.tags.tag1
  }
}
`, name, tagValue)
}
