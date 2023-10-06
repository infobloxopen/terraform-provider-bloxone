package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigAuthNsg_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthNsgBasicTestStep(t),
			{
				Config: `
					data "b1ddi_dns_auth_nsgs" "tf_acc_auth_nsg" {
						filters = {
							name = "tf_acc_test_auth_nsg"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.0.name", "tf_acc_test_auth_nsg"),
				),
			},
		},
	})
}

func TestAccDataSourceConfigAuthNsg_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthNsgFullConfigTestStep(t),
			{
				Config: `
					data "b1ddi_dns_auth_nsgs" "tf_acc_auth_nsg" {
						filters = {
							name = "tf_acc_test_auth_nsg"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_nsgs.tf_acc_auth_nsg", "results.0.name", "tf_acc_test_auth_nsg"),
				),
			},
		},
	})
}
