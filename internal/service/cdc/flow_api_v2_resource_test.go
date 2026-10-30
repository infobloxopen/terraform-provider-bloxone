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
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2BasicConfig(name, serviceID, sourceID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_ha"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Import and Read
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_flow.test"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2BasicConfig(name, serviceID, sourceID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					testAccCheckFlowApiV2Disappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFlowApiV2Resource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_description"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Description(name, serviceID, sourceID, "Test flow description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Test flow description"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Description(name, serviceID, sourceID, "Updated test flow description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Updated test flow description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Destinations(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_destinations"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Destinations(name, serviceID, sourceID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destinations.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "destinations.0", "bloxone_cdc_destination_syslog.test_destinations", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_enabled"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Enabled(name, serviceID, sourceID, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Enabled(name, serviceID, sourceID, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_EtlFilters(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_etl_filters"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	etlFilterID := testAccCreateCdcEtlFilterFqdn(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2EtlFilters(name, serviceID, sourceID, etlFilterID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "etl_filters.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "etl_filters.0", strconv.FormatInt(etlFilterID, 10)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_FilterExpression(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_filter_expression"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read: filter_expression refers to the applied ETL filter by name
			{
				Config: testAccFlowApiV2FilterExpression(name, serviceID, sourceID, "a"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "etl_filters.0", "bloxone_cdc_etl_filter_record_type.test_filter_expression_a", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "filter_expression", "bloxone_cdc_etl_filter_record_type.test_filter_expression_a", "name"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2FilterExpression(name, serviceID, sourceID, "b"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "etl_filters.0", "bloxone_cdc_etl_filter_record_type.test_filter_expression_b", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "filter_expression", "bloxone_cdc_etl_filter_record_type.test_filter_expression_b", "name"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_name"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Name(name, serviceID, sourceID, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Name(name, serviceID, sourceID, name+"-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Source(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_source"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID1 := testAccCreateCdcSourceNios(t, name+"-1")
	sourceID2 := testAccCreateCdcSourceNios(t, name+"-2")
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Source(name, serviceID, sourceID1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source", strconv.FormatInt(sourceID1, 10)),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Source(name, serviceID, sourceID2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source", strconv.FormatInt(sourceID2, 10)),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_SourceDataTypes(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_source_data_types"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2SourceDataTypes(name, serviceID, sourceID, `["QUERY_RESP_LOG"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_data_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_data_types.0", "QUERY_RESP_LOG"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2SourceDataTypes(name, serviceID, sourceID, `["RPZ_LOG"]`),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "source_data_types.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "source_data_types.0", "RPZ_LOG"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFlowApiV2Resource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_tags"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	sourceID := testAccCreateCdcSourceNios(t, name)
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Tags(name, serviceID, sourceID, map[string]string{
					"tag1": "value1",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Tags(name, serviceID, sourceID, map[string]string{
					"tag1": "value1changed",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
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

// testAccCreateCdcEtlFilterFqdn creates an FQDN ETL filter directly via the SDK (this provider
// does not yet expose a resource for ETL filters) and registers its cleanup.
func testAccCreateCdcEtlFilterFqdn(t *testing.T, name string) int64 {
	t.Helper()
	ctx := context.Background()

	data := []string{"example.com"}
	apiRes, _, err := acctest.BloxOneClient.CDCAPI.EtlFilterApiAPI.
		CreateEtlFilterFqdn(ctx).
		Body(cdc.EtlFilterFqdn{
			Name: &name,
			Data: data,
		}).
		Execute()
	if err != nil {
		t.Fatalf("unable to create CDC ETL filter: %v", err)
	}
	result := apiRes.GetResults()
	t.Cleanup(func() {
		_, err := acctest.BloxOneClient.CDCAPI.EtlFilterApiAPI.DeleteEtlFilterFqdn(ctx, *result.Id).Execute()
		if err != nil {
			t.Logf("unable to clean up CDC ETL filter %d: %v", *result.Id, err)
		}
	})
	return *result.Id
}

func testAccFlowApiV2BasicConfig(name string, serviceID string, sourceID int64) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test.id]

  depends_on = [bloxone_cdc_destination_syslog.test]
}
`, name, name, name, serviceID, sourceID)
}

func testAccFlowApiV2Description(name string, serviceID string, sourceID int64, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_description" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_description" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_description.id]
  description       = %q

  depends_on = [bloxone_cdc_destination_syslog.test_description]
}
`, name, name, name, serviceID, sourceID, description)
}

func testAccFlowApiV2Destinations(name string, serviceID string, sourceID int64) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_destinations" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_destinations" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_destinations.id]

  depends_on = [bloxone_cdc_destination_syslog.test_destinations]
}
`, name, name, name, serviceID, sourceID)
}

func testAccFlowApiV2Enabled(name string, serviceID string, sourceID int64, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_enabled" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_enabled" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_enabled.id]
  enabled           = %s

  depends_on = [bloxone_cdc_destination_syslog.test_enabled]
}
`, name, name, name, serviceID, sourceID, enabled)
}

func testAccFlowApiV2EtlFilters(name string, serviceID string, sourceID int64, etlFilterID int64) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_etl_filters" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_etl_filters" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_etl_filters.id]
  etl_filters       = [%d]

  depends_on = [bloxone_cdc_destination_syslog.test_etl_filters]
}
`, name, name, name, serviceID, sourceID, etlFilterID)
}

// testAccFlowApiV2FilterExpressionPrereqs creates just the destination and ETL filter that
// testAccFlowApiV2FilterExpression's flow depends on, applied as their own TestStep first.
func testAccFlowApiV2FilterExpression(name string, serviceID string, sourceID int64, filter string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_filter_expression" {
  name    = %[1]q
  address = "%[1]s.example.com"
}

resource "bloxone_cdc_etl_filter_record_type" "test_filter_expression_a" {
  name = "%[1]s-a"
  data = ["A"]
}

resource "bloxone_cdc_etl_filter_record_type" "test_filter_expression_b" {
  name = "%[1]s-b"
  data = ["AAAA"]
}

resource "bloxone_cdc_flow" "test_filter_expression" {
  name              = %[1]q
  cdc_services      = [%[2]q]
  source            = %[3]d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_filter_expression.id]
  etl_filters       = [bloxone_cdc_etl_filter_record_type.test_filter_expression_%[4]s.id]
  filter_expression = bloxone_cdc_etl_filter_record_type.test_filter_expression_%[4]s.name
}
`, name, serviceID, sourceID, filter)
}

func testAccFlowApiV2Name(resourceSuffix string, serviceID string, sourceID int64, name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_name" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_name" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_name.id]

  depends_on = [bloxone_cdc_destination_syslog.test_name]
}
`, resourceSuffix, resourceSuffix, name, serviceID, sourceID)
}

func testAccFlowApiV2Source(name string, serviceID string, sourceID int64) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_source" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_source" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_source.id]

  depends_on = [bloxone_cdc_destination_syslog.test_source]
}
`, name, name, name, serviceID, sourceID)
}

