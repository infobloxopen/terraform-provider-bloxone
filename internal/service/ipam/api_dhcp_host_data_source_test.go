package ipam_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDhcpHostDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_hosts.test"
	var v ipam.IpamsvcHost

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpHostDataSourceConfigFilters("TF_TEST_HOST_01"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpHostExists(context.Background(), dataSourceName, &v),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "TF_TEST_HOST_01"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
				),
			},
		},
	})
}

func TestAccDhcpHostDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dhcp_hosts.test_tag"
	var v ipam.IpamsvcHost

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDhcpHostDataSourceConfigTagFilters("Terraform Acceptance Testing"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDhcpHostExists(context.Background(), dataSourceName, &v),
					resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.name", "TF_TEST_HOST_01"),
					resource.TestCheckResourceAttr(dataSourceName, "results.0.tags.used_for", "Terraform Acceptance Testing"),
					resource.TestCheckResourceAttrSet(dataSourceName, "results.0.id"),
				),
			},
		},
	})
}

func testAccCheckDhcpHostExists(ctx context.Context, dataSourceName string, v *ipam.IpamsvcHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("not found: %s", dataSourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			DhcpHostAPI.
			DhcpHostRead(ctx, rs.Primary.Attributes["results.0.id"]).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", dataSourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

// below all TestAcc functions

func testAccDhcpHostDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`

data "bloxone_dhcp_hosts" "test" {
  filters = {
	 name = %q
  }
}
`, name)
}

func testAccDhcpHostDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
data "bloxone_dhcp_hosts" "test_tag" {
  tag_filters = {
	used_for = %q
  }
}
`, tagValue)
}
