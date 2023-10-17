package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigView_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsViewBasicTestStep(),
			{
				Config: `data "b1ddi_dns_views" "tf_acc_dns_views" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_dns_view"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_views.tf_acc_dns_views", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.0.name", "tf_acc_test_dns_view"),
				),
			},
		},
	})
}

func TestAccDataSourceConfigView_FullConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsViewFullConfigTestStep(),
			{
				Config: `data "b1ddi_dns_views" "tf_acc_dns_views" {
						filters = {
							# Check string filter
							"name" = "tf_acc_test_dns_view_full_config"
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_views.tf_acc_dns_views", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_views.tf_acc_dns_views", "results.0.name", "tf_acc_test_dns_view_full_config"),
				),
			},
		},
	})
}
