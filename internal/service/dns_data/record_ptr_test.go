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

func TestAccRecordPTRResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_ptr_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordPTRRdata(zoneFqdn, "google.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "google.com."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordPTRRdata(zoneFqdn, "apple.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.dname", "apple.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordPTRDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_ptr_records.test"
	resourceName := "bloxone_dns_ptr_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("ptr")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordPTRDataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccRecordPTRRdata(zoneFqdn string, dname string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_ptr_record" "test_rdata" {
	name_in_zone = "ptr"
    rdata = {
		dname = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, dname)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordPTRDataSourceConfigFilters(zoneFqdn, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_ptr_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    dname = "example.com."
  }
}

data "bloxone_dns_ptr_records" "test" {
  filters = {
    name_in_zone = %[1]q
	zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_ptr_record.test]
}`, nameInZone)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
