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

func TestAccFlowApiV2Resource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2BasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_ha"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_flow.test"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2BasicConfig(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					testAccCheckFlowApiV2Disappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFlowApiV2Resource_CdcServices(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_cdc_services"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2CdcServices("CDC_SERVICES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cdc_services", "CDC_SERVICES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2CdcServices("CDC_SERVICES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cdc_services", "CDC_SERVICES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_description"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Description("DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Description("DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Destinations(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_destinations"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Destinations("DESTINATIONS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destinations", "DESTINATIONS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Destinations("DESTINATIONS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destinations", "DESTINATIONS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_enabled"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Enabled("ENABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "ENABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Enabled("ENABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "ENABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_EtlFilters(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_etl_filters"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2EtlFilters("ETL_FILTERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "etl_filters", "ETL_FILTERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2EtlFilters("ETL_FILTERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "etl_filters", "ETL_FILTERS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_FilterExpression(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_filter_expression"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2FilterExpression("FILTER_EXPRESSION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_expression", "FILTER_EXPRESSION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2FilterExpression("FILTER_EXPRESSION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "filter_expression", "FILTER_EXPRESSION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_name"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Name("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Name("NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_ScheduleOption(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_schedule_option"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2ScheduleOption("SCHEDULE_OPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "schedule_option", "SCHEDULE_OPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2ScheduleOption("SCHEDULE_OPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "schedule_option", "SCHEDULE_OPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_ScriptSchedule(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_script_schedule"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2ScriptSchedule("SCRIPT_SCHEDULE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "script_schedule", "SCRIPT_SCHEDULE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2ScriptSchedule("SCRIPT_SCHEDULE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "script_schedule", "SCRIPT_SCHEDULE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Source(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_source"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Source("SOURCE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source", "SOURCE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Source("SOURCE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source", "SOURCE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_SourceDataTypes(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_source_data_types"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2SourceDataTypes("SOURCE_DATA_TYPES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_data_types", "SOURCE_DATA_TYPES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2SourceDataTypes("SOURCE_DATA_TYPES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_data_types", "SOURCE_DATA_TYPES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_SourceLogTypes(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_source_log_types"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2SourceLogTypes("SOURCE_LOG_TYPES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_log_types", "SOURCE_LOG_TYPES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2SourceLogTypes("SOURCE_LOG_TYPES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_log_types", "SOURCE_LOG_TYPES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_tags"
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Tags("TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Tags("TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFlowApiV2Exists(ctx context.Context, resourceName string, v *cdc.FlowDataV2) resource.TestCheckFunc {
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
			FlowApiV2API.
			ReadFlowDataV2(ctx, id).
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

func testAccCheckFlowApiV2Destroy(ctx context.Context, v *cdc.FlowDataV2) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			FlowApiV2API.
			ReadFlowDataV2(ctx, *v.Id).
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

func testAccCheckFlowApiV2Disappears(ctx context.Context, v *cdc.FlowDataV2) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			FlowApiV2API.
			DeleteFlowDataV2(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFlowApiV2BasicConfig(string) string {
	// TODO: create basic resource with required fields
	return `
resource "bloxone_cdc_flow" "test" {
}
`
}

func testAccFlowApiV2CdcServices(cdcServices string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_cdc_services" {
    cdc_services = %q
}
`, cdcServices)
}

func testAccFlowApiV2Description(description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_description" {
    description = %q
}
`, description)
}

func testAccFlowApiV2Destinations(destinations string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_destinations" {
    destinations = %q
}
`, destinations)
}

func testAccFlowApiV2Enabled(enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_enabled" {
    enabled = %q
}
`, enabled)
}

func testAccFlowApiV2EtlFilters(etlFilters string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_etl_filters" {
    etl_filters = %q
}
`, etlFilters)
}

func testAccFlowApiV2FilterExpression(filterExpression string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_filter_expression" {
    filter_expression = %q
}
`, filterExpression)
}

func testAccFlowApiV2Name(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_name" {
    name = %q
}
`, name)
}

func testAccFlowApiV2ScheduleOption(scheduleOption string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_schedule_option" {
    schedule_option = %q
}
`, scheduleOption)
}

func testAccFlowApiV2ScriptSchedule(scriptSchedule string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_script_schedule" {
    script_schedule = %q
}
`, scriptSchedule)
}

func testAccFlowApiV2Source(source string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_source" {
    source = %q
}
`, source)
}

func testAccFlowApiV2SourceDataTypes(sourceDataTypes string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_source_data_types" {
    source_data_types = %q
}
`, sourceDataTypes)
}

func testAccFlowApiV2SourceLogTypes(sourceLogTypes string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_source_log_types" {
    source_log_types = %q
}
`, sourceLogTypes)
}

func testAccFlowApiV2Tags(tags string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_flow" "test_tags" {
    tags = %q
}
`, tags)
}
