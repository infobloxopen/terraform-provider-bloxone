package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBasicTestStep(),
			{
				Config: `
					data "b1ddi_addresses" "tf_acc_addresses" {
						filters = {
							# Check string filter
							"address" = "192.168.1.15"
							"comment" = "This Address is created by terraform provider acceptance test"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_addresses.tf_acc_addresses", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_addresses.tf_acc_addresses", "results.0.comment", "This Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}
