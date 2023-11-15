package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO:add tests
// The following require additional resource/data source objects to be supported.
// - auto_generate_records
// - addresses
// - host_namesFi

// Test case
func TestAccIpamHostResource_basic(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpamHost_disappears(t *testing.T) {
	resourceName := "bloxone_ipam_host.test"
	var v ipam.IpamsvcIpamHost
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckIpamHostDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccIpamHostBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIpamHostExists(context.Background(), resourceName, &v),
					testAccCheckIpamHostDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccIpamHostResource_Comment(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test_comment"
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostComment(name, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccIpamHostComment(name, "test comment updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccIpamHostResource_Tags(t *testing.T) {
	var resourceName = "bloxone_ipam_host.test_tags"
	var name = acctest.RandomNameWithPrefix("ipam_host")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccIpamHostTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccIpamHostTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckIpamHostExists(ctx context.Context, resourceName string, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostRead(ctx, rs.Primary.ID).
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

func testAccCheckIpamHostDestroy(ctx context.Context, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostRead(ctx, *v.Id).
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

func testAccCheckIpamHostDisappears(ctx context.Context, v *ipam.IpamsvcIpamHost) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccIpamHostBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test" {
    name = "%s"
}
`, name)
}

func testAccIpamHostComment(name, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test_comment" {
    name = "%s"
    comment = "%s"
}
`, name, comment)
}

func testAccIpamHostTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s=%q
`, k, v)
	}
	tagsStr += "\t}"
	return fmt.Sprintf(`
resource "bloxone_ipam_host" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
