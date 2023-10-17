package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcFixedAddress_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressBasicTestStep(),
			{
				Config: `
					data "b1ddi_fixed_addresses" "tf_acc_fixed_addresses" {
						filters = {
							"name" = "tf_acc_test_fixed_address"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.comment", "This Fixed Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcFixedAddress_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceFixedAddressFullConfigTestStep(),
			{
				Config: `
					data "b1ddi_fixed_addresses" "tf_acc_fixed_addresses" {
						filters = {
							"name" = "tf_acc_test_fixed_address_full_config"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_fixed_addresses.tf_acc_fixed_addresses", "results.0.comment", "This Fixed Address is created by terraform provider acceptance test"),
				),
			},
		},
	})
}
