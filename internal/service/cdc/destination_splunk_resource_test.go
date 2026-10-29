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

func TestAccDestinationSplunkResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CSV"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_splunk.test"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					testAccCheckDestinationSplunkDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationSplunkResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_address"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkAddress(name, fmt.Sprintf("%s.example.com:9997", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s.example.com:9997", name)),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkAddress(name, fmt.Sprintf("%s-updated.example.com:9997", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s-updated.example.com:9997", name)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_description"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_enabled"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_IndexerName(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_indexer_name"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkIndexerName(name, "main"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "indexer_name", "main"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkIndexerName(name, "main_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "indexer_name", "main_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_output_data_format"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkOutputDataFormat(name, "CSV"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CSV"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkOutputDataFormat(name, "CIM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CIM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_name"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk.test_tags"
	var v cdc.DestinationSplunk
	name := acctest.RandomNameWithPrefix("destination-splunk")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationSplunkExists(ctx context.Context, resourceName string, v *cdc.DestinationSplunk) resource.TestCheckFunc {
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
			ReadDestinationSplunk(ctx, id).
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

func testAccCheckDestinationSplunkDestroy(ctx context.Context, v *cdc.DestinationSplunk) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationSplunk(ctx, *v.Id).
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

func testAccCheckDestinationSplunkDisappears(ctx context.Context, v *cdc.DestinationSplunk) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationSplunk(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccDestinationSplunkBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
}
`, name)
}

func testAccDestinationSplunkAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_address" {
  name         = %[1]q
  address      = %[2]q
  indexer_name = "main"
}
`, name, address)
}

func testAccDestinationSplunkIndexerName(name string, indexerName string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_indexer_name" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = %[2]q
}
`, name, indexerName)
}

func testAccDestinationSplunkDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_description" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
  description  = %[2]q
}
`, name, description)
}

func testAccDestinationSplunkEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_enabled" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
  enabled      = %[2]s
}
`, name, enabled)
}

func testAccDestinationSplunkOutputDataFormat(name string, outputDataFormat string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_output_data_format" {
  name               = %[1]q
  address            = "%[1]s.example.com:9997"
  indexer_name       = "main"
  output_data_format = %[2]q
}
`, name, outputDataFormat)
}

func testAccDestinationSplunkName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_name" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
}
`, name)
}

func testAccDestinationSplunkTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk" "test_tags" {
  name         = %[1]q
  address      = "%[1]s.example.com:9997"
  indexer_name = "main"
  tags         = %[2]s
}
`, name, tagsStr)
}
