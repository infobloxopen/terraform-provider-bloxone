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

// TODO:
// - Federated Pool

func TestAccOverlappingBlockResource_basic(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockBasicConfig("10.20.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "address", "10.20.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.test", "id"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "network_compliant"),
					resource.TestCheckResourceAttrSet(resourceName, "protocol"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_disappears(t *testing.T) {
	resourceName := "bloxone_federation_overlapping_block.test"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOverlappingBlockDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOverlappingBlockBasicConfig("10.20.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					testAccCheckOverlappingBlockDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOverlappingBlockResource_Address(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_address"
	var v1 ipamfederation.OverlappingBlock
	var v2 ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockAddress("10.21.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, "address", "10.21.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockAddress("10.22.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockDestroy(context.Background(), &v1),
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v2),
					resource.TestCheckResourceAttr(resourceName, "address", "10.22.0.0"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_Cidr(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_cidr"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockCidr("10.23.0.0", 16, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.23.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "16"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockCidr("10.23.0.0", 18, realmName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "10.23.0.0"),
					resource.TestCheckResourceAttr(resourceName, "cidr", "18"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_comment"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockComment("10.24.0.0", 16, realmName, "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockComment("10.24.0.0", 16, realmName, "COMMENT_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_FederatedPoolId(t *testing.T) {
	t.Skip("TODO: Federated Pool support is yet to be added")

	var resourceName = "bloxone_federation_overlapping_block.test_federated_pool_id"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockFederatedPoolId("10.25.0.0", 16, realmName, "FEDERATED_POOL_ID_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_pool_id", "FEDERATED_POOL_ID_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockFederatedPoolId("10.25.0.0", 16, realmName, "FEDERATED_POOL_ID_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_pool_id", "FEDERATED_POOL_ID_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_FederatedRealm(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_federated_realm"
	var v ipamfederation.OverlappingBlock
	realmName1 := acctest.RandomNameWithPrefix("federated-realm")
	realmName2 := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.one"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.one", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockFederatedRealm(realmName1, realmName2, "bloxone_federation_federated_realm.two"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttrPair(resourceName, "federated_realm", "bloxone_federation_federated_realm.two", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_Name(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_name"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockName("10.26.0.0", 16, realmName, "NAME_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockName("10.26.0.0", 16, realmName, "NAME_TEST_UPDATED"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_TEST_UPDATED"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOverlappingBlockResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federation_overlapping_block.test_tags"
	var v ipamfederation.OverlappingBlock
	realmName := acctest.RandomNameWithPrefix("federated-realm")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOverlappingBlockTags("10.27.0.0", 16, realmName, map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccOverlappingBlockTags("10.27.0.0", 16, realmName, map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOverlappingBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOverlappingBlockExists(ctx context.Context, resourceName string, v *ipamfederation.OverlappingBlock) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			OverlappingBlockAPI.
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

func testAccCheckOverlappingBlockDestroy(ctx context.Context, v *ipamfederation.OverlappingBlock) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			OverlappingBlockAPI.
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

func testAccCheckOverlappingBlockDisappears(ctx context.Context, v *ipamfederation.OverlappingBlock) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			OverlappingBlockAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOverlappingBlockBasicConfig(address string, cidr int, federatedRealm string) string {
	// TODO: create basic resource with required fields
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockAddress(address string, cidr int, federatedRealm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_address" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockCidr(address string, cidr int, federatedRealm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_cidr" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
}
`, address, cidr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockComment(address string, cidr int, federatedRealm string, comment string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_comment" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    comment = %q
}
`, address, cidr, comment)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockFederatedPoolId(address string, cidr int, federatedRealm string, federatedPoolId string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_federated_pool_id" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    federated_pool_id = %q
}
`, address, cidr, federatedPoolId)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockFederatedRealm(federatedRealm1, federatedRealm2, realm string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_federated_realm" {
   address = "10.28.0.0"
   cidr = 16
   federated_realm = %s.id
}
`, realm)
	return strings.Join([]string{testAccBaseWithTwoFederatedRealm(federatedRealm1, federatedRealm2), config}, "")
}

func testAccOverlappingBlockName(address string, cidr int, federatedRealm string, name string) string {
	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_name" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    name = %q
}
`, address, cidr, name)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}

func testAccOverlappingBlockTags(address string, cidr int, federatedRealm string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	config := fmt.Sprintf(`
resource "bloxone_federation_overlapping_block" "test_tags" {
    address = %q
    cidr = %d
    federated_realm = bloxone_federation_federated_realm.test.id
    tags = %s
}
`, address, cidr, tagsStr)
	return strings.Join([]string{testAccBaseWithFederatedRealm(federatedRealm), config}, "")
}
