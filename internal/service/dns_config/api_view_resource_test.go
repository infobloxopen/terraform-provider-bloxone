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

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - zone_authority : Mname and rname provide inconsistent result after apply

func TestAccViewResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_view.test"
	var v dnsconfig.View
	var name = acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "dnssec_root_keys.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
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

func TestAccViewResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_view.test"
	var v dnsconfig.View
	var name = acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckViewDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccViewBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					testAccCheckViewDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccViewResource_Name(t *testing.T) {
	var resourceName = "bloxone_dns_view.test"
	var name1 = acctest.RandomNameWithPrefix("view")
	var name2 = acctest.RandomNameWithPrefix("view")
	var v1 dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewBasicConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccViewBasicConfig(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_AddEdnsOptionInOutgoingQuery(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_add_edns_option_in_outgoing_query"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewAddEdnsOptionInOutgoingQuery(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "add_edns_option_in_outgoing_query", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewAddEdnsOptionInOutgoingQuery(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "add_edns_option_in_outgoing_query", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_comment"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewComment(name, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccViewComment(name, "another test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "another test comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_CustomRootNs(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_custom_root_ns"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewCustomRootNs(name, "192.168.10.10", "tf-example.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns.0.fqdn", "tf-example.com."),
				),
			},
			// Update and Read
			{
				Config: testAccViewCustomRootNsUpdate(name, "192.168.11.11", "tf-infoblox.com.", "192.168.11.12", "tf-infoblox-acc.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
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

func TestAccViewResource_CustomRootNsEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_custom_root_ns_enabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewCustomRootNsEnabled(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccViewCustomRootNsEnabled(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Custom Root NS list"),
			},
			// Update and Read
			{
				Config: testAccViewCustomRootNsEnabled(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "custom_root_ns_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_disabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDisabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDisabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecEnableValidation(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_dnssec_enable_validation"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecEnableValidation(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enable_validation", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecEnableValidation(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enable_validation", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_dnssec_enabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecTrustAnchors(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_dnssec_trust_anchors"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecTrustAnchors(name, "8", "tf-infoblox.com.", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.algorithm", "8"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.zone", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.sep", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecTrustAnchors(name, "7", "tf-infoblox.com.", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.algorithm", "7"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.zone", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "dnssec_trust_anchors.0.sep", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DnssecValidateExpiry(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_dnssec_validate_expiry"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDnssecValidateExpiry(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validate_expiry", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDnssecValidateExpiry(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_validate_expiry", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_DtcConfig(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_dtc_config"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewDtcConfig(name, 700),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_config.default_ttl", "700"),
				),
			},
			// Update and Read
			{
				Config: testAccViewDtcConfig(name, 500),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dtc_config.default_ttl", "500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EcsEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ecs_enabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEcsEnabled(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccViewEcsEnabled(name, "true", false),
				ExpectError: regexp.MustCompile("should not be empty if ECS is enabled"),
			},
			// Update and Read
			{
				Config: testAccViewEcsEnabled(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EcsForwarding(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ecs_forwarding"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEcsForwarding(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_forwarding", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEcsForwarding(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_forwarding", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EcsPrefixV4(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ecs_prefix_v4"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEcsPrefixV4(name, 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v4", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEcsPrefixV4(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v4", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EcsPrefixV6(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ecs_prefix_v6"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEcsPrefixV6(name, 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v6", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEcsPrefixV6(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_prefix_v6", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EcsZones(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ecs_zones"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEcsZones(name, "allow", "tf-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.fqdn", "tf-infoblox.com."),
				),
			},
			// Update and Read
			{
				Config: testAccViewEcsZones(name, "deny", "tf-test-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "ecs_zones.0.fqdn", "tf-test-infoblox.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_EdnsUdpSize(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_edns_udp_size"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewEdnsUdpSize(name, 1200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "1200"),
				),
			},
			// Update and Read
			{
				Config: testAccViewEdnsUdpSize(name, 1000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "edns_udp_size", "1000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaaAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_filter_aaaa_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "filter_aaaa_acl", name, "allow", "192.168.10.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.address", "192.168.10.10"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "filter_aaaa_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "filter_aaaa_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "filter_aaaa_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "filter_aaaa_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_FilterAaaaOnV4(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_filter_aaaa_on_v4"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewFilterAaaaOnV4(name, "no"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_on_v4", "no"),
				),
			},
			// Update and Read
			{
				Config: testAccViewFilterAaaaOnV4(name, "break_dnssec"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_aaaa_on_v4", "break_dnssec"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Forwarders(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_forwarders"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwarders(name, "192.168.10.10", "tf-example.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.fqdn", "tf-example.com."),
				),
			},
			// Update and Read
			{
				Config: testAccViewForwarders(name, "192.168.11.11", "tf-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "forwarders.0.fqdn", "tf-infoblox.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ForwardersOnly(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_forwarders_only"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewForwardersOnly(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccViewForwardersOnly(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Forwarders list"),
			},
			// Update and Read
			{
				Config: testAccViewForwardersOnly(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_GssTsigEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_gss_tsig_enabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewGssTsigEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewGssTsigEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_inheritance_sources"
	var v dnsconfig.View
	var name = acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewInheritanceSources(name, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.add_edns_option_in_outgoing_query.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.custom_root_ns_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dnssec_validation_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ecs_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.filter_aaaa_on_v4.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.forwarders_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.lame_ttl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.edns_udp_size.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.match_recursive_only.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.max_cache_ttl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.max_negative_ttl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.minimal_responses.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.notify.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.recursion_enabled.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.sort_list.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.synthesize_address_records_from_https.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.transfer_acl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.use_forwarders_for_subzones.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.default_ttl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.expire.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.mname_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.negative_ttl.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.refresh.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.retry.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.rname.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccViewInheritanceSources(name, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.add_edns_option_in_outgoing_query.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.custom_root_ns_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dnssec_validation_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ecs_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.filter_aaaa_on_v4.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.forwarders_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.lame_ttl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.edns_udp_size.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.match_recursive_only.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.max_cache_ttl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.max_negative_ttl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.minimal_responses.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.notify.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.recursion_enabled.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.sort_list.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.synthesize_address_records_from_https.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.transfer_acl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.use_forwarders_for_subzones.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.default_ttl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.expire.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.mname_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.negative_ttl.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.refresh.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.retry.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.zone_authority.rname.action", "override"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_IpSpaces(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_ip_spaces"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View
	var ipSpaceName = acctest.RandomNameWithPrefix("ip_space")
	var ipSpaceName2 = acctest.RandomNameWithPrefix("ip_space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewIpSpaces(ipSpaceName, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_spaces.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "ip_spaces.0", "bloxone_ipam_ip_space.test_space", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccViewIpSpaces(ipSpaceName2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ip_spaces.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "ip_spaces.0", "bloxone_ipam_ip_space.test_space", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_LameTtl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_lame_ttl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewLameTtl(name, 3000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lame_ttl", "3000"),
				),
			},
			// Update and Read
			{
				Config: testAccViewLameTtl(name, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "lame_ttl", "3600"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchClientsAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_match_clients_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "match_clients_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "match_clients_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "match_clients_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "match_clients_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "match_clients_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "match_clients_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "match_clients_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchDestinationsAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_match_destinations_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "match_destinations_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "match_destinations_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "match_destinations_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "match_destinations_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "match_destinations_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "match_destinations_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "match_destinations_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MatchRecursiveOnly(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_match_recursive_only"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMatchRecursiveOnly(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_recursive_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMatchRecursiveOnly(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "match_recursive_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxCacheTtl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_max_cache_ttl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxCacheTtl(name, 600000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "600000"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxCacheTtl(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_cache_ttl", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxNegativeTtl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_max_negative_ttl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxNegativeTtl(name, 10000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_negative_ttl", "10000"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxNegativeTtl(name, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_negative_ttl", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MaxUdpSize(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_max_udp_size"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMaxUdpSize(name, 1232),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "1232"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMaxUdpSize(name, 512),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_udp_size", "512"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_MinimalResponses(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_minimal_responses"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewMinimalResponses(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "minimal_responses", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewMinimalResponses(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "minimal_responses", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Notify(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_notify"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewNotify(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewNotify(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_QueryAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_query_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "query_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "query_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "query_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "query_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "query_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "query_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RecursionAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_recursion_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "recursion_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "recursion_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "recursion_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "recursion_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "recursion_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "recursion_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "recursion_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_RecursionEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_recursion_enabled"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewRecursionEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewRecursionEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "recursion_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_SortList(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_sort_list"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewSortList(name, "ip", "192.168.12.12"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.source", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.prioritized_networks.0", "192.168.12.12"),
				),
			},
			// Update and Read
			{
				Config: testAccViewSortList(name, "any", "192.168.13.13"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.element", "any"),
					resource.TestCheckResourceAttr(resourceName, "sort_list.0.prioritized_networks.0", "192.168.13.13"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_SynthesizeAddressRecordsFromHttps(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_synthesize_address_records_from_https"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewSynthesizeAddressRecordsFromHttps(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synthesize_address_records_from_https", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewSynthesizeAddressRecordsFromHttps(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "synthesize_address_records_from_https", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_tags"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccViewTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_TransferAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_transfer_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "transfer_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "transfer_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "transfer_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "transfer_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "transfer_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "transfer_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UpdateAcl(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_update_acl"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("view", "update_acl", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.address", "192.168.11.11"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAny("view", "update_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAclAcl("view", "update_acl", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "update_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("view", "update_acl", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "update_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseForwardersForSubzones(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_use_forwarders_for_subzones"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseForwardersForSubzones(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccViewUseForwardersForSubzones(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_UseRootForwardersForLocalResolutionWithB1td(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_use_root_forwarders_for_local_resolution_with_b1td"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewUseRootForwardersForLocalResolutionWithB1td(name, "false", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", "false"),
				),
			},
			// Update and Read
			{
				Config:      testAccViewUseRootForwardersForLocalResolutionWithB1td(name, "true", false),
				ExpectError: regexp.MustCompile("Cannot use empty Forwarders list"),
			},
			// Update and Read
			{
				Config: testAccViewUseRootForwardersForLocalResolutionWithB1td(name, "true", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_root_forwarders_for_local_resolution_with_b1td", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccViewResource_ZoneAuthority(t *testing.T) {
	var resourceName = "bloxone_dns_view.test_zone_authority"
	var name = acctest.RandomNameWithPrefix("view")
	var v dnsconfig.View

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccViewZoneAuthority(name, 28600, 2519200, "test.b1ddi", 700,
					10500, 3500, "host", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.default_ttl", "28600"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.expire", "2519200"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.mname", "test.b1ddi"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.negative_ttl", "700"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.refresh", "10500"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.retry", "3500"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.rname", "host"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.use_default_mname", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccViewZoneAuthority(name, 30000, 2519200,
					"test-infoblox.b1ddi", 800, 11800, 3700, "host-test", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckViewExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.default_ttl", "30000"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.expire", "2519200"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.mname", "test-infoblox.b1ddi"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.negative_ttl", "800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.refresh", "11800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.retry", "3700"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.rname", "host-test"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.use_default_mname", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckViewExists(ctx context.Context, resourceName string, v *dnsconfig.View) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ViewAPI.
			Read(ctx, rs.Primary.ID).
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

func testAccCheckViewDestroy(ctx context.Context, v *dnsconfig.View) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ViewAPI.
			Read(ctx, *v.Id).
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

func testAccCheckViewDisappears(ctx context.Context, v *dnsconfig.View) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ViewAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccViewBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
    name = %q
}
`, name)
}

func testAccViewAddEdnsOptionInOutgoingQuery(name, addEdnsOptionInOutgoingQuery string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_add_edns_option_in_outgoing_query" {
    name = %q
    add_edns_option_in_outgoing_query = %q
}
`, name, addEdnsOptionInOutgoingQuery)
}

func testAccViewComment(name, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccViewCustomRootNs(name string, address string, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_custom_root_ns" {
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

func testAccViewCustomRootNsUpdate(name string, address string, fqdn string, address2 string, fqdn2 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_custom_root_ns" {
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

func testAccViewCustomRootNsEnabled(name string, customRootNsEnabled string, addCustomRootNSBlock bool) string {
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
resource "bloxone_dns_view" "test_custom_root_ns_enabled" {
    name = %q
    custom_root_ns_enabled = %q
	%s
	
}
`, name, customRootNsEnabled, customRootNS)
}

func testAccViewDisabled(name, disabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_disabled" {
    name = %q
    disabled = %q
}
`, name, disabled)
}

func testAccViewDnssecEnableValidation(name, dnssecEnableValidation string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_dnssec_enable_validation" {
    name = %q
    dnssec_enable_validation = %q
}
`, name, dnssecEnableValidation)
}

func testAccViewDnssecEnabled(name, dnssecEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_dnssec_enabled" {
    name = %q
    dnssec_enabled = %q
}
`, name, dnssecEnabled)
}

func testAccViewDnssecTrustAnchors(name string, algorithm string, zone string, sep string) string {
	public_key := "AwEAAaz/tAm8yTn4Mfeh5eyI96WSVexTBAvkMgJzkKTOiW1vkIbzxeF3+/4RgWOq7HrxRixHlFlExOLAJr5emLvN7SWXgnLh4+B5xQlNVz8Og8kvArMtNROxVQuCaSnIDdD5LKyWbRd2n9WGe2R8PzgCmr3EgVLrjyBxWezF0jLHwVN8efS3rCj/EWgvIWgb9tarpVUDK/b58Da+sqqls3eNbuv7pr+eoZG+SrDK6nWeL3c6H5Apxz7LjVc1uTIdsIXxuOLYA4/ilBmSVIzuDWfdRUfhHdY6+cn8HFRm+2hM8AnXGXws9555KrUB5qihylGa8subX2Nn6UwNR1AkUTV74bU="
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_dnssec_trust_anchors" {
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

func testAccViewDnssecValidateExpiry(name, dnssecValidateExpiry string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_dnssec_validate_expiry" {
    name = %q
    dnssec_validate_expiry = %q
}
`, name, dnssecValidateExpiry)
}

func testAccViewDtcConfig(name string, default_ttl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_dtc_config" {
    name = %q
    dtc_config = {
			default_ttl = %d
		}
}
`, name, default_ttl)
}

func testAccViewEcsEnabled(name string, ecsEnabled string, addEcsZonesBlock bool) string {
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
resource "bloxone_dns_view" "test_ecs_enabled" {
    name = %q
    ecs_enabled = %q
	%s
}
`, name, ecsEnabled, ecsBlock)
}

func testAccViewEcsForwarding(name, ecsForwarding string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_ecs_forwarding" {
    name = %q
    ecs_forwarding = %q
}
`, name, ecsForwarding)
}

func testAccViewEcsPrefixV4(name string, ecsPrefixV4 int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_ecs_prefix_v4" {
    name = %q
    ecs_prefix_v4 = %d
}
`, name, ecsPrefixV4)
}

func testAccViewEcsPrefixV6(name string, ecsPrefixV6 int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_ecs_prefix_v6" {
    name = %q
    ecs_prefix_v6 = %d
}
`, name, ecsPrefixV6)
}

func testAccViewEcsZones(name, access, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_ecs_zones" {
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

func testAccViewEdnsUdpSize(name string, ednsUdpSize int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_edns_udp_size" {
    name = %q
    edns_udp_size = %d
}
`, name, ednsUdpSize)
}

func testAccViewFilterAaaaOnV4(name, filterAaaaOnV4 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_filter_aaaa_on_v4" {
    name = %q
    filter_aaaa_on_v4 = %q
}
`, name, filterAaaaOnV4)
}

func testAccViewForwarders(name, address, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_forwarders" {
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

func testAccViewForwardersOnly(name string, forwarderOnly string, addForwardersBlock bool) string {
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
resource "bloxone_dns_view" "test_forwarders_only" {
    name = %q
    forwarders_only = %q
	%s
	
}
`, name, forwarderOnly, forwardersBlock)
}
func testAccViewGssTsigEnabled(name, gssTsigEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_gss_tsig_enabled" {
    name = %q
    gss_tsig_enabled = %q
}
`, name, gssTsigEnabled)
}

func testAccViewInheritanceSources(name, action string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_inheritance_sources" {
    name = %[1]q
	inheritance_sources = {
		add_edns_option_in_outgoing_query = {
			action = %[2]q
		}
		custom_root_ns_block = {
			action = %[2]q
		}
		dnssec_validation_block	= {
			action = %[2]q
		}
		ecs_block = {
			action = %[2]q
		}
		edns_udp_size	= {
			action = %[2]q
		}
		filter_aaaa_on_v4 = {
			action = %[2]q
		}
		forwarders_block = {
			action = %[2]q
		}
		gss_tsig_enabled = {
			action = %[2]q
		}
		kerberos_keys	= {
			action = %[2]q
		}
		lame_ttl	= {
			action = %[2]q
		}
		match_recursive_only	= {
			action = %[2]q
		}
		max_cache_ttl	= {
			action = %[2]q
		}
		max_negative_ttl	= {
			action = %[2]q
		}
		minimal_responses	= {
			action = %[2]q
		}
		notify = {
			action = %[2]q
		}
		recursion_enabled = {
			action = %[2]q
		}
		sort_list = {
			action = %[2]q
		}
		synthesize_address_records_from_https = {
			action = %[2]q
		}
		transfer_acl = {
			action = %[2]q
		}
		use_forwarders_for_subzones = {
			action = %[2]q
		}
		zone_authority = {
			default_ttl	= {
				action = %[2]q
			}
			expire	= {
				action = %[2]q
			}
			mname_block	= {
				action = %[2]q
			}
			negative_ttl	= {
				action = %[2]q
			}
			refresh	= {
				action = %[2]q
			}
			retry	= {
				action = %[2]q
			}
			rname	= {
				action = %[2]q
			}
		}
	}

}
`, name, action)
}

func testAccViewIpSpaces(ipSpaceName, viewName string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_space" {
	name = %q
}
resource "bloxone_dns_view" "test_ip_spaces" {
    name = %q
    ip_spaces = [
		bloxone_ipam_ip_space.test_space.id
]
}
`, ipSpaceName, viewName)
}

func testAccViewLameTtl(name string, lameTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_lame_ttl" {
    name = %q
    lame_ttl = %d
}
`, name, lameTtl)
}

func testAccViewMatchRecursiveOnly(name, matchRecursiveOnly string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_match_recursive_only" {
    name = %q
    match_recursive_only = %q
}
`, name, matchRecursiveOnly)
}

func testAccViewMaxCacheTtl(name string, maxCacheTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_max_cache_ttl" {
    name = %q
    max_cache_ttl = %d
}
`, name, maxCacheTtl)
}

func testAccViewMaxNegativeTtl(name string, maxNegativeTtl int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_max_negative_ttl" {
    name = %q
    max_negative_ttl = %d
}
`, name, maxNegativeTtl)
}

func testAccViewMaxUdpSize(name string, maxUdpSize int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_max_udp_size" {
    name = %q
    max_udp_size = %d
}
`, name, maxUdpSize)
}

func testAccViewMinimalResponses(name, minimalResponses string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_minimal_responses" {
    name = %q
    minimal_responses = %q
}
`, name, minimalResponses)
}

func testAccViewNotify(name, notify string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_notify" {
    name = %q
    notify = %q
}
`, name, notify)
}

func testAccViewRecursionEnabled(name, recursionEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_recursion_enabled" {
    name = %q
    recursion_enabled = %q
}
`, name, recursionEnabled)
}

func testAccViewSortList(name string, element string, addressPrioritizedNetworks string) string {
	sourceAdd := ""
	if element == "ip" {
		sourceAdd = "source = \"192.168.11.11\""
	}
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_sort_list" {
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

func testAccViewSynthesizeAddressRecordsFromHttps(name, synthesizeAddressRecordsFromHttps string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_synthesize_address_records_from_https" {
    name = %q
    synthesize_address_records_from_https = %q
}
`, name, synthesizeAddressRecordsFromHttps)
}

func testAccViewTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}

func testAccViewUseForwardersForSubzones(name, useForwardersForSubzones string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_use_forwarders_for_subzones" {
    name = %q
    use_forwarders_for_subzones = %q
}
`, name, useForwardersForSubzones)
}

func testAccViewUseRootForwardersForLocalResolutionWithB1td(name string, useRootForwardersForLocalResolutionWithB1td string, addForwardersBlock bool) string {
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
resource "bloxone_dns_view" "test_use_root_forwarders_for_local_resolution_with_b1td" {
    name = %q
    use_root_forwarders_for_local_resolution_with_b1td = %q
	%s
	
}
`, name, useRootForwardersForLocalResolutionWithB1td, forwardersBlock)
}

func testAccViewZoneAuthority(name string, defaultTTL int64, expire int64, mName string, negativeTTL int64,
	refresh int64, retry int64, rName string, useDefaultMName string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test_zone_authority" {
    name = %q
    zone_authority = {
			default_ttl = %d
			expire = %d
			mname = %q             
			negative_ttl = %d       
			refresh = %d           
			retry = %d             
			rname = %q             
			use_default_mname = %q  
}
}
`, name, defaultTTL, expire, mName, negativeTTL, refresh, retry, rName, useDefaultMName)
}