func testAccFlowApiV2SourceDataTypes(name string, serviceID string, sourceID int64, sourceDataTypes string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_source_data_types" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_source_data_types" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = %s
  destinations      = [bloxone_cdc_destination_syslog.test_source_data_types.id]

  depends_on = [bloxone_cdc_destination_syslog.test_source_data_types]
}
`, name, name, name, serviceID, sourceID, sourceDataTypes)
}

func testAccFlowApiV2Tags(name string, serviceID string, sourceID int64, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_destination_syslog" "test_tags" {
  name    = %q
  address = "%s.example.com"
}

resource "bloxone_cdc_flow" "test_tags" {
  name              = %q
  cdc_services      = [%q]
  source            = %d
  source_data_types = ["QUERY_RESP_LOG"]
  destinations      = [bloxone_cdc_destination_syslog.test_tags.id]
  tags              = %s

  depends_on = [bloxone_cdc_destination_syslog.test_tags]
}
`, name, name, name, serviceID, sourceID, tagsStr)
}

func TestAccFlowApiV2Resource_Schedule(t *testing.T) {
	var resourceName = "bloxone_cdc_flow.test_schedule"
	var v cdc.FlowDataV2
	acctest.PreCheck(t)
	name := acctest.RandomNameWithPrefix("flow")
	serviceID := testAccFlowApiV2ServiceID(t)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFlowApiV2Schedule(name, serviceID, "0 0 0 */2 * * *"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "schedule_option", "custom"),
					resource.TestCheckResourceAttr(resourceName, "script_schedule", "0 0 0 */2 * * *"),
					resource.TestCheckResourceAttr(resourceName, "source_data_types.#", "0"),
				),
			},
			// Update and Read
			{
				Config: testAccFlowApiV2Schedule(name, serviceID, "0 0 0 */4 * * *"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "script_schedule", "0 0 0 */4 * * *"),
				),
			},
			// Import and Read
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccFlowApiV2Schedule(name string, serviceID string, scriptSchedule string) string {
	return fmt.Sprintf(`
data "bloxone_cdc_sources" "schedule" {
  filters = {
    type = "SOURCE_SCHEDULE"
  }
}

resource "bloxone_cdc_application" "test_schedule" {
  name = "%[1]s-app"
  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
}

resource "bloxone_cdc_destination_application" "test_schedule" {
  name        = %[1]q
  application = bloxone_cdc_application.test_schedule.id
}

resource "bloxone_cdc_flow" "test_schedule" {
  name              = %[1]q
  cdc_services      = [%[2]q]
  source            = data.bloxone_cdc_sources.schedule.results.0.id
  source_data_types = []
  destinations      = [bloxone_cdc_destination_application.test_schedule.id]
  schedule_option   = "custom"
  script_schedule   = %[3]q
}
`, name, serviceID, scriptSchedule)
}
