package b1ddi

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccDataSourceIpamsvcDhcpHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					data "b1ddi_dhcp_hosts" "dhcp_hosts" {}
				`,
			},
		},
	})
}

func TestAccDataSourceIpamsvcDhcpHostByName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					data "b1ddi_dhcp_hosts" "dhcp_hosts" {
						filters = {
							"name" = "%s"
						}
					}
				`, testAccReadDhcpHost(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.b1ddi_dhcp_hosts.dhcp_hosts", "results.#", "1"),
					resource.TestCheckResourceAttr("data.b1ddi_dhcp_hosts.dhcp_hosts", "results.0.name", testAccReadDhcpHost(t)),
				),
			},
		},
	})
}
