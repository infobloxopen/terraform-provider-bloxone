package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

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
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_server.test"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

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

func TestAccServerResource_ClientPrincipal(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_client_principal"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")
	var clientPrincipal = acctest.RandomNameWithPrefix("CLIENT_PRINCIPAL")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerClientPrincipal(name, clientPrincipal),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_principal", clientPrincipal),
				),
			},
			// Update and Read
			{
				Config: testAccServerClientPrincipal("name", "CLIENT_PRINCIPAL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "client_principal", "CLIENT_PRINCIPAL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_comment"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp_server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsConflictResolutionMode(name, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsConflictResolutionMode(name, "no_check_with_dhcid"),
				//Check: resource.ComposeTestCheckFunc(
				//	testAccCheckServerExists(context.Background(), resourceName, &v),
				//	resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				//	),
				ExpectError: regexp.MustCompile("Conflicting values provided"),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsDomain(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_domain"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsDomain(name, "Test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "Test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsDomain(name, "Test_update.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "Test_update.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsEnabled(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_enabled"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsGeneratedPrefix(name, "myhost"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsGeneratedPrefix(name, "myhost.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "myhost.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsSendUpdates(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_send_updates"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsSendUpdates(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsSendUpdates(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsTtlPercent(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_ttl_percent"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDdnsUseConflictResolution(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDdnsUseConflictResolution(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_DdnsZones(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_ddns_zones"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")
	var zoneFQDN = acctest.RandomNameWithPrefix("auth-zone") + "."

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerDhcpConfig("NAME_REPLACE_ME", "DHCP_CONFIG_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config", "DHCP_CONFIG_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerDhcpConfig("NAME_REPLACE_ME", "DHCP_CONFIG_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config", "DHCP_CONFIG_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_GssTsigFallback(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_gss_tsig_fallback"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionFilename(name, "TEST_HEADER_OPTION_FILENAME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "TEST_HEADER_OPTION_FILENAME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionFilename(name, "TEST_HEADER_OPTION_FILENAME_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "TEST_HEADER_OPTION_FILENAME_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_header_option_server_address"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionServerAddress(name, "192.28.4.3"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "192.28.4.3"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionServerAddress(name, "192.28.4.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "192.28.4.4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_header_option_server_name"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHeaderOptionServerName(name, "TEST_HEADER_OPTION_SERVER_NAME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "TEST_HEADER_OPTION_SERVER_NAME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHeaderOptionServerName(name, "TEST_HEADER_OPTION_SERVER_NAME_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "TEST_HEADER_OPTION_SERVER_NAME_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HostnameRewriteChar(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_hostname_rewrite_char"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHostnameRewriteChar(name, "#"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "#"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHostnameRewriteChar(name, "*"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "*"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_HostnameRewriteEnabled(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_hostname_rewrite_enabled"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerHostnameRewriteRegex(name, "[^a-z_.]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z_.]"),
				),
			},
			// Update and Read
			{
				Config: testAccServerHostnameRewriteRegex(name, "[^a-zA-Z0-9_.]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-zA-Z0-9_.]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_inheritance_sources"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerInheritanceSources(name, "INHERITANCE_SOURCES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources", "INHERITANCE_SOURCES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerInheritanceSources(name, "INHERITANCE_SOURCES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources", "INHERITANCE_SOURCES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosKdc(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_kdc"
	var v ipam.IpamsvcServer

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosKdc("NAME_REPLACE_ME", "KERBEROS_KDC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_kdc", "KERBEROS_KDC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosKdc("NAME_REPLACE_ME", "KERBEROS_KDC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_kdc", "KERBEROS_KDC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosKeys(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_keys"
	var v ipam.IpamsvcServer

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosKeys("NAME_REPLACE_ME", "KERBEROS_KEYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_keys", "KERBEROS_KEYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosKeys("NAME_REPLACE_ME", "KERBEROS_KEYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_keys", "KERBEROS_KEYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosRekeyInterval(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_rekey_interval"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosRekeyInterval(name, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_rekey_interval", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosRekeyInterval(name, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_rekey_interval", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosRetryInterval(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_retry_interval"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosRetryInterval(name, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_retry_interval", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosRetryInterval(name, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_retry_interval", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosTkeyLifetime(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_tkey_lifetime"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosTkeyLifetime(name, "10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_tkey_lifetime", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosTkeyLifetime(name, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_tkey_lifetime", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_KerberosTkeyProtocol(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_kerberos_tkey_protocol"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerKerberosTkeyProtocol(name, "TCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_tkey_protocol", "TCP"),
				),
			},
			// Update and Read
			{
				Config: testAccServerKerberosTkeyProtocol(name, "UDP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "kerberos_tkey_protocol", "UDP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_name"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
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

func TestAccServerResource_ServerPrincipal(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_server_principal"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerServerPrincipal(name, "TEST_SERVER_PRINCIPAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_principal", "TEST_SERVER_PRINCIPAL"),
				),
			},
			// Update and Read
			{
				Config: testAccServerServerPrincipal(name, "TEST_SERVER_PRINCIPAL_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "server_principal", "TEST_SERVER_PRINCIPAL_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_tags"
	var v ipam.IpamsvcServer

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerTags("NAME_REPLACE_ME", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerTags("NAME_REPLACE_ME", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServerResource_VendorSpecificOptionOptionSpace(t *testing.T) {
	var resourceName = "bloxone_dhcp_server.test_vendor_specific_option_option_space"
	var v ipam.IpamsvcServer
	var name = acctest.RandomNameWithPrefix("dhcp-server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServerVendorSpecificOptionOptionSpace(name, "VENDOR_SPECIFIC_OPTION_OPTION_SPACE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_specific_option_option_space", "VENDOR_SPECIFIC_OPTION_OPTION_SPACE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccServerVendorSpecificOptionOptionSpace(name, "VENDOR_SPECIFIC_OPTION_OPTION_SPACE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "vendor_specific_option_option_space", "VENDOR_SPECIFIC_OPTION_OPTION_SPACE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckServerExists(ctx context.Context, resourceName string, v *ipam.IpamsvcServer) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
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

func testAccCheckServerDestroy(ctx context.Context, v *ipam.IpamsvcServer) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
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

func testAccCheckServerDisappears(ctx context.Context, v *ipam.IpamsvcServer) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
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
resource "bloxone_dhcp_server" "test" {
    name = %q
}
`, name)
}

func testAccServerClientPrincipal(name string, clientPrincipal string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_client_principal" {
    name = %q
    client_principal = %q
}
`, name, clientPrincipal)
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

func testAccServerDdnsConflictResolutionMode(name string, ddnsConflictResolutionMode string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_conflict_resolution_mode" {
    name = %q
    ddns_conflict_resolution_mode = %q
}
`, name, ddnsConflictResolutionMode)
}

func testAccServerDdnsDomain(name string, ddnsDomain string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_domain" {
    name = %q
    ddns_domain = %q
}
`, name, ddnsDomain)
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

func testAccServerDdnsUseConflictResolution(name string, ddnsUseConflictResolution string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_ddns_use_conflict_resolution" {
    name = %q
    ddns_use_conflict_resolution = %q
}
`, name, ddnsUseConflictResolution)
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

func testAccServerDhcpConfig(name string, dhcpConfig string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_config" {
    name = %q
    dhcp_config = %q
}
`, name, dhcpConfig)
}

func testAccServerDhcpOptions(name string, dhcpOptions string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options" {
    name = %q
    dhcp_options = %q
}
`, name, dhcpOptions)
}

func testAccServerDhcpOptionsV6(name string, dhcpOptionsV6 string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_dhcp_options_v6" {
    name = %q
    dhcp_options_v6 = %q
}
`, name, dhcpOptionsV6)
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

func testAccServerInheritanceSources(name string, inheritanceSources string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_inheritance_sources" {
    name = %q
    inheritance_sources = %q
}
`, name, inheritanceSources)
}

func testAccServerKerberosKdc(name string, kerberosKdc string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_kdc" {
    name = %q
    kerberos_kdc = %q
}
`, name, kerberosKdc)
}

func testAccServerKerberosKeys(name string, kerberosKeys string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_keys" {
    name = %q
    kerberos_keys = %q
}
`, name, kerberosKeys)
}

func testAccServerKerberosRekeyInterval(name string, kerberosRekeyInterval string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_rekey_interval" {
    name = %q
    kerberos_rekey_interval = %q
}
`, name, kerberosRekeyInterval)
}

func testAccServerKerberosRetryInterval(name string, kerberosRetryInterval string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_retry_interval" {
    name = %q
    kerberos_retry_interval = %q
}
`, name, kerberosRetryInterval)
}

func testAccServerKerberosTkeyLifetime(name string, kerberosTkeyLifetime string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_tkey_lifetime" {
    name = %q
    kerberos_tkey_lifetime = %q
}
`, name, kerberosTkeyLifetime)
}

func testAccServerKerberosTkeyProtocol(name string, kerberosTkeyProtocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_kerberos_tkey_protocol" {
    name = %q
    kerberos_tkey_protocol = %q
}
`, name, kerberosTkeyProtocol)
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

func testAccServerTags(name string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}

func testAccServerVendorSpecificOptionOptionSpace(name string, vendorSpecificOptionOptionSpace string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_server" "test_vendor_specific_option_option_space" {
    name = %q
    vendor_specific_option_option_space = %q
}
`, name, vendorSpecificOptionOptionSpace)
}
