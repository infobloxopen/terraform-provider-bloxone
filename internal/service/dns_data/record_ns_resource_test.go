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

func TestAccRecordNSResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_ns_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNSRdata(zoneFqdn, "ns1.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "ns1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNSRdata(zoneFqdn, "ns2.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "ns2.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordNSRdata(zoneFqdn string, dname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_ns_record" "test_rdata" {
	name_in_zone = "ns"
    rdata = {
		dname = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, dname)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
