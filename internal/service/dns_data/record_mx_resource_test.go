package dns_data_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dnsdata"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordMXResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_mx_record.test_rdata"
	var v dnsdata.Record
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordMXRdata(zoneFqdn, "m1.example.com", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.exchange", "m1.example.com"),
					resource.TestCheckResourceAttr(resourceName, "rdata.preference", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordMXRdata(zoneFqdn, "m2.example.com", 20),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.exchange", "m2.example.com"),
					resource.TestCheckResourceAttr(resourceName, "rdata.preference", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordMXRdata(zoneFqdn string, exchange string, preference int) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_mx_record" "test_rdata" {
	name_in_zone = "mx"
    rdata = {
       exchange = %q
       preference = %d
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, exchange, preference)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
