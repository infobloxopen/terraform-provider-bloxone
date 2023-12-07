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

	resource.Test(t, resource.TestCase{
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

func TestAccRecordAAAADataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_aaaa_records.test"
	resourceName := "bloxone_dns_aaaa_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("aaaa")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordAAAADataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
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

func testAccRecordAAAADataSourceConfigFilters(zoneFqdn, niz string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_aaaa_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    address = "2001:db8::1"
  }
}

data "bloxone_dns_aaaa_records" "test" {
  filters = {
    name_in_zone = %[1]q
    zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_aaaa_record.test]
}`, niz)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
