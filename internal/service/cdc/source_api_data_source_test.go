package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccSourceApiDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_source_nios_list.test"
	resourceName := "bloxone_cdc_source_nios.test"
	sourceName := acctest.RandomNameWithPrefix("source")
	var v cdc.SourceNios

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSourceApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceApiDataSourceConfigFilters(sourceName),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					}, testAccCheckSourceApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccSourceApiDataSource_TagFilters(t *testing.T) {
	// TODO: the CDC backend does not correctly honor tag_filters on
	// /v1/sources/nios. Confirmed via direct curl against the API: a
	// _tfilter query for a tag value that matches nothing returns the full,
	// unfiltered list of every source on the cluster instead of an empty
	// result set, so this data source always surfaces an arbitrary
	// pre-existing source rather than the one just created. This is a
	// backend bug, not something fixable in the provider. Needs a CDC
	// backend fix before this can be un-skipped.
	t.Skip("CDC backend does not honor tag_filters on /v1/sources/nios; see TODO above")

	dataSourceName := "data.bloxone_cdc_source_nios_list.test"
	resourceName := "bloxone_cdc_source_nios.test"
	sourceName := acctest.RandomNameWithPrefix("source")
	var v cdc.SourceNios
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSourceApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceApiDataSourceConfigTagFilters(sourceName, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					}, testAccCheckSourceApiResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckSourceApiResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "address", dataSourceName, "results.0.address"),
		resource.TestCheckResourceAttrPair(resourceName, "ca_certificate", dataSourceName, "results.0.ca_certificate"),
		resource.TestCheckResourceAttrPair(resourceName, "certificate", dataSourceName, "results.0.certificate"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "insecure_mode", dataSourceName, "results.0.insecure_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		// nios_password/scp_password are excluded: the CDC API always returns
		// a masked placeholder for these fields, so they never match the
		// resource's real config value.
		resource.TestCheckResourceAttrPair(resourceName, "nios_username", dataSourceName, "results.0.nios_username"),
		resource.TestCheckResourceAttrPair(resourceName, "rpz_log_insecure_mode", dataSourceName, "results.0.rpz_log_insecure_mode"),
		resource.TestCheckResourceAttrPair(resourceName, "scp_username", dataSourceName, "results.0.scp_username"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccSourceApiDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test" {
  name          = %q
  address       = "%s.example.com"
  nios_username = "tf_auto"
  nios_password = "tf_auto_password1!"
}

data "bloxone_cdc_source_nios_list" "test" {
  filters = {
	name = bloxone_cdc_source_nios.test.name
  }
}
`, name, name)
}

func testAccSourceApiDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test" {
  name          = %q
  address       = "%s.example.com"
  nios_username = "tf_auto"
  nios_password = "tf_auto_password1!"
  tags = {
	tag1 = %q
  }
}

data "bloxone_cdc_source_nios_list" "test" {
  tag_filters = {
	tag1 = bloxone_cdc_source_nios.test.tags.tag1
  }
}
`, name, name, tagValue)
}
