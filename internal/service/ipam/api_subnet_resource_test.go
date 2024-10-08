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
// - dhcp_host
// - dhcp_options
// - external_keys
// - federated_realms
// - config_profiles

func TestAccSubnetResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetBasicConfig(spaceName, "10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
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

func TestAccSubnetResource_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_subnet.test"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubnetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetBasicConfig(spaceName, "10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					testAccCheckSubnetDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSubnetResource_Address(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test"
	var v1 ipam.Subnet
	var v2 ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetBasicConfig(spaceName, "10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Update and Read
			// Update should destroy previous subnet and create new subnet
			{
				Config: testAccSubnetBasicConfig(spaceName, "11.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetDestroy(context.Background(), &v1),
					testAccCheckSubnetExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", "11.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetBasicConfig(spaceName, "10.0.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetBasicConfig(spaceName, "10.0.0.0", 26),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "26"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_Space(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test"
	var v1 ipam.Subnet
	var v2 ipam.Subnet
	spaceName1 := acctest.RandomNameWithPrefix("ip-space")
	spaceName2 := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetSpace(spaceName1, spaceName2, "bloxone_ipam_ip_space.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetSpace(spaceName1, spaceName2, "bloxone_ipam_ip_space.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetDestroy(context.Background(), &v1),
					testAccCheckSubnetExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "space", "bloxone_ipam_ip_space.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_AsmConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_asm_config"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetAsmConfig(spaceName, "10.0.0.0", 24, 70, true, true, 12, 40, "count", 40, 30, 30, "2020-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "asm_config.asm_threshold", "70"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.enable_notification", "true"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.forecast_period", "12"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_factor", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.growth_type", "count"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.history", "40"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_total", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.min_unused", "30"),
					resource.TestCheckResourceAttr(resourceName, "asm_config.reenable_date", "2020-01-10T10:11:22Z"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetAsmConfig(spaceName, "10.0.0.0", 24, 80, false, false, 10, 50, "percent", 50, 10, 10, "2021-01-10T10:11:22Z"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
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

func TestAccSubnetResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_comment"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetComment(spaceName, "10.0.0.0", 24, "some comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "some comment"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetComment(spaceName, "10.0.0.0", 24, "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsClientUpdate(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_client_update"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsClientUpdate(spaceName, "10.0.0.0", 24, "server"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "server"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsClientUpdate(spaceName, "10.0.0.0", 24, "over_client_update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_client_update", "over_client_update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsConflictResolutionMode(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_conflict_resolution_mode"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsConflictResolutionMode(spaceName, "10.0.0.0", 24, false, "check_exists_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_exists_with_dhcid"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsConflictResolutionMode(spaceName, "10.0.0.0", 24, true, "check_with_dhcid"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_use_conflict_resolution", "true"),
					resource.TestCheckResourceAttr(resourceName, "ddns_conflict_resolution_mode", "check_with_dhcid"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsDomain(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_domain"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsDomain(spaceName, "10.0.0.0", 24, "abc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "abc"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsDomain(spaceName, "10.0.0.0", 24, "xyz"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_domain", "xyz"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsGenerateName(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_generate_name"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsGenerateName(spaceName, "10.0.0.0", 24, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsGenerateName(spaceName, "10.0.0.0", 24, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generate_name", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsGeneratedPrefix(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_generated_prefix"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsGeneratedPrefix(spaceName, "10.0.0.0", 24, "host-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-prefix"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsGeneratedPrefix(spaceName, "10.0.0.0", 24, "host-another-prefix"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_generated_prefix", "host-another-prefix"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DhcpOptions(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_dhcp_options"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDhcpOptionsOption(spaceName, "10.0.0.0", 24, "option_group_test", "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDhcpOptionsGroup(spaceName, "10.0.0.0", 24, "option_group_test", "group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsSendUpdates(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_send_updates"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsSendUpdates(spaceName, "10.0.0.0", 24, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsSendUpdates(spaceName, "10.0.0.0", 24, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_send_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsTtlPercent(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_ttl_percent"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsTtlPercent(spaceName, "10.0.0.0", 24, "20"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "20"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsTtlPercent(spaceName, "10.0.0.0", 24, "40"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_ttl_percent", "40"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DdnsUpdateOnRenew(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_ddns_update_on_renew"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDdnsUpdateOnRenew(spaceName, "10.0.0.0", 24, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDdnsUpdateOnRenew(spaceName, "10.0.0.0", 24, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_update_on_renew", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_DhcpConfig(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_dhcp_config"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDhcpConfig(spaceName, "10.0.0.0", 24, true, true, true, 50, 60),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.allow_unknown_v6", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.ignore_client_uid", "true"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time", "50"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_config.lease_time_v6", "60")),
			},
			// Update and Read
			{
				Config: testAccSubnetDhcpConfig(spaceName, "10.0.0.0", 24, false, false, false, 55, 65),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
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

func TestAccSubnetResource_DisableDhcp(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_disable_dhcp"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetDisableDhcp(spaceName, "10.0.0.0", 24, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetDisableDhcp(spaceName, "10.0.0.0", 24, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_dhcp", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HeaderOptionFilename(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_header_option_filename"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHeaderOptionFilename(spaceName, "10.0.0.0", 24, "HEADER_OPTION_FILENAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILENAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHeaderOptionFilename(spaceName, "10.0.0.0", 24, "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_filename", "HEADER_OPTION_FILENAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HeaderOptionServerAddress(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_header_option_server_address"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHeaderOptionServerAddress(spaceName, "10.0.0.0", 24, "12.0.0.4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.4"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHeaderOptionServerAddress(spaceName, "10.0.0.0", 24, "12.0.0.5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_address", "12.0.0.5"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HeaderOptionServerName(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_header_option_server_name"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHeaderOptionServerName(spaceName, "10.0.0.0", 24, "HEADER_OPTION_SERVER_NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHeaderOptionServerName(spaceName, "10.0.0.0", 24, "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "header_option_server_name", "HEADER_OPTION_SERVER_NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HostnameRewriteChar(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_hostname_rewrite_char"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHostnameRewriteChar(spaceName, "10.0.0.0", 24, "+"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "+"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHostnameRewriteChar(spaceName, "10.0.0.0", 24, "/"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_char", "/"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HostnameRewriteEnabled(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_hostname_rewrite_enabled"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHostnameRewriteEnabled(spaceName, "10.0.0.0", 24, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHostnameRewriteEnabled(spaceName, "10.0.0.0", 24, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_HostnameRewriteRegex(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_hostname_rewrite_regex"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetHostnameRewriteRegex(spaceName, "10.0.0.0", 24, "[^a-z]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^a-z]"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetHostnameRewriteRegex(spaceName, "10.0.0.0", 24, "[^0-9]"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hostname_rewrite_regex", "[^0-9]"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_InheritanceSources(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_inheritance_sources"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetInheritanceSources(spaceName, "10.0.0.0", 24, "inherit"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_enabled.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_hostname_block.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "inherit"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_block.action", "inherit"),
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
				),
			},
			// Update and Read
			{
				Config: testAccSubnetInheritanceSources(spaceName, "10.0.0.0", 24, "override"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_enable_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_growth_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.asm_threshold.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.forecast_period.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.history.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_total.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.asm_config.min_unused.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_client_update.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_conflict_resolution_mode.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_hostname_block.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_ttl_percent.action", "override"),
					resource.TestCheckResourceAttr(resourceName, "inheritance_sources.ddns_update_block.action", "override"),
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
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_MultipleFederatedRealms(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_federated_realms"
	var v ipam.Subnet
	var ipSpaceName = acctest.RandomNameWithPrefix("ip-space")
	var address = "192.168.0.0"
	var cidr = "16"
	var realmName1 = acctest.RandomNameWithPrefix("realm1")
	var realmName2 = acctest.RandomNameWithPrefix("realm2")
	var realmName3 = acctest.RandomNameWithPrefix("realm3")
	var realmName4 = acctest.RandomNameWithPrefix("realm4")
	var realmName5 = acctest.RandomNameWithPrefix("realm5")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSubnetDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			// Step 1: Create subnet with multiple federated realms and verify
			{
				Config: testAccSubnetMultipleFederatedRealms(ipSpaceName, address, cidr, realmName1, realmName2, realmName3, realmName4, realmName5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms.#", "5"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.0", "bloxone_federation_federated_realm."+realmName1, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.1", "bloxone_federation_federated_realm."+realmName2, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.2", "bloxone_federation_federated_realm."+realmName3, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.3", "bloxone_federation_federated_realm."+realmName4, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.4", "bloxone_federation_federated_realm."+realmName5, "id"),
				),
			},
			// Step 2: Update subnet with federated realms with different order and verify
			{
				Config: testAccSubnetMultipleFederatedRealms(ipSpaceName, address, cidr, realmName3, realmName5, realmName1, realmName2, realmName4),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realms.#", "5"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.0", "bloxone_federation_federated_realm."+realmName3, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.1", "bloxone_federation_federated_realm."+realmName5, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.2", "bloxone_federation_federated_realm."+realmName1, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.3", "bloxone_federation_federated_realm."+realmName2, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realms.4", "bloxone_federation_federated_realm."+realmName4, "id"),
				),
			},
		},
	})
}

func TestAccSubnetResource_Name(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_name"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetName(spaceName, "10.0.0.0", 24, "subnet_name"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "subnet_name"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetName(spaceName, "10.0.0.0", 24, "subnet_name_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "subnet_name_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_RenewTimeAndRebindTime(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_renew_time"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetRebindTimeAndRenewTime(spaceName, "10.0.0.0", 24, "60", "50"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rebind_time", "60"),
					resource.TestCheckResourceAttr(resourceName, "renew_time", "50"),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetRebindTimeAndRenewTime(spaceName, "10.0.0.0", 24, "90", "80"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rebind_time", "90"),
					resource.TestCheckResourceAttr(resourceName, "renew_time", "80"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSubnetResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_tags"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSubnetTags(spaceName, "10.0.0.0", 24, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccSubnetTags(spaceName, "10.0.0.0", 24, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName, &v),
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

func TestAccSubnetResource_NextAvailableId(t *testing.T) {
	var resourceName = "bloxone_ipam_subnet.test_next_available_id"
	var v ipam.Subnet
	spaceName := acctest.RandomNameWithPrefix("ip-space")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetNextAvailableId(spaceName, "bloxone_ipam_address_block.one", 24, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName+".0", &v),
					resource.TestCheckResourceAttrPair(resourceName+".0", "next_available_id", "bloxone_ipam_address_block.one", "id"),
					resource.TestCheckResourceAttr(resourceName+".0", "address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName+".0", "cidr", "24"),
				),
			},
			{
				Config: testAccSubnetNextAvailableId(spaceName, "bloxone_ipam_address_block.two", 24, 1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSubnetExists(context.Background(), resourceName+".0", &v),
					resource.TestCheckResourceAttrPair(resourceName+".0", "next_available_id", "bloxone_ipam_address_block.two", "id"),
					resource.TestCheckResourceAttr(resourceName+".0", "address", "11.0.0.0"),
					resource.TestCheckResourceAttr(resourceName+".0", "cidr", "24"),
				),
			},
			// Test when count > 1
			{
				Config: testAccSubnetNextAvailableId(spaceName, "bloxone_ipam_address_block.two", 24, 5),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName+".0", "next_available_id", "bloxone_ipam_address_block.two", "id"),
					resource.TestCheckResourceAttrPair(resourceName+".1", "next_available_id", "bloxone_ipam_address_block.two", "id"),
					resource.TestCheckResourceAttrPair(resourceName+".2", "next_available_id", "bloxone_ipam_address_block.two", "id"),
					resource.TestCheckResourceAttrPair(resourceName+".3", "next_available_id", "bloxone_ipam_address_block.two", "id"),
					resource.TestCheckResourceAttrPair(resourceName+".4", "next_available_id", "bloxone_ipam_address_block.two", "id"),
				),
			},
		},
	})
}

func testAccCheckSubnetExists(ctx context.Context, resourceName string, v *ipam.Subnet) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			SubnetAPI.
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

func testAccCheckSubnetDestroy(ctx context.Context, v *ipam.Subnet) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			SubnetAPI.
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

func testAccCheckSubnetDisappears(ctx context.Context, v *ipam.Subnet) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			SubnetAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithIPSpace(name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test" {
    name = %q
}
`, name)
}

func testAccBaseWithTwoIPSpace(name1, name2 string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "one" {
    name = %q
}
resource "bloxone_ipam_ip_space" "two" {
    name = %q
}`, name1, name2)
}

func testAccSubnetBasicConfig(spaceName, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetSpace(spaceName1, spaceName2, space string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test" {
    address = "10.0.0.0"
    cidr = 24
    space = %s.id
}
`, space)
	return strings.Join([]string{testAccBaseWithTwoIPSpace(spaceName1, spaceName2), config}, "")
}

func testAccSubnetAsmConfig(spaceName, address string, cidr int, asmThreshold int, enable, enableNotification bool, forecastPeriod, growthFactor int, growthType string, history, minTotal, minUnused int, reenableDate string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_asm_config" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
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
`, address, cidr, asmThreshold, enable, enableNotification, forecastPeriod, growthFactor, growthType, history, minTotal, minUnused, reenableDate)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetComment(spaceName, address string, cidr int, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_comment" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    comment = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsClientUpdate(spaceName, address string, cidr int, ddnsClientUpdate string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_client_update" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_client_update = %q
}
`, address, cidr, ddnsClientUpdate)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsConflictResolutionMode(spaceName, address string, cidr int, ddnsUseConflictResolution bool, ddnsConflictResolutionMode string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_conflict_resolution_mode" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
	ddns_use_conflict_resolution = %t
    ddns_conflict_resolution_mode = %q
}
`, address, cidr, ddnsUseConflictResolution, ddnsConflictResolutionMode)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDhcpOptionsOption(spaceName, address string, cidr int, name, optionItemType, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_dhcp_options" {
	address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    name = %q
    dhcp_options = [
      {
       type = %q
       option_code = bloxone_dhcp_option_code.test.id
       option_value = %q
      }
    ]
}
`, address, cidr, name, optionItemType, optValue)

	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip4"), config}, "")
}

func testAccSubnetDhcpOptionsGroup(spaceName, address string, cidr int, name, optionItemType string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_dhcp_options" {
	address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    name = %q
    dhcp_options = [
      {
       type = %q
       group = bloxone_dhcp_option_group.test.id
      }
    ]
}
`, address, cidr, name, optionItemType)

	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), testAccBaseWithOptionSpaceAndCode("og-"+name, "os-"+name, "ip4"), config}, "")
}

func testAccSubnetMultipleFederatedRealms(spaceName, address, cidr, realmName1, realmName2, realmName3, realmName4, realmName5 string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_federated_realms" {
    address = %q
    cidr = %q
    space = bloxone_ipam_ip_space.test.id
    federated_realms = [
		bloxone_federation_federated_realm.%s.id,
		bloxone_federation_federated_realm.%s.id,
		bloxone_federation_federated_realm.%s.id,
		bloxone_federation_federated_realm.%s.id,
		bloxone_federation_federated_realm.%s.id
	]
}
`, address, cidr, realmName1, realmName2, realmName3, realmName4, realmName5)
	return strings.Join([]string{
		testAccBaseWithIPSpace(spaceName),
		testAccBaseWithFederatedRealm(realmName1, realmName1),
		testAccBaseWithFederatedRealm(realmName2, realmName2),
		testAccBaseWithFederatedRealm(realmName3, realmName3),
		testAccBaseWithFederatedRealm(realmName4, realmName4),
		testAccBaseWithFederatedRealm(realmName5, realmName5),
		config,
	}, "")
}

func testAccBaseWithOptionSpaceAndCode(optionGroup, optionSpace, protocol string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test" {
    name = %q
    protocol = %q
}

resource "bloxone_dhcp_option_code" "test" {
    code = "234"
    name = "test_dhcp_option_code"
    option_space = bloxone_dhcp_option_space.test.id
    type = "boolean"
}
`, optionGroup, protocol)
	return strings.Join([]string{testAccBaseWithOptionSpace(optionSpace, protocol), config}, "")
}

func testAccSubnetDdnsDomain(spaceName, address string, cidr int, ddnsDomain string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_domain" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_domain = %q
}
`, address, cidr, ddnsDomain)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsGenerateName(spaceName, address string, cidr int, ddnsGenerateName bool) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_generate_name" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_generate_name = %t
}
`, address, cidr, ddnsGenerateName)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsGeneratedPrefix(spaceName, address string, cidr int, ddnsGeneratedPrefix string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_generated_prefix" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_generated_prefix = %q
}
`, address, cidr, ddnsGeneratedPrefix)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsSendUpdates(spaceName, address string, cidr int, ddnsSendUpdates bool) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_send_updates" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_send_updates = %t
}
`, address, cidr, ddnsSendUpdates)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsTtlPercent(spaceName, address string, cidr int, ddnsTtlPercent string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_ttl_percent" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_ttl_percent = %q
}
`, address, cidr, ddnsTtlPercent)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDdnsUpdateOnRenew(spaceName, address string, cidr int, ddnsUpdateOnRenew bool) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_ddns_update_on_renew" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    ddns_update_on_renew = %t
}
`, address, cidr, ddnsUpdateOnRenew)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDhcpConfig(spaceName, address string, cidr int, allowUnknown, allowUnknownV6, ignoreClientUid bool,
	leaseTime, leaseTimeV6 int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_dhcp_config" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    dhcp_config = {
		allow_unknown = %t
		allow_unknown_v6 = %t
		ignore_client_uid = %t
		lease_time = %d
		lease_time_v6 = %d
	}
}
`, address, cidr, allowUnknown, allowUnknownV6, ignoreClientUid, leaseTime, leaseTimeV6)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetDisableDhcp(spaceName, address string, cidr int, disableDhcp bool) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_disable_dhcp" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    disable_dhcp = %t
}
`, address, cidr, disableDhcp)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHeaderOptionFilename(spaceName, address string, cidr int, headerOptionFilename string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_header_option_filename" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    header_option_filename = %q
}
`, address, cidr, headerOptionFilename)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHeaderOptionServerAddress(spaceName, address string, cidr int, headerOptionServerAddress string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_header_option_server_address" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    header_option_server_address = %q
}
`, address, cidr, headerOptionServerAddress)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHeaderOptionServerName(spaceName, address string, cidr int, headerOptionServerName string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_header_option_server_name" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    header_option_server_name = %q
}
`, address, cidr, headerOptionServerName)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHostnameRewriteChar(spaceName, address string, cidr int, hostnameRewriteChar string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_hostname_rewrite_char" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_char = %q
}
`, address, cidr, hostnameRewriteChar)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHostnameRewriteEnabled(spaceName, address string, cidr int, hostnameRewriteEnabled bool) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_hostname_rewrite_enabled" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_enabled = %t
}
`, address, cidr, hostnameRewriteEnabled)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetHostnameRewriteRegex(spaceName, address string, cidr int, hostnameRewriteRegex string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_hostname_rewrite_regex" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    hostname_rewrite_regex = %q
}
`, address, cidr, hostnameRewriteRegex)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetInheritanceSources(spaceName, address string, cidr int, action string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_inheritance_sources" {
    address = %[1]q
    cidr = %[2]d
    space = bloxone_ipam_ip_space.test.id
	inheritance_sources = {
		asm_config = {
			action = %[3]q
			asm_enable_block = {
				action = %[3]q
			}
			asm_growth_block = {
				action = %[3]q
			}
			asm_threshold = {
				action = %[3]q
			}
			forecast_period = {
				action = %[3]q
			}
			history = {
				action = %[3]q
			}
			min_total = {
				action = %[3]q
			}
			min_unused = {
				action = %[3]q
			}
		}
		dhcp_config = {
			allow_unknown = {
				action = %[3]q
			}
			allow_unknown_v6 = {
				action = %[3]q
			}
			filters	= {
				action = %[3]q
			}
			filters_v6	= {
				action = %[3]q
			}
			ignore_client_uid = {
				action = %[3]q
			}
			ignore_list	= {
				action = %[3]q
			}
			lease_time = {
				action = %[3]q
			}
			lease_time_v6 = {
				action = %[3]q
			}
		}
		ddns_client_update = {
			action = %[3]q
		}
		ddns_conflict_resolution_mode = {
			action = %[3]q
		}
		ddns_enabled = {
			action = "inherit"
		}
		ddns_hostname_block = {
			action = %[3]q
		}
		ddns_ttl_percent = {
			action = %[3]q
		}
		ddns_update_block = {
			action = %[3]q
		}
		ddns_update_on_renew = {
			action = %[3]q
		}
		//dhcp_option
		ddns_use_conflict_resolution = {
			action = %[3]q
		}
		header_option_filename = {
			action = %[3]q
		}
		header_option_server_address = {
			action = %[3]q
		}
		header_option_server_name = {
			action = %[3]q
		}
		hostname_rewrite_block = {
			action = %[3]q
		}
	}
}
`, address, cidr, action)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetName(spaceName, address string, cidr int, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_name" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    name = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetRebindTimeAndRenewTime(spaceName, address string, cidr int, rebindTime, renewTime string) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_renew_time" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    rebind_time = %q
    renew_time = %q
}
`, address, cidr, rebindTime, renewTime)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetTags(spaceName, address string, cidr int, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_ipam_subnet" "test_tags" {
    address = %q
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
    tags = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}

func testAccSubnetNextAvailableId(spaceName string, addressBlockResourceName string, cidr int, count int) string {
	config := fmt.Sprintf(`
resource "bloxone_ipam_address_block" "one" {
	space = bloxone_ipam_ip_space.test.id
	address = "10.0.0.0"
	cidr = 16
}

resource "bloxone_ipam_address_block" "two" {
	space = bloxone_ipam_ip_space.test.id
	address = "11.0.0.0"
	cidr = 16
}

resource "bloxone_ipam_subnet" "test_next_available_id" {
	count = %d
    next_available_id = %s.id
    cidr = %d
    space = bloxone_ipam_ip_space.test.id
}
`, count, addressBlockResourceName, cidr)
	return strings.Join([]string{testAccBaseWithIPSpace(spaceName), config}, "")
}
