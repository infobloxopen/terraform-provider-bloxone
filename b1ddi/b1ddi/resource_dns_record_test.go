package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_data/record"
	"regexp"
	"testing"
)

func TestAccResourceDnsRecord_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsRecordBasicTestStep(t *testing.T) resource.TestStep {
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
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_record" {
						zone = b1ddi_dns_auth_zone.tf_acc_test_auth_zone.id
						name_in_zone = "tf_acc_test_a_record"
						rdata = {
							"address" = "192.168.1.15"
						}
						type = "A"
					}`, testAccReadDnsHost(t),
		),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_zone_name", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "comment", ""),
			resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "created_at"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "delegation", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "disabled", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_zone_name", "tf-acc-test.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_name_in_zone", "tf_acc_test_a_record"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_rdata", "192.168.1.15"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "inheritance_sources.#"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "name_in_zone", "tf_acc_test_a_record"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "options"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "rdata.address", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "source.0", "STATIC"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "tags"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "ttl", "28800"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "type", "A"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "updated_at"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "view"),
			resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "view_name", "tf_acc_test_dns_view"),
			resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "zone"),
		),
	}
}

func TestAccResourceDnsRecord_full_config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
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
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_record" {
						comment = "This DNS Record is created by the terraform provider acceptance test"
						disabled = true
						name_in_zone = "tf_acc_test_a_record"
						options = {
							"create_ptr" = true
 							"check_rmz" = true
						}
						rdata = {
							"address" = "192.168.1.15"
						}
						tags = {
							TestType = "Acceptance"
						}
						ttl = 24400
						type = "A"
						zone = b1ddi_dns_auth_zone.tf_acc_test_auth_zone.id
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_zone_name", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "comment", "This DNS Record is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "created_at"),
					// ToDo Add check for custom delegation
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "delegation", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "disabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_zone_name", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_rdata", "192.168.1.15"),
					// ToDo Add check for custom inheritance_sources
					resource.TestCheckNoResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "options.create_ptr", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "options.check_rmz", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "rdata.address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "source.0", "STATIC"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "ttl", "24400"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "type", "A"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "updated_at"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "view"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "view_name", "tf_acc_test_dns_view"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "zone"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsRecord_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
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
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_record" {
						comment = "This DNS Record is created by the terraform provider acceptance test"
						disabled = true
						name_in_zone = "tf_acc_test_a_record"
						options = {
							"create_ptr" = true
						}
						rdata = {
							"address" = "192.168.1.15"
						}
						tags = {
							TestType = "Acceptance"
						}
						ttl = 24400
						type = "A"
						zone = b1ddi_dns_auth_zone.tf_acc_test_auth_zone.id
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "absolute_zone_name", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "comment", "This DNS Record is created by the terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "created_at"),
					// ToDo Add check for custom delegation
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "delegation", ""),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "disabled", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_name_spec", "tf_acc_test_a_record.tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_absolute_zone_name", "tf-acc-test.com."),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "dns_rdata", "192.168.1.15"),
					// ToDo Add check for custom inheritance_sources
					resource.TestCheckNoResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "inheritance_sources.#"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "name_in_zone", "tf_acc_test_a_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "options.create_ptr", "true"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "rdata.address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "source.0", "STATIC"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "ttl", "24400"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "type", "A"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "updated_at"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "view"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "view_name", "tf_acc_test_dns_view"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "zone"),
					// Test SRV record
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "name_in_zone", "tf_acc_test_srv_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "rdata.port", "2100"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "rdata.weight", "100"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "source.0", "STATIC"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "ttl", "24400"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "type", "SRV"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "updated_at"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "view"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "view_name", "tf_acc_test_dns_view"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "zone"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsSRVRecord_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
			{
				Config: fmt.Sprintf(`resource "b1ddi_dns_view" "tf_acc_test_dns_view" {
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
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_srv_record" {
						comment = "This DNS Record is created by the terraform provider acceptance test"
						disabled = true
						name_in_zone = "tf_acc_test_srv_record"
						rdata = {
							"weight" = 100
							"port" = 2100
							"priority" = 0
						}
						tags = {
							TestType = "Acceptance"
						}
						ttl = 24400
						type = "SRV"
						zone = b1ddi_dns_auth_zone.tf_acc_test_auth_zone.id
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "name_in_zone", "tf_acc_test_srv_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "rdata.port", "2100"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "rdata.weight", "100"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "rdata.priority", "0"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "source.0", "STATIC"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "ttl", "24400"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "type", "SRV"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "updated_at"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "view"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_srv_record", "view_name", "tf_acc_test_dns_view"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_srv_record", "zone"),
				),
			},
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_srv_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsRecord_UpdateTypeExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
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
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_record" {
						comment = "This DNS Record is created by the terraform provider acceptance test"
						disabled = true
						name_in_zone = "tf_acc_test_a_record"
						rdata = {
							"address" = "192.168.1.15"
						}
						type = "PTR"
						zone = b1ddi_dns_auth_zone.tf_acc_test_auth_zone.id
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
					resource.TestCheckResourceAttr("b1ddi_dns_record.tf_acc_test_dns_record", "type", "A"),
				),
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
			},
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceDnsRecord_UpdateViewExpectError(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsRecordBasicTestStep(t),
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
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}
					resource "b1ddi_dns_auth_zone" "tf_acc_test_new_auth_zone" {
						internal_secondaries {
							host = data.b1ddi_dns_hosts.dns_host.results.0.id
						}
						fqdn = "tf-new-acc-test.com."
						primary_type = "cloud"
						view = b1ddi_dns_view.tf_acc_test_dns_view.id
					}
					resource "b1ddi_dns_record" "tf_acc_test_dns_record" {
						comment = "This DNS Record is created by the terraform provider acceptance test"
						disabled = true
						name_in_zone = "tf_acc_test_a_record"
						rdata = {
							"address" = "192.168.1.15"
						}
						type = "PTR"
						zone = b1ddi_dns_auth_zone.tf_acc_test_new_auth_zone.id
					}`, testAccReadDnsHost(t),
				),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccDnsRecordExists("b1ddi_dns_record.tf_acc_test_dns_record"),
					resource.TestCheckResourceAttrSet("b1ddi_dns_record.tf_acc_test_dns_record", "zone"),
				),
				ExpectError: regexp.MustCompile("changing the value of '[a-z]*' field is not allowed"),
			},
			{
				ResourceName:      "b1ddi_dns_record.tf_acc_test_dns_record",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsRecordExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSDataAPI.Record.RecordRead(
			&record.RecordReadParams{ID: res.Primary.ID, Context: context.TODO()},
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
