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

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccEtlFilterApiResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_DEFAULT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_etl_filter_record_type.test"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					testAccCheckEtlFilterApiDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccEtlFilterApiResource_Data(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_data"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiData(name, `["A"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "A"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiData(name, `["A", "AAAA"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_description"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiDescription(name, "DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiDescription(name, "DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_enabled"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_FilterSource(t *testing.T) {
	// TODO: setting filter_source to SOURCE_INGRESS does not persist on the
	// CDC backend. Confirmed via direct curl against the API, with no
	// provider/Terraform involved: a PUT with filter_source=SOURCE_INGRESS
	// returns 200 and echoes back filter_source="SOURCE_INGRESS" in the same
	// response, but an immediately following GET on the same object shows it
	// reverted to filter_source="SOURCE_DEFAULT". This causes Terraform's
	// post-apply refresh to see a stale diff on every apply. This is a
	// backend persistence bug, not something fixable in the provider. Needs
	// a CDC backend fix before this can be un-skipped.
	t.Skip("CDC backend does not persist filter_source=SOURCE_INGRESS; see TODO above")

	var resourceName = "bloxone_cdc_etl_filter_record_type.test_filter_source"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiFilterSource(name, "SOURCE_DEFAULT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_DEFAULT"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiFilterSource(name, "SOURCE_INGRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_INGRESS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_name"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_record_type.test_tags"
	var v cdc.EtlFilterRecordType
	name := acctest.RandomNameWithPrefix("etl-filter")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterApiTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterApiTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckEtlFilterApiExists(ctx context.Context, resourceName string, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
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

func testAccCheckEtlFilterApiDestroy(ctx context.Context, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
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

func testAccCheckEtlFilterApiDisappears(ctx context.Context, v *cdc.EtlFilterRecordType) resource.TestCheckFunc {
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

func testAccEtlFilterApiBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test" {
    name = %q
    data = ["A"]
}
`, name)
}

func testAccEtlFilterApiData(name string, data string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_data" {
    name = %q
    data = %s
}
`, name, data)
}

func testAccEtlFilterApiDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_description" {
    name        = %q
    data        = ["A"]
    description = %q
}
`, name, description)
}

func testAccEtlFilterApiEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_enabled" {
    name    = %q
    data    = ["A"]
    enabled = %s
}
`, name, enabled)
}

func testAccEtlFilterApiFilterSource(name string, filterSource string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_filter_source" {
    name          = %q
    data          = ["A"]
    filter_source = %q
}
`, name, filterSource)
}

func testAccEtlFilterApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_name" {
    name = %q
    data = ["A"]
}
`, name)
}

func testAccEtlFilterApiTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_record_type" "test_tags" {
    name = %q
    data = ["A"]
    tags = %s
}
`, name, tagsStr)
}
