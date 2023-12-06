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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneBasicConfig(fqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckForwardZoneDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccForwardZoneBasicConfig(fqdn),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v1 dns_config.ConfigForwardZone
	var v2 dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneBasicConfig(fqdn),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", fqdn),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneComment(fqdn, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneComment(fqdn, "test comment update"),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneDisabled(fqdn, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneDisabled(fqdn, "false"),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneExternalForwarders(fqdn, "192.168.10.10", "tf-infoblox-test.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.address", "192.168.10.10"),
					resource.TestCheckResourceAttr(resourceName, "external_forwarders.0.fqdn", "tf-infoblox-test.com."),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneExternalForwarders(fqdn, "192.168.11.11", "tf-infoblox.com."),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneForwardOnly(fqdn, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "forward_only", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneForwardOnly(fqdn, "true"),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneHosts(fqdn, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.0", "data.bloxone_dns_hosts.one", "results.0.id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneHosts(fqdn, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "hosts.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "hosts.0", "data.bloxone_dns_hosts.two", "results.0.id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_InternalForwarders(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_internal_forwarders"
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneInternalForwarders(fqdn, "one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forwarders.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "internal_forwarders.0", "data.bloxone_dns_hosts.one", "results.0.id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneInternalForwarders(fqdn, "two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "internal_forwarders.#", "1"),
					resource.TestCheckResourceAttrPair(resourceName, "internal_forwarders.0", "data.bloxone_dns_hosts.two", "results.0.id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccForwardZoneResource_Nsgs(t *testing.T) {
	var resourceName = "bloxone_dns_forward_zone.test_nsgs"
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneNsgs(fqdn, "bloxone_dns_forward_nsg.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "nsgs.0", "bloxone_dns_forward_nsg.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneNsgs(fqdn, "bloxone_dns_forward_nsg.two"),
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneTags(fqdn, map[string]string{
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
				Config: testAccForwardZoneTags(fqdn, map[string]string{
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
	var fqdn = acctest.RandomNameWithPrefix("fw-zone") + ".com."
	var v1 dns_config.ConfigForwardZone
	var v2 dns_config.ConfigForwardZone

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccForwardZoneView(fqdn, "bloxone_dns_view.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckForwardZoneExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccForwardZoneView(fqdn, "bloxone_dns_view.two"),
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
	return `
data "bloxone_dns_hosts" "one" {
    filters = {
        name = "TF_TEST_HOST_01"
    }
}
data "bloxone_dns_hosts" "two" {
    filters = {
        name = "TF_TEST_HOST_02"
    }
}
    `
}

func testAccForwardZoneHosts(fqdn, host string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_hosts" {
    fqdn = %q
    hosts = [data.bloxone_dns_hosts.%s.results.0.id]
}
`, fqdn, host)
	return strings.Join([]string{testAccBaseWithHost(), config}, "")
}

func testAccForwardZoneInternalForwarders(fqdn, host string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_forward_zone" "test_internal_forwarders" {
    fqdn = %q
    internal_forwarders = [data.bloxone_dns_hosts.%s.results.0.id]
}
`, fqdn, host)
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
`, acctest.RandomNameWithPrefix("fw-nsg"), acctest.RandomNameWithPrefix("fw-nsg"), fqdn, nsgs)
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
