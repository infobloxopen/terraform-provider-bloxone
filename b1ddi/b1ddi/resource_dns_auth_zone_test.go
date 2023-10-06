package b1ddi

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_zone"
)

func TestAccResourceDnsAuthZone_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsAuthZoneBasicTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
		resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
			name = "tf_acc_test_dns_view"
		}
		data "b1ddi_dns_hosts" "dns_host" {
			filters = {
				"name" = "%s"
			}
		}
		resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
			internal_secondaries {
				host = data.b1ddi_dns_hosts.dns_host.results.0.id
			}
			fqdn = "tf-acc-test.com."
			primary_type = "cloud"
			view = b1ddi_dns_view.tf_acc_test_dns_view.id
		}`, testAccReadDnsHost(t)),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "comment", ""),
			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "disabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "gss_tsig_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_assigned_hosts.#", "0"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "initial_soa_serial", "1"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "internal_secondaries.0.host"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "notify", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "primary_type", "cloud"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "protocol_fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.#", "0"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "tags"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.#", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "updated_at"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "use_forwarders_for_subzones", "true"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "view"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.default_ttl", "28800"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.expire", "2419200"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.mname", "ns.b1ddi.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.negative_ttl", "900"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_mname", "ns.b1ddi.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_rname", "hostmaster.tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.refresh", "10800"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.retry", "3600"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.rname", "hostmaster@tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.use_default_mname", "true"),
		),
	}
}

func TestAccResourceDnsAuthZone_FullConfigCloud(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneFullConfigCloudTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsAuthZoneFullConfigCloudTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						comment = "This Auth Zone is created by the terraform provider acceptance test"
						disabled = true

						external_secondaries {
							address = "192.168.1.50"
							fqdn = "tf_test_external_secondary."
						}

						fqdn = "tf-acc-test.com."
						gss_tsig_enabled = true

						#inheritance_assigned_hosts {}
						#inheritance_sources {}

						initial_soa_serial = 3

						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}

						notify = true
						
						primary_type = "cloud"

						query_acl {
							access = "deny"
							address = "192.168.1.10"
							element = "ip"
						}

						tags = {
							TestType = "Acceptance"
						}

						transfer_acl {
							access = "allow"
							address = "192.168.1.20"
							element = "ip"
						}
						update_acl {
							access = "allow"
							address = "192.168.1.30"
							element = "ip"
						}

						use_forwarders_for_subzones = false

						view = b1ddi_dns_view.tf_acc_test_dns_view.id

						zone_authority {
							default_ttl = 14400
							expire = 1209600
							mname = "mname.tf-acc-test.com."
							negative_ttl = 700
							refresh = 5400
							retry = 1800
							rname = "rname@tf-acc-test.com"
							use_default_mname = false
						}
						
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_auth_zone_with_is" {
  						internal_secondaries {
    						host = data.b1ddi_dns_hosts.dns_host.results.0.id
  						}
  						fqdn = "tf-example.com."
  						primary_type = "cloud"
  						view = b1ddi_dns_view.tf_example_dns_view.id
  						inheritance_sources {
    						update_acl {
      							action = "override"
    						}
  						} 
					}`, testAccReadDnsHost(t),
		),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "comment", "This Auth Zone is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "disabled", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.#", "0"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.address", "192.168.1.50"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.fqdn", "tf_test_external_secondary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.protocol_fqdn", "tf_test_external_secondary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.stealth", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_enabled", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_key.#"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "gss_tsig_enabled", "true"),
			// ToDo Add check for custom inheritance_assigned_hosts
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_assigned_hosts.#", "0"),
			// ToDo Add check for custom inheritance_sources
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "initial_soa_serial", "3"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "internal_secondaries.0.host"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "notify", "true"),
			// ToDo Add check for custom nsgs
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "nsgs.#", "0"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "primary_type", "cloud"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "protocol_fqdn", "tf-acc-test.com."),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.access", "deny"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.address", "192.168.1.10"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.element", "ip"),

			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "tags.#"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.access", "allow"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.address", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.element", "ip"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.access", "allow"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.address", "192.168.1.30"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.element", "ip"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "updated_at"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "use_forwarders_for_subzones", "false"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "view"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.default_ttl", "14400"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.expire", "1209600"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.mname", "mname.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.negative_ttl", "700"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_mname", "mname.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_rname", "rname.tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.refresh", "5400"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.retry", "1800"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.rname", "rname@tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.use_default_mname", "false"),

			testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_auth_zone_with_is"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "primary_type", "cloud"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.update_acl.action", "override"),
		),
	}
}

