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

func TestAccAclResource_basic(t *testing.T) {
	var resourceName = "bloxone_dns_acl.test"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),

					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAclResource_disappears(t *testing.T) {
	resourceName := "bloxone_dns_acl.test"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAclDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccAclBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					testAccCheckAclDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccAclResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dns_acl.test_comment"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclComment(name, "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccAclComment(name, "updated test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated test comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAclResource_List(t *testing.T) {
	var resourceName = "bloxone_dns_acl.test_list"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclIP("acl", "list", name, "allow", "192.168.11.11"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list.0.access", "allow"),
					resource.TestCheckResourceAttr(resourceName, "list.0.element", "ip"),
					resource.TestCheckResourceAttr(resourceName, "list.0.address", "192.168.11.11")),
			},
			// Update and Read
			{
				Config: testAccAclAny("acl", "list", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "list.0.element", "any"),
				),
			},
			//Update and Read
			{
				Config: testAccAclAcl("acl", "list", name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list.0.element", "acl"),
					resource.TestCheckResourceAttrPair(resourceName, "list.0.acl", "bloxone_dns_acl.test", "id"),
				),
			},
			//Update and Read
			{
				Config: testAccAclTsigKey("acl", "list", name, "deny"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "list.0.access", "deny"),
					resource.TestCheckResourceAttr(resourceName, "list.0.element", "tsig_key"),
					resource.TestCheckResourceAttrPair(resourceName, "list.0.tsig_key.key", "bloxone_keys_tsig.test", "id"),
				),
			},
			//Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAclResource_Name(t *testing.T) {
	var resourceName = "bloxone_dns_acl.test_name"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccAclName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccAclResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dns_acl.test_tags"
	var v dns_config.ConfigACL
	var name = acctest.RandomNameWithPrefix("acl")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccAclTags(name, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccAclTags(name, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAclExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckAclExists(ctx context.Context, resourceName string, v *dns_config.ConfigACL) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AclAPI.
			AclRead(ctx, rs.Primary.ID).
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

func testAccCheckAclDestroy(ctx context.Context, v *dns_config.ConfigACL) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AclAPI.
			AclRead(ctx, *v.Id).
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

func testAccCheckAclDisappears(ctx context.Context, v *dns_config.ConfigACL) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.DNSConfigurationAPI.
			AclAPI.
			AclDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccAclBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test" {
    name = %q
}
`, name)
}

func testAccAclComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccAclName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test_name" {
    name = %q
}
`, name)
}

func testAccAclIP(objectName, aclFieldName, name, access, address string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_%[1]s" "test_%[2]s" {
    name = %[3]q
    %[2]s = [
		{
			access = %[4]q
			element = "ip"
			address = %[5]q
		}
]
}
`, objectName, aclFieldName, name, access, address)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+name, "acl-"+name), config}, "")
}

func testAccAclAny(objectName, aclFieldName, name, access string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_%[1]s" "test_%[2]s" {
    name = %[3]q
    %[2]s = [
		{
			access = %[4]q
			element = "any"
		}
]
}
`, objectName, aclFieldName, name, access)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+name, "acl-"+name), config}, "")
}

func testAccAclAcl(objectName, aclFieldName, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_%[1]s" "test_%[2]s" {
    name = %[3]q
    %[2]s = [
		{
			element = "acl"
			acl = bloxone_dns_acl.test.id
		}
]
}
`, objectName, aclFieldName, name)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+name, "acl-"+name), config}, "")
}

func testAccAclTsigKey(objectName, aclFieldName, name, access string) string {
	config := fmt.Sprintf(`
resource "bloxone_dns_%[1]s" "test_%[2]s" {
    name = %[3]q
    %[2]s = [
		{
			element = "tsig_key"
			access = %[4]q
			tsig_key = {
				key = bloxone_keys_tsig.test.id
			}
		}
]
}
`, objectName, aclFieldName, name, access)
	return strings.Join([]string{testAccBaseWithTsigAndAcl("tsig-"+name, "acl-"+name), config}, "")

}

func testAccAclTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dns_acl" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}

func testAccBaseWithTsigAndAcl(tsigKeyName, aclName string) string {
	return fmt.Sprintf(`
resource "bloxone_keys_tsig" "test" {
    name = %q
}

resource "bloxone_dns_acl" "test" {
	name = %q
	list = [
		{			
			element = "ip"
			access = "allow"
			address = "10.0.0.0/24"
		}
	]
}
`, tsigKeyName+".", aclName)
}
