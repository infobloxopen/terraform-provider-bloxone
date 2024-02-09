package dns_data_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordAAAAResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_aaaa_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAAAARdata(zoneFqdn, "2001:db8::1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "2001:db8::1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordAAAARdata(zoneFqdn, "2001:db8::2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.address", "2001:db8::2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordAAAARdata(zoneFqdn string, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_aaaa_record" "test_rdata" {
    rdata = {
		"address" = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, address)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
