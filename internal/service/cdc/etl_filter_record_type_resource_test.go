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

func TestAccEtlFilterRecordTypeResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "A"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_DEFAULT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_etl_filter_record_type.test"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterRecordTypeDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterRecordTypeBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					testAccCheckEtlFilterRecordTypeDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccEtlFilterRecordTypeResource_Data(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_data"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeData(name, `["A"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "A"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeData(name, `["A", "AAAA"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "A"),
					resource.TestCheckResourceAttr(resourceName, "data.1", "AAAA"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_description"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Removing description from config resets it to the default
			{
				Config: testAccEtlFilterRecordTypeDescriptionRemoved(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_enabled"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_FilterSource(t *testing.T) {
	// TODO: setting filter_source to SOURCE_INGRESS does not persist on the
	// CDC backend: a PUT with filter_source=SOURCE_INGRESS returns 200 and
	// echoes back SOURCE_INGRESS, but the next GET shows SOURCE_DEFAULT again
	// (confirmed with curl for record_type, no Terraform involved). Needs a
	// CDC backend fix before this can be un-skipped.
	t.Skip("CDC backend does not persist filter_source=SOURCE_INGRESS; see TODO above")

	var resourceName = "bloxone_cdc_etl_filter_record_type.test_filter_source"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeFilterSource(name, "SOURCE_DEFAULT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_DEFAULT"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeFilterSource(name, "SOURCE_INGRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_INGRESS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_name"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterRecordTypeResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_tags"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter-record-type")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterRecordTypeTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterRecordTypeTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterRecordTypeExists(context.Background(), resourceName, &v),
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

func testAccCheckEtlFilterRecordTypeExists(ctx context.Context, resourceName string, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
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
			EtlFilterApiAPI.
			ReadEtlFilterRecordType(ctx, id).
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

func testAccCheckEtlFilterRecordTypeDestroy(ctx context.Context, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			EtlFilterApiAPI.
			ReadEtlFilterRecordType(ctx, *v.Id).
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

func testAccCheckEtlFilterRecordTypeDisappears(ctx context.Context, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			EtlFilterApiAPI.
			DeleteEtlFilterRecordType(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccEtlFilterRecordTypeBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test" {
  name = %q
  data = ["A"]
}
`, name)
}

func testAccEtlFilterRecordTypeData(name string, data string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_data" {
  name = %q
  data = %s
}
`, name, data)
}

func testAccEtlFilterRecordTypeDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_description" {
  name        = %q
  data        = ["A"]
  description = %q
}
`, name, description)
}

func testAccEtlFilterRecordTypeDescriptionRemoved(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_description" {
  name = %q
  data = ["A"]
}
`, name)
}

func testAccEtlFilterRecordTypeEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_enabled" {
  name    = %q
  data    = ["A"]
  enabled = %s
}
`, name, enabled)
}

func testAccEtlFilterRecordTypeFilterSource(name string, filterSource string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_filter_source" {
  name          = %q
  data          = ["A"]
  filter_source = %q
}
`, name, filterSource)
}

func testAccEtlFilterRecordTypeName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_name" {
  name = %q
  data = ["A"]
}
`, name)
}

func testAccEtlFilterRecordTypeTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_tags" {
  name = %q
  data = ["A"]
  tags = %s
}
`, name, tagsStr)
}
