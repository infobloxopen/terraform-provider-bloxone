package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address_block"
	"regexp"
	"testing"
)

func TestAccResourceAddressBlock_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockBasicTestStep(),
			{
				ResourceName:            "b1ddi_address_block.tf_acc_test_address_block",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceAddressBlockBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
  						name = "tf_acc_test_address_block"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Address Block is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testAccAddressBlockExists("b1ddi_address_block.tf_acc_test_address_block"),
			testCheckAddressBlockInSpace("b1ddi_address_block.tf_acc_test_address_block", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "address", "192.168.1.0"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.asm_threshold", "90"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable", "true"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable_notification", "true"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.forecast_period", "14"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_factor", "20"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_type", "percent"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.history", "30"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_total", "10"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_unused", "10"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_scope_flag", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "cidr", "24"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "comment", "This Address Block is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_client_update", "client"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_domain", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generate_name", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generated_prefix", "myhost"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_send_updates", "true"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_update_on_renew", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_use_conflict_resolution", "true"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.allow_unknown", "true"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.filters.#"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.ignore_list.#"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.lease_time", "3600"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_host"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_free", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_total", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_used", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_filename", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_address", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_name", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_char", "-"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_regex", "[^a-zA-Z0-9.-]"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_parent", ""),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "name", "tf_acc_test_address_block"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "protocol", "ip4"),
			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "space"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "tags.%", "1"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "updated_at"),
		),
	}
}

func TestAccResourceAddressBlock_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockFullConfigTestStep(),
			{
				ResourceName:            "b1ddi_address_block.tf_acc_test_address_block",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceAddressBlockFullConfigTestStep() resource.TestStep {
	return resource.TestStep{
		Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
						
						asm_config {
							asm_threshold = 80
							enable = false
							enable_notification = false
							forecast_period = 9
							growth_factor = 15
							growth_type = "count"
							history = 50
							min_total = 20
							min_unused = 20
						}

						cidr = 24
						comment = "This Address Block is created by terraform provider acceptance test"
						ddns_client_update = "ignore"
						ddns_domain = "domain"
						ddns_generate_name = true
						ddns_generated_prefix = "tf_acc_host"
						ddns_send_updates = false
						ddns_update_on_renew = true
						ddns_use_conflict_resolution = false

						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

  						name = "tf_acc_test_address_block"
						
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						
						tags = {
							TestType = "Acceptance"
						}
					}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testAccAddressBlockExists("b1ddi_address_block.tf_acc_test_address_block"),
			testCheckAddressBlockInSpace("b1ddi_address_block.tf_acc_test_address_block", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "address", "192.168.1.0"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.asm_threshold", "80"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable_notification", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.forecast_period", "9"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_factor", "15"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_type", "count"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.history", "50"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_total", "20"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_unused", "20"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_scope_flag", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "cidr", "24"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "comment", "This Address Block is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_client_update", "ignore"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_domain", "domain"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generate_name", "true"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generated_prefix", "tf_acc_host"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_send_updates", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_update_on_renew", "true"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_use_conflict_resolution", "false"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.allow_unknown", "true"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.filters.#"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.ignore_list.#"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.lease_time", "3600"),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_host"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.0.option_value", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.0.type", "option"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_free", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_total", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_used", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_filename", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_address", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_name", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_char", "-"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_regex", "[^a-zA-Z0-9.-]"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_parent", ""),
			resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "name", "tf_acc_test_address_block"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "parent", ""),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "protocol", "ip4"),
			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "space"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "tags.%", "1"),

			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "updated_at"),
		),
	}
}

func TestAccResourceAddressBlock_UpdateAddressExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockBasicTestStep(),
			{
				Config: ` resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.15.0"
  						name = "tf_acc_test_address_block"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Address Block is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}`,
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "address", "192.168.1.0"),
				),
			},
			{
				ResourceName:            "b1ddi_address_block.tf_acc_test_address_block",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func TestAccResourceAddressBlock_UpdateSpaceExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockBasicTestStep(),
			{
				Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_ip_space" "tf_acc_new_test_space" {
  						name = "tf_acc_new_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
  						name = "tf_acc_test_address_block"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_new_test_space.id
  						comment = "This Address Block is created by terraform provider acceptance test"
						tags = {
							TestType = "Acceptance"
						}
					}`,
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
			},
			{
				ResourceName:            "b1ddi_address_block.tf_acc_test_address_block",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func TestAccResourceAddressBlock_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockBasicTestStep(),
			{
				Config: `
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_address_block" "tf_acc_test_address_block" {
						address = "192.168.1.0"
						
						asm_config {
							asm_threshold = 80
							enable = false
							enable_notification = false
							forecast_period = 9
							growth_factor = 15
							growth_type = "count"
							history = 50
							min_total = 20
							min_unused = 20
						}

						cidr = 24
						comment = "This Address Block is created by terraform provider acceptance test"
						ddns_client_update = "ignore"
						ddns_domain = "domain"
						ddns_generate_name = true
						ddns_generated_prefix = "tf_acc_host"
						ddns_send_updates = false
						ddns_update_on_renew = true
						ddns_use_conflict_resolution = false

						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

  						name = "tf_acc_test_address_block"
						
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						
						tags = {
							TestType = "Acceptance"
						}
					}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testAccAddressBlockExists("b1ddi_address_block.tf_acc_test_address_block"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "address", "192.168.1.0"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.asm_threshold", "80"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.enable_notification", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.forecast_period", "9"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_factor", "15"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.growth_type", "count"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.history", "50"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_total", "20"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.min_unused", "20"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "asm_scope_flag", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "cidr", "24"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "comment", "This Address Block is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_client_update", "ignore"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_domain", "domain"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generate_name", "true"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_generated_prefix", "tf_acc_host"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_send_updates", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_update_on_renew", "true"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "ddns_use_conflict_resolution", "false"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.allow_unknown", "true"),
					resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.filters.#"),
					resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.ignore_list.#"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_config.0.lease_time", "3600"),
					resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_host"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.0.option_value", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_options.0.type", "option"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_free", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_total", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_used", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "dhcp_utilization.0.dhcp_utilization", "0"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_filename", ""),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_address", ""),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "header_option_server_name", ""),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_char", "-"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "hostname_rewrite_regex", "[^a-zA-Z0-9.-]"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_parent", ""),
					resource.TestCheckNoResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "name", "tf_acc_test_address_block"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "parent", ""),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "protocol", "ip4"),
					resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "space"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "tags.%", "1"),

					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.high", "0"),
					resource.TestCheckResourceAttr("b1ddi_address_block.tf_acc_test_address_block", "threshold.0.low", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_address_block.tf_acc_test_address_block", "updated_at"),
				),
			},
			{
				ResourceName:            "b1ddi_address_block.tf_acc_test_address_block",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func testAccAddressBlockExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.AddressBlock.AddressBlockRead(
			&address_block.AddressBlockReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

// Checks if the specified AddressBlock resides in the specified IP Space
func testCheckAddressBlockInSpace(addressBlockPath, spacePath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		addressBlockResource, found := s.RootModule().Resources[addressBlockPath]
		if !found {
			return fmt.Errorf("not found %s", addressBlockPath)
		}
		if addressBlockResource.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", addressBlockPath)
		}
		space, found := s.RootModule().Resources[spacePath]
		if !found {
			return fmt.Errorf("not found %s", spacePath)
		}
		if space.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", spacePath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.AddressBlock.AddressBlockRead(
			&address_block.AddressBlockReadParams{ID: addressBlockResource.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.ID != addressBlockResource.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				addressBlockResource.Primary.ID)
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
