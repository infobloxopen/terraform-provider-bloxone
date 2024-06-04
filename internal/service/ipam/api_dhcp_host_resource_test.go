package ipam_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/infobloxopen/bloxone-go-client/ipam"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDhcpHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_host.test"
	var v ipam.Host
	var dhcpServerName = acctest.RandomNameWithPrefix("dhcp_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpHostBasicConfig("TF_TEST_HOST_02", dhcpServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDhcpHostResource_Server(t *testing.T) {
	var resourceName = "bloxone_dhcp_host.test_server"
	var v ipam.Host
	var dhcpServerName1 = acctest.RandomNameWithPrefix("dhcp_host")
	var dhcpServerName2 = acctest.RandomNameWithPrefix("dhcp_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDhcpHostServer("TF_TEST_HOST_02", "server", dhcpServerName1, dhcpServerName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "server", "bloxone_dhcp_server.server", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDhcpHostServer("TF_TEST_HOST_02", "server2", dhcpServerName1, dhcpServerName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "server", "bloxone_dhcp_server.server2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDhcpHostExists(ctx context.Context, resourceName string, v *ipam.Host) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			DhcpHostAPI.
			Read(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccBaseWithDhcpConfig(hostName, name string) string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "test_host" {
    filters = {
		display_name = %q 
	}
}

resource "bloxone_dhcp_server" "server" {
    name = %q
}
`, hostName, name)
}

func testAccDhcpHostBasicConfig(hostName, dhcpServerName string) string {
	config := `
resource "bloxone_dhcp_host" "test" {
	id = data.bloxone_infra_hosts.test_host.results.0.legacy_id
	server = bloxone_dhcp_server.server.id
}
`
	return strings.Join([]string{testAccBaseWithDhcpConfig(hostName, dhcpServerName), config}, "")
}

func testAccDhcpHostServer(hostName, dhcpServer, dhcpServerName1, dhcpServerName2 string) string {
	config := fmt.Sprintf(`

resource "bloxone_dhcp_server" "server2" {
	name = %q
}

resource "bloxone_dhcp_host" "test_server" {
	id = data.bloxone_infra_hosts.test_host.results.0.legacy_id
	server = bloxone_dhcp_server.%s.id
}
`, dhcpServerName2, dhcpServer)
	return strings.Join([]string{testAccBaseWithDhcpConfig(hostName, dhcpServerName1), config}, "")

}
