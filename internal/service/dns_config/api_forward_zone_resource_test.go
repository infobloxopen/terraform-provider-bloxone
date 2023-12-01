package dns_config_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

//TODO: add tests
// The following require additional resource/data source objects to be supported.
// - internal_forwarders
// - hosts

func TestAccForwardZoneResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneBasicConfig("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "mapping"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol_fqdn"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_forward_zone.test"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardZoneBasicConfig("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					testAccCheckForwardZoneDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccForwardZoneResource_FQDN(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test"
	var v1 dns_config.ConfigForwardZone
	var v2 dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneBasicConfig("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneBasicConfig("tf-infoblox-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneDestroy(context.Background(), &v1),
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-infoblox-test.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_comment"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneComment("tf-acc-test.com.", "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneComment("tf-acc-test.com.", "test comment update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_disabled"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneDisabled("tf-acc-test.com.", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneDisabled("tf-acc-test.com.", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_ExternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_external_forwarders"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneExternalForwarders("tf-acc-test.com.", "192.168.10.10", "tf-infoblox-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "tf-infoblox-test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneExternalForwarders("tf-acc-test.com.", "192.168.11.11", "tf-infoblox.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.11.11"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "tf-infoblox.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_ForwardOnly(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_forward_only"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneForwardOnly("tf-acc-test.com.", "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "forward_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneForwardOnly("tf-acc-test.com.", "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "forward_only", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Hosts(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_hosts"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneHosts("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneHosts("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_InternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_internal_forwarders"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneInternalForwarders("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneInternalForwarders("tf-acc-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "tf-acc-test.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_nsgs"
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneNsgs("tf-acc-test.com.", "bloxone_dns_forward_nsg.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneNsgs("tf-acc-test.com.", "bloxone_dns_forward_nsg.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_tags"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneTags("tf-acc-test.com.", map[string]string{
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
				Config: testAccForwardZoneTags("tf-acc-test.com.", map[string]string{
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

func TestAccForwardZoneResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_view"
	var v1 dns_config.ConfigForwardZone
	var v2 dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneView("tf-acc-test.com.", "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneView("tf-acc-test.com.", "bloxone_dns_view.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneDestroy(context.Background(), &v1),
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckForwardZoneExists(ctx context.Context, resourceName string, v *dns_config.ConfigForwardZone) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardZoneAPI.
			ForwardZoneRead(ctx, rs.Primary.ID).
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

func testAccCheckForwardZoneDestroy(ctx context.Context, v *dns_config.ConfigForwardZone) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardZoneAPI.
			ForwardZoneRead(ctx, *v.Id).
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

func testAccCheckForwardZoneDisappears(ctx context.Context, v *dns_config.ConfigForwardZone) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			ForwardZoneAPI.
			ForwardZoneDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccForwardZoneBasicConfig(fqdn string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test" {
    fqdn = %q
}
`, fqdn)
}

func testAccForwardZoneComment(fqdn, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_comment" {
    fqdn = %q
    comment = %q
}
`, fqdn, comment)
}

func testAccForwardZoneDisabled(fqdn, disabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_disabled" {
    fqdn = %q
    disabled = %q
}
`, fqdn, disabled)
}

func testAccForwardZoneExternalForwarders(fqdn, address, externalForwardersFQDN string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_external_forwarders" {
    fqdn = %q
    external_forwarders = [
		{
			address = %q
			fqdn = %q
		}
]
}
`, fqdn, address, externalForwardersFQDN)
}

func testAccForwardZoneForwardOnly(fqdn, forwardOnly string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_forward_only" {
    fqdn = %q
    forward_only = %q
}
`, fqdn, forwardOnly)
}
func testAccBaseWithHost() string {
	return fmt.Sprintf(`
data "bloxone_dns_hosts" "all_hosts" {
	filters = {
		name = "TF_TEST_HOST_01"
}
}
`)
}

func testAccForwardZoneHosts(fqdn string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_hosts" {
    fqdn = %q
    hosts = [data.bloxone_dns_hosts.all_hosts.results.0.id]
}
`, fqdn)
	return strings.Join([]string{testAccBaseWithHost(), config}, "")
}

func testAccForwardZoneInternalForwarders(fqdn string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_internal_forwarders" {
    fqdn = %q
    internal_forwarders = [data.bloxone_dns_hosts.all_hosts.results.0.id]
}
`, fqdn)
	return strings.Join([]string{testAccBaseWithHost(), config}, "")
}

func testAccForwardZoneNsgs(fqdn, nsgs string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_forward_nsg" "one"{
	name = %q
}

resource "bloxone_dns_forward_nsg" "two"{
	name = %q
}

resource "bloxone_dns_forward_zone" "test_nsgs" {
	fqdn = %q
    nsgs = [%s.id]
}
`, acctest.RandomNameWithPrefix("auth-nsg"), acctest.RandomNameWithPrefix("auth-nsg"), fqdn, nsgs)
}

func testAccForwardZoneTags(fqdn string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_tags" {
    fqdn = %q
    tags = %s
}
`, fqdn, tagsStr)
}

func testAccForwardZoneView(fqdn, view string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_view" "one" {
	name = %q
}

resource "bloxone_dns_view" "two" {
	name = %q
}

resource "bloxone_dns_forward_zone" "test_view" {
	fqdn = %q
    view = %s.id
}
`, acctest.RandomNameWithPrefix("view"), acctest.RandomNameWithPrefix("view"), fqdn, view)
}
