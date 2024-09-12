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

func TestAccFederatedRealmResource_basic(t *testing.T) {
	var resourceName = "bloxone_federated_realm.test"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedRealmBasicConfig("REALM_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "REALM_TEST"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedRealmResource_disappears(t *testing.T) {
	resourceName := "bloxone_federated_realm.test"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckFederatedRealmDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccFederatedRealmBasicConfig("REALM_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					testAccCheckFederatedRealmDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

// readonly fields
func TestAccFederatedRealmResource_AllocationV4(t *testing.T) {
	var resourceName = "bloxone_federated_realm.test_allocation_v4"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedRealmAllocationV4("NAME_REPLACE_ME", "ALLOCATION_V4_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allocation_v4", "ALLOCATION_V4_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedRealmAllocationV4("NAME_REPLACE_ME", "ALLOCATION_V4_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allocation_v4", "ALLOCATION_V4_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedRealmResource_Comment(t *testing.T) {
	var resourceName = "bloxone_federated_realm.test_comment"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedRealmComment("REALM_TEST", "COMMENT_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedRealmComment("REALM_TEST", "COMMENT_UPDATE_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_TEST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedRealmResource_Name(t *testing.T) {
	var resourceName = "bloxone_federated_realm.test_name"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedRealmName("REALM_TEST"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "REALM_TEST"),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedRealmName("REALM_TEST_UPDATE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "REALM_TEST_UPDATE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccFederatedRealmResource_Tags(t *testing.T) {
	var resourceName = "bloxone_federated_realm.test_tags"
	var v ipamfederation.FederatedRealm

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactoriesWithTags,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccFederatedRealmTags("REALM_TEST", map[string]string{
					"site": "NA",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.site", "NA"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.site", "NA"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Update and Read
			{
				Config: testAccFederatedRealmTags("NAME_REPLACE_ME", map[string]string{
					"site": "CA",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFederatedRealmExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.site", "CA"),
					resource.TestCheckResourceAttr(resourceName, "tags_all.site", "CA"),
					acctest.VerifyDefaultTag(resourceName),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckFederatedRealmExists(ctx context.Context, resourceName string, v *ipamfederation.FederatedRealm) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedRealmAPI.
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

func testAccCheckFederatedRealmDestroy(ctx context.Context, v *ipamfederation.FederatedRealm) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedRealmAPI.
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

func testAccCheckFederatedRealmDisappears(ctx context.Context, v *ipamfederation.FederatedRealm) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAMFederationAPI.
			FederatedRealmAPI.
			Delete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccFederatedRealmBasicConfig(name string) string {
	// TODO: create basic resource with required fields
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test" {
    name = %q
}
`, name)
}

func testAccFederatedRealmAllocationV4(name string, allocationV4 string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test_allocation_v4" {
    name = %q
    allocation_v4 = %q
}
`, name, allocationV4)
}

func testAccFederatedRealmComment(name string, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test_comment" {
    name = %q
    comment = %q
}
`, name, comment)
}

func testAccFederatedRealmName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test_name" {
    name = %q
}
`, name)
}

func testAccFederatedRealmTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"
	return fmt.Sprintf(`
resource "bloxone_federated_realm" "test_tags" {
    name = %q
    tags = %s
}
`, name, tagsStr)
}
