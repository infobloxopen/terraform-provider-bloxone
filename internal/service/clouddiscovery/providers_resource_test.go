package clouddiscovery_test

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccProvidersResource_basic(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersBasicConfig(name, "Amazon Web Services",
					"single", "role_arn", "dynamic", configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "destination_types_enabled.#"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "desired_state", "enabled"),
					resource.TestCheckResourceAttr(resourceName, "sync_interval", "Auto"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_disappears(t *testing.T) {
	resourceName := "bloxone_cloud_discovery_provider.test"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckProvidersDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccProvidersBasicConfig(name, "Amazon Web Services",
					"single", "role_arn", "dynamic", configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					testAccCheckProvidersDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccProvidersResource_AccountPreference(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_account_preference"
	var v1, v2 clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersAccountPreference(name, "Amazon Web Services",
					"single", "role_arn", "dynamic", configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "account_preference", "single"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersAccountPreference(name, "Amazon Web Services",
					"auto_discover_multiple", "role_arn", "dynamic", configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersDestroy(context.Background(), &v1),
					testAccCheckProvidersExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "account_preference", "auto_discover_multiple"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_AdditionalConfig(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_additional_config"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersAdditionalConfig(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "storage", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_config.object_type.objects.0.category.excluded", "true"),
					resource.TestCheckResourceAttr(resourceName, "additional_config.object_type.objects.0.category.id", "storage"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersAdditionalConfig(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "security", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "additional_config.object_type.objects.0.category.excluded", "true"),
					resource.TestCheckResourceAttr(resourceName, "additional_config.object_type.objects.0.category.id", "security")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_CredentialPreference(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_credential_preference"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersCredentialPreference(name, "Amazon Web Services",
					"single", "role_arn", "dynamic", configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "credential_preference.access_identifier_type", "role_arn"),
					resource.TestCheckResourceAttr(resourceName, "credential_preference.credential_type", "dynamic"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_Description(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_description"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersDescription(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "test description"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersDescription(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "updated test description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "updated test description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_DesiredState(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_desired_state"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersDesiredState(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_state", "disabled"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersDesiredState(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "enabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_state", "enabled"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_Destinations(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_destinations"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersDestinations(viewName, name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "IPAM/DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destinations.0.destination_type", "IPAM/DHCP"),
				),
			},
			{
				Config: testAccProvidersDestinations(viewName, name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "destinations.0.destination_type", "IPAM/DHCP"),
					resource.TestCheckResourceAttr(resourceName, "destinations.1.destination_type", "DNS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_Name(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_name"
	var v1, v2 clouddiscovery.DiscoveryConfig
	name1 := acctest.RandomName()
	name2 := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersName(name1, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersName(name2, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersDestroy(context.Background(), &v1),
					testAccCheckProvidersExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_ProviderType(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_provider_type"
	var v1, v2, v3 clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	randNumber := randomNumber()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randNumber)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create and Read
			{
				Config: testAccProvidersProviderType(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, randNumber),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "provider_type", "Amazon Web Services"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersProviderType(name, "Google Cloud Platform",
					"single", "project_id", "dynamic",
					"33333333333", "33333333333"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersDestroy(context.Background(), &v1),
					testAccCheckProvidersExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "provider_type", "Google Cloud Platform"),
				),
			},
			//Update and Read
			{
				Config: testAccProvidersProviderType(name, "Microsoft Azure",
					"single", "tenant_id", "dynamic",
					"44444444444", "2222222222"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersDestroy(context.Background(), &v2),
					testAccCheckProvidersExists(context.Background(), resourceName, &v3),
					resource.TestCheckResourceAttr(resourceName, "provider_type", "Microsoft Azure"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_SourceConfigs(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_source_configs"
	var v1, v2 clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersSourceConfigs(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "source_configs.0.credential_config.access_identifier", configAccessId),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersSourceConfigs(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					"arn:aws:iam::987654321098:role/infoblox_discovery"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersDestroy(context.Background(), &v1),
					testAccCheckProvidersExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "source_configs.0.credential_config.access_identifier", "arn:aws:iam::987654321098:role/infoblox_discovery"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_SyncInterval(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_sync_interval"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersSyncInterval(name, "Amazon Web Services",
					"single", configAccessId, "15"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_interval", "15"),
				),
			},
			// Update and Read
			{
				Config: testAccProvidersSyncInterval(name, "Amazon Web Services",
					"single", configAccessId, "Auto"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_interval", "Auto"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccProvidersResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cloud_discovery_provider.test_tags"
	var v clouddiscovery.DiscoveryConfig
	name := acctest.RandomName()
	configAccessId := fmt.Sprintf("arn:aws:iam::%s:role/infoblox_discovery", randomNumber())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccProvidersTags(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, map[string]string{
						"tag1": "value1",
						"tag2": "value2",
					}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2")),
			},
			// Update and Read
			{
				Config: testAccProvidersTags(name, "Amazon Web Services",
					"single", "role_arn", "dynamic",
					configAccessId, map[string]string{
						"tag2": "value2changed",
						"tag3": "value3",
					}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckProvidersExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3")),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckProvidersExists(ctx context.Context, resourceName string, v *clouddiscovery.DiscoveryConfig) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DiscoveryConfigurationAPIV2.
			ProvidersAPI.
			Read(ctx, rs.Primary.ID).
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

func testAccCheckProvidersDestroy(ctx context.Context, v *clouddiscovery.DiscoveryConfig) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DiscoveryConfigurationAPIV2.
			ProvidersAPI.
			Read(ctx, *v.Id).
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

func testAccCheckProvidersDisappears(ctx context.Context, v *clouddiscovery.DiscoveryConfig) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DiscoveryConfigurationAPIV2.
			ProvidersAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

// RandomNumber Function to generate a random 12-digit number for cloud discovery Role ARN
func randomNumber() string {
	return fmt.Sprintf("%d", rand.Intn(999999999999))
}

func testAccProvidersBasicConfig(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
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
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}

func testAccProvidersAccountPreference(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_account_preference" {
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
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}

func testAccProvidersAdditionalConfig(name, providerType, accountPreference, accessIdType, credType, configAccessId, addConfigCategoryId, addConfigCategoryExcluded string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_additional_config" {
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
    additional_config = {
		object_type = {
			objects = [ {
					category = {
						excluded = %q
						id = %q
					}
				}
			] 
		}
	}
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId, addConfigCategoryExcluded, addConfigCategoryId)
}

func testAccProvidersCredentialPreference(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_credential_preference" {
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
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}

func testAccProvidersDescription(name, providerType, accountPreference, accessIdType, credType, configAccessId, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_description" {
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
    description = %q
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId, description)
}

func testAccProvidersDesiredState(name, providerType, accountPreference, accessIdType, credType, configAccessId, desiredState string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_desired_state" {
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
    desired_state = %q
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId, desiredState)
}

func testAccProvidersDestinations(viewName, name, providerType, accountPreference, accessIdType, credType, configAccessId, destinationType string) string {
	destinationsStr := ""
	destinationTypeEnabledStr := ""
	if destinationType == "IPAM/DHCP" {
		destinationsStr = "{\n\t\t\tconfig = {}\n\t\t\tdestination_type = \"IPAM/DHCP\"\n\t\t}"
		destinationTypeEnabledStr = "\"IPAM/DHCP\""
	}
	if destinationType == "DNS" {
		destinationTypeEnabledStr = "\"IPAM/DHCP\" , \"DNS\""
		destinationsStr = "{\n\t\t\tconfig = {}\n\t\t\tdestination_type = \"IPAM/DHCP\"\n\t\t},\n\t\t{\n\t\t\tconfig = {\n\t\t\t\tdns = {\n\t\t\t\t\tview_id = bloxone_dns_view.test.id\n\t\t\t\t}\n\t\t\t}\n\t\t\tdestination_type = \"DNS\"\n\t\t}"
	}
	return fmt.Sprintf(`
resource "bloxone_dns_view" "test" {
    name = %q
}

resource "bloxone_cloud_discovery_provider" "test_destinations" {
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
	destination_types_enabled = [%s]
	destinations = [
		%s
	]
}
`, viewName, name, providerType, accountPreference, accessIdType, credType, configAccessId, destinationTypeEnabledStr, destinationsStr)
}

func testAccProvidersName(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_name" {
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
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}

func testAccProvidersProviderType(name, providerType, accountPreference, accessIdType, credType, configAccessId, restrictedAccounts string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_provider_type" {
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
		restricted_to_accounts = [%q]
	}]
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId, restrictedAccounts)
}

func testAccProvidersSourceConfigs(name, providerType, accountPreference, accessIdType, credType, configAccessId string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_source_configs" {
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
`, name, providerType, accountPreference, accessIdType, credType, configAccessId)
}

func testAccProvidersSyncInterval(name string, providerType, accountPreference, configAccessId, syncInterval string) string {
	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_sync_interval" {
    name = %q
	provider_type = %q
	account_preference = %q
	credential_preference = {
		access_identifier_type = "role_arn"
		credential_type = "dynamic"
	}
	source_configs = [ {
		credential_config = {
				access_identifier = %q
			}
	}]
    sync_interval = %q
}
`, name, providerType, accountPreference, configAccessId, syncInterval)
}

func testAccProvidersTags(name string, providerType, accountPreference, accessIdType, credType, configAccessId string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cloud_discovery_provider" "test_tags" {
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
    tags = %s
}
`, name, providerType, accountPreference, accessIdType, credType, configAccessId, tagsStr)
}
