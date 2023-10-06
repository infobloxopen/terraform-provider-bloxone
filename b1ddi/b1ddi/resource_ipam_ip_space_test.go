package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/ip_space"
	"testing"
)

// ToDo add test case for IP Space with DHCP options
// ToDo add test case for IP Space with Inheritance Sources
// ToDo add check deleted

func TestAccResourceIPSpace_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceIPSpaceBasicTestStep(),
			{
				ResourceName:            "b1ddi_ip_space.tf_acc_test_space",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceIPSpaceBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			// Check default values
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.asm_threshold", "90"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable_notification", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.forecast_period", "14"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_factor", "20"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_type", "percent"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.history", "30"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_total", "10"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_unused", "10"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_scope_flag", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "comment", "This IP Space is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_client_update", "client"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_domain", ""),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generate_name", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generated_prefix", "myhost"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_send_updates", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_update_on_renew", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_use_conflict_resolution", "true"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.allow_unknown", "true"),
			resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.filters.#"),
			resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.#"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.lease_time", "3600"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_filename", ""),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_address", ""),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_name", ""),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_char", "_"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_regex", "[^a-zA-Z0-9_.]"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "inheritance_sources.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "name", "tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "tags.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "updated_at"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandon_utilization", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandoned", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.dynamic", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.free", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.static", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.total", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.used", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "vendor_specific_option_option_space", ""),
		),
	}
}

func TestAccResourceIPSpace_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceIPSpaceFullConfigTestStep(),
			{
				ResourceName:            "b1ddi_ip_space.tf_acc_test_space",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceIPSpaceFullConfigTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
				data "b1ddi_option_codes" "tf_acc_option_code" {
					filters = {
						"name" = "routers"
					}
				}

				resource "b1ddi_ip_space" "tf_acc_test_space" {
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
					comment = "This IP Space is created by terraform provider acceptance test"
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
						ignore_list {
							type = "hardware"
							value = "00:00:00:00:00:00"
						}
					}

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
			
					#inheritance_sources {
					#	asm_config {
					#		asm_enable_block {
					#			action = "inherit"
					#		}
					#	}
					#}

					name = "tf_acc_test_space"
					tags = {
						TestType = "Acceptance"
					}

					
				}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.asm_threshold", "80"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable_notification", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.forecast_period", "9"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_factor", "20"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_type", "count"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.history", "50"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_total", "20"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_unused", "20"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_scope_flag", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "comment", "This IP Space is created by terraform provider acceptance test"),

			resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_client_update", "ignore"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_domain", "domain"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generate_name", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generated_prefix", "tf_acc_host"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_send_updates", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_update_on_renew", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_use_conflict_resolution", "false"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.allow_unknown", "false"),
			resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.filters.#"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.lease_time", "1800"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.0.type", "hardware"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.0.value", "00:00:00:00:00:00"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.0.option_value", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.0.type", "option"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_filename", "Acc Test Header"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_address", "192.168.1.10"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_name", "Header Option Server Name"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_char", " "),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_enabled", "true"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_regex", "[aaa bbb]"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "inheritance_sources.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "name", "tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "tags.%", "1"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "tags.TestType", "Acceptance"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "updated_at"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandon_utilization", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandoned", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.dynamic", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.free", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.static", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.total", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.used", "0"),
			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.utilization", "0"),

			resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "vendor_specific_option_option_space", ""),
		),
	}
}

func TestAccResourceIPSpace_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceIPSpaceBasicTestStep(),
			{
				Config: fmt.Sprintf(`
				data "b1ddi_option_codes" "tf_acc_option_code" {
					filters = {
						"name" = "routers"
					}
				}
				resource "b1ddi_ip_space" "tf_acc_test_space" {
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
					comment = "This IP Space is updated by terraform provider acceptance test"
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
						ignore_list {
							type = "hardware"
							value = "00:00:00:00:00:00"
						}
					}

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

					name = "tf_acc_test_space_updated"
					tags = {
						TestType = "Acceptance"
					}

					
				}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.asm_threshold", "80"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.enable_notification", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.forecast_period", "9"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_factor", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.growth_type", "count"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.history", "50"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_total", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.min_unused", "20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_config.0.reenable_date", "1970-01-01T00:00:00.000Z"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "asm_scope_flag", "0"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "comment", "This IP Space is updated by terraform provider acceptance test"),

					resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "created_at"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_client_update", "ignore"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_domain", "domain"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generate_name", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_generated_prefix", "tf_acc_host"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_send_updates", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_update_on_renew", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "ddns_use_conflict_resolution", "false"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.allow_unknown", "false"),
					resource.TestCheckNoResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.filters.#"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.lease_time", "1800"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.0.type", "hardware"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_config.0.ignore_list.0.value", "00:00:00:00:00:00"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.0.option_value", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "dhcp_options.0.type", "option"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_filename", "Acc Test Header"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_address", "192.168.1.10"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "header_option_server_name", "Header Option Server Name"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_char", " "),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_enabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "hostname_rewrite_regex", "[aaa bbb]"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "inheritance_sources.%", "0"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "name", "tf_acc_test_space_updated"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "tags.TestType", "Acceptance"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.high", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "threshold.0.low", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_ip_space.tf_acc_test_space", "updated_at"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandon_utilization", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.abandoned", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.dynamic", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.free", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.static", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.total", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.used", "0"),
					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "utilization.0.utilization", "0"),

					resource.TestCheckResourceAttr("b1ddi_ip_space.tf_acc_test_space", "vendor_specific_option_option_space", ""),
				),
			},
			{
				ResourceName:            "b1ddi_ip_space.tf_acc_test_space",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func testCheckIPSpaceExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.IPSpace.IPSpaceRead(
			&ip_space.IPSpaceReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.ID != res.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				res.Primary.ID,
			)
		}

		return nil
	}
}
