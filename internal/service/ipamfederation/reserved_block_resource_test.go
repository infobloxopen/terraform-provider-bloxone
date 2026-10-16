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

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

func TestAccReservedBlockResource_basic(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockBasicConfig(realmName, "10.90.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.90.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_reserved_block.test"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckReservedBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccReservedBlockBasicConfig(realmName, "10.91.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					testAccCheckReservedBlockDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccReservedBlockResource_Address(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_address"
	var v1 ipamfederation.ReservedBlock
	var v2 ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockAddress(realmName, "10.92.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", "10.92.0.0"),
				),
			},
			// Update and Read (address is immutable, so this forces a replace)
			{
				Config: testAccReservedBlockAddress(realmName, "10.93.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockDestroy(context.Background(), &v1),
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", "10.93.0.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_cidr"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockCidr(realmName, "10.94.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "24"),
				),
			},
			// Update and Read
			{
				Config: testAccReservedBlockCidr(realmName, "10.94.0.0", 25),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "25"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_comment"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockComment(realmName, "10.95.0.0", 24, "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccReservedBlockComment(realmName, "10.95.0.0", 24, "COMMENT_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_FederatedPoolId(t *testing.T) {
	t.Skip("FederatedPoolId requires a valid federated pool resource ID; to be added when a pool fixture is available")
}

func TestAccReservedBlockResource_FederatedRealm(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_federated_realm"
	var v ipamfederation.ReservedBlock
	realmName1 := acctest.RandomNameWithPrefix("federated-realm")
	realmName2 := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.one", "10.96.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccReservedBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.two", "10.96.0.0", 24),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_name"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockName(realmName, "10.98.0.0", 24, "NAME_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccReservedBlockName(realmName, "10.98.0.0", 24, "NAME_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccReservedBlockResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_reserved_block.test_tags"
	var v ipamfederation.ReservedBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccReservedBlockTags(realmName, "10.101.0.0", 24, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.tag2", "value2"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccReservedBlockTags(realmName, "10.101.0.0", 24, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckReservedBlockExists(context.Background(), resourceName, &v),
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

func testAccCheckReservedBlockExists(ctx context.Context, resourceName string, v *ipamfederation.ReservedBlock) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			ReservedBlockAPI.
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

func testAccCheckReservedBlockDestroy(ctx context.Context, v *ipamfederation.ReservedBlock) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			ReservedBlockAPI.
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

func testAccCheckReservedBlockDisappears(ctx context.Context, v *ipamfederation.ReservedBlock) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			ReservedBlockAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccReservedBlockBasicConfig(federatedRealm string, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockAddress(federatedRealm string, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_address" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockCidr(federatedRealm string, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_cidr" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockComment(federatedRealm string, address string, cidr int, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_comment" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
    comment = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockFederatedRealm(federatedRealm1, federatedRealm2, realm string, address string, cidr int) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_federated_realm" {
    federated_realm = %s.id
    address = %q
    cidr = %d
}
`, realm, address, cidr)
	return strings.Join([]string{testAccBaseWithTwoFederatedRealm(federatedRealm1, federatedRealm2), config}, "")
}

func testAccReservedBlockName(federatedRealm string, address string, cidr int, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_name" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
    name = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccReservedBlockTags(federatedRealm string, address string, cidr int, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_federation_reserved_block" "test_tags" {
    federated_realm = bloxone_federation_federated_realm.test.id
    address = %q
    cidr = %d
    tags = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}
