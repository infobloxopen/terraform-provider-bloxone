package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceConfigAuthZone_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneBasicTestStep(t),
			{
				Config: `
					data "b1ddi_dns_auth_zones" "tf_acc_auth_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

func TestAccDataSourceConfigAuthZone_FullConfigCloud(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneFullConfigCloudTestStep(t),
			{
				Config: `
					data "b1ddi_dns_auth_zones" "tf_acc_auth_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}

func TestAccDataSourceConfigAuthZone_FullConfigExternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceDnsAuthZoneFullConfigExternalTestStep(t),
			{
				Config: `
					data "b1ddi_dns_auth_zones" "tf_acc_auth_zones" {
						filters = {
							fqdn = "tf-acc-test.com."
						}
					}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.#", "1"),
					resource.TestCheckResourceAttrSet("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.id"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_auth_zones.tf_acc_auth_zones", "results.0.fqdn", "tf-acc-test.com."),
				),
			},
		},
	})
}
