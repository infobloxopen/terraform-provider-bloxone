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

func TestAccRecordSVCBResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_svcb_record.test_rdata"
	var v dnsdata.Record
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSVCBRdata(zoneFqdn, "google.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.target_name", "google.com."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSVCBRdata(zoneFqdn, "apple.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.target_name", "apple.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordSVCBRdata(zoneFqdn string, svcb string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_svcb_record" "test_rdata" {
	name_in_zone = "svcb"
    rdata = {
		target_name = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, svcb)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
