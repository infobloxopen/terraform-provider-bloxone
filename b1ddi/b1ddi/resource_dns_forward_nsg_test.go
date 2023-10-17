package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/dns_config/forward_nsg"
	"testing"
)

func TestAccResourceDnsForwardNsg_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardNsgBasicTestStep(),
			{
				ResourceName:      "b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsForwardNsgBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: `
		resource "b1ddi_dns_forward_nsg" "tf_acc_test_forward_nsg" {
			name = "tf_acc_test_forward_nsg"
		}`,
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsForwardNsgExists("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "comment", ""),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "external_forwarders.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "forwarders_only", "false"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "hosts.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "internal_forwarders.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "name", "tf_acc_test_forward_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "nsgs.#", "0"),
			resource.TestCheckNoResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "tags"),
		),
	}
}

func TestAccResourceDnsForwardNsg_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardNsgFullConfigTestStep(t),
			{
				ResourceName:      "b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func resourceDnsForwardNsgFullConfigTestStep(t *testing.T) resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
			data "b1ddi_dns_hosts" "dns_host" {
				filters = {
					"name" = "%s"
				}
			}
			resource "b1ddi_dns_forward_nsg" "tf_acc_test_forward_nsg" {
				comment = "This Auth NSG is created by the terraform provider acceptance test"
				external_forwarders {
					address = "192.168.1.70"
					fqdn = "tf_acc_test_forwarder.infolbox.com."
				}
				forwarders_only = true
				hosts = [data.b1ddi_dns_hosts.dns_host.results.0.id]
				name = "tf_acc_test_forward_nsg"
				tags = {
					TestType = "Acceptance"
				}
		}`, testAccReadDnsHost(t)),
		Check: resource.ComposeAggregateTestCheckFunc(
			testAccDnsForwardNsgExists("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "comment", "This Auth NSG is created by the terraform provider acceptance test"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "external_forwarders.#", "1"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "external_forwarders.0.address", "192.168.1.70"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "external_forwarders.0.fqdn", "tf_acc_test_forwarder.infolbox.com."),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "forwarders_only", "true"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "hosts.#", "1"),
			// ToDo Add internal forwarders
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "internal_forwarders.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "name", "tf_acc_test_forward_nsg"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "nsgs.#", "0"),
			resource.TestCheckResourceAttr("b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg", "tags.TestType", "Acceptance"),
		),
	}
}

func TestAccResourceDnsForwardNsg_Update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardNsgBasicTestStep(),
			resourceDnsForwardNsgFullConfigTestStep(t),
			{
				ResourceName:      "b1ddi_dns_forward_nsg.tf_acc_test_forward_nsg",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccDnsForwardNsgExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.DNSConfigurationAPI.ForwardNsg.ForwardNsgRead(
			&forward_nsg.ForwardNsgReadParams{ID: res.Primary.ID, Context: context.TODO()},
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
