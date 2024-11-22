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

	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccDelegationResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test"
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationBasicConfig(viewName),
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

	t.Skip("Test Skipped due to inconsistent error codes returned by the API [NORTHSTAR-12575]")

	resourceName := "bloxone_dns_delegation.test"
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckDelegationDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccDelegationBasicConfig(viewName),
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
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationComment(viewName, "Delegation zone is created by Terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "Delegation zone is created by Terraform"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationComment(viewName, "Delegation zone was created by Terraform"),
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
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationDelegationServers(viewName, "12.0.0.0", "ns1.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.address", "12.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.fqdn", "ns1.com."),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDelegationServers(viewName, "12.0.0.1", "ns2.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.address", "12.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.fqdn", "ns2.com."),
				),
			},
			// fqdn with no address
			{
				Config: testAccDelegationDelegationServers(viewName, "", "ns3.com."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.address", ""),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers.0.fqdn", "ns3.com."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Disabled(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_disabled"
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationDisabled(viewName, "test.123.", "false", "123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDisabled(viewName, "test.123.", "true", "123."),
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
	var v1, v2 dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationFqdn(viewName, "test.123."),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "test.123."),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationFqdn(viewName, "test1.123."),
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
	var v dnsconfig.Delegation
	viewName := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationTags(viewName, map[string]string{
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
				Config: testAccDelegationTags(viewName, map[string]string{
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
	var v1, v2 dnsconfig.Delegation
	view1 := acctest.RandomNameWithPrefix("view")
	view2 := acctest.RandomNameWithPrefix("view")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationView(view1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view."+view1, "id"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationView(view2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationDestroy(context.Background(), &v1),
					testAccCheckDelegationExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttrPair(resourceName, "view", "bloxone_dns_view."+view2, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckDelegationExists(ctx context.Context, resourceName string, v *dnsconfig.Delegation) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
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

func testAccCheckDelegationDestroy(ctx context.Context, v *dnsconfig.Delegation) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
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

func testAccCheckDelegationDisappears(ctx context.Context, v *dnsconfig.Delegation) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			DelegationAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithViewAndAuthZone(viewName string) string {
	return fmt.Sprintf(`
	resource "bloxone_dns_view" "test" {
		name = %q
	}

	resource "bloxone_dns_auth_zone" test {
		fqdn = "123."
		primary_type = "cloud"
		view = bloxone_dns_view.test.id
		depends_on = [bloxone_dns_view.test]
	}`, viewName)
}

func testAccDelegationBasicConfig(viewName string) string {
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
	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationComment(viewName string, comment string) string {
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
	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationDelegationServers(viewName, delegationServersAddrs, delegationServersFqdn string) string {
	addressStr := ""
	if delegationServersAddrs != "" {
		addressStr = fmt.Sprintf(`"address": %q`, delegationServersAddrs)
	}

	config := fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_delegation_servers" {
    fqdn = "test.123."
    delegation_servers = [{
		%s
		"fqdn": %q
	}]
    view = bloxone_dns_view.test.id
    depends_on = [bloxone_dns_view.test, bloxone_dns_auth_zone.test]
}
`, addressStr, delegationServersFqdn)
	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationDisabled(viewName, fqdn, disabled, authFqdn string) string {
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
	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationFqdn(viewName, fqdn string) string {
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

	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationTags(viewName string, tags map[string]string) string {
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
	return strings.Join([]string{testAccBaseWithViewAndAuthZone(viewName), config}, "")
}

func testAccDelegationView(view string) string {
	return fmt.Sprintf(`
	resource "bloxone_dns_view" %[1]q {
     name = %[1]q
    }

	resource "bloxone_dns_auth_zone" test {
     fqdn = "123."
     primary_type = "cloud"
     view = bloxone_dns_view.%[1]s.id
	}

resource "bloxone_dns_delegation" "test_view" {
    fqdn = "test.123."
    view = bloxone_dns_view.%[1]s.id
    delegation_servers = [{
		address = "12.0.0.0"
		fqdn = "ns1.com."
  	}]
    depends_on = [bloxone_dns_auth_zone.test]
}
`, view)
}
