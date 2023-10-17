package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"os"
	"regexp"
	"testing"
)

func TestAccResourceSubnet_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetBasicTestStep(),
			{
				ResourceName:            "b1ddi_subnet.tf_acc_test_subnet",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceSubnetBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
			testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "address", "192.168.1.0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.asm_threshold", "90"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable_notification", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.forecast_period", "14"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_factor", "20"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_type", "percent"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.history", "30"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_total", "10"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_unused", "10"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_scope_flag", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "cidr", "24"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "comment", "This Subnet is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "created_at"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_client_update", "client"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_domain", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generate_name", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generated_prefix", "myhost"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_send_updates", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_update_on_renew", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_use_conflict_resolution", "true"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.allow_unknown", "true"),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.filters.#"),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.ignore_list.#"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.lease_time", "3600"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_host", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_free", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_total", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_used", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_filename", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_address", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_name", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_char", "-"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_regex", "[^a-zA-Z0-9.-]"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_parent", ""),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "name", "tf_acc_test_subnet"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "protocol", "ip4"),
			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "space"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "tags.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "updated_at"),
		),
	}
}

func TestAccResourceSubnet_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetFullConfigTestStep(t),
			{
				ResourceName:            "b1ddi_subnet.tf_acc_test_subnet",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceSubnetFullConfigTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					data "b1ddi_dhcp_hosts" "dhcp_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						address = "192.168.1.0"
						asm_config {
							asm_threshold = 80
							enable = false
							enable_notification = false
							forecast_period = 9
							growth_type = "count"
							history = 50
							min_total = 20
							min_unused = 20
						}
						cidr = 24
						comment = "This Subnet is created by terraform provider acceptance test"
						ddns_client_update = "ignore"
						ddns_domain = "domain"
						ddns_generate_name = true
						ddns_generated_prefix = "tf_acc_host"
						ddns_send_updates = false
						ddns_update_on_renew = true
						ddns_use_conflict_resolution = false
						dhcp_config {
							allow_unknown = false
							#filters = ["filter1"]
							lease_time = 1800
						}
						dhcp_host = data.b1ddi_dhcp_hosts.dhcp_host.results.0.id
						
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						header_option_filename = "Acc Test Header"
						header_option_server_address = "192.168.1.10"
						header_option_server_name = "Header Option Server Name"
						hostname_rewrite_char = " "
						hostname_rewrite_enabled = true
						hostname_rewrite_regex = "[aaa bbb]"
						#inheritance_sources {}
						name = "tf_acc_test_subnet"
						space = b1ddi_ip_space.tf_acc_test_space.id
						tags = {
							TestType = "Acceptance"
						}
						#threshold {}
					}`, testAccReadDhcpHost(t)),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
			testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "address", "192.168.1.0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.asm_threshold", "80"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable_notification", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.forecast_period", "9"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_factor", "20"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_type", "count"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.history", "50"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_total", "20"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_unused", "20"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_scope_flag", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "cidr", "24"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "comment", "This Subnet is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "created_at"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_client_update", "ignore"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_domain", "domain"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generate_name", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generated_prefix", "tf_acc_host"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_send_updates", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_update_on_renew", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_use_conflict_resolution", "false"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.allow_unknown", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.filters.#"),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.ignore_list.#"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.lease_time", "1800"),

			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "dhcp_host"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.0.option_value", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.0.type", "option"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_free", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_total", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_used", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_filename", "Acc Test Header"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_address", "192.168.1.10"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_name", "Header Option Server Name"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_char", " "),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_enabled", "true"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_regex", "[aaa bbb]"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_parent", ""),
			resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "name", "tf_acc_test_subnet"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "protocol", "ip4"),
			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "space"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "tags.%", "1"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "tags.TestType", "Acceptance"),

			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "updated_at"),
		),
	}
}

func TestAccResourceSubnet_UpdateAddressExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetBasicTestStep(),
			{
				Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.15.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}`,
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "address", "192.168.1.0"),
				),
			},
			{
				ResourceName:            "b1ddi_subnet.tf_acc_test_subnet",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func TestAccResourceSubnet_UpdateSpaceExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetBasicTestStep(),
			// Create second IP Space, and recreate the subnet in this space
			{
				// ToDo Add graceful deletion process
				Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_ip_space" "tf_acc_new_test_space" {
  						name = "tf_acc_new_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_new_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
						depends_on = [b1ddi_ip_space.tf_acc_test_space]
					}`,
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
				),
			},
			{
				ResourceName:            "b1ddi_subnet.tf_acc_test_subnet",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func TestAccResourceSubnet_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceSubnetBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					data "b1ddi_dhcp_hosts" "dhcp_host" {
						filters = {
							"name" = "%s"
						}
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						address = "192.168.1.0"
						asm_config {
							asm_threshold = 80
							enable = false
							enable_notification = false
							forecast_period = 9
							growth_type = "count"
							history = 50
							min_total = 20
							min_unused = 20
						}
						cidr = 25
						comment = "This Subnet is updated by terraform provider acceptance test"
						ddns_client_update = "ignore"
						ddns_domain = "domain"
						ddns_generate_name = true
						ddns_generated_prefix = "tf_acc_host"
						ddns_send_updates = false
						ddns_update_on_renew = true
						ddns_use_conflict_resolution = false
						dhcp_config {
							allow_unknown = false
							#filters = ["filter1"]
							lease_time = 1800
						}
						dhcp_host = data.b1ddi_dhcp_hosts.dhcp_host.results.0.id
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						header_option_filename = "Acc Test Header Updated"
						header_option_server_address = "192.168.1.10"
						header_option_server_name = "Header Option Server Name"
						hostname_rewrite_char = " "
						hostname_rewrite_enabled = true
						hostname_rewrite_regex = "[aaa bbb]"
						#inheritance_sources {}
						name = "tf_acc_test_subnet"
						space = b1ddi_ip_space.tf_acc_test_space.id
						tags = {
							TestType = "Acceptance"
						}
						#threshold {}
					}`, testAccReadDhcpHost(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "address", "192.168.1.0"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.asm_threshold", "80"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.enable_notification", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.forecast_period", "9"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_factor", "20"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.growth_type", "count"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.history", "50"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_total", "20"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.min_unused", "20"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "asm_scope_flag", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "cidr", "25"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "comment", "This Subnet is updated by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "created_at"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_client_update", "ignore"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_domain", "domain"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generate_name", "true"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_generated_prefix", "tf_acc_host"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_send_updates", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_update_on_renew", "true"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "ddns_use_conflict_resolution", "false"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.allow_unknown", "false"),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.filters.#"),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.ignore_list.#"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_config.0.lease_time", "1800"),
					// ToDo Add dhcp_host parameter
					//resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_host", "dhcp_host"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.0.option_value", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_options.0.type", "option"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_free", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_total", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_used", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "dhcp_utilization.0.dhcp_utilization", "0"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_filename", "Acc Test Header Updated"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_address", "192.168.1.10"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "header_option_server_name", "Header Option Server Name"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_char", " "),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_enabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "hostname_rewrite_regex", "[aaa bbb]"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_assigned_hosts.%", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_parent", ""),
					resource.TestCheckNoResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "name", "tf_acc_test_subnet"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "parent", ""),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "protocol", "ip4"),
					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "space"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "tags.TestType", "Acceptance"),

					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.high", "0"),
					resource.TestCheckResourceAttr("b1ddi_subnet.tf_acc_test_subnet", "threshold.0.low", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_subnet.tf_acc_test_subnet", "updated_at"),
				),
			},
			{
				ResourceName:            "b1ddi_subnet.tf_acc_test_subnet",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func testCheckSubnetExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.Subnet.SubnetRead(
			&subnet.SubnetReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

// Checks if the specified Subnet resides in the specified IP Space
func testCheckSubnetInSpace(subnetPath, spacePath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		subnetResource, found := s.RootModule().Resources[subnetPath]
		if !found {
			return fmt.Errorf("not found %s", subnetPath)
		}
		if subnetResource.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", subnetPath)
		}
		space, found := s.RootModule().Resources[spacePath]
		if !found {
			return fmt.Errorf("not found %s", spacePath)
		}
		if space.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", spacePath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.Subnet.SubnetRead(
			&subnet.SubnetReadParams{ID: subnetResource.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.ID != subnetResource.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				subnetResource.Primary.ID)
		}

		if *resp.Payload.Result.Space != space.Primary.ID {
			return fmt.Errorf(
				"'space' does not match: \n got: '%s', \nexpected: '%s'",
				*resp.Payload.Result.Space,
				space.Primary.ID)
		}

		return nil
	}
}

// Read Dhcp Host name from the env. If env is not specified, skip the test.
func testAccReadDhcpHost(t *testing.T) string {
	internalSecondary := os.Getenv("DHCP_HOST")

	if internalSecondary == "" {
		t.Skipf("No DHCP_HOST env is set for the %s acceptance test", t.Name())
	}

	return internalSecondary
}
