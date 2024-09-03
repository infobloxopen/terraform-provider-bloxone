package ipamfederation_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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
				Config: testAccFederatedBlockBasicConfig("FEDERATED_REALM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "federated_realm", "FEDERATED_REALM_REPLACE_ME"),
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
				Config: testAccFederatedBlockBasicConfig("FEDERATED_REALM_REPLACE_ME"),
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
	var resourceName = "bloxone_federated_block.test_address"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockAddress("FEDERATED_REALM_REPLACE_ME", "ADDRESS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockAddress("FEDERATED_REALM_REPLACE_ME", "ADDRESS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", "ADDRESS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

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
	var resourceName = "bloxone_federated_block.test_cidr"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockCidr("FEDERATED_REALM_REPLACE_ME", "CIDR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "CIDR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockCidr("FEDERATED_REALM_REPLACE_ME", "CIDR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cidr", "CIDR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_comment"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockComment("FEDERATED_REALM_REPLACE_ME", "COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockComment("FEDERATED_REALM_REPLACE_ME", "COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_FederatedRealm(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_federated_realm"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockFederatedRealm("FEDERATED_REALM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realm", "FEDERATED_REALM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockFederatedRealm("FEDERATED_REALM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "federated_realm", "FEDERATED_REALM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedBlockResource_Name(t *testing.T) {
	var resourceName = "bloxone_federated_block.test_name"
	var v ipamfederation.FederatedBlock

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockName("FEDERATED_REALM_REPLACE_ME", "NAME_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockName("FEDERATED_REALM_REPLACE_ME", "NAME_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "NAME_UPDATE_REPLACE_ME"),
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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedBlockTags("FEDERATED_REALM_REPLACE_ME", "TAGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedBlockTags("FEDERATED_REALM_REPLACE_ME", "TAGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedBlockExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags", "TAGS_UPDATE_REPLACE_ME"),
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
		apiRes, _, err := acctest.BloxOneClient.IPAMFederation.
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
		_, httpRes, err := acctest.BloxOneClient.IPAMFederation.
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
		_, err := acctest.BloxOneClient.IPAMFederation.
			FederatedBlockAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFederatedBlockBasicConfig(federatedRealm string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test" {
    federated_realm = %q
}
`, federatedRealm)
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

func testAccFederatedBlockComment(federatedRealm string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_comment" {
    federated_realm = %q
    comment = %q
}
`, federatedRealm, comment)
}

func testAccFederatedBlockFederatedRealm(federatedRealm string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_federated_realm" {
    federated_realm = %q
}
`, federatedRealm)
}

func testAccFederatedBlockName(federatedRealm string, name string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_name" {
    federated_realm = %q
    name = %q
}
`, federatedRealm, name)
}

func testAccFederatedBlockParent(federatedRealm string, parent string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_parent" {
    federated_realm = %q
    parent = %q
}
`, federatedRealm, parent)
}

func testAccFederatedBlockTags(federatedRealm string, tags string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_block" "test_tags" {
    federated_realm = %q
    tags = %q
}
`, federatedRealm, tags)
}
