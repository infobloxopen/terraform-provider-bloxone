package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/fixed_address"
	"testing"
)

func TestAccResourceFixedAddress_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressBasicTestStep(),
			{
				ResourceName:      "b1ddi_fixed_address.tf_acc_test_fixed_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceFixedAddressBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_fixed_address" "tf_acc_test_fixed_address" {
						name = "tf_acc_test_fixed_address"						
						address = "192.168.1.15"
						ip_space = b1ddi_ip_space.tf_acc_test_space.id
						match_type = "mac"
						match_value = "00:00:00:00:00:00"
						comment = "This Fixed Address is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckFixedAddressExists("b1ddi_fixed_address.tf_acc_test_fixed_address"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "address", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "comment", "This Fixed Address is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_filename", ""),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_address", ""),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_name", ""),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "hostname", ""),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_parent"),
			resource.TestCheckNoResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_sources.#"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "ip_space"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_type", "mac"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_value", "00:00:00:00:00:00"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "name", "tf_acc_test_fixed_address"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "parent"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "updated_at"),
		),
	}
}

func TestAccResourceFixedAddress_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressFullConfigTestStep(),
			{
				ResourceName:      "b1ddi_fixed_address.tf_acc_test_fixed_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceFixedAddressFullConfigTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_fixed_address" "tf_acc_test_fixed_address" {
						address = "192.168.1.15"
						comment = "This Fixed Address is created by terraform provider acceptance test"
						
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						header_option_filename = "Acc Test Header"
						header_option_server_address = "192.168.1.10"
						header_option_server_name = "Header Option Server Name"

						hostname = ""

						#inheritance_sources {}

						ip_space = b1ddi_ip_space.tf_acc_test_space.id
						match_type = "client_text"
						match_value = "Client Text"
						name = "tf_acc_test_fixed_address_full_config"
						tags = {
							TestType = "Acceptance"
						}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckFixedAddressExists("b1ddi_fixed_address.tf_acc_test_fixed_address"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "address", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "comment", "This Fixed Address is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "created_at"),

			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.0.option_value", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.0.type", "option"),

			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_filename", "Acc Test Header"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_address", "192.168.1.10"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_name", "Header Option Server Name"),
			// ToDo Check hostname
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "hostname", ""),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_parent"),
			resource.TestCheckNoResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_sources.#"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "ip_space"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_type", "client_text"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_value", "Client Text"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "name", "tf_acc_test_fixed_address_full_config"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "parent"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.%", "1"),
			resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.TestType", "Acceptance"),
			resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "updated_at"),
		),
	}
}

func TestAccResourceFixedAddress_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						cidr = 24
						space = b1ddi_ip_space.tf_acc_test_space.id
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_fixed_address" "tf_acc_test_fixed_address" {
						address = "192.168.1.15"
						comment = "This Fixed Address is created by terraform provider acceptance test"
						
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						header_option_filename = "Acc Test Header"
						header_option_server_address = "192.168.1.10"
						header_option_server_name = "Header Option Server Name"

						hostname = ""

						#inheritance_sources {}

						ip_space = b1ddi_ip_space.tf_acc_test_space.id
						match_type = "client_text"
						match_value = "Client Text"
						name = "tf_acc_test_fixed_address_full_config"
						tags = {
							TestType = "Acceptance"
						}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckFixedAddressExists("b1ddi_fixed_address.tf_acc_test_fixed_address"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "comment", "This Fixed Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.0.option_value", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "dhcp_options.0.type", "option"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_filename", "Acc Test Header"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_address", "192.168.1.10"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "header_option_server_name", "Header Option Server Name"),
					// ToDo Check hostname
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "hostname", ""),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_assigned_hosts.%", "0"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_parent"),
					resource.TestCheckNoResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "inheritance_sources.#"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "ip_space"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_type", "client_text"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "match_value", "Client Text"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "name", "tf_acc_test_fixed_address_full_config"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "parent"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_fixed_address.tf_acc_test_fixed_address", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttrSet("b1ddi_fixed_address.tf_acc_test_fixed_address", "updated_at"),
				),
			},
			{
				ResourceName:      "b1ddi_fixed_address.tf_acc_test_fixed_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckFixedAddressExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.FixedAddress.FixedAddressRead(
			&fixed_address.FixedAddressReadParams{ID: res.Primary.ID, Context: context.TODO()},
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
