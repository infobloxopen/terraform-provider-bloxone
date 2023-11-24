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

func TestAccForwardNsgResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgBasicConfig("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Test fields with default value
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
				Config: testAccForwardNsgBasicConfig("NAME_REPLACE_ME"),
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
				Config: testAccForwardNsgComment("NAME_REPLACE_ME", "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgComment("NAME_REPLACE_ME", "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
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
				Config: testAccForwardNsgExternalForwarders("NAME_REPLACE_ME", "EXTERNAL_FORWARDERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders", "EXTERNAL_FORWARDERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgExternalForwarders("NAME_REPLACE_ME", "EXTERNAL_FORWARDERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders", "EXTERNAL_FORWARDERS_UPDATE_REPLACE_ME"),
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
				Config: testAccForwardNsgForwardersOnly("NAME_REPLACE_ME", "FORWARDERS_ONLY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "FORWARDERS_ONLY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgForwardersOnly("NAME_REPLACE_ME", "FORWARDERS_ONLY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "FORWARDERS_ONLY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Hosts(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_hosts"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgHosts("NAME_REPLACE_ME", "HOSTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts", "HOSTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgHosts("NAME_REPLACE_ME", "HOSTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts", "HOSTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_InternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_internal_forwarders"
	var v dns_config.ConfigForwardNSG

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgInternalForwarders("NAME_REPLACE_ME", "INTERNAL_FORWARDERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forwarders", "INTERNAL_FORWARDERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgInternalForwarders("NAME_REPLACE_ME", "INTERNAL_FORWARDERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forwarders", "INTERNAL_FORWARDERS_UPDATE_REPLACE_ME"),
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
				Config: testAccForwardNsgName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgName("NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
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
				Config: testAccForwardNsgNsgs("NAME_REPLACE_ME", "NSGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs", "NSGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgNsgs("NAME_REPLACE_ME", "NSGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs", "NSGS_UPDATE_REPLACE_ME"),
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
				Config: testAccForwardNsgTags("NAME_REPLACE_ME", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgTags("NAME_REPLACE_ME", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
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

func testAccForwardNsgComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccForwardNsgExternalForwarders(name string, externalForwarders string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_external_forwarders" {
    name = %q
    external_forwarders = %q
}
`, name, externalForwarders)
}

func testAccForwardNsgForwardersOnly(name string, forwardersOnly string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_forwarders_only" {
    name = %q
    forwarders_only = %q
}
`, name, forwardersOnly)
}

func testAccForwardNsgHosts(name string, hosts string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_hosts" {
    name = %q
    hosts = %q
}
`, name, hosts)
}

func testAccForwardNsgInternalForwarders(name string, internalForwarders string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_internal_forwarders" {
    name = %q
    internal_forwarders = %q
}
`, name, internalForwarders)
}

func testAccForwardNsgName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_name" {
    name = %q
}
`, name)
}

func testAccForwardNsgNsgs(name string, nsgs string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_nsgs" {
    name = %q
    nsgs = %q
}
`, name, nsgs)
}

func testAccForwardNsgTags(name string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_tags" {
    name = %q
    tags = %q
}
`, name, tags)
}
