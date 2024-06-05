package keys_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccTsigDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_keys_tsigs.test"
	resourceName := "bloxone_keys_tsig.test"
	var v keys.TSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTsigDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTsigDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckTsigExists(context.Background(), resourceName, &v),
					}, testAccCheckTsigResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccTsigDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_keys_tsigs.test"
	resourceName := "bloxone_keys_tsig.test"
	var v keys.TSIGKey
	name := acctest.RandomNameWithPrefix("key") + "."

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckTsigDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccTsigDataSourceConfigTagFilters(name, acctest.RandomName()),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckTsigExists(context.Background(), resourceName, &v),
					}, testAccCheckTsigResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckTsigResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "algorithm", dataSourceName, "results.0.algorithm"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "results.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol_name", dataSourceName, "results.0.protocol_name"),
		resource.TestCheckResourceAttrPair(resourceName, "secret", dataSourceName, "results.0.secret"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccTsigDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test" {
  name = %q
}

data "bloxone_keys_tsigs" "test" {
  filters = {
	name = bloxone_keys_tsig.test.name
  }
}
`, name)
}

func testAccTsigDataSourceConfigTagFilters(name, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_keys_tsigs" "test" {
  tag_filters = {
	tag1 = bloxone_keys_tsig.test.tags.tag1
  }
}
`, name, tagValue)
}
