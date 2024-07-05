package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/infobloxopen/bloxone-go-client/dnsconfig"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDnsHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_host.test"
	var v dnsconfig.Host
	var dnsServerName = acctest.RandomNameWithPrefix("dns_server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDnsHostBasicConfig("TF_TEST_HOST_02", dnsServerName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrSet(resourceName, "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDnsHostResource_Server(t *testing.T) {
	var resourceName = "bloxone_dns_host.test_server"
	var v dnsconfig.Host
	var dnsServerName1 = acctest.RandomNameWithPrefix("dns_server")
	var dnsServerName2 = acctest.RandomNameWithPrefix("dns_server")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDnsHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDnsHostServer("TF_TEST_HOST_02", "server", dnsServerName1, dnsServerName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "server", "bloxone_dns_server.server", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDnsHostServer("TF_TEST_HOST_02", "server2", dnsServerName1, dnsServerName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "server", "bloxone_dns_server.server2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDnsHostExists(ctx context.Context, resourceName string, v *dnsconfig.Host) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			HostAPI.
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

func testAccCheckDnsHostDestroy(ctx context.Context, v *dnsconfig.Host) resource.TestCheckFunc {
	// Verify the server is unassigned from the host
	return func(state *terraform.State) error {
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			HostAPI.
			Read(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		if apiRes.Result.Server == nil {
			// Server object was deleted
			return nil
		}
		return errors.New("server expected to be unassigned from host")
	}
}

func testAccBaseWithDnsConfig(hostName, name string) string {
	return fmt.Sprintf(`
data "bloxone_infra_hosts" "test_host" {
    filters = {
		display_name = %q 
	}
}

resource "bloxone_dns_server" "server" {
    name = %q
}
`, hostName, name)
}

func testAccDnsHostBasicConfig(hostName, dnsServerName string) string {
	config := `
resource "bloxone_dns_host" "test" {
	id = data.bloxone_infra_hosts.test_host.results.0.legacy_id
	server = bloxone_dns_server.server.id
}
`
	return strings.Join([]string{testAccBaseWithDnsConfig(hostName, dnsServerName), config}, "")
}

func testAccDnsHostServer(hostName, dnsServer, dnsServerName1, dnsServerName2 string) string {
	config := fmt.Sprintf(`

resource "bloxone_dns_server" "server2" {
	name = %q
}

resource "bloxone_dns_host" "test_server" {
	id = data.bloxone_infra_hosts.test_host.results.0.legacy_id
	server = bloxone_dns_server.%s.id
}
`, dnsServerName2, dnsServer)
	return strings.Join([]string{testAccBaseWithDnsConfig(hostName, dnsServerName1), config}, "")
}
