package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/auth_nsg"
	"testing"
)

func TestAccResourceDnsAuthNsg_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthNsgBasicTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsAuthNsgBasicTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
		data "b1ddi_dns_hosts" "dns_host" {
			filters = {
				"name" = "%s"
			}
		}
		resource "b1ddi_dns_auth_nsg" "tf_acc_test_auth_nsg" {
			name = "tf_acc_test_auth_nsg"
			internal_secondaries {
				host = data.b1ddi_dns_hosts.dns_host.results.0.id
			}
		}`, testAccReadDnsHost(t)),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsAuthNsgExists("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "comment", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_primaries.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_secondaries.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "internal_secondaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "name", "tf_acc_test_auth_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "nsgs.#", "0"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "tags"),
		),
	}
}

func TestAccResourceDnsAuthNsg_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthNsgFullConfigTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsAuthNsgFullConfigTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
		data "b1ddi_dns_hosts" "dns_host" {
			filters = {
				"name" = "%s"
			}
		}
		resource "b1ddi_dns_auth_nsg" "tf_acc_test_auth_nsg" {
			comment = "This Auth NSG is created by the terraform provider acceptance test"

			external_primaries {
				address = "192.168.1.60"
				fqdn = "tf_test_external_primary."
				type = "primary"
			}

			external_secondaries {
				address = "192.168.1.50"
				fqdn = "tf_test_external_secondary."
			}

			name = "tf_acc_test_auth_nsg"
			internal_secondaries {
				host = data.b1ddi_dns_hosts.dns_host.results.0.id
			}
			tags = {
				TestType = "Acceptance"
			}
		}`, testAccReadDnsHost(t)),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsAuthNsgExists("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "comment", "This Auth NSG is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_primaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_primaries.0.address", "192.168.1.60"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_primaries.0.fqdn", "tf_test_external_primary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_primaries.0.type", "primary"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_secondaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_secondaries.0.address", "192.168.1.50"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "external_secondaries.0.fqdn", "tf_test_external_secondary."),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "internal_secondaries.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "name", "tf_acc_test_auth_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg", "tags.TestType", "Acceptance"),
		),
	}
}

func TestAccResourceDnsAuthNsg_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthNsgBasicTestStep(t),
			resourceDnsAuthNsgFullConfigTestStep(t),
			{
				ResourceName:      "b1ddi_dns_auth_nsg.tf_acc_test_auth_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsAuthNsgExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSConfigurationAPI.AuthNsg.AuthNsgRead(
			&auth_nsg.AuthNsgReadParams{ID: res.Primary.ID, Context: context.TODO()},
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
