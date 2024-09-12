package ipamfederation_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccFederatedBlockResource_basic(t *testing.T) {
	var resourceName = "bloxone_federated_block.test"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, "FEDERATED_REALM_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federated_realm.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_disappears(t *testing.T) {
	resourceName := "bloxone_federated_block.test"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, "FEDERATED_REALM_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					testAccCheckFederatedBlockDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccFederatedBlockResource_Address(t *testing.T) {
	var resourceName = "bloxone_federated_block.test"
	var v1 ipamfederation.FederatedBlock
	var v2 ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.11.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockDestroy(context.Background(), &v1),
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", "10.11.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// read only fields
func TestAccFederatedBlockResource_AllocationV4(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_allocation_v4"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockAllocationV4("FEDERATED_REALM_REPLACE_ME", "ALLOCATION_V4_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allocation_v4", "ALLOCATION_V4_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockAllocationV4("FEDERATED_REALM_REPLACE_ME", "ALLOCATION_V4_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allocation_v4", "ALLOCATION_V4_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_federated_block.test"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockBasicConfig("10.10.0.0", 18, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.10.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "18"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_comment"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockComment("10.10.0.0", 16, realmName, "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockComment("10.10.0.0", 16, realmName, "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_FederatedRealm(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_federated_realm"
	var v ipamfederation.FederatedBlock
	realmName1 := acctest.RandomNameWithPrefix("federated-realm")
	realmName2 := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockFederatedRealm(realmName1, realmName2, "bloxone_federated_realm.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federated_realm.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockFederatedRealm(realmName1, realmName2, "bloxone_federated_realm.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federated_realm.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Name(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_name"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockName("10.0.0.0", 24, realmName, "NAME_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockName("10.0.0.0", 24, realmName, "NAME_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Parent(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_parent"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockParent("FEDERATED_REALM_REPLACE_ME", "PARENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parent", "PARENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockParent("FEDERATED_REALM_REPLACE_ME", "PARENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "parent", "PARENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_tags"
	var v ipamfederation.FederatedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockTags("10.0.0.0", realmName, 24, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockTags("10.0.0.0", realmName, 24, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag3", "value3"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFederatedBlockExists(ctx context.Context, resourceName string, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
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

func testAccCheckFederatedBlockDestroy(ctx context.Context, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
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

func testAccCheckFederatedBlockDisappears(ctx context.Context, v *ipamfederation.FederatedBlock) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedBlockAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccBaseWithFederatedRealm(name string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test" {
    name = %q
}
`, name)
}

func testAccFederatedBlockBasicConfig(address string, cidr int, federatedRealm string) string {
	// TODO: create basic resource with required fields
	config := fmt.Sprintf(`
resource "bloxone_federated_block" "test" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federated_realm.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockAddress(federatedRealm string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_address" {
    federated_realm = %q
    address = %q
}
`, federatedRealm, address)
}

func testAccFederatedBlockAllocationV4(federatedRealm string, allocationV4 string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_allocation_v4" {
    federated_realm = %q
    allocation_v4 = %q
}
`, federatedRealm, allocationV4)
}

func testAccFederatedBlockCidr(federatedRealm string, cidr string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_cidr" {
    federated_realm = %q
    cidr = %q
}
`, federatedRealm, cidr)
}

func testAccFederatedBlockComment(address string, cidr int, federatedRealm string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_federated_block" "test_comment" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federated_realm.test.id
    comment = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockFederatedRealm(federatedRealm1, federatedRealm2, realm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federated_block" "test_federated_realm" {
   address = "10.0.0.0"
    cidr = 16
    federated_realm =%s.id
}
`, realm)
	return strings.Join([]string{testAccBaseWithTwoFederatedRealm(federatedRealm1, federatedRealm2), config}, "")
}

func testAccBaseWithTwoFederatedRealm(name1, name2 string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "one" {
	name = %q
}
resource "bloxone_federated_realm" "two" {
	name = %q
}`, name1, name2)
}

func testAccFederatedBlockName(address string, cidr int, federatedRealm string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federated_block" "test_name" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federated_realm.test.id
    name = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccFederatedBlockParent(federatedRealm string, parent string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_parent" {
    federated_realm = %q
    parent = %q
}
`, federatedRealm, parent)
}

func testAccFederatedBlockTags(address string, federatedRealm string, cidr int, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_federated_block" "test_tags" {
    address = %q
    federated_realm = bloxone_federated_realm.test.id
    cidr = %d
    tags = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}
