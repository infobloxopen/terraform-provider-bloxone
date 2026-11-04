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

func TestAccEtlFilterThreatClassPropertyResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "Malicious"),
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

func TestAccEtlFilterThreatClassPropertyResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_etl_filter_threat_class_property.test"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckEtlFilterThreatClassPropertyDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccEtlFilterThreatClassPropertyBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					testAccCheckEtlFilterThreatClassPropertyDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_Data(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_data"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyData(name, `["Malicious"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "Malicious"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyData(name, `["DDoS", "Whitelist"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "data.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "data.0", "DDoS"),
					resource.TestCheckResourceAttr(resourceName, "data.1", "Whitelist"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_description"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyDescription(name, "Test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyDescription(name, "Test description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test description updated"),
				),
			},
			// Removing description from config resets it to the default
			{
				Config: testAccEtlFilterThreatClassPropertyDescriptionRemoved(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_enabled"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_FilterSource(t *testing.T) {
	// TODO: setting filter_source to SOURCE_INGRESS does not persist on the
	// CDC backend: a PUT with filter_source=SOURCE_INGRESS returns 200 and
	// echoes back SOURCE_INGRESS, but the next GET shows SOURCE_DEFAULT again
	// (confirmed with curl for record_type, no Terraform involved). Needs a
	// CDC backend fix before this can be un-skipped.
	t.Skip("CDC backend does not persist filter_source=SOURCE_INGRESS; see TODO above")

	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_filter_source"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyFilterSource(name, "SOURCE_DEFAULT"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_DEFAULT"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyFilterSource(name, "SOURCE_INGRESS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_source", "SOURCE_INGRESS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_name"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccEtlFilterThreatClassPropertyResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_etl_filter_threat_class_property.test_tags"
	var v cdc.EtlFilterThreatClassProperty
	name := acctest.RandomNameWithPrefix("etl-filter-threat-class-property")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccEtlFilterThreatClassPropertyTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccEtlFilterThreatClassPropertyTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEtlFilterThreatClassPropertyExists(context.Background(), resourceName, &v),
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

func testAccCheckEtlFilterThreatClassPropertyExists(ctx context.Context, resourceName string, v *cdc.EtlFilterThreatClassProperty) resource.TestCheckFunc {
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
			ReadEtlFilterThreatClassProperty(ctx, id).
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

func testAccCheckEtlFilterThreatClassPropertyDestroy(ctx context.Context, v *cdc.EtlFilterThreatClassProperty) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			EtlFilterApiAPI.
			ReadEtlFilterThreatClassProperty(ctx, *v.Id).
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

func testAccCheckEtlFilterThreatClassPropertyDisappears(ctx context.Context, v *cdc.EtlFilterThreatClassProperty) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			EtlFilterApiAPI.
			DeleteEtlFilterThreatClassProperty(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccEtlFilterThreatClassPropertyBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test" {
  name = %q
  data = ["Malicious"]
}
`, name)
}

func testAccEtlFilterThreatClassPropertyData(name string, data string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_data" {
  name = %q
  data = %s
}
`, name, data)
}

func testAccEtlFilterThreatClassPropertyDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_description" {
  name        = %q
  data        = ["Malicious"]
  description = %q
}
`, name, description)
}

func testAccEtlFilterThreatClassPropertyDescriptionRemoved(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_description" {
  name = %q
  data = ["Malicious"]
}
`, name)
}

func testAccEtlFilterThreatClassPropertyEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_enabled" {
  name    = %q
  data    = ["Malicious"]
  enabled = %s
}
`, name, enabled)
}

func testAccEtlFilterThreatClassPropertyFilterSource(name string, filterSource string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_filter_source" {
  name          = %q
  data          = ["Malicious"]
  filter_source = %q
}
`, name, filterSource)
}

func testAccEtlFilterThreatClassPropertyName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_name" {
  name = %q
  data = ["Malicious"]
}
`, name)
}

func testAccEtlFilterThreatClassPropertyTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("    %s = %q\n", k, v)
	}
	tagsStr += "  }"

	return fmt.Sprintf(`
resource "bloxone_cdc_etl_filter_threat_class_property" "test_tags" {
  name = %q
  data = ["Malicious"]
  tags = %s
}
`, name, tagsStr)
}
