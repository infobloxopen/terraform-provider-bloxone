package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - internal_secondaries
// - nsgs
// - zone_authority : Mname and rname provide inconsistent results after apply

func TestAccAuthZoneResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig(fqdn, "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
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

	t.Skip("Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]")

	resourceName := "bloxone_dns_auth_zone.test"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthZoneBasicConfig(fqdn, "cloud"),
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
	var v1 dnsconfig.AuthZone
	var v2 dnsconfig.AuthZone
	var fqdn1 = acctest.RandomNameWithPrefix("auth-zone") + ".com."
	var fqdn2 = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig(fqdn1, "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn1),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneBasicConfig(fqdn2, "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn2),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_PrimaryType(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v1 dnsconfig.AuthZone
	var v2 dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneBasicConfig(fqdn, "cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneBasicConfig(fqdn, "external"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "external"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_CompartmentId(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_compartment_id"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneCompartmentId(fqdn, "cloud", "c4695."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", "c4695."),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneCompartmentId(fqdn, "cloud", ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
					resource.TestCheckResourceAttr(resourceName, "primary_type", "cloud"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_comment"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneComment(fqdn, "cloud", "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneComment(fqdn, "cloud", "test comment update"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneDisabled(fqdn, "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneDisabled(fqdn, "cloud", "true"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneExternalPrimaries(fqdn, "external", "tf-infoblox-test.com.", "192.168.10.10", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "tf-infoblox-test.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneExternalPrimaries(fqdn, "external", "tf-infoblox.com.", "192.168.11.11", "primary"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "tf-infoblox.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
				),
			},
			// Update and Read : External Primaries Type - nsg
			{
				Config:      testAccAuthZoneExternalPrimaries(fqdn, "external", "tf-infoblox-test.com.", "192.168.10.10", "nsg"),
				ExpectError: regexp.MustCompile("External primary type should be 'primary'"),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_external_secondaries"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneExternalSecondaries(fqdn, "external", "tf-infoblox-test.com.", "192.168.10.10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.fqdn", "tf-infoblox-test.com."),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "192.168.10.10"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneExternalSecondaries(fqdn, "external", "tf-infoblox.com.", "192.168.11.11"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneGssTsigEnabled(fqdn, "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneGssTsigEnabled(fqdn, "cloud", "true"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneInheritanceSources(fqdn, "cloud", "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneInheritanceSources(fqdn, "cloud", "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "override"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_InitialSoaSerial(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_initial_soa_serial"
	var v1 dnsconfig.AuthZone
	var v2 dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneInitialSoaSerial(fqdn, "cloud", 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "initial_soa_serial", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneInitialSoaSerial(fqdn, "cloud", 2),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneNotify(fqdn, "cloud", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneNotify(fqdn, "cloud", "true"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneNsgs(fqdn, "cloud", "bloxone_dns_auth_nsg.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_auth_nsg.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneNsgs(fqdn, "cloud", "bloxone_dns_auth_nsg.two"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP(fqdn, "cloud", "query_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny(fqdn, "cloud", "query_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAcl(fqdn, "cloud", "query_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "query_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneAclTsigKey(fqdn, "cloud", "query_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "query_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_tags"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneTags(fqdn, "cloud", map[string]string{
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
				Config: testAccAuthZoneTags(fqdn, "cloud", map[string]string{
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP(fqdn, "cloud", "transfer_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny(fqdn, "cloud", "transfer_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAcl(fqdn, "cloud", "transfer_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "transfer_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneAclTsigKey(fqdn, "cloud", "transfer_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "transfer_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_UpdateAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_update_acl"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneAclIP(fqdn, "cloud", "update_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAny(fqdn, "cloud", "update_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneAclAcl(fqdn, "cloud", "update_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "update_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneAclTsigKey(fqdn, "cloud", "update_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "update_acl.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_UseForwardersForSubzones(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_use_forwarders_for_subzones"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneUseForwardersForSubzones(fqdn, "cloud", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneUseForwardersForSubzones(fqdn, "cloud", "false"),
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
	var v1 dnsconfig.AuthZone
	var v2 dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneView(fqdn, "cloud", "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneView(fqdn, "cloud", "bloxone_dns_view.two"),
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
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneZoneAuthority(fqdn, "cloud", 28800, 2419200, "ns.b1ddi", 900,
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
				Config: testAccAuthZoneZoneAuthority(fqdn, "cloud", 30000, 2519200, "ns.b1ddi", 800, 11800, 3700, "hostmaster", "false"),
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

func TestAccAuthZoneResource_DnssecSigningPolicy(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_dnssec_signing_policy"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneDnssecSigningPolicy(fqdn, "cloud", "NSEC"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_signing_policy.nsec_type", "NSEC"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneDnssecSigningPolicy(fqdn, "cloud", "NSEC3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_signing_policy.nsec_type", "NSEC3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_GridPrimaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_grid_primaries"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneGridPrimaries(fqdn, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primaries.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "grid_primaries.0.host", "data.bloxone_dns_hosts.three", "results.0.id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneGridPrimaries(fqdn, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primaries.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "grid_primaries.0.host", "data.bloxone_dns_hosts.three", "results.0.id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_GridSecondaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_grid_secondaries"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneGridSecondaries(fqdn, "one", "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "grid_secondaries.0.host", "data.bloxone_dns_hosts.one", "results.0.id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneGridSecondaries(fqdn, "two", "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "grid_secondaries.0.host", "data.bloxone_dns_hosts.two", "results.0.id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_MaxRecordsPerType(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_max_records_per_type"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneMaxRecordsPerType(fqdn, "cloud", 100),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_records_per_type", "100"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneMaxRecordsPerType(fqdn, "cloud", 500),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_records_per_type", "500"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_MaxTypesPerName(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_max_types_per_name"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneMaxTypesPerName(fqdn, "cloud", 50),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_types_per_name", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneMaxTypesPerName(fqdn, "cloud", 200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "max_types_per_name", "200"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Nameservers(t *testing.T) {
	t.Skip("Skipping Tests for NameServers until for v1 Phase 1")
	var resourceName = "bloxone_dns_auth_zone.test_nameservers"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read - basic nameserver with role and external origin
			{
				Config: testAccAuthZoneNameservers(fqdn, "1.1.1.1", "a.com.", "primary", "false", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.fqdn", "a.com."),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.origin", "external"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.role", "primary"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.stealth", "false"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.tsig_enabled", "false"),
				),
			},
			// Update - with stealth enabled
			{
				Config: testAccAuthZoneNameservers(fqdn, "2.2.2.2", "b.com.", "secondary", "true", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.fqdn", "b.com."),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.origin", "cloud"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.role", "primary"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.stealth", "true"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.tsig_enabled", "false"),
				),
			},
			// Update - with tsig_enabled enabled
			{
				Config: testAccAuthZoneNameservers(fqdn, "3.3.3.3", "c.com.", "secondary", "true", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.address", "3.3.3.3"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.fqdn", "c.com."),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.role", "primary"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.stealth", "true"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.tsig_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_Nsg(t *testing.T) {
	t.Skip("Skipping Tests for NSG until for v1 Phase 1")
	var resourceName = "bloxone_dns_auth_zone.test_nsg"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."
	nsg1Name := acctest.RandomNameWithPrefix("auth-nsg")
	nsg2Name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneNsg(fqdn, nsg1Name, nsg2Name, "bloxone_dns_auth_nsg.nsg_one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsg", "bloxone_dns_auth_nsg.nsg_one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneNsg(fqdn, nsg1Name, nsg2Name, "bloxone_dns_auth_nsg.nsg_two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsg", "bloxone_dns_auth_nsg.nsg_two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneResource_SecondaryZoneRecordsSync(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_secondary_zone_records_sync"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneSecondaryZoneRecordsSync(fqdn, "external", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_zone_records_sync", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneSecondaryZoneRecordsSync(fqdn, "external", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "secondary_zone_records_sync", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAuthZoneExists(ctx context.Context, resourceName string, v *dnsconfig.AuthZone) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
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

func testAccCheckAuthZoneDestroy(ctx context.Context, v *dnsconfig.AuthZone) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
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

func testAccCheckAuthZoneDisappears(ctx context.Context, v *dnsconfig.AuthZone) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthZoneAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAuthZoneAclAcl(fqdn, primaryType, aclFieldName string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[3]s" {
    fqdn = %[1]q
    primary_type = %[2]q
    %[3]s = [
		{
			element = "acl"
			acl = bloxone_dns_acl.test.id
		}
]
}
`, fqdn, primaryType, aclFieldName)
	return strings.Join([]string{testAccAclBasicConfig(acctest.RandomNameWithPrefix("acl")), config}, "")
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

func testAccAuthZoneAclTsigKey(fqdn, primaryType, aclFieldName, access string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[3]s" {
    fqdn = %[1]q
    primary_type = %[2]q
    %[3]s = [
		{
			element = "tsig_key"
			access = %[4]q
			tsig_key = {
				key = bloxone_keys_tsig.test.id
			}
		}
]
}
`, fqdn, primaryType, aclFieldName, access)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+acctest.RandomNameWithPrefix("auth-zone"), "acl-"+fqdn), config}, "")

}

func testAccAuthZoneBasicConfig(fqdn, primaryType string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
    fqdn = %q
    primary_type = %q
}
`, fqdn, primaryType)
}

func testAccAuthZoneCompartmentId(fqdn, primaryType, compartmentId string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_compartment_id" {
    fqdn = %q
    primary_type = %q
    compartment_id = %q
}
`, fqdn, primaryType, compartmentId)
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
	name = "one"
}

resource "bloxone_dns_auth_nsg" "two"{
	name = "two"
}

resource "bloxone_dns_auth_zone" "test_nsgs" {
	fqdn = %q
    primary_type = %q
    nsgs = [%s.id]
}
`, fqdn, primaryType, nsgs)
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

func testAccAuthZoneDnssecSigningPolicy(fqdn, primaryType, nsecType string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_dnssec_signing_policy" {
    fqdn         = %q
    primary_type = %q
    dnssec_signing_policy = {
        nsec_type = %q
    }
}
`, fqdn, primaryType, nsecType)
}

func testAccAuthZoneGridPrimaries(fqdn, host string) string {
	config := fmt.Sprintf(`
data "bloxone_dns_views" "test_grid_secondaries" {
  tag_filters = {
	"nios/imported" = "true"
  }
}

data "bloxone_dns_hosts" "three" {
    tag_filters = {
        "host/deployment_type" = "CNIOS"
    }
}

resource "bloxone_dns_auth_zone" "test_grid_primaries" {
    fqdn         = %q
    primary_type = "cloud"
    grid_primaries = [
        {
            host = data.bloxone_dns_hosts.three.results.0.id
        }
    ]
	view = data.bloxone_dns_views.test_grid_secondaries.results.0.id
}
`, fqdn)
	return strings.Join([]string{testAccBaseWithHost(), config}, "")
}

func testAccAuthZoneGridSecondaries(fqdn, host, hostPrimary string) string {
	config := fmt.Sprintf(`
data "bloxone_dns_views" "test_grid_secondaries" {
  tag_filters = {
	"nios/imported" = "true"
  }
}

data "bloxone_dns_hosts" "three" {
    tag_filters = {
        "host/deployment_type" = "CNIOS"
    }
}

resource "bloxone_dns_auth_zone" "test_grid_secondaries" {
    fqdn         = %q
    primary_type = "cloud"
    grid_secondaries = [
        {
            host = data.bloxone_dns_hosts.three.results.0.id
        }
    ]
	grid_primaries = [
        {
            host = data.bloxone_dns_hosts.three.results.0.id
        }
    ]
	view = data.bloxone_dns_views.test_grid_secondaries.results.0.id
}
`, fqdn)
	return strings.Join([]string{testAccBaseWithHost(), config}, "")
}

func testAccAuthZoneMaxRecordsPerType(fqdn, primaryType string, maxRecords int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_max_records_per_type" {
    fqdn                = %q
    primary_type        = %q
    max_records_per_type = %d
}
`, fqdn, primaryType, maxRecords)
}

func testAccAuthZoneMaxTypesPerName(fqdn, primaryType string, maxTypes int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_max_types_per_name" {
    fqdn             = %q
    primary_type     = %q
    max_types_per_name = %d
}
`, fqdn, primaryType, maxTypes)
}

func testAccAuthZoneNameservers(fqdn, address, nsqdn, role, stealth, tsigEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_nameservers" {
    fqdn         = %q
	primary_type = "cloud"
    nameservers = [
        {
            address      = %q
            fqdn         = %q
            role         = %q
            tsig_enabled = %s
        }
    ]
}
`, fqdn, address, nsqdn, role, tsigEnabled)
}

func testAccAuthZoneNsg(fqdn, nsg1Name, nsg2Name, nsgRef string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "nsg_one" {
    name = %q
	nameservers = [
        {
            address = "1.4.3.22"
            fqdn    = "aa.com."
            role    = "primary"
        }
    ]
}

resource "bloxone_dns_auth_nsg" "nsg_two" {
    name = %q
	nameservers = [
        {
            address = "1.4.3.22"
            fqdn    = "bb.com."
            role    = "primary"
        }
    ]
}

resource "bloxone_dns_auth_zone" "test_nsg" {
    fqdn         = %q
    nsg          = %s.id
}
`, nsg1Name, nsg2Name, fqdn, nsgRef)
}

func testAccAuthZoneSecondaryZoneRecordsSync(fqdn, primaryType, secondaryZoneRecordsSync string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_secondary_zone_records_sync" {
    fqdn                       = %q
    primary_type               = %q
    secondary_zone_records_sync = %s
}
`, fqdn, primaryType, secondaryZoneRecordsSync)
}
