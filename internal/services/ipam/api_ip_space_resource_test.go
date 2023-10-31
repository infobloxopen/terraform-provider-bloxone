package ipam_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccIpSpaceResource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config:             testIpSpaceConfig("one"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bloxone_ipam_ip_space.test", "name", "one"),
				),
			},
			// Update and Read testing
			{
				Config:             testIpSpaceConfig("two"),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("bloxone_ipam_ip_space.test", "name", "two"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testIpSpaceConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" "test" {
	name = "%s"
}
`, name)
}
