package cdc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDestinationSplunkCloudResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "config_filename", ""),
					resource.TestCheckResourceAttr(resourceName, "config_option", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CSV"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_destination_splunk_cloud.test"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDestinationSplunkCloudDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDestinationSplunkCloudBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					testAccCheckDestinationSplunkCloudDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDestinationSplunkCloudResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_address"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudAddress(name, fmt.Sprintf("%s.example.com:9997", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s.example.com:9997", name)),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudAddress(name, fmt.Sprintf("%s-updated.example.com:9997", name)),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", fmt.Sprintf("%s-updated.example.com:9997", name)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_description"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_enabled"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_IndexerName(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_indexer_name"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudIndexerName(name, "main"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "indexer_name", "main"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudIndexerName(name, "main_updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "indexer_name", "main_updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_ConfigOption(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_config_option"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudConfigOption(name, "DEFAULT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_option", "DEFAULT"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudConfigOption(name, "MANUAL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "config_option", "MANUAL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_OutputDataFormat(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_output_data_format"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudOutputDataFormat(name, "CSV"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CSV"),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudOutputDataFormat(name, "CIM"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "output_data_format", "CIM"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_name"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDestinationSplunkCloudResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_destination_splunk_cloud.test_tags"
	var v cdc.DestinationSplunkCloud
	name := acctest.RandomNameWithPrefix("destination-splunk-cloud")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDestinationSplunkCloudTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccDestinationSplunkCloudTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDestinationSplunkCloudExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDestinationSplunkCloudExists(ctx context.Context, resourceName string, v *cdc.DestinationSplunkCloud) resource.TestCheckFunc {
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
			ReadDestinationSplunkCloud(ctx, id).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckDestinationSplunkCloudDestroy(ctx context.Context, v *cdc.DestinationSplunkCloud) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			ReadDestinationSplunkCloud(ctx, *v.Id).
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

func testAccCheckDestinationSplunkCloudDisappears(ctx context.Context, v *cdc.DestinationSplunkCloud) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			DestinationApiAPI.
			DeleteDestinationSplunkCloud(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

// testAccDestinationSplunkCloudBaseConfig loads the sample certificate used for
// both required certificates. The path is absolute because Terraform runs test
// configs from a temporary working directory.
func testAccDestinationSplunkCloudBaseConfig() string {
	_, thisFile, _, _ := runtime.Caller(0)
	certPath := filepath.Join(filepath.Dir(thisFile), "..", "..", "acctest", "testdata", "bloxone_cdc", "ca_cert.pem")
	return fmt.Sprintf(`
locals {
  splunk_cloud_cert = file(%q)
}
`, certPath)
}

func testAccDestinationSplunkCloudConfigOption(name string, configOption string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_config_option" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  config_option  = %[2]q
}
`, name, configOption)
}

func testAccDestinationSplunkCloudOutputDataFormat(name string, outputDataFormat string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_output_data_format" {
  name               = %[1]q
  address            = "%[1]s.example.com:9997"
  indexer_name       = "main"
  certificate        = local.splunk_cloud_cert
  ca_certificate     = local.splunk_cloud_cert
  output_data_format = %[2]q
}
`, name, outputDataFormat)
}

func testAccDestinationSplunkCloudBasicConfig(name string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
}
`, name)
}

func testAccDestinationSplunkCloudAddress(name string, address string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_address" {
  name           = %[1]q
  address        = %[2]q
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
}
`, name, address)
}

func testAccDestinationSplunkCloudIndexerName(name string, indexerName string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_indexer_name" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  indexer_name   = %[2]q
}
`, name, indexerName)
}

func testAccDestinationSplunkCloudDescription(name string, description string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_description" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  description    = %[2]q
}
`, name, description)
}

func testAccDestinationSplunkCloudEnabled(name string, enabled string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_enabled" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  enabled        = %[2]s
}
`, name, enabled)
}

func testAccDestinationSplunkCloudName(name string) string {
	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_name" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
}
`, name)
}

func testAccDestinationSplunkCloudTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return testAccDestinationSplunkCloudBaseConfig() + fmt.Sprintf(`
resource "bloxone_cdc_destination_splunk_cloud" "test_tags" {
  name           = %[1]q
  address        = "%[1]s.example.com:9997"
  indexer_name   = "main"
  certificate    = local.splunk_cloud_cert
  ca_certificate = local.splunk_cloud_cert
  tags           = %[2]s
}
`, name, tagsStr)
}
