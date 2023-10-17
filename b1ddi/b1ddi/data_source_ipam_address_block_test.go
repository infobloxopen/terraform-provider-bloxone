package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcAddressBlock_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockBasicTestStep(),
			{
				Config: `data "b1ddi_address_blocks" "tf_acc_address_blocks" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_address_block"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.0.comment", "This Address Block is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcAddressBlock_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBlockFullConfigTestStep(),
			{
				Config: `
					data "b1ddi_address_blocks" "tf_acc_address_blocks" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_address_block"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_address_blocks.tf_acc_address_blocks", "results.0.comment", "This Address Block is created by terraform provider acceptance test"),
				),
			},
		},
	})
}
