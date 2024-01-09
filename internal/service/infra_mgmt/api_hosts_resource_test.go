package infra_mgmt_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/infra_mgmt"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

// TODO: add tests
// The following require additional resource/data source objects to be supported.
// location
// pool_id - Creating pools is not supported for all accounts

func TestAccHostsResource_basic(t *testing.T) {
	var resourceName = "bloxone_infra_host.test"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "display_name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "legacy_id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "maintenance_mode", "disabled"),
					resource.TestCheckResourceAttr(resourceName, "timezone", "UTC"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHostsResource_disappears(t *testing.T) {
	resourceName := "bloxone_infra_host.test"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckHostsDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccHostsBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					testAccCheckHostsDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccHostsResource_Description(t *testing.T) {
	var resourceName = "bloxone_infra_host.test_description"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsDescription(name, "some description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "some description"),
				),
			},
			// Update and Read
			{
				Config: testAccHostsDescription(name, "some updated description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "some updated description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHostsResource_IpSpace(t *testing.T) {
	var resourceName = "bloxone_infra_host.test_ip_space"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsIpSpace(name, "space_1", acctest.RandomNameWithPrefix("ip-space")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.space_1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccHostsIpSpace(name, "space_2", acctest.RandomNameWithPrefix("ip-space")),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "ip_space", "bloxone_ipam_ip_space.space_2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHostsResource_MaintenanceMode(t *testing.T) {
	var resourceName = "bloxone_infra_host.test_maintenance_mode"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsMaintenanceMode(name, "enabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintenance_mode", "enabled"),
				),
			},
			// Update and Read
			{
				Config: testAccHostsMaintenanceMode(name, "disabled"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "maintenance_mode", "disabled"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHostsResource_SerialNumber(t *testing.T) {
	var resourceName = "bloxone_infra_host.test_serial_number"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsSerialNumber(name, "abcd"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "abcd"),
				),
			},
			// Update and Read
			{
				Config: testAccHostsSerialNumber(name, "xyzf"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "serial_number", "xyzf"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccHostsResource_Tags(t *testing.T) {
	var resourceName = "bloxone_infra_host.test_tags"
	var v infra_mgmt.InfraHost
	name := acctest.RandomNameWithPrefix("host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccHostsTags(name, "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccHostsTags(name, "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHostsExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckHostsExists(ctx context.Context, resourceName string, v *infra_mgmt.InfraHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.InfraManagementAPI.
			HostsAPI.
			HostsRead(ctx, rs.Primary.ID).
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

func testAccCheckHostsDestroy(ctx context.Context, v *infra_mgmt.InfraHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.InfraManagementAPI.
			HostsAPI.
			HostsRead(ctx, *v.Id).
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

func testAccCheckHostsDisappears(ctx context.Context, v *infra_mgmt.InfraHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.InfraManagementAPI.
			HostsAPI.
			HostsDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccHostsBasicConfig(displayName string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test" {
    display_name = %q
}
`, displayName)
}

func testAccHostsDescription(displayName, description string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test_description" {
    display_name = %q
    description = %q
}
`, displayName, description)
}

func testAccHostsIpSpace(displayName, ipSpaceResourceName, ipSpace string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_ip_space" %[1]s {
	name = %q
}

resource "bloxone_infra_host" "test_ip_space" {
    display_name = %q
    ip_space = bloxone_ipam_ip_space.%[1]s.id
}
`, ipSpaceResourceName, ipSpace, displayName)
}

func testAccHostsMaintenanceMode(displayName, maintenanceMode string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test_maintenance_mode" {
    display_name = %q
    maintenance_mode = %q
}
`, displayName, maintenanceMode)
}

func testAccHostsSerialNumber(displayName, serialNumber string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test_serial_number" {
    display_name = %q
    serial_number = %q
}
`, displayName, serialNumber)
}

func testAccHostsTags(displayName, tagValue string) string {
	return fmt.Sprintf(`
resource "bloxone_infra_host" "test_tags" {
    display_name = %q
    tags = {
		tag1 = %q
	}
}
`, displayName, tagValue)
}
