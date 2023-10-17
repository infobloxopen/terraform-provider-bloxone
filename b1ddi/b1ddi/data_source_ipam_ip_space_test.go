package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcIPSpace_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create IP Space resource
			resourceIPSpaceBasicTestStep(),
			// Test IP Space Data Source
			{
				Config: `
					data "b1ddi_ip_spaces" "tf_acc_spaces" {
						filters = {
							# Check integer filter
							"asm_scope_flag" = 0
							# Check string filter
							"name" = "tf_acc_test_space"
							# Check bool filter
							"hostname_rewrite_enabled" = false
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_ip_spaces.tf_acc_spaces", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_ip_spaces.tf_acc_spaces", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_ip_spaces.tf_acc_spaces", "results.0.comment", "This IP Space is created by terraform provider acceptance test"),
				),
			},
		},
	})
}

func TestAccDataSourceIpamsvcIPSpace_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			// Create IP Space resource
			resourceIPSpaceFullConfigTestStep(),
			// Test IP Space Data Source
			{
				Config: `
					data "b1ddi_ip_spaces" "tf_acc_spaces" {
						filters = {
							# Check integer filter
							"asm_scope_flag" = 0
							# Check string filter
							"name" = "tf_acc_test_space"
							# Check bool filter
							"hostname_rewrite_enabled" = true
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_ip_spaces.tf_acc_spaces", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_ip_spaces.tf_acc_spaces", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_ip_spaces.tf_acc_spaces", "results.0.comment", "This IP Space is created by terraform provider acceptance test"),
				),
			},
		},
	})
}
