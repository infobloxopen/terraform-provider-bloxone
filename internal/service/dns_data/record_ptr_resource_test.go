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

func TestAccRecordPTRResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_ptr_record.test_rdata"
	var v dnsdata.Record
	view := acctest.RandomNameWithPrefix("view")
	zoneFqdn := "10.in-addr.arpa."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPTRRdata(view, zoneFqdn, "domain.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "domain.com."),
					resource.TestCheckResourceAttr(resourceName, `options.address`, "10.0.0.1"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPTRRdata(view, zoneFqdn, "apple.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "apple.com."),
					resource.TestCheckResourceAttr(resourceName, `options.address`, "10.0.0.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPTRResource_Options(t *testing.T) {
	var resourceName = "bloxone_dns_ptr_record.test_options"
	var v dnsdata.Record
	view := acctest.RandomNameWithPrefix("view")
	zoneFqdn := "10.in-addr.arpa."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPtrOptions(view, zoneFqdn, "domain.com.", "10.0.0.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, `options.address`, "10.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "1.0.0"),
					resource.TestCheckResourceAttrPair(resourceName, "zone", "bloxone_dns_auth_zone.test", "id"),
				),
			},
			{
				Config: testAccRecordPtrOptions(view, zoneFqdn, "domain.com.", "10.0.0.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, `options.address`, "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "name_in_zone", "2.0.0"),
					resource.TestCheckResourceAttrPair(resourceName, "zone", "bloxone_dns_auth_zone.test", "id"),
				),
			},
		},
	})
}

func testAccRecordPTRRdata(view, zoneFqdn string, dname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_ptr_record" "test_rdata" {
	name_in_zone = "1.0.0"
    rdata = {
		dname = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, dname)
	return strings.Join([]string{testAccBaseWithZoneAndView(view, zoneFqdn), config}, "")
}

func testAccRecordPtrOptions(view, zoneFqdn string, dname string, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_ptr_record" "test_options" {
    rdata = {
		dname = %[1]q
	}
	options = {
		address = %[2]q
	}
    view = bloxone_dns_auth_zone.test.view
	depends_on = [bloxone_dns_auth_zone.test]
}
`, dname, address)
	return strings.Join([]string{testAccBaseWithZoneAndView(view, zoneFqdn), config}, "")
}
