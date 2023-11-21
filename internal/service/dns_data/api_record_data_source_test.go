package dns_data_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dns_data"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccRecordDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_records.test"
	resourceName := "bloxone_dns_record.test"
	var v dns_data.DataRecord

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordDataSourceConfigFilters("RDATA_REPLACE_ME", "TYPE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccRecordDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dns_records.test"
	resourceName := "bloxone_dns_record.test"
	var v dns_data.DataRecord
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckRecordDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccRecordDataSourceConfigTagFilters("RDATA_REPLACE_ME", "TYPE_REPLACE_ME", "value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckRecordExists(context.Background(), resourceName, &v),
					}, testAccCheckRecordResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

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

func testAccRecordDataSourceConfigFilters(rdata, type_ string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test" {
  rdata = %q
  type = %q
}

data "bloxone_dns_records" "test" {
  filters = {
	rdata = bloxone_dns_record.test.rdata
  }
}
`, rdata, type_)
}

func testAccRecordDataSourceConfigTagFilters(rdata, type_ string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_record" "test" {
  rdata = %q
  type = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_dns_records" "test" {
  tag_filters = {
	tag1 = bloxone_dns_record.test.tags.tag1
  }
}
`, rdata, type_, tagValue)
}
