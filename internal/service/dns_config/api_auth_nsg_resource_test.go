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

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - internal_secondaries
// - TSIG keys

func TestAccAuthNsgResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_auth_nsg.test"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAuthNsgDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAuthNsgBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					testAccCheckAuthNsgDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAuthNsgResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_comment"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgComment(name, "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgComment(name, "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_external_primaries"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")
	nsg1 := acctest.RandomNameWithPrefix("auth-nsg")
	nsg2 := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgExternalPrimaries(name, "1.1.1.1", "a.com.", nsg1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "a.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
					resource.TestCheckResourceAttrPair(resourceName, "external_primaries.1.nsg", "bloxone_dns_auth_nsg.one", "id"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.1.type", "nsg"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgExternalPrimaries(name, "2.2.2.2", "b.com.", nsg2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.fqdn", "b.com."),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.type", "primary"),
					resource.TestCheckResourceAttrPair(resourceName, "external_primaries.1.nsg", "bloxone_dns_auth_nsg.one", "id"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.1.type", "nsg"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_external_secondaries"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgExternalSecondaries(name, "1.1.1.1", "a.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "1.1.1.1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.fqdn", "a.com."),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgExternalSecondaries(name, "2.2.2.2", "b.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "2.2.2.2"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.fqdn", "b.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_Name(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_name"
	var v1 dnsconfig.AuthNSG
	var v2 dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgName("nsg2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgDestroy(context.Background(), &v1),
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "name", "nsg2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_nsgs"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")
	nsg1 := acctest.RandomNameWithPrefix("auth-nsg")
	nsg2 := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgNsgs(name, nsg1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_auth_nsg.nsg1", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgNsgsUpdate(name, nsg1, nsg2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "nsgs.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_auth_nsg.nsg1", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.1", "bloxone_dns_auth_nsg.nsg2", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAuthNsgResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_auth_nsg.test_tags"
	var v dnsconfig.AuthNSG
	name := acctest.RandomNameWithPrefix("auth-nsg")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAuthNsgTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccAuthNsgTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAuthNsgExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAuthNsgExists(ctx context.Context, resourceName string, v *dnsconfig.AuthNSG) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthNsgAPI.
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

func testAccCheckAuthNsgDestroy(ctx context.Context, v *dnsconfig.AuthNSG) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthNsgAPI.
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

func testAccCheckAuthNsgDisappears(ctx context.Context, v *dnsconfig.AuthNSG) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AuthNsgAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAuthNsgBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test" {
    name = %q
}
`, name)
}

func testAccAuthNsgComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccAuthNsgExternalPrimaries(name string, address string, fqdn string, referencedNsgName string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "one" {
	name = %[4]q
}

resource "bloxone_dns_auth_nsg" "test_external_primaries" {
    name = %[1]q
    external_primaries = [
		{
			address = %[2]q,
			fqdn = %[3]q
			type = "primary"
		},
		{
			nsg = bloxone_dns_auth_nsg.one.id
			type = "nsg"
		}
	]
	lifecycle {
		create_before_destroy = true
	}
}
`, name, address, fqdn, referencedNsgName)
}

func testAccAuthNsgExternalSecondaries(name string, address string, fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test_external_secondaries" {
    name = %q
    external_secondaries = [
		{
			address = %q,
			fqdn = %q
		},
	]
}
`, name, address, fqdn)
}

func testAccAuthNsgName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test_name" {
    name = %q
}
`, name)
}

func testAccAuthNsgNsgs(name string, nsg1 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "nsg1" {
    name = %q
}
resource "bloxone_dns_auth_nsg" "test_nsgs" {
    name = %q
    nsgs = [bloxone_dns_auth_nsg.nsg1.id]
}
`, nsg1, name)
}

func testAccAuthNsgNsgsUpdate(name string, nsg1 string, nsg2 string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "nsg1" {
    name = %q
}
resource "bloxone_dns_auth_nsg" "nsg2" {
    name = %q
}
resource "bloxone_dns_auth_nsg" "test_nsgs" {
    name = %q
    nsgs = [
		bloxone_dns_auth_nsg.nsg1.id,
		bloxone_dns_auth_nsg.nsg2.id
	]
}
`, nsg1, nsg2, name)
}

func testAccAuthNsgTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_auth_nsg" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
