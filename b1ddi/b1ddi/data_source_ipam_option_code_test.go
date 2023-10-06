package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcOptionCode(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_option_codes" "tf_acc_option_code" {
						filters = {
							"name" = "routers"
						}
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.b1ddi_option_codes.tf_acc_option_code", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_option_codes.tf_acc_option_code", "results.0.code", "3"),
				),
			},
		},
	})
}
