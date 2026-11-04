package cdc_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccApplicationApiResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "file_data"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "type"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_application.test"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckApplicationApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccApplicationApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					testAccCheckApplicationApiDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccApplicationApiResource_Data(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_data"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiData(name, `print("automation")`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data", testAccApplicationApiDataJSON("script_1.py", `print("automation")`)),
					testAccCheckApplicationApiFileData(resourceName, "script_1.py", `print("automation")`),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiData(name, `print("automation updated")`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data", testAccApplicationApiDataJSON("script_1.py", `print("automation updated")`)),
					testAccCheckApplicationApiFileData(resourceName, "script_1.py", `print("automation updated")`),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_DataFromFile(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_data_file"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")
	scriptPath := acctest.TestdataPath(t, "bloxone_cdc", "cdc_application_handler_sample.py")
	scriptContent := acctest.ReadTestdata(t, "bloxone_cdc", "cdc_application_handler_sample.py")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiDataFromFile(name, scriptPath),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data", testAccApplicationApiDataJSON("cdc_application_handler_sample.py", scriptContent)),
					testAccCheckApplicationApiFileData(resourceName, "cdc_application_handler_sample.py", scriptContent),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_description"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiDescription(name, "Test application description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test application description"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiDescription(name, "Updated test application description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test application description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_enabled"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_name"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_tags"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccApplicationApiResource_Variables(t *testing.T) {
	var resourceName = "bloxone_cdc_application.test_variables"
	var v cdc.ApplicationData
	name := acctest.RandomNameWithPrefix("application")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccApplicationApiVariables(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "variables.var1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccApplicationApiVariables(name, "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckApplicationApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "variables.var1", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckApplicationApiExists(ctx context.Context, resourceName string, v *cdc.ApplicationData) resource.TestCheckFunc {
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
			ApplicationApiAPI.
			ReadApplicationData(ctx, id).
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

// testAccApplicationApiDataJSON returns the value jsonencode() produces for a
// single-script data object; Terraform's jsonencode sorts keys the same way
// encoding/json does.
func testAccApplicationApiDataJSON(filename, content string) string {
	b, _ := json.Marshal(map[string]interface{}{
		"scripts": []map[string]string{{"content": content, "filename": filename}},
	})
	return string(b)
}

// testAccCheckApplicationApiFileData verifies that the backend response stored
// in file_data contains the given script.
func testAccCheckApplicationApiFileData(resourceName, filename, content string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttrWith(resourceName, "file_data", func(value string) error {
		var data struct {
			Scripts []struct {
				Content  string `json:"content"`
				Filename string `json:"filename"`
			} `json:"scripts"`
		}
		if err := json.Unmarshal([]byte(value), &data); err != nil {
			return fmt.Errorf("file_data is not valid JSON: %w", err)
		}
		if len(data.Scripts) != 1 {
			return fmt.Errorf("expected 1 script in file_data, got %d", len(data.Scripts))
		}
		if data.Scripts[0].Filename != filename || data.Scripts[0].Content != content {
			return fmt.Errorf("unexpected script in file_data: %+v", data.Scripts[0])
		}
		return nil
	})
}

func testAccCheckApplicationApiDestroy(ctx context.Context, v *cdc.ApplicationData) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			ApplicationApiAPI.
			ReadApplicationData(ctx, *v.Id).
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

func testAccCheckApplicationApiDisappears(ctx context.Context, v *cdc.ApplicationData) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			ApplicationApiAPI.
			DeleteApplicationData(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccApplicationApiBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
}
`, name)
}

func testAccApplicationApiData(name string, content string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_data" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = %q
        }]
    })
}
`, name, content)
}

func testAccApplicationApiDataFromFile(name string, scriptPath string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_data_file" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = basename(%[2]q)
            content  = file(%[2]q)
        }]
    })
}
`, name, scriptPath)
}

func testAccApplicationApiDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_description" {
    name        = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
    description = %q
}
`, name, description)
}

func testAccApplicationApiEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_enabled" {
    name    = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
    enabled = %s
}
`, name, enabled)
}

func testAccApplicationApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_name" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
}
`, name)
}

func testAccApplicationApiTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_tags" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
    tags = %s
}
`, name, tagsStr)
}

func testAccApplicationApiVariables(name string, variablesValue string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_application" "test_variables" {
    name = %q
    data = jsonencode({
        scripts = [{
            filename = "script_1.py"
            content  = "print(\"automation\")"
        }]
    })
    variables = {
        var1 = %q
    }
}
`, name, variablesValue)
}
