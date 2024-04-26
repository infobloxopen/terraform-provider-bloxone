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

func TestAccRecordADataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_a_records.test"
	resourceName := "bloxone_dns_a_record.test"
	var v dnsdata.Record
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."
	niz := acctest.RandomNameWithPrefix("a")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordADataSourceConfigFilters(zoneFqdn, niz),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordADataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_a_records.test"
	resourceName := "bloxone_dns_a_record.test"
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com."

	var v dnsdata.Record
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordADataSourceConfigTagFilters(zoneFqdn, "10.0.0.15", acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func testAccCheckRecordResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "absolute_name_spec", dataSourceName, "results.0.absolute_name_spec"),
		resource.TestCheckResourceAttrPair(resourceName, "absolute_zone_name", dataSourceName, "results.0.absolute_zone_name"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "delegation", dataSourceName, "results.0.delegation"),
		resource.TestCheckResourceAttrPair(resourceName, "disabled", dataSourceName, "results.0.disabled"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_absolute_name_spec", dataSourceName, "results.0.dns_absolute_name_spec"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_absolute_zone_name", dataSourceName, "results.0.dns_absolute_zone_name"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_name_in_zone", dataSourceName, "results.0.dns_name_in_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_rdata", dataSourceName, "results.0.dns_rdata"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "inheritance_sources", dataSourceName, "results.0.inheritance_sources"),
		resource.TestCheckResourceAttrPair(resourceName, "ipam_host", dataSourceName, "results.0.ipam_host"),
		resource.TestCheckResourceAttrPair(resourceName, "name_in_zone", dataSourceName, "results.0.name_in_zone"),
		resource.TestCheckResourceAttrPair(resourceName, "options", dataSourceName, "results.0.options"),
		resource.TestCheckResourceAttrPair(resourceName, "provider_metadata", dataSourceName, "results.0.provider_metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "rdata", dataSourceName, "results.0.rdata"),
		resource.TestCheckResourceAttrPair(resourceName, "source", dataSourceName, "results.0.source"),
		resource.TestCheckResourceAttrPair(resourceName, "subtype", dataSourceName, "results.0.subtype"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "ttl", dataSourceName, "results.0.ttl"),
		resource.TestCheckResourceAttrPair(resourceName, "type", dataSourceName, "results.0.type"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
		resource.TestCheckResourceAttrPair(resourceName, "view", dataSourceName, "results.0.view"),
		resource.TestCheckResourceAttrPair(resourceName, "view_name", dataSourceName, "results.0.view_name"),
		resource.TestCheckResourceAttrPair(resourceName, "zone", dataSourceName, "results.0.zone"),
	}
}

func testAccRecordADataSourceConfigFilters(zoneFqdn, niz string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test" {
  name_in_zone = %[1]q
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    address = "10.0.0.15"
  }
}

data "bloxone_dns_a_records" "test" {
  filters = {
    name_in_zone = %[1]q
    zone = bloxone_dns_auth_zone.test.id
  }
  depends_on = [bloxone_dns_a_record.test]
}`, niz)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}

func testAccRecordADataSourceConfigTagFilters(zoneFqdn, address string, tagValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_a_record" "test" {
  zone = bloxone_dns_auth_zone.test.id
  rdata = {
    address = %[1]q
  }
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_a_records" "test" {
  tag_filters = {
	tag1 = bloxone_dns_a_record.test.tags.tag1
  }
}
`, address, tagValue)
	return strings.Join([]string{config, testAccBaseWithZone(zoneFqdn)}, "")
}
