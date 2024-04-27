package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

/* TODO: Add tests for the following
   - hosts
   - internal forwarders
*/

func TestAccForwardNsgResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test"
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("name")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
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
	var v dnsconfig.ForwardNSG

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardNsgBasicConfig(acctest.RandomNameWithPrefix("nsg")),
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
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("nsg")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgComment(name, "This Forward NSG is created through Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This Forward NSG is created through Terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgComment(name, "This Forward NSG was created through Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This Forward NSG was created through Terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_ExternalForwarders_Address(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_external_forwarders_address"
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("nsg")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgExternalForwardersAddress(name, "192.168.1.0"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.0"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgExternalForwardersAddress(name, "192.168.1.1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_ExternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_external_forwarders_fqdn"
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("nsg")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgExternalForwardersFqdn(name, "192.168.1.0", "terraform-acc-forward-ext."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "terraform-acc-forward-ext."),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.0"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgExternalForwardersFqdn(name, "192.168.1.0", "terraform-acc-forward-ext-1."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "terraform-acc-forward-ext-1."),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.1.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_ForwardersOnly(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_forwarders_only"
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgForwardersOnly(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "forwarders_only", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgForwardersOnly(name, "false"),
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
	var v dnsconfig.ForwardNSG
	name1 := acctest.RandomNameWithPrefix("nsg")
	name2 := acctest.RandomNameWithPrefix("nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgName(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgName(name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_nsgs"
	var v dnsconfig.ForwardNSG
	name1 := acctest.RandomNameWithPrefix("nsg")
	name2 := acctest.RandomNameWithPrefix("nsg")
	name3 := acctest.RandomNameWithPrefix("nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgNsgs(name1, name3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg."+name1, "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardNsgNsgs(name2, name3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg."+name2, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardNsgResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_forward_nsg.test_tags"
	var v dnsconfig.ForwardNSG
	name := acctest.RandomNameWithPrefix("nsg")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardNsgTags(name, map[string]string{
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
				Config: testAccForwardNsgTags(name, map[string]string{
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

func testAccCheckForwardNsgExists(ctx context.Context, resourceName string, v *dnsconfig.ForwardNSG) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
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

func testAccCheckForwardNsgDestroy(ctx context.Context, v *dnsconfig.ForwardNSG) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
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

func testAccCheckForwardNsgDisappears(ctx context.Context, v *dnsconfig.ForwardNSG) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardNsgAPI.
			Delete(ctx, *v.Id).
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

func testAccForwardNsgExternalForwardersAddress(name, address string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_external_forwarders_address" {
    name = %q
    external_forwarders = [{
		address = %q
    }]
}
`, name, address)
}

func testAccForwardNsgExternalForwardersFqdn(name, address, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "test_external_forwarders_fqdn" {
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

func testAccForwardNsgNsgs(name1, name2 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" %[1]q {
    name = %[1]q

	lifecycle {
      // The old resource can't be destroyed unless it is unlinked, so create new resource and link it before destroying
	  create_before_destroy = true
	}
}

resource "bloxone_dns_forward_nsg" "test_nsgs" {
    name = %[2]q
    nsgs = [bloxone_dns_forward_nsg.%[1]s.id]
}
`, name1, name2)
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
