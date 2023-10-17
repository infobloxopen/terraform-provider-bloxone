package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigForwardNsg_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardNsgBasicTestStep(),
			{
				Config: `
					data "b1ddi_dns_forward_nsgs" "tf_acc_forward_nsg" {
						filters = {
							name = "tf_acc_test_forward_nsg"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.0.name", "tf_acc_test_forward_nsg"),
				),
			},
		},
	})
}

func TestAccDataSourceConfigForwardNsg_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsForwardNsgFullConfigTestStep(t),
			{
				Config: `
					data "b1ddi_dns_forward_nsgs" "tf_acc_forward_nsg" {
						filters = {
							name = "tf_acc_test_forward_nsg"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_forward_nsgs.tf_acc_forward_nsg", "results.0.name", "tf_acc_test_forward_nsg"),
				),
			},
		},
	})
}
