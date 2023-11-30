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

func TestAccDelegationResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "test.123."),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.address", "12.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.fqdn", "ns1.com."),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol_fqdn"),
					resource.TestCheckResourceAttrSet(resourceName, "parent"),
					resource.TestCheckResourceAttrPair(resourceName, "parent", "bloxone_dns_auth_zone.test", "id"),
					// Test fields with default value
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_delegation.test"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					testAccCheckDelegationDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccDelegationResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_comment"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationComment("Delegation zone is created by Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Delegation zone is created by Terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationComment("Delegation zone was created by Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Delegation zone was created by Terraform"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_DelegationServers(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_delegation_servers"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationDelegationServers("12.0.0.0", "ns1.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDelegationServers("12.0.0.1", "ns2.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_disabled"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationDisabled("test.123.", "false", "123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDisabled("test.123.", "true", "123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Fqdn(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_fqdn"
	var v1, v2 dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationFqdn("test.123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "test.123."),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationFqdn("test1.123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationDestroy(context.Background(), &v1),
					testAccCheckDelegationExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "test1.123."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_tags"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationTags(map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationTags(map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_view"
	var v1, v2 dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationView("bloxone_dns_view.test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationView("bloxone_dns_view.test1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationDestroy(context.Background(), &v1),
					testAccCheckDelegationExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view.test1", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDelegationExists(ctx context.Context, resourceName string, v *dns_config.ConfigDelegation) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
			DelegationRead(ctx, rs.Primary.ID).
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

func testAccCheckDelegationDestroy(ctx context.Context, v *dns_config.ConfigDelegation) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
			DelegationRead(ctx, *v.Id).
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

func testAccCheckDelegationDisappears(ctx context.Context, v *dns_config.ConfigDelegation) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
			DelegationDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccViewAndAuthZone() string {
	return `
	resource "bloxone_dns_view" "test" {
		name = "example-view"
	}

	resource "bloxone_dns_auth_zone" test {
		fqdn = "123."
		primary_type = "cloud"
		view = bloxone_dns_view.test.id
		depends_on = [bloxone_dns_view.test]
	}`
}

func testAccDelegationBasicConfig() string {
	// TODO: create basic resource with required fields
	config := `
resource "bloxone_dns_delegation" "test" {
    fqdn = "test.123."
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`
	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

func testAccDelegationComment(comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_comment" {
    fqdn = "test.123."
    comment = %q
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
    }]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, comment)
	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

// venkat
func testAccDelegationDelegationServers(delegationServersAddrs, delegationServersFqdn string) string {

	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_delegation_servers" {
    fqdn = "test.123."
    delegation_servers = [{
		"address": %q
		"fqdn": %q
	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, delegationServersAddrs, delegationServersFqdn)
	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

func testAccDelegationDisabled(fqdn, disabled, authFqdn string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_disabled" {
    fqdn = %q
    disabled = %q
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, fqdn, disabled)
	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

func testAccDelegationFqdn(fqdn string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_fqdn" {
    fqdn = %q
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, fqdn)

	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

func testAccDelegationTags(tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"
	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_tags" {
    fqdn = "test.123."
    tags = %s
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, tagsStr)
	return strings.Join([]string{testAccViewAndAuthZone(), config}, "")
}

func testAccDelegationView(view string) string {
	return fmt.Sprintf(`
	resource "bloxone_dns_view" "test" {
     name = "example-view"
    }

	resource "bloxone_dns_view" "test1" {
     name = "example-view-1"
    }

	resource "bloxone_dns_auth_zone" test {
     fqdn = "123."
     primary_type = "cloud"
     view = %s.id
	}
resource "bloxone_dns_delegation" "test_view" {
    fqdn = "test.123."
    view = %s.id
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, view, view)
}
