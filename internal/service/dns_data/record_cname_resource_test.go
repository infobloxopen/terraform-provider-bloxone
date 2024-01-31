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

func TestAccRecordCNAMEResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_cname_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCNAMERdata(zoneFqdn, "c1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.cname", "c1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCNAMERdata(zoneFqdn, "c2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.cname", "c2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordCNAMERdata(zoneFqdn string, cname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_cname_record" "test_rdata" {
	name_in_zone = "cname"
    rdata = {
		cname = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, cname)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
