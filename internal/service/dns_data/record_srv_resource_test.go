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

func TestAccRecordSRVResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_srv_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordSRVRdata(zoneFqdn, 80, 10, "abc.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.port", "80"),
					resource.TestCheckResourceAttr(resourceName, "rdata.priority", "10"),
					resource.TestCheckResourceAttr(resourceName, "rdata.target", "abc.com."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordSRVRdata(zoneFqdn, 90, 20, "xyz.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.port", "90"),
					resource.TestCheckResourceAttr(resourceName, "rdata.priority", "20"),
					resource.TestCheckResourceAttr(resourceName, "rdata.target", "xyz.com."),
				),
			},
			// Update with optional fields and Read
			{
				Config: testAccRecordSRVRdataWithWeight(zoneFqdn, 90, 20, "xyz.com.", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.port", "90"),
					resource.TestCheckResourceAttr(resourceName, "rdata.priority", "20"),
					resource.TestCheckResourceAttr(resourceName, "rdata.target", "xyz.com."),
					resource.TestCheckResourceAttr(resourceName, "rdata.weight", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordSRVRdata(zoneFqdn string, port int, priority int, target string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_srv_record" "test_rdata" {
	name_in_zone = "srv"
    rdata = {
        port = %d
        priority = %d
        target = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, port, priority, target)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordSRVRdataWithWeight(zoneFqdn string, port int, priority int, target string, weight int) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_srv_record" "test_rdata" {
	name_in_zone = "srv"
    rdata = {
        port = %d
        priority = %d
        target = %q
        weight = %d
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, port, priority, target, weight)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
