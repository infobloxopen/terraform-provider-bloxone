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
// - dhcp_config.filters
// - dhcp_config.filters_v6
// - dhcp_config.ignore_items
// - vendor_specific_option

func TestAccIpSpaceResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceBasicConfig("test", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
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

func TestAccIpSpaceResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_ip_space.test"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpSpaceBasicConfig("test", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					testAccCheckIpSpaceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpSpaceResource_AsmConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_asm_config"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceAsmConfig(name, 70, true, true, 12, 40, "count", 40, 30, 30, "2020-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "70"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "12"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "count"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.history", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2020-01-10T10:11:22Z")),
			},
			// Update and Read
			{
				Config: testAccIpSpaceAsmConfig(name, 80, false, false, 10, 50, "percent", 50, 10, 10, "2021-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "80"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "false"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "50"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "percent"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.history", "50"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "10"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2021-01-10T10:11:22Z"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_comment"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceComment(name, "some comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceComment(name, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Update and Read  (unset)
			{
				Config: testAccIpSpaceComment(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Update and Read  (unset null)
			{
				Config: testAccIpSpaceBasicConfig("test_comment", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsClientUpdate(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_client_update"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsClientUpdate(name, "server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "server"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsClientUpdate(name, "over_client_update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "over_client_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsConflictResolutionMode(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_conflict_resolution_mode"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsConflictResolutionMode(name, false, "check_exists_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_exists_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsConflictResolutionMode(name, true, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsDomain(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_domain"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsDomain(name, "abc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "abc"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsDomain(name, "xyz"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "xyz"),
				),
			},
			// Update and Read  (unset empty string)
			{
				Config: testAccIpSpaceDdnsDomain(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", ""),
				),
			},
			// Update and Read  (unset null)
			{
				Config: testAccIpSpaceBasicConfig("test_ddns_domain", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsGenerateName(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_generate_name"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsGenerateName(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsGenerateName(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsGeneratedPrefix(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_generated_prefix"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsGeneratedPrefix(name, "host-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-prefix"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsGeneratedPrefix(name, "host-another-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-another-prefix"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DhcpOptions(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_dhcp_options"
	var v1 ipam.IpamsvcIPSpace
	name := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDhcpOptionsOption(name, "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDhcpOptionsGroup(name, "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DhcpOptionsV6(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_dhcp_options_v6"
	var v1 ipam.IpamsvcIPSpace
	name := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDhcpOptionsOptionV6(name, "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options_v6.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDhcpOptionsGroupV6(name, "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options_v6.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsSendUpdates(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_send_updates"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsSendUpdates(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsSendUpdates(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsTtlPercent(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_ttl_percent"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsTtlPercent(name, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsTtlPercent(name, "40"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "40"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DdnsUpdateOnRenew(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_ddns_update_on_renew"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDdnsUpdateOnRenew(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDdnsUpdateOnRenew(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_DhcpConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_dhcp_config"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceDhcpConfig(name, 1000, 2000, true, true, true, 50, 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time", "1000"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time_v6", "2000"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "50"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "60"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceDhcpConfig(name, 1500, 2500, false, false, false, 55, 65),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time", "1500"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.abandoned_reclaim_time_v6", "2500"),
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

func TestAccIpSpaceResource_HeaderOptionFilename(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_filename"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionFilename(name, "HEADER_OPTION_FILEip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILEip_space_name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionFilename(name, "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read  (unset)
			{
				Config: testAccIpSpaceHeaderOptionFilename(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", ""),
				),
			},
			// Update and Read  (unset null)
			{
				Config: testAccIpSpaceBasicConfig("test_header_option_filename", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_server_address"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionServerAddress(name, "10.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionServerAddress(name, "12.0.0.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.0"),
				),
			},
			// Update and Read  (unset)
			{
				Config: testAccIpSpaceHeaderOptionServerAddress(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", ""),
				),
			},
			// Update and Read  (unset null)
			{
				Config: testAccIpSpaceBasicConfig("test_header_option_server_address", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_header_option_server_name"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHeaderOptionServerName(name, "HEADER_OPTION_SERVER_ip_space_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_ip_space_name"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHeaderOptionServerName(name, "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read  (unset)
			{
				Config: testAccIpSpaceHeaderOptionServerName(name, ""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", ""),
				),
			},
			// Update and Read  (unset null)
			{
				Config: testAccIpSpaceBasicConfig("test_header_option_server_name", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteChar(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_char"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteChar(name, "+"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "+"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteChar(name, "/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "/"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteEnabled(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_enabled"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteEnabled(name, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteEnabled(name, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_HostnameRewriteRegex(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_hostname_rewrite_regex"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceHostnameRewriteRegex(name, "[^a-z]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z]"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceHostnameRewriteRegex(name, "[^0-9]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^0-9]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpSpaceResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_inheritance_sources"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceInheritanceSources(name, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_filename.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_address.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.header_option_server_name.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.hostname_rewrite_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.vendor_specific_option_option_space.action", "inherit"),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceInheritanceSources(name, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_on_renew.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_use_conflict_resolution.action", "override"),
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

func TestAccIpSpaceResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_ip_space.test_tags"
	var v ipam.IpamsvcIPSpace
	var name = acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpSpaceTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccIpSpaceTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpSpaceExists(context.Background(), resourceName, &v),
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

func testAccCheckIpSpaceExists(ctx context.Context, resourceName string, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceRead(ctx, rs.Primary.ID).
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

func testAccCheckIpSpaceDestroy(ctx context.Context, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceRead(ctx, *v.Id).
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

func testAccCheckIpSpaceDisappears(ctx context.Context, v *ipam.IpamsvcIPSpace) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpSpaceAPI.
			IpSpaceDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpSpaceBasicConfig(rName, name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" %q {
    name = %q
}
`, rName, name)
}

func testAccIpSpaceAsmConfig(name string, asmThreshold int, enable, enableNotification bool, forecastPeriod, growthFactor int, growthType string, history, minTotal, minUnused int, reenableDate string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_asm_config" {
    name = %q
    asm_config = {
		asm_threshold = %d
		enable = %t
		enable_notification = %t
		forecast_period = %d
		growth_factor = %d
		growth_type = %q
		history = %d
		min_total = %d
		min_unused = %d
		reenable_date = %q
	}
}
`, name, asmThreshold, enable, enableNotification, forecastPeriod, growthFactor, growthType, history, minTotal, minUnused, reenableDate)
}

func testAccIpSpaceComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccIpSpaceDdnsClientUpdate(name, ddnsClientUpdate string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_client_update" {
    name = %q
    ddns_client_update = %q
}
`, name, ddnsClientUpdate)
}

func testAccIpSpaceDdnsConflictResolutionMode(name string, ddnsUseConflictResolution bool, ddnsConflictResolutionMode string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_conflict_resolution_mode" {
    name = %q
	ddns_use_conflict_resolution = %t
    ddns_conflict_resolution_mode = %q
}
`, name, ddnsUseConflictResolution, ddnsConflictResolutionMode)
}

func testAccIpSpaceDdnsDomain(name, ddnsDomain string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_domain" {
    name = %q
    ddns_domain = %q
}
`, name, ddnsDomain)
}

func testAccIpSpaceDdnsGenerateName(name string, ddnsGenerateName bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_generate_name" {
    name = %q
    ddns_generate_name = %t
}
`, name, ddnsGenerateName)
}

func testAccIpSpaceDdnsGeneratedPrefix(name, ddnsGeneratedPrefix string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_generated_prefix" {
    name = %q
    ddns_generated_prefix = %q
}
`, name, ddnsGeneratedPrefix)
}

func testAccIpSpaceDdnsSendUpdates(name string, ddnsSendUpdates bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_send_updates" {
    name = %q
    ddns_send_updates = %t
}
`, name, ddnsSendUpdates)
}

func testAccIpSpaceDdnsTtlPercent(name, ddnsTtlPercent string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_ttl_percent" {
    name = %q
    ddns_ttl_percent = %s
}
`, name, ddnsTtlPercent)
}

func testAccIpSpaceDdnsUpdateOnRenew(name string, ddnsUpdateOnRenew bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_ddns_update_on_renew" {
    name = %q
    ddns_update_on_renew = %t
}
`, name, ddnsUpdateOnRenew)
}

func testAccIpSpaceDhcpConfig(name string,
	abandonedReclaimTime, abandonedReclaimTimeV6 int,
	allowUnknown, allowUnknownV6, ignoreClientUid bool,
	leaseTime, leaseTimeV6 int) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_config" {
    name = %q
    dhcp_config = {
		abandoned_reclaim_time = %d
		abandoned_reclaim_time_v6 = %d
		allow_unknown = %t
		allow_unknown_v6 = %t
		ignore_client_uid = %t
		lease_time = %d
		lease_time_v6 = %d
	}
}
`, name, abandonedReclaimTime, abandonedReclaimTimeV6, allowUnknown, allowUnknownV6, ignoreClientUid, leaseTime, leaseTimeV6)
}

func testAccIpSpaceDhcpOptionsOption(name string, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_options" {
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
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip4"), config}, "")

}

func testAccIpSpaceDhcpOptionsGroup(name string, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_options" {
    name = %q
    dhcp_options = [
      {
       type = %q
       group = bloxone_dhcp_option_group.test.id
      }
    ]
}
`, name, optionItemType)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip4"), config}, "")

}

func testAccIpSpaceDhcpOptionsOptionV6(name string, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_options_v6" {
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
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip6"), config}, "")
}

func testAccIpSpaceDhcpOptionsGroupV6(name string, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_dhcp_options_v6" {
   name = %q
   dhcp_options_v6 = [
     {
      type = %q
      group = bloxone_dhcp_option_group.test.id
     }
   ]
}
`, name, optionItemType)
	return strings.Join([]string{testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip6"), config}, "")
}

func testAccIpSpaceHeaderOptionFilename(name, headerOptionFilename string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_filename" {
    name = %q
    header_option_filename = %q
}
`, name, headerOptionFilename)
}

func testAccIpSpaceHeaderOptionServerAddress(name, headerOptionServerAddress string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_server_address" {
    name = %q
    header_option_server_address = %q
}
`, name, headerOptionServerAddress)
}

func testAccIpSpaceHeaderOptionServerName(name, headerOptionServerName string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_header_option_server_name" {
    name = %q
    header_option_server_name = %q
}
`, name, headerOptionServerName)
}

func testAccIpSpaceHostnameRewriteChar(name, hostnameRewriteChar string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_char" {
    name = %q
    hostname_rewrite_char = %q
}
`, name, hostnameRewriteChar)
}

func testAccIpSpaceHostnameRewriteEnabled(name string, hostnameRewriteEnabled bool) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_enabled" {
    name = %q
    hostname_rewrite_enabled = %t
}
`, name, hostnameRewriteEnabled)
}

func testAccIpSpaceHostnameRewriteRegex(name, hostnameRewriteRegex string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_hostname_rewrite_regex" {
    name = %q
    hostname_rewrite_regex = %q
}
`, name, hostnameRewriteRegex)
}

func testAccIpSpaceInheritanceSources(name, action string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_inheritance_sources" {
	name = %[1]q
	inheritance_sources = {
		asm_config = {
			action = %[2]q
			asm_enable_block = {
				action = %[2]q
			}
			asm_growth_block = {
				action = %[2]q
			}
			asm_threshold = {
				action = %[2]q
			}
			forecast_period = {
				action = %[2]q
			}
			history = {
				action = %[2]q
			}
			min_total = {
				action = %[2]q
			}
			min_unused = {
				action = %[2]q
			}
		}
		ddns_client_update = {
			action = %[2]q
		}
		ddns_conflict_resolution_mode = {
			action = %[2]q
		}
		ddns_enabled = {
			action = "inherit"
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
		vendor_specific_option_option_space	= {
			action = %[2]q
		}
	}

}
`, name, action)
}

func testAccIpSpaceTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
