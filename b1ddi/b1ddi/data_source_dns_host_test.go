package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceDnsHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_hosts" "dns_hosts" {}
				`),
			},
		},
	})
}

func TestAccDataSourceDnsHostByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dns_hosts" "dns_host" {
						filters = {
							"name" = "%s"
						}
					}
				`, testAccReadDnsHost(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dns_hosts.dns_host", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dns_hosts.dns_host", "results.0.name", testAccReadDnsHost(t)),
				),
			},
		},
	})
}
