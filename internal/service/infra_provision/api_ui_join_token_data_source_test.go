package infra_provision_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/infraprovision"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccUIJoinTokenDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_join_tokens.test"
	resourceName := "bloxone_infra_join_token.test"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUIJoinTokenDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUIJoinTokenDataSourceConfigFilters(name),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
						resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					}, testAccCheckUIJoinTokenResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccUIJoinTokenDataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_infra_join_tokens.test"
	resourceName := "bloxone_infra_join_token.test"
	var v infraprovision.JoinToken
	name := acctest.RandomNameWithPrefix("jt")
	tagValue := acctest.RandomNameWithPrefix("tag ")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckUIJoinTokenDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccUIJoinTokenDataSourceConfigTagFilters(name, tagValue),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckUIJoinTokenExists(context.Background(), resourceName, &v),
						resource.TestCheckResourceAttr(dataSourceName, "results.#", "1"),
					}, testAccCheckUIJoinTokenResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckUIJoinTokenResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "deleted_at", dataSourceName, "results.0.deleted_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "expires_at", dataSourceName, "results.0.expires_at"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "last_used_at", dataSourceName, "results.0.last_used_at"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "status", dataSourceName, "results.0.status"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "token_id", dataSourceName, "results.0.token_id"),
		resource.TestCheckResourceAttrPair(resourceName, "use_counter", dataSourceName, "results.0.use_counter"),
	}
}

func testAccUIJoinTokenDataSourceConfigFilters(name string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test" {
  name = %q
}

data "bloxone_infra_join_tokens" "test" {
  filters = {
	name = bloxone_infra_join_token.test.name
  }
}
`, name)
}

func testAccUIJoinTokenDataSourceConfigTagFilters(name string, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_join_token" "test" {
  name = %q
  tags = {
	tag1 = %q
  }
}

data "bloxone_infra_join_tokens" "test" {
  tag_filters = {
	tag1 = bloxone_infra_join_token.test.tags.tag1
  }
}
`, name, tagValue)
}
