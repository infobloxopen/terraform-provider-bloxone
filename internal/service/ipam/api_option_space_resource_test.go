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

func TestAccOptionSpaceResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_space.test"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionSpaceBasicConfig("option_space_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "option_space_test"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
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

func TestAccOptionSpaceResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_option_space.test"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionSpaceDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionSpaceBasicConfig("option_space_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					testAccCheckOptionSpaceDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOptionSpaceResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_space.test_comment"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionSpaceComment("option_space_test", "ip4", "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionSpaceComment("option_space_test", "ip4", "test comment update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionSpaceResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_space.test_name"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionSpaceName("option_space_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "option_space_test"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionSpaceName("option_space_test_1", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "option_space_test_1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionSpaceResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_space.test_tags"
	var v ipam.IpamsvcOptionSpace

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionSpaceTags("option_space_test", "ip4", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccOptionSpaceTags("option_space_test", "ip4", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionSpaceExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOptionSpaceExists(ctx context.Context, resourceName string, v *ipam.IpamsvcOptionSpace) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionSpaceAPI.
			OptionSpaceRead(ctx, rs.Primary.ID).
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

func testAccCheckOptionSpaceDestroy(ctx context.Context, v *ipam.IpamsvcOptionSpace) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionSpaceAPI.
			OptionSpaceRead(ctx, *v.Id).
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

func testAccCheckOptionSpaceDisappears(ctx context.Context, v *ipam.IpamsvcOptionSpace) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionSpaceAPI.
			OptionSpaceDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOptionSpaceBasicConfig(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test" {
    name = %q
    protocol = %q
}
`, name, protocol)
}

func testAccOptionSpaceComment(name, protocol, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test_comment" {
    name = %q
    protocol = %q
    comment = %q
}
`, name, protocol, comment)
}

func testAccOptionSpaceName(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test_name" {
    name = %q
    protocol = %q
}
`, name, protocol)
}

func testAccOptionSpaceTags(name, protocol string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_space" "test_tags" {
    name = %q
    protocol = %q
    tags = %s
}
`, name, protocol, tagsStr)
}
