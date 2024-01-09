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

func TestAccRecordTXTResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_txt_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordTXTRdata(zoneFqdn, "abc"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.text", "abc"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordTXTRdata(zoneFqdn, "xyz"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.text", "xyz"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordTXTRdata(zoneFqdn string, txt string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_txt_record" "test_rdata" {
	name_in_zone = "txt"
    rdata = {
		text = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, txt)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
