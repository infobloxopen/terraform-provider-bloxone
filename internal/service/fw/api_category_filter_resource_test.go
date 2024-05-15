package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccCategoryFiltersResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_category_filter.test"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCategoryFiltersBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCategoryFiltersResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_category_filter.test"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCategoryFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccCategoryFiltersBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					testAccCheckCategoryFiltersDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccCategoryFiltersResource_Categories(t *testing.T) {
	var resourceName = "bloxone_td_category_filter.test_categories"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCategoryFiltersCategories(name, "1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "categories.0", "data.bloxone_td_content_categories.test", "results.1.category_name"),
				),
			},
			// Update and Read
			{
				Config: testAccCategoryFiltersCategories(name, "2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "categories.0", "data.bloxone_td_content_categories.test", "results.2.category_name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCategoryFiltersResource_Description(t *testing.T) {
	var resourceName = "bloxone_td_category_filter.test_description"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCategoryFiltersDescription(name, "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccCategoryFiltersDescription(name, "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCategoryFiltersResource_Name(t *testing.T) {
	var resourceName = "bloxone_td_category_filter.test_name"
	var v fw.CategoryFilter
	name1 := acctest.RandomNameWithPrefix("category-filter")
	name2 := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCategoryFiltersName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccCategoryFiltersName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccCategoryFiltersResource_Tags(t *testing.T) {
	var resourceName = "bloxone_td_category_filter.test_tags"
	var v fw.CategoryFilter
	name := acctest.RandomNameWithPrefix("category-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccCategoryFiltersTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccCategoryFiltersTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCategoryFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckCategoryFiltersExists(ctx context.Context, resourceName string, v *fw.CategoryFilter) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.Atoi(rs.Primary.ID)
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			CategoryFiltersAPI.
			ReadCategoryFilter(ctx, int32(id)).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResults()
		return nil
	}
}

func testAccCheckCategoryFiltersDestroy(ctx context.Context, v *fw.CategoryFilter) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			CategoryFiltersAPI.
			ReadCategoryFilter(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckCategoryFiltersDisappears(ctx context.Context, v *fw.CategoryFilter) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			CategoryFiltersAPI.
			DeleteSingleCategoryFilters(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithContentCategories() string {
	return fmt.Sprintf(`
data "bloxone_td_content_categories" "test" {
}
`)
}

func testAccCategoryFiltersBasicConfig(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
}
`, name)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}

func testAccCategoryFiltersCategories(name, category string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test_categories" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.%s.category_name]
}
`, name, category)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}

func testAccCategoryFiltersDescription(name, description string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test_description" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
	description = %q
}
`, name, description)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}

func testAccCategoryFiltersName(name string) string {
	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test_name" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
}
`, name)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}

func testAccCategoryFiltersTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_td_category_filter" "test_tags" {
	name = %q
	categories = [data.bloxone_td_content_categories.test.results.0.category_name]
	tags = %s
}
`, name, tagsStr)
	return strings.Join([]string{testAccBaseWithContentCategories(), config}, "")
}
