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

func TestAccDelegationResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationBasicConfig("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol_fqdn"),
					// Test fields with default value
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
				Config: testAccDelegationBasicConfig("FQDN_REPLACE_ME"),
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
				Config: testAccDelegationComment("FQDN_REPLACE_ME", "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationComment("FQDN_REPLACE_ME", "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
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
				Config: testAccDelegationDelegationServers("FQDN_REPLACE_ME", "DELEGATION_SERVERS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers", "DELEGATION_SERVERS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDelegationServers("FQDN_REPLACE_ME", "DELEGATION_SERVERS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "delegation_servers", "DELEGATION_SERVERS_UPDATE_REPLACE_ME"),
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
				Config: testAccDelegationDisabled("FQDN_REPLACE_ME", "DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationDisabled("FQDN_REPLACE_ME", "DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disabled", "DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Fqdn(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_fqdn"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_Parent(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_parent"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationParent("FQDN_REPLACE_ME", "PARENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parent", "PARENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationParent("FQDN_REPLACE_ME", "PARENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parent", "PARENT_UPDATE_REPLACE_ME"),
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
				Config: testAccDelegationTags("FQDN_REPLACE_ME", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationTags("FQDN_REPLACE_ME", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccDelegationResource_View(t *testing.T) {
	var resourceName = "bloxone_dns_delegation.test_view"
	var v dns_config.ConfigDelegation

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccDelegationView("FQDN_REPLACE_ME", "VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccDelegationView("FQDN_REPLACE_ME", "VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDelegationExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
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

func testAccDelegationBasicConfig(fqdn string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test" {
    fqdn = %q
}
`, fqdn)
}

func testAccDelegationComment(fqdn string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_comment" {
    fqdn = %q
    comment = %q
}
`, fqdn, comment)
}

func testAccDelegationDelegationServers(fqdn string, delegationServers string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_delegation_servers" {
    fqdn = %q
    delegation_servers = %q
}
`, fqdn, delegationServers)
}

func testAccDelegationDisabled(fqdn string, disabled string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_disabled" {
    fqdn = %q
    disabled = %q
}
`, fqdn, disabled)
}

func testAccDelegationFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}

func testAccDelegationParent(fqdn string, parent string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_parent" {
    fqdn = %q
    parent = %q
}
`, fqdn, parent)
}

func testAccDelegationTags(fqdn string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_tags" {
    fqdn = %q
    tags = %q
}
`, fqdn, tags)
}

func testAccDelegationView(fqdn string, view string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_delegation" "test_view" {
    fqdn = %q
    view = %q
}
`, fqdn, view)
}
