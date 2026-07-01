package dns_config_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - zone_authority : Mname and rname provide inconsistent results after apply

func TestAccAuthZoneV2Resource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2BasicConfig(fqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
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

func TestAccAuthZoneV2Resource_disappears(t *testing.T) {

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
				Config: testAccAuthZoneV2BasicConfig(fqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					testAccCheckAuthZoneDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAuthZoneV2Resource_FQDN(t *testing.T) {
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
				Config: testAccAuthZoneV2BasicConfig(fqdn1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn1),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2BasicConfig(fqdn2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneDestroy(context.Background(), &v1),
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_comment"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2Comment(fqdn, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2Comment(fqdn, "test comment update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_disabled"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2Disabled(fqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2Disabled(fqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_GssTsigEnabled(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_gss_tsig_enabled"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2GssTsigEnabled(fqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2GssTsigEnabled(fqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_inheritance_sources"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2InheritanceSources(fqdn, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2InheritanceSources(fqdn, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.gss_tsig_enabled.action", "override"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_InitialSoaSerial(t *testing.T) {
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
				Config: testAccAuthZoneV2InitialSoaSerial(fqdn, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "initial_soa_serial", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2InitialSoaSerial(fqdn, 2),
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

func TestAccAuthZoneV2Resource_Notify(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_notify"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2Notify(fqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2Notify(fqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_Nameservers(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_nameservers"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2Nameservers(fqdn, "1.1.1.1", "a.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.fqdn", "a.com."),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.origin", "external"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.role", "primary"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2Nameservers(fqdn, "2.2.2.2", "b.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nameservers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.fqdn", "b.com."),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.origin", "external"),
					resource.TestCheckResourceAttr(resourceName, "nameservers.0.role", "primary"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_Nsg(t *testing.T) {
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
				Config: testAccAuthZoneV2Nsg(fqdn, nsg1Name, nsg2Name, "bloxone_dns_auth_nsg.nsg_one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsg", "bloxone_dns_auth_nsg.nsg_one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2Nsg(fqdn, nsg1Name, nsg2Name, "bloxone_dns_auth_nsg.nsg_two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsg", "bloxone_dns_auth_nsg.nsg_two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_QueryAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_query_acl"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2AclIP(fqdn, "query_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAny(fqdn, "query_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAcl(fqdn, "query_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "query_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "query_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneV2AclTsigKey(fqdn, "query_acl", "deny"),
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

func TestAccAuthZoneV2Resource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_tags"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2Tags(fqdn, map[string]string{
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
				Config: testAccAuthZoneV2Tags(fqdn, map[string]string{
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

func TestAccAuthZoneV2Resource_TransferAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_transfer_acl"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2AclIP(fqdn, "transfer_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAny(fqdn, "transfer_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAcl(fqdn, "transfer_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "transfer_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "transfer_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneV2AclTsigKey(fqdn, "transfer_acl", "deny"),
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

func TestAccAuthZoneV2Resource_UpdateAcl(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_update_acl"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2AclIP(fqdn, "update_acl", "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAny(fqdn, "update_acl", "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "any"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2AclAcl(fqdn, "update_acl"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_acl.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "update_acl.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAuthZoneV2AclTsigKey(fqdn, "update_acl", "deny"),
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

func TestAccAuthZoneV2Resource_UseForwardersForSubzones(t *testing.T) {
	var resourceName = "bloxone_dns_auth_zone.test_use_forwarders_for_subzones"
	var v dnsconfig.AuthZone
	var fqdn = acctest.RandomNameWithPrefix("auth-zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthZoneV2UseForwardersForSubzones(fqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2UseForwardersForSubzones(fqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_forwarders_for_subzones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthZoneV2Resource_View(t *testing.T) {
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
				Config: testAccAuthZoneV2View(fqdn, "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthZoneV2View(fqdn, "bloxone_dns_view.two"),
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

func TestAccAuthZoneV2Resource_ZoneAuthority(t *testing.T) {
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
				Config: testAccAuthZoneV2ZoneAuthority(fqdn, 28800, 2419200, "ns.b1ddi", 900,
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
				Config: testAccAuthZoneV2ZoneAuthority(fqdn, 30000, 2519200, "ns.b1ddi", 800, 11800, 3700, "hostmaster", "false"),
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

func testAccAuthZoneV2AclAcl(fqdn, aclFieldName string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[2]s" {
    fqdn = %[1]q
    %[2]s = [
		{
			element = "acl"
			acl = bloxone_dns_acl.test.id
		}
]
}
`, fqdn, aclFieldName)
	return strings.Join([]string{testAccAclBasicConfig(acctest.RandomNameWithPrefix("acl")), config}, "")
}

func testAccAuthZoneV2AclAny(fqdn, aclFieldName, access string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[2]s" {
    fqdn = %[1]q
    %[2]s = [
		{
			access = %[3]q
			element = "any"
		}
]
}
`, fqdn, aclFieldName, access)
}

func testAccAuthZoneV2AclIP(fqdn, aclFieldName, access, address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[2]s" {
    fqdn = %[1]q
    %[2]s = [
		{
			access = %[3]q
			element = "ip"
			address = %[4]q
		}
]
}
`, fqdn, aclFieldName, access, address)
}

func testAccAuthZoneV2AclTsigKey(fqdn, aclFieldName, access string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_%[2]s" {
    fqdn = %[1]q
    %[2]s = [
		{
			element = "tsig_key"
			access = %[3]q
			tsig_key = {
				key = bloxone_keys_tsig.test.id
			}
		}
]
}
`, fqdn, aclFieldName, access)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+acctest.RandomNameWithPrefix("auth-zone"), "acl-"+fqdn), config}, "")

}

func testAccAuthZoneV2BasicConfig(fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test" {
    fqdn = %q
}
`, fqdn)
}

func testAccAuthZoneV2Comment(fqdn, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_comment" {
    fqdn = %q
    comment = %q
}
`, fqdn, comment)
}

func testAccAuthZoneV2Disabled(fqdn, disabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_disabled" {
    fqdn = %q
    disabled = %q
}
`, fqdn, disabled)
}

func testAccAuthZoneV2GssTsigEnabled(fqdn, gssTsigEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_gss_tsig_enabled" {
    fqdn = %q
    gss_tsig_enabled = %q
}
`, fqdn, gssTsigEnabled)
}

func testAccAuthZoneV2InheritanceSources(fqdn, action string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_inheritance_sources" {
    fqdn = %[1]q
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
`, fqdn, action)
}

func testAccAuthZoneV2InitialSoaSerial(fqdn string, initialSoaSerial int64) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_initial_soa_serial" {
    fqdn = %q
    initial_soa_serial = %d
}
`, fqdn, initialSoaSerial)
}

func testAccAuthZoneV2Notify(fqdn, notify string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_notify" {
    fqdn = %q
    notify = %q
}
`, fqdn, notify)
}

func testAccAuthZoneV2Nameservers(fqdn, address, nsqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_nameservers" {
    fqdn         = %q
    nameservers = [
        {
            address = %q
            fqdn    = %q
            role    = "primary"
        }
    ]
}
`, fqdn, address, nsqdn)
}

func testAccAuthZoneV2Nsg(fqdn, nsg1Name, nsg2Name, nsgRef string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "nsg_one" {
    name = %q
	nameservers = [
        {
            address = "1.14.13.220"
            fqdn    = "cc.com."
            role    = "primary"
        }
    ]
}

resource "bloxone_dns_auth_nsg" "nsg_two" {
    name = %q
	nameservers = [
        {
            address = "1.14.13.222"
            fqdn    = "dd.com."
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

func testAccAuthZoneV2Tags(fqdn string, tags map[string]string) string {
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
    tags = %s
}
`, fqdn, tagsStr)
}

func testAccAuthZoneV2UseForwardersForSubzones(fqdn, useForwardersForSubzones string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_use_forwarders_for_subzones" {
    fqdn = %q
    use_forwarders_for_subzones = %q
}
`, fqdn, useForwardersForSubzones)
}

func testAccAuthZoneV2View(fqdn, view string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "one" {
	name = %q
}

resource "bloxone_dns_view" "two" {
	name = %q
}

resource "bloxone_dns_auth_zone" "test_view" {
	fqdn = %q
    view = %s.id
}
`, acctest.RandomNameWithPrefix("view"), acctest.RandomNameWithPrefix("view"), fqdn, view)
}

func testAccAuthZoneV2ZoneAuthority(fqdn string, defaultTTL int64, expire int64, mName string, negativeTTL int64,
	refresh int64, retry int64, rName string, useDefaultMName string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_zone_authority" {
	fqdn = %q
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
`, fqdn, defaultTTL, expire, mName, negativeTTL, refresh, retry, rName, useDefaultMName)
}
