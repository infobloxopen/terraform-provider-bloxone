package fw_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccApplicationFiltersResource_basic(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersBasicConfig(name, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_time"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_time"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "readonly", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationFiltersResource_disappears(t *testing.T) {
	resourceName := "bloxone_td_application_filter.test"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationFiltersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationFiltersBasicConfig(name, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					testAccCheckApplicationFiltersDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccApplicationFiltersResource_CriteriaCategory(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test_criteria"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersCriteriaCategory(name, "Email"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.category", "Email"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationFiltersCriteriaCategory(name, "Communication"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.category", "Communication"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationFiltersResource_CriteriaName(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test_criteria"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersCriteriaName(name, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.name", "Microsoft 365"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationFiltersCriteriaName(name, "163 Cloud"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "criteria.0.name", "163 Cloud"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationFiltersResource_Description(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test_description"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersDescription(name, "Microsoft 365", "Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test Description"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationFiltersDescription(name, "Microsoft 365", "Updated Test Description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated Test Description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationFiltersResource_Name(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test_name"
	var v fw.ApplicationFilter
	name1 := acctest.RandomNameWithPrefix("app-filter")
	name2 := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersName(name1, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationFiltersName(name2, "Microsoft 365"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationFiltersResource_Tags(t *testing.T) {
	var resourceName = "bloxone_td_application_filter.test_tags"
	var v fw.ApplicationFilter
	name := acctest.RandomNameWithPrefix("app-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationFiltersTags(name, "Microsoft 365", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationFiltersTags(name, "Microsoft 365", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationFiltersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckApplicationFiltersExists(ctx context.Context, resourceName string, v *fw.ApplicationFilter) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.FWAPI.
			ApplicationFiltersAPI.
			ReadApplicationFilter(ctx, int32(id)).
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

func testAccCheckApplicationFiltersDestroy(ctx context.Context, v *fw.ApplicationFilter) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.FWAPI.
			ApplicationFiltersAPI.
			ReadApplicationFilter(ctx, *v.Id).
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

func testAccCheckApplicationFiltersDisappears(ctx context.Context, v *fw.ApplicationFilter) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.FWAPI.
			ApplicationFiltersAPI.
			DeleteSingleApplicationFilters(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccApplicationFiltersBasicConfig(name, criteriaName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
}
`, name, criteriaName)
}

func testAccApplicationFiltersCriteriaCategory(name, criteriaName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test_criteria" {
	name = %q
	criteria  = [
	{
		category = %q
	}
]
}
`, name, criteriaName)
}

func testAccApplicationFiltersCriteriaName(name, criteriaName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test_criteria" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
}
`, name, criteriaName)
}

func testAccApplicationFiltersDescription(name, criteriaName, description string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test_description" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
	description = %q
}
`, name, criteriaName, description)
}

func testAccApplicationFiltersName(name, criteriaName string) string {
	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test_name" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
}
`, name, criteriaName)
}

func testAccApplicationFiltersTags(name, criteriaName string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_td_application_filter" "test_tags" {
	name = %q
	criteria  = [
	{
		name = %q
	}
]
	tags = %s
}
`, name, criteriaName, tagsStr)
}
