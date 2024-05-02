package anycast_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOnPremAnycastManagerResource_basic(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerBasicConfig(anycastName, "DHCP", "10.0.0.8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
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

func TestAccOnPremAnycastManagerResource_disappears(t *testing.T) {
	resourceName := "bloxone_anycast_config.test"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOnPremAnycastManagerDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOnPremAnycastManagerBasicConfig(anycastName, "DHCP", "10.0.0.7"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					testAccCheckOnPremAnycastManagerDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOnPremAnycastManagerResource_AnycastIpAddress(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test_anycast_ip_address"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerAnycastIpAddress("10.0.0.2", anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "anycast_ip_address", "10.0.0.2"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastManagerAnycastIpAddress("10.0.0.3", anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName),
					resource.TestCheckResourceAttr(resourceName, "anycast_ip_address", "10.0.0.3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastManagerResource_Description(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test_description"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerDescription("10.0.0.2", anycastName, "DNS", "Anycast comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Anycast comment"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastManagerDescription("10.0.0.2", anycastName, "DNS", "Anycast comment updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "Anycast comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastManagerResource_Name(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test_name"
	var v anycast.AnycastConfig
	anycastName1 := acctest.RandomNameWithPrefix("anycast")
	anycastName2 := acctest.RandomNameWithPrefix("anycast")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerName("10.0.0.1", anycastName1, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName1),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastManagerName("10.0.0.1", anycastName2, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", anycastName2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastManagerResource_Service(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test_service"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")
	anycastIP := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerService(anycastIP, anycastName, "DHCP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DHCP"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastManagerService(anycastIP, anycastName, "DNS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service", "DNS"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOnPremAnycastManagerResource_Tags(t *testing.T) {
	var resourceName = "bloxone_anycast_config.test_tags"
	var v anycast.AnycastConfig
	anycastName := acctest.RandomNameWithPrefix("anycast")
	anycastIP := acctest.RandomIP()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOnPremAnycastManagerTags(anycastIP, anycastName, "DNS", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccOnPremAnycastManagerTags(anycastIP, anycastName, "DNS", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOnPremAnycastManagerExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOnPremAnycastManagerExists(ctx context.Context, resourceName string, v *anycast.AnycastConfig) resource.TestCheckFunc {
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
			ReadAnycastConfigWithRuntimeStatus(ctx, id).
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

func testAccCheckOnPremAnycastManagerDestroy(ctx context.Context, v *anycast.AnycastConfig) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			ReadAnycastConfigWithRuntimeStatus(ctx, *v.Id).
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

func testAccCheckOnPremAnycastManagerDisappears(ctx context.Context, v *anycast.AnycastConfig) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, _, err := acctest.BloxOneClient.AnycastAPI.
			OnPremAnycastManagerAPI.
			DeleteAnycastConfig(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOnPremAnycastManagerBasicConfig(name, service, anycastIpAddress string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test" {
    name = %q
    service= %q
    anycast_ip_address = %q
}
`, name, service, anycastIpAddress)
}

func testAccOnPremAnycastManagerAnycastIpAddress(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_anycast_ip_address" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastManagerDescription(anycastIpAddress, name, service, description string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_description" {
    anycast_ip_address = %q
    name = %q
    service = %q
    description = %q
}
`, anycastIpAddress, name, service, description)
}

func testAccOnPremAnycastManagerName(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_name" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastManagerService(anycastIpAddress, name, service string) string {
	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_service" {
    anycast_ip_address = %q
    name = %q
    service = %q
}
`, anycastIpAddress, name, service)
}

func testAccOnPremAnycastManagerTags(anycastIpAddress, name, service string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_anycast_config" "test_tags" {
    anycast_ip_address = %q
    name = %q
    service = %q
    tags = %s
}
`, anycastIpAddress, name, service, tagsStr)
}
