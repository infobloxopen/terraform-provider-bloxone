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

	resource.ParallelTest(t, resource.TestCase{
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

func TestAccRecordAAAAResource_Options(t *testing.T) {
	var resourceName = "bloxone_dns_aaaa_record.test_options"
	var datasourceName = "data.bloxone_dns_ptr_records.test"
	var v1 dns_data.DataRecord
	var v2 dns_data.DataRecord
	viewName := acctest.RandomNameWithPrefix("view")
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordAAAAOptions(viewName, zoneFqdn, "2001:db8::1", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "options.create_ptr", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.check_rmz", "true"),
					resource.TestCheckResourceAttr(datasourceName, "results.0.name_in_zone", "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0"),
				),
			},
			{
				Config: testAccRecordAAAAOptions(viewName, zoneFqdn, "2001:db8::1", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "options.create_ptr", "true"),
					resource.TestCheckResourceAttr(resourceName, "options.check_rmz", "false"),
					resource.TestCheckResourceAttr(datasourceName, "results.0.name_in_zone", "1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.8.b.d.0"),
				),
			},
			{
				// The value of create_ptr has changed from `true` to false, so the resource will be recreated.
				Config: testAccRecordAAAAOptions(viewName, zoneFqdn, "2001:db8::1", false, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordDestroy(context.Background(), &v1),
					testAccCheckRecordExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "options.create_ptr", "false"),
					resource.TestCheckResourceAttr(resourceName, "options.check_rmz", "false"),
					resource.TestCheckResourceAttr(datasourceName, "results.#", "0"),
				),
			},
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

func testAccRecordAAAAOptions(view string, zoneFqdn string, address string, createPtr, checkRmz bool) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_aaaa_record" "test_options" {
    rdata = {
		"address" = %q
	}
	options = {
		"create_ptr" = %t
		"check_rmz" = %t
	}
	zone = bloxone_dns_auth_zone.test.id
	depends_on = [bloxone_dns_auth_zone.rmz, bloxone_dns_auth_zone.test]
}

data "bloxone_dns_ptr_records" "test" {
	filters = {
		zone = bloxone_dns_auth_zone.rmz.id
	}   
	depends_on = [bloxone_dns_aaaa_record.test_options]
}

resource "bloxone_dns_auth_zone" "rmz" {
	fqdn = "1.0.0.2.ip6.arpa."
	primary_type = "cloud"
	view = bloxone_dns_view.test.id
}
`, address, createPtr, checkRmz)
	return strings.Join([]string{testAccBaseWithZoneAndView(view, zoneFqdn), config}, "")
}
