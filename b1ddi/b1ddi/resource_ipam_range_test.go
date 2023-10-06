package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/range_operations"
	"regexp"
	"testing"
)

func TestAccResourceRange_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceRangeBasicTestStep(),
			{
				ResourceName:            "b1ddi_range.tf_acc_test_range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceRangeBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
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
					}
					resource "b1ddi_range" "tf_acc_test_range" {
						start = "192.168.1.15"
						end = "192.168.1.30"
  						name = "tf_acc_test_range"
						space = b1ddi_ip_space.tf_acc_test_space.id 
  						comment = "This Range is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
			testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
			testAccRangeExists("b1ddi_range.tf_acc_test_range"),
			testCheckRangeInSpace("b1ddi_range.tf_acc_test_range", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "comment", "This Range is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_host", ""),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "end", "192.168.1.30"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.%", "0"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "inheritance_parent"),
			resource.TestCheckNoResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "name", "tf_acc_test_range"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "parent"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "protocol", "ip4"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "start", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "tags.%", "0"),

			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "updated_at"),
		),
	}
}

func TestAccResourceRange_full_config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceRangeFullConfigTestStep(),
			{
				ResourceName:            "b1ddi_range.tf_acc_test_range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func resourceRangeFullConfigTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
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
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_range" "tf_acc_test_range" {
						comment = "This Range is created by terraform provider acceptance test"
						#dhcp_host = "dhcp_host"
						
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						end = "192.168.1.30"
						exclusion_ranges {
							comment = "This Exclusion Range is created by terraform provider acceptance test"
							end = "192.168.1.25"
							start = "192.168.1.20"
						}
						#inheritance_assigned_hosts
						#inheritance_parent
						#inheritance_sources {}
						name = "tf_acc_test_range"
						#parent
						space = b1ddi_ip_space.tf_acc_test_space.id
						start = "192.168.1.15"
						tags = {
							TestType = "Acceptance"
						}
						#threshold {}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
			testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
			testAccRangeExists("b1ddi_range.tf_acc_test_range"),
			testCheckRangeInSpace("b1ddi_range.tf_acc_test_range", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "comment", "This Range is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_host", ""),

			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.0.option_value", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.0.type", "option"),

			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "end", "192.168.1.30"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.comment", "This Exclusion Range is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.end", "192.168.1.25"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.start", "192.168.1.20"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_assigned_hosts.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "inheritance_parent"),
			resource.TestCheckNoResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "name", "tf_acc_test_range"),
			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "parent"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "protocol", "ip4"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "start", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "tags.%", "1"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "tags.TestType", "Acceptance"),

			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.enabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.high", "0"),
			resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.low", "0"),

			resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "updated_at"),
		),
	}
}

func TestAccResourceRange_UpdateSpaceExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceRangeBasicTestStep(),
			{
				Config: fmt.Sprintf(`
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
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_range" "tf_acc_test_range" {
						start = "192.168.1.15"
						end = "192.168.1.30"
  						name = "tf_acc_test_range"
						space = b1ddi_ip_space.tf_acc_new_test_space.id 
  						comment = "This Range is created by terraform provider acceptance test"
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_new_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_new_test_subnet"),
					testCheckSubnetInSpace("b1ddi_subnet.tf_acc_new_test_subnet", "b1ddi_ip_space.tf_acc_new_test_space"),
					testAccRangeExists("b1ddi_range.tf_acc_test_range"),
					testCheckRangeInSpace("b1ddi_range.tf_acc_test_range", "b1ddi_ip_space.tf_acc_new_test_space"),
				),
			},
			{
				ResourceName:            "b1ddi_range.tf_acc_test_range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func TestAccResourceRange_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceRangeBasicTestStep(),
			{
				Config: fmt.Sprintf(`
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
					}
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
					resource "b1ddi_range" "tf_acc_test_range" {
						comment = "This Range is created by terraform provider acceptance test"
						#dhcp_host = "dhcp_host"
						
						dhcp_options {
							option_code = data.b1ddi_option_codes.tf_acc_option_code.results.0.id
							option_value = "192.168.1.20"
							type = "option"
						}

						end = "192.168.1.30"
						exclusion_ranges {
							comment = "This Exclusion Range is created by terraform provider acceptance test"
							end = "192.168.1.25"
							start = "192.168.1.20"
						}
						#inheritance_assigned_hosts
						#inheritance_parent
						#inheritance_sources {}
						name = "tf_acc_test_range"
						#parent
						space = b1ddi_ip_space.tf_acc_test_space.id
						start = "192.168.1.15"
						tags = {
							TestType = "Acceptance"
						}
						#threshold {}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testCheckSubnetInSpace("b1ddi_subnet.tf_acc_test_subnet", "b1ddi_ip_space.tf_acc_test_space"),
					testAccRangeExists("b1ddi_range.tf_acc_test_range"),
					testCheckRangeInSpace("b1ddi_range.tf_acc_test_range", "b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "comment", "This Range is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "created_at"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_host", ""),

					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.0.option_value", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "dhcp_options.0.type", "option"),

					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "end", "192.168.1.30"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.comment", "This Exclusion Range is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.end", "192.168.1.25"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "exclusion_ranges.0.start", "192.168.1.20"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_assigned_hosts.%", "0"),
					resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "inheritance_parent"),
					resource.TestCheckNoResourceAttr("b1ddi_range.tf_acc_test_range", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "name", "tf_acc_test_range"),
					resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "parent"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "protocol", "ip4"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "start", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "tags.TestType", "Acceptance"),

					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.enabled", "false"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.high", "0"),
					resource.TestCheckResourceAttr("b1ddi_range.tf_acc_test_range", "threshold.0.low", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_range.tf_acc_test_range", "updated_at"),
				),
			},
			{
				ResourceName:            "b1ddi_range.tf_acc_test_range",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"updated_at", "utilization"},
			},
		},
	})
}

func testAccRangeExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.RangeOperations.RangeRead(
			&range_operations.RangeReadParams{ID: res.Primary.ID, Context: context.TODO()},
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

// Checks if the specified Range resides in the specified IP Space
func testCheckRangeInSpace(rangePath, spacePath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rangeResource, found := s.RootModule().Resources[rangePath]
		if !found {
			return fmt.Errorf("not found %s", rangePath)
		}
		if rangeResource.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", rangePath)
		}
		space, found := s.RootModule().Resources[spacePath]
		if !found {
			return fmt.Errorf("not found %s", spacePath)
		}
		if space.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", spacePath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.RangeOperations.RangeRead(
			&range_operations.RangeReadParams{ID: rangeResource.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.ID != rangeResource.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				rangeResource.Primary.ID)
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
