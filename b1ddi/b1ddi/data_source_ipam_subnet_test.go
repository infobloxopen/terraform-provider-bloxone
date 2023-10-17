package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcSubnet_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create Subnet resource
			resourceSubnetBasicTestStep(),
			// Check Subnet data source
			{
				Config: `
					data "b1ddi_subnets" "tf_acc_subnets" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_subnet"
							"address" = "192.168.1.0"
							"cidr" = 24
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_subnets.tf_acc_subnets", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.0.comment", "This Subnet is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcSubnet_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create Subnet resource
			resourceSubnetFullConfigTestStep(t),
			// Check Subnet data source
			{
				Config: `
					data "b1ddi_subnets" "tf_acc_subnets" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_subnet"
							"address" = "192.168.1.0"
							"cidr" = 24
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_subnets.tf_acc_subnets", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_subnets.tf_acc_subnets", "results.0.comment", "This Subnet is created by terraform provider acceptance test"),
				),
			},
		},
	})
}
