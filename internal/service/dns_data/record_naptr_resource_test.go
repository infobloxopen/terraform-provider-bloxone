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

func TestAccRecordNAPTRResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_naptr_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNAPTRRdata(zoneFqdn, 10, 10, "+", "SIP+D2U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.order", "10"),
					resource.TestCheckResourceAttr(resourceName, "rdata.preference", "10"),
					resource.TestCheckResourceAttr(resourceName, "rdata.replacement", "+"),
					resource.TestCheckResourceAttr(resourceName, "rdata.services", "SIP+D2U"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNAPTRRdata(zoneFqdn, 20, 20, ".", "SIP+E2U"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.order", "20"),
					resource.TestCheckResourceAttr(resourceName, "rdata.preference", "20"),
					resource.TestCheckResourceAttr(resourceName, "rdata.replacement", "."),
					resource.TestCheckResourceAttr(resourceName, "rdata.services", "SIP+E2U"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordNAPTRResource_Rdata_FlagsAndRegexp(t *testing.T) {
	var resourceName = "bloxone_dns_naptr_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordNAPTRRdataWithFlagsAndRegex(zoneFqdn, "U", "!^.*$!sip:jdoe@corpxyz.com!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.flags", "U"),
					resource.TestCheckResourceAttr(resourceName, "rdata.regexp", "!^.*$!sip:jdoe@corpxyz.com!"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordNAPTRRdataWithFlagsAndRegex(zoneFqdn, "A", "!^.*$!sip:jdoe@corpabc.com!"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.flags", "A"),
					resource.TestCheckResourceAttr(resourceName, "rdata.regexp", "!^.*$!sip:jdoe@corpabc.com!"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccRecordNAPTRRdata(zoneFqdn string, order int, preference int, replacement string, services string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_naptr_record" "test_rdata" {
	name_in_zone = "naptr"
    rdata = {
        order = %d
        preference = %d
        replacement = %q
        services = %q
    }
    zone = bloxone_dns_auth_zone.test.id
}
`, order, preference, replacement, services)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordNAPTRRdataWithFlagsAndRegex(zoneFqdn string, flags string, regexp string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_naptr_record" "test_rdata" {
	name_in_zone = "naptr"
    rdata = {
        flags = %q
        order = 100
        preference = 10
        regexp = %q
        replacement = "."
        services = "SIP+D2U"
    }
    zone = bloxone_dns_auth_zone.test.id
}
`, flags, regexp)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}
