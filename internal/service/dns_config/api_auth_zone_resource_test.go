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
// - ACL Type - TSIG Key
// - ACL Type - ACL
// - internal_secondaries
// - nsgs
// - zone_authority : Mname and rname provide inconsistent results after apply
// - view

func TestAccAuthZoneResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig("tf-acc-test.com.", "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "mapping"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol_fqdn"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "initial_soa_serial", "1"),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_auth_zone.test"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthZoneBasicConfig("tf-acc-test.com.", "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					testAccCheckAuthZoneDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAuthZoneResource_FQDN(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v1 dns_config.ConfigAuthZone
	var v2 dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig("tf-acc-test.com.", "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneBasicConfig("tf-infoblox-test.com.", "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-infoblox-test.com."),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_PrimaryType(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v1 dns_config.ConfigAuthZone
	var v2 dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig("tf-acc-test.com.", "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneBasicConfig("tf-acc-test.com.", "external"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "external"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_comment"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneComment("tf-acc-test.com.", "cloud", "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneComment("tf-acc-test.com.", "cloud", "test comment update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_disabled"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneDisabled("tf-acc-test.com.", "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneDisabled("tf-acc-test.com.", "cloud", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_external_primaries"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneExternalPrimaries("tf-acc-test.com.", "external", "tf-infoblox-test.com.", "192.168.10.10", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "tf-infoblox-test.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneExternalPrimaries("tf-acc-test.com.", "external", "tf-infoblox.com.", "192.168.11.11", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
				),
			},
			// Update and Read : External Primaries Type - nsg
			{
				Config:      testAccAuthZoneExternalPrimaries("tf-acc-test.com.", "external", "tf-infoblox-test.com.", "192.168.10.10", "nsg"),
				ExpectError: regexp.MustCompile("External primary type should be 'primary'"),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_external_secondaries"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneExternalSecondaries("tf-acc-test.com.", "external", "tf-infoblox-test.com.", "192.168.10.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.fqdn", "tf-infoblox-test.com."),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "192.168.10.10"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneExternalSecondaries("tf-acc-test.com.", "external", "tf-infoblox.com.", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.fqdn", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "192.168.11.11"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_GssTsigEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_gss_tsig_enabled"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneGssTsigEnabled("tf-acc-test.com.", "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneGssTsigEnabled("tf-acc-test.com.", "cloud", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_inheritance_sources"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneInheritanceSources("tf-acc-test.com.", "cloud", "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "inherit"),
				),
				ExpectNonEmptyPlan: true,
			},
			// Update and Read
			{
				Config: testAccAuthZoneInheritanceSources("tf-acc-test.com.", "cloud", "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "override"),
				),
				ExpectNonEmptyPlan: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_InitialSoaSerial(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_initial_soa_serial"
	var v1 dns_config.ConfigAuthZone
	var v2 dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneInitialSoaSerial("tf-acc-test.com.", "cloud", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "initial_soa_serial", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneInitialSoaSerial("tf-acc-test.com.", "cloud", 2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "initial_soa_serial", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Notify(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_notify"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneNotify("tf-acc-test.com.", "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneNotify("tf-acc-test.com.", "cloud", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_nsgs"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneNsgs("tf-acc-test.com.", "cloud", "bloxone_dns_auth_nsg.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_auth_nsg.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneNsgs("tf-acc-test.com.", "cloud", "bloxone_dns_auth_nsg.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_auth_nsg.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_QueryAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_query_acl"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP("tf-acc-test.com.", "cloud", "query_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny("tf-acc-test.com.", "cloud", "query_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_tags"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneTags("tf-acc-test.com.", "cloud", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneTags("tf-acc-test.com.", "cloud", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_TransferAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_transfer_acl"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP("tf-acc-test.com.", "cloud", "transfer_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny("tf-acc-test.com.", "cloud", "transfer_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_UpdateAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_update_acl"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP("tf-acc-test.com.", "cloud", "update_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny("tf-acc-test.com.", "cloud", "update_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "any"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_UseForwardersForSubzones(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_use_forwarders_for_subzones"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneUseForwardersForSubzones("tf-acc-test.com.", "cloud", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneUseForwardersForSubzones("tf-acc-test.com.", "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_view"
	var v1 dns_config.ConfigAuthZone
	var v2 dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneView("tf-acc-test.com.", "cloud", "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneView("tf-acc-test.com.", "cloud", "bloxone_dns_view.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_ZoneAuthority(t *testing.T) {
	t.Skipf("Mname and rname provide incosistent result after apply")
	var resourceName = "bloxone_dns_auth_zone.test_zone_authority"
	var v dns_config.ConfigAuthZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneZoneAuthority("tf-acc-test.com.", "cloud", 28800, 2419200, "ns.b1ddi", 900,
					10800, 3600, "hostmaster", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.default_ttl", "28800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.expire", "2419200"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.mname", "ns.b1ddi"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.negative_ttl", "900"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.refresh", "10800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.retry", "3600"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.rname", "hostmaster"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.use_default_mname", "false"),
				),
				// mname and rname gets appended with fqdn
				ExpectNonEmptyPlan: true,
			},
			// Update and Read
			{
				Config: testAccAuthZoneZoneAuthority("tf-acc-test.com.", "cloud", 30000, 2519200, "ns.b1ddi", 800, 11800, 3700, "hostmaster", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.default_ttl", "30000"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.expire", "2519200"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.mname", "ns.b1ddi"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.negative_ttl", "800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.refresh", "11800"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.retry", "3700"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.rname", "hostmaster"),
					resource.TestCheckResourceAttr(resourceName, "zone_authority.use_default_mname", "false"),
				),
				// mname and rname gets appended with fqdn
				ExpectNonEmptyPlan: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAuthZoneExists(ctx context.Context, resourceName string, v *dns_config.ConfigAuthZone) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
			AuthZoneRead(ctx, rs.Primary.ID).
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

func testAccCheckAuthZoneDestroy(ctx context.Context, v *dns_config.ConfigAuthZone) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
			AuthZoneRead(ctx, *v.Id).
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

func testAccCheckAuthZoneDisappears(ctx context.Context, v *dns_config.ConfigAuthZone) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
			AuthZoneDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAuthZoneBasicConfig(fqdn, primaryType string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
    fqdn = %q
    primary_type = %q
}
`, fqdn, primaryType)
}

func testAccAuthZoneComment(fqdn, primaryType, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_comment" {
    fqdn = %q
    primary_type = %q
    comment = %q
}
`, fqdn, primaryType, comment)
}

func testAccAuthZoneDisabled(fqdn, primaryType, disabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_disabled" {
    fqdn = %q
    primary_type = %q
    disabled = %q
}
`, fqdn, primaryType, disabled)
}

func testAccAuthZoneExternalPrimaries(fqdn, primaryType, fqdnExternalPrimaries, addressExternalPrimaries, typeExternalPrimaries string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_external_primaries" {
    fqdn = %q
    primary_type = %q
    external_primaries = [
		{
			fqdn = %q
			address = %q
			type = %q
			
		}
]
}
`, fqdn, primaryType, fqdnExternalPrimaries, addressExternalPrimaries, typeExternalPrimaries)
}

func testAccAuthZoneExternalSecondaries(fqdn, primaryType, fqdnExternalSecondaries, addressExternalSecondaries string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_external_secondaries" {
    fqdn = %q
    primary_type = %q
    external_secondaries = [
		{
			fqdn = %q
			address = %q
			
		}
]
}
`, fqdn, primaryType, fqdnExternalSecondaries, addressExternalSecondaries)
}

func testAccAuthZoneGssTsigEnabled(fqdn, primaryType, gssTsigEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_gss_tsig_enabled" {
    fqdn = %q
    primary_type = %q
    gss_tsig_enabled = %q
}
`, fqdn, primaryType, gssTsigEnabled)
}

func testAccAuthZoneInheritanceSources(fqdn, primaryType, action string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_inheritance_sources" {
    fqdn = %[1]q
    primary_type = %[2]q
	inheritance_sources = { 
		gss_tsig_enabled = {
			action = %[3]q
		}
		notify = {
			action = %[3]q
		}
		transfer_acl = {
			action = %[3]q
		}
		useforwardersforsubzones = {
			action = %[3]q
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
`, fqdn, primaryType, action)
}

func testAccAuthZoneInitialSoaSerial(fqdn string, primaryType string, initialSoaSerial int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_initial_soa_serial" {
    fqdn = %q
    primary_type = %q
    initial_soa_serial = %d
}
`, fqdn, primaryType, initialSoaSerial)
}

func testAccAuthZoneNotify(fqdn, primaryType, notify string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_notify" {
    fqdn = %q
    primary_type = %q
    notify = %q
}
`, fqdn, primaryType, notify)
}

func testAccAuthZoneNsgs(fqdn, primaryType, nsgs string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "one"{
	name = %q
}

resource "bloxone_dns_auth_nsg" "two"{
	name = %q
}

resource "bloxone_dns_auth_zone" "test_nsgs" {
	fqdn = %q
    primary_type = %q
    nsgs = [%s.id]
}
`, acctest.RandomNameWithPrefix("auth-nsg"), acctest.RandomNameWithPrefix("auth-nsg"), fqdn, primaryType, nsgs)
}

func testAccAuthZoneAclIP(fqdn, primaryType, aclFieldName, access, address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[3]s" {
	fqdn = %[1]q
    primary_type = %[2]q
    %[3]s = [
		{
			access = %[4]q
			element = "ip"
			address = %[5]q
		}
]
}
`, fqdn, primaryType, aclFieldName, access, address)
}

func testAccAuthZoneAclAny(fqdn, primaryType, aclFieldName, access string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[3]s" {
	fqdn = %[1]q
    primary_type = %[2]q
    %[3]s = [
		{
			access = %[4]q
			element = "any"
		}
]
}
`, fqdn, primaryType, aclFieldName, access)
}

func testAccAuthZoneTags(fqdn string, primaryType string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_tags" {
    fqdn = %q
    primary_type = %q
    tags = %s
}
`, fqdn, primaryType, tagsStr)
}

func testAccAuthZoneUseForwardersForSubzones(fqdn, primaryType, useForwardersForSubzones string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_use_forwarders_for_subzones" {
    fqdn = %q
    primary_type = %q
    use_forwarders_for_subzones = %q
}
`, fqdn, primaryType, useForwardersForSubzones)
}

func testAccAuthZoneView(fqdn, primaryType, view string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "one" {
	name = %q
}

resource "bloxone_dns_view" "two" {
	name = %q
}

resource "bloxone_dns_auth_zone" "test_view" {
	fqdn = %q
    primary_type = %q
    view = %s.id
}
`, acctest.RandomNameWithPrefix("view"), acctest.RandomNameWithPrefix("view"), fqdn, primaryType, view)
}

func testAccAuthZoneZoneAuthority(fqdn string, primaryType string, defaultTTL int64, expire int64, mName string, negativeTTL int64,
	refresh int64, retry int64, rName string, useDefaultMName string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_zone_authority" {
	fqdn = %q
    primary_type = %q
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
`, fqdn, primaryType, defaultTTL, expire, mName, negativeTTL, refresh, retry, rName, useDefaultMName)
}
