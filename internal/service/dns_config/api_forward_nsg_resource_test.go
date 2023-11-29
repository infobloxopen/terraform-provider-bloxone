package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

/* TODO: Add tests for the following
   - hosts
   - internal forwarders
*/

func TestAccForwardNsgResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgBasicConfig("terraform-acc-test.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "terraform-acc-test.com"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_forward_nsg.test"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardNsgBasicConfig("terraform-acc-test.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					testAccCheckForwardNsgDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccForwardNsgResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_comment"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgComment("terraform-acc-test.com", "This Forward zone is created through Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This Forward zone is created through Terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgComment("terraform-acc-test.com", "This Forward zone was created through Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This Forward zone was created through Terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_ExternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_external_forwarders"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgExternalForwarders("terraform-acc-test.com", "192.168.1.0", "terraform-acc-forward-ext."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "terraform-acc-forward-ext."),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.0"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgExternalForwarders("terraform-acc-test.com", "192.168.1.1", "terraform-acc-forward-ext-1."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "terraform-acc-forward-ext-1."),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_ForwardersOnly(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_forwarders_only"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgForwardersOnly("terraform-acc-test.com", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgForwardersOnly("terraform-acc-test.com", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Name(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_name"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgName("terraform-acc-test.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "terraform-acc-test.com"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgName("test.change.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "test.change.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_nsgs"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgNsgs("terraform-acc-test-1.com", "test-nsg"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg.nsg", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgNsgs("terraform-acc-test-1.com", "test-nsg-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg.nsg", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_tags"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgTags("terraform-acc-test.com", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgTags("terraform-acc-test.com", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckForwardNsgExists(ctx context.Context, resourceName string, v *dns_config.ConfigForwardNSG) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
			ForwardNsgRead(ctx, rs.Primary.ID).
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

func testAccCheckForwardNsgDestroy(ctx context.Context, v *dns_config.ConfigForwardNSG) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
			ForwardNsgRead(ctx, *v.Id).
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

func testAccCheckForwardNsgDisappears(ctx context.Context, v *dns_config.ConfigForwardNSG) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
			ForwardNsgDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccForwardNsgBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test" {
    name = %q
}
`, name)
}

func testAccForwardNsgComment(name, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccForwardNsgExternalForwarders(name, address, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_external_forwarders" {
    name = %q
    external_forwarders = [{
		address = %q
		fqdn = %q
    }]
}
`, name, address, fqdn)
}

func testAccForwardNsgForwardersOnly(name, forwardersOnly string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_forwarders_only" {
    name = %q
    forwarders_only = %q
}
`, name, forwardersOnly)
}

func testAccForwardNsgName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_name" {
    name = %q
}
`, name)
}

func testAccForwardNsgNsgs(name, nsgs string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "nsg" {
    name = %q
}

resource "bloxone_dns_forward_nsg" "test_nsgs" {
    name = %q
    nsgs = [bloxone_dns_forward_nsg.nsg.id]
}
`, nsgs, name)
}

func testAccForwardNsgTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
