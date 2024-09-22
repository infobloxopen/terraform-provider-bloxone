package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - dhcp_server.client_principal
// - dhcp_server.Inheritance_sources
// - dhcp_Server.Kerberos_kdc
// - dhcp_server.Kerberos_keys
// - dhcp_server.Kerberos_rekey_interval
// - dhcp_server.Kerberos_retry_interval
// - dhcp_server.Kerberos_tkey_lifetime
// - dhcp_server.Kerberos_tkey_protocol
// - dhcp_server.Vendor_specific_option_option_space
// - dhcp_server.Dhcp_options., both v4 & v6

func TestAccServerResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
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
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "client"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost"),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "-"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-zA-Z0-9_.]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_server.test"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
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

func TestAccServerResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_comment"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerComment(name, "TEST_COMMENT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "TEST_COMMENT"),
				),
			},
			// Update and Read
			{
				Config: testAccServerComment(name, "TEST_COMMENT_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "TEST_COMMENT_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsClientUpdate(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_client_update"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp_server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsClientUpdate(name, "client"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "client"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsClientUpdate(name, "server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "server"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsConflictResolutionMode(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_conflict_resolution_mode"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsConflictResolutionMode(name, false, "check_exists_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_exists_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsConflictResolutionMode(name, true, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsDomain(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_domain"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsDomain(name, "test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsDomain(name, "test-update.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "test-update.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DhcpOptions(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_dhcp_options"
	var v1 ipam.Server
	optionSpaceName := acctest.RandomNameWithPrefix("os")
	serverName := acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDhcpOptionsOption(serverName, optionSpaceName, "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDhcpOptionsGroup(serverName, optionSpaceName, "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DhcpOptionsV6(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_dhcp_options"
	var v1 ipam.Server
	optionSpaceName := acctest.RandomNameWithPrefix("os")
	serverName := acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDhcpOptionsOptionV6(serverName, optionSpaceName, "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options_v6.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDhcpOptionsGroupV6(serverName, optionSpaceName, "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsEnabled(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_enabled"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsGenerateName(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_generate_name"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsGenerateName(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsGenerateName(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsGeneratedPrefix(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_generated_prefix"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsGeneratedPrefix(name, "myhost-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost-prefix"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsGeneratedPrefix(name, "myhost-another-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost-another-prefix"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsSendUpdates(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_send_updates"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsSendUpdates(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsSendUpdates(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsTtlPercent(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_ttl_percent"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsTtlPercent(name, "25"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "25"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsTtlPercent(name, "75"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "75"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsUpdateOnRenew(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_update_on_renew"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsUpdateOnRenew(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsUpdateOnRenew(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsUseConflictResolution(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_use_conflict_resolution"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsUseConflictResolution(name, false, "check_exists_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_exists_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsUseConflictResolution(name, true, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsZones(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_zones"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")
	var zoneFQDN = acctest.RandomNameWithPrefix("auth-zone") + "."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,

		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsZones(name, zoneFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zones.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "ddns_zones.0.zone", "bloxone_dns_auth_zone.test_zone", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsZones(name, zoneFQDN),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_zones.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "ddns_zones.0.zone", "bloxone_dns_auth_zone.test_zone", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DhcpConfig(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_dhcp_config"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDhcpConfig(name, true, true, true, 50, 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "50"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDhcpConfig(name, false, false, false, 55, 65),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "false"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "55"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "65"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_GssTsigFallback(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_gss_tsig_fallback"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerGssTsigFallback(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_fallback", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerGssTsigFallback(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "gss_tsig_fallback", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HeaderOptionFilename(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_header_option_filename"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionFilename(name, "HEADER_OPTION.txt"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION.txt"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionFilename(name, "HEADER_OPTION_UPDATE.txt"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_UPDATE.txt"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_header_option_server_address"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionServerAddress(name, "12.0.0.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.4"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionServerAddress(name, "12.0.0.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_header_option_server_name"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionServerName(name, "test-server-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "test-server-1"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionServerName(name, "test-server-2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "test-server-2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HostnameRewriteChar(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_hostname_rewrite_char"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHostnameRewriteChar(name, "+"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "+"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHostnameRewriteChar(name, "/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "/"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HostnameRewriteEnabled(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_hostname_rewrite_enabled"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHostnameRewriteEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHostnameRewriteEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HostnameRewriteRegex(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_hostname_rewrite_regex"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHostnameRewriteRegex(name, "[^a-z]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z]"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHostnameRewriteRegex(name, "[^0-9]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^0-9]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_inheritance_sources"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerInheritanceSources(name, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_hostname_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.allow_unknown.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.allow_unknown_v6.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.filters.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.filters_v6.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.ignore_client_uid.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.lease_time.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.lease_time_v6.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.hostname_rewrite_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.vendor_specific_option_option_space.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccServerInheritanceSources(name, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.allow_unknown.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.allow_unknown_v6.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.filters.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.filters_v6.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.ignore_client_uid.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.lease_time.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.dhcp_config.lease_time_v6.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.hostname_rewrite_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.vendor_specific_option_option_space.action", "override"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_name"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccServerName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_ProfileType(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_comment"
	var v1, v2 ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerProfileType(name, "server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "profile_type", "server"),
				),
			},
			// Update and Read
			{
				Config: testAccServerProfileType(name, "subnet"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerDestroy(context.Background(), &v1),
					testAccCheckServerExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "profile_type", "subnet"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_ServerPrincipal(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_server_principal"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerServerPrincipal(name, "test.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_principal", "test.com"),
				),
			},
			// Update and Read
			{
				Config: testAccServerServerPrincipal(name, "test-update.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_principal", "test-update.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_tags"
	var v ipam.Server
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
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
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
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
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckServerExists(ctx context.Context, resourceName string, v *ipam.Server) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			ServerAPI.
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

func testAccCheckServerDestroy(ctx context.Context, v *ipam.Server) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			ServerAPI.
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

func testAccCheckServerDisappears(ctx context.Context, v *ipam.Server) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			ServerAPI.
			Delete(ctx, *v.Id).
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
resource "bloxone_dhcp_server" "test" {
    name = %q
}
`, name)
}

func testAccServerComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccServerDdnsClientUpdate(name string, ddnsClientUpdate string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_client_update" {
    name = %q
    ddns_client_update = %q
}
`, name, ddnsClientUpdate)
}

func testAccServerDdnsConflictResolutionMode(name string, ddnsUseConflictResolution bool, ddnsConflictResolutionMode string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_conflict_resolution_mode" {
    name = %q
	ddns_use_conflict_resolution = %t
    ddns_conflict_resolution_mode = %q
}
`, name, ddnsUseConflictResolution, ddnsConflictResolutionMode)
}

func testAccServerDdnsDomain(name string, ddnsDomain string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_domain" {
    name = %q
    ddns_domain = %q
}
`, name, ddnsDomain)
}

func testAccServerDhcpOptionsOption(name string, optionSpaceName, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options" {
    name = %q
    dhcp_options = [
      {
       type = %q
       option_code = bloxone_dhcp_option_code.test.id
       option_value = %q
      }
    ]
}
`, name, optionItemType, optValue)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+optionSpaceName, optionSpaceName, "ip4"), config}, "")
}

func testAccServerDhcpOptionsGroup(name string, optionSpaceName, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options" {
    name = %q
    dhcp_options = [
      {
       type = %q
       group = bloxone_dhcp_option_group.test.id
      }
    ]
}
`, name, optionItemType)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+optionSpaceName, optionSpaceName, "ip4"), config}, "")
}

func testAccServerDhcpOptionsOptionV6(name string, optionSpaceName, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options" {
    name = %q
    dhcp_options_v6 = [
      {
       type = %q
       option_code = bloxone_dhcp_option_code.test.id
       option_value = %q
      }
    ]
}
`, name, optionItemType, optValue)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+optionSpaceName, optionSpaceName, "ip6"), config}, "")
}

func testAccServerDhcpOptionsGroupV6(name string, optionSpaceName, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options" {
    name = %q
    dhcp_options_v6 = [
      {
       type = %q
       group = bloxone_dhcp_option_group.test.id
      }
    ]
}
`, name, optionItemType)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+optionSpaceName, optionSpaceName, "ip6"), config}, "")
}

func testAccServerDdnsEnabled(name string, ddnsEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_enabled" {
    name = %q
    ddns_enabled = %q
}
`, name, ddnsEnabled)
}

func testAccServerDdnsGenerateName(name string, ddnsGenerateName string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_generate_name" {
    name = %q
    ddns_generate_name = %q
}
`, name, ddnsGenerateName)
}

func testAccServerDdnsGeneratedPrefix(name string, ddnsGeneratedPrefix string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_generated_prefix" {
    name = %q
    ddns_generated_prefix = %q
}
`, name, ddnsGeneratedPrefix)
}

func testAccServerDdnsSendUpdates(name string, ddnsSendUpdates string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_send_updates" {
    name = %q
    ddns_send_updates = %q
}
`, name, ddnsSendUpdates)
}

func testAccServerDdnsTtlPercent(name string, ddnsTtlPercent string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_ttl_percent" {
    name = %q
    ddns_ttl_percent = %q
}
`, name, ddnsTtlPercent)
}

func testAccServerDdnsUpdateOnRenew(name string, ddnsUpdateOnRenew string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_update_on_renew" {
    name = %q
    ddns_update_on_renew = %q
}
`, name, ddnsUpdateOnRenew)
}

func testAccServerDdnsUseConflictResolution(name string, ddnsUseConflictResolution bool, ddnsConflictResolutionMode string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_use_conflict_resolution" {
	name = %q
	ddns_use_conflict_resolution = %t
	ddns_conflict_resolution_mode = %q
}
`, name, ddnsUseConflictResolution, ddnsConflictResolutionMode)
}

func testAccServerDdnsZones(name string, zoneFQDN string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_zone" "test_zone" {
	fqdn = %q
	primary_type = "cloud"
}
resource "bloxone_dhcp_server" "test_ddns_zones" {
	name = %q
	ddns_zones = [{
		zone = bloxone_dns_auth_zone.test_zone.id
	}]
	depends_on = [bloxone_dns_auth_zone.test_zone]
}
`, zoneFQDN, name)
}

func testAccServerDhcpConfig(name string, allowUnknown, allowUnknownV6, ignoreClientUid bool,
	leaseTime, leaseTimeV6 int) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_config" {
    name = %q
    dhcp_config = {
		allow_unknown = %t
		allow_unknown_v6 = %t
		ignore_client_uid = %t
		lease_time = %d
		lease_time_v6 = %d
	}
}
`, name, allowUnknown, allowUnknownV6, ignoreClientUid, leaseTime, leaseTimeV6)
}

func testAccServerGssTsigFallback(name string, gssTsigFallback string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_gss_tsig_fallback" {
    name = %q
    gss_tsig_fallback = %q
}
`, name, gssTsigFallback)
}

func testAccServerHeaderOptionFilename(name string, headerOptionFilename string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_header_option_filename" {
    name = %q
    header_option_filename = %q
}
`, name, headerOptionFilename)
}

func testAccServerHeaderOptionServerAddress(name string, headerOptionServerAddress string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_header_option_server_address" {
    name = %q
    header_option_server_address = %q
}
`, name, headerOptionServerAddress)
}

func testAccServerHeaderOptionServerName(name string, headerOptionServerName string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_header_option_server_name" {
    name = %q
    header_option_server_name = %q
}
`, name, headerOptionServerName)
}

func testAccServerHostnameRewriteChar(name string, hostnameRewriteChar string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_hostname_rewrite_char" {
    name = %q
    hostname_rewrite_char = %q
}
`, name, hostnameRewriteChar)
}

func testAccServerHostnameRewriteEnabled(name string, hostnameRewriteEnabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_hostname_rewrite_enabled" {
    name = %q
    hostname_rewrite_enabled = %q
}
`, name, hostnameRewriteEnabled)
}

func testAccServerHostnameRewriteRegex(name string, hostnameRewriteRegex string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_hostname_rewrite_regex" {
    name = %q
    hostname_rewrite_regex = %q
}
`, name, hostnameRewriteRegex)
}

func testAccServerInheritanceSources(name, action string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_inheritance_sources" {
	name = %[1]q
	inheritance_sources = {
		ddns_block = {
			action = %[2]q
		}
		ddns_client_update = {
			action = %[2]q
		}
		ddns_conflict_resolution_mode = {
			action = %[2]q
		}
		ddns_hostname_block = {
			action = %[2]q
		}
		ddns_ttl_percent = {
			action = %[2]q
		}
		ddns_update_on_renew = {
			action = %[2]q
		}
		ddns_use_conflict_resolution = {
			action = %[2]q
		}
		dhcp_config = {
			allow_unknown = {
				action = %[2]q
			}
			allow_unknown_v6 = {
				action = %[2]q
			}
			filters	= {
				action = %[2]q
			}
			filters_v6	= {
				action = %[2]q
			}
			ignore_client_uid = {
				action = %[2]q
			}
			ignore_list	= {
				action = %[2]q
			}
			lease_time = {
				action = %[2]q
			}
			lease_time_v6 = {
				action = %[2]q
			}
		}
		header_option_filename = {
			action = %[2]q
		}
		header_option_server_address = {
			action = %[2]q
		}
		header_option_server_name = {
			action = %[2]q
		}
		hostname_rewrite_block = {
			action = %[2]q
		}
		vendor_specific_option_option_space = {
			action = %[2]q
		}
	}

}
`, name, action)
}

func testAccServerName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_name" {
    name = %q
}
`, name)
}

func testAccServerServerPrincipal(name string, serverPrincipal string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_server_principal" {
    name = %q
    server_principal = %q
}
`, name, serverPrincipal)
}

func testAccServerProfileType(name string, profileType string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_comment" {
    name = %q
    profile_type = %q
}
`, name, profileType)
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
resource "bloxone_dhcp_server" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