func TestAccResourceDnsAuthZone_FullConfigExternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneFullConfigExternalTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsAuthZoneFullConfigExternalTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						comment = "This Auth Zone is created by the terraform provider acceptance test"
						disabled = true

						external_primaries {
							address = "192.168.1.60"
							fqdn = "tf_test_external_primary."
							type = "primary"
						}

						external_secondaries {
							address = "192.168.1.50"
							fqdn = "tf_test_external_secondary."
						}

						fqdn = "tf-acc-test.com."
						gss_tsig_enabled = true

						#inheritance_assigned_hosts {}
						#inheritance_sources {}

						initial_soa_serial = 3

						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}

						notify = true
						
						primary_type = "external"

						query_acl {
							access = "deny"
							address = "192.168.1.10"
							element = "ip"
						}

						tags = {
							TestType = "Acceptance"
						}

						transfer_acl {
							access = "allow"
							address = "192.168.1.20"
							element = "ip"
						}
						update_acl {
							access = "allow"
							address = "192.168.1.30"
							element = "ip"
						}

						use_forwarders_for_subzones = false

						view = b1ddi_dns_view.tf_acc_test_dns_view.id

						zone_authority {
							default_ttl = 14400
							expire = 1209600
							mname = "mname.tf-acc-test.com."
							negative_ttl = 700
							refresh = 5400
							retry = 1800
							rname = "rname@tf-acc-test.com"
							use_default_mname = false
						}
						
					}`, testAccReadDnsHost(t),
		),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "comment", "This Auth Zone is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "disabled", "true"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.address", "192.168.1.60"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.fqdn", "tf_test_external_primary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.nsg", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.protocol_fqdn", "tf_test_external_primary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.tsig_enabled", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.tsig_key.#"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.0.type", "primary"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.address", "192.168.1.50"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.fqdn", "tf_test_external_secondary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.protocol_fqdn", "tf_test_external_secondary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.stealth", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_enabled", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_key.#"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "gss_tsig_enabled", "true"),
			// ToDo Add check for custom inheritance_assigned_hosts
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_assigned_hosts.#", "0"),
			// ToDo Add check for custom inheritance_sources
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "initial_soa_serial", "3"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "internal_secondaries.0.host"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapped_subnet", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapping", "forward"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "notify", "true"),
			// ToDo Add check for custom nsgs
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "nsgs.#", "0"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "primary_type", "external"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "protocol_fqdn", "tf-acc-test.com."),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.access", "deny"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.address", "192.168.1.10"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.element", "ip"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "tags.TestType", "Acceptance"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.access", "allow"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.address", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.element", "ip"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.access", "allow"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.address", "192.168.1.30"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.element", "ip"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "updated_at"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "use_forwarders_for_subzones", "false"),

			resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "view"),

			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.default_ttl", "14400"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.expire", "1209600"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.mname", "mname.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.negative_ttl", "700"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_mname", "mname.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_rname", "rname.tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.refresh", "5400"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.retry", "1800"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.rname", "rname@tf-acc-test.com"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.use_default_mname", "false"),
		),
	}
}

func TestAccResourceDnsAuthZone_UpdateFQDNExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}
						fqdn = "tf-acc-test2.com."
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`, testAccReadDnsHost(t),
				),
				ExpectError: regexp.MustCompile("changing the value of 'fqdn' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
				),
			},
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsAuthZone_UpdateInitialSoaSerialExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}
						fqdn = "tf-acc-test.com."
						initial_soa_serial = 3
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`, testAccReadDnsHost(t),
				),
				ExpectError: regexp.MustCompile("changing the value of 'initial_soa_serial' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "initial_soa_serial", "1"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsAuthZone_UpdatePrimaryTypeExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}
						fqdn = "tf-acc-test.com."
						primary_type = "external"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}`, testAccReadDnsHost(t),
				),
				ExpectError: regexp.MustCompile("changing the value of 'primary_type' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
				),
			},
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsAuthZone_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
						name = "tf_acc_test_dns_view"
					}
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_auth_zone" {
						comment = "This Auth Zone is created by the terraform provider acceptance test"
						disabled = true

						#external_primaries {}
						external_secondaries {
							address = "192.168.1.50"
							fqdn = "tf_test_external_secondary."
						}

						fqdn = "tf-acc-test.com."
						gss_tsig_enabled = true

						#inheritance_assigned_hosts {}
						#inheritance_sources {}

						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}

						notify = true
						
						primary_type = "cloud"

						query_acl {
							access = "deny"
							address = "192.168.1.10"
							element = "ip"
						}

						tags = {
							TestType = "Acceptance"
						}

						transfer_acl {
							access = "allow"
							address = "192.168.1.20"
							element = "ip"
						}
						update_acl {
							access = "allow"
							address = "192.168.1.30"
							element = "ip"
						}

						use_forwarders_for_subzones = false

						view = b1ddi_dns_view.tf_acc_test_dns_view.id

						zone_authority {
							default_ttl = 14400
							expire = 1209600
							mname = "mname.tf-acc-test.com."
							negative_ttl = 700
							refresh = 5400
							retry = 1800
							rname = "rname@tf-acc-test.com"
							use_default_mname = false
						}
						
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsAuthZoneExists("b1ddi_dns_auth_zone.tf_acc_test_auth_zone"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "comment", "This Auth Zone is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "disabled", "true"),
					// ToDo Add check for custom external_primaries
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_primaries.#", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.address", "192.168.1.50"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.fqdn", "tf_test_external_secondary."),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.protocol_fqdn", "tf_test_external_secondary."),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.stealth", "false"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_enabled", "false"),
					resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "external_secondaries.0.tsig_key.#"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "fqdn", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "gss_tsig_enabled", "true"),
					// ToDo Add check for custom inheritance_assigned_hosts
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_assigned_hosts.#", "0"),
					// ToDo Add check for custom inheritance_sources
					resource.TestCheckNoResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "initial_soa_serial", "1"),

					resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "internal_secondaries.0.host"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapped_subnet", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "mapping", "forward"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "notify", "true"),
					// ToDo Add check for custom nsgs
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "nsgs.#", "0"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "parent", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "primary_type", "cloud"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "protocol_fqdn", "tf-acc-test.com."),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.access", "deny"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.address", "192.168.1.10"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "query_acl.0.element", "ip"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "tags.TestType", "Acceptance"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.access", "allow"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.address", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "transfer_acl.0.element", "ip"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.access", "allow"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.address", "192.168.1.30"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "update_acl.0.element", "ip"),

					resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "updated_at"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "use_forwarders_for_subzones", "false"),

					resource.TestCheckResourceAttrSet("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "view"),

					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.default_ttl", "14400"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.expire", "1209600"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.mname", "mname.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.negative_ttl", "700"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_mname", "mname.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.protocol_rname", "rname.tf-acc-test.com"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.refresh", "5400"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.retry", "1800"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.rname", "rname@tf-acc-test.com"),
					resource.TestCheckResourceAttr("b1ddi_dns_auth_zone.tf_acc_test_auth_zone", "zone_authority.0.use_default_mname", "false"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_auth_zone.tf_acc_test_auth_zone",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsAuthZoneExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSConfigurationAPI.AuthZone.AuthZoneRead(
			&auth_zone.AuthZoneReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		if resp.Payload.Result.ID != res.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				res.Primary.ID)
		}
		return nil
	}
}

// Read Dhcp Host name from the env. If env is not specified, skip the test.
func testAccReadDnsHost(t *testing.T) string {
	internalSecondary := os.Getenv("DNS_HOST")

	if internalSecondary == "" {
		t.Skipf("No DNS_HOST env is set for the %s acceptance test", t.Name())
	}

	return internalSecondary
}
