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

func TestAccRecordHTTPSResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_https_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordHTTPSRdata(zoneFqdn, "google.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.target_name", "google.com."),
				),
			},
			// Update and Read
			{
				Config: testAccRecordHTTPSRdata(zoneFqdn, "apple.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.target_name", "apple.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordHTTPSDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_https_records.test"
	resourceName := "bloxone_dns_https_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("https")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordHTTPSDataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccRecordHTTPSRdata(zoneFqdn string, https string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_https_record" "test_rdata" {
	name_in_zone = "https"
    rdata = {
		target_name = %q
	}
    zone = bloxone_dns_auth_zone.test.id
}
`, https)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordHTTPSDataSourceConfigFilters(zoneFqdn, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_https_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    target_name = "example.com."
  }
}

data "bloxone_dns_https_records" "test" {
  filters = {
    name_in_zone = %[1]q
	zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_https_record.test]
}`, nameInZone)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
