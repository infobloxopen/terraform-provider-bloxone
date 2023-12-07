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

func TestAccRecordGenericResource_Rdata(t *testing.T) {
	var resourceName = "bloxone_dns_record.test_rdata"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccRecordGenericRdataPresentation(zoneFqdn, "TYPE256", "10 1 \"https://example.com\""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRecordExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.0.type", "PRESENTATION"),
					resource.TestCheckResourceAttr(resourceName, "rdata.subfields.0.value", "10 1 \"https://example.com\""),
					resource.TestCheckResourceAttr(resourceName, "type", "TYPE256"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccRecordGenericDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_records.test"
	resourceName := "bloxone_dns_record.test"
	var v dns_data.DataRecord
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("naptr")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordGenericDataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccRecordGenericRdataPresentation(zoneFqdn string, flags string, regexp string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_record" "test_rdata" {
	name_in_zone = "naptr"
    type = %q
    rdata        = {
      subfields = [
        {
          type  = "PRESENTATION"
          value = %q
        }
      ]
    }
    zone = bloxone_dns_auth_zone.test.id
}
`, flags, regexp)
	return strings.Join([]string{testAccBaseWithZone(zoneFqdn), config}, "")
}

func testAccRecordGenericDataSourceConfigFilters(zoneFqdn, nameInZone string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_record" "test" {
  type = "TYPE256"
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata        = {
    subfields = [
      {
        type  = "PRESENTATION"
        value = "10 1 \"https://example.com\""
      }
    ]
  }
}

data "bloxone_dns_records" "test" {
  filters = {
    name_in_zone = %[1]q
	zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_record.test]
}`, nameInZone)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
