package dfp_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDfpDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_dfp_services.test"
	resourceName := "bloxone_dfp_service.test"
	var v dfp.Dfp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDfpDataSourceConfigFilters(),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDfpExists(context.Background(), resourceName, &v),
					}, testAccCheckDfpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccDfpDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_dfp_services.test"
	resourceName := "bloxone_dfp_service.test"
	var v dfp.Dfp

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDfpDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDfpDataSourceConfigTagFilters("value1"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckDfpExists(context.Background(), resourceName, &v),
					}, testAccCheckDfpResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckDfpResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{}
}

func testAccDfpDataSourceConfigFilters() string {
	return fmt.Sprintf(`
resource "bloxone_dfp_service" "test" {
}

data "bloxone_dfp_services" "test" {
  filters = {
	 = bloxone_dfp_service.test.
  }
}
`)
}

func testAccDfpDataSourceConfigTagFilters(tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_dfp_service" "test" {
  tags = {
	tag1 = %q
  }
}

data "bloxone_dfp_services" "test" {
  tag_filters = {
	tag1 = bloxone_dfp_service.test.tags.tag1
  }
}
`, tagValue)
}
