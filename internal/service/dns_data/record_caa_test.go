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

func TestAccRecordCAAResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_caa_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordCAARdata(zoneFqdn, "issue", "ca.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.tag", "issue"),
					resource.TestCheckResourceAttr(resourceName, "rdata.value", "ca.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccRecordCAARdata(zoneFqdn, "issuewild", "*.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.tag", "issuewild"),
					resource.TestCheckResourceAttr(resourceName, "rdata.value", "*.example.com"),
				),
			},
			// Update with optional fields and Read
			{
				Config: testAccRecordCAARdataWithFlags(zoneFqdn, 1, "issuewild", "*.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.flags", "1"),
					resource.TestCheckResourceAttr(resourceName, "rdata.tag", "issuewild"),
					resource.TestCheckResourceAttr(resourceName, "rdata.value", "*.example.com"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordCAADataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_caa_records.test"
	resourceName := "bloxone_dns_caa_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("caa")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordCAADataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccRecordCAARdataWithFlags(zoneFqdn string, flags int, tag string, value string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_caa_record" "test_rdata" {
    rdata = {
		flags = %d
        tag = %q
        value = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, flags, tag, value)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCAARdata(zoneFqdn string, tag string, value string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_caa_record" "test_rdata" {
    rdata = {
        tag = %q
        value = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, tag, value)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordCAADataSourceConfigFilters(zoneFqdn, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_caa_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
    rdata = {
        tag = "issue"
        value = "ca.example.com"
	}
}

data "bloxone_dns_caa_records" "test" {
  filters = {
    name_in_zone = %[1]q
    zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_caa_record.test]
}`, nameInZone)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
