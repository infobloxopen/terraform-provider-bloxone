package infra_mgmt_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccServicesResource_basic(t *testing.T) {
	var resourceName = "bloxone_infra_service.test"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesBasicConfig(hostName, serviceName, "dhcp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttrPair(resourceName, "pool_id", "bloxone_infra_host.test", "pool_id"),
					resource.TestCheckResourceAttr(resourceName, "service_type", "dhcp"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "desired_state", "stop"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServicesResource_disappears(t *testing.T) {
	resourceName := "bloxone_infra_service.test"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckServicesDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccServicesBasicConfig(hostName, serviceName, "dhcp"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					testAccCheckServicesDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccServicesResource_Description(t *testing.T) {
	var resourceName = "bloxone_infra_service.test_description"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesDescription(hostName, serviceName, "dhcp", "service description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "service description"),
				),
			},
			// Update and Read
			{
				Config: testAccServicesDescription(hostName, serviceName, "dhcp", "service description updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "service description updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServicesResource_DesiredState(t *testing.T) {
	var resourceName = "bloxone_infra_service.test_desired_state"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesDesiredState(hostName, serviceName, "dhcp", "start"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_state", "start"),
				),
			},
			// Update and Read
			{
				Config: testAccServicesDesiredState(hostName, serviceName, "dhcp", "stop"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_state", "stop"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServicesResource_DesiredVersion(t *testing.T) {
	var resourceName = "bloxone_infra_service.test_desired_version"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesDesiredVersion(hostName, serviceName, "dhcp", "3.4.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_version", "3.4.0"),
				),
			},
			// Update and Read
			{
				Config: testAccServicesDesiredVersion(hostName, serviceName, "dhcp", "3.5.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "desired_version", "3.5.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServicesResource_InterfaceLabels(t *testing.T) {
	var resourceName = "bloxone_infra_service.test_interface_labels"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesInterfaceLabels(hostName, serviceName, "dhcp", []string{"WAN", "LAN"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interface_labels.0", "WAN"),
					resource.TestCheckResourceAttr(resourceName, "interface_labels.1", "LAN"),
				),
			},
			// Update and Read
			{
				Config: testAccServicesInterfaceLabels(hostName, serviceName, "dhcp", []string{"label1", "label2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "interface_labels.0", "label1"),
					resource.TestCheckResourceAttr(resourceName, "interface_labels.1", "label2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccServicesResource_Tags(t *testing.T) {
	var resourceName = "bloxone_infra_service.test_tags"
	var v infra_mgmt.InfraService
	var serviceName = acctest.RandomNameWithPrefix("service")
	var hostName = acctest.RandomNameWithPrefix("host")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccServicesTags(hostName, serviceName, "dhcp", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccServicesTags(hostName, serviceName, "dhcp", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckServicesExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckServicesExists(ctx context.Context, resourceName string, v *infra_mgmt.InfraService) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.InfraManagementAPI.
			ServicesAPI.
			ServicesRead(ctx, rs.Primary.ID).
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

func testAccCheckServicesDestroy(ctx context.Context, v *infra_mgmt.InfraService) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.InfraManagementAPI.
			ServicesAPI.
			ServicesRead(ctx, *v.Id).
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

func testAccCheckServicesDisappears(ctx context.Context, v *infra_mgmt.InfraService) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.InfraManagementAPI.
			ServicesAPI.
			ServicesDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}
func testAccServicesBaseWithHost(hostName string) string {
	sn := hex.EncodeToString(sha256.New().Sum([]byte(hostName)))
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test" {
  display_name = %q
  serial_number = %q
}
`, hostName, sn)
}

func testAccServicesBasicConfig(hostName, serviceName, serviceType string) string {
	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
}
`, serviceName, serviceType),
	}, "")
}

func testAccServicesDescription(hostName, serviceName, serviceType, description string) string {
	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test_description" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
    description = %q
}
`, serviceName, serviceType, description),
	}, "")
}

func testAccServicesDesiredState(hostName, serviceName, serviceType, desiredState string) string {
	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test_desired_state" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
    desired_state = %q
}
`, serviceName, serviceType, desiredState),
	}, "")
}

func testAccServicesDesiredVersion(hostName, serviceName, serviceType, desiredVersion string) string {
	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test_desired_version" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
    desired_version = %q
}
`, serviceName, serviceType, desiredVersion),
	}, "")
}

func testAccServicesInterfaceLabels(hostName, serviceName, serviceType string, interfaceLabels []string) string {
	interfaceLabelsBlock := strings.Builder{}
	for _, l := range interfaceLabels {
		interfaceLabelsBlock.WriteString(fmt.Sprintf("%q,", l))
	}

	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test_interface_labels" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
    interface_labels = [%s]
}
`, serviceName, serviceType, interfaceLabelsBlock.String()),
	}, "")
}

func testAccServicesTags(hostName, serviceName, serviceType, tags string) string {
	return strings.Join([]string{
		testAccServicesBaseWithHost(hostName),
		fmt.Sprintf(`
resource "bloxone_infra_service" "test_tags" {
    name = %q
    pool_id = bloxone_infra_host.test.pool_id
    service_type = %q
	wait_for_state = false
    tags = {
		tag1 = %q
	}
}
`, serviceName, serviceType, tags),
	}, "")
}
