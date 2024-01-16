package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - inheritance_sources
// - Kerberos_keys

func TestAccServerResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_server.test"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "dnssec_root_keys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ecs_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_on_v4", "no"),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_server.test"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServerBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					testAccCheckServerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccServerResource_AddEdnsOptionInOutgoingQuery(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_add_edns_option_in_outgoing_query"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAddEdnsOptionInOutgoingQuery(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "add_edns_option_in_outgoing_query", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAddEdnsOptionInOutgoingQuery(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "add_edns_option_in_outgoing_query", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_AutoSortViews(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_auto_sort_views"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAutoSortViews(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_sort_views", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAutoSortViews(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_sort_views", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_comment"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerComment(name, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccServerComment(name, "test updated commentE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test updated commentE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_CustomRootNs(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_custom_root_ns"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerCustomRootNs(name, "192.168.10.10", "tf-example.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.fqdn", "tf-example.com."),
				),
			},
			// Update and Read
			{
				Config: testAccServerCustomRootNsUpdate(name, "192.168.11.11", "tf-infoblox.com.", "192.168.11.12", "tf-infoblox-acc.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.fqdn", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.1.address", "192.168.11.12"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.1.fqdn", "tf-infoblox-acc.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_CustomRootNsEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_custom_root_ns_enabled"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerCustomRootNsEnabled(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccServerCustomRootNsEnabled(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Custom Root NS list"),
			},
			// Update and Read
			{
				Config: testAccServerCustomRootNsEnabled(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DnssecEnableValidation(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_dnssec_enable_validation"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDnssecEnableValidation(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enable_validation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDnssecEnableValidation(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enable_validation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DnssecEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_dnssec_enabled"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDnssecEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDnssecEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DnssecTrustAnchors(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_dnssec_trust_anchors"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDnssecTrustAnchors(name, "8", "tf-infoblox.com.", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.algorithm", "8"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.zone", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.sep", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDnssecTrustAnchors(name, "7", "tf-infoblox.com.", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.algorithm", "7"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.zone", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.sep", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DnssecValidateExpiry(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_dnssec_validate_expiry"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDnssecValidateExpiry(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validate_expiry", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDnssecValidateExpiry(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validate_expiry", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_EcsEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_ecs_enabled"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerEcsEnabled(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccServerEcsEnabled(name, "true", false),
				ExpectError: regexp.MustCompile("ECS zones list should not be empty"),
			},
			// Update and Read
			{
				Config: testAccServerEcsEnabled(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_EcsForwarding(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_ecs_forwarding"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerEcsForwarding(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_forwarding", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerEcsForwarding(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_forwarding", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_EcsPrefixV4(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_ecs_prefix_v4"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerEcsPrefixV4(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v4", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccServerEcsPrefixV4(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v4", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_EcsPrefixV6(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_ecs_prefix_v6"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerEcsPrefixV6(name, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v6", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccServerEcsPrefixV6(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v6", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_EcsZones(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_ecs_zones"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerEcsZones(name, "allow", "tf-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.fqdn", "tf-infoblox.com."),
				),
			},
			// Update and Read
			{
				Config: testAccServerEcsZones(name, "deny", "tf-test-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.fqdn", "tf-test-infoblox.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_FilterAaaaAcl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_filter_aaaa_acl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAclIP("filter_aaaa_acl", name, "allow", "192.168.10.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.address", "192.168.10.10"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAclAny("filter_aaaa_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_FilterAaaaOnV4(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_filter_aaaa_on_v4"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerFilterAaaaOnV4(name, "no"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_on_v4", "no"),
				),
			},
			// Update and Read
			{
				Config: testAccServerFilterAaaaOnV4(name, "break_dnssec"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_on_v4", "break_dnssec"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Forwarders(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_forwarders"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerForwarders(name, "192.168.10.10", "tf-example.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.fqdn", "tf-example.com."),
				),
			},
			// Update and Read
			{
				Config: testAccServerForwarders(name, "192.168.11.11", "tf-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.fqdn", "tf-infoblox.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_ForwardersOnly(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_forwarders_only"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerForwardersOnly(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccServerForwardersOnly(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Forwarders list"),
			},
			// Update and Read
			{
				Config: testAccServerForwardersOnly(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_GssTsigEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_gss_tsig_enabled"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerGssTsigEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerGssTsigEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_inheritance_sources"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerInheritanceSources(name, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "inherit"),
				),
				ExpectNonEmptyPlan: true,
			},
			// Update and Read
			{
				Config: testAccServerInheritanceSources(name, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "override"),
				),
				ExpectNonEmptyPlan: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_LameTtl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_lame_ttl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerLameTtl(name, 3000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lame_ttl", "3000"),
				),
			},
			// Update and Read
			{
				Config: testAccServerLameTtl(name, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lame_ttl", "3600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_LogQueryResponse(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_log_query_response"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerLogQueryResponse(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_query_response", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerLogQueryResponse(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "log_query_response", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_MatchRecursiveOnly(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_match_recursive_only"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerMatchRecursiveOnly(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_recursive_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerMatchRecursiveOnly(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_recursive_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_MaxCacheTtl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_max_cache_ttl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerMaxCacheTtl(name, 600000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "600000"),
				),
			},
			// Update and Read
			{
				Config: testAccServerMaxCacheTtl(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_MaxNegativeTtl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_max_negative_ttl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerMaxNegativeTtl(name, 10000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_negative_ttl", "10000"),
				),
			},
			// Update and Read
			{
				Config: testAccServerMaxNegativeTtl(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_negative_ttl", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_MinimalResponses(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_minimal_responses"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerMinimalResponses(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "minimal_responses", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerMinimalResponses(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "minimal_responses", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Name(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_name"
	var v dns_config.ConfigServer
	var name1 = acctest.RandomNameWithPrefix("dns-server")
	var name2 = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccServerName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Notify(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_notify"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerNotify(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerNotify(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_QueryAcl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_query_acl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAclIP("query_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAclAny("query_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_QueryPort(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_query_port"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerQueryPort(name, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_port", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccServerQueryPort(name, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_port", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_RecursionAcl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_recursion_acl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAclIP("recursion_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAclAny("recursion_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_RecursionEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_recursion_enabled"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerRecursionEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerRecursionEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_RecursiveClients(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_recursive_clients"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerRecursiveClients(name, "100"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursive_clients", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccServerRecursiveClients(name, "200"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursive_clients", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_ResolverQueryTimeout(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_resolver_query_timeout"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerResolverQueryTimeout(name, "15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolver_query_timeout", "15"),
				),
			},
			// Update and Read
			{
				Config: testAccServerResolverQueryTimeout(name, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "resolver_query_timeout", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_SecondaryAxfrQueryLimit(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_secondary_axfr_query_limit"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerSecondaryAxfrQueryLimit(name, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_axfr_query_limit", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccServerSecondaryAxfrQueryLimit(name, "3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_axfr_query_limit", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_SecondarySoaQueryLimit(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_secondary_soa_query_limit"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerSecondarySoaQueryLimit(name, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_soa_query_limit", "2"),
				),
			},
			// Update and Read
			{
				Config: testAccServerSecondarySoaQueryLimit(name, "3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_soa_query_limit", "3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_SortList(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_sort_list"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerSortList(name, "ip", "192.168.12.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.source", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.prioritized_networks.0", "192.168.12.12"),
				),
			},
			// Update and Read
			{
				Config: testAccServerSortList(name, "any", "192.168.13.13"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.element", "any"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.prioritized_networks.0", "192.168.13.13"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_SynthesizeAddressRecordsFromHttps(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_synthesize_address_records_from_https"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerSynthesizeAddressRecordsFromHttps(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synthesize_address_records_from_https", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerSynthesizeAddressRecordsFromHttps(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synthesize_address_records_from_https", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_tags"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccServerTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_TransferAcl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_transfer_acl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAclIP("transfer_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAclAny("transfer_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_UpdateAcl(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_update_acl"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerAclIP("update_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccServerAclAny("update_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_UseForwardersForSubzones(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_use_forwarders_for_subzones"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerUseForwardersForSubzones(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerUseForwardersForSubzones(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_UseRootForwardersForLocalResolutionWithB1td(t *testing.T) {
	var resourceName = "bloxone_dns_server.test_use_root_forwarders_for_local_resolution_with_b1td"
	var v dns_config.ConfigServer
	var name = acctest.RandomNameWithPrefix("dns-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerUseRootForwardersForLocalResolutionWithB1td(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccServerUseRootForwardersForLocalResolutionWithB1td(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Forwarders list"),
			},
			// Update and Read
			{
				Config: testAccServerUseRootForwardersForLocalResolutionWithB1td(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckServerExists(ctx context.Context, resourceName string, v *dns_config.ConfigServer) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ServerAPI.
			ServerRead(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckServerDestroy(ctx context.Context, v *dns_config.ConfigServer) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ServerAPI.
			ServerRead(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckServerDisappears(ctx context.Context, v *dns_config.ConfigServer) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ServerAPI.
			ServerDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccServerBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test" {
    name = %q
}
`, name)
}

func testAccServerAddEdnsOptionInOutgoingQuery(name string, addEdnsOptionInOutgoingQuery string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_add_edns_option_in_outgoing_query" {
    name = %q
    add_edns_option_in_outgoing_query = %q
}
`, name, addEdnsOptionInOutgoingQuery)
}

func testAccServerAutoSortViews(name string, autoSortViews string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_auto_sort_views" {
    name = %q
    auto_sort_views = %q
}
`, name, autoSortViews)
}

func testAccServerComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccServerCustomRootNs(name string, address string, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_custom_root_ns" {
    name = %q
    custom_root_ns = [
		{
			address = %q
			fqdn = %q
		}
]
}
`, name, address, fqdn)
}

func testAccServerCustomRootNsUpdate(name string, address string, fqdn string, address2 string, fqdn2 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_custom_root_ns" {
    name = %q
    custom_root_ns = [
		{
			address = %q
			fqdn = %q
		},
		{
			address = %q
			fqdn = %q
		}
]
}
`, name, address, fqdn, address2, fqdn2)
}

func testAccServerCustomRootNsEnabled(name string, customRootNsEnabled string, addCustomRootNSBlock bool) string {
	customRootNS := ""
	if addCustomRootNSBlock {
		customRootNS = `custom_root_ns = [
		{
			address = "192.168.10.10"
			fqdn = "tf-infoblox.com."
		}
]`
	}
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_custom_root_ns_enabled" {
    name = %q
    custom_root_ns_enabled = %q
	%s
	
}
`, name, customRootNsEnabled, customRootNS)
}

func testAccServerDnssecEnableValidation(name string, dnssecEnableValidation string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_dnssec_enable_validation" {
    name = %q
    dnssec_enable_validation = %q
}
`, name, dnssecEnableValidation)
}

func testAccServerDnssecEnabled(name string, dnssecEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_dnssec_enabled" {
    name = %q
    dnssec_enabled = %q
}
`, name, dnssecEnabled)
}

func testAccServerDnssecTrustAnchors(name string, algorithm string, zone string, sep string) string {
	public_key := "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_dnssec_trust_anchors" {
    name = %q
    dnssec_trust_anchors = [
		{
			algorithm = %q
			public_key = %q
			zone = %q
			sep = %q
		}
]
}
`, name, algorithm, public_key, zone, sep)
}

func testAccServerDnssecValidateExpiry(name string, dnssecValidateExpiry string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_dnssec_validate_expiry" {
    name = %q
    dnssec_validate_expiry = %q
}
`, name, dnssecValidateExpiry)
}

func testAccServerEcsEnabled(name string, ecsEnabled string, addEcsZonesBlock bool) string {
	ecsBlock := ""
	if addEcsZonesBlock {
		ecsBlock = `ecs_zones = [
		{
			access = "allow"
			fqdn = "tf-infoblox.com."
		}
]`
	}
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_ecs_enabled" {
    name = %q
    ecs_enabled = %q
	%s
}
`, name, ecsEnabled, ecsBlock)
}

func testAccServerEcsForwarding(name string, ecsForwarding string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_ecs_forwarding" {
    name = %q
    ecs_forwarding = %q
}
`, name, ecsForwarding)
}

func testAccServerEcsPrefixV4(name string, ecsPrefixV4 int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_ecs_prefix_v4" {
    name = %q
    ecs_prefix_v4 = %d
}
`, name, ecsPrefixV4)
}

func testAccServerEcsPrefixV6(name string, ecsPrefixV6 int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_ecs_prefix_v6" {
    name = %q
    ecs_prefix_v6 = %d
}
`, name, ecsPrefixV6)
}

func testAccServerEcsZones(name string, access, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_ecs_zones" {
    name = %q
	ecs_zones = [
		{
			access = %q
			fqdn = %q
		}
]
}
`, name, access, fqdn)
}

func testAccServerFilterAaaaOnV4(name string, filterAaaaOnV4 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_filter_aaaa_on_v4" {
    name = %q
    filter_aaaa_on_v4 = %q
}
`, name, filterAaaaOnV4)
}

func testAccServerForwarders(name string, address, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_forwarders" {
    name = %q
	forwarders = [
		{
			address = %q
			fqdn = %q
		}
]
}
`, name, address, fqdn)
}

func testAccServerForwardersOnly(name string, forwarderOnly string, addForwardersBlock bool) string {
	forwardersBlock := ""
	if addForwardersBlock {
		forwardersBlock = `forwarders = [
		{
			address = "192.168.11.11"
			fqdn = "tf-infoblox.com."
		}
]`
	}
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_forwarders_only" {
    name = %q
    forwarders_only = %q
	%s
	
}
`, name, forwarderOnly, forwardersBlock)
}

func testAccServerGssTsigEnabled(name string, gssTsigEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_gss_tsig_enabled" {
    name = %q
    gss_tsig_enabled = %q
}
`, name, gssTsigEnabled)
}

func testAccServerInheritanceSources(name, action string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_inheritance_sources" {
    name = %[1]q
	inheritance_sources = { 
		gss_tsig_enabled = {
			action = %[2]q
		}
		notify = {
			action = %[2]q
		}
		transfer_acl = {
			action = %[2]q
		}
		useforwardersforsubzones = {
			action = %[2]q
		}
	}
	gss_tsig_enabled = true
	notify = true
	transfer_acl = [
		{
			access = "allow"
			element = "ip"
			address = "192.168.11.11"
		}
	]
	use_forwarders_for_subzones = true
		

}
`, name, action)
}

func testAccServerLameTtl(name string, lameTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_lame_ttl" {
    name = %q
    lame_ttl = %d
}
`, name, lameTtl)
}

func testAccServerLogQueryResponse(name string, logQueryResponse string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_log_query_response" {
    name = %q
    log_query_response = %q
}
`, name, logQueryResponse)
}

func testAccServerMatchRecursiveOnly(name string, matchRecursiveOnly string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_match_recursive_only" {
    name = %q
    match_recursive_only = %q
}
`, name, matchRecursiveOnly)
}

func testAccServerMaxCacheTtl(name string, maxCacheTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_max_cache_ttl" {
    name = %q
    max_cache_ttl = %d
}
`, name, maxCacheTtl)
}

func testAccServerMaxNegativeTtl(name string, maxNegativeTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_max_negative_ttl" {
    name = %q
    max_negative_ttl = %d
}
`, name, maxNegativeTtl)
}

func testAccServerMinimalResponses(name string, minimalResponses string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_minimal_responses" {
    name = %q
    minimal_responses = %q
}
`, name, minimalResponses)
}

func testAccServerName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_name" {
    name = %q
}
`, name)
}

func testAccServerNotify(name string, notify string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_notify" {
    name = %q
    notify = %q
}
`, name, notify)
}

func testAccServerQueryPort(name string, queryPort string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_query_port" {
    name = %q
    query_port = %q
}
`, name, queryPort)
}

func testAccServerAclIP(aclFieldName, name, access, address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_%[1]s" {
    name = %[2]q
    %[1]s = [
		{
			access = %[3]q
			element = "ip"
			address = %[4]q
		}
]
}
`, aclFieldName, name, access, address)
}

func testAccServerAclAny(aclFieldName, name, access string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_%[1]s" {
    name = %[2]q
    %[1]s = [
		{
			access = %[3]q
			element = "any"
		}
]
}
`, aclFieldName, name, access)
}

func testAccServerRecursionEnabled(name string, recursionEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_recursion_enabled" {
    name = %q
    recursion_enabled = %q
}
`, name, recursionEnabled)
}

func testAccServerRecursiveClients(name string, recursiveClients string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_recursive_clients" {
    name = %q
    recursive_clients = %q
}
`, name, recursiveClients)
}

func testAccServerResolverQueryTimeout(name string, resolverQueryTimeout string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_resolver_query_timeout" {
    name = %q
    resolver_query_timeout = %q
}
`, name, resolverQueryTimeout)
}

func testAccServerSecondaryAxfrQueryLimit(name string, secondaryAxfrQueryLimit string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_secondary_axfr_query_limit" {
    name = %q
    secondary_axfr_query_limit = %q
}
`, name, secondaryAxfrQueryLimit)
}

func testAccServerSecondarySoaQueryLimit(name string, secondarySoaQueryLimit string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_secondary_soa_query_limit" {
    name = %q
    secondary_soa_query_limit = %q
}
`, name, secondarySoaQueryLimit)
}

func testAccServerSortList(name string, element string, addressPrioritizedNetworks string) string {
	sourceAdd := ""
	if element == "ip" {
		sourceAdd = "source = \"192.168.11.11\""
	}
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_sort_list" {
    name = %q
    sort_list = [
		{
			
			element = %q
			%s
			prioritized_networks = [ "%s" ]
		}
]
}
`, name, element, sourceAdd, addressPrioritizedNetworks)
}

func testAccServerSynthesizeAddressRecordsFromHttps(name string, synthesizeAddressRecordsFromHttps string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_synthesize_address_records_from_https" {
    name = %q
    synthesize_address_records_from_https = %q
}
`, name, synthesizeAddressRecordsFromHttps)
}

func testAccServerTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}

func testAccServerUseForwardersForSubzones(name string, useForwardersForSubzones string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_use_forwarders_for_subzones" {
    name = %q
    use_forwarders_for_subzones = %q
}
`, name, useForwardersForSubzones)
}

func testAccServerUseRootForwardersForLocalResolutionWithB1td(name string, useRootForwardersForLocalResolutionWithB1td string, addForwardersBlock bool) string {
	forwardersBlock := ""
	if addForwardersBlock {
		forwardersBlock = `forwarders = [
		{
			address = "192.168.11.11"
			fqdn = "tf-infoblox.com."
		}
]`
	}
	return fmt.Sprintf(`
resource "bloxone_dns_server" "test_use_root_forwarders_for_local_resolution_with_b1td" {
    name = %q
    use_root_forwarders_for_local_resolution_with_b1td = %q
	%s
	
}
`, name, useRootForwardersForLocalResolutionWithB1td, forwardersBlock)
}
