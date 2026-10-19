package cdc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func TestAccFlowApiV2DataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_flows.test"
	resourceName := "bloxone_cdc_flow.test"
	flowName := acctest.RandomNameWithPrefix("flow")
	acctest.PreCheck(t)
	sourceID := testAccCreateCdcSourceNios(t, flowName)
	serviceID := testAccFlowApiV2ServiceID(t)
	var v cdc.FlowDataV2

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2DataSourceConfigFilters(flowName, serviceID, sourceID),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					}, testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

func TestAccFlowApiV2DataSource_TagFilters(t *testing.T) {
	dataSourceName := "data.bloxone_cdc_flows.test"
	resourceName := "bloxone_cdc_flow.test"
	flowName := acctest.RandomNameWithPrefix("flow")
	acctest.PreCheck(t)
	sourceID := testAccCreateCdcSourceNios(t, flowName)
	serviceID := testAccFlowApiV2ServiceID(t)
	var v cdc.FlowDataV2
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFlowApiV2Destroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFlowApiV2DataSourceConfigTagFilters(flowName, acctest.RandomName(), serviceID, sourceID),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckFlowApiV2Exists(context.Background(), resourceName, &v),
					}, testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// testAccFlowApiV2ServiceID looks up a CDC service (cdc_services) already present on the
// cluster. This provider does not have a resource to create one (the /v2/services API only
// supports GET - a CdcService only exists once an on-prem CDC host agent registers itself,
// which isn't something Terraform can drive), so these tests depend on at least one already
// existing on the test cluster.
func testAccFlowApiV2ServiceID(t *testing.T) string {
	t.Helper()
	apiRes, _, err := acctest.BloxOneClient.CDCAPI.ServiceApiAPI.
		ListCdcService(context.Background()).
		Execute()
	if err != nil {
		t.Fatalf("unable to list CDC services: %v", err)
	}
	results := apiRes.GetResults()
	if len(results) == 0 {
		t.Fatalf("no CDC service found on the cluster; these tests require at least one on-prem CDC host agent to have registered")
	}
	return *results[0].Id
}

// testAccCreateCdcSourceNios creates a NIOS source directly via the SDK (rather than through the
// bloxone_cdc_source_nios Terraform resource) and registers its cleanup. address/nios_username/
// nios_password don't need to point at a real NIOS grid - the API only validates format, not
// reachability.
func testAccCreateCdcSourceNios(t *testing.T, name string) int64 {
	t.Helper()
	ctx := context.Background()

	address := acctest.RandomNameWithPrefix("source") + ".example.com"
	niosUsername := "tf_auto"
	niosPassword := "tf_auto_password1!"
	scpUsername := "tf_auto_scp"
	scpPassword := "tf_auto_scp_password1!"
	// RPZ_LOG (one of the source_data_types values exercised by these tests)
	// requires the source to have a client certificate configured.
	certificate := acctest.ReadTestdata(t, "bloxone_cdc", "certkey.pem")
	apiRes, _, err := acctest.BloxOneClient.CDCAPI.SourceApiAPI.
		CreateSourceNios(ctx).
		Body(cdc.SourceNios{
			Name:         &name,
			Address:      &address,
			NiosUsername: &niosUsername,
			NiosPassword: &niosPassword,
			ScpUsername:  &scpUsername,
			ScpPassword:  &scpPassword,
			Certificate:  &certificate,
		}).
		Execute()
	if err != nil {
		t.Fatalf("unable to create CDC NIOS source: %v", err)
	}
	result := apiRes.GetResults()
	t.Cleanup(func() {
		_, err := acctest.BloxOneClient.CDCAPI.SourceApiAPI.DeleteSourceNios(ctx, *result.Id).Execute()
		if err != nil {
			t.Logf("unable to clean up CDC NIOS source %d: %v", *result.Id, err)
		}
	})
	return *result.Id
}

// below all TestAcc functions

func testAccCheckFlowApiV2ResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "cdc_services", dataSourceName, "results.0.cdc_services"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "destination_ha", dataSourceName, "results.0.destination_ha"),
		resource.TestCheckResourceAttrPair(resourceName, "destinations", dataSourceName, "results.0.destinations"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "results.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "etl_filters", dataSourceName, "results.0.etl_filters"),
		resource.TestCheckResourceAttrPair(resourceName, "filter_expression", dataSourceName, "results.0.filter_expression"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "schedule_option", dataSourceName, "results.0.schedule_option"),
		resource.TestCheckResourceAttrPair(resourceName, "script_schedule", dataSourceName, "results.0.script_schedule"),
		resource.TestCheckResourceAttrPair(resourceName, "source", dataSourceName, "results.0.source"),
		resource.TestCheckResourceAttrPair(resourceName, "source_data_types", dataSourceName, "results.0.source_data_types"),
		resource.TestCheckResourceAttrPair(resourceName, "source_log_types", dataSourceName, "results.0.source_log_types"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccFlowApiV2DataSourceConfigFilters(name string, serviceID string, sourceID int64) string {
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
}

data "bloxone_cdc_flows" "test" {
  filters = {
	name = bloxone_cdc_flow.test.name
  }
}
`, name, name, name, serviceID, sourceID)
}

func testAccFlowApiV2DataSourceConfigTagFilters(name string, tagValue string, serviceID string, sourceID int64) string {
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
  tags = {
	tag1 = %q
  }
}

data "bloxone_cdc_flows" "test" {
  tag_filters = {
	tag1 = bloxone_cdc_flow.test.tags.tag1
  }
}
`, name, name, name, serviceID, sourceID, tagValue)
}
