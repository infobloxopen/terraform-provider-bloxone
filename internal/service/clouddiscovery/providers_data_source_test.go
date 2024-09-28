package clouddiscovery_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccProvidersDataSource_Filters(t *testing.T) {
	dataSourceName := "data.bloxone_cloud_discovery_providers.test"
	resourceName := "bloxone_cloud_discovery_provider.test"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckProvidersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccProvidersDataSourceConfigFilters(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					"arn:aws:iam::123456789012:role/infoblox_discovery"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckProvidersExists(context.Background(), resourceName, &v),
					}, testAccCheckProvidersResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckProvidersResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "account_preference", dataSourceName, "results.0.account_preference"),
		resource.TestCheckResourceAttrPair(resourceName, "additional_config", dataSourceName, "results.0.additional_config"),
		resource.TestCheckResourceAttrPair(resourceName, "created_at", dataSourceName, "results.0.created_at"),
		resource.TestCheckResourceAttrPair(resourceName, "credential_preference", dataSourceName, "results.0.credential_preference"),
		resource.TestCheckResourceAttrPair(resourceName, "deleted_at", dataSourceName, "results.0.deleted_at"),
		resource.TestCheckResourceAttrPair(resourceName, "description", dataSourceName, "results.0.description"),
		resource.TestCheckResourceAttrPair(resourceName, "desired_state", dataSourceName, "results.0.desired_state"),
		resource.TestCheckResourceAttrPair(resourceName, "destination_types_enabled", dataSourceName, "results.0.destination_types_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "destinations", dataSourceName, "results.0.destinations"),
		resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "results.0.id"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "results.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "provider_type", dataSourceName, "results.0.provider_type"),
		resource.TestCheckResourceAttrPair(resourceName, "source_configs", dataSourceName, "results.0.source_configs"),
		resource.TestCheckResourceAttrPair(resourceName, "status_message", dataSourceName, "results.0.status_message"),
		resource.TestCheckResourceAttrPair(resourceName, "sync_interval", dataSourceName, "results.0.sync_interval"),
		resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "results.0.tags"),
		resource.TestCheckResourceAttrPair(resourceName, "updated_at", dataSourceName, "results.0.updated_at"),
	}
}

func testAccProvidersDataSourceConfigFilters(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test" {
    name = %q
	provider_type = %q
	account_preference = %q
	credential_preference = {
		access_identifier_type = %q
		credential_type = %q
	}
	source_configs = [ {
		credential_config = {
				access_identifier = %q
			}
	}]
}

data "bloxone_cloud_discovery_providers" "test" {
  filters = {
	name = bloxone_cloud_discovery_provider.test.name
  }
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}
