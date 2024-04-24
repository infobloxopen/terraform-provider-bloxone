package anycast_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"net/http"
	"strconv"
	"testing"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOnPremAnycastHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostBasicConfig("anycast1", "DHCP", "10.0.0.7"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "account_id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_disappears(t *testing.T) {
	resourceName := "bloxone_anycast_ac_config.test"
	var v anycast.ProtoOnpremHost

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnPremAnycastHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOnPremAnycastHostBasicConfig("anycast_test", "DHCP", "10.0.0.7"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					testAccCheckOnPremAnycastHostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOnPremAnycastHostResource_AnycastIpAddress(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_anycast_ip_address"
	var v anycast.ProtoOnpremHost
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostAnycastIpAddress("10.0.0.2", anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "anycast_ip_address", "10.0.0.2"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostAnycastIpAddress("10.0.0.3", anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "anycast_ip_address", "10.0.0.3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_Description(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_description"
	var v anycast.ProtoOnpremHost
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostDescription("10.0.0.2", anycastName, "DNS", "Anycast comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Anycast comment"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostDescription("10.0.0.2", anycastName, "DNS", "Anycast comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Anycast comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_Name(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_name"
	var v anycast.ProtoOnpremHost
	anycastName1 := acctest.RandomNameWithPrefix("anycast")
	anycastName2 := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostName("10.0.0.1", anycastName1, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName1),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostName("10.0.0.1", anycastName2, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_OnpremHosts(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_onprem_hosts"
	var v anycast.ProtoOnpremHost
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostOnpremHosts("10.0.0.1", anycastName, "DNS", "anycastHost1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostOnpremHosts("10.0.0.1", anycastName, "DNS", "anycastHost2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_Service(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_service"
	var v anycast.ProtoOnpremHost
	anycastName := acctest.RandomNameWithPrefix("anycast")
	anycastIP := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostService(anycastIP, anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DHCP"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostService(anycastIP, anycastName, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastHostResource_Tags(t *testing.T) {
	var resourceName = "bloxone_anycast_ac_config.test_tags"
	var v anycast.ProtoOnpremHost
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastHostTags("10.0.0.1", anycastName, "DNS", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastHostTags("10.0.0.1", anycastName, "DNS", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastHostExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOnPremAnycastHostExists(ctx context.Context, resourceName string, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return fmt.Errorf("error parsing ID: %v", err)
		}
		apiRes, _, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerGetOnpremHost(ctx, id). //OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus
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

func testAccCheckOnPremAnycastHostDestroy(ctx context.Context, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerGetOnpremHost(ctx, *v.Id). //OnPremAnycastManagerReadAnycastConfigWithRuntimeStatus
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

func testAccCheckOnPremAnycastHostDisappears(ctx context.Context, v *anycast.ProtoOnpremHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, _, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			OnPremAnycastManagerDeleteOnpremHost(ctx, *v.Id). //testAccCheckOnPremAnycastHostDisappears
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOnPremAnycastHostBasicConfig(name, service, anycastIpAddress string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
data "bloxone_infra_services" "anycast_services" {
    filters = {
      service_type = "anycast"
    }
}

data "bloxone_infra_hosts" "anycast_hosts" {
    filters = {
      pool_id = data.bloxone_infra_services.anycast_services.results.0.pool_id
    }
}

resource "bloxone_anycast_host "test" {
    name = %q
    service= %q
    anycast_ip_address = %q
}
`, name, service, anycastIpAddress)
}

func testAccOnPremAnycastHostAnycastIpAddress(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_ac_config" "test_anycast_ip_address" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastHostDescription(anycastIpAddress, name, service, description string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_ac_config" "test_description" {
    anycast_ip_address = %q
    name = %q
    service = %q
    description = %q
}
`, anycastIpAddress, name, service, description)
}

func testAccOnPremAnycastHostName(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_ac_config" "test_name" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastHostOnpremHosts(anycastIpAddress, name, service, opHost string) string {
	return fmt.Sprintf(`
data "bloxone_infra_services" "anycast_services" {
    filters = {
      service_type = "anycast"
    }
}

data "bloxone_infra_hosts" "anycast_hosts" {
    filters = {
      pool_id = data.bloxone_infra_services.anycast_services.results.0.pool_id
    }
}

resource "bloxone_anycast_ac_config" "test_onprem_hosts" {
    anycast_ip_address = %q
    name = %q
    service = %q
    onprem_hosts = [
	{
		id = data.bloxone_infra_hosts.anycast_hosts.results.0.legacy_id
		name = data.bloxone_infra_hosts.anycast_hosts.results.0.display_name
	}
	]
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastHostService(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_ac_config" "test_service" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastHostTags(anycastIpAddress, name, service string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_anycast_ac_config" "test_tags" {
    anycast_ip_address = %q
    name = %q
    service = %q
    tags = %s
}
`, anycastIpAddress, name, service, tagsStr)
}
