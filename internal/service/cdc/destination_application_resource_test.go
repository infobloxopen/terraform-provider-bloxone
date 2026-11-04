package cdc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationApplicationResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrPair(resourceName, "application", "bloxone_cdc_application.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "STREAM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_application.test"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationApplicationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationApplicationBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					testAccCheckDestinationApplicationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationApplicationResource_Application(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_application"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationApplication(name, "test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "application", "bloxone_cdc_application.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationApplication(name, "test2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "application", "bloxone_cdc_application.test2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_Context(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_context"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationContext(name, map[string]string{"var1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "context.var1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationContext(name, map[string]string{"var1": "value1changed", "var2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "context.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "context.var1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "context.var2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_description"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Removing description from config resets it to the default
			{
				Config: testAccDestinationApplicationDescriptionRemoved(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_enabled"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_name"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationApplicationResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_application.test_tags"
	var v cdc.DestinationApplication
	name := acctest.RandomNameWithPrefix("destination-application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationApplicationTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationApplicationTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationApplicationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationApplicationExists(ctx context.Context, resourceName string, v *cdc.DestinationApplication) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationApplication(ctx, id).
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

func testAccCheckDestinationApplicationDestroy(ctx context.Context, v *cdc.DestinationApplication) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationApplication(ctx, *v.Id).
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

func testAccCheckDestinationApplicationDisappears(ctx context.Context, v *cdc.DestinationApplication) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationApplication(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDestinationApplicationBaseConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test" {
  name = "%[1]s-app"
  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
}

resource "bloxone_cdc_application" "test2" {
  name = "%[1]s-app2"
  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
}
`, name)
}

func testAccDestinationApplicationApplication(name string, application string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_application" {
  name        = %q
  application = bloxone_cdc_application.%s.id
}
`, name, application)
}

func testAccDestinationApplicationBasicConfig(name string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test" {
  name        = %q
  application = bloxone_cdc_application.test.id
}
`, name)
}

func testAccDestinationApplicationContext(name string, contextVars map[string]string) string {
	contextStr := "{\n"
	for k, v := range contextVars {
		contextStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	contextStr += "  }"

	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_context" {
  name        = %q
  application = bloxone_cdc_application.test.id
  context     = %s
}
`, name, contextStr)
}

func testAccDestinationApplicationDescription(name string, description string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_description" {
  name        = %q
  application = bloxone_cdc_application.test.id
  description = %q
}
`, name, description)
}

func testAccDestinationApplicationDescriptionRemoved(name string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_description" {
  name        = %q
  application = bloxone_cdc_application.test.id
}
`, name)
}

func testAccDestinationApplicationEnabled(name string, enabled string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_enabled" {
  name        = %q
  application = bloxone_cdc_application.test.id
  enabled     = %s
}
`, name, enabled)
}

func testAccDestinationApplicationName(name string) string {
	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_name" {
  name        = %q
  application = bloxone_cdc_application.test.id
}
`, name)
}

func testAccDestinationApplicationTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return testAccDestinationApplicationBaseConfig(name) + fmt.Sprintf(`
resource "bloxone_cdc_destination_application" "test_tags" {
  name        = %q
  application = bloxone_cdc_application.test.id
  tags        = %s
}
`, name, tagsStr)
}
